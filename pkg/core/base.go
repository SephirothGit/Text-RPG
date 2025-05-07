package core

import (
	"fmt"
	"strings"
)

type Character interface {
	Fight()
	Sleep()
	Drink()
	Eat()
	ChooseLocation()
}

type BaseCharacter struct {
	Name           string
	HP             int
	MaxHP          int
	DMG            int
	Silver         int
	Gold           int
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

type MenuItem struct {
	Cost int
	Hp   int
}

var drinks = map[string]MenuItem{
	"water":  {Cost: 1, Hp: 10},
	"coffee": {Cost: 1, Hp: 10},
	"beer":   {Cost: 1, Hp: 15},
	"wine":   {Cost: 2, Hp: 30},
}

var food = map[string]MenuItem{
	"bread":    {Cost: 1, Hp: 20},
	"salad":    {Cost: 2, Hp: 30},
	"sandwich": {Cost: 3, Hp: 70},
	"steak":    {Cost: 5, Hp: 100},
}

func (c *BaseCharacter) Fight() {
	fmt.Printf("%s attacks the enemy for %d DMG!\n", c.Name, c.DMG)
}

func (c *BaseCharacter) Sleep() {
	c.HP = c.MaxHP
	fmt.Printf("%s sleep and restores full HP %d/%d\n", c.Name, c.HP, c.MaxHP)
}

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
			fmt.Printf("%s drinks %s and restores %d HP (-%d silver coins) HP(%d/%d)\n", 
				c.Name, drinkChoice, item.Hp, item.Cost, c.HP, c.MaxHP)
		} else {
			fmt.Println("Not enough silver coins")
		}
	} else {
		fmt.Println("We don't have this drink")
	}
}

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
			fmt.Printf("%s eats %s and restores %d HP (-%d silver) HP(%d/%d) \n", 
				c.Name, foodChoice, item.Hp, item.Cost, c.HP, c.MaxHP)
		} else {
			fmt.Println("Not enough silver")
		}
	} else {
		fmt.Println("We don't have that meal")
	}
}