package domain

type Troop struct {
	id     int
	roles  []*RoleImpl
	battle *Battle
}

func (t *Troop) removeRoleById(id int) {
	for i, role := range t.roles {
		if role.getId() == id {
			t.roles = append(t.roles[:i], t.roles[i+1:]...)
			break
		}
	}
}

func (t *Troop) isAnnihilated() bool {
	for _, role := range t.roles {
		if role.isAlive() {
			return false
		}
	}
	return true
}

func (t *Troop) setBattle(battle *Battle) {
	t.battle = battle
}

func (t *Troop) getBattle() *Battle {
	return t.battle
}

func (t *Troop) addRole(role *RoleImpl) {
	t.roles = append(t.roles, role)
	role.setTroop(t)
}

func (t *Troop) getId() int {
	return t.id
}

func NewTroop(id int, roles []*RoleImpl) *Troop {
	return &Troop{id: id, roles: roles}
}
