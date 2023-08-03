package domain

type ISprite interface {
	getCoordinate() int
}

type Sprite struct {
	coordinate int
}

func (s *Sprite) getCoordinate() int {
	return s.coordinate
}