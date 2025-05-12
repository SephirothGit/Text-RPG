package location

import (
	"fmt"
	"strings"
	"github.com/SephirothGit/Text-RPG/pkg/core"
)

// Method for going to the Town
func (c *core.BaseCharacter) GoToTown() {
	fmt.Printf("%s goes to the Town...\n", c.Name)

	if c.LichDragonDead {
		fmt.Println("Guard: Welcome to the Fawn, our hero!")
	}

	for {
		//Choose a place to visit in town
		var choosePlace string

		fmt.Print("Which place you want to visit? (Dungeon/Inn/Tavern/Exit)")
		_, err := fmt.Scan(&choosePlace)
		if err != nil {
			fmt.Println("Input error when choose a place to visit in town: %v", err)
			continue
		}

		switch strings.ToLower(choosePlace) {
		case "dungeon":
			core.c.GoToDungeon()
			return

		case "inn":
			fmt.Printf("%s goes to the Mermaid's inn...\n", c.Name)
			c.GoToInn()
			return

		case "tavern":
			fmt.Printf("%s goes to the Brick's tavern...\n", c.Name)
			c.GoToTavernB()
			return

		case "exit":
			fmt.Print("You came out\n")
			c.ChooseLocation()
			return

		default:
			fmt.Print("Unknown location. Choose from familiar places\n")
		}
	}
}
