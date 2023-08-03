package domain

type Poisoned struct {
	*AbstractState
}

func NewPoisoned(role Role) *IState {
	return &IState{&Poisoned{
		NewAbstractState(3, "中毒", role),
	}}
}

func (state *Poisoned) beforeTakeTurn() int {
	state.role.subtractHP(15)
	return 1
}
