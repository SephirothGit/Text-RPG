package location

import (
	"fmt"
	"strings"
	"github.com/SephirothGit/Text-RPG/internal/character"
)

// Go to Brick's tavern
func (c *character.BaseCharacter) GoToTavernB() {
	fmt.Printf("Waitress: Welcome to the Brick's Tavern %s!\n", c.Name)

	if c.LichDragonDead {
		fmt.Println("Waitress: I am happy to see dragon slayer in our tavern!")
	}

	for {
		var tavernActionB string

		fmt.Print("Brick: Good day, traveler! Do you want to drink a glass of cold beer with a snacks? (Drink/Eat/Exit)")
		_, err := fmt.Scan(&tavernActionB)
		if err != nil {
			fmt.Println("input err when choose a Brick's tavern action: %v", err)
			continue
		}

		switch strings.ToLower(tavernActionB) {
		case "drink":
			fmt.Printf("Brick: What do you want to drink, %s?\n", c.Name)
			c.Drink()
		case "eat":
			fmt.Printf("Brick: What would you like to eat, %s\n", c.Name)
			c.Eat()
		case "exit":
			fmt.Print("Brick: Come back anytime!\n")
			c.ChooseLocation()
			return
		default:
			fmt.Print("Unknown action. Choose from (Drink/Eat/Exit)\n")
		}
	}
}