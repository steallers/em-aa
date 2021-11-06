package usecase

import "github.com/steallers/em-aa/servers/api/model"

type (
	ManagerUseCaseInterface interface {
		AddNewEmployee(employee model.Employee) (employeeID uint, err error)
		RemoveEmployee(employeeID uint) error
	}
)
