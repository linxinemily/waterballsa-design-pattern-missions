package domain

type Employee interface {
	GetSubordinates() []Employee
}
