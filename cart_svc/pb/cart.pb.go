// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.15.8
// source: pb/cart.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ViewCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *ViewCartRequest) Reset() {
	*x = ViewCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCartRequest) ProtoMessage() {}

func (x *ViewCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCartRequest.ProtoReflect.Descriptor instead.
func (*ViewCartRequest) Descriptor() ([]byte, []int) {
	return file_pb_cart_proto_rawDescGZIP(), []int{0}
}

func (x *ViewCartRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID   int64   `protobuf:"varint,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	ProductName string  `protobuf:"bytes,2,opt,name=ProductName,proto3" json:"ProductName,omitempty"`
	Category    string  `protobuf:"bytes,3,opt,name=Category,proto3" json:"Category,omitempty"`
	Quantity    int64   `protobuf:"varint,4,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Price       float32 `protobuf:"fixed32,5,opt,name=Price,proto3" json:"Price,omitempty"`
	FinalPrice  float32 `protobuf:"fixed32,6,opt,name=FinalPrice,proto3" json:"FinalPrice,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_pb_cart_proto_rawDescGZIP(), []int{1}
}

func (x *Cart) GetProductID() int64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *Cart) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Cart) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Cart) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Cart) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Cart) GetFinalPrice() float32 {
	if x != nil {
		return x.FinalPrice
	}
	return 0
}

type ViewCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPrice float32 `protobuf:"fixed32,1,opt,name=TotalPrice,proto3" json:"TotalPrice,omitempty"`
	Cart       []*Cart `protobuf:"bytes,2,rep,name=cart,proto3" json:"cart,omitempty"`
}

func (x *ViewCartResponse) Reset() {
	*x = ViewCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCartResponse) ProtoMessage() {}

func (x *ViewCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCartResponse.ProtoReflect.Descriptor instead.
func (*ViewCartResponse) Descriptor() ([]byte, []int) {
	return file_pb_cart_proto_rawDescGZIP(), []int{2}
}

func (x *ViewCartResponse) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *ViewCartResponse) GetCart() []*Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type AddToCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Pid   string `protobuf:"bytes,2,opt,name=Pid,proto3" json:"Pid,omitempty"`
}

func (x *AddToCartRequest) Reset() {
	*x = AddToCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartRequest) ProtoMessage() {}

func (x *AddToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartRequest.ProtoReflect.Descriptor instead.
func (*AddToCartRequest) Descriptor() ([]byte, []int) {
	return file_pb_cart_proto_rawDescGZIP(), []int{3}
}

func (x *AddToCartRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AddToCartRequest) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

type AddToCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPrice float32 `protobuf:"fixed32,1,opt,name=TotalPrice,proto3" json:"TotalPrice,omitempty"`
	Cart       []*Cart `protobuf:"bytes,2,rep,name=cart,proto3" json:"cart,omitempty"`
}

func (x *AddToCartResponse) Reset() {
	*x = AddToCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartResponse) ProtoMessage() {}

func (x *AddToCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartResponse.ProtoReflect.Descriptor instead.
func (*AddToCartResponse) Descriptor() ([]byte, []int) {
	return file_pb_cart_proto_rawDescGZIP(), []int{4}
}

func (x *AddToCartResponse) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *AddToCartResponse) GetCart() []*Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

type RemoveProductsFromCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveProductsFromCartRequest) Reset() {
	*x = RemoveProductsFromCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProductsFromCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProductsFromCartRequest) ProtoMessage() {}

func (x *RemoveProductsFromCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProductsFromCartRequest.ProtoReflect.Descriptor instead.
func (*RemoveProductsFromCartRequest) Descriptor() ([]byte, []int) {
	return file_pb_cart_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveProductsFromCartRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RemoveProductsFromCartRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemoveProductsFromCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPrice float32 `protobuf:"fixed32,1,opt,name=TotalPrice,proto3" json:"TotalPrice,omitempty"`
	Cart       []*Cart `protobuf:"bytes,2,rep,name=cart,proto3" json:"cart,omitempty"`
}

func (x *RemoveProductsFromCartResponse) Reset() {
	*x = RemoveProductsFromCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveProductsFromCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveProductsFromCartResponse) ProtoMessage() {}

func (x *RemoveProductsFromCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveProductsFromCartResponse.ProtoReflect.Descriptor instead.
func (*RemoveProductsFromCartResponse) Descriptor() ([]byte, []int) {
	return file_pb_cart_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveProductsFromCartResponse) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *RemoveProductsFromCartResponse) GetCart() []*Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

var File_pb_cart_proto protoreflect.FileDescriptor

var file_pb_cart_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x27, 0x0a, 0x0f, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb4, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12,
	0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x4d, 0x0a, 0x10, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x3a,
	0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x50, 0x69, 0x64, 0x22, 0x4e, 0x0a, 0x11, 0x41, 0x64,
	0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x19, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x45, 0x0a, 0x1d, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x5b, 0x0a, 0x1e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x32, 0xcd,
	0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f,
	0x0a, 0x08, 0x56, 0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x56, 0x69, 0x65,
	0x77, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x56,
	0x69, 0x65, 0x77, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x11, 0x2e, 0x41,
	0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1e, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x46, 0x72,
	0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x46, 0x72,
	0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_cart_proto_rawDescOnce sync.Once
	file_pb_cart_proto_rawDescData = file_pb_cart_proto_rawDesc
)

func file_pb_cart_proto_rawDescGZIP() []byte {
	file_pb_cart_proto_rawDescOnce.Do(func() {
		file_pb_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_cart_proto_rawDescData)
	})
	return file_pb_cart_proto_rawDescData
}

var file_pb_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_cart_proto_goTypes = []interface{}{
	(*ViewCartRequest)(nil),                // 0: ViewCartRequest
	(*Cart)(nil),                           // 1: Cart
	(*ViewCartResponse)(nil),               // 2: ViewCartResponse
	(*AddToCartRequest)(nil),               // 3: AddToCartRequest
	(*AddToCartResponse)(nil),              // 4: AddToCartResponse
	(*RemoveProductsFromCartRequest)(nil),  // 5: RemoveProductsFromCartRequest
	(*RemoveProductsFromCartResponse)(nil), // 6: RemoveProductsFromCartResponse
}
var file_pb_cart_proto_depIdxs = []int32{
	1, // 0: ViewCartResponse.cart:type_name -> Cart
	1, // 1: AddToCartResponse.cart:type_name -> Cart
	1, // 2: RemoveProductsFromCartResponse.cart:type_name -> Cart
	0, // 3: CartService.ViewCart:input_type -> ViewCartRequest
	3, // 4: CartService.AddToCart:input_type -> AddToCartRequest
	5, // 5: CartService.RemoveProductsFromCart:input_type -> RemoveProductsFromCartRequest
	2, // 6: CartService.ViewCart:output_type -> ViewCartResponse
	4, // 7: CartService.AddToCart:output_type -> AddToCartResponse
	6, // 8: CartService.RemoveProductsFromCart:output_type -> RemoveProductsFromCartResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_cart_proto_init() }
func file_pb_cart_proto_init() {
	if File_pb_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProductsFromCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveProductsFromCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_cart_proto_goTypes,
		DependencyIndexes: file_pb_cart_proto_depIdxs,
		MessageInfos:      file_pb_cart_proto_msgTypes,
	}.Build()
	File_pb_cart_proto = out.File
	file_pb_cart_proto_rawDesc = nil
	file_pb_cart_proto_goTypes = nil
	file_pb_cart_proto_depIdxs = nil
}
