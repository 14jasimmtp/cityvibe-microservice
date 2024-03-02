package main

import (
	"fmt"
	"log"
	"net"

	"github.com/14jasimmtp/cityvibe-microservice/user_svc/pb"
	"github.com/14jasimmtp/cityvibe-microservice/user_svc/pkg/config"
	"github.com/14jasimmtp/cityvibe-microservice/user_svc/pkg/db"
	"github.com/14jasimmtp/cityvibe-microservice/user_svc/pkg/repository"
	"github.com/14jasimmtp/cityvibe-microservice/user_svc/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error while loading config", err)
	}

	db := db.DbConnection(c.DB_URL)

	repo := repository.NewUserRepo(db)
	svc := service.NewService(repo, db)

	lis, err := net.Listen("tcp", c.PORT)
	if err != nil {
		log.Fatal("error while listening to port", err)
	}
	fmt.Println("listening on port 4001")

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, svc)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("error while serving grpc server", err)
	}
}
