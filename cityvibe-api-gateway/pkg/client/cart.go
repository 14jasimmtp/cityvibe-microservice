package client

import (
	"fmt"

	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/cart_service/pb"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitCartClient(c *config.Config) pb.CartServiceClient {
	conn, err := grpc.Dial(c.CartSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error connecting auth client : ", err)
	}

	client := pb.NewCartServiceClient(conn)

	return client

}
