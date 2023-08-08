package domain

type RoleState interface {
	beforeTakeTurn() (canGoOn bool)
	updateRound()
	attack(target Role, damageUnit int)
	setRole(r Role)
}

type AbstractRoleState struct {
	expiredRound int
	round        int
	role         Role
}

func NewAbstractRoleState(role Role) *AbstractRoleState {
	state := &AbstractRoleState{
		expiredRound: 3,
		round:        0,
		role:         role,
	}
	return state
}

func (s *AbstractRoleState) beforeTakeTurn() (canGoOn bool) {
	return true
}

func (s *AbstractRoleState) updateRound() {
	s.round += 1
	if s.round >= s.expiredRound {
		s.role.setState(NewNormalState(s.role))
	}
}

func (s *AbstractRoleState) attack(target Role, damageUnit int) {
	target.getDamaged(damageUnit)
}

func (s *AbstractRoleState) setRole(r Role) {
	s.role = r
}
