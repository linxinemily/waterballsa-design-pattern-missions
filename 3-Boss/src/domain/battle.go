package domain

type Battle struct {
	round  int
	troop1 *Troop
	troop2 *Troop
}

func NewBattle(troop1, troop2 *Troop) *Battle {
	battle := &Battle{troop1: troop1, troop2: troop2}
	troop1.setBattle(battle)
	troop2.setBattle(battle)
	return battle
}

func (b *Battle) takeRound() (hasNextRound bool) {

	roles := make([]*RoleImpl, len(b.troop1.roles)+len(b.troop2.roles))
	roles = append(roles, b.troop1.roles...)
	roles = append(roles, b.troop2.roles...)

	for _, role := range roles {
		role.takeTurn()
		if !role.isAlive() || b.troop1.isAnnihilated() || b.troop2.isAnnihilated() {
			return false
		}
	}

	return true
}

func (b *Battle) updateRound() {
	for _, role := range b.troop1.roles {
		role.getState().updateRound()
	}
	for _, role := range b.troop2.roles {
		role.getState().updateRound()
	}
}
