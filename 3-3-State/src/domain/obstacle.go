package domain

type Obstacle struct {
	*AbstractMapObject
}

func NewObstacle() *Obstacle {
	return &Obstacle{
		AbstractMapObject: NewAbstractMapObject("🚧"),
	}
}
