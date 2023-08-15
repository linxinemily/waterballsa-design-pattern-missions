package domain

type PetrochemicalState struct {
	*AbstractRoleState
}

func NewPetrochemicalState(role Role) *PetrochemicalState {
	return &PetrochemicalState{NewAbstractRoleState(role, "石化")}
}

func (s *PetrochemicalState) beforeTakeTurn() (canGoOn bool) {
	return false
}
