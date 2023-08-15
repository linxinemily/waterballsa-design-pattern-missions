package domain

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

var skillFactory = map[string]func(Role) Skill{
	"鼓舞": func(role Role) Skill {
		return NewCheerUpSkill(role)
	},
	"詛咒": func(role Role) Skill {
		return NewCurseSkill(role)
	},
	"下毒": func(role Role) Skill {
		return NewPoisonSkill(role)
	},
	"石化": func(role Role) Skill {
		return NewPetrochemicalSkill(role)
	},
	"自爆": func(role Role) Skill {
		return NewSelfExplosionSkill(role)
	},
	"自我治療": func(role Role) Skill {
		return NewSelfHealingSkill(role)
	},
	"召喚": func(role Role) Skill {
		return NewSummonSkill(role)
	},
	"水球": func(role Role) Skill {
		return NewWaterBallSkill(role)
	},
	"火球": func(role Role) Skill {
		return NewFireBallSkill(role)
	},
	"一拳攻擊": func(role Role) Skill {
		return NewOnePunchSkill(role)
	},
}

func TestRpg(t *testing.T) {

	testcasesDirPath := os.Getenv("TESTCASES_DIR_PATH")
	if testcasesDirPath == "" {
		panic("TESTCASES_DIR_PATH is not set")
	}

	table := []struct {
		filename string
	}{
		{"only-basic-attack"},
		{"cheerup"},
		{"curse"},
		{"poison"},
		{"petrochemical"},
		{"self-explosion"},
		{"self-healing"},
		{"summon"},
		{"waterball-and-fireball-1v2"},
		{"one-punch"},
	}

	for _, tt := range table {
		t.Run(tt.filename, func(t *testing.T) {

			inputFile := fmt.Sprintf("%s/%s.in", testcasesDirPath, tt.filename)

			file, err := os.Open(inputFile)
			if err != nil {
				return
			}
			defer file.Close()

			var buf bytes.Buffer

			rpg := NewRPG(&buf)
			troops := make(map[string]*Troop)
			var currentTroop *Troop
			var hero *RoleImpl
			lines := make([]string, 0)

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				line = strings.TrimSpace(line)
				lines = append(lines, line)
			}

			startHeroInputIndex := 0

			for index, line := range lines {
				if _, err = strconv.Atoi(line); err == nil { // 英雄開始輸入 input
					startHeroInputIndex = index
					break
				} else {
					if strings.HasPrefix(line, "#軍隊-") {
						strArr := strings.Split(line, "-")
						troopId := strArr[1]
						if strArr[2] == "開始" {
							currentTroop = rpg.CreateTroop()
							troops[troopId] = currentTroop
						} else { // 結束
							currentTroop = nil
						}
					} else if len(strings.Split(line, " ")) >= 4 { //建立角色
						strArr := strings.Split(line, " ")
						hp, _ := strconv.Atoi(strArr[1])
						mp, _ := strconv.Atoi(strArr[2])
						str, _ := strconv.Atoi(strArr[3])

						var role *RoleImpl
						if strArr[0] == "英雄" {
							hero = rpg.CreateHero("英雄", hp, mp, str)
							role = hero
							currentTroop.addRole(hero)
						} else {
							role = rpg.CreateAI(strArr[0], hp, mp, str)
							currentTroop.addRole(role)
						}

						for i := 4; i < len(strArr); i++ {
							skill, exist := skillFactory[strArr[i]]
							if !exist {
								panic(fmt.Sprintf("skill %s not exist", strArr[i]))
							}
							role.addSkill(&SkillImpl{skill(role)})
						}
					}
				}
			}

			heroReader := strings.NewReader(strings.Join(lines[startHeroInputIndex:], "\n"))
			hero.setScanner(bufio.NewScanner(heroReader))

			rpg.StartBattle(troops["1"], troops["2"])

			content, err := os.ReadFile(fmt.Sprintf("../testcases/%s.out", tt.filename))
			if err != nil {
				log.Fatal(err)
			}

			assert.Equal(t, string(content), buf.String())
		})
	}
}
