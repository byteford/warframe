package db

import (
	"encoding/json"
	"os"

	inventory "github.com/byteford/warframe/inventory"
	"github.com/byteford/warframe/player"
)

func LoadItems(name string) ([]inventory.Item, error) {

	res, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var items []inventory.Item
	err = json.Unmarshal(res, &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func LoadPlayer(name string) (player.Player, error) {

	res, err := os.ReadFile(name)
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
