package domain

import "fmt"

type Telecom struct{}

func (t *Telecom) connect() {
	fmt.Println("The telecom has been turned on.")
}

func (t *Telecom) disconnect() {
	fmt.Println("The telecom has been turned off.")
}
