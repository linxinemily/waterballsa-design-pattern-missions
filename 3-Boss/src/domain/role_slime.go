package domain

import "strconv"

type Slime struct {
	*AbstractRole
}

func NewSlime(id int, rpg *RPG) *Slime {
	slime := &Slime{NewAbstractRole(id, "slime"+strconv.Itoa(id), 10, 0, 50, rpg)}
	slime.setState(NewNormalState(slime))
	return slime
}

func (r *Slime) getSkillFromInput() *SkillImpl {
	return nil
}

func (r *Slime) getTargetsFromInput(candidates []*RoleImpl, amount int) []*RoleImpl {
	return nil
}
