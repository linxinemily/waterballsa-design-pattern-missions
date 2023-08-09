package domain

import "fmt"

type BasicSkill struct {
	*AbstractSkill
}

func NewBasicSkill(owner *RoleImpl) *BasicSkill {
	return &BasicSkill{NewAbstractSkill(owner, 10, "普通攻擊")}
}

func (s *BasicSkill) execute(targets []*RoleImpl) {
	s.owner.attack(targets[0], s.owner.getStr())
}

func (s *BasicSkill) getTargets(allRolesOnBattle []*RoleImpl) (targets []*RoleImpl) {
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

func (s *BasicSkill) getCustomResultString(targets []*RoleImpl) string {
	return fmt.Sprintf("%s 攻擊 %s。\n", s.owner.getNameWithTroopId(), targets[0].getNameWithTroopId())
}
