package domain

import (
	"fmt"
	"log"
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
	scanner := r.getScanner()
	for {
		fmt.Fprintf(r.getRPG().getWriter(), "選擇行動：")
		for i, skill := range skills {
			fmt.Fprintf(r.getRPG().getWriter(), "(%d) %s ", i, skill.getName())
		}
		fmt.Fprintln(r.getRPG().getWriter())

		scanner.Scan()
		userInput := scanner.Text()

		choice, err := strconv.Atoi(userInput)
		if err != nil {
			log.Println("輸入錯誤")
			continue // 重新要求输入
		}

		if choice >= 0 && choice < len(skills) {
			return skills[choice]
		}

		log.Println("無效的技能編號")
	}
}

func (r *Hero) getTargetsFromInput(candidates []Role, amount int) []Role {
	selectedTargets := make([]Role, 0)
	scanner := r.getScanner()

	for len(selectedTargets) < amount {
		fmt.Fprintf(r.getRPG().getWriter(), "選擇 %d 位目標: ", amount)
		for i, candidate := range candidates {
			fmt.Fprintf(r.getRPG().getWriter(), "(%d) %s ", i, candidate.getNameWithTroopId())
		}
		fmt.Fprintln(r.getRPG().getWriter())

		scanner.Scan()
		userInput := scanner.Text()

		indices := strings.Split(userInput, ",")
		for _, indexStr := range indices {
			index := strings.TrimSpace(indexStr)
			if idx, err := strconv.Atoi(index); err == nil && idx >= 0 && idx < len(candidates) {
				selectedTargets = append(selectedTargets, candidates[idx])
			} else {
				log.Printf("無效的目標編號：%s\n", index)
			}
		}
	}

	return selectedTargets
}
