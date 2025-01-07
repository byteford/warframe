package inventory

import (
	"fmt"
	"strings"

	"github.com/byteford/warframe/print"
)

type Item struct {
	Name         string   `json:"name"`
	Amount       int      `json:"amount,omitempty"`
	Crafting     Crafting `json:"crafting,omitempty"`
	FarmingNotes string   `json:"farmingNotes,omitempty"`
}

type Items []Item

func ItemIndexFromList(items Items, name string) int {
	for i, v := range items {
		if strings.EqualFold(v.Name, name) {
			return i
		}
	}
	return -1
}

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

func (items Items) UpdateItem(name string, amount int) (Items, error) {
	i := items
	index := ItemIndexFromList(items, name)
	if index == -1 {
		i = append(items, Item{Name: name, Amount: amount})
		return i, nil
	}
	items[index].Amount = amount
	return i, nil
}
