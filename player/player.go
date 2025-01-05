package player

import "github.com/byteford/warframe/inventory"

type Player struct {
	Inventory []inventory.Item `json:"inventory"`
}
