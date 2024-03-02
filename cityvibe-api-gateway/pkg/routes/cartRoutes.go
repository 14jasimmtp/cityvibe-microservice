package routes

import (
	cartHandler "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/cart_service/handler"
	middlewares "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Cart(r *gin.Engine, handler cartHandler.CartHandler,mid *middlewares.UserMiddleware) {
	r.POST("/cart",mid.UserAuthMiddleware,handler.AddToCart)
	r.GET("/cart",mid.UserAuthMiddleware,handler.ViewCart)
	r.DELETE("/cart",mid.UserAuthMiddleware,handler.RemoveProductsFromCart)
}
