package domain

type CheerUpState struct {
	*AbstractRoleState
}

func NewCheerUpState(role Role) *CheerUpState {
	return &CheerUpState{NewAbstractRoleState(role)}
}

func (s *CheerUpState) attack(target Role, damageUnit int) {
	target.getDamaged(damageUnit + 50)
}
