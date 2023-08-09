package domain

type SelfExplosionSkill struct {
	*AbstractSkill
}

func NewSelfExplosionSkill(owner *RoleImpl) *SelfExplosionSkill {
	return &SelfExplosionSkill{NewAbstractSkill(owner, 200, "自爆")}
}

func (s *SelfExplosionSkill) execute(targets []*RoleImpl) {
	s.owner.setHp(0)
	for _, target := range targets {
		s.owner.attack(target, 150)
	}
}

func (s *SelfExplosionSkill) getTargets(allRolesOnBattle []*RoleImpl) []*RoleImpl {
	return allRolesOnBattle
}
