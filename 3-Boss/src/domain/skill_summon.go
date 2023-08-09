package domain

type SummonSkill struct {
	*AbstractSkill
}

func NewSummonSkill(owner *RoleImpl) *SummonSkill {
	return &SummonSkill{NewAbstractSkill(owner, 150, "召喚")}
}

func (s *SummonSkill) execute(targets []*RoleImpl) {
	slime := s.owner.getRPG().CreateSlime()
	s.owner.getTroop().addRole(slime)
}

func (s *SummonSkill) getTargets(allRolesOnBattle []*RoleImpl) []*RoleImpl {
	// no targets
	return nil
}
