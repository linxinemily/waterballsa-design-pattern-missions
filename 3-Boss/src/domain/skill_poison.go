package domain

type PoisonSkill struct {
	*AbstractSkill
}

func NewPoisonSkill(owner *RoleImpl) *PoisonSkill {
	return &PoisonSkill{NewAbstractSkill(owner, 80, "中毒")}
}

func (s *PoisonSkill) execute(targets []*RoleImpl) {
	target := targets[0]
	target.setState(NewPoisonedState(target))
}

func (s *PoisonSkill) getTargets(allRolesOnBattle []*RoleImpl) []*RoleImpl {
	candidates := make([]*RoleImpl, 0)

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
