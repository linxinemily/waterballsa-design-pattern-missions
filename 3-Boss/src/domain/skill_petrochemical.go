package domain

type PetrochemicalSkill struct {
	*AbstractSkill
}

func NewPetrochemicalSkill(owner Role) *PetrochemicalSkill {
	return &PetrochemicalSkill{NewAbstractSkill(owner, 100)}
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

	if len(candidates) <= 3 {
		return candidates
	}

	return s.owner.getTargetsFromInput(candidates, 3)
}

func (s *PetrochemicalSkill) getName() string {
	return "石化"
}
