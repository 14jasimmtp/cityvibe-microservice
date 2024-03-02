package OrderHandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/models"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/order_service/pb"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	client pb.OrderServiceClient
}

func NewOrderClient(client pb.OrderServiceClient) *OrderHandler {
	fmt.Println(client)
	return &OrderHandler{client: client}
}

// OrderFromCart godoc
// @Summary Place an order from the user's cart
// @Description Place an order using the provided checkout details.
// @Tags User Order
// @Accept json
// @Produce json
// @Param OrderInput body models.CheckOut true "Details for the order checkout"
// @Success 200 {object} string "message": "Ordered products successfully", "order Details": OrderDetails
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /orders [post]
func (h *OrderHandler) OrderFromCart(c *gin.Context) {
	var OrderInput models.CheckOut

	if c.ShouldBindJSON(&OrderInput) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Enter constraints correctly"})
		return
	}
	Error, err := utils.Validation(OrderInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
		return
	}
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if OrderInput.PaymentID == 1 || OrderInput.PaymentID == 2 {
		OrderDetails, err := h.client.ExecutePurchase(context.Background(), &pb.ExecutePurchaseReq{
			Token: Token,
			OrderInput: &pb.Checkout{
				PaymentId: int64(OrderInput.PaymentID),
				AddressId: int64(OrderInput.AddressID),
			},
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "ordered products successfully", "order Details": OrderDetails})
		}

	} else if OrderInput.PaymentID == 3 {
		OrderDetails, err := h.client.ExecutePurchaseWallet(context.Background(), &pb.ExecutePurchaseWalletReq{
			Token: Token,
			OrderInput: &pb.Checkout{
				PaymentId: int64(OrderInput.PaymentID),
				AddressId: int64(OrderInput.AddressID),
			},
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "ordered products successfully", "order Details": OrderDetails})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enter a valid payment method"})
		return
	}
}

// ViewCheckOut godoc
// @Summary View the checkout page
// @Description Retrieve details for the user's checkout page.
// @Tags User Order
// @Accept json
// @Produce json
// @Success 200 {object} string "message": "CheckOut Page loaded successfully", "order Details": OrderDetails
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /checkout [get]
func (h *OrderHandler) ViewCheckOut(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	OrderDetails, err := h.client.CheckOut(context.Background(), &pb.CheckOutReq{Token: Token})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CheckOut Page loaded successfully", "order Details": OrderDetails})
}

// ViewOrders godoc
// @Summary View user orders
// @Description Retrieve details of orders for the authenticated user.
// @Tags User Order
// @Accept json
// @Produce json
// @Success 200 {object} string "message": "Orders", "order Details": OrderDetails
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /orders [get]
func (h *OrderHandler) ViewOrders(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	OrderDetails, err := h.client.ViewUserOrders(context.Background(), &pb.ViewUserOrdersReq{Token: Token})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "orders", "order Details": OrderDetails})

}

// CancelOrder godoc
// @Summary Cancel an order
// @Description Cancel an order for the authenticated user based on the provided order and product IDs.
// @Tags User Order
// @Accept json
// @Produce json
// @Param order_id query string true "Order ID to be cancelled"
// @Param product_id query string true "Product ID in the order to be cancelled"
// @Success 200 {object} string "message": "Order cancelled successfully"
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 401 {object} string "error": "Unauthorized"
// @Failure 500 {object} string "error": "Internal Server Error"
// @Router /orders/cancel [put]
func (h *OrderHandler) CancelOrder(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderId := c.Query("order_id")
	product_id := c.Query("product_id")

	_, err = h.client.CancelOrder(context.Background(), &pb.CancelOrderReq{Token: Token, OrderId: orderId, ProductId: product_id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order cancelled successfully"})

}
