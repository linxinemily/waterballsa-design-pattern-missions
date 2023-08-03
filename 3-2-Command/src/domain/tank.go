package domain

import "fmt"

type Tank struct{ Name string }

func (t *Tank) moveForward() {
	fmt.Println("The tank has moved forward.")
}

func (t *Tank) moveBackward() {
	fmt.Println("The tank has moved backward.")
}
