package domain

import (
	"fmt"
	"strconv"
	"strings"
)

type Hero struct {
	*AbstractRole
}

func NewHero(id int, name string, HP int, MP int, STR int, rpg *RPG) *Hero {
	hero := &Hero{NewAbstractRole(id, name, HP, MP, STR, rpg)}
	hero.setState(NewNormalState(hero))
	return hero
}

func (r *Hero) getSkillFromInput() *SkillImpl {
	skills := r.getSkills()

	for {
		fmt.Println("選擇行動：")
		for i, skill := range skills {
			fmt.Printf("(%d) %s ", i, skill.getName())
		}
		fmt.Println()

		var userInput int
		fmt.Print("請輸入技能編號：")
		_, err := fmt.Scanf("%d", &userInput)
		if err != nil {
			fmt.Println("輸入錯誤")
			continue // 重新要求輸入
		}

		if userInput >= 0 && userInput < len(skills) {
			return skills[userInput]
		}

		fmt.Println("無效的技能編號")
	}
}

func (r *Hero) getTargetsFromInput(candidates []*RoleImpl, amount int) []*RoleImpl {
	selectedTargets := make([]*RoleImpl, 0)

	for len(selectedTargets) < amount {
		fmt.Printf("選擇 %d 位目標: ", amount)
		for i, candidate := range candidates {
			fmt.Printf("(%d) [%d]%s ", i, i+1, candidate.getName())
		}
		fmt.Println()

		var userInput string
		fmt.Print("請輸入目標編號，用逗號分隔：")
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			fmt.Println("輸入錯誤")
			continue
		}

		indices := strings.Split(userInput, ",")
		for _, indexStr := range indices {
			index := strings.TrimSpace(indexStr)
			if idx, err := strconv.Atoi(index); err == nil && idx >= 0 && idx < len(candidates) {
				selectedTargets = append(selectedTargets, candidates[idx])
			} else {
				fmt.Printf("無效的目標編號：%s\n", index)
			}
		}
	}

	return selectedTargets
}
