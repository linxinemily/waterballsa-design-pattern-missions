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
	skillNames := make([]string, len(skills))
	for i, skill := range skills {
		skillNames[i] = fmt.Sprintf("(%d) %s", i, skill.getName())
	}
	fmt.Fprintln(r.getRPG().getWriter(), fmt.Sprintf("選擇行動：%s", strings.Join(skillNames, " ")))

	scanner.Scan()
	userInput := scanner.Text()

	choice, err := strconv.Atoi(userInput)
	if err != nil {
		panic(err)
	}

	if choice >= 0 && choice < len(skills) {
		return skills[choice]
	}

	panic(fmt.Sprintf("無效的選擇: %d", choice))
}

func (r *Hero) getTargetsFromInput(candidates []Role, amount int) []Role {
	selectedTargets := make([]Role, 0)
	scanner := r.getScanner()

	for len(selectedTargets) < amount {
		candidatesStr := make([]string, len(candidates))
		for i, candidate := range candidates {
			candidatesStr[i] = fmt.Sprintf("(%d) %s", i, candidate.getNameWithTroopId())
		}
		fmt.Fprintln(r.getRPG().getWriter(), fmt.Sprintf("選擇 %d 位目標: %s", amount, strings.Join(candidatesStr, " ")))

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
