package domain

import "math"

type Normal struct {
	*AbstractState
}

func NewNormal(role Role) *IState {
	return &IState{&Normal{
		NewAbstractState(math.MaxInt, "正常", role),
	}}
}
