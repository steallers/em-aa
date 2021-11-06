package pb

import (
	"context"

	"github.com/steallers/em-aa/servers/api/usecase"
)

type Handler struct {
	Manager usecase.ManagerUseCaseInterface
}

func (h Handler) UpdateEmployeeData(ctx context.Context, request *UpdateEmployeeDataRequest) (*UpdateEmployeeDataResponses, error) {
	panic("implement me")
}

func (h Handler) UpdateEmployeeProfilePicture(ctx context.Context, request *UpdateEmployeeProfilePictureRequest) (*UpdateEmployeeProfilePictureResponses, error) {
	panic("implement me")
}

func (h Handler) AddEmployeeContact(ctx context.Context, request *AddEmployeeContactRequest) (*AddEmployeeContactResponses, error) {
	panic("implement me")
}

func (h Handler) RemoveEmployeeContact(ctx context.Context, request *RemoveEmployeeContactRequest) (*RemoveEmployeeContactResponses, error) {
	panic("implement me")
}

func (h Handler) UpdateEmployeeContact(ctx context.Context, request *UpdateEmployeeContactRequest) (*UpdateEmployeeContactResponses, error) {
	panic("implement me")
}

func (h Handler) mustEmbedUnimplementedEmployeeServiceServer() {
	panic("implement me")
}
