package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/14jasimmtp/cityvibe-microservice/order-svc/pb"
	"github.com/14jasimmtp/cityvibe-microservice/order-svc/pkg/domain"
	interfaceRepo "github.com/14jasimmtp/cityvibe-microservice/order-svc/pkg/repository/interface"
	"github.com/14jasimmtp/cityvibe-microservice/order-svc/utils"
)

type Service struct {
	r interfaceRepo.Repo
	pb.UnimplementedOrderServiceServer
}

func NewService(r interfaceRepo.Repo) Service {
	return Service{r: r}
}

func (s *Service) CheckOut(ctx context.Context, req *pb.CheckOutReq) (*pb.CheckOutRes, error) {
	userId, err := utils.ExtractUserIdFromToken(req.Token)
	if err != nil {
		return nil, err
	}

	AllUserAddress, err := s.r.ViewAddress(userId)
	if err != nil {
		return nil, err
	}

	AllCartProducts, err := s.r.DisplayCart(userId)
	if err != nil {
		return nil, err
	}

	TotalAmount, err := s.r.CartTotalAmount(userId)
	if err != nil {
		return nil, err
	}
	var addreass []*pb.Addresss
	for _,v:=range AllUserAddress{
		a:=&pb.Addresss{
			Name: v.Name,
			Id: int64(v.ID),
			City: v.City,
			HouseName: v.House_name,
			Pin: v.Pin,
			Street: v.Street,
			State: v.State,
			Phone: v.Phone,
		}
		addreass = append(addreass, a)
	}

	var cart []*pb.Carts
	for _,v:=range AllCartProducts{
		c:=&pb.Carts{
			ProductID: int64(v.ProductID),
			ProductName: v.ProductName,
			Category: v.Category,
			Price: float32(v.Price),
			FinalPrice: float32(v.FinalPrice),
			Quantity: int64(v.Quantity),

		}
		cart = append(cart, c)
	}

	if AllCartProducts[0].FinalPrice == 0 {
		return &pb.CheckOutRes{
			Address:     addreass,
			Cart:        cart,
			TotalAmount: float32(TotalAmount),
		}, nil
	} else {

		finalPrice, err := s.r.CartFinalPrice(userId)
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(AllCartProducts); i++ {
			AllCartProducts[i].Price = AllCartProducts[i].FinalPrice
		}
		return &pb.CheckOutRes{
			Address:        addreass,
			Cart:           cart,
			TotalAmount:    float32(TotalAmount),
			DiscountAmount: float32(finalPrice),
		}, nil
	}
}

func (s *Service) ExecutePurchase(ctx context.Context, req *pb.ExecutePurchaseReq) (*pb.ExecutePurchaseResponse, error) {
	var TotalAmount float64
	var method string
	userId, err := utils.ExtractUserIdFromToken(req.Token)
	if err != nil {
		return nil, err
	}

	addressExist := s.r.CheckAddressExist(userId, uint(req.OrderInput.AddressId))
	if !addressExist {
		return nil, errors.New(`address doesn't exist`)
	}

	paymentExist := s.r.CheckPaymentMethodExist(uint(req.OrderInput.AddressId))
	if !paymentExist {
		return &pb.ExecutePurchaseResponse{}, errors.New(`payment method doesn't exist`)
	}
	if req.OrderInput.PaymentId == 1 {
		method = "COD"
	} else {
		method = "Razorpay"
	}

	cartExist := s.r.CheckCartExist(userId)
	if !cartExist {
		return nil, errors.New(`cart is empty`)
	}

	cartItems, err := s.r.DisplayCart(userId)
	if err != nil {
		return nil, err
	}
	if cartItems[0].FinalPrice != 0 {

		for i := 0; i < len(cartItems); i++ {
			cartItems[i].Price = cartItems[i].FinalPrice
		}

		TotalAmount, err = s.r.CartFinalPrice(userId)
		if err != nil {
			return nil, errors.New(`error while calculating total amount`)
		}
	} else {
		TotalAmount, err = s.r.CartTotalAmount(userId)
		if err != nil {
			return nil, errors.New(`error while calculating total amount`)
		}
	}

	OrderID, err := s.r.OrderFromCart(uint(req.OrderInput.AddressId), uint(req.OrderInput.PaymentId), userId, TotalAmount)
	if err != nil {
		return nil, err
	}

	if err := s.r.AddOrderProducts(userId, OrderID, cartItems); err != nil {
		return nil, err
	}

	var orderItemDetails domain.OrderItem
	for _, c := range cartItems {
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = c.Quantity
		err := s.r.UpdateCartAndStockAfterOrder(userId, int(orderItemDetails.ProductID), orderItemDetails.Quantity)
		if err != nil {
			return nil, err
		}
	}
	return &pb.ExecutePurchaseResponse{
		OrderId:       int64(OrderID),
		Paymentmthd:   method,
		TotalAmount:   float32(TotalAmount),
		PaymentStatus: "not paid",
	}, nil
}

func (s *Service) ExecutePurchaseWallet(ctx context.Context, req *pb.ExecutePurchaseWalletReq) (*pb.ExecutePurchaseWalletRes, error) {
	var TotalAmount float64
	userId, err := utils.ExtractUserIdFromToken(req.Token)
	if err != nil {
		return nil, err
	}

	user, err := s.r.GetUserById(int(userId))
	if err != nil {
		return nil, err
	}

	addressExist := s.r.CheckAddressExist(userId, uint(req.OrderInput.AddressId))
	if !addressExist {
		return nil, errors.New(`address doesn't exist`)
	}

	paymentExist := s.r.CheckPaymentMethodExist(uint(req.OrderInput.PaymentId))
	if !paymentExist {
		return nil, errors.New(`payment method doesn't exist`)
	}

	cartExist := s.r.CheckCartExist(userId)
	if !cartExist {
		return nil, errors.New(`cart is empty`)
	}

	cartItems, err := s.r.DisplayCart(userId)
	if err != nil {
		return nil, err
	}

	if cartItems[1].FinalPrice != 0 {

		for i := 0; i < len(cartItems); i++ {
			cartItems[i].Price = cartItems[i].FinalPrice
		}

		TotalAmount, err = s.r.CartFinalPrice(userId)
		if err != nil {
			return nil, errors.New(`error while calculating total amount`)
		}
	} else {
		TotalAmount, err = s.r.CartTotalAmount(userId)
		if err != nil {
			return nil, errors.New(`error while calculating total amount`)
		}
	}

	if user.Wallet < TotalAmount {
		return nil, errors.New(`insufficient Balance in Wallet.Add amount to wallet to purchase`)
	}

	OrderID, err := s.r.OrderFromCart(uint(req.OrderInput.AddressId), uint(req.OrderInput.PaymentId), userId, TotalAmount)
	if err != nil {
		return nil, err
	}

	if err := s.r.AddOrderProducts(userId, OrderID, cartItems); err != nil {
		return nil, err
	}
	_, err = s.r.UpdateShipmentAndPaymentByOrderID("processing", "paid", OrderID)
	if err != nil {
		return nil, err
	}

	user.Wallet -= TotalAmount

	err = s.r.UpdateWallet(user.Wallet, userId)
	if err != nil {
		return nil, err
	}

	var orderItemDetails domain.OrderItem
	for _, c := range cartItems {
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = c.Quantity
		err := s.r.UpdateCartAndStockAfterOrder(userId, int(orderItemDetails.ProductID), orderItemDetails.Quantity)
		if err != nil {
			return nil, err
		}
	}
	return &pb.ExecutePurchaseWalletRes{
		OrderId:       int64(OrderID),
		Paymentmthd:   "Wallet",
		TotalAmount:   float32(TotalAmount),
		PaymentStatus: "paid",
	}, nil
}

func (s *Service) ViewUserOrders(ctx context.Context, req *pb.ViewUserOrdersReq) (*pb.ViewUserOrdersRes, error) {
	userId, err := utils.ExtractUserIdFromToken(req.Token)
	if err != nil {
		return nil, err
	}

	OrderDetails, err := s.r.GetOrderDetails(userId)
	if err != nil {
		return nil, err
	}
os := &pb.ViewUserOrdersRes{}

for _, v := range OrderDetails {
    order := &pb.ViewUserOrders{
        OrderDetails: &pb.OrderDetails{
            Id:            int32(v.OrderDetails.Id),
            FinalPrice:    float32(v.OrderDetails.FinalPrice),
            PaymentMethod: v.OrderDetails.PaymentMethod,
            PaymentStatus: v.OrderDetails.PaymentStatus,
        },
        OrderProductDetails: make([]*pb.OrderProductDetails, 0), 
    }

    for _, productDetail := range v.OrderProductDetails {
        orderProductDetail := &pb.OrderProductDetails{
            ProductId: uint32(productDetail.ProductID),
			ProductName: productDetail.ProductName,
			OrderStatus: productDetail.OrderStatus,
			Quantity: int32(productDetail.Quantity),
			TotalPrice: float32(productDetail.TotalPrice),
        }
        order.OrderProductDetails = append(order.OrderProductDetails, orderProductDetail)
    }

    	os.Orders = append(os.Orders, order)
	}
	return os,nil
}


func (s *Service) CancelOrder(ctx context.Context,req *pb.CancelOrderReq) (*pb.CancelOrderRes,error) {
	UserID, err := utils.ExtractUserIdFromToken(req.Token)
	if err != nil {
		return nil,err
	}
	err = s.r.CheckOrder(req.OrderId, UserID)
	if err != nil {
		return nil,errors.New(`no orders found with this id`)
	}

	OrderDetails, err := s.r.CancelOrderDetails(UserID, req.OrderId, req.ProductId)
	if err != nil {
		return nil,err
	}

	if OrderDetails.OrderStatus == "Delivered" {
		return nil,errors.New(`the order is delivered .Can't Cancel`)
	}

	if OrderDetails.OrderStatus == "Cancelled" {
		return nil,errors.New(`the order is already cancelled`)
	}

	if OrderDetails.PaymentStatus == "paid" {
		err := s.r.ReturnAmountToWallet(UserID, req.OrderId, req.ProductId)
		if err != nil {
			return nil,err
		}

	}
	err = s.r.UpdateOrderFinalPrice(OrderDetails.OrderID, OrderDetails.TotalPrice)
	if err != nil {
		return nil,err
	}
	proid, _ := strconv.Atoi(req.ProductId)
	err = s.r.UpdateStock(proid, OrderDetails.Quantity)
	if err != nil {
		return nil,err
	}

	err = s.r.CancelOrder(req.OrderId, req.ProductId, UserID)
	if err != nil {
		return nil,err
	}

	return nil,nil
}
