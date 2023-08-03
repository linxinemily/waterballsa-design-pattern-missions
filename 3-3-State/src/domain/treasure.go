package domain

type Treasure struct {
	*TreasureType
	*AbstractMapObject
}

func NewTreasure(treasureType *TreasureType) *Treasure {
	return &Treasure{
		TreasureType:      treasureType,
		AbstractMapObject: NewAbstractMapObject("🎁"),
	}
}
