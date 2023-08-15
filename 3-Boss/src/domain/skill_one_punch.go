package domain

type OnePunchSkill struct {
	*AbstractSkill
	*OnePunchEffectImpl
}

func NewOnePunchSkill(owner Role) *OnePunchSkill {
	skill := &OnePunchSkill{AbstractSkill: NewAbstractSkill(owner, 180, "一拳攻擊")}
	skill.OnePunchEffectImpl = &OnePunchEffectImpl{NewHpMoreThan500Effect(
		&OnePunchEffectImpl{NewPoisonedOrPetrochemicalEffect(
			&OnePunchEffectImpl{NewCheerUpEffect(
				&OnePunchEffectImpl{NewNormalStateEffect(nil, skill)},
				skill)},
			skill)},
		skill)}

	return skill
}

func (s *OnePunchSkill) execute(targets []Role) {
	s.handle(targets)
}

func (s *OnePunchSkill) getTargets(allRolesOnBattle []Role) []Role {
	candidates := make([]Role, 0)

	for _, role := range allRolesOnBattle {
		if role.isEnemyOf(s.owner) {
			candidates = append(candidates, role)
		}
	}

	if len(candidates) <= 1 {
		return candidates
	}

	return s.owner.getTargetsFromInput(candidates, 1)
}
