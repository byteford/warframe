package inventory

import (
	"fmt"
)

type Crafting struct {
	Blueprint     Blueprint `json:"blueprint,omitempty"`
	Materials     Materials `json:"materials,omitempty"`
	BaseMaterials Materials
}

type Blueprint struct {
	Name string `json:"name"`
}

type Material struct {
	Name   string `json:"name"`
	Amount int    `json:"amount,omitempty"`
}

type Materials []Material

const enderr = "end of tree"

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
