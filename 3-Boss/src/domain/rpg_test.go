package domain

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestRpg(t *testing.T) {
	var buf bytes.Buffer
	//
	//// 创建一个自定义的 log.Logger，将输出重定向到缓冲区
	//logger := log.New(&buf, "custom-log-prefix: ", log.LstdFlags)
	//
	//// 设置 log 输出到自定义 logger
	//log.SetOutput(logger.Writer())
	log.SetFlags(0)

	inputFile := "../testcases/only-basic-attack.in"

	file, err := os.Open(inputFile)
	if err != nil {
		return
	}
	defer file.Close()

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
			} else if len(strings.Split(line, " ")) == 4 { //建立角色
				strArr := strings.Split(line, " ")
				hp, _ := strconv.Atoi(strArr[1])
				mp, _ := strconv.Atoi(strArr[2])
				str, _ := strconv.Atoi(strArr[3])
				if strArr[0] == "英雄" {
					hero = rpg.CreateHero("英雄", hp, mp, str)
					currentTroop.addRole(hero)
				} else {
					currentTroop.addRole(rpg.CreateAI(strArr[0], hp, mp, str))
				}
			}
		}
	}

	heroReader := strings.NewReader(strings.Join(lines[startHeroInputIndex:], "\n"))
	hero.setScanner(bufio.NewScanner(heroReader))

	rpg.StartBattle(troops["1"], troops["2"])

	fmt.Println(buf.String())
}
