package repository

import (
	"github.com/steallers/employee-management/api/model"
	"gorm.io/gorm"
	"log"
)

type EmployeeRepository struct {
	database *gorm.DB
}

func CreateEmployeeRepository(connectors *gorm.DB) EmployeeRepositoryInterface {
	return EmployeeRepository{
		connectors,
	}
}
// Deprecated: InsertEmployeeData
func (rep EmployeeRepository) InsertEmployeeData(employee model.Employee) (newEmployee model.Employee, err error) {
	log.Printf("employee %+v", &employee)
	err = rep.database.Create(&employee).Error
	log.Printf("employee %+v", &employee)
	newEmployee = employee
	return
}

func (rep EmployeeRepository) DeleteEmployeeData(employeeID uint) error {
	return rep.database.Delete(&model.Employee{}, employeeID).Error
}

func (rep EmployeeRepository) FindEmployeeWithEmployeeID(employeeID uint) (employee model.Employee, err error) {
	err = rep.database.Where("id = ?", employeeID).Take(&employee).Error
	return
}
