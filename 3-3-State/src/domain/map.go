package domain

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Map struct {
	objects [][]MapObject
	size    int
	game    *Game
	roles   []Role
}

func NewMap(size int, game *Game) *Map {
	m := make([][]MapObject, size)
	for i := 0; i < size; i++ {
		row := make([]MapObject, size)
		m[i] = row
	}

	return &Map{
		objects: m,
		size:    size,
		game:    game,
	}
}

func (m *Map) getRandomVacancy() (row, col int) {
	rand.Seed(time.Now().UnixNano())
	for {
		row = rand.Intn(m.Size())
		col = rand.Intn(m.Size())
		if obj, err := m.getObjectAt(row, col); err == nil && obj == nil {
			return row, col
		}
	}
}

func (m *Map) moveRoleToRandomPosition(object Role) {
	row, col := m.getRandomVacancy()
	m.moveRoleTo(object, row, col)
}

func (m *Map) putRoleInRandomPosition(object Role) {
	row, col := m.getRandomVacancy()
	m.putRoleAt(object, row, col)
}

func (m *Map) putObjectInRandomPosition(object MapObject) {
	row, col := m.getRandomVacancy()
	m.putObjectAt(object, row, col)
}

func (m *Map) putRoleAt(role Role, row, col int) error {
	m.roles = append(m.roles, role)
	err := m.putObjectAt(role, row, col)
	return err
}

func (m *Map) putObjectAt(object MapObject, row, col int) error {
	size := len(m.objects)
	if row < 0 || row >= size {
		return errors.New("not valid x")
	}

	if col < 0 || col >= size {
		return errors.New("not valid x")
	}

	m.objects[row][col] = object
	object.setPosition(row, col)

	return nil
}

func (m *Map) getObjectAt(row int, col int) (MapObject, error) {
	size := len(m.objects)
	if row < 0 || row >= size {
		return nil, errors.New("not valid x")
	}

	if col < 0 || col >= size {
		return nil, errors.New("not valid x")
	}

	return m.objects[row][col], nil
}

func (m *Map) Size() int {
	return m.size
}

func (m *Map) print() {
	fmt.Printf("MonsterCount: %d\n", m.monstersCount())
	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			fmt.Print(" === ")
		}
		fmt.Println()

		fmt.Print("|")
		for j := 0; j < m.size; j++ {
			if o := m.objects[i][j]; o != nil {
				fmt.Printf(" %s |", o.Symbol())
			} else {
				fmt.Print("    |")
			}
		}
		fmt.Println()
	}
	for j := 0; j < m.size; j++ {
		fmt.Print(" === ")
	}
	fmt.Println()
}

func (m *Map) moveRoleTo(role Role, toRow, toCol int) error {
	fromRow := role.Row()
	fromCol := role.Col()
	size := len(m.objects)
	if toRow < 0 || toRow >= size {
		return errors.New("not valid x")
	}

	if toCol < 0 || toCol >= size {
		return errors.New("not valid x")
	}

	m.objects[toRow][toCol] = role
	role.setPosition(toRow, toCol)

	m.objects[fromRow][fromCol] = nil

	return nil
}

func (m *Map) removeObject(role MapObject) {
	removed := m.objects[role.Row()][role.Col()]

	for i := 0; i < len(m.roles); i++ {
		if m.roles[i] == removed {
			m.roles = append(m.roles[:i], m.roles[i+1:]...)
		}
	}

	m.objects[role.Row()][role.Col()] = nil
}

func (m *Map) monstersCount() int {
	monsterCount := 0
	for _, role := range m.roles {
		if _, isM := role.(*Monster); isM {
			monsterCount += 1
		}
	}
	return monsterCount
}
