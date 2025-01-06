package db

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/byteford/warframe/inventory"
	"github.com/byteford/warframe/player"
)

func LoadItems(name string) (inventory.Items, error) {
	res, err := os.ReadFile(fmt.Sprintf("%s.json", name))
	if err != nil {
		return nil, err
	}

	var items inventory.Items
	err = json.Unmarshal(res, &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func SaveItems(name string, items inventory.Items) error {
	jsonOutput, err := json.Marshal(items)
	if err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s.json", name), jsonOutput, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadPlayer(name string) (player.Player, error) {
	res, err := os.ReadFile(fmt.Sprintf("%s.json", name))
	if err != nil {
		return player.Player{}, err
	}

	var p player.Player
	err = json.Unmarshal(res, &p)
	if err != nil {
		return player.Player{}, err
	}
	return p, nil
}

func SavePlayer(name string, p player.Player) error {
	jsonOutput, err := json.Marshal(p)
	if err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s.json", name), jsonOutput, 0644)
	if err != nil {
		return err
	}
	return nil
}
