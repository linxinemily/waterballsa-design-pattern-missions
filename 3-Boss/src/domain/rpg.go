package domain

import (
	"fmt"
	"io"
)

type RPG struct {
	battle         *Battle
	hero           *Hero
	roleIdCounter  int
	troopIdCounter int
	writer         io.Writer
}

func NewRPG(writer io.Writer) *RPG {
	return &RPG{roleIdCounter: 1, troopIdCounter: 1, writer: writer}
}

func (r *RPG) StartBattle(troop1 *Troop, troop2 *Troop) {
	r.battle = NewBattle(troop1, troop2, r.hero)

	for r.battle.takeRound() {
		r.battle.updateRound()
	}

	if r.hero.isAlive() {
		fmt.Fprintln(r.getWriter(), "你獲勝了！")
	} else {
		fmt.Fprintln(r.getWriter(), "你失敗了！")
	}
}

func (r *RPG) getAllAliveRolesOnBattle() []Role {
	var allRoles []Role
	for _, troop := range []*Troop{r.battle.troop1, r.battle.troop2} {
		for _, role := range troop.roles {
			if role.isAlive() {
				allRoles = append(allRoles, role)
			}
		}
	}
	return allRoles
}

func (r *RPG) CreateSlime() *RoleImpl {
	slime := NewSlime(r.roleIdCounter, r)
	slime.addSkill(&SkillImpl{NewBasicSkill(slime)})
	r.roleIdCounter++
	return &RoleImpl{Role: slime}
}

func (r *RPG) CreateHero(name string, HP int, MP int, STR int) *RoleImpl {
	if r.hero != nil {
		panic("Hero already exists")
	}
	hero := NewHero(r.roleIdCounter, name, HP, MP, STR, r)
	hero.addSkill(&SkillImpl{NewBasicSkill(hero)})
	r.roleIdCounter++
	r.hero = hero
	return &RoleImpl{Role: hero}
}

func (r *RPG) CreateAI(name string, HP int, MP int, STR int) *RoleImpl {
	ai := NewAI(r.roleIdCounter, name, HP, MP, STR, r)
	ai.addSkill(&SkillImpl{NewBasicSkill(ai)})
	r.roleIdCounter++
	return &RoleImpl{Role: ai}
}

func (r *RPG) CreateTroop(roles ...*RoleImpl) *Troop {
	troop := NewTroop(r.troopIdCounter, roles)
	for _, role := range roles {
		role.setTroop(troop)
	}
	r.troopIdCounter++
	return troop
}

func (r *RPG) getWriter() io.Writer {
	return r.writer
}
