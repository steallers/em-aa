package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/steallers/em-aa/servers/api/model"
	"github.com/steallers/em-aa/servers/api/repository"
	"github.com/steallers/em-aa/servers/api/usecase"
	"github.com/steallers/em-aa/servers/services/database"
	"gorm.io/gorm"

	server "github.com/steallers/em-aa/pb"
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
	log.Printf("error is : %e", err)
	//Building All Servers
	f := server.Handler{
		//Register all Use Case here
		Manager: managerUseCase,
	}
	grpcServer := grpc.NewServer()

	server.RegisterEmployeeServiceServer(grpcServer, &f)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
