// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: pkg/payment_service/pb/payment.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PaymentService_PaymentMethodID_FullMethodName     = "/PaymentService/PaymentMethodID"
	PaymentService_PaymentAlreadyPaid_FullMethodName  = "/PaymentService/PaymentAlreadyPaid"
	PaymentService_MakePaymentRazorPay_FullMethodName = "/PaymentService/MakePaymentRazorPay"
	PaymentService_VerifyPayment_FullMethodName       = "/PaymentService/VerifyPayment"
)

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	PaymentMethodID(ctx context.Context, in *PaymentMethodIdReq, opts ...grpc.CallOption) (*PaymentmethodIdRes, error)
	PaymentAlreadyPaid(ctx context.Context, in *PAPreq, opts ...grpc.CallOption) (*PapRes, error)
	MakePaymentRazorPay(ctx context.Context, in *MprReq, opts ...grpc.CallOption) (*MprRes, error)
	VerifyPayment(ctx context.Context, in *VpReq, opts ...grpc.CallOption) (*VpRes, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) PaymentMethodID(ctx context.Context, in *PaymentMethodIdReq, opts ...grpc.CallOption) (*PaymentmethodIdRes, error) {
	out := new(PaymentmethodIdRes)
	err := c.cc.Invoke(ctx, PaymentService_PaymentMethodID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) PaymentAlreadyPaid(ctx context.Context, in *PAPreq, opts ...grpc.CallOption) (*PapRes, error) {
	out := new(PapRes)
	err := c.cc.Invoke(ctx, PaymentService_PaymentAlreadyPaid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) MakePaymentRazorPay(ctx context.Context, in *MprReq, opts ...grpc.CallOption) (*MprRes, error) {
	out := new(MprRes)
	err := c.cc.Invoke(ctx, PaymentService_MakePaymentRazorPay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) VerifyPayment(ctx context.Context, in *VpReq, opts ...grpc.CallOption) (*VpRes, error) {
	out := new(VpRes)
	err := c.cc.Invoke(ctx, PaymentService_VerifyPayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility
type PaymentServiceServer interface {
	PaymentMethodID(context.Context, *PaymentMethodIdReq) (*PaymentmethodIdRes, error)
	PaymentAlreadyPaid(context.Context, *PAPreq) (*PapRes, error)
	MakePaymentRazorPay(context.Context, *MprReq) (*MprRes, error)
	VerifyPayment(context.Context, *VpReq) (*VpRes, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (UnimplementedPaymentServiceServer) PaymentMethodID(context.Context, *PaymentMethodIdReq) (*PaymentmethodIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaymentMethodID not implemented")
}
func (UnimplementedPaymentServiceServer) PaymentAlreadyPaid(context.Context, *PAPreq) (*PapRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaymentAlreadyPaid not implemented")
}
func (UnimplementedPaymentServiceServer) MakePaymentRazorPay(context.Context, *MprReq) (*MprRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakePaymentRazorPay not implemented")
}
func (UnimplementedPaymentServiceServer) VerifyPayment(context.Context, *VpReq) (*VpRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPayment not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_PaymentMethodID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentMethodIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).PaymentMethodID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_PaymentMethodID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).PaymentMethodID(ctx, req.(*PaymentMethodIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_PaymentAlreadyPaid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PAPreq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).PaymentAlreadyPaid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_PaymentAlreadyPaid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).PaymentAlreadyPaid(ctx, req.(*PAPreq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_MakePaymentRazorPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MprReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).MakePaymentRazorPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_MakePaymentRazorPay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).MakePaymentRazorPay(ctx, req.(*MprReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_VerifyPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).VerifyPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_VerifyPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).VerifyPayment(ctx, req.(*VpReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PaymentMethodID",
			Handler:    _PaymentService_PaymentMethodID_Handler,
		},
		{
			MethodName: "PaymentAlreadyPaid",
			Handler:    _PaymentService_PaymentAlreadyPaid_Handler,
		},
		{
			MethodName: "MakePaymentRazorPay",
			Handler:    _PaymentService_MakePaymentRazorPay_Handler,
		},
		{
			MethodName: "VerifyPayment",
			Handler:    _PaymentService_VerifyPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/payment_service/pb/payment.proto",
}
