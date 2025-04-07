package main

import (
	"fmt"
	"strings"
)

//TODO:

//!!!Add legendary items and money reward in dungeon, fights, actions
//!!!Finally make good project architecture

// Character interface that defines methods for all characters
type Character interface {
	Fight()
	Sleep()
	Drink()
	Eat()
	ChooseLocation()
}

type BaseCharacter struct {
	Name   string
	HP     int
	MaxHP  int
	DMG    int
	Gold   int
	Silver int
}

type Knight struct {
	BaseCharacter
}

// Structure for Paladin
type Paladin struct {
	BaseCharacter
}

// Structure for Mage
type Mage struct {
	BaseCharacter
}

// Structure for Priest
type Priest struct {
	BaseCharacter
}

// Structure for item info
type MenuItem struct {
	Cost int
	Hp   int
}

// Map for drinks
var drinks = map[string]MenuItem{
	"water":  {Cost: 1, Hp: 10},
	"coffee": {Cost: 1, Hp: 10},
	"beer":   {Cost: 1, Hp: 15},
	"wine":   {Cost: 2, Hp: 30},
}

// Map for food
var food = map[string]MenuItem{
	"bread":    {Cost: 1, Hp: 20},
	"salad":    {Cost: 2, Hp: 30},
	"sandwich": {Cost: 3, Hp: 70},
	"steak":    {Cost: 5, Hp: 100},
}

// Fight method
func (c *BaseCharacter) Fight() {
	fmt.Printf("%s fights\n", c.Name)
}

// Sleep and restore full HP method
func (c *BaseCharacter) Sleep() {
	c.HP = c.MaxHP
	fmt.Printf("%s sleep and restores full HP %d/%d\n", c.Name, c.HP, c.MaxHP)
}

// Drink method
func (c *BaseCharacter) Drink() {

	fmt.Println("Available drinks: ")
	for name, item := range drinks {
		fmt.Printf("- %s (%d HP, %d silver)\n", name, item.Hp, item.Cost)
	}

	var drinkChoice string

	fmt.Print("Choose a drink: ")
	fmt.Scan(&drinkChoice)

	if item, exists := drinks[strings.ToLower(drinkChoice)]; exists {
		if c.Silver >= item.Cost {
			c.Silver -= item.Cost
			c.HP += item.Hp
			if c.HP > c.MaxHP {
				c.HP = c.MaxHP
			}
			fmt.Printf("%s drinks %s and restores %d HP (-%d silver coins)\n", c.Name, drinkChoice, item.Hp, item.Cost)
		} else {
			fmt.Println("Not enough silver coins")
		}
	} else {
		fmt.Println("We don't have this drink")
	}

}

// Eat method
func (c *BaseCharacter) Eat() {

	fmt.Println("Available food: ")
	for name, item := range food {
		fmt.Printf("- %s (%d HP, %d silver)\n", name, item.Hp, item.Cost)
	}

	var choice string

	fmt.Print("Choose a meal: ")
	fmt.Scan(&choice)

	if item, exists := food[strings.ToLower(choice)]; exists {
		if c.Silver >= item.Cost {
			c.Silver -= item.Cost
			c.HP += item.Hp
			if c.HP > c.MaxHP {
				c.HP = c.MaxHP
			}
			fmt.Printf("%s eats %s and restores %d HP (-%d silver)\n", c.Name, choice, item.Hp, item.Cost)
		} else {
			fmt.Println("Not enough silver")
		}
	} else {
		fmt.Println("We don't have that meal")
	}
}

// Fight overridden method for Warrior
func (k *Knight) Fight() {
	fmt.Printf("%s (Knight) hit the enemy with a stick for %d DMG\n", k.Name, k.DMG)
}

// Fight overridden method for Paladin
func (p *Paladin) Fight() {
	fmt.Printf("%s (Paladin) hit the enemy with a holy sword for %d DMG\n", p.Name, p.DMG)
}

// Fight overridden method for Mage
func (m *Mage) Fight() {
	fmt.Printf("%s (Mage) cast a fireball for %d DMG\n", m.Name, m.DMG)
}

func (r *Priest) Fight() {
	fmt.Printf("%s (Priest) hit the enemy with a holy hammer for %d DMG\n", r.Name, r.DMG)
}

// Choose location func
func (c *BaseCharacter) ChooseLocation() {
	for {
		var location string

		fmt.Print("Where do you want to go? (Forest/Town/Tavern/Exit)")
		_, err := fmt.Scan(&location)
		if err != nil {
			fmt.Println("input error", err)
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

// Method for going to the forbidden forest
func (c *BaseCharacter) GoToForest() {
	fmt.Printf("%s goes to the Forbidden forest...\n", c.Name)

	//Improved fight system
	goblinHP := 50
	fmt.Printf("A goblin (HP: %d) has attacked you!\n", goblinHP)

	for {
		fmt.Printf("%s HP: %d/%d | Goblin HP %d\n", c.Name, c.HP, c.MaxHP, goblinHP)
		fmt.Print("What will you do? (Fight/Run)")

		var action string
		fmt.Scan(&action)

		switch strings.ToLower(action) {
		case "fight":
			fmt.Printf("%s attacks the goblin for %d DMG!\n", c.Name, c.DMG)
			goblinHP -= c.DMG

			if goblinHP <= 0 {
				fmt.Println("You defeated the goblin!\n")
				c.ChooseLocation()
				return
			}
			//Monster fights
			goblinDMG := 10
			c.HP -= goblinDMG
			fmt.Printf("Goblin hits you with a club for %d damage!\n", goblinDMG)

			if c.HP <= 0 {
				fmt.Println("You died")
				return
			}
		case "run":
			fmt.Print("You run away from the goblin!\n")
			c.ChooseLocation()
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}

// Method for going to the town
func (c *BaseCharacter) GoToTown() {
	fmt.Printf("%s goes to the Town...\n", c.Name)

	for {
		//Choose a place to visit in town
		var choosePlace string

		fmt.Print("Which place you want to visit? (Dungeon/Inn/Tavern/Exit)")
		_, err := fmt.Scan(&choosePlace)
		if err != nil {
			fmt.Println("input error:", err)
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

	//Loop for actions
	for {

		var tavernAction string

		fmt.Print("Olaf: Hey, traveler! Do you want to drink a cold beer or a room for a night? (Eat/Drink/Sleep/Exit)")

		//Error handling
		_, err := fmt.Scan(&tavernAction)
		if err != nil {
			fmt.Println("Input error:", err)
		}

		switch strings.ToLower(tavernAction) {
		case "drink":
			fmt.Printf("Olaf: What would you like to drink, %s?\n", c.Name)
			c.Drink()
		case "sleep":
			fmt.Print("Olaf: your room is first door on the right side of second floor\n")
			c.Sleep()
		case "eat":
			fmt.Print("Waitress: What would you like to eat today, %s?\n", c.Name)
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

func (c *BaseCharacter) GoToTavernB() {
	fmt.Printf("Waitress: Welcome to the Brick's Tavern %s!\n", c.Name)

	for {
		var tavernActionB string

		fmt.Print("Brick: Good day, traveler! Do you want to drink a glass of cold beer with a snacks? (Drink/Eat/Exit)")
		_, err := fmt.Scan(&tavernActionB)
		if err != nil {
			fmt.Println("input err", err)
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

func (c *BaseCharacter) GoToDungeon() {

	for {
		var dungeonAction string

		fmt.Printf("%s went down into the dungeon...\n", c.Name)

		fmt.Print("Choose floor: (First/Second/Third/Teleport)")

		_, err := fmt.Scan(&dungeonAction)
		if err != nil {
			fmt.Println("Input error:", err)
		}

		switch strings.ToLower(dungeonAction) {
		case "first":
			fmt.Printf("%s went down to the 1 floor of the Ancient Elven Dungeon\n", c.Name)
		case "second":
			fmt.Printf("%s went to the middle floor of the Ancient Elven Dungeon\n", c.Name)
		case "third":
			fmt.Printf("%s went to the lower floor of the Ancient Elven Dungeon\n", c.Name)
		case "teleport":
			fmt.Print("Choose a place to teleport:\n")
			c.ChooseLocation()
			return
		default:
			fmt.Print("Unknown command. Choose from known floors\n")
		}
	}
}
func (c *BaseCharacter) GoToInn() {
	fmt.Print("Welcome to the Mermaid's Inn!\n")

	//How many paid days left
	remainingNights := 0

	for {
		if remainingNights > 0 {
			fmt.Printf("Freya: You have %d remaining nights, would you like go to sleep? (Yes/No)\n", remainingNights)

			var usePaidNight string
			_, err := fmt.Scan(&usePaidNight)
			if err != nil {
				fmt.Println("Input error", err)
				continue
			}
			if strings.ToLower(usePaidNight) == "Yes" {
				remainingNights--
				c.Sleep()
				fmt.Print("You have %d nights remaining.\n", remainingNights)
				continue
			}
		}

		for {
			var innDuration string
			fmt.Print("Freya: We offer best rooms in Fawn! For how many days you want to stay in town? (1/7/30/Exit)")

			_, err := fmt.Scan(&innDuration)
			if err != nil {
				fmt.Println("Input error:", err)
				continue
			}

			// Choose for how long you want to stay in Fawn
			switch strings.ToLower(innDuration) {
			case "1":
				fmt.Print("Freya: It will cost you 1 silver coin\n")
				if c.Silver >= 1 {
					c.Silver -= 1
					fmt.Print("1 silver coin was given away\n")
					c.Sleep()
				} else {
					fmt.Println("Not enough silver")
					continue
				}

			case "7":
				fmt.Print("Freya: It will cost you 7 silver coins\n")
				if c.Silver >= 7 {
					c.Silver -= 7
					fmt.Print("7 silver coins were given away\n")
					c.Sleep()
				} else {
					fmt.Println("Not enough silver")
					continue
				}

			case "30":
				fmt.Print("Freya: Great choice! It will cost you only 27 silver coins!\n")
				if c.Silver >= 27 {
					c.Silver -= 27
					fmt.Print("27 silver coins were given away\n")
					c.Sleep()
				} else {
					fmt.Println("Not enough silver")
					continue
				}

			case "exit":
				if remainingNights > 0 {
					fmt.Printf("You still have %d paid night(s) in our inn", remainingNights)
				}
				fmt.Print("You came out from the Mermaid's inn")
				c.ChooseLocation()
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
						fmt.Println("Input error:", err)
						continue
					}

					switch strings.ToLower(innAction) {
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
					}
				}
			}
		}
	}

// Start game function
func StartGame() {

	var name, characterClass string

	//Type a name of your character
	fmt.Print("Give a name for your character: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("input error:", err)
		return
	}

	//Choose a class of character
	fmt.Print("Choose class of your character: (Knight/Paladin/Mage/Priest)")
	_, err = fmt.Scan(&characterClass)
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	var character Character

	switch strings.ToLower(characterClass) {
	case "knight":
		character = &Knight{BaseCharacter{Name: name, HP: 180, MaxHP: 180, DMG: 20, Gold: 10, Silver: 120}}
	case "paladin":
		character = &Paladin{BaseCharacter{Name: name, HP: 220, MaxHP: 220, DMG: 24, Gold: 13, Silver: 146}}
	case "mage":
		character = &Mage{BaseCharacter{Name: name, HP: 120, MaxHP: 120, DMG: 36, Gold: 34, Silver: 293}}
	case "priest":
		character = &Priest{BaseCharacter{Name: name, HP: 80, MaxHP: 80, DMG: 8, Gold: 2, Silver: 91}}
	default:
		fmt.Println("Unknown class. Create Knight by default")
		character = &Knight{BaseCharacter{Name: name, HP: 180, MaxHP: 180, DMG: 20, Gold: 10, Silver: 120}}
	}
	character.ChooseLocation()
}

func main() {
	StartGame()
}
