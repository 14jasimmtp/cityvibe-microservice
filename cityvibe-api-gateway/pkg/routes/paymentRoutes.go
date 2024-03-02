package routes

import (
	PaymentHandler "github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/payment_service/handler"
	"github.com/gin-gonic/gin"
)

func Payment(r *gin.Engine, handler *PaymentHandler.PaymentHandler) {
	r.GET("/razorpay/payment",handler.ExecuteRazorPayPayment)
	r.POST("/razorpay/verify",handler.VerifyPayment)

	r.LoadHTMLFiles("/home/jasim/Cityvibe-microservice/cityvibe-api-gateway/template/notfound.html")
}
