package domain

type BasicSkill struct {
	*AbstractSkill
}

func NewBasicSkill(owner Role) *BasicSkill {
	return &BasicSkill{NewAbstractSkill(owner, 10)}
}

func (s *BasicSkill) execute(targets []Role) {
	s.owner.attack(targets[0], s.owner.getStr())
}

func (s *BasicSkill) getTargets(allRolesOnBattle []Role) (targets []Role) {
	candidates := make([]Role, 0)

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

func (s *BasicSkill) getName() string {
	return "普通攻擊"
}
