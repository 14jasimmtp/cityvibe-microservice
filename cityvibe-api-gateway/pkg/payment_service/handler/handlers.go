package PaymentHandler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/models"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/payment_service/pb"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	client pb.PaymentServiceClient
}

func NewPaymentClient(client pb.PaymentServiceClient) *PaymentHandler {
	fmt.Println(client)
	return &PaymentHandler{client: client}
}

func (h *PaymentHandler) ExecuteRazorPayPayment(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error from orderID"})
		return
	}
	paymentMethodID, err := h.client.PaymentMethodID(context.Background(), &pb.PaymentMethodIdReq{OrderId: int64(orderID)})
	if err != nil {
		log.Print(err)
		c.HTML(http.StatusNotFound, "notfound.html", nil)
		return
	}
	if paymentMethodID.Paymentid == 2 {
		payment, _ := h.client.PaymentAlreadyPaid(context.Background(), &pb.PAPreq{OrderId: int64(orderID)})
		if payment.Status {
			c.HTML(http.StatusOK, "pay.html", nil)
			return

		}
		res, err := h.client.MakePaymentRazorPay(context.Background(), &pb.MprReq{OrderId: int64(orderID)})
		if err != nil {
			c.HTML(http.StatusNotFound, "notfound.html", nil)
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"final_price": res.FinalPrice,
			"razor_id":    res.RzpayOID,
			"user_name":   res.Username,
			"total":       res.TotalPrice,
		})
		return
	}
	c.HTML(http.StatusNotFound, "notfound.html", nil)
}

func (h *PaymentHandler) VerifyPayment(c *gin.Context) {
	var Verify models.PaymentVerify
	if c.ShouldBindJSON(&Verify) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enter fields correctly"})
		return
	}
	Error, err := utils.Validation(Verify)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
		return
	}
	order := c.Query("orderId")
	order_id, err := strconv.Atoi(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something went wrong"})
	}
	Order, err := h.client.VerifyPayment(context.Background(), &pb.VpReq{OrderId: Verify.OrderID, PaymentId: Verify.PaymentID, Signature: Verify.Signature, OGID: int64(order_id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated payment details successfully", "Order Details": Order})
}
