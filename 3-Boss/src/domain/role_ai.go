package domain

import (
	"fmt"
)

type AI struct {
	*AbstractRole
	seed int
}

func NewAI(id int, name string, HP int, MP int, STR int, rpg *RPG) *AI {
	ai := &AI{AbstractRole: NewAbstractRole(id, name, HP, MP, STR, rpg)}
	ai.setState(NewNormalState(ai))
	return ai
}

func (r *AI) getSkillFromInput() *SkillImpl {

	fmt.Fprintf(r.getRPG().getWriter(), "選擇行動：")
	for i, skill := range r.getSkills() {
		fmt.Fprintf(r.getRPG().getWriter(), "(%d) %s ", i, skill.getName())
	}
	fmt.Fprintln(r.getRPG().getWriter())

	skill := r.getSkills()[r.seed%len(r.getSkills())]
	r.seed++

	return skill
}

func (r *AI) getTargetsFromInput(candidates []Role, amount int) []Role {
	targets := make([]Role, 0)

	for index := range candidates {
		if len(targets) >= amount {
			break
		}

		targets = append(targets, candidates[(r.seed+index)%len(candidates)])
	}

	r.seed++

	return targets
}

func (r *AI) setSeed(seed int) {
	r.seed = seed
}
