package routes

import (
	middlewares "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/middleware"
	orderHandler "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/order_service/handler"
	"github.com/gin-gonic/gin"
)

func OrderRoute(r *gin.Engine, handler *orderHandler.OrderHandler,mid *middlewares.UserMiddleware) {
	r.POST("/order",mid.UserAuthMiddleware,handler.OrderFromCart)
	r.GET("/checkout",mid.UserAuthMiddleware,handler.ViewCheckOut)
	r.GET("/order",mid.UserAuthMiddleware,handler.ViewOrders)
	r.DELETE("/order",mid.UserAuthMiddleware,handler.CancelOrder)
}
