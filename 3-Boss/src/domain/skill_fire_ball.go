package domain

type FireBallSkill struct {
	*AbstractSkill
}

func NewFireBallSkill(owner Role) *FireBallSkill {
	return &FireBallSkill{NewAbstractSkill(owner, 50, "火球")}
}

func (s *FireBallSkill) execute(targets []Role) {
	for _, target := range targets {
		s.owner.attack(target, 50)
	}
}

func (s *FireBallSkill) getTargets(allRolesOnBattle []Role) []Role {
	candidates := make([]Role, 0)

	for _, role := range allRolesOnBattle {
		if role.isEnemyOf(s.owner) {
			candidates = append(candidates, role)
		}
	}

	return candidates
}
