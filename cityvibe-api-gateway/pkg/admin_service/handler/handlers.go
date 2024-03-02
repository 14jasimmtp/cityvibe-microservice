package AdminHandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/admin_service/pb"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/models"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	client pb.AdminServiceClient
}

func NewAdminClient(client pb.AdminServiceClient) *AdminHandler {
	fmt.Println(client)
	return &AdminHandler{client: client}
}

func (h *AdminHandler) AdminLogin(c *gin.Context) {
	var admin models.AdminLogin

	if c.ShouldBindJSON(&admin) != nil {
		fmt.Println("binding error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter details correctly"})
		return
	}

	Error, err := utils.Validation(admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
		return
	}
	fmt.Println("hello")
	fmt.Println(h.client)
	admindetails, err := h.client.AdminLogin(context.Background(), &pb.AdminLoginReq{Email: admin.Email,Password: admin.Password})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("Authorisation", admindetails.Message, 36000, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Admin logged in successfully"})

}

func (h *AdminHandler) GetAllUsers(c *gin.Context) {
	Users, err := h.client.GetAllUsers(context.Background(), &pb.GetAllUsersReq{})
	if err != nil {
		fmt.Println("h.client error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(int(Users.Status), gin.H{"message": "users are", "users": Users.Users})
}

func (h *AdminHandler) BlockUser(c *gin.Context) {
	id := c.Query("id")
	_, err := h.client.BlockUser(context.Background(), &pb.BlockUserReq{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user successfully blocked"})
}

func (h *AdminHandler) UnBlockUser(c *gin.Context) {
	id := c.Query("id")
	_, err := h.client.UnBlockUser(context.Background(), &pb.UnBlockUserReq{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user successfully unblocked"})
}

// func (h *AdminHandler) OrderDetailsForAdmin(c *gin.Context) {
// 	allOrderDetails, err := h.client.GetAllOrderDetailsForAdmin(context.Background(),&pb.NoParam{})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't retrieve order details"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Order details retrieved successfully", "All orders": allOrderDetails})
// }

// func (h *AdminHandler) OrderDetailsforAdminWithID(c *gin.Context) {
// 	orderID := c.Query("orderID")

// 	OrderDetails, err := h.client.GetOrderDetails(orderID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"Order Products": OrderDetails})
// }

// func (h *AdminHandler) AddOffer(c *gin.Context) {
// 	var offer models.Offer

// 	if err := c.ShouldBindJSON(&offer); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	err := h.client.ExecuteAddOffer(&offer)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "offer added sussefully"})
// }

// func (h *AdminHandler) AllOffer(c *gin.Context) {

// 	offerlist, err := h.client.ExecuteGetOffers()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"offers": offerlist})
// }

func (h *AdminHandler) DashBoard(c *gin.Context) {
	adminDashboard, err := h.client.DashBoard(context.Background(), &pb.NoParam{})
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "admin dashboard ", "dashboard": adminDashboard})
}
