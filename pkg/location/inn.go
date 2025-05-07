package location

import (
	"fmt"
	"strings"
	"github.com/SephirothGit/Text-RPG/pkg/character"
)

func (c *BaseCharacter) GoToInn() {
	fmt.Print("Welcome to the Mermaid's Inn!\n")

	if c.LichDragonDead {
		fmt.Print("Freya: I'm glad to see you, hero of Fawn!")
	}

	for {
		var sleepAction string
		fmt.Print("Freya: We offer best rooms in town! Do you want to stay in Fawn for a night? (Yes/No/Exit)")

		_, err := fmt.Scan(&sleepAction)
		if err != nil {
			fmt.Println("Input error sleep action: %v", err)
			continue
		}

		// Choose for how long you want to stay in Fawn
		sleepAction = strings.ToLower(sleepAction)

		switch sleepAction {

		case "yes":
			fmt.Print("Freya: It will cost you 1 silver coin\n")

			if c.Silver >= 1 {
				c.Silver -= 1
				fmt.Print("1 silver coin was given away\n")
				c.Sleep()
			} else {
				fmt.Println("Not enough silver")
			}
			continue

		case "no":
			fmt.Print("Ok, maybe another time\n")
			break

		case "exit":
			fmt.Print("You came out from the Mermaid's inn")
			c.GoToTown()
			return

		default:
			fmt.Println("Unknown action")
			continue
		}

		for {
			var innAction string
			fmt.Printf("Freya: %s, do you want to eat and drink, or maybe you want to hear rumors? (Drink/Eat/Rumors/Exit)", c.Name)

			_, err = fmt.Scan(&innAction)
			if err != nil {
				fmt.Println("Input error inn action: %v", err)
				continue
			}

			innAction = strings.ToLower(innAction)

			switch innAction {
			case "drink":
				c.Drink()

			case "eat":
				c.Eat()

			case "rumors":
				fmt.Print("Freya: You want to hear the rumors, ok it will cost 2 silver coins\n")
				if c.Silver >= 2 {
					c.Silver -= 2
					fmt.Print("2 silver coins were given away\n")

					fmt.Print("Freya: Some strangers that were here two days ago asked travelers about man with a sword named Steel Rose. I heard that he stole this item from the head of a royal guard of Shangri-La. A reward of 20 gold coins has been offered for the return of this sword. If you fast and smart enough to overtake those men, you can get this reward.\n")
				} else {
					fmt.Println("Not enough silver")
				}

			case "exit":
				fmt.Print("Freya: Come back whenever you want to stay in town, traveler!\n")
				c.ChooseLocation()
				return

			default:
				fmt.Println("Unknown action")
				continue
			}
		}
	}
}
