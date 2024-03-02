package service

import (
	"context"

	"github.com/14jasimmtp/cityvibe-microservice/product_svc/pb"
	"github.com/14jasimmtp/cityvibe-microservice/product_svc/pkg/models"
	infRepo "github.com/14jasimmtp/cityvibe-microservice/product_svc/pkg/repository/interface"
	"github.com/jinzhu/copier"
)

type Service struct {
	r infRepo.Repo
	pb.UnimplementedProductServiceServer
}

func NewService(r infRepo.Repo) *Service {
	return &Service{r: r}
}

func (s *Service) AddProduct(ctx context.Context, req *pb.AddProductReq) (*pb.AddProductRes, error) {

	product := models.AddProduct{
		Name:        req.Name,
		Description: req.Description,
		CategoryID:  int(req.CategoryId),
		Size:        int(req.Size),
		Stock:       int(req.Stock),
		Price:       int(req.Price),
		Color:       req.Color,
	}

	ProductResponse, err := s.r.AddProduct(product)
	if err != nil {
		return nil, err
	}
	var res pb.AddProductRes
	copier.Copy(&res, &ProductResponse)
	return &res, nil
}

func (s *Service) GetAllProducts(ctx context.Context, req *pb.NoParam) (*pb.GetAllProductsRes, error) {
	ProductDetails, err := s.r.GetAllProducts()
	if err != nil {
		return nil, err
	}
	var res pb.GetAllProductsRes

	copier.Copy(&res.Product, &ProductDetails)
	return &res, nil
}

func (s *Service) DeleteProduct(ctx context.Context, req *pb.DeleteProductReq) (*pb.DeleteProductRes, error) {

	err := s.r.DeleteProduct(int(req.Pid))
	if err != nil {
		return nil, err
	}

	return &pb.DeleteProductRes{}, nil
}

func (s *Service) GetSingleProduct(ctx context.Context, req *pb.GetSingleProductReq) (*pb.GetSingleProductRes, error) {
	product, err := s.r.GetSingleProduct(req.Pid)
	if err != nil {
		return nil, err
	}

	var res pb.GetSingleProductRes

	res.Product.Name = product.Name
	res.Product.Id = int64(product.ID)
	res.Product.Color = product.Color
	res.Product.Price = float32(product.Price)
	res.Product.Stock = int64(product.Stock)
	res.Product.Description = product.Description

	return &res, nil
}
