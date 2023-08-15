package domain

type CheerUpSkill struct {
	*AbstractSkill
}

func NewCheerUpSkill(owner Role) *CheerUpSkill {
	return &CheerUpSkill{NewAbstractSkill(owner, 100, "鼓舞")}
}

func (s *CheerUpSkill) execute(targets []Role) {
	for _, target := range targets {
		target.setState(NewCheerUpState(target))
	}
}

func (s *CheerUpSkill) getTargets(allRolesOnBattle []Role) []Role {
	candidates := make([]Role, 0)

	for _, role := range allRolesOnBattle {
		if role.isAllyOf(s.owner) {
			candidates = append(candidates, role)
		}
	}

	if len(candidates) <= 3 {
		return candidates
	}

	return s.owner.getTargetsFromInput(candidates, 3)
}
