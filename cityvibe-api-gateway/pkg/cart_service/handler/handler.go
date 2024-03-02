package cartHandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/cart_service/pb"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	client pb.CartServiceClient
}

func NewCartHandler(client pb.CartServiceClient) CartHandler {
	fmt.Println(client)
	return CartHandler{client: client}
}

func (cart *CartHandler) AddToCart(c *gin.Context) {
	pid := c.Query("product_id")

	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Cart, err := cart.client.AddToCart(context.Background(), &pb.AddToCartRequest{Token: Token, Pid: pid})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product added to cart successfully", "Cart": Cart})
}

func (cart *CartHandler) ViewCart(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UserCart, err := cart.client.ViewCart(context.Background(), &pb.ViewCartRequest{Token: Token})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart details", "Cart": UserCart})

}

func (cart *CartHandler) RemoveProductsFromCart(c *gin.Context) {
	id := c.Query("product_id")
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Cart, err := cart.client.RemoveProductsFromCart(context.Background(), &pb.RemoveProductsFromCartRequest{Token: Token, Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product removed from cart successfully", "Cart": Cart})

}
