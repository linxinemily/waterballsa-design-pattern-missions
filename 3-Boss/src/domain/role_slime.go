package domain

import (
	"fmt"
)

type Slime struct {
	*AbstractRole
}

func NewSlime(id int, rpg *RPG) *Slime {
	slime := &Slime{NewAbstractRole(id, "Slime", 100, 0, 50, rpg)}
	slime.setState(NewNormalState(slime))
	return slime
}

func (r *Slime) getSkillFromInput() *SkillImpl {
	fmt.Fprintln(r.getRPG().getWriter(), fmt.Sprintf("選擇行動：(0) %s", r.getSkills()[0].getName()))
	return r.getSkills()[0]
}

func (r *Slime) getTargetsFromInput(candidates []Role, amount int) []Role {
	return candidates[:amount]
}
