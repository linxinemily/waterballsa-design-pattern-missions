package domain

type PetrochemicalSkill struct {
	*AbstractSkill
}

func NewPetrochemicalSkill(owner *RoleImpl) *PetrochemicalSkill {
	return &PetrochemicalSkill{NewAbstractSkill(owner, 100, "石化")}
}

func (s *PetrochemicalSkill) execute(targets []*RoleImpl) {
	for _, target := range targets {
		target.setState(NewPetrochemicalState(target))
	}
}

func (s *PetrochemicalSkill) getTargets(allRolesOnBattle []*RoleImpl) []*RoleImpl {
	candidates := make([]*RoleImpl, 0)

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
