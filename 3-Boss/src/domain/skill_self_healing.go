package domain

type SelfHealingSkill struct {
	*AbstractSkill
}

func NewSelfHealingSkill(owner Role) *SelfHealingSkill {
	return &SelfHealingSkill{NewAbstractSkill(owner, 50, "自我治療")}
}

func (s *SelfHealingSkill) execute(targets []Role) {
	s.owner.addHp(150)
}

func (s *SelfHealingSkill) getTargets(allRolesOnBattle []Role) []Role {
	return nil
}
