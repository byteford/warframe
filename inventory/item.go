package inventory

import (
	"fmt"
	"strings"
)

type Item struct {
	Name     string   `json:"name"`
	Amount   int      `json:"amount"`
	Crafting Crafting `json:"crafting"`
}

func ItemFromList(items []Item, name string) (Item, error) {
	for _, v := range items {
		if strings.EqualFold(v.Name, name) {
			return v, nil
		}
	}
	return Item{}, fmt.Errorf("item: %s, not found", name)
}
