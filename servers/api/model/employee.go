package model

import (
	"gorm.io/gorm"
	"time"
)

type Employee struct {
	gorm.Model
	FirstName  string
	MiddleName string
	LastName   string
	Contact    []EmployeeContact
	Gender     string //male and female
	BirthDate  time.Time
}

func (Employee) TableName() string {
	return "maindb.employee"
}

type EmployeeContact struct {
	ID          string
	EmployeeID  uint
	ContactType string
	Value       string
}

func (EmployeeContact) TableName() string {
	return "maindb.employee_contact"
}
