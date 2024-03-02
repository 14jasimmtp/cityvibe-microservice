package routes

import (
	adminHandler "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/admin_service/handler"
	middlewares "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Admin(r *gin.Engine,handler *adminHandler.AdminHandler,mid *middlewares.AdminMiddleware){
	r.POST("admin/login",handler.AdminLogin)
	r.GET("admin/users",mid.AdminAuthMiddleware,handler.GetAllUsers)
	r.PUT("admin/blockuser",mid.AdminAuthMiddleware,handler.BlockUser)
	r.PUT("admin/unblockuser",mid.AdminAuthMiddleware,handler.UnBlockUser)
	r.GET("admin/Dashboard",mid.AdminAuthMiddleware,handler.DashBoard)
}