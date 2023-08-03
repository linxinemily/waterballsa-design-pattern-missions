package enum

type Rank int

const (
	Three Rank = iota
	Four
	Five
	Six
	Seven
	Eight
	Night
	Ten
	J
	Q
	K
	A
	Two
)

func (r Rank) String() string {
	return map[Rank]string{
		Three: "3",
		Four:  "4",
		Five:  "5",
		Six:   "6",
		Seven: "7",
		Eight: "8",
		Night: "9",
		Ten:   "10",
		J:     "J",
		Q:     "Q",
		K:     "K",
		A:     "A",
		Two:   "2",
	}[r]
}

func RankStringToVal(str string) Rank {
	return map[string]Rank{
		"3":  Three,
		"4":  Four,
		"5":  Five,
		"6":  Six,
		"7":  Seven,
		"8":  Eight,
		"9":  Night,
		"10": Ten,
		"J":  J,
		"Q":  Q,
		"K":  K,
		"A":  A,
		"2":  Two,
	}[str]
}
