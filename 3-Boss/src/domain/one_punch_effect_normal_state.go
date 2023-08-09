package domain

type NormalStateEffect struct {
	*AbstractOnePunchEffect
}

func NewNormalStateEffect(next *OnePunchEffectImpl, skill *OnePunchSkill) *NormalStateEffect {
	return &NormalStateEffect{NewAbstractOnePunchEffect(next, skill)}
}

func (e *NormalStateEffect) match(targets []*RoleImpl) bool {
	switch targets[0].getState().(type) {
	case *NormalState:
		return true
	default:
		return false
	}
}

func (e *NormalStateEffect) doHandling(targets []*RoleImpl) {
	e.skill.getOwner().attack(targets[0], 100)
}
