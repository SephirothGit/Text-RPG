package character

import (
	"fmt"
	"github.com/SephirothGit/Text-RPG/pkg/core"
)

// Structure for Knight
type Knight struct {
	core.BaseCharacter
}

// Fight overridden method for Warrior
func (k *core.Knight) Fight() {
	fmt.Printf("%s (Knight) hits the enemy with a sword for %d DMG!\n", k.Name, k.DMG)
}
