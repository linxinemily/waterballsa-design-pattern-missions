package domain

type RoleState interface {
	beforeTakeTurn() (canGoOn bool)
	updateRound()
	attack(target Role, damageUnit int)
	setRole(r Role)
	getName() string
}

type AbstractRoleState struct {
	expiredRound int
	round        int
	role         Role
	name         string
}

func NewAbstractRoleState(role Role, name string) *AbstractRoleState {
	state := &AbstractRoleState{
		expiredRound: 3,
		round:        0,
		role:         role,
		name:         name,
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
	target.getDamagedBy(damageUnit, s.role)
}

func (s *AbstractRoleState) setRole(r Role) {
	s.role = r
}

func (s *AbstractRoleState) getName() string {
	return s.name
}
