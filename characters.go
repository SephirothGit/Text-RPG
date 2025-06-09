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

// Cgaracter stats
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

// Structure for Knight
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


// Start game function
func StartGame() {

	var name, characterClass string

	//Type a name of your character
	fmt.Print("Give a name for your character: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Printf("input error when give a name to character: %v", err)
		return
	}

	//Choose a class of character
	fmt.Print("Choose class of your character: (Knight/Paladin/Mage/Priest)")
	_, err = fmt.Scan(&characterClass)
	if err != nil {
		fmt.Printf("Input error when choose a class: %v", err)
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
		fmt.Printf("Input error when choose a drink: %v", err)
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
		fmt.Printf("Input error when choose a meal: %v", err)
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

// Fight overridden method for Warrior
func (k *Knight) Fight() {
	fmt.Printf("%s (Knight) hits the enemy with a sword for %d DMG!\n", k.Name, k.DMG)
}

// Fight overridden method for Paladin
func (p *Paladin) Fight() {
	fmt.Printf("%s (Paladin) hits the enemy with a holy sword for %d DMG!\n", p.Name, p.DMG)
}

// Fight overridden method for Mage
func (m *Mage) Fight() {
	fmt.Printf("%s (Mage) cast a fireball for %d DMG!\n", m.Name, m.DMG)
}

// Fight overridden method for Priest
func (r *Priest) Fight() {
	fmt.Printf("%s (Priest) hit the enemy with a holy hammer for %d DMG!\n", r.Name, r.DMG)
}
