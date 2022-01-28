// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/merchant_center_link_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [MerchantCenterLinkService.ListMerchantCenterLinks][google.ads.googleads.v1.services.MerchantCenterLinkService.ListMerchantCenterLinks].
type ListMerchantCenterLinksRequest struct {
	// The ID of the customer onto which to apply the Merchant Center link list
	// operation.
	CustomerId           string   `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMerchantCenterLinksRequest) Reset()         { *m = ListMerchantCenterLinksRequest{} }
func (m *ListMerchantCenterLinksRequest) String() string { return proto.CompactTextString(m) }
func (*ListMerchantCenterLinksRequest) ProtoMessage()    {}
func (*ListMerchantCenterLinksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f76fd8f3030942d, []int{0}
}

func (m *ListMerchantCenterLinksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMerchantCenterLinksRequest.Unmarshal(m, b)
}
func (m *ListMerchantCenterLinksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMerchantCenterLinksRequest.Marshal(b, m, deterministic)
}
func (m *ListMerchantCenterLinksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMerchantCenterLinksRequest.Merge(m, src)
}
func (m *ListMerchantCenterLinksRequest) XXX_Size() int {
	return xxx_messageInfo_ListMerchantCenterLinksRequest.Size(m)
}
func (m *ListMerchantCenterLinksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMerchantCenterLinksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMerchantCenterLinksRequest proto.InternalMessageInfo

func (m *ListMerchantCenterLinksRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

// Response message for [MerchantCenterLinkService.ListMerchantCenterLinks][google.ads.googleads.v1.services.MerchantCenterLinkService.ListMerchantCenterLinks].
type ListMerchantCenterLinksResponse struct {
	// Merchant Center links available for the requested customer
	MerchantCenterLinks  []*resources.MerchantCenterLink `protobuf:"bytes,1,rep,name=merchant_center_links,json=merchantCenterLinks,proto3" json:"merchant_center_links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ListMerchantCenterLinksResponse) Reset()         { *m = ListMerchantCenterLinksResponse{} }
func (m *ListMerchantCenterLinksResponse) String() string { return proto.CompactTextString(m) }
func (*ListMerchantCenterLinksResponse) ProtoMessage()    {}
func (*ListMerchantCenterLinksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f76fd8f3030942d, []int{1}
}

func (m *ListMerchantCenterLinksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMerchantCenterLinksResponse.Unmarshal(m, b)
}
func (m *ListMerchantCenterLinksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMerchantCenterLinksResponse.Marshal(b, m, deterministic)
}
func (m *ListMerchantCenterLinksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMerchantCenterLinksResponse.Merge(m, src)
}
func (m *ListMerchantCenterLinksResponse) XXX_Size() int {
	return xxx_messageInfo_ListMerchantCenterLinksResponse.Size(m)
}
func (m *ListMerchantCenterLinksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMerchantCenterLinksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMerchantCenterLinksResponse proto.InternalMessageInfo

func (m *ListMerchantCenterLinksResponse) GetMerchantCenterLinks() []*resources.MerchantCenterLink {
	if m != nil {
		return m.MerchantCenterLinks
	}
	return nil
}

// Request message for [MerchantCenterLinkService.GetMerchantCenterLink][google.ads.googleads.v1.services.MerchantCenterLinkService.GetMerchantCenterLink].
type GetMerchantCenterLinkRequest struct {
	// Resource name of the Merchant Center link.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMerchantCenterLinkRequest) Reset()         { *m = GetMerchantCenterLinkRequest{} }
func (m *GetMerchantCenterLinkRequest) String() string { return proto.CompactTextString(m) }
func (*GetMerchantCenterLinkRequest) ProtoMessage()    {}
func (*GetMerchantCenterLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f76fd8f3030942d, []int{2}
}

func (m *GetMerchantCenterLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMerchantCenterLinkRequest.Unmarshal(m, b)
}
func (m *GetMerchantCenterLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMerchantCenterLinkRequest.Marshal(b, m, deterministic)
}
func (m *GetMerchantCenterLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMerchantCenterLinkRequest.Merge(m, src)
}
func (m *GetMerchantCenterLinkRequest) XXX_Size() int {
	return xxx_messageInfo_GetMerchantCenterLinkRequest.Size(m)
}
func (m *GetMerchantCenterLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMerchantCenterLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMerchantCenterLinkRequest proto.InternalMessageInfo

func (m *GetMerchantCenterLinkRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [MerchantCenterLinkService.MutateMerchantCenterLink][google.ads.googleads.v1.services.MerchantCenterLinkService.MutateMerchantCenterLink].
type MutateMerchantCenterLinkRequest struct {
	// The ID of the customer being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The operation to perform on the link
	Operation            *MerchantCenterLinkOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MutateMerchantCenterLinkRequest) Reset()         { *m = MutateMerchantCenterLinkRequest{} }
func (m *MutateMerchantCenterLinkRequest) String() string { return proto.CompactTextString(m) }
func (*MutateMerchantCenterLinkRequest) ProtoMessage()    {}
func (*MutateMerchantCenterLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f76fd8f3030942d, []int{3}
}

func (m *MutateMerchantCenterLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMerchantCenterLinkRequest.Unmarshal(m, b)
}
func (m *MutateMerchantCenterLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMerchantCenterLinkRequest.Marshal(b, m, deterministic)
}
func (m *MutateMerchantCenterLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMerchantCenterLinkRequest.Merge(m, src)
}
func (m *MutateMerchantCenterLinkRequest) XXX_Size() int {
	return xxx_messageInfo_MutateMerchantCenterLinkRequest.Size(m)
}
func (m *MutateMerchantCenterLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMerchantCenterLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMerchantCenterLinkRequest proto.InternalMessageInfo

func (m *MutateMerchantCenterLinkRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateMerchantCenterLinkRequest) GetOperation() *MerchantCenterLinkOperation {
	if m != nil {
		return m.Operation
	}
	return nil
}

// A single update on a Merchant Center link.
type MerchantCenterLinkOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The operation to perform
	//
	// Types that are valid to be assigned to Operation:
	//	*MerchantCenterLinkOperation_Update
	//	*MerchantCenterLinkOperation_Remove
	Operation            isMerchantCenterLinkOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *MerchantCenterLinkOperation) Reset()         { *m = MerchantCenterLinkOperation{} }
func (m *MerchantCenterLinkOperation) String() string { return proto.CompactTextString(m) }
func (*MerchantCenterLinkOperation) ProtoMessage()    {}
func (*MerchantCenterLinkOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f76fd8f3030942d, []int{4}
}

func (m *MerchantCenterLinkOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantCenterLinkOperation.Unmarshal(m, b)
}
func (m *MerchantCenterLinkOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantCenterLinkOperation.Marshal(b, m, deterministic)
}
func (m *MerchantCenterLinkOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantCenterLinkOperation.Merge(m, src)
}
func (m *MerchantCenterLinkOperation) XXX_Size() int {
	return xxx_messageInfo_MerchantCenterLinkOperation.Size(m)
}
func (m *MerchantCenterLinkOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantCenterLinkOperation.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantCenterLinkOperation proto.InternalMessageInfo

func (m *MerchantCenterLinkOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isMerchantCenterLinkOperation_Operation interface {
	isMerchantCenterLinkOperation_Operation()
}

type MerchantCenterLinkOperation_Update struct {
	Update *resources.MerchantCenterLink `protobuf:"bytes,1,opt,name=update,proto3,oneof"`
}

type MerchantCenterLinkOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*MerchantCenterLinkOperation_Update) isMerchantCenterLinkOperation_Operation() {}

func (*MerchantCenterLinkOperation_Remove) isMerchantCenterLinkOperation_Operation() {}

func (m *MerchantCenterLinkOperation) GetOperation() isMerchantCenterLinkOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *MerchantCenterLinkOperation) GetUpdate() *resources.MerchantCenterLink {
	if x, ok := m.GetOperation().(*MerchantCenterLinkOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *MerchantCenterLinkOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*MerchantCenterLinkOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MerchantCenterLinkOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MerchantCenterLinkOperation_Update)(nil),
		(*MerchantCenterLinkOperation_Remove)(nil),
	}
}

// Response message for Merchant Center link mutate.
type MutateMerchantCenterLinkResponse struct {
	// Result for the mutate.
	Result               *MutateMerchantCenterLinkResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *MutateMerchantCenterLinkResponse) Reset()         { *m = MutateMerchantCenterLinkResponse{} }
func (m *MutateMerchantCenterLinkResponse) String() string { return proto.CompactTextString(m) }
func (*MutateMerchantCenterLinkResponse) ProtoMessage()    {}
func (*MutateMerchantCenterLinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f76fd8f3030942d, []int{5}
}

func (m *MutateMerchantCenterLinkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMerchantCenterLinkResponse.Unmarshal(m, b)
}
func (m *MutateMerchantCenterLinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMerchantCenterLinkResponse.Marshal(b, m, deterministic)
}
func (m *MutateMerchantCenterLinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMerchantCenterLinkResponse.Merge(m, src)
}
func (m *MutateMerchantCenterLinkResponse) XXX_Size() int {
	return xxx_messageInfo_MutateMerchantCenterLinkResponse.Size(m)
}
func (m *MutateMerchantCenterLinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMerchantCenterLinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMerchantCenterLinkResponse proto.InternalMessageInfo

func (m *MutateMerchantCenterLinkResponse) GetResult() *MutateMerchantCenterLinkResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// The result for the Merchant Center link mutate.
type MutateMerchantCenterLinkResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateMerchantCenterLinkResult) Reset()         { *m = MutateMerchantCenterLinkResult{} }
func (m *MutateMerchantCenterLinkResult) String() string { return proto.CompactTextString(m) }
func (*MutateMerchantCenterLinkResult) ProtoMessage()    {}
func (*MutateMerchantCenterLinkResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f76fd8f3030942d, []int{6}
}

func (m *MutateMerchantCenterLinkResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateMerchantCenterLinkResult.Unmarshal(m, b)
}
func (m *MutateMerchantCenterLinkResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateMerchantCenterLinkResult.Marshal(b, m, deterministic)
}
func (m *MutateMerchantCenterLinkResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateMerchantCenterLinkResult.Merge(m, src)
}
func (m *MutateMerchantCenterLinkResult) XXX_Size() int {
	return xxx_messageInfo_MutateMerchantCenterLinkResult.Size(m)
}
func (m *MutateMerchantCenterLinkResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateMerchantCenterLinkResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateMerchantCenterLinkResult proto.InternalMessageInfo

func (m *MutateMerchantCenterLinkResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*ListMerchantCenterLinksRequest)(nil), "google.ads.googleads.v1.services.ListMerchantCenterLinksRequest")
	proto.RegisterType((*ListMerchantCenterLinksResponse)(nil), "google.ads.googleads.v1.services.ListMerchantCenterLinksResponse")
	proto.RegisterType((*GetMerchantCenterLinkRequest)(nil), "google.ads.googleads.v1.services.GetMerchantCenterLinkRequest")
	proto.RegisterType((*MutateMerchantCenterLinkRequest)(nil), "google.ads.googleads.v1.services.MutateMerchantCenterLinkRequest")
	proto.RegisterType((*MerchantCenterLinkOperation)(nil), "google.ads.googleads.v1.services.MerchantCenterLinkOperation")
	proto.RegisterType((*MutateMerchantCenterLinkResponse)(nil), "google.ads.googleads.v1.services.MutateMerchantCenterLinkResponse")
	proto.RegisterType((*MutateMerchantCenterLinkResult)(nil), "google.ads.googleads.v1.services.MutateMerchantCenterLinkResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/merchant_center_link_service.proto", fileDescriptor_2f76fd8f3030942d)
}

var fileDescriptor_2f76fd8f3030942d = []byte{
	// 696 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcf, 0x4f, 0x13, 0x41,
	0x14, 0x76, 0x4b, 0x24, 0x61, 0xaa, 0x97, 0x31, 0x84, 0xb5, 0x10, 0x68, 0x56, 0x0e, 0xa4, 0x87,
	0xd9, 0x14, 0x42, 0xd4, 0xc5, 0x1a, 0xb7, 0x8d, 0x82, 0x09, 0x08, 0x59, 0x13, 0x62, 0xb4, 0x49,
	0xb3, 0xec, 0x0e, 0x75, 0xc3, 0xee, 0x4c, 0xdd, 0x99, 0xed, 0x05, 0xb9, 0x98, 0x78, 0xf3, 0xe6,
	0x1f, 0x60, 0xe2, 0xd1, 0x3f, 0x85, 0x84, 0x8b, 0x9e, 0xbc, 0x1b, 0x0f, 0xfe, 0x15, 0x66, 0xe7,
	0x47, 0x8b, 0xa1, 0xdb, 0x22, 0xdc, 0xde, 0xce, 0xbc, 0xf9, 0xbe, 0xf9, 0xde, 0xfb, 0xde, 0x2c,
	0x68, 0x75, 0x29, 0xed, 0xc6, 0xd8, 0xf6, 0x43, 0x66, 0xcb, 0x30, 0x8f, 0xfa, 0x75, 0x9b, 0xe1,
	0xb4, 0x1f, 0x05, 0x98, 0xd9, 0x09, 0x4e, 0x83, 0xb7, 0x3e, 0xe1, 0x9d, 0x00, 0x13, 0x8e, 0xd3,
	0x4e, 0x1c, 0x91, 0xa3, 0x8e, 0xda, 0x45, 0xbd, 0x94, 0x72, 0x0a, 0xab, 0xf2, 0x24, 0xf2, 0x43,
	0x86, 0x06, 0x20, 0xa8, 0x5f, 0x47, 0x1a, 0xa4, 0xf2, 0xa8, 0x88, 0x26, 0xc5, 0x8c, 0x66, 0x69,
	0x11, 0x8f, 0xc4, 0xaf, 0x2c, 0xe8, 0xd3, 0xbd, 0xc8, 0xf6, 0x09, 0xa1, 0xdc, 0xe7, 0x11, 0x25,
	0x4c, 0xed, 0x2a, 0x76, 0x5b, 0x7c, 0x1d, 0x64, 0x87, 0xf6, 0x61, 0x84, 0xe3, 0xb0, 0x93, 0xf8,
	0x4c, 0x9f, 0x9f, 0x3b, 0x77, 0x3e, 0x88, 0x23, 0x4c, 0xb8, 0xdc, 0xb0, 0x5c, 0xb0, 0xb8, 0x1d,
	0x31, 0xbe, 0xa3, 0xa8, 0x5b, 0x82, 0x79, 0x3b, 0x22, 0x47, 0xcc, 0xc3, 0xef, 0x32, 0xcc, 0x38,
	0x5c, 0x02, 0xe5, 0x20, 0x63, 0x9c, 0x26, 0x38, 0xed, 0x44, 0xa1, 0x69, 0x54, 0x8d, 0x95, 0x19,
	0x0f, 0xe8, 0xa5, 0xe7, 0xa1, 0xf5, 0xc9, 0x00, 0x4b, 0x85, 0x18, 0xac, 0x47, 0x09, 0xc3, 0x30,
	0x02, 0xb3, 0xa3, 0xd4, 0x31, 0xd3, 0xa8, 0x4e, 0xad, 0x94, 0x57, 0xd7, 0x51, 0x51, 0xfd, 0x06,
	0xd5, 0x41, 0x17, 0xe1, 0xbd, 0x3b, 0xc9, 0x45, 0x4a, 0xab, 0x05, 0x16, 0x36, 0xf1, 0x88, 0xcb,
	0x68, 0x3d, 0xf7, 0xc0, 0x6d, 0x0d, 0xda, 0x21, 0x7e, 0x82, 0x95, 0xa2, 0x5b, 0x7a, 0xf1, 0x85,
	0x9f, 0x60, 0xeb, 0x8b, 0x01, 0x96, 0x76, 0x32, 0xee, 0x73, 0x5c, 0x0c, 0x34, 0xa9, 0x30, 0xf0,
	0x0d, 0x98, 0xa1, 0x3d, 0x9c, 0x8a, 0x56, 0x99, 0xa5, 0xaa, 0xb1, 0x52, 0x5e, 0x6d, 0xa0, 0x49,
	0x46, 0x19, 0xa1, 0x73, 0x57, 0x83, 0x78, 0x43, 0x3c, 0xeb, 0xbb, 0x01, 0xe6, 0xc7, 0xa4, 0xc2,
	0x0d, 0x50, 0xce, 0x7a, 0xa1, 0xcf, 0xb1, 0xb0, 0x81, 0x39, 0x25, 0xe8, 0x2b, 0x9a, 0x5e, 0x3b,
	0x05, 0x3d, 0xcb, 0x9d, 0xb2, 0xe3, 0xb3, 0x23, 0x0f, 0xc8, 0xf4, 0x3c, 0x86, 0xbb, 0x60, 0x5a,
	0x7e, 0x09, 0x55, 0x57, 0xed, 0xcf, 0xd6, 0x0d, 0x4f, 0xc1, 0x40, 0x13, 0x4c, 0xa7, 0x38, 0xa1,
	0x7d, 0x2c, 0xea, 0x30, 0x93, 0xef, 0xc8, 0xef, 0x66, 0xf9, 0x5c, 0x91, 0xac, 0xf7, 0xa0, 0x5a,
	0x5c, 0x75, 0x65, 0xa5, 0x57, 0x39, 0x14, 0xcb, 0x62, 0xae, 0x4a, 0xfa, 0xe4, 0x12, 0x25, 0x2d,
	0xc6, 0xcc, 0x62, 0xee, 0x29, 0x3c, 0xeb, 0x29, 0x58, 0x1c, 0x9f, 0x79, 0x29, 0xef, 0xac, 0x9e,
	0xdd, 0x04, 0x77, 0x2f, 0x22, 0xbc, 0x94, 0x97, 0x81, 0x3f, 0x0d, 0x30, 0x57, 0x30, 0x2d, 0xf0,
	0x12, 0x52, 0xc6, 0x0f, 0x6b, 0xc5, 0xbd, 0x06, 0x82, 0xac, 0xaf, 0xf5, 0xf0, 0xc3, 0x8f, 0x5f,
	0x9f, 0x4b, 0x6b, 0xb0, 0x9e, 0x3f, 0x4e, 0xda, 0xcd, 0xcc, 0x3e, 0x3e, 0xe7, 0xf5, 0x46, 0xed,
	0xc4, 0x1e, 0x31, 0x7a, 0xf0, 0xcc, 0x00, 0xb3, 0x23, 0x67, 0x0f, 0x3e, 0x9e, 0x7c, 0xaf, 0x71,
	0x43, 0x5b, 0xb9, 0x9a, 0x01, 0xad, 0x86, 0xd0, 0x72, 0x1f, 0xae, 0xe7, 0x5a, 0x8e, 0xff, 0x69,
	0x5d, 0x63, 0x28, 0xad, 0x36, 0x4a, 0x8c, 0x5d, 0x3b, 0x81, 0xbf, 0x0d, 0x60, 0x16, 0x39, 0x02,
	0xba, 0xd7, 0xf1, 0x9d, 0x54, 0xd5, 0xbc, 0x96, 0x75, 0x65, 0xbb, 0x5a, 0x42, 0x62, 0xc3, 0x7a,
	0xf0, 0xdf, 0xed, 0x72, 0x12, 0x81, 0xed, 0x18, 0xb5, 0xca, 0xfc, 0xa9, 0x6b, 0x0e, 0xf9, 0x55,
	0xd4, 0x8b, 0x18, 0x0a, 0x68, 0xd2, 0xfc, 0x58, 0x02, 0xcb, 0x01, 0x4d, 0x26, 0xde, 0xb5, 0xb9,
	0x58, 0xe8, 0xfa, 0xbd, 0xfc, 0xb9, 0xd9, 0x33, 0x5e, 0x6f, 0x29, 0x8c, 0x2e, 0x8d, 0x7d, 0xd2,
	0x45, 0x34, 0xed, 0xda, 0x5d, 0x4c, 0xc4, 0x63, 0x64, 0x0f, 0x59, 0x8b, 0x7f, 0xc5, 0x1b, 0x3a,
	0xf8, 0x5a, 0x9a, 0xda, 0x74, 0xdd, 0x6f, 0xa5, 0xea, 0xa6, 0x04, 0x74, 0x43, 0x86, 0x64, 0x98,
	0x47, 0xfb, 0x75, 0xa4, 0x88, 0xd9, 0xa9, 0x4e, 0x69, 0xbb, 0x21, 0x6b, 0x0f, 0x52, 0xda, 0xfb,
	0xf5, 0xb6, 0x4e, 0xf9, 0x53, 0x5a, 0x96, 0xeb, 0x8e, 0xe3, 0x86, 0xcc, 0x71, 0x06, 0x49, 0x8e,
	0xb3, 0x5f, 0x77, 0x1c, 0x9d, 0x76, 0x30, 0x2d, 0xee, 0xb9, 0xf6, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xfa, 0xe5, 0x47, 0x66, 0x31, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MerchantCenterLinkServiceClient is the client API for MerchantCenterLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MerchantCenterLinkServiceClient interface {
	// Returns Merchant Center links available for this customer.
	ListMerchantCenterLinks(ctx context.Context, in *ListMerchantCenterLinksRequest, opts ...grpc.CallOption) (*ListMerchantCenterLinksResponse, error)
	// Returns the Merchant Center link in full detail.
	GetMerchantCenterLink(ctx context.Context, in *GetMerchantCenterLinkRequest, opts ...grpc.CallOption) (*resources.MerchantCenterLink, error)
	// Updates status or removes a Merchant Center link.
	MutateMerchantCenterLink(ctx context.Context, in *MutateMerchantCenterLinkRequest, opts ...grpc.CallOption) (*MutateMerchantCenterLinkResponse, error)
}

type merchantCenterLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantCenterLinkServiceClient(cc grpc.ClientConnInterface) MerchantCenterLinkServiceClient {
	return &merchantCenterLinkServiceClient{cc}
}

func (c *merchantCenterLinkServiceClient) ListMerchantCenterLinks(ctx context.Context, in *ListMerchantCenterLinksRequest, opts ...grpc.CallOption) (*ListMerchantCenterLinksResponse, error) {
	out := new(ListMerchantCenterLinksResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.MerchantCenterLinkService/ListMerchantCenterLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantCenterLinkServiceClient) GetMerchantCenterLink(ctx context.Context, in *GetMerchantCenterLinkRequest, opts ...grpc.CallOption) (*resources.MerchantCenterLink, error) {
	out := new(resources.MerchantCenterLink)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.MerchantCenterLinkService/GetMerchantCenterLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantCenterLinkServiceClient) MutateMerchantCenterLink(ctx context.Context, in *MutateMerchantCenterLinkRequest, opts ...grpc.CallOption) (*MutateMerchantCenterLinkResponse, error) {
	out := new(MutateMerchantCenterLinkResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.MerchantCenterLinkService/MutateMerchantCenterLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantCenterLinkServiceServer is the server API for MerchantCenterLinkService service.
type MerchantCenterLinkServiceServer interface {
	// Returns Merchant Center links available for this customer.
	ListMerchantCenterLinks(context.Context, *ListMerchantCenterLinksRequest) (*ListMerchantCenterLinksResponse, error)
	// Returns the Merchant Center link in full detail.
	GetMerchantCenterLink(context.Context, *GetMerchantCenterLinkRequest) (*resources.MerchantCenterLink, error)
	// Updates status or removes a Merchant Center link.
	MutateMerchantCenterLink(context.Context, *MutateMerchantCenterLinkRequest) (*MutateMerchantCenterLinkResponse, error)
}

// UnimplementedMerchantCenterLinkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMerchantCenterLinkServiceServer struct {
}

func (*UnimplementedMerchantCenterLinkServiceServer) ListMerchantCenterLinks(ctx context.Context, req *ListMerchantCenterLinksRequest) (*ListMerchantCenterLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMerchantCenterLinks not implemented")
}
func (*UnimplementedMerchantCenterLinkServiceServer) GetMerchantCenterLink(ctx context.Context, req *GetMerchantCenterLinkRequest) (*resources.MerchantCenterLink, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMerchantCenterLink not implemented")
}
func (*UnimplementedMerchantCenterLinkServiceServer) MutateMerchantCenterLink(ctx context.Context, req *MutateMerchantCenterLinkRequest) (*MutateMerchantCenterLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateMerchantCenterLink not implemented")
}

func RegisterMerchantCenterLinkServiceServer(s *grpc.Server, srv MerchantCenterLinkServiceServer) {
	s.RegisterService(&_MerchantCenterLinkService_serviceDesc, srv)
}

func _MerchantCenterLinkService_ListMerchantCenterLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMerchantCenterLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterLinkServiceServer).ListMerchantCenterLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.MerchantCenterLinkService/ListMerchantCenterLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterLinkServiceServer).ListMerchantCenterLinks(ctx, req.(*ListMerchantCenterLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantCenterLinkService_GetMerchantCenterLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMerchantCenterLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterLinkServiceServer).GetMerchantCenterLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.MerchantCenterLinkService/GetMerchantCenterLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterLinkServiceServer).GetMerchantCenterLink(ctx, req.(*GetMerchantCenterLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantCenterLinkService_MutateMerchantCenterLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateMerchantCenterLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantCenterLinkServiceServer).MutateMerchantCenterLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.MerchantCenterLinkService/MutateMerchantCenterLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantCenterLinkServiceServer).MutateMerchantCenterLink(ctx, req.(*MutateMerchantCenterLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MerchantCenterLinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.MerchantCenterLinkService",
	HandlerType: (*MerchantCenterLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMerchantCenterLinks",
			Handler:    _MerchantCenterLinkService_ListMerchantCenterLinks_Handler,
		},
		{
			MethodName: "GetMerchantCenterLink",
			Handler:    _MerchantCenterLinkService_GetMerchantCenterLink_Handler,
		},
		{
			MethodName: "MutateMerchantCenterLink",
			Handler:    _MerchantCenterLinkService_MutateMerchantCenterLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/merchant_center_link_service.proto",
}