package domain

type PoisonSkill struct {
	*AbstractSkill
}

func NewPoisonSkill(owner Role) *PoisonSkill {
	return &PoisonSkill{NewAbstractSkill(owner, 80)}
}

func (s *PoisonSkill) execute(targets []Role) {
	for _, target := range targets {
		target.setState(NewPoisonedState(target))
	}
}

func (s *PoisonSkill) getTargets(allRolesOnBattle []Role) []Role {
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

func (s *PoisonSkill) getName() string {
	return "中毒"
}
