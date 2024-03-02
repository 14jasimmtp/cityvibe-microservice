package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/14jasimmtp/cityvibe-microservice/user_svc/pb"
	"github.com/14jasimmtp/cityvibe-microservice/user_svc/pkg/domain"
	"github.com/14jasimmtp/cityvibe-microservice/user_svc/pkg/models"
	RepoIfc "github.com/14jasimmtp/cityvibe-microservice/user_svc/pkg/repository/interface"
	"github.com/14jasimmtp/cityvibe-microservice/user_svc/pkg/utils"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
	r RepoIfc.Repo
	pb.UnimplementedUserServiceServer
}

func NewService(r RepoIfc.Repo,db *gorm.DB) *Service{
	return &Service{r: r,DB: db}
}

func (s *Service) Signup(ctx context.Context, req *pb.UserSignupReq) (*pb.UserSignupRes, error) {
	fmt.Println("hi",)
	CheckEmail, err := s.r.CheckUserExistsEmail(req.Email)
	if err != nil {
		fmt.Println("server error")
		return nil, errors.New("server error")
	}
	fmt.Println("email doesn't exist")
	if CheckEmail != nil {
		fmt.Println("user already exist")
		return nil, errors.New("user already exist with this email")
	}

	CheckPhone, err := s.r.CheckUserExistsByPhone(req.Phone)
	if err != nil {
		fmt.Println("server error")
		return nil, errors.New("server error")
	}
	if CheckPhone != nil {
		fmt.Println("user already exist with this number")
		return nil, errors.New("user already exist with this number")
	}

	if req.Password != req.ConfirmPassword {
		fmt.Println("passwords doesn't match")
		return nil, errors.New("paswords doesn't match")
	}

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		fmt.Println("error while hashing ")
		return nil, errors.New("server error occured(password hashing)")
	}
	req.Password = string(HashedPassword)

	var Userdt domain.User
	err = copier.Copy(&Userdt, &req)
	if err != nil {
		return nil,err
	}
	s.DB.Create(&Userdt)
	return &pb.UserSignupRes{Message: "user successfully signed up .Login to shop"}, nil
}

func (s *Service) UserLogin(ctx context.Context, req *pb.UserLoginReq) (*pb.UserLoginRes, error) {
	CheckPhone, err := s.r.CheckUserExistsByPhone(req.Phone)
	if err != nil {
		return nil, errors.New("error with server")
	}
	if CheckPhone == nil {
		return nil, errors.New("phone number doesn't exist")
	}
	userdetails, err := s.r.FindUserByPhone(req.Phone)
	if err != nil {
		return nil, err
	}

	if userdetails.Blocked {
		return nil, errors.New("user is blocked")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userdetails.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("password not matching")
	}
	Tokenstring, err := utils.TokenGenerate(userdetails,"user")
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error generating token")
	}
	ResUser:=&pb.UserLoginRes{
		User: &pb.User{
			Firstname: userdetails.Firstname,
			Lastname: userdetails.Lastname,
			Email: userdetails.Email,
			Phone: userdetails.Phone,
			Wallet: float32(userdetails.Wallet),
		},
		Message: Tokenstring,
	}


	return ResUser, nil
}

func (s *Service) AddAddress(ctx context.Context, req *pb.AddAddressRequest) (*pb.AddAddressResponse, error) {
	UserId, err := utils.ExtractUserIdFromToken(req.Token)
	if err != nil {
		return nil, err
	}
	address:=models.Address{
		Name: req.Address.Name,
		Housename: req.Address.HouseName,
		Phone: req.Address.Phone,
		Street: req.Address.Street,
		City: req.Address.City,
		State: req.Address.State,
		Pin: req.Address.Pin,
	}

	AddressRes, err := s.r.AddAddress(address, UserId)
	if err != nil {
		return nil, err
	}
	var res pb.AddAddressResponse
	res.Address.City=AddressRes.City

	return  &res,nil
}

func (s *Service) ViewUserAddress(ctx context.Context, req *pb.ViewUserAddressRequest) (*pb.ViewUserAddressResponse, error) {
	UserId, err := utils.ExtractUserIdFromToken(req.Token)
	if err != nil {
		return nil, err
	}

	Address, err := s.r.ViewAddress(UserId)
	if err != nil {
		return nil, err
	}

	var res pb.ViewUserAddressResponse

	copier.Copy(&res.Addresses,&Address)

	return &res, nil
}
