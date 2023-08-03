package domain

type RealEmployee struct {
	id           int
	name         string
	age          int
	subordinates []Employee
}

func NewRealEmployee(id int, name string, age int, subordinateIds []int, database Database) *RealEmployee {
	re := &RealEmployee{id: id, name: name, age: age}

	re.subordinates = make([]Employee, 0)
	for _, subordinateId := range subordinateIds {
		re.subordinates = append(re.subordinates, database.GetEmployeeById(subordinateId))
	}

	return re
}

func (re *RealEmployee) GetSubordinates() []Employee {
	return re.subordinates
}
