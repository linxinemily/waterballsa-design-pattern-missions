package domain

import "fmt"

type RPG struct {
	battle         *Battle
	hero           *Hero
	roleIdCounter  int
	troopIdCounter int
}

func NewRPG() *RPG {
	return &RPG{}
}

func (r *RPG) StartBattle(troop1 *Troop, troop2 *Troop) {
	r.battle = NewBattle(troop1, troop2)

	for r.battle.takeRound() {
		r.battle.updateRound()
	}

	if r.hero.isAlive() {
		fmt.Println("你獲勝了！")
	} else {
		fmt.Println("你失敗了！")
	}
}

func (r *RPG) getAllRolesOnBattle() []*RoleImpl {
	var allRoles []*RoleImpl
	for _, role := range r.battle.troop1.roles {
		allRoles = append(allRoles, role)
	}
	for _, role := range r.battle.troop2.roles {
		allRoles = append(allRoles, role)
	}
	return allRoles
}

func (r *RPG) CreateSlime() *RoleImpl {
	slime := NewSlime(r.roleIdCounter, r)
	r.roleIdCounter++
	return &RoleImpl{Role: slime}
}

func (r *RPG) CreateHero(name string, HP int, MP int, STR int) *RoleImpl {
	if r.hero != nil {
		panic("Hero already exists")
	}
	hero := NewHero(r.roleIdCounter, name, HP, MP, STR, r)
	r.roleIdCounter++
	r.hero = hero
	return &RoleImpl{Role: hero}
}

func (r *RPG) CreateAI(name string, HP int, MP int, STR int) *RoleImpl {
	ai := NewAI(r.roleIdCounter, name, HP, MP, STR, r)
	r.roleIdCounter++
	return &RoleImpl{Role: ai}
}

func (r *RPG) CreateTroop(roles ...*RoleImpl) *Troop {
	troop := NewTroop(r.troopIdCounter, roles)
	r.troopIdCounter++
	return troop
}
