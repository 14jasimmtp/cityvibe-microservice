package main

import (
	"log"
	"net"

	"github.com/14jasimmtp/cityvibe-microservice/payment-svc/pb"
	"github.com/14jasimmtp/cityvibe-microservice/payment-svc/pkg/config"
	"github.com/14jasimmtp/cityvibe-microservice/payment-svc/pkg/db"
	"github.com/14jasimmtp/cityvibe-microservice/payment-svc/pkg/repository"
	"github.com/14jasimmtp/cityvibe-microservice/payment-svc/pkg/service"
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
	pb.RegisterPaymentServiceServer(grpcServer, svc.UnimplementedPaymentServiceServer)
	log.Printf("listening on port %s", c.PORT)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("error while serving grpc server", err)
	}
}
