package characters

import "fmt"

// Structure for Knight
type Knight struct {
	BaseCharacter
}

// Fight overridden method for Warrior
func (k *Knight) Fight() {
	fmt.Printf("%s (Knight) hits the enemy with a sword for %d DMG!\n", k.Name, k.DMG)
}