package location

import (
	"fmt"
	"strings"

	"github.com/SephirothGit/Text-RPG/internal/character"
)

// Method for going to the Forbidden forest
func (c *character.BaseCharacter) GoToForest() {

	fmt.Printf("%s goes to the Forbidden forest...\n", c.Name)

	//Improved fight system
	HobgoblinHP := 120
	fmt.Printf("A Hobgoblin (HP: %d) has attacked you!\n", HobgoblinHP)

	for {
		fmt.Printf("%s HP: %d/%d | Hobgoblin HP %d\n", c.Name, c.HP, c.MaxHP, HobgoblinHP)
		fmt.Print("What will you do? (Attack/Run)")

		var action string
		_, err := fmt.Scan(&action)
		if err != nil {
			fmt.Println("Input error after meeting HobGoblin: %v", err)
		}

		switch strings.ToLower(action) {
		case "attack":
			c.Fight()

			fmt.Printf("%s attacks the Hobgoblin for %d DMG!\n", c.Name, c.DMG)
			HobgoblinHP -= c.DMG

			if HobgoblinHP <= 0 {
				fmt.Println("You defeated the Hobgoblin!\n")
				c.ChooseLocation()
				return
			}
			//Monster fights
			HobgoblinDMG := 15
			c.HP -= HobgoblinDMG
			fmt.Printf("Hobgoblin hits you with a club for %d damage!\n", HobgoblinDMG)

			if c.HP <= 0 {
				fmt.Println("You died")
				return
			}

		case "run":
			fmt.Print("You run away from the Hobgoblin!\n")
			c.ChooseLocation()
			return

		default:
			fmt.Println("Unknown command")
		}
	}
}
