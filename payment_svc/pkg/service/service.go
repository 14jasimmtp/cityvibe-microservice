package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/14jasimmtp/cityvibe-microservice/payment-svc/pb"
	"github.com/14jasimmtp/cityvibe-microservice/payment-svc/pkg/config"
	interfaceRepo "github.com/14jasimmtp/cityvibe-microservice/payment-svc/pkg/repository/interface"
	"github.com/14jasimmtp/cityvibe-microservice/payment-svc/pkg/utils"
	"github.com/razorpay/razorpay-go"
)

type Service struct {
	r interfaceRepo.Repo
	pb.UnimplementedPaymentServiceServer
}

func NewService(r interfaceRepo.Repo) Service {
	return Service{r: r}
}

func (s *Service) MakePaymentRazorPay(ctx context.Context, req *pb.MprReq) (*pb.MprRes, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	PaymentDetails, err := s.r.GetPaymentDetails(int(req.OrderId))
	if err != nil {
		return nil, err
	}

	client := razorpay.NewClient(cfg.KEY_ID_FOR_PAY, cfg.SECRET_KEY_FOR_PAY)

	data := map[string]interface{}{
		"amount":   int(PaymentDetails.Final_price * 100),
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		fmt.Println("hello")
		fmt.Println(err)
		return nil, err
	}

	razorPayOrderID := body["id"].(string)

	err = s.r.AddRazorPayDetails(int(req.OrderId), razorPayOrderID)
	if err != nil {
		fmt.Println("hig")
		return nil, err
	}

	return &pb.MprRes{TotalPrice: float32(PaymentDetails.Total_price), FinalPrice: float32(PaymentDetails.Final_price), Username: PaymentDetails.Username, Phone: PaymentDetails.Userphone}, nil

}

func (s *Service) PaymentMethodID(ctx context.Context, req *pb.PaymentMethodIdReq) (*pb.PaymentmethodIdRes, error) {
	PaymethodID, err := s.r.PayMethod(int(req.OrderId))
	if err != nil {
		return nil, err
	}
	return &pb.PaymentmethodIdRes{Paymentid: int64(PaymethodID)}, nil
}

func (s *Service) PaymentAlreadyPaid(ctx context.Context, req *pb.PAPreq) (*pb.PapRes, error) {
	AlreadyPayed, err := s.r.PaymentAlreadyPaid(int(req.OrderId))
	if err != nil {
		return &pb.PapRes{Status: false}, err
	}
	return &pb.PapRes{Status: AlreadyPayed}, nil
}

func (s *Service) VerifyPayment(ctx context.Context, req *pb.VpReq) (*pb.VpRes, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	paid, err := s.r.CheckVerifiedPayment(int(req.OGID))
	if err != nil {
		return nil, err
	}
	if paid {
		return nil, errors.New(`already payment verified`)
	}

	result := utils.VerifyPayment(req.OrderId, req.PaymentId, req.Signature, cfg.SECRET_KEY_FOR_PAY)
	if !result {
		return nil, errors.New("payment is unsuccessful")
	}

	orders, err := s.r.UpdateShipmentAndPaymentByOrderID("processing", "paid", int(req.OGID))
	if err != nil {
		return nil, err
	}
	return &pb.VpRes{
		Id:            int64(orders.Id),
		FinalPrice:    float32(orders.FinalPrice),
		PaymentMethod: orders.PaymentMethod,
		PaymentStatus: orders.PaymentStatus,
	}, nil
}
