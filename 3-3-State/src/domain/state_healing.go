package domain

type Healing struct {
	*AbstractState
}

func NewHealing(role Role) *IState {
	return &IState{&Healing{
		NewAbstractState(5, "恢復", role),
	}}
}

func (state *Healing) beforeTakeTurn() int {
	if state.role.getHP()+30 >= state.role.getFullHP() {
		state.role.applyState(NewNormal(state.role))
	}
	state.role.addHP(30)

	return 1
}
