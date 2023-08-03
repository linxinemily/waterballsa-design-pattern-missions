package domain

import (
	"C3M2H1/stack"
	"errors"
	"fmt"
)

type MainController struct {
	bindMap map[byte]*CommandWrapper
	s1      *stack.Stack[*CommandWrapper]
	s2      *stack.Stack[*CommandWrapper]
}

func NewMainController() *MainController {
	return &MainController{
		bindMap: getInitialBindMap(),
		s1:      stack.NewStack[*CommandWrapper](),
		s2:      stack.NewStack[*CommandWrapper](),
	}
}

func (controller *MainController) Bind(key byte, command *CommandWrapper) {
	if _, exists := controller.bindMap[key]; exists {
		controller.bindMap[key] = command
	}
	printBindMap(controller.bindMap)
}

func (controller *MainController) Reset() {
	controller.bindMap = getInitialBindMap()
	printBindMap(controller.bindMap)
}

func (controller *MainController) Press(key byte) error {
	if _, exists := controller.bindMap[key]; exists {
		controller.s1.Push(controller.bindMap[key])
		controller.bindMap[key].execute()
		controller.s2 = stack.NewStack[*CommandWrapper]()
		printBindMap(controller.bindMap)
		return nil
	} else {
		return errors.New("invalid key")
	}
}

func (controller *MainController) Undo() {
	if !controller.s1.IsEmpty() {
		previousCmd := controller.s1.Pop()
		previousCmd.undo()
		controller.s2.Push(previousCmd)
		printBindMap(controller.bindMap)
	}
}

func (controller *MainController) Redo() {
	if !controller.s2.IsEmpty() {
		nextCmd := controller.s2.Pop()
		nextCmd.execute()
		controller.s1.Push(nextCmd)
		printBindMap(controller.bindMap)
	}
}

func getInitialBindMap() map[byte]*CommandWrapper {
	validKeys := []byte{
		'a', 'b', 'c', 'd', 'e',
		'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o',
		'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y',
		'z',
	}

	bindMap := make(map[byte]*CommandWrapper)

	for _, key := range validKeys {
		bindMap[key] = nil
	}

	return bindMap
}

func printBindMap(bindMap map[byte]*CommandWrapper) {
	for key, val := range bindMap {
		if val != nil {
			fmt.Printf("%s: %s", string(key), val.Name())
			current := val.next
			for current != nil {
				fmt.Printf(" & %s", current.Name())
				current = current.next
			}
			fmt.Println()
		}
	}
}
