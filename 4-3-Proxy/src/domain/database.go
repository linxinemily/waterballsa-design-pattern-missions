package domain

type Database interface {
	GetEmployeeById(id int) Employee
}
