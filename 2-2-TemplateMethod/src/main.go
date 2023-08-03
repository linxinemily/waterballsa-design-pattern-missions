package main

import (
	"bigger-or-smaller-game/domain"
)

func main() {
	game := domain.NewUnoGame()

	game.Start()
}
