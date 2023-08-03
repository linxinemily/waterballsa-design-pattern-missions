package domain

type Invincible struct {
	*AbstractState
}

func NewInvincible(role Role) *IState {
	return &IState{&Invincible{
		NewAbstractState(2, "無敵", role),
	}}
}

func (state *Invincible) attacked() {
	//
}
