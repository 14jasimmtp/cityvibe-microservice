package main

import (
	"fmt"
	"log"
	"net"

	"github.com/14jasimmtp/cityvibe-microservice/product_svc/pb"
	"github.com/14jasimmtp/cityvibe-microservice/product_svc/pkg/config"
	"github.com/14jasimmtp/cityvibe-microservice/product_svc/pkg/db"
	"github.com/14jasimmtp/cityvibe-microservice/product_svc/pkg/repository"
	"github.com/14jasimmtp/cityvibe-microservice/product_svc/pkg/service"
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

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, svc)
	fmt.Println("listening on port 4002")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("error while serving grpc server", err)
	}
}
