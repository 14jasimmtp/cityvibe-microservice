package routes

import (
	middlewares "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/middleware"
	product "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/product_service/handler"
	"github.com/gin-gonic/gin"
)

func Product(r *gin.Engine, handler product.ProductHandler, mid *middlewares.AdminMiddleware) {
	r.POST("/product",mid.AdminAuthMiddleware, handler.AddProduct)
	r.DELETE("/product/:id",mid.AdminAuthMiddleware, handler.DeleteProduct)
	r.GET("/product/:id", handler.ShowSingleProduct)
	r.GET("/product/all", handler.GetAllProducts)//working
}
