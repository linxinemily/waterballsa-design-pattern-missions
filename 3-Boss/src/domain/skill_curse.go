package domain

type CurseSkill struct {
	*AbstractSkill
}

func NewCurseSkill(owner *RoleImpl) *CurseSkill {
	return &CurseSkill{NewAbstractSkill(owner, 100, "詛咒")}
}

func (s *CurseSkill) execute(targets []*RoleImpl) {
	target := targets[0]
	//先檢查如果 taker 已被相同 giver 詛咒，就跳過
	if _, ok := target.getAfflictedObservers()[s.owner.getId()]; !ok {
		target.setAfflictedObserver(s.owner.getId(), NewCursed(s.owner, target))
	}
}

func (s *CurseSkill) getTargets(allRolesOnBattle []*RoleImpl) []*RoleImpl {

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
