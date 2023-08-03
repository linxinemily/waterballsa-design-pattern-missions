package domain

import (
	"C3M3H1/domain/enum"
	"math/rand"
	"time"
)

type Orderless struct {
	*AbstractState
}

func NewOrderless(role Role) *IState {
	return &IState{&Orderless{
		NewAbstractState(3, "混亂", role),
	}}
}

func (state *Orderless) getValidDirections() []enum.RoleDirection {
	rand.Seed(time.Now().UnixNano())
	return [][]enum.RoleDirection{
		{enum.Top, enum.Down},
		{enum.Left, enum.Right},
	}[rand.Intn(2)]
}
