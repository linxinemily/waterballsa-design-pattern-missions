package domain

type FireBallSkill struct {
	*AbstractSkill
}

func NewFireBallSkill(owner *RoleImpl) *FireBallSkill {
	return &FireBallSkill{NewAbstractSkill(owner, 50, "火球")}
}

func (s *FireBallSkill) execute(targets []*RoleImpl) {
	for _, target := range targets {
		s.owner.attack(target, 50)
	}
}

func (s *FireBallSkill) getTargets(allRolesOnBattle []*RoleImpl) []*RoleImpl {
	candidates := make([]*RoleImpl, 0)

	for _, role := range allRolesOnBattle {
		if role.isEnemyOf(s.owner) {
			candidates = append(candidates, role)
		}
	}

	return candidates
}
