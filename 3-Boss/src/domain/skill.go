package domain

import (
	"fmt"
	"strings"
)

type SkillImpl struct {
	Skill
}

func (s *SkillImpl) takeAction(allRolesOnBattle []*RoleImpl) {
	targets := s.getTargets(allRolesOnBattle)

	s.getOwner().decreaseMp(s.getConsumeMp())

	s.execute(targets)
}

func (s *SkillImpl) printResult(targets []*RoleImpl) {
	resultString := s.getCustomResultString(targets)
	if resultString == "" { // default result string
		if len(targets) == 0 {
			fmt.Printf("%s 使用了 %s。\n", s.getOwner().getNameWithTroopId(), s.getName())
		} else {
			targetNames := make([]string, 0)
			for _, target := range targets {
				targetNames = append(targetNames, target.getNameWithTroopId())
			}
			fmt.Printf("%s 對 %s 使用了 %s。\n", s.getOwner().getNameWithTroopId(), strings.Join(targetNames, ","), s.getName())
		}
	} else {
		fmt.Println(s.getCustomResultString(targets))
	}
}

type Skill interface {
	getTargets(allRolesOnBattle []*RoleImpl) (targets []*RoleImpl)
	getConsumeMp() int
	execute(targets []*RoleImpl)
	getOwner() *RoleImpl
	getName() string
	getCustomResultString(targets []*RoleImpl) string
}

type AbstractSkill struct {
	owner     *RoleImpl
	consumeMp int
	name      string
}

func NewAbstractSkill(owner *RoleImpl, consumeMp int, name string) *AbstractSkill {
	return &AbstractSkill{
		owner:     owner,
		consumeMp: consumeMp,
		name:      name,
	}
}

func (s *AbstractSkill) getOwner() *RoleImpl {
	return s.owner
}

func (s *AbstractSkill) getConsumeMp() int {
	return s.consumeMp
}

func (s *AbstractSkill) getName() string {
	return s.name
}

func (s *AbstractSkill) getCustomResultString([]*RoleImpl) string {
	return ""
}
