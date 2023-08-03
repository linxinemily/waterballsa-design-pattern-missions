package domain

type Teleport struct {
	*AbstractState
}

func NewTeleport(role Role) *IState {
	return &IState{&Teleport{
		NewAbstractState(1, "瞬身", role),
	}}
}

func (state *Teleport) afterExpired() {
	state.role.getMap().moveRoleToRandomPosition(state.role)
	state.role.applyState(&IState{NewNormal(state.role)})
}
