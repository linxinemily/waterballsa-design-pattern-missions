package domain

type SelfExplosionSkill struct {
	*AbstractSkill
}

func NewSelfExplosionSkill(owner Role) *SelfExplosionSkill {
	return &SelfExplosionSkill{NewAbstractSkill(owner, 200)}
}

func (s *SelfExplosionSkill) execute(targets []Role) {
	s.owner.setHp(0)
	for _, target := range targets {
		s.owner.attack(target, 150)
	}
}

func (s *SelfExplosionSkill) getTargets(allRolesOnBattle []Role) []Role {
	return allRolesOnBattle
}

func (s *SelfExplosionSkill) getName() string {
	return "自爆"
}
