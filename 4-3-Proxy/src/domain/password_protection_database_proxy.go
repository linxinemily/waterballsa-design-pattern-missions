package domain

import (
	"os"
)

type PasswordProtectionDatabaseProxy struct {
	realDatabase *RealDatabase
}

func NewPasswordProtectionDatabaseProxy(filename string) *PasswordProtectionDatabaseProxy {
	return &PasswordProtectionDatabaseProxy{NewRealDatabase(filename)}
}

func (p *PasswordProtectionDatabaseProxy) GetEmployeeById(id int) Employee {
	pw := os.Getenv("PASSWORD")
	if pw == "1qaz2wsx" {
		return p.realDatabase.GetEmployeeById(id)
	}
	panic("invalid password")
	return nil
}
