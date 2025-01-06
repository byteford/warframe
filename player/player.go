package player

import "github.com/byteford/warframe/inventory"

type Player struct {
	Name      string          `json:"name"`
	Inventory inventory.Items `json:"inventory"`
}
