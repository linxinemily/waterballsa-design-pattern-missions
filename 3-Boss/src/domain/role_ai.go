package domain

type AI struct {
	*AbstractRole
}

func NewAI(id int, name string, HP int, MP int, STR int, rpg *RPG) *AI {
	ai := &AI{NewAbstractRole(id, name, HP, MP, STR, rpg)}
	ai.setState(NewNormalState(ai))
	return ai
}

func (r *AI) getSkillFromInput() *SkillImpl {
	return nil
}

func (r *AI) getTargetsFromInput(candidates []*RoleImpl, amount int) []*RoleImpl {
	return nil
}
