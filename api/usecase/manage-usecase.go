package usecase

import (
	"errors"
	"fmt"
	. "github.com/steallers/employee-management/api/constant"
	"github.com/steallers/employee-management/api/model"
	"github.com/steallers/employee-management/api/repository"
	"gorm.io/gorm"
)

type ManagerUseCase struct {
	EmployeeRepo repository.EmployeeRepositoryInterface
}

func CreateManagerUseCase(employeeRepo repository.EmployeeRepositoryInterface) ManagerUseCaseInterface {
	return ManagerUseCase{
		EmployeeRepo: employeeRepo,
	}
}

func (b ManagerUseCase) AddNewEmployee(employee model.Employee) (employeeID uint, err error) {
	employee, err = b.EmployeeRepo.InsertEmployeeData(employee)
	if err != nil {
		return 0, err
	}
	return employee.ID, nil
}

func (b ManagerUseCase) RemoveEmployee(employeeID uint) error {
	// Validate if employee data available in our database
	_, err := b.EmployeeRepo.FindEmployeeWithEmployeeID(employeeID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return fmt.Errorf(EmployeeNotFound, employeeID)
		}
		return fmt.Errorf("%s %w", UnDescribeErrors, err)
	}
	//Add Validator User Task Here

	//End Validator User Task Here
	return b.EmployeeRepo.DeleteEmployeeData(employeeID)
}

func (b ManagerUseCase) UpdateEmployeeData() {

}
