package enum

type Suit int

const (
	Club Suit = iota
	Diamond
	Heart
	Spade
)

func (s Suit) String() string {
	switch s {
	case Club:
		return "♣︎"
	case Diamond:
		return "♦︎"
	case Heart:
		return "♥︎"
	case Spade:
		return "♠️"
	default:
		return "Unknown"
	}
}
