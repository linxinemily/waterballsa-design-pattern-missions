package domain

import (
	"bufio"
	"fmt"
	"math/rand"
	"time"

	"os"
	"strconv"
)

type World struct {
	length    int
	collision *ICollision
	sprites   []ISprite
}

func NewWorld(c *ICollision) *World {
	world := &World{collision: c, length: 30}
	world.sprites = world.initSprites()
	world.collision.setWorld(world)
	// fmt.Println(world.sprites)
	return world
}

func (w *World) initSprites() []ISprite {

	sprites := make([]ISprite, w.length)

	rand.Seed(time.Now().UnixNano())
	temp := []ISprite{NewHero(), NewHero(), NewHero(), NewHero(), NewWater(), NewWater(), NewWater(), NewFire(), NewFire(), NewFire()}

	for i := 0; i < len(temp); i++ {
		j := rand.Intn(w.length)

		for sprites[j] != nil {
			j = rand.Intn(w.length)
		}

		sprites[j] = temp[i]
	}

	return sprites
}

func (w *World) Move(x1 int, x2 int) {
	s1 := w.sprites[x1]
	s2 := w.sprites[x2]
	if s1 != nil && s2 != nil {
		w.collision.handle(x1, x2)
	} else if s1 != nil && s2 == nil {
		w.sprites[x2] = s1
		w.sprites[x1] = nil
	} else {
		fmt.Println("Nothing happen")
	}
}

func (w *World) GetCoordinateFromUserInput() (x1, x2 *int) {

	for x1 == nil && x2 == nil {
		var input1, input2 int

		fmt.Println("Enter x1:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input1, err := strconv.Atoi(scanner.Text())
		fmt.Println("Enter x2:")
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}

		if input1 >= len(w.sprites) || input2 >= len(w.sprites) {
			fmt.Printf("input should >= 0, < %d \n", w.length)
			continue
		}

		if input1 == input2 {
			fmt.Println("x1 and x2 cannot be equal")
			continue
		}

		x1, x2 = &input1, &input2
	}

	return x1, x2
}

func (w *World) setSpriteInPosition(idx int, sprite ISprite) {
	w.sprites[idx] = sprite
}

func (w *World) removeSpriteInPosition(idx int) {
	w.sprites[idx] = nil
}

func (w *World) getSpriteInPosition(idx int) ISprite {
	return w.sprites[idx]
}
