package characters

import "fmt"

// Structure for Mage
type Mage struct {
	BaseCharacter
}

// Fight overridden method for Mage
func (m *Mage) Fight() {
	fmt.Printf("%s (Mage) cast a fireball for %d DMG!\n", m.Name, m.DMG)
}