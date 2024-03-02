package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/14jasimmtp/cityvibe-microservice/cart_svc/pb"
	interfaceRepo "github.com/14jasimmtp/cityvibe-microservice/cart_svc/pkg/repository/interface"
	"github.com/14jasimmtp/cityvibe-microservice/cart_svc/pkg/utils"
)

type Service struct {
	r interfaceRepo.Repo
	pb.UnimplementedCartServiceServer
}

func NewService(r interfaceRepo.Repo) Service {
	return Service{r: r}
}

func (s *Service) ViewCart(ctx context.Context, req *pb.ViewCartRequest) (*pb.ViewCartResponse, error) {
	UserId, err := utils.ExtractUserIdFromToken(req.Token)
	if err != nil {
		return nil, err
	}

	Cart, err := s.r.DisplayCart(UserId)
	if err != nil {
		return nil, err
	}

	cartTotal, err := s.r.CartTotalAmount(UserId)
	if err != nil {
		return nil, err
	}
	var cart []*pb.Cart

	for _, v := range Cart {
		h := pb.Cart{
			ProductID:   int64(v.ProductID),
			ProductName: v.ProductName,
			FinalPrice:  float32(v.FinalPrice),
			Quantity:    int64(v.Quantity),
			Price:       float32(v.Price),
			Category:    v.Category,
		}
		cart = append(cart, &h)
	}

	return &pb.ViewCartResponse{
		TotalPrice: float32(cartTotal),
		Cart:       cart,
	}, nil

}

func (s *Service) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {

	_, err := s.r.CheckSingleProduct(req.Pid)
	if err != nil {
		return nil, errors.New("product doesn't exist")
	}

	UserId, err := utils.ExtractUserIdFromToken(req.Token)
	if err != nil {
		return nil, err
	}

	ProId, err := strconv.Atoi(req.Pid)
	if err != nil {
		return nil, err
	}

	productPrize, err := s.r.GetCartProductAmountFromID(req.Pid)
	if err != nil {
		return nil, err
	}
	true, err := s.r.CheckProductExistInCart(UserId, req.Pid)
	if err != nil {
		return nil, err

	}
	fmt.Println(true)
	if true {
		TotalProductAmount, err := s.r.TotalPrizeOfProductInCart(UserId, req.Pid)
		if err != nil {
			return nil, err
		}

		err = s.r.UpdateCart(1, TotalProductAmount+productPrize, UserId, req.Pid)
		if err != nil {
			return nil, err
		}
	} else {
		if err := s.r.CheckCartStock(ProId); err != nil {
			return nil, err
		}
		err := s.r.AddToCart(ProId, UserId, productPrize)
		if err != nil {
			return nil, err
		}
	}

	CartDetails, err := s.r.DisplayCart(UserId)
	if err != nil {
		return nil, err
	}

	cartTotalAmount, err := s.r.CartTotalAmount(UserId)
	if err != nil {
		return nil, err
	}

	var cart []*pb.Cart

	for _, v := range CartDetails {
		h := pb.Cart{
			ProductID:   int64(v.ProductID),
			ProductName: v.ProductName,
			FinalPrice:  float32(v.FinalPrice),
			Quantity:    int64(v.Quantity),
			Price:       float32(v.Price),
			Category:    v.Category,
		}
		cart = append(cart, &h)
	}

	return &pb.AddToCartResponse{
		TotalPrice: float32(cartTotalAmount),
		Cart:       cart,
	}, nil
}

func (s *Service) RemoveProductsFromCart(ctx context.Context, req *pb.RemoveProductsFromCartRequest) (*pb.RemoveProductsFromCartResponse, error) {
	ProId, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}

	UserId, err := utils.ExtractUserIdFromToken(req.Token)
	if err != nil {
		return nil, err
	}

	err = s.r.RemoveProductFromCart(ProId, UserId)
	if err != nil {
		return nil, err
	}

	updatedCart, err := s.r.DisplayCart(UserId)
	if err != nil {
		return nil, err
	}
	cartTotal, err := s.r.CartTotalAmount(UserId)
	if err != nil {
		return nil, err
	}
	var cart []*pb.Cart

	for _, v := range updatedCart {
		h := pb.Cart{
			ProductID:   int64(v.ProductID),
			ProductName: v.ProductName,
			FinalPrice:  float32(v.FinalPrice),
			Quantity:    int64(v.Quantity),
			Price:       float32(v.Price),
			Category:    v.Category,
		}
		cart = append(cart, &h)
	}

	return &pb.RemoveProductsFromCartResponse{
		TotalPrice: float32(cartTotal),
		Cart:       cart,
	}, nil
}
