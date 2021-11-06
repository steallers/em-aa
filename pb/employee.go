package pb

import (
	"context"
	"fmt"
	"github.com/steallers/employee-management/api/model"
	"gorm.io/gorm"
	"log"
	"time"
)

func (h Handler) AddNewEmployee(ctx context.Context, in *AddNewEmployeeRequest) (*AddNewEmployeeResponses, error) {
	log.Printf(NewRequestMark, "add new employee", in.GetUserID(), in.String())
	birthDate, err := time.Parse("2006/01/02", in.BirthDate)
	if err != nil {
		log.Printf("invalid birthdate input: -> %e", err)
	}
	newEmployeeModel := model.Employee{
		Model:      gorm.Model{},
		FirstName:  in.FirstName,
		MiddleName: in.MiddleName,
		LastName:   in.LastName,
		Contact:    nil,
		Gender:     in.Gender,
		BirthDate:  birthDate,
	}
	employeeID, err := h.Manager.AddNewEmployee(newEmployeeModel)
	return &AddNewEmployeeResponses{Status: RequestSuccess, Message: fmt.Sprintf("new employee data added with id (%d)", employeeID)}, nil
}

func (h Handler) RemoveEmployee(ctx context.Context, in *RemoveEmployeeRequest) (*RemoveEmployeeResponse, error) {
	log.Printf(NewRequestMark, "remove employee", in.GetUserID(), in.String())
	err := h.Manager.RemoveEmployee(uint(in.GetEmployeeID()))
	if err != nil {
		return &RemoveEmployeeResponse{
			Status:  RequestFailed,
			Message: err.Error(),
		}, nil
	}
	return &RemoveEmployeeResponse{
		Status:  RequestSuccess,
		Message: fmt.Sprintf("user with id (%d) removed", in.GetEmployeeID()),
	}, nil
}
