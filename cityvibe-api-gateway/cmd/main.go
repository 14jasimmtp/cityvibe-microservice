package main

import (
	"log"

	adminHandler "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/admin_service/handler"
	cartHandler "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/cart_service/handler"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/client"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/config"
	middlewares "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/middleware"
	orderHandler "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/order_service/handler"
	paymentHandler "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/payment_service/handler"
	product "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/product_service/handler"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/routes"
	UserHandler "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/user_service/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal("error while configuring env")
	}



	productc := client.InitProductClient(c)
	productHandler := product.NewProductClient(productc)
	
	userClient := client.InitUserClient(c)
	userHandler := UserHandler.NewUserClient(userClient)

	carClient:=client.InitCartClient(c)
	cartHandler:=cartHandler.NewCartHandler(carClient)

	adminClient:=client.InitAdminClient(c)
	adminHandler:=adminHandler.NewAdminClient(adminClient)

	paymantClient:=client.InitPaymentClient(c)
	paymentHandler:=paymentHandler.NewPaymentClient(paymantClient)

	orderClient:=client.InitOrderClient(c)
	orderHandler:=orderHandler.NewOrderClient(orderClient)
	UserMiddleware:=middlewares.NewUserMiddleware()
	adminMiddleware:=middlewares.NewAdminMiddleware()

	r := gin.Default()
	routes.UserRoute(r, userHandler,UserMiddleware)
	routes.Product(r, productHandler,adminMiddleware)
	routes.OrderRoute(r,orderHandler,UserMiddleware)
	routes.Admin(r,adminHandler,adminMiddleware)
	routes.Payment(r,paymentHandler)
	routes.Cart(r,cartHandler,UserMiddleware)
	r.LoadHTMLGlob("/home/jasim/Cityvibe-microservice/cityvibe-api-gateway/template/*")
	

	r.Run(c.PORT)
}
