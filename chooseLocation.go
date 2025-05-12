package location

import (
	"fmt"
	"strings"
	"github.com/SephirothGit/Text-RPG/pkg/core"

)

// Choose location func
func (c *core.BaseCharacter) ChooseLocation() {

	for {
		var location string
		fmt.Print("Where do you want to go? (Forest/Town/Tavern/Exit)")
		_, err := fmt.Scan(&location)
		if err != nil {
			fmt.Println("input error when choose a place to go: %v", err)
			return
		}

		switch strings.ToLower(location) {
		case "forest":
			location.GoToForest()
			return
		case "town":
			c.GoToTown()
			return
		case "tavern":
			c.GoToTavern()
			return
		case "exit":
			fmt.Print("Goodbye")
			return
		default:
			fmt.Print("Unknown location. Use fast travel to a familiar place\n")
			c.ChooseLocation()
		}
	}
}
