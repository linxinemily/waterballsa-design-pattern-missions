package domain

type Hero struct {
	Sprite
	HP int
}

func NewHero() *Hero {
	return &Hero{HP: 30}
}

func (h *Hero) substractHP(val int) {
	h.HP -= val
	// fmt.Println("after substract HP:", h.HP)
}

func (h *Hero) addHP(val int) {
	h.HP += val
	// fmt.Println("after add HP:", h.HP)
}
