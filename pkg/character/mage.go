package character

import (
	"fmt"
	"github.com/SephirothGit/Text-RPG/pkg/core"
)

// Structure for Mage
type Mage struct {
	core.BaseCharacter
}

// Fight overridden method for Mage
func (m *Mage) Fight() {
	fmt.Printf("%s (Mage) cast a fireball for %d DMG!\n", m.Name, m.DMG)
}