package domain

import (
	"matching-system/domain/enum"
	"regexp"
	"strconv"
	"strings"
)

type Individual struct {
	id     int
	gender enum.Gender
	age    int
	intro  string
	habits string
	coord  string
}

func NewIndividual(id int, gender enum.Gender, age int, intro string, habits string, coord string) *Individual {
	return &Individual{
		id:     id,
		gender: gender,
		age:    validatedAge(age),
		intro:  validatedIntro(intro),
		habits: validatedHabits(habits),
		coord:  validatedCoord(coord),
	}
}

func validatedAge(age int) int {
	if age <= 18 {
		panic("age cannot under 18")
	}
	return age
}

func validatedIntro(intro string) string {
	if len(intro) > 200 {
		panic("intro cannot be more than 200 characters")
	}
	return intro
}

func validatedHabits(habits string) string {
	re := regexp.MustCompile(`^[^\s,]+(?:,\s[^\s,]+)*$`)
	if !re.Match([]byte(habits)) {
		panic("invalid habits format, valid example: \"A, B, C\"")
	}
	return habits
}

func validatedCoord(coord string) string {
	re := regexp.MustCompile(`\([0-9]*,[0-9]*\)`)
	if !re.Match([]byte(coord)) {
		panic("invalid coord format, valid example: \"(2, 4)\"")
	}
	return coord
}

func (i *Individual) GetXY() (x, y int) {
	s := strings.Split(i.coord, ",")

	_x := strings.TrimPrefix(s[0], "(")
	_y := strings.TrimSuffix(s[1], ")")

	intX, _ := strconv.Atoi(_x)
	intY, _ := strconv.Atoi(_y)

	return intX, intY
}

func (i *Individual) GetHabits() []string {
	return strings.Split(i.habits, ", ")
}

func (i *Individual) GetId() int {
	return i.id
}
