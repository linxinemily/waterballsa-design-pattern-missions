package main

import (
	"collision-detecting/domain"
)

func main() {
	world := domain.NewWorld(
		domain.NewICollision(domain.NewWaterAndFireCollision(
			domain.NewICollision(domain.NewWaterAndWaterCollision(
				domain.NewICollision(domain.NewFireAndFireCollision(
					domain.NewICollision(domain.NewHeroAndFireCollision(
						domain.NewICollision(domain.NewHeroAndWaterCollision(
							domain.NewICollision(domain.NewHeroAndHeroCollision(nil)),
						)),
					)),
				)),
			)),
		)),
	)

	for {
		x1, x2 := world.GetCoordinateFromUserInput()
		world.Move(*x1, *x2)
	}
}
