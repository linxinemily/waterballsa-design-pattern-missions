package main

import (
	"C3M2H1/domain"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	tank := &domain.Tank{}
	telecom := &domain.Telecom{}
	ctr := domain.NewMainController()

	commandList := []domain.Command{
		domain.NewTankMoveForwardCommand(tank),
		domain.NewTankMoveBackwardCommand(tank),
		domain.NewTelecomConnectCommand(telecom),
		domain.NewTelecomDisconnectCommand(telecom),
		domain.NewResetMainControllerKeyboardCommand(ctr),
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			//快捷鍵設置
			fmt.Print("設置巨集指令 (y/n): ")
			scanner.Scan()
			setMacro, valid := map[string]bool{"y": true, "n": false}[scanner.Text()]
			fmt.Print("Key: ")
			scanner.Scan()
			key := scanner.Text()
			var bindCmd *domain.CommandWrapper

			if valid && setMacro {
				fmt.Printf("要將哪些指令設置成快捷鍵 %s 的巨集（輸入多個數字，以空白隔開）:\n", key)
				for i, wrapper := range commandList {
					fmt.Printf("(%d) %s\n", i, wrapper.Name())
				}
				scanner.Scan()
				cmdIds := strings.Split(scanner.Text(), " ") // [0 , 1]

				i := 0
				var head *domain.CommandWrapper
				var current *domain.CommandWrapper

				for {
					if i >= len(cmdIds) {
						break
					}

					if head == nil {
						intVar, _ := strconv.Atoi(cmdIds[i])
						head = domain.NewICommandWrapper(commandList[intVar])
						current = head
						i += 1
					}

					intVar, _ := strconv.Atoi(cmdIds[i])
					temp := domain.NewICommandWrapper(commandList[intVar])
					current.SetNext(temp)
					current = temp
					i += 1
				}
				bindCmd = head
			} else {
				fmt.Printf("要將哪一道指令設置到快捷鍵 %s 上: \n", key)
				for i, wrapper := range commandList {
					fmt.Printf("(%d) %s\n", i, wrapper.Name())
				}
				scanner.Scan()
				intVar, _ := strconv.Atoi(scanner.Text())
				bindCmd = domain.NewICommandWrapper(commandList[intVar])
			}

			ctr.Bind([]byte(key)[0], bindCmd)

		case "2":
			//undo
			ctr.Undo()
		case "3":
			//redo
			ctr.Redo()
		default:
			//按下按鍵
			err := ctr.Press([]byte(option)[0])
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

}
