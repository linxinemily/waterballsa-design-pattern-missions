package domain

type SummonSkill struct {
	*AbstractSkill
}

func NewSummonSkill(owner Role) *SummonSkill {
	return &SummonSkill{NewAbstractSkill(owner, 150)}
}

func (s *SummonSkill) execute(targets []Role) {
	slime := s.owner.getRPG().CreateSlime()
	s.owner.getTroop().addRole(slime)
}

func (s *SummonSkill) getTargets(allRolesOnBattle []Role) []Role {
	// no targets
	return nil
}

func (s *SummonSkill) getName() string {
	return "召喚"
}
