package domain

type CheerUpSkill struct {
	*AbstractSkill
}

func NewCheerUpSkill(owner *RoleImpl) *CheerUpSkill {
	return &CheerUpSkill{NewAbstractSkill(owner, 100, "鼓舞")}
}

func (s *CheerUpSkill) execute(targets []*RoleImpl) {
	for _, target := range targets {
		target.setState(NewCheerUpState(target))
	}
}

func (s *CheerUpSkill) getTargets(allRolesOnBattle []*RoleImpl) []*RoleImpl {
	candidates := make([]*RoleImpl, 0)

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
