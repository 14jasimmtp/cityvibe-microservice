package main

import (
	"fmt"
	"log"
	"net"

	"github.com/14jasimmtp/cityvibe-microservice/cart_svc/pb"
	"github.com/14jasimmtp/cityvibe-microservice/cart_svc/pkg/config"
	"github.com/14jasimmtp/cityvibe-microservice/cart_svc/pkg/db"
	"github.com/14jasimmtp/cityvibe-microservice/cart_svc/pkg/repository"
	"github.com/14jasimmtp/cityvibe-microservice/cart_svc/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error while loading config", err)
	}

	db := db.DbConnection(c.DB_URL)

	repo := repository.NewRepo(db)
	svc := service.NewService(repo)

	lis, err := net.Listen("tcp", c.PORT)
	if err != nil {
		log.Fatal("error while listening to port", err)
	}
	fmt.Println("listening on port 4000")

	grpcServer := grpc.NewServer()
	pb.RegisterCartServiceServer(grpcServer, &svc)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("error while serving grpc server", err)
	}
}
