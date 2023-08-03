package uno_enum

type Number int

const (
	Zero Number = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Night
)

func (s Number) String() string {
	switch s {
	case Zero:
		return "0"
	case One:
		return "1"
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
	default:
		return "Unknown"
	}
}
