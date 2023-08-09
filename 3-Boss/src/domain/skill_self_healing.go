package domain

type SelfHealingSkill struct {
	*AbstractSkill
}

func NewSelfHealingSkill(owner *RoleImpl) *SelfHealingSkill {
	return &SelfHealingSkill{NewAbstractSkill(owner, 50, "自我治療")}
}

func (s *SelfHealingSkill) execute(targets []*RoleImpl) {
	targets[0].addHp(150)
}

func (s *SelfHealingSkill) getTargets(allRolesOnBattle []Role) []*RoleImpl {
	return []*RoleImpl{s.owner}
}
