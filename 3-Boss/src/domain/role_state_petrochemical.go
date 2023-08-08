package domain

type PetrochemicalState struct {
	*AbstractRoleState
}

func NewPetrochemicalState(role Role) *PetrochemicalState {
	return &PetrochemicalState{NewAbstractRoleState(role)}
}

func (s *PetrochemicalState) beforeTakeTurn() (canGoOn bool) {
	return false
}
