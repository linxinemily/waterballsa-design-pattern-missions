package domain

type PoisonedState struct {
	*AbstractRoleState
}

func NewPoisonedState(role Role) *PoisonedState {
	return &PoisonedState{NewAbstractRoleState(role, "中毒")}
}

func (s *PoisonedState) beforeTakeTurn() (canGoOn bool) {
	s.role.getDamagedBy(30, nil)
	if s.role.isAlive() {
		return true
	} else {
		return false
	}
}
