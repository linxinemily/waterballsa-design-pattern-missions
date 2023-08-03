package uno_enum

type Color int

const (
	Blue Color = iota
	Red
	Yellow
	Green
)

func (s Color) String() string {
	switch s {
	case Blue:
		return "BLUE"
	case Red:
		return "RED"
	case Yellow:
		return "YELLOW"
	case Green:
		return "GREEN"
	default:
		return "Unknown"
	}
}
