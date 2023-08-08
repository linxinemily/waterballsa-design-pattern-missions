package domain

type NormalState struct {
	*AbstractRoleState
}

func NewNormalState(role Role) *NormalState {
	return &NormalState{NewAbstractRoleState(role)}
}
