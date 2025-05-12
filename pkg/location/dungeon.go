package location

import (
	"fmt"
	"strings"
	"github.com/SephirothGit/Text-RPG/pkg/core"
)

// Go to the Ancient Elves Dungeon
func (c *core.BaseCharacter) GoToDungeon() {

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
