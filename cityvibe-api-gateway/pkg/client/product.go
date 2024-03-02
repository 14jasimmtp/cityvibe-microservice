package client

import (
	"fmt"

	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/config"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/product_service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitProductClient(c *config.Config) pb.ProductServiceClient {
	conn, err := grpc.Dial(c.ProductSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error connecting auth client : ", err)
	}

	client := pb.NewProductServiceClient(conn)

	return client

}
