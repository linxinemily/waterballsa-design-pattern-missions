package enum

type Suit int

const (
	Club Suit = iota
	Diamond
	Heart
	Spade
)

func (s Suit) String() string {
	return map[Suit]string{
		Club:    "C",
		Diamond: "D",
		Heart:   "H",
		Spade:   "S",
	}[s]
}

func SuitStringToVal(str string) Suit {
	return map[string]Suit{
		"C": Club,
		"D": Diamond,
		"H": Heart,
		"S": Spade,
	}[str]
}
