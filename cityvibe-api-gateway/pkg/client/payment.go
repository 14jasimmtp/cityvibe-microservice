package client

import (
	"fmt"

	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/config"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/payment_service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitPaymentClient(c *config.Config) pb.PaymentServiceClient {
	conn, err := grpc.Dial(c.PaymentSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error connecting auth client : ", err)
	}

	client := pb.NewPaymentServiceClient(conn)

	return client

}
