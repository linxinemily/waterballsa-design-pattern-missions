package domain

type CheerUpEffect struct {
	*AbstractOnePunchEffect
}

func NewCheerUpEffect(next *OnePunchEffectImpl, skill *OnePunchSkill) *CheerUpEffect {
	return &CheerUpEffect{NewAbstractOnePunchEffect(next, skill)}
}

func (e *CheerUpEffect) match(targets []Role) bool {
	switch targets[0].getState().(type) {
	case *CheerUpState:
		return true
	default:
		return false
	}
}

func (e *CheerUpEffect) doHandling(targets []Role) {
	target := targets[0]
	e.skill.getOwner().attack(target, 100)
	target.setState(NewNormalState(target))
}
