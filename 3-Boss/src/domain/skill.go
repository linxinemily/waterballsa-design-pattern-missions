package domain

type SkillImpl struct {
	Skill
}

func (s *SkillImpl) takeAction(allRolesOnBattle []Role) {
	targets := s.getTargets(allRolesOnBattle)

	s.getOwner().decreaseMp(s.getConsumeMp())

	s.execute(targets)
}

type Skill interface {
	getTargets(allRolesOnBattle []Role) (targets []Role)
	getConsumeMp() int
	execute(targets []Role)
	getOwner() Role
	getName() string
}

type AbstractSkill struct {
	owner     Role
	consumeMp int
}

func NewAbstractSkill(owner Role, consumeMp int) *AbstractSkill {
	return &AbstractSkill{
		owner:     owner,
		consumeMp: consumeMp,
	}
}

func (s *AbstractSkill) getOwner() Role {
	return s.owner
}

func (s *AbstractSkill) getConsumeMp() int {
	return s.consumeMp
}
