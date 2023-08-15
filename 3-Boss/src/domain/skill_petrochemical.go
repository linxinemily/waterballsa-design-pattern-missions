package domain

type PetrochemicalSkill struct {
	*AbstractSkill
}

func NewPetrochemicalSkill(owner Role) *PetrochemicalSkill {
	return &PetrochemicalSkill{NewAbstractSkill(owner, 100, "石化")}
}

func (s *PetrochemicalSkill) execute(targets []Role) {
	for _, target := range targets {
		target.setState(NewPetrochemicalState(target))
	}
}

func (s *PetrochemicalSkill) getTargets(allRolesOnBattle []Role) []Role {
	candidates := make([]Role, 0)

	for _, role := range allRolesOnBattle {
		if role.isEnemyOf(s.owner) {
			candidates = append(candidates, role)
		}
	}

	amount := 1
	if len(candidates) <= amount {
		return candidates
	}

	return s.owner.getTargetsFromInput(candidates, amount)
}
