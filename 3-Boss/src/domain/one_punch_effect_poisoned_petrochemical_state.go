package domain

type PoisonedOrPetrochemicalEffect struct {
	*AbstractOnePunchEffect
}

func NewPoisonedOrPetrochemicalEffect(next *OnePunchEffectImpl, skill *OnePunchSkill) *PoisonedOrPetrochemicalEffect {
	return &PoisonedOrPetrochemicalEffect{NewAbstractOnePunchEffect(next, skill)}
}

func (e *PoisonedOrPetrochemicalEffect) match(targets []Role) bool {
	switch targets[0].getState().(type) {
	case *PoisonedState:
		return true
	case *PetrochemicalState:
		return true
	default:
		return false
	}
}

func (e *PoisonedOrPetrochemicalEffect) doHandling(targets []Role) {
	for i := 0; i < 3; i++ {
		e.skill.getOwner().attack(targets[0], 80)
	}
}
