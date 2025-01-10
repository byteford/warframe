package inventory

import (
	"fmt"
	"sort"
	"strings"
)

type Crafting struct {
	Blueprint     Blueprint `json:"blueprint,omitempty"`
	Materials     Materials `json:"materials,omitempty"`
	BaseMaterials Materials `json:"BaseMaterials,omitempty"`
}

type Blueprint struct {
	Name string `json:"name,omitempty"`
	Have bool   `json:"have,omitempty"`
}

type Material struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type Materials []Material

const enderr = "end of tree"

func (m Materials) Sort() Materials {
	sort.Slice(m, func(i, j int) bool {
		return m[i].Name < m[j].Name
	})
	return m
}

func (c Crafting) GetBaseMaterials(items Items) (Materials, error) {
	var mats Materials

	if len(c.Materials) == 0 {
		return Materials{}, fmt.Errorf("%s", enderr)
	}
	for _, v := range c.Materials {
		item, err := ItemFromList(items, v.Name)
		if err != nil {
			mats = append(mats, v)
		}
		res, err := item.Crafting.GetBaseMaterials(items)
		if err != nil {
			if err.Error() != enderr {
				return Materials{}, err
			}
			mats = append(mats, v)
		}
		mats = append(mats, res...)
	}
	materials, err := mats.Unique()
	if err != nil {
		return Materials{}, err
	}
	return materials, nil
}

func (m Materials) Unique() (Materials, error) {
	tmp := map[string]int{}
	for _, v := range m {
		if _, ok := tmp[v.Name]; !ok {
			tmp[v.Name] = 0
		}
		tmp[v.Name] = tmp[v.Name] + v.Amount
	}
	var mats Materials
	for k, v := range tmp {
		mats = append(mats, Material{Name: k, Amount: v})
	}
	return mats, nil
}

func (m Materials) Required(have Items) (Materials, error) {
	mats := Materials{}
	for _, v := range m {
		have_item, err := ItemFromList(have, v.Name)
		if err != nil {
			if !strings.Contains(err.Error(), "not found") {
				return Materials{}, err
			}
		} else {
			if have_item.Amount >= v.Amount {
				continue
			}
		}
		mats = append(mats, v)
	}
	return mats, nil
}
