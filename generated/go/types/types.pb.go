// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: types.proto

package types

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Custom Types
type Fee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Asset      string `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	BasisPoint int64  `protobuf:"varint,2,opt,name=basis_point,json=basisPoint,proto3" json:"basis_point,omitempty"`
}

func (x *Fee) Reset() {
	*x = Fee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fee) ProtoMessage() {}

func (x *Fee) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fee.ProtoReflect.Descriptor instead.
func (*Fee) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *Fee) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

func (x *Fee) GetBasisPoint() int64 {
	if x != nil {
		return x.BasisPoint
	}
	return 0
}

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAmount  int64 `protobuf:"varint,1,opt,name=base_amount,json=baseAmount,proto3" json:"base_amount,omitempty"`
	QuoteAmount int64 `protobuf:"varint,2,opt,name=quote_amount,json=quoteAmount,proto3" json:"quote_amount,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *Balance) GetBaseAmount() int64 {
	if x != nil {
		return x.BaseAmount
	}
	return 0
}

func (x *Balance) GetQuoteAmount() int64 {
	if x != nil {
		return x.QuoteAmount
	}
	return 0
}

type BalanceWithFee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance *Balance `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Fee     *Fee     `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (x *BalanceWithFee) Reset() {
	*x = BalanceWithFee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceWithFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceWithFee) ProtoMessage() {}

func (x *BalanceWithFee) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceWithFee.ProtoReflect.Descriptor instead.
func (*BalanceWithFee) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{2}
}

func (x *BalanceWithFee) GetBalance() *Balance {
	if x != nil {
		return x.Balance
	}
	return nil
}

func (x *BalanceWithFee) GetFee() *Fee {
	if x != nil {
		return x.Fee
	}
	return nil
}

type Market struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAsset  string `protobuf:"bytes,1,opt,name=base_asset,json=baseAsset,proto3" json:"base_asset,omitempty"`
	QuoteAsset string `protobuf:"bytes,2,opt,name=quote_asset,json=quoteAsset,proto3" json:"quote_asset,omitempty"`
}

func (x *Market) Reset() {
	*x = Market{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Market) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Market) ProtoMessage() {}

func (x *Market) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Market.ProtoReflect.Descriptor instead.
func (*Market) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{3}
}

func (x *Market) GetBaseAsset() string {
	if x != nil {
		return x.BaseAsset
	}
	return ""
}

func (x *Market) GetQuoteAsset() string {
	if x != nil {
		return x.QuoteAsset
	}
	return ""
}

type MarketWithFee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Market *Market `protobuf:"bytes,1,opt,name=market,proto3" json:"market,omitempty"`
	Fee    *Fee    `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (x *MarketWithFee) Reset() {
	*x = MarketWithFee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarketWithFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarketWithFee) ProtoMessage() {}

func (x *MarketWithFee) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarketWithFee.ProtoReflect.Descriptor instead.
func (*MarketWithFee) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{4}
}

func (x *MarketWithFee) GetMarket() *Market {
	if x != nil {
		return x.Market
	}
	return nil
}

func (x *MarketWithFee) GetFee() *Fee {
	if x != nil {
		return x.Fee
	}
	return nil
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasePrice  float32 `protobuf:"fixed32,1,opt,name=base_price,json=basePrice,proto3" json:"base_price,omitempty"`
	QuotePrice float32 `protobuf:"fixed32,2,opt,name=quote_price,json=quotePrice,proto3" json:"quote_price,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{5}
}

func (x *Price) GetBasePrice() float32 {
	if x != nil {
		return x.BasePrice
	}
	return 0
}

func (x *Price) GetQuotePrice() float32 {
	if x != nil {
		return x.QuotePrice
	}
	return 0
}

type PriceWithFee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price  *Price `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	Fee    *Fee   `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
	Amount uint64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Asset  string `protobuf:"bytes,4,opt,name=asset,proto3" json:"asset,omitempty"`
}

func (x *PriceWithFee) Reset() {
	*x = PriceWithFee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceWithFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceWithFee) ProtoMessage() {}

func (x *PriceWithFee) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceWithFee.ProtoReflect.Descriptor instead.
func (*PriceWithFee) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{6}
}

func (x *PriceWithFee) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *PriceWithFee) GetFee() *Fee {
	if x != nil {
		return x.Fee
	}
	return nil
}

func (x *PriceWithFee) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PriceWithFee) GetAsset() string {
	if x != nil {
		return x.Asset
	}
	return ""
}

type AddressWithBlindingKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The confidential address encoded using a blech32 format.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The blinding private key for the given address encoded in hex format
	Blinding string `protobuf:"bytes,2,opt,name=blinding,proto3" json:"blinding,omitempty"`
}

func (x *AddressWithBlindingKey) Reset() {
	*x = AddressWithBlindingKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressWithBlindingKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressWithBlindingKey) ProtoMessage() {}

func (x *AddressWithBlindingKey) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressWithBlindingKey.ProtoReflect.Descriptor instead.
func (*AddressWithBlindingKey) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{7}
}

func (x *AddressWithBlindingKey) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddressWithBlindingKey) GetBlinding() string {
	if x != nil {
		return x.Blinding
	}
	return ""
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a,
	0x03, 0x46, 0x65, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61,
	0x73, 0x69, 0x73, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x62, 0x61, 0x73, 0x69, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x07, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x62, 0x61, 0x73,
	0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x0e, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x74, 0x68, 0x46, 0x65, 0x65, 0x12, 0x22, 0x0a, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e,
	0x46, 0x65, 0x65, 0x52, 0x03, 0x66, 0x65, 0x65, 0x22, 0x48, 0x0a, 0x06, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x22, 0x48, 0x0a, 0x0d, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x46, 0x65, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x04, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x03, 0x66, 0x65, 0x65, 0x22, 0x47, 0x0a, 0x05,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x72, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x46, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x04, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x22, 0x4e, 0x0a, 0x16, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x57, 0x69, 0x74, 0x68, 0x42, 0x6c, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x6c, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x64, 0x65, 0x78, 0x2d, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x64, 0x65, 0x78, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_types_proto_goTypes = []interface{}{
	(*Fee)(nil),                    // 0: Fee
	(*Balance)(nil),                // 1: Balance
	(*BalanceWithFee)(nil),         // 2: BalanceWithFee
	(*Market)(nil),                 // 3: Market
	(*MarketWithFee)(nil),          // 4: MarketWithFee
	(*Price)(nil),                  // 5: Price
	(*PriceWithFee)(nil),           // 6: PriceWithFee
	(*AddressWithBlindingKey)(nil), // 7: AddressWithBlindingKey
}
var file_types_proto_depIdxs = []int32{
	1, // 0: BalanceWithFee.balance:type_name -> Balance
	0, // 1: BalanceWithFee.fee:type_name -> Fee
	3, // 2: MarketWithFee.market:type_name -> Market
	0, // 3: MarketWithFee.fee:type_name -> Fee
	5, // 4: PriceWithFee.price:type_name -> Price
	0, // 5: PriceWithFee.fee:type_name -> Fee
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fee); i {
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
		file_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
		file_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceWithFee); i {
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
		file_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Market); i {
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
		file_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarketWithFee); i {
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
		file_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceWithFee); i {
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
		file_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressWithBlindingKey); i {
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
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
