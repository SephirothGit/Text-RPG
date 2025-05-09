package core

import (
	"fmt"
	"strings"
	"github.com/SephirothGit/Text-RPG/pkg/character"
)

func StartGame() {
	var name, characterClass string

	fmt.Print("Give a name for your character: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("input error when give a name to character: %v", err)
		return
	}

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
			BaseCharacter: BaseCharacter{
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
			BaseCharacter: BaseCharacter{
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
			BaseCharacter: BaseCharacter{
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
			BaseCharacter: BaseCharacter{
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
			BaseCharacter: BaseCharacter{
				Name:   name,
				HP:     200,
				MaxHP:  200,
				DMG:    22,
				Gold:   10,
				Silver: 120,
			},
		}
	}
	
	c.ChooseLocation()
}