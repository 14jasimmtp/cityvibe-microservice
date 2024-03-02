package client

import (
	"fmt"

	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/config"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/order_service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitOrderClient(c *config.Config) pb.OrderServiceClient {
	conn, err := grpc.Dial(c.OrderSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error connecting auth client : ", err)
	}

	client := pb.NewOrderServiceClient(conn)

	return client

}
