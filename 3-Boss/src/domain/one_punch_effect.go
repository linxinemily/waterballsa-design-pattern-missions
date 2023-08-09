package domain

type OnePunchEffectImpl struct {
	OnePunchEffect
}

func NewOnePunchEffectImpl(effect OnePunchEffect) *OnePunchEffectImpl {
	return &OnePunchEffectImpl{
		effect,
	}
}

func (e *OnePunchEffectImpl) handle(targets []*RoleImpl) {
	if e.match(targets) {
		e.doHandling(targets)
	} else if e.getNext() != nil {
		e.getNext().handle(targets)
	}
}

type OnePunchEffect interface {
	match(targets []*RoleImpl) bool
	doHandling(targets []*RoleImpl)
	getNext() *OnePunchEffectImpl
}

type AbstractOnePunchEffect struct {
	next  *OnePunchEffectImpl
	skill *OnePunchSkill
}

func NewAbstractOnePunchEffect(n *OnePunchEffectImpl, skill *OnePunchSkill) *AbstractOnePunchEffect {
	return &AbstractOnePunchEffect{
		next:  n,
		skill: skill,
	}
}

func (e *AbstractOnePunchEffect) getNext() *OnePunchEffectImpl {
	return e.next
}
