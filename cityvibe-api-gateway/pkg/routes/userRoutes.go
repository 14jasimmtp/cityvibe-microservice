package routes

import (
	middlewares "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/middleware"
	UserHandler "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/user_service/handler"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine, handler UserHandler.UserHandler,mid *middlewares.UserMiddleware) {
	r.POST("/user/login", handler.UserLogin)
	r.POST("/user/signup", handler.UserSignUp)
	r.GET("/user/address",mid.UserAuthMiddleware, handler.ViewUserAddress)
	r.POST("/user/address",mid.UserAuthMiddleware, handler.AddNewAddressDetails)
}
