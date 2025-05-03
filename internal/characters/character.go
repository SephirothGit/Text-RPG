package characters

import (
	"fmt"
	"strings"
)

// This file contains Character interface, BaseCharacter structure, inventory and items, methods and overridden methods for BaseCharacter
// Also contains ChooseLocation function

// Cgaracter stats
type BaseCharacter struct {
	Name   string
	HP     int
	MaxHP  int
	DMG    int
	Silver int
	Gold   int
	SkeletonDead   bool
	SalamanderDead bool
	LichDragonDead bool
}


type Item struct {
	Name        string
	Description string
	Value       int
}

type Inventory struct {
	Items   []Item
	MaxSize int
}

type Equipment struct {
	Weapon *Item
	Armor  *Item
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
	fmt.Printf("%s attacks the enemy for %d DMG!\n", c.Name, c.DMG)
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
	_, err := fmt.Scan(&drinkChoice)
	if err != nil {
		fmt.Println("Input error when choose a drink: %v", err)
	}

	if item, exists := drinks[strings.ToLower(drinkChoice)]; exists {
		if c.Silver >= item.Cost {
			c.Silver -= item.Cost
			c.HP += item.Hp
			if c.HP > c.MaxHP {
				c.HP = c.MaxHP
			}
			fmt.Printf("%s drinks %s and restores %d HP (-%d silver coins) HP(%d/%d)\n", c.Name, drinkChoice, item.Hp, item.Cost, c.HP, c.MaxHP)
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

	var foodChoice string

	fmt.Print("Choose a meal: ")
	_, err := fmt.Scan(&foodChoice)
	if err != nil {
		fmt.Println("Input error when choose a meal: %v", err)
	}

	if item, exists := food[strings.ToLower(foodChoice)]; exists {
		if c.Silver >= item.Cost {
			c.Silver -= item.Cost
			c.HP += item.Hp
			if c.HP > c.MaxHP {
				c.HP = c.MaxHP
			}
			fmt.Printf("%s eats %s and restores %d HP (-%d silver) HP(%d/%d) \n", c.Name, foodChoice, item.Hp, item.Cost, c.HP, c.MaxHP)
		} else {
			fmt.Println("Not enough silver")
		}
	} else {
		fmt.Println("We don't have that meal")
	}
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