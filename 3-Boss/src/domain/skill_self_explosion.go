package domain

type SelfExplosionSkill struct {
	*AbstractSkill
}

func NewSelfExplosionSkill(owner Role) *SelfExplosionSkill {
	return &SelfExplosionSkill{NewAbstractSkill(owner, 200, "自爆")}
}

func (s *SelfExplosionSkill) execute(targets []Role) {
	for _, target := range targets {
		s.owner.attack(target, 150)
	}
	s.owner.getDamagedBy(s.owner.getHp(), nil)
}

func (s *SelfExplosionSkill) getTargets(allRolesOnBattle []Role) []Role {
	targetsExceptOwner := make([]Role, 0)
	for i, target := range allRolesOnBattle {
		if target.getId() != s.owner.getId() {
			targetsExceptOwner = append(targetsExceptOwner, allRolesOnBattle[i])
		}
	}
	return targetsExceptOwner
}
