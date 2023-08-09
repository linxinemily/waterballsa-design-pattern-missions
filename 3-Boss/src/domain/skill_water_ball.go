package domain

type WaterBallSkill struct {
	*AbstractSkill
}

func NewWaterBallSkill(owner *RoleImpl) *WaterBallSkill {
	return &WaterBallSkill{NewAbstractSkill(owner, 50, "水球")}
}

func (s *WaterBallSkill) execute(targets []*RoleImpl) {
	for _, target := range targets {
		s.owner.attack(target, 120)
	}
}

func (s *WaterBallSkill) getTargets(allRolesOnBattle []*RoleImpl) []*RoleImpl {
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
