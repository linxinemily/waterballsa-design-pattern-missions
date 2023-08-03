package domain

import (
	"C3M3H1/domain/enum"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Character struct {
	*AbstractRole
	direction enum.RoleDirection
}

func NewCharacter(m *Map) *Character {
	rand.Seed(time.Now().UnixNano())
	direction := enum.RoleDirection(rand.Intn(len(enum.AllDirections)))
	char := &Character{
		NewAbstractRole(m, direction.String(), 300),
		direction,
	}
	char.applyState(NewNormal(char))
	return char
}

func (char *Character) takeTurn() {
	char.m.print()
	directions := char.state.getValidDirections()
	fmt.Print("按下按鍵選擇動作：")
	for _, direction := range directions {
		fmt.Printf("[%s] %s", direction.Button().Key, direction.Button().Text)
	}
	fmt.Println(" [k] 攻擊")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := scanner.Text()

		if input == "k" {
			char.attack()
			break
		} else {
			var match bool
			for _, d := range enum.AllDirections {
				if d.Button().Key == input {
					stateGenerator, err := char.moveTo(char, d)
					if err != nil {
						fmt.Println(err.Error())
					} else {
						if stateGenerator != nil {
							char.applyState(stateGenerator.generate(char))
						}
						char.setDirection(d)
						match = true
						break
					}
				}
			}
			if match {
				break
			} else {
				fmt.Println("invalid input")
			}
		}
	}
	char.printState()
	char.m.print()
}

func (char *Character) printState() {
	fmt.Printf("主角 HP: %d, State: %s\n", char.hp, char.state)
}

func (char *Character) setDirection(direction enum.RoleDirection) {
	char.direction = direction
	char.symbol = direction.String()
}

func (char *Character) getValidEnemies() []Role {
	// 取得面向方位前方所有怪物，但不能穿越障礙物
	var monsters []Role

	switch char.direction {
	case enum.Right:
		for i := char.col + 1; i < char.m.size; i++ {
			if _, isObstacle := char.m.objects[char.row][i].(*Obstacle); isObstacle {
				break
			}
			if monster, isMonster := char.m.objects[char.row][i].(*Monster); isMonster {
				monsters = append(monsters, monster)
			}
		}
	case enum.Left:
		for i := char.col - 1; i >= 0; i-- {
			if _, isObstacle := char.m.objects[char.row][i].(*Obstacle); isObstacle {
				break
			}
			if monster, isMonster := char.m.objects[char.row][i].(*Monster); isMonster {
				monsters = append(monsters, monster)
			}
		}
	case enum.Top:
		for i := char.row - 1; i >= 0; i-- {
			if _, isObstacle := char.m.objects[i][char.col].(*Obstacle); isObstacle {
				break
			}
			if monster, isMonster := char.m.objects[i][char.col].(*Monster); isMonster {
				monsters = append(monsters, monster)
			}
		}
	case enum.Down:
		for i := char.row + 1; i < char.m.size; i++ {
			if _, isObstacle := char.m.objects[i][char.col].(*Obstacle); isObstacle {
				break
			}
			if monster, isMonster := char.m.objects[i][char.col].(*Monster); isMonster {
				monsters = append(monsters, monster)
			}
		}
	}
	return monsters
}

func (char *Character) afterAttacked() {
	char.hp -= 50
	char.applyState(NewInvincible(char))
}

func (char *Character) getAllEnemies() []Role {
	var monsters []Role
	for _, role := range char.m.roles {
		if monster, isMonster := role.(*Monster); isMonster {
			monsters = append(monsters, monster)
		}
	}
	return monsters
}
