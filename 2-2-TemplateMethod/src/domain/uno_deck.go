package domain

import uno_enum "bigger-or-smaller-game/domain/enum/uno"

func NewUnoDeck() (d *Deck[UnoCard]) {
	cards := make([]*UnoCard, 40)
	var count int
	for color := uno_enum.Color(0); color < uno_enum.Green+1; color++ {
		for number := uno_enum.Number(0); number < uno_enum.Night+1; number++ {
			cards[count] = NewUnoCard(color, number)
			count++
		}
	}

	return &Deck[UnoCard]{
		cards: cards,
	}
}
