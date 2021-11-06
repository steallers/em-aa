package repository

import "github.com/steallers/em-aa/servers/api/model"

type (
	EmployeeRepositoryInterface interface {
		InsertEmployeeData(employee model.Employee) (newEmployee model.Employee, err error) // Deprecated : Will be removed soon
		DeleteEmployeeData(employeeID uint) error                                           // Not Implemented : Working
		FindEmployeeWithEmployeeID(employeeID uint) (employee model.Employee, err error)
	}
)
