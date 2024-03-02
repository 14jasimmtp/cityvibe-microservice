package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/14jasimmtp/cityvibe-microservices/admin-svc/pb"
	"github.com/14jasimmtp/cityvibe-microservices/admin-svc/pkg/models"
	interfaceRepo "github.com/14jasimmtp/cityvibe-microservices/admin-svc/pkg/repository/interface"
	"github.com/14jasimmtp/cityvibe-microservices/admin-svc/utils"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	r interfaceRepo.Repo
	pb.UnimplementedAdminServiceServer
}

func NewAdminService(R interfaceRepo.Repo) Service {
	return Service{
		r: R,
	}
}

func (svc *Service) AdminLogin(ctx context.Context, req *pb.AdminLoginReq) (*pb.AdminLoginRes, error) {

	AdminDetails, err := svc.r.AdminLogin(models.AdminLogin{Email: req.Email, Password: req.Password})
	fmt.Println(err)
	if err != nil {
		fmt.Println("Admin doesn't exist")
		return nil, errors.New("admin not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(AdminDetails.Password), []byte(req.Password)) != nil {
		fmt.Println("wrong password")
		return nil, errors.New("wrong password")
	}

	tokenString, err := utils.AdminTokenGenerate(AdminDetails, "admin")
	if err != nil {
		fmt.Println("error generating token")
		return nil, errors.New("error generating token")
	}

	return &pb.AdminLoginRes{
		Message: tokenString,
	}, nil

}

func (svc *Service) DashBoard(ctx context.Context, req *pb.NoParam) (*pb.DashBoardRes, error) {
	userDetails, err := svc.r.DashBoardUserDetails()
	if err != nil {
		return nil, err
	}
	productDetails, err := svc.r.DashBoardProductDetails()
	if err != nil {
		return nil, err
	}
	orderDetails, err := svc.r.DashBoardOrder()
	if err != nil {
		return nil, err
	}
	totalRevenue, err := svc.r.TotalRevenue()
	if err != nil {
		return nil, err
	}
	amountDetails, err := svc.r.AmountDetails()
	if err != nil {
		return nil, err
	}
	return &pb.DashBoardRes{
		DashboardUser:    &pb.DashBoardUser{TotalUsers: int32(userDetails.TotalUsers), BlockedUser: userDetails.BlockedUser},
		DashboardProduct: &pb.DashBoardProduct{TotalProducts: int32(productDetails.TotalProducts), OutofStockProductID: productDetails.OutofStockProductID},
		DashboardOrder:   &pb.DashboardOrder{DeliveredOrderProducts: int32(orderDetails.DeliveredOrderProducts), PendingOrderProducts: int32(orderDetails.PendingOrderProducts), CancelledOrderProducts: int32(orderDetails.CancelledOrderProducts), TotalOrderItems: int32(orderDetails.TotalOrderItems), TotalOrderQuantity: int32(orderDetails.TotalOrderQuantity)},
		DashboardRevenue: &pb.DashboardRevenue{TodayRevenue: totalRevenue.TodayRevenue, MonthRevenue: totalRevenue.MonthRevenue, YearRevenue: totalRevenue.YearRevenue},
		DashboardAmount:  &pb.DashboardAmount{CreditedAmount: amountDetails.CreditedAmount, PendingAmount: amountDetails.PendingAmount},
	}, nil
}

func (svc *Service) GetAllUsers(ctx context.Context, req *pb.GetAllUsersReq) (*pb.GetAllUsersRes, error) {
	user, err := svc.r.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var res []*pb.User

	for _,v:=range user{
		users:=&pb.User{
			Id: uint64(v.ID),
			Firstname: v.Firstname,
			Lastname: v.Lastname,
			Email: v.Email,
			Blocked: v.Blocked,
			Wallet: float32(v.Wallet),
			Phone: v.Phone,
		}
		res = append(res, users)
	}

	return &pb.GetAllUsersRes{Users: res}, nil
}

func (svc *Service) BlockUser(ctx context.Context,req *pb.BlockUserReq )(*pb.NoParam,error) {
	id, _ := strconv.Atoi(req.Id)
	user, err := svc.r.GetUserById(id)
	if err != nil {
		return nil,err
	}
	if user.Blocked {
		return nil,errors.New("already blocked")
	} else {
		user.Blocked = true
	}
	err = svc.r.BlockUserByID(*user)
	if err != nil {
		return nil,err
	}
	return nil,nil

}

func (svc *Service) UnBlockUser(ctx context.Context,req *pb.UnBlockUserReq) (*pb.NoParam,error) {
	id, _ := strconv.Atoi(req.Id)
	user, err := svc.r.GetUserById(id)
	if err != nil {
		return nil,err
	}
	if !user.Blocked {
		return nil,errors.New("already unblocked")
	} else {
		user.Blocked = false
	}
	err = svc.r.UnBlockUserByID(*user)
	if err != nil {
		return nil,err
	}
	return nil,nil

}

// func (svc *Service) GetAllOrderDetailsForAdmin(ctx context.Context,req *pb.NoParam) (*pb., error) {
// 	orderDetail, err := svc.r.GetAllOrderDetailsBrief()
// 	if err != nil {
// 		return []models.ViewAdminOrderDetails{}, err
// 	}
// 	return &pb.ViewAdminOrderDetails{}, nil
// }

// func (svc *Service) GetOrderDetails(orderID string) ([]models.OrderProductDetails, error) {
// 	orderDetails, err := svc.r.GetSingleOrderDetails(orderID)
// 	if err != nil {
// 		return []models.OrderProductDetails{}, err
// 	}

// 	return orderDetails, nil
// }

