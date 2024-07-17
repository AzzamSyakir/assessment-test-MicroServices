// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.27.2
// source: office.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Office struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchCode string                 `protobuf:"bytes,1,opt,name=BranchCode,proto3" json:"BranchCode,omitempty"`
	BranchName string                 `protobuf:"bytes,2,opt,name=BranchName,proto3" json:"BranchName,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Office) Reset() {
	*x = Office{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Office) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Office) ProtoMessage() {}

func (x *Office) ProtoReflect() protoreflect.Message {
	mi := &file_office_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Office.ProtoReflect.Descriptor instead.
func (*Office) Descriptor() ([]byte, []int) {
	return file_office_proto_rawDescGZIP(), []int{0}
}

func (x *Office) GetBranchCode() string {
	if x != nil {
		return x.BranchCode
	}
	return ""
}

func (x *Office) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

func (x *Office) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Office) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type OfficeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Office `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *OfficeResponse) Reset() {
	*x = OfficeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfficeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfficeResponse) ProtoMessage() {}

func (x *OfficeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_office_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfficeResponse.ProtoReflect.Descriptor instead.
func (*OfficeResponse) Descriptor() ([]byte, []int) {
	return file_office_proto_rawDescGZIP(), []int{1}
}

func (x *OfficeResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OfficeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OfficeResponse) GetData() *Office {
	if x != nil {
		return x.Data
	}
	return nil
}

type OfficeResponseRepeated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Office `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *OfficeResponseRepeated) Reset() {
	*x = OfficeResponseRepeated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfficeResponseRepeated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfficeResponseRepeated) ProtoMessage() {}

func (x *OfficeResponseRepeated) ProtoReflect() protoreflect.Message {
	mi := &file_office_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfficeResponseRepeated.ProtoReflect.Descriptor instead.
func (*OfficeResponseRepeated) Descriptor() ([]byte, []int) {
	return file_office_proto_rawDescGZIP(), []int{2}
}

func (x *OfficeResponseRepeated) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OfficeResponseRepeated) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OfficeResponseRepeated) GetData() []*Office {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateOfficeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BranchName *string `protobuf:"bytes,2,opt,name=BranchName,proto3,oneof" json:"BranchName,omitempty"`
}

func (x *UpdateOfficeRequest) Reset() {
	*x = UpdateOfficeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOfficeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOfficeRequest) ProtoMessage() {}

func (x *UpdateOfficeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_office_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOfficeRequest.ProtoReflect.Descriptor instead.
func (*UpdateOfficeRequest) Descriptor() ([]byte, []int) {
	return file_office_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateOfficeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateOfficeRequest) GetBranchName() string {
	if x != nil && x.BranchName != nil {
		return *x.BranchName
	}
	return ""
}

type CreateOfficeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchName string `protobuf:"bytes,1,opt,name=BranchName,proto3" json:"BranchName,omitempty"`
}

func (x *CreateOfficeRequest) Reset() {
	*x = CreateOfficeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOfficeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOfficeRequest) ProtoMessage() {}

func (x *CreateOfficeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_office_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOfficeRequest.ProtoReflect.Descriptor instead.
func (*CreateOfficeRequest) Descriptor() ([]byte, []int) {
	return file_office_proto_rawDescGZIP(), []int{4}
}

func (x *CreateOfficeRequest) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

var File_office_proto protoreflect.FileDescriptor

var file_office_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x79, 0x0a, 0x0e,
	0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63,
	0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x81, 0x01, 0x0a, 0x16, 0x4f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x59, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x23, 0x0a, 0x0a, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xaa, 0x04,
	0x0a, 0x0d, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x63, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x23, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x2d, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x12, 0x32, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x32, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x73,
	0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69,
	0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x0c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x23, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63,
	0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a,
	0x2d, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x73, 0x12, 0x24, 0x2e,
	0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x35, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x05, 0x5a, 0x03, 0x70, 0x62,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_office_proto_rawDescOnce sync.Once
	file_office_proto_rawDescData = file_office_proto_rawDesc
)

func file_office_proto_rawDescGZIP() []byte {
	file_office_proto_rawDescOnce.Do(func() {
		file_office_proto_rawDescData = protoimpl.X.CompressGZIP(file_office_proto_rawDescData)
	})
	return file_office_proto_rawDescData
}

var file_office_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_office_proto_goTypes = []interface{}{
	(*Office)(nil),                 // 0: assesement_test_MicroServices.Office
	(*OfficeResponse)(nil),         // 1: assesement_test_MicroServices.OfficeResponse
	(*OfficeResponseRepeated)(nil), // 2: assesement_test_MicroServices.OfficeResponseRepeated
	(*UpdateOfficeRequest)(nil),    // 3: assesement_test_MicroServices.UpdateOfficeRequest
	(*CreateOfficeRequest)(nil),    // 4: assesement_test_MicroServices.CreateOfficeRequest
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(*ById)(nil),                   // 6: assesement_test_MicroServices.ById
	(*Empty)(nil),                  // 7: assesement_test_MicroServices.empty
}
var file_office_proto_depIdxs = []int32{
	5, // 0: assesement_test_MicroServices.Office.CreatedAt:type_name -> google.protobuf.Timestamp
	5, // 1: assesement_test_MicroServices.Office.UpdatedAt:type_name -> google.protobuf.Timestamp
	0, // 2: assesement_test_MicroServices.OfficeResponse.data:type_name -> assesement_test_MicroServices.Office
	0, // 3: assesement_test_MicroServices.OfficeResponseRepeated.data:type_name -> assesement_test_MicroServices.Office
	6, // 4: assesement_test_MicroServices.OfficeService.GetOfficeById:input_type -> assesement_test_MicroServices.ById
	3, // 5: assesement_test_MicroServices.OfficeService.UpdateOffice:input_type -> assesement_test_MicroServices.UpdateOfficeRequest
	4, // 6: assesement_test_MicroServices.OfficeService.CreateOffice:input_type -> assesement_test_MicroServices.CreateOfficeRequest
	6, // 7: assesement_test_MicroServices.OfficeService.DeleteOffice:input_type -> assesement_test_MicroServices.ById
	7, // 8: assesement_test_MicroServices.OfficeService.ListOffices:input_type -> assesement_test_MicroServices.empty
	1, // 9: assesement_test_MicroServices.OfficeService.GetOfficeById:output_type -> assesement_test_MicroServices.OfficeResponse
	1, // 10: assesement_test_MicroServices.OfficeService.UpdateOffice:output_type -> assesement_test_MicroServices.OfficeResponse
	1, // 11: assesement_test_MicroServices.OfficeService.CreateOffice:output_type -> assesement_test_MicroServices.OfficeResponse
	1, // 12: assesement_test_MicroServices.OfficeService.DeleteOffice:output_type -> assesement_test_MicroServices.OfficeResponse
	2, // 13: assesement_test_MicroServices.OfficeService.ListOffices:output_type -> assesement_test_MicroServices.OfficeResponseRepeated
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_office_proto_init() }
func file_office_proto_init() {
	if File_office_proto != nil {
		return
	}
	file_commons_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_office_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Office); i {
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
		file_office_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfficeResponse); i {
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
		file_office_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfficeResponseRepeated); i {
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
		file_office_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOfficeRequest); i {
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
		file_office_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOfficeRequest); i {
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
	file_office_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_office_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_office_proto_goTypes,
		DependencyIndexes: file_office_proto_depIdxs,
		MessageInfos:      file_office_proto_msgTypes,
	}.Build()
	File_office_proto = out.File
	file_office_proto_rawDesc = nil
	file_office_proto_goTypes = nil
	file_office_proto_depIdxs = nil
}
