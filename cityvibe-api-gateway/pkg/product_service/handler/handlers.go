package product

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/models"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/product_service/pb"
	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	client pb.ProductServiceClient
}

func NewProductClient(client pb.ProductServiceClient) ProductHandler {
	fmt.Println(client)
	return ProductHandler{client: client}
}

func (h *ProductHandler) AddProduct(c *gin.Context) {
	var product models.AddProduct

	if c.ShouldBind(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enter Product details correctly"})
		return
	}

	// image, _ := c.FormFile("image")

	Error, err := utils.Validation(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": Error})
		return
	}

	req := pb.AddProductReq{
		Name:        product.Name,
		Color:       product.Color,
		Stock:       int64(product.Stock),
		Size:        int64(product.Size),
		Price:       float32(product.Price),
		Description: product.Description,
		CategoryId:  int64(product.CategoryID),
	}

	NewProduct, err := h.client.AddProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product added successfully", "product": NewProduct})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	idnum, _ := strconv.Atoi(id)

	_, err := h.client.DeleteProduct(context.Background(), &pb.DeleteProductReq{Pid: int64(idnum)})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product removed successfully"})
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	fmt.Println("hi", h.client)
	products, err := h.client.GetAllProducts(context.Background(), &pb.NoParam{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "products list", "products": products.Product})

}

func (h *ProductHandler) ShowSingleProduct(c *gin.Context) {
	id := c.Param("id")
	idnum, _ := strconv.Atoi(id)
	product, err := h.client.GetSingleProduct(context.Background(), &pb.GetSingleProductReq{Pid: int64(idnum)})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product details", "product": product.Product})
}
