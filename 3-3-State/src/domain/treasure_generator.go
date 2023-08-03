package domain

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type TreasureGenerator struct {
	*TreasureType
	rateDivisor  int
	rateDividend int
	arr          []bool
	count        int
}

func NewTreasureGenerator(treasureType *TreasureType) *TreasureGenerator {
	s := strings.Split(treasureType.rate, "/")
	dividend, _ := strconv.Atoi(s[0]) // 被除數
	divisor, _ := strconv.Atoi(s[1])  // 除數

	generator := &TreasureGenerator{
		TreasureType: treasureType,
		rateDivisor:  divisor,
		rateDividend: dividend,
	}
	generator.init()
	return generator
}

func (generator *TreasureGenerator) init() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]bool, generator.rateDivisor)

	count := 0
	for count < generator.rateDividend {
		randNum := rand.Intn(generator.rateDivisor)
		if !arr[randNum] {
			arr[randNum] = true
			count += 1
		}
	}
	generator.count = 0
	generator.arr = arr
}

func (generator *TreasureGenerator) generate() *Treasure {
	if generator.count >= generator.rateDivisor {
		generator.init()
	}
	var treasure *Treasure
	if generator.arr[generator.count] {
		treasure = NewTreasure(generator.TreasureType)
	}
	generator.count += 1

	return treasure
}
