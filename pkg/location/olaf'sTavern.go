package location

import (
	"fmt"
	"strings"
	"github.com/SephirothGit/Text-RPG/pkg/character"
)

// Method for going to the Olaf's tavern
func (c *BaseCharacter) GoToTavern() {
	fmt.Printf("%s goes to the Olaf's tavern...\n", c.Name)

	if c.LichDragonDead {
		fmt.Println("Aren't you a hero who defeated the Fartissax?! Because of his miasma, monsters from the Forbidden Forest often attacked our city walls. It is an honor to serve you, what would you like today?")
	}

	//Loop for actions
	for {

		var tavernAction string

		fmt.Print("Olaf: Hey, traveler! Do you want to drink a cold beer or a room for a night? (Eat/Drink/Sleep/Exit)")

		//Error handling
		_, err := fmt.Scan(&tavernAction)
		if err != nil {
			fmt.Println("Input error when choose a tavern action: %v", err)
			continue
		}

		switch strings.ToLower(tavernAction) {
		case "drink":
			fmt.Printf("Olaf: What would you like to drink, %s?\n", c.Name)
			c.Drink()
		case "sleep":
			fmt.Print("Olaf: your room is first door on the right side of second floor\n")
			c.Sleep()
		case "eat":
			fmt.Printf("Waitress: What would you like to eat today, %s?\n", c.Name)
			c.Eat()
		case "exit":
			fmt.Print("Olaf: Goodbye, traveler!\n")
			c.ChooseLocation()
			return
		default:
			fmt.Print("Olaf: We don't offer this option\n")
		}
	}
}
