package exercise

import (
	"fmt"
	"slices"
)

func main() {

}

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name      string
	Inventory []Item
}

func (p *Player) addItemToInventory(item Item) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("Picked up %s\n", item.Name)
}

func (p *Player) dropItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = slices.Delete(p.Inventory, i, i+1)
			fmt.Printf("%s dropped %s\n", p.Name, itemName)
			return
		}
	}
	fmt.Printf("%s does not have item %s in their inventory", p.Name, itemName)
}

func (p *Player) useItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			if item.Type == "potion" {
				p.Inventory = slices.Delete(p.Inventory, i, i+1)
				fmt.Printf("%s used %s and feel rejuvenated\n", p.Name, itemName)
			} else {
				fmt.Printf("%s used %s\n", p.Name, itemName)
			}
			return
		}
	}
	fmt.Printf("%s does not have %s in their inventory", p.Name, itemName)
}
