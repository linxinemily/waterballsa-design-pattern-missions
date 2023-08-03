package enum

type Gender int

const (
	Male Gender = iota
	Female
)

func (g Gender) String() string {
	switch g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Unknown"
	}
}
