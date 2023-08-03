package domain

type Erupting struct {
	*AbstractState
}

func NewErupting(role Role) *IState {
	return &IState{&Erupting{
		NewAbstractState(3, "爆發", role),
	}}
}

func (state *Erupting) afterExpired() {
	state.role.applyState(NewTeleport(state.role))
}

func (state *Erupting) attack() {
	for _, enemy := range state.role.getAllEnemies() {
		enemy.attacked()
	}
}
