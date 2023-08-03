package domain

type MapObject interface {
	Symbol() string
	Row() int
	Col() int
	setPosition(row, col int)
}

type AbstractMapObject struct {
	row    int
	col    int
	symbol string
	m      *Map
}

func NewAbstractMapObject(symbol string) *AbstractMapObject {
	return &AbstractMapObject{
		symbol: symbol,
	}
}

func (object *AbstractMapObject) Symbol() string {
	return object.symbol
}

func (object *AbstractMapObject) Row() int {
	return object.row
}

func (object *AbstractMapObject) Col() int {
	return object.col
}

func (object *AbstractMapObject) setPosition(row, col int) {
	object.row = row
	object.col = col
}
