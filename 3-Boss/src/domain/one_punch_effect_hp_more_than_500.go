package domain

type HpMoreThan500Effect struct {
	*AbstractOnePunchEffect
}

func NewHpMoreThan500Effect(next *OnePunchEffectImpl, skill *OnePunchSkill) *HpMoreThan500Effect {
	return &HpMoreThan500Effect{NewAbstractOnePunchEffect(next, skill)}
}

func (e *HpMoreThan500Effect) match(targets []Role) bool {
	return targets[0].getHp() >= 500
}

func (e *HpMoreThan500Effect) doHandling(targets []Role) {
	e.skill.getOwner().attack(targets[0], 300)
}
