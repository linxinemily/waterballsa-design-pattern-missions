package domain

type PoisonedState struct {
	*AbstractRoleState
}

func NewPoisonedState(role Role) *PoisonedState {
	return &PoisonedState{NewAbstractRoleState(role, "中毒")}
}

func (s *PoisonedState) beforeTakeTurn() (canGoOn bool) {
	s.role.getDamaged(30)
	if s.role.isAlive() {
		return true
	} else {
		return false
	}
}
