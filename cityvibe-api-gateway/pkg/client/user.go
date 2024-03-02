package client

import (
	"fmt"

	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/config"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/user_service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitUserClient(c *config.Config) pb.UserServiceClient {
	conn, err := grpc.Dial(c.UserSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error connecting auth client : ", err)
	}

	client := pb.NewUserServiceClient(conn)

	return client

}
