package main

import (
	"fmt"
	"strings"
)

// Character interface that defines methods for all characters
type Character interface {
	Fight()
	Sleep()
	Drink()
	Eat()
	ChooseLocation()
}

// Choose location func
func (c *BaseCharacter) ChooseLocation() {

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
			c.GoToForest()
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

// Method for going to the Town
func (c *BaseCharacter) GoToTown() {
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
			c.GoToDungeon()
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

// Go to Brick's tavern
func (c *BaseCharacter) GoToTavernB() {
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

// Go to the Ancient Elves Dungeon
func (c *BaseCharacter) GoToDungeon() {

	if c.LichDragonDead {
		fmt.Println("The dungeon feels quieter now, when the LichDragon is gone...")
	}

	for {
		var DungeonAction string

		fmt.Printf("%s stands at the magic gate that able to teleport between floors of the dungeon...\n", c.Name)

		fmt.Print("Choose floor: (First/Second/Third/Exit)")
		_, err := fmt.Scan(&DungeonAction)
		if err != nil {
			fmt.Println("Input error when choose a floor in dungeon: %v", err)
		}

		switch strings.ToLower(DungeonAction) {

		// First floor
		case "first":

			// If Skeleton was killed
			if c.SkeletonDead {
				fmt.Println("The skeleton has already been slain! There's nothing more for you on this floor...")
				continue
			}

			fmt.Printf("%s went down to the 1 floor of the Ancient Elves Dungeon\n", c.Name)

			// Fight with a skeleton
			skeletonHP := 65
			fmt.Printf("Skeleton (HP: %d) has attacked you!\n", skeletonHP)

			for {
				fmt.Printf("%s HP: %d/%d | skeleton HP: %d\n", c.Name, c.HP, c.MaxHP, skeletonHP)
				fmt.Print("What will you do (Attack/Run)")

				var firstFloorAction string
				fmt.Scan(&firstFloorAction)

				switch strings.ToLower(firstFloorAction) {
				case "attack":
					fmt.Printf("%s attacks the skeleton for %d DMG\n", c.Name, c.DMG)
					skeletonHP -= c.DMG

					if skeletonHP <= 0 {
						c.SkeletonDead = true
						fmt.Println("You defeated the skeleton!")
						c.Silver += 3
						fmt.Println("You found 3 silver coins in the skeleton's bones")
						break
					}

					// Monster attacks
					skeletonDMG := 8
					fmt.Printf("Skeleton attacks you with a short sword for %d DMG!\n", skeletonDMG)
					c.HP -= skeletonDMG

					if c.HP <= 0 {
						fmt.Println("You died!")
						return
					}

				case "run":
					fmt.Print("You run away from the skeleton...\n")
					c.ChooseLocation()
					return

				default:
					fmt.Println("Unknown action")
					continue
				}

				if skeletonHP <= 0 {
					break
				}
			}

			// Second floor
		case "second":

			// If Salamander was killed
			if c.SalamanderDead {
				fmt.Println("The Salamander has already been slain! There's nothing more for you on this floor...")
				continue
			}

			fmt.Printf("%s went to the middle floor of the Ancient Elves Dungeon\n", c.Name)

			// Fight with a salamander
			salamanderHP := 110
			fmt.Printf("Salamander (HP: %d) has attacked you!\n", salamanderHP)

			for {
				fmt.Printf("%s HP: %d/%d | salamander HP: %d\n", c.Name, c.HP, c.MaxHP, salamanderHP)
				fmt.Print("What will you do (Attack/Run)")

				var secondFloorAction string
				fmt.Scan(&secondFloorAction)

				switch strings.ToLower(secondFloorAction) {
				case "attack":
					fmt.Printf("%s attacks the salamander for %d DMG\n", c.Name, c.DMG)
					salamanderHP -= c.DMG

					if salamanderHP <= 0 {
						c.SalamanderDead = true
						fmt.Println("You defeated the salamander!")
						c.Gold += 1
						c.Silver += 7
						fmt.Println("You found 1 gold and 7 silver coins")
						break
					}

					// Monster attacks
					salamanderDMG := 15
					fmt.Printf("Salamander attacks you with a fire ball for %d DMG!\n", salamanderDMG)
					c.HP -= salamanderDMG

					if c.HP <= 0 {
						fmt.Println("You died!")
						return
					}

				case "run":
					fmt.Print("You run away from the salamander...\n")
					c.ChooseLocation()
					return

				default:
					fmt.Println("Unknown action")
					continue
				}
				if salamanderHP <= 0 {
					break
				}
			}

			// Third floor
		case "third":

			// If Fortissax was killed
			if c.LichDragonDead {
				fmt.Println("The LichDragon Fortissax has already been slain! There's nothing more for you on this floor...")
				continue
			}

			fmt.Printf("%s went to the lower floor of the Ancient Elves Dungeon\n", c.Name)

			// Fight with a LichDradon Fortissax
			LichDradonFortissaxHP := 180
			fmt.Printf("A LichDragon Fortissax (HP: %d) has attacked you!\n", LichDradonFortissaxHP)

			for {
				fmt.Printf("%s HP: %d/%d | LichDragon Fortisax HP: %d\n", c.Name, c.HP, c.MaxHP, LichDradonFortissaxHP)
				fmt.Print("What will you do (Attack/Run)")

				var thirdFloorAction string
				fmt.Scan(&thirdFloorAction)

				switch strings.ToLower(thirdFloorAction) {

				case "attack":
					fmt.Printf("%s attacks the LichDragon Fortissax for %d DMG\n", c.Name, c.DMG)
					LichDradonFortissaxHP -= c.DMG

					if LichDradonFortissaxHP <= 0 {
						break
					}

					// Monster attacks
					LichDradonFortissaxDMG := 30
					fmt.Printf("LichDragon Fortissax attacks you with a lightning spear for %d DMG!\n", LichDradonFortissaxDMG)
					c.HP -= LichDradonFortissaxDMG

					if c.HP <= 0 {
						fmt.Println("You died!")
						return
					}

				case "run":
					fmt.Print("You run away from the LichDragon Fortissax...\n")
					c.ChooseLocation()
					return

				default:
					fmt.Println("Unknown action")
					continue
				}

				if LichDradonFortissaxHP <= 0 {
					c.LichDragonDead = true
					fmt.Println("You defeated the LichDragon Fortissax!")
					c.Gold += 10
					c.Silver += 235
					fmt.Println("You found 10 gold, 235 silver coins and claim soul of the LichDragon that improves your luck forever!")
					break
				}
			}

		case "exit":
			fmt.Print("Choose a place to teleport:\n")
			c.ChooseLocation()
			return

		default:
			fmt.Print("Unknown command. Choose from known floors\n")
			continue
		}
	}
}

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

// Start game function
func StartGame() {

	var name, characterClass string

	//Type a name of your character
	fmt.Print("Give a name for your character: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("input error when give a name to character: %v", err)
		return
	}

	//Choose a class of character
	fmt.Print("Choose class of your character: (Knight/Paladin/Mage/Priest)")
	_, err = fmt.Scan(&characterClass)
	if err != nil {
		fmt.Println("Input error when choose a class: %v", err)
		return
	}

	var character Character

	switch strings.ToLower(characterClass) {
	case "knight":
		character = &Knight{
			BaseCharacter{
				Name:   name,
				HP:     200,
				MaxHP:  200,
				DMG:    22,
				Gold:   10,
				Silver: 120,
			},
		}

	case "paladin":
		character = &Paladin{
			BaseCharacter{
				Name:   name,
				HP:     160,
				MaxHP:  160,
				DMG:    35,
				Gold:   13,
				Silver: 146,
			},
		}

	case "mage":
		character = &Mage{
			BaseCharacter{
				Name:   name,
				HP:     125,
				MaxHP:  125,
				DMG:    45,
				Gold:   34,
				Silver: 293,
			},
		}

	case "priest":
		character = &Priest{
			BaseCharacter{
				Name:   name,
				HP:     80,
				MaxHP:  80,
				DMG:    12,
				Gold:   2,
				Silver: 91,
			},
		}

	default:
		fmt.Println("Failed to create character! Using default class Knight")
		character = &Knight{
			BaseCharacter{
				Name:   name,
				HP:     200,
				MaxHP:  200,
				DMG:    22,
				Gold:   10,
				Silver: 120,
			},
		}
	}
	character.ChooseLocation()
}

func main() {
	StartGame()
}
