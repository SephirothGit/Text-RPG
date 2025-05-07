package character

import "fmt"

// Structure for Priest
type Priest struct {
	BaseCharacter
}

// Fight overridden method for Priest
func (r *Priest) Fight() {
	fmt.Printf("%s (Priest) hit the enemy with a holy hammer for %d DMG!\n", r.Name, r.DMG)
}