package domain

type LazyLoadSubordinatesEmployeeProxy struct {
	id             int
	name           string
	age            int
	subordinateIds []int
	realEmployee   *RealEmployee
	database       Database
}

func NewLazyLoadSubordinatesEmployeeProxy(id int, name string, age int, subordinateIds []int, database Database) *LazyLoadSubordinatesEmployeeProxy {
	return &LazyLoadSubordinatesEmployeeProxy{
		id, name, age, subordinateIds, nil, database,
	}
}

func (re *LazyLoadSubordinatesEmployeeProxy) GetSubordinates() []Employee {
	if re.realEmployee == nil {
		re.realEmployee = NewRealEmployee(re.id, re.name, re.age, re.subordinateIds, re.database)
	}
	return re.realEmployee.subordinates
}
