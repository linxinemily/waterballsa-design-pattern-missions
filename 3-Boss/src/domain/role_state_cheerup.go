package domain

type CheerUpState struct {
	*AbstractRoleState
}

func NewCheerUpState(role Role) *CheerUpState {
	return &CheerUpState{NewAbstractRoleState(role, "受到鼓舞")}
}

func (s *CheerUpState) attack(target Role, damageUnit int) {
	target.getDamagedBy(damageUnit+50, s.role)
}
