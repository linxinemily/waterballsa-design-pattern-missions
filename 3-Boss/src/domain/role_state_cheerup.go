package domain

type CheerUpState struct {
	*AbstractRoleState
}

func NewCheerUpState(role Role) *CheerUpState {
	return &CheerUpState{NewAbstractRoleState(role, "鼓舞")}
}

func (s *CheerUpState) attack(target Role, damageUnit int) {
	target.getDamaged(damageUnit + 50)
}
