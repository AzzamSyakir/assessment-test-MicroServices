// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.27.2
// source: screen.proto

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

type Screen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ScreenCode string                 `protobuf:"bytes,2,opt,name=screenCode,proto3" json:"screenCode,omitempty"`
	ScreenName string                 `protobuf:"bytes,3,opt,name=ScreenName,proto3" json:"ScreenName,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Screen) Reset() {
	*x = Screen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Screen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Screen) ProtoMessage() {}

func (x *Screen) ProtoReflect() protoreflect.Message {
	mi := &file_screen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Screen.ProtoReflect.Descriptor instead.
func (*Screen) Descriptor() ([]byte, []int) {
	return file_screen_proto_rawDescGZIP(), []int{0}
}

func (x *Screen) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Screen) GetScreenCode() string {
	if x != nil {
		return x.ScreenCode
	}
	return ""
}

func (x *Screen) GetScreenName() string {
	if x != nil {
		return x.ScreenName
	}
	return ""
}

func (x *Screen) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Screen) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ScreenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Screen `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ScreenResponse) Reset() {
	*x = ScreenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreenResponse) ProtoMessage() {}

func (x *ScreenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_screen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreenResponse.ProtoReflect.Descriptor instead.
func (*ScreenResponse) Descriptor() ([]byte, []int) {
	return file_screen_proto_rawDescGZIP(), []int{1}
}

func (x *ScreenResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ScreenResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ScreenResponse) GetData() *Screen {
	if x != nil {
		return x.Data
	}
	return nil
}

type ScreenResponseRepeated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Screen `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ScreenResponseRepeated) Reset() {
	*x = ScreenResponseRepeated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScreenResponseRepeated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScreenResponseRepeated) ProtoMessage() {}

func (x *ScreenResponseRepeated) ProtoReflect() protoreflect.Message {
	mi := &file_screen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScreenResponseRepeated.ProtoReflect.Descriptor instead.
func (*ScreenResponseRepeated) Descriptor() ([]byte, []int) {
	return file_screen_proto_rawDescGZIP(), []int{2}
}

func (x *ScreenResponseRepeated) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ScreenResponseRepeated) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ScreenResponseRepeated) GetData() []*Screen {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateScreenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ScreenName *string `protobuf:"bytes,2,opt,name=ScreenName,proto3,oneof" json:"ScreenName,omitempty"`
}

func (x *UpdateScreenRequest) Reset() {
	*x = UpdateScreenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScreenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScreenRequest) ProtoMessage() {}

func (x *UpdateScreenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_screen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScreenRequest.ProtoReflect.Descriptor instead.
func (*UpdateScreenRequest) Descriptor() ([]byte, []int) {
	return file_screen_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateScreenRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateScreenRequest) GetScreenName() string {
	if x != nil && x.ScreenName != nil {
		return *x.ScreenName
	}
	return ""
}

type CreateScreenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScreenName string `protobuf:"bytes,1,opt,name=ScreenName,proto3" json:"ScreenName,omitempty"`
}

func (x *CreateScreenRequest) Reset() {
	*x = CreateScreenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_screen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScreenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScreenRequest) ProtoMessage() {}

func (x *CreateScreenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_screen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScreenRequest.ProtoReflect.Descriptor instead.
func (*CreateScreenRequest) Descriptor() ([]byte, []int) {
	return file_screen_proto_rawDescGZIP(), []int{4}
}

func (x *CreateScreenRequest) GetScreenName() string {
	if x != nil {
		return x.ScreenName
	}
	return ""
}

var File_screen_proto protoreflect.FileDescriptor

var file_screen_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01,
	0x0a, 0x06, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x79,
	0x0a, 0x0e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d,
	0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x72,
	0x65, 0x65, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x81, 0x01, 0x0a, 0x16, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x59, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0a, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x53, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x32,
	0xaa, 0x04, 0x0a, 0x0d, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x63, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x23, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x2d, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x32, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x72,
	0x65, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63,
	0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x32, 0x2e, 0x61, 0x73, 0x73, 0x65,
	0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x12, 0x23, 0x2e, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d,
	0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x79, 0x49,
	0x64, 0x1a, 0x2d, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x6a, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x12,
	0x24, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x35, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x05, 0x5a, 0x03,
	0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_screen_proto_rawDescOnce sync.Once
	file_screen_proto_rawDescData = file_screen_proto_rawDesc
)

func file_screen_proto_rawDescGZIP() []byte {
	file_screen_proto_rawDescOnce.Do(func() {
		file_screen_proto_rawDescData = protoimpl.X.CompressGZIP(file_screen_proto_rawDescData)
	})
	return file_screen_proto_rawDescData
}

var file_screen_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_screen_proto_goTypes = []interface{}{
	(*Screen)(nil),                 // 0: assesement_test_MicroServices.Screen
	(*ScreenResponse)(nil),         // 1: assesement_test_MicroServices.ScreenResponse
	(*ScreenResponseRepeated)(nil), // 2: assesement_test_MicroServices.ScreenResponseRepeated
	(*UpdateScreenRequest)(nil),    // 3: assesement_test_MicroServices.UpdateScreenRequest
	(*CreateScreenRequest)(nil),    // 4: assesement_test_MicroServices.CreateScreenRequest
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(*ById)(nil),                   // 6: assesement_test_MicroServices.ById
	(*Empty)(nil),                  // 7: assesement_test_MicroServices.empty
}
var file_screen_proto_depIdxs = []int32{
	5, // 0: assesement_test_MicroServices.Screen.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: assesement_test_MicroServices.Screen.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: assesement_test_MicroServices.ScreenResponse.data:type_name -> assesement_test_MicroServices.Screen
	0, // 3: assesement_test_MicroServices.ScreenResponseRepeated.data:type_name -> assesement_test_MicroServices.Screen
	6, // 4: assesement_test_MicroServices.ScreenService.GetScreenById:input_type -> assesement_test_MicroServices.ById
	3, // 5: assesement_test_MicroServices.ScreenService.UpdateScreen:input_type -> assesement_test_MicroServices.UpdateScreenRequest
	4, // 6: assesement_test_MicroServices.ScreenService.CreateScreen:input_type -> assesement_test_MicroServices.CreateScreenRequest
	6, // 7: assesement_test_MicroServices.ScreenService.DeleteScreen:input_type -> assesement_test_MicroServices.ById
	7, // 8: assesement_test_MicroServices.ScreenService.ListScreens:input_type -> assesement_test_MicroServices.empty
	1, // 9: assesement_test_MicroServices.ScreenService.GetScreenById:output_type -> assesement_test_MicroServices.ScreenResponse
	1, // 10: assesement_test_MicroServices.ScreenService.UpdateScreen:output_type -> assesement_test_MicroServices.ScreenResponse
	1, // 11: assesement_test_MicroServices.ScreenService.CreateScreen:output_type -> assesement_test_MicroServices.ScreenResponse
	1, // 12: assesement_test_MicroServices.ScreenService.DeleteScreen:output_type -> assesement_test_MicroServices.ScreenResponse
	2, // 13: assesement_test_MicroServices.ScreenService.ListScreens:output_type -> assesement_test_MicroServices.ScreenResponseRepeated
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_screen_proto_init() }
func file_screen_proto_init() {
	if File_screen_proto != nil {
		return
	}
	file_commons_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_screen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Screen); i {
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
		file_screen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreenResponse); i {
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
		file_screen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScreenResponseRepeated); i {
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
		file_screen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScreenRequest); i {
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
		file_screen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScreenRequest); i {
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
	file_screen_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_screen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_screen_proto_goTypes,
		DependencyIndexes: file_screen_proto_depIdxs,
		MessageInfos:      file_screen_proto_msgTypes,
	}.Build()
	File_screen_proto = out.File
	file_screen_proto_rawDesc = nil
	file_screen_proto_goTypes = nil
	file_screen_proto_depIdxs = nil
}
