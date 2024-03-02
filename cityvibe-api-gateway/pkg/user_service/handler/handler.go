package UserHandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/models"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/user_service/pb"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	client pb.UserServiceClient
}

func NewUserClient(client pb.UserServiceClient) UserHandler {
	fmt.Println(client)
	return UserHandler{client: client}
}

func (h *UserHandler) UserLogin(c *gin.Context) {
	fmt.Println(h.client)

	var User models.UserLoginDetails

	if c.ShouldBindJSON(&User) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Details in correct format"})
		return
	}
	Error, err := utils.Validation(&User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
		return
	}

	user, err := h.client.UserLogin(context.Background(), &pb.UserLoginReq{
		Phone:    User.Phone,
		Password: User.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("Authorisation", user.Message, 3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "user successfully logged in", "user": user.User})
}

func (h *UserHandler) UserSignUp(c *gin.Context) {
	fmt.Println(h.client)

	var User models.UserSignUpDetails

	if c.ShouldBindJSON(&User) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Details in correct format"})
		return
	}
	data, err := utils.Validation(&User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": data})
	}

	resp, err := h.client.Signup(context.Background(), &pb.UserSignupReq{
		FirstName:       User.FirstName,
		LastName:        User.LastName,
		Email:           User.Email,
		Phone:           User.Phone,
		Password:        User.Password,
		ConfirmPassword: User.ConfirmPassword,
	})
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": resp.Message})
}

func (h *UserHandler) ViewUserAddress(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Address, err := h.client.ViewUserAddress(context.Background(), &pb.ViewUserAddressRequest{Token: Token})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Address", "Address": Address.Addresses})

}

func (h *UserHandler) AddNewAddressDetails(c *gin.Context) {
	var Address models.Address

	if c.ShouldBindJSON(&Address) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Details correctly"})
	}

	Error, err := utils.Validation(Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
		return
	}

	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	AddressRes, err := h.client.AddAddress(context.Background(), &pb.AddAddressRequest{Address: &pb.Address{
		Name:      Address.Name,
		HouseName: Address.Housename,
		Phone:     Address.Phone,
		Street:    Address.Street,
		City:      Address.City,
		State:     Address.State,
		Pin:       Address.Pin,
	}, Token: Token})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address added successfully", "Address": AddressRes.Address})
}
