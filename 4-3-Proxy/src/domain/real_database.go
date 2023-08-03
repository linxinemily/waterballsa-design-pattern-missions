package domain

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type RealDatabase struct {
	filename string
}

func NewRealDatabase(filename string) *RealDatabase {
	return &RealDatabase{filename}
}

func (rd *RealDatabase) GetEmployeeById(id int) Employee {
	file, err := os.Open(rd.filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; i < id; i++ {
		if !scanner.Scan() {
			return nil
		}
	}

	if scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		id, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}

		name := row[1]

		age, err := strconv.Atoi(row[2])
		if err != nil {
			panic(err)
		}

		var subordinateIds []int
		if len(row) > 3 {
			subordinateIdsStr := strings.Split(row[3], ",")
			for _, s := range subordinateIdsStr {
				id, _ := strconv.Atoi(s)
				subordinateIds = append(subordinateIds, id)
			}
		}

		return NewLazyLoadSubordinatesEmployeeProxy(id, name, age, subordinateIds, rd)

	} else {
		return nil
	}
}
