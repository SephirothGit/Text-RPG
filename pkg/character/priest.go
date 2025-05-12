package character

import (
	"fmt"
	"github.com/SephirothGit/Text-RPG/pkg/core"
)

// Structure for Priest
type Priest struct {
	core.BaseCharacter
}

// Fight overridden method for Priest
func (r *Priest) Fight() {
	fmt.Printf("%s (Priest) hit the enemy with a holy hammer for %d DMG!\n", r.Name, r.DMG)
}