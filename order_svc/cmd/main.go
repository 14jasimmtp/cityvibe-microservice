package main

import (
	"log"
	"net"

	"github.com/14jasimmtp/cityvibe-microservice/order-svc/pb"
	"github.com/14jasimmtp/cityvibe-microservice/order-svc/pkg/config"
	"github.com/14jasimmtp/cityvibe-microservice/order-svc/pkg/db"
	"github.com/14jasimmtp/cityvibe-microservice/order-svc/pkg/repository"
	"github.com/14jasimmtp/cityvibe-microservice/order-svc/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error while loading config", err)
	}

	db := db.DbConnection(c.DB_URL)
	repo:=repository.NewRepo(db)
	svc:=service.NewService(repo)

	lis, err := net.Listen("tcp", c.PORT)
	if err != nil {
		log.Fatal("error while listening to port", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer,&svc)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("error while serving grpc server", err)
	}
}
