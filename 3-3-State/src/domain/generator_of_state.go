package domain

type StateGenerator interface {
	generate(role Role) *IState
}

type StateInvincibleGenerator struct{}

func (g StateInvincibleGenerator) generate(role Role) *IState {
	return NewInvincible(role)
}

type StatePoisonedGenerator struct{}

func (g StatePoisonedGenerator) generate(role Role) *IState {
	return NewPoisoned(role)
}

type StateAcceleratedGenerator struct{}

func (g StateAcceleratedGenerator) generate(role Role) *IState {
	return NewAccelerated(role)
}

type StateHealingGenerator struct{}

func (g StateHealingGenerator) generate(role Role) *IState {
	return NewHealing(role)
}

type StateOrderlessGenerator struct{}

func (g StateOrderlessGenerator) generate(role Role) *IState {
	return NewOrderless(role)
}

type StateStockpileGenerator struct{}

func (g StateStockpileGenerator) generate(role Role) *IState {
	return NewStockpile(role)
}

type StateEruptingGenerator struct{}

func (g StateEruptingGenerator) generate(role Role) *IState {
	return NewErupting(role)
}

type StateTeleportGenerator struct{}

func (g StateTeleportGenerator) generate(role Role) *IState {
	return NewTeleport(role)
}
