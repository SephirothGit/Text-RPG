package character

import (
	"fmt"
	"github.com/SephirothGit/Text-RPG/pkg/core"
)

// Structure for Paladin
type Paladin struct {
	core.BaseCharacter
}

// Fight overridden method for Paladin
func (p *Paladin) Fight() {
	fmt.Printf("%s (Paladin) hits the enemy with a holy sword for %d DMG!\n", p.Name, p.DMG)
}