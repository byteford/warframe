package inventory

import (
	"fmt"
	"strings"

	"github.com/byteford/warframe/print"
)

type Item struct {
	Name     string   `json:"name"`
	Amount   int      `json:"amount"`
	Crafting Crafting `json:"crafting"`
}

type Items []Item

func ItemFromList(items Items, name string) (Item, error) {
	for _, v := range items {
		if strings.EqualFold(v.Name, name) {
			return v, nil
		}
	}
	return Item{}, fmt.Errorf("item: %s, not found", name)
}

func LoadInBase(items Items) (Items, error) {
	mats := Materials{}
	for _, v := range items {
		mats = append(mats, Material{Name: v.Name, Amount: 0})
		m, err := v.Crafting.GetBaseMaterials(items)
		if err != nil {
			if err.Error() != enderr {
				return Items{}, err
			}
			continue
		}
		mats = append(mats, m...)
	}
	print.Printf("ALL\n%+v\n", mats)
	uMats, err := mats.Unique()
	if err != nil {
		return Items{}, err
	}
	print.Printf("Unique\n%+v\n", uMats)
	for _, v := range uMats {
		_, err = ItemFromList(items, v.Name)
		if err == nil {
			continue
		}
		items = append(items, Item{Name: v.Name})
	}
	return items, nil
}
