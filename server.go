package main

import (
	"fmt"
	"github.com/steallers/employee-management/api/model"
	"github.com/steallers/employee-management/api/repository"
	"github.com/steallers/employee-management/api/usecase"
	"github.com/steallers/employee-management/services/database"
	"gorm.io/gorm"
	"log"
	"net"
	"time"

	"github.com/steallers/employee-management/chat"
	server "github.com/steallers/employee-management/pb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8007))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//Building All Services
	databaseService := database.StartDatabaseService()
	databaseService.MigrateDatabase()

	//Building All Repository
	employeeRepository := repository.CreateEmployeeRepository(databaseService.GetConnectors())
	s := chat.Server{}
	//

	//Building All UseCase
	managerUseCase := usecase.CreateManagerUseCase(employeeRepository)
	newEmployeeModel := model.Employee{
		Model:      gorm.Model{},
		FirstName:  "Jupriano",
		MiddleName: "Abelio",
		LastName:   "Ginting",
		Contact:    nil,
		Gender:     "Female",
		BirthDate:  time.Now(),
	}

	eid, err := managerUseCase.AddNewEmployee(newEmployeeModel)
	log.Printf("employee id is :%d", eid)
	log.Printf("error is : %e",err)
	//Building All Servers
	f := server.Handler{
		//Register all Use Case here
		Manager: managerUseCase,
	}
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)
	server.RegisterEmployeeServiceServer(grpcServer, &f)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
