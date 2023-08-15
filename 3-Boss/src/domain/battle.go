package domain

type Battle struct {
	round  int
	troop1 *Troop
	troop2 *Troop
	hero   *Hero
}

func NewBattle(troop1, troop2 *Troop, hero *Hero) *Battle {
	battle := &Battle{troop1: troop1, troop2: troop2, hero: hero}
	troop1.setBattle(battle)
	troop2.setBattle(battle)
	return battle
}

func (b *Battle) takeRound() (hasNextRound bool) {
	for _, troop := range []*Troop{b.troop1, b.troop2} {
		for i := 0; i < len(troop.roles); i++ {
			role := troop.roles[i]
			if !role.isAlive() {
				continue
			}
			role.takeTurn()
			if !b.hero.isAlive() || b.troop1.isAnnihilated() || b.troop2.isAnnihilated() {
				return false
			}
		}
	}

	return true
}

func (b *Battle) updateRound() {
	for _, troop := range []*Troop{b.troop1, b.troop2} {
		for _, role := range troop.roles {
			role.getState().updateRound()
		}
	}
}
