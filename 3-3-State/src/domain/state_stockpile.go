package domain

type Stockpile struct {
	*AbstractState
}

func NewStockpile(role Role) *IState {
	return &IState{&Stockpile{
		NewAbstractState(2, "蓄力", role),
	}}
}

func (state *Stockpile) afterExpired() {
	state.role.applyState(NewErupting(state.role))
}

func (state *Stockpile) attacked() {
	state.role.applyState(NewNormal(state.role))
}
