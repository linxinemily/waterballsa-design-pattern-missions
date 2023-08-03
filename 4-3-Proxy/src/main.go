package main

import (
	"C4M3H1/domain"
	"fmt"
)

func main() {
	db := domain.NewPasswordProtectionDatabaseProxy("employee_data.txt")
	employee := db.GetEmployeeById(4)
	fmt.Println(employee)
	for _, e := range employee.GetSubordinates() {
		fmt.Println(e)
	}
}
