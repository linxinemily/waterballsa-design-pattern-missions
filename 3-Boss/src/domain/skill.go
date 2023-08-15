package domain

import (
	"fmt"
	"strings"
)

type SkillImpl struct {
	Skill
}

func NewSkillImpl(skill Skill) *SkillImpl {
	return &SkillImpl{skill}
}

func (s *SkillImpl) takeAction(allRolesOnBattle []Role) {
	targets := s.getTargets(allRolesOnBattle)

	s.getOwner().decreaseMp(s.getConsumeMp())

	s.printResult(targets)
	s.execute(targets)
}

func (s *SkillImpl) printResult(targets []Role) {
	if len(targets) == 0 {
		fmt.Fprintf(s.getOwner().getRPG().getWriter(), "%s 使用了 %s。\n", s.getOwner().getNameWithTroopId(), s.getName())
	} else if _, ok := s.Skill.(*BasicSkill); !ok {
		targetNames := make([]string, 0)
		for _, target := range targets {
			targetNames = append(targetNames, target.getNameWithTroopId())
		}
		fmt.Fprintf(s.getOwner().getRPG().getWriter(), "%s 對 %s 使用了 %s。\n", s.getOwner().getNameWithTroopId(), strings.Join(targetNames, ", "), s.getName())
	}
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
	name      string
}

func NewAbstractSkill(owner Role, consumeMp int, name string) *AbstractSkill {
	return &AbstractSkill{
		owner:     owner,
		consumeMp: consumeMp,
		name:      name,
	}
}

func (s *AbstractSkill) getOwner() Role {
	return s.owner
}

func (s *AbstractSkill) getConsumeMp() int {
	return s.consumeMp
}

func (s *AbstractSkill) getName() string {
	return s.name
}
