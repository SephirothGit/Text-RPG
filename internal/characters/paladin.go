package characters

import "fmt"

// Structure for Paladin
type Paladin struct {
	BaseCharacter
}

// Fight overridden method for Paladin
func (p *Paladin) Fight() {
	fmt.Printf("%s (Paladin) hits the enemy with a holy sword for %d DMG!\n", p.Name, p.DMG)
}