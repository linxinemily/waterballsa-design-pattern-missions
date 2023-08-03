package enum

type RoleDirection int

const (
	Top RoleDirection = iota
	Down
	Left
	Right
)

var AllDirections = []RoleDirection{Top, Down, Left, Right}

func (direction RoleDirection) String() string {
	return map[RoleDirection]string{
		Top:   "⬆️",
		Down:  "⬇️",
		Left:  "⬅️",
		Right: "➡️",
	}[direction]
}

type Button struct {
	Key  string
	Text string
}

func (direction RoleDirection) Button() Button {
	return map[RoleDirection]Button{
		Top:   {Key: "w", Text: "往上移動"},
		Down:  {Key: "s", Text: "往下移動"},
		Left:  {Key: "a", Text: "往左移動"},
		Right: {Key: "d", Text: "往右移動"},
	}[direction]
}

func GetMovementByDirection(originRow, originCol int, direction RoleDirection) (row, col int) {
	rowCol := map[RoleDirection][2]int{
		Top:   {originRow - 1, originCol},
		Down:  {originRow + 1, originCol},
		Left:  {originRow, originCol - 1},
		Right: {originRow, originCol + 1},
	}[direction]
	return rowCol[0], rowCol[1]
}
