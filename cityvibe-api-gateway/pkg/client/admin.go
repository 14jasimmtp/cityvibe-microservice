package client

import (
	"fmt"

	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/admin_service/pb"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitAdminClient(c *config.Config) pb.AdminServiceClient {
	conn, err := grpc.Dial(c.AdminSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error connecting auth client : ", err)
	}

	client := pb.NewAdminServiceClient(conn)

	return client

}
