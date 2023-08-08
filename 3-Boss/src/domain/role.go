package domain

type RoleImpl struct {
	Role
}

func (r *RoleImpl) takeTurn() {
	canGoOn := r.getState().beforeTakeTurn()
	if !canGoOn {
		return
	}

	var skill *SkillImpl
	for {
		skill = r.getSkillFromInput()
		if skill != nil && skill.getConsumeMp() <= r.getMp() {
			break
		}
	}

	allRolesOnBattle := r.getRPG().getAllRolesOnBattle()
	skill.takeAction(allRolesOnBattle)
}

type Role interface {
	getSkillFromInput() *SkillImpl                            // implement in concrete role
	getTargetsFromInput(candidates []Role, amount int) []Role // implement in concrete role
	attack(target Role, damageUnit int)
	getDamaged(damageUnit int)
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
	setAfflictedObserver(roleId int, cursed *Cursed)
	setHp(i int)
	setTroop(t *Troop)
	getHp() int
	getSkills() []*SkillImpl
	getName() string
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
}

func NewAbstractRole(id int, name string, HP int, MP int, STR int, rpg *RPG) *AbstractRole {
	return &AbstractRole{
		id:   id,
		name: name,
		HP:   HP,
		MP:   MP,
		STR:  STR,
		rpg:  rpg,
	}
}

func (r *AbstractRole) attack(target Role, damageUnit int) {
	r.state.attack(target, damageUnit)
}

func (r *AbstractRole) getDamaged(damageUnit int) {
	r.HP -= damageUnit
	if r.HP <= 0 {
		r.afterDied()
	}
}

func (r *AbstractRole) isAllyOf(role Role) bool {
	return r.troop.id == role.getTroop().id
}

func (r *AbstractRole) isEnemyOf(role Role) bool {
	return r.troop.id != role.getTroop().id
}

func (r *AbstractRole) setState(state RoleState) {
	r.state = state
}

func (r *AbstractRole) afterDied() {
	r.troop.removeRoleById(r.id)
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

func (r *AbstractRole) setAfflictedObserver(roleId int, cursed *Cursed) {
	r.afflictedObservers[roleId] = cursed
}

func (r *AbstractRole) setHp(i int) {
	r.HP = i
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
