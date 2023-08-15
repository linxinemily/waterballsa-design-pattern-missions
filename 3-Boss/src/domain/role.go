package domain

import (
	"bufio"
	"fmt"
	"os"
)

type RoleImpl struct {
	Role
}

func (r *RoleImpl) takeTurn() {
	fmt.Fprintln(r.getRPG().getWriter(), fmt.Sprintf("輪到 %s。", r.getNameWithTroopIdAndStatus()))

	canGoOn := r.getState().beforeTakeTurn()
	if !canGoOn {
		return
	}

	var skill *SkillImpl
	for {
		skill = r.getSkillFromInput()
		if skill != nil && skill.getConsumeMp() <= r.getMp() {
			break
		} else {
			fmt.Fprintln(r.getRPG().getWriter(), "你缺乏 MP，不能進行此行動。")
		}
	}

	allRolesOnBattle := r.getRPG().getAllAliveRolesOnBattle()
	skill.takeAction(allRolesOnBattle)
}

type Role interface {
	getSkillFromInput() *SkillImpl                            // implement in concrete role
	getTargetsFromInput(candidates []Role, amount int) []Role // implement in concrete role
	attack(target Role, damageUnit int)
	getDamagedBy(damageUnit int, giver Role)
	isAllyOf(role Role) bool
	isEnemyOf(role Role) bool
	isAlive() bool
	setState(state RoleState)
	afterDied()
	getState() RoleState
	getMp() int
	decreaseMp(mp int)
	getRPG() *RPG
	getTroop() *Troop
	getId() int
	addHp(hp int)
	getStr() int
	getAfflictedObservers() map[int]AfflictedObserver
	setAfflictedObserver(roleId int, observer AfflictedObserver)
	setTroop(t *Troop)
	getHp() int
	getSkills() []*SkillImpl
	getName() string
	getNameWithTroopId() string
	getNameWithTroopIdAndStatus() string
	SetSkills(skills ...*SkillImpl)
	addSkill(skill *SkillImpl)
	getScanner() *bufio.Scanner
	setScanner(scanner *bufio.Scanner)
}

type AbstractRole struct {
	id                 int
	name               string
	HP                 int
	MP                 int
	STR                int
	state              RoleState
	rpg                *RPG
	troop              *Troop
	afflictedObservers map[int]AfflictedObserver
	skills             []*SkillImpl
	scanner            *bufio.Scanner
}

func NewAbstractRole(id int, name string, HP int, MP int, STR int, rpg *RPG) *AbstractRole {
	return &AbstractRole{
		id:                 id,
		name:               name,
		HP:                 HP,
		MP:                 MP,
		STR:                STR,
		rpg:                rpg,
		scanner:            bufio.NewScanner(os.Stdin),
		afflictedObservers: make(map[int]AfflictedObserver),
	}
}

func (r *AbstractRole) attack(target Role, damageUnit int) {
	r.state.attack(target, damageUnit)
}

func (r *AbstractRole) getDamagedBy(damageUnit int, giver Role) {
	r.HP -= damageUnit
	if giver != nil {
		fmt.Fprintf(r.getRPG().getWriter(), "%s 對 %s 造成 %d 點傷害。\n", giver.getNameWithTroopId(), r.getNameWithTroopId(), damageUnit)
	}
	if r.HP <= 0 {
		fmt.Fprintln(r.getRPG().getWriter(), r.getNameWithTroopId(), "死亡。")
		r.afterDied()
	}
}

func (r *AbstractRole) isAllyOf(role Role) bool {
	return r.troop.id == role.getTroop().id && r.id != role.getId()
}

func (r *AbstractRole) isEnemyOf(role Role) bool {
	return r.troop.id != role.getTroop().id
}

func (r *AbstractRole) setState(state RoleState) {
	r.state = state
}

func (r *AbstractRole) afterDied() {
	for _, observer := range r.afflictedObservers {
		observer.reward()
	}
}

func (r *AbstractRole) getState() RoleState {
	return r.state
}

func (r *AbstractRole) getMp() int {
	return r.MP
}

func (r *AbstractRole) decreaseMp(mp int) {
	r.MP -= mp
}

func (r *AbstractRole) isAlive() bool {
	return r.HP > 0
}

func (r *AbstractRole) getRPG() *RPG {
	return r.rpg
}

func (r *AbstractRole) getTroop() *Troop {
	return r.troop
}

func (r *AbstractRole) getId() int {
	return r.id
}

func (r *AbstractRole) addHp(hp int) {
	r.HP += hp
}

func (r *AbstractRole) getStr() int {
	return r.STR
}

func (r *AbstractRole) getAfflictedObservers() map[int]AfflictedObserver {
	return r.afflictedObservers
}

func (r *AbstractRole) setAfflictedObserver(roleId int, observer AfflictedObserver) {
	r.afflictedObservers[roleId] = observer
}

func (r *AbstractRole) setTroop(t *Troop) {
	r.troop = t
}

func (r *AbstractRole) getHp() int {
	return r.HP
}

func (r *AbstractRole) getSkills() []*SkillImpl {
	return r.skills
}

func (r *AbstractRole) getName() string {
	return r.name
}

func (r *AbstractRole) getNameWithTroopId() string {
	return fmt.Sprintf("[%d]%s", r.getTroop().getId(), r.getName())
}

func (r *AbstractRole) getNameWithTroopIdAndStatus() string {
	return fmt.Sprintf("%s (HP: %d, MP: %d, STR: %d, State: %s)",
		r.getNameWithTroopId(), r.getHp(), r.getMp(),
		r.getStr(), r.getState().getName())
}

func (r *AbstractRole) SetSkills(skills ...*SkillImpl) {
	r.skills = skills
}

func (r *AbstractRole) getScanner() *bufio.Scanner {
	return r.scanner
}

func (r *AbstractRole) setScanner(scanner *bufio.Scanner) {
	r.scanner = scanner
}

func (r *AbstractRole) addSkill(skill *SkillImpl) {
	r.skills = append(r.skills, skill)
}
