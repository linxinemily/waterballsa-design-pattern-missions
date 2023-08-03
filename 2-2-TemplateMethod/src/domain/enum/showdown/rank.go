package showdown_enum

type Rank int

const (
	Two Rank = iota
	Three
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
)

func (s Rank) String() string {
	switch s {
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Night:
		return "9"
	case Ten:
		return "10"
	case J:
		return "J"
	case Q:
		return "Q"
	case K:
		return "K"
	case A:
		return "A"
	default:
		return "Unknown"
	}
}
