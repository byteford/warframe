package player

import (
	"strings"

	"github.com/byteford/warframe/inventory"
)

type Player struct {
	Name      string          `json:"name"`
	Inventory inventory.Items `json:"inventory"`
	Plan      Plan            `json:"plan"`
}

type Plan struct {
	Craft inventory.Materials `json:"craft"`
}

func Init(name string) (Player, error) {
	return Player{Name: name, Inventory: inventory.Items{}}, nil
}

func (p *Player) AddCraft(name string, amount int) {
	if p.hasCraft(name) {
		return
	}
	p.Plan.Craft = append(p.Plan.Craft, inventory.Material{Name: name, Amount: amount})
}

func (p *Player) DeleteCraft(name string, amount int) {
	index := p.CraftIndex(name)
	if index == -1 {
		return
	}
	p.Plan.Craft = append(p.Plan.Craft[:index], p.Plan.Craft[index+1:]...)
}

func (p *Player) CraftIndex(name string) int {
	for i, v := range p.Plan.Craft {
		if strings.EqualFold(v.Name, name) {
			return i
		}
	}
	return -1
}

func (p *Player) hasCraft(name string) bool {
	for _, v := range p.Plan.Craft {
		if strings.EqualFold(v.Name, name) {
			return true
		}
	}
	return false
}
