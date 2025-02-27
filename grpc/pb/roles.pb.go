// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: roles.proto

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

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RoleCode  string                 `protobuf:"bytes,2,opt,name=RoleCode,proto3" json:"RoleCode,omitempty"`
	RoleName  string                 `protobuf:"bytes,3,opt,name=RoleName,proto3" json:"RoleName,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_roles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_roles_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetRoleCode() string {
	if x != nil {
		return x.RoleCode
	}
	return ""
}

func (x *Role) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *Role) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Role) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type RoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Role  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RoleResponse) Reset() {
	*x = RoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleResponse) ProtoMessage() {}

func (x *RoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_roles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleResponse.ProtoReflect.Descriptor instead.
func (*RoleResponse) Descriptor() ([]byte, []int) {
	return file_roles_proto_rawDescGZIP(), []int{1}
}

func (x *RoleResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RoleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RoleResponse) GetData() *Role {
	if x != nil {
		return x.Data
	}
	return nil
}

type RoleResponseRepeated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Role `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *RoleResponseRepeated) Reset() {
	*x = RoleResponseRepeated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleResponseRepeated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleResponseRepeated) ProtoMessage() {}

func (x *RoleResponseRepeated) ProtoReflect() protoreflect.Message {
	mi := &file_roles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleResponseRepeated.ProtoReflect.Descriptor instead.
func (*RoleResponseRepeated) Descriptor() ([]byte, []int) {
	return file_roles_proto_rawDescGZIP(), []int{2}
}

func (x *RoleResponseRepeated) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RoleResponseRepeated) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RoleResponseRepeated) GetData() []*Role {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RoleName *string `protobuf:"bytes,2,opt,name=RoleName,proto3,oneof" json:"RoleName,omitempty"`
}

func (x *UpdateRoleRequest) Reset() {
	*x = UpdateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleRequest) ProtoMessage() {}

func (x *UpdateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_roles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoleRequest) Descriptor() ([]byte, []int) {
	return file_roles_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRoleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRoleRequest) GetRoleName() string {
	if x != nil && x.RoleName != nil {
		return *x.RoleName
	}
	return ""
}

type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleName string `protobuf:"bytes,1,opt,name=RoleName,proto3" json:"RoleName,omitempty"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_roles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_roles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_roles_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRoleRequest) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

var File_roles_proto protoreflect.FileDescriptor

var file_roles_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d,
	0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a,
	0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x75, 0x0a, 0x0c, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7d, 0x0a, 0x14, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d,
	0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x51, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x08,
	0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x90, 0x04, 0x0a, 0x0b,
	0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x23, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63,
	0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a,
	0x2b, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x30, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63,
	0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d,
	0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x30, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x73, 0x73, 0x65,
	0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x23, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x2b, 0x2e, 0x61, 0x73, 0x73, 0x65,
	0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x33, 0x2e, 0x61, 0x73, 0x73, 0x65,
	0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x4d, 0x69, 0x63, 0x72,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x05,
	0x5a, 0x03, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_roles_proto_rawDescOnce sync.Once
	file_roles_proto_rawDescData = file_roles_proto_rawDesc
)

func file_roles_proto_rawDescGZIP() []byte {
	file_roles_proto_rawDescOnce.Do(func() {
		file_roles_proto_rawDescData = protoimpl.X.CompressGZIP(file_roles_proto_rawDescData)
	})
	return file_roles_proto_rawDescData
}

var file_roles_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_roles_proto_goTypes = []any{
	(*Role)(nil),                  // 0: assesement_test_MicroServices.Role
	(*RoleResponse)(nil),          // 1: assesement_test_MicroServices.RoleResponse
	(*RoleResponseRepeated)(nil),  // 2: assesement_test_MicroServices.RoleResponseRepeated
	(*UpdateRoleRequest)(nil),     // 3: assesement_test_MicroServices.UpdateRoleRequest
	(*CreateRoleRequest)(nil),     // 4: assesement_test_MicroServices.CreateRoleRequest
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*ById)(nil),                  // 6: assesement_test_MicroServices.ById
	(*Empty)(nil),                 // 7: assesement_test_MicroServices.empty
}
var file_roles_proto_depIdxs = []int32{
	5, // 0: assesement_test_MicroServices.Role.CreatedAt:type_name -> google.protobuf.Timestamp
	5, // 1: assesement_test_MicroServices.Role.UpdatedAt:type_name -> google.protobuf.Timestamp
	0, // 2: assesement_test_MicroServices.RoleResponse.data:type_name -> assesement_test_MicroServices.Role
	0, // 3: assesement_test_MicroServices.RoleResponseRepeated.data:type_name -> assesement_test_MicroServices.Role
	6, // 4: assesement_test_MicroServices.RoleService.GetRoleById:input_type -> assesement_test_MicroServices.ById
	3, // 5: assesement_test_MicroServices.RoleService.UpdateRole:input_type -> assesement_test_MicroServices.UpdateRoleRequest
	4, // 6: assesement_test_MicroServices.RoleService.CreateRole:input_type -> assesement_test_MicroServices.CreateRoleRequest
	6, // 7: assesement_test_MicroServices.RoleService.DeleteRole:input_type -> assesement_test_MicroServices.ById
	7, // 8: assesement_test_MicroServices.RoleService.ListRoles:input_type -> assesement_test_MicroServices.empty
	1, // 9: assesement_test_MicroServices.RoleService.GetRoleById:output_type -> assesement_test_MicroServices.RoleResponse
	1, // 10: assesement_test_MicroServices.RoleService.UpdateRole:output_type -> assesement_test_MicroServices.RoleResponse
	1, // 11: assesement_test_MicroServices.RoleService.CreateRole:output_type -> assesement_test_MicroServices.RoleResponse
	1, // 12: assesement_test_MicroServices.RoleService.DeleteRole:output_type -> assesement_test_MicroServices.RoleResponse
	2, // 13: assesement_test_MicroServices.RoleService.ListRoles:output_type -> assesement_test_MicroServices.RoleResponseRepeated
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_roles_proto_init() }
func file_roles_proto_init() {
	if File_roles_proto != nil {
		return
	}
	file_commons_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_roles_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Role); i {
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
		file_roles_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RoleResponse); i {
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
		file_roles_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RoleResponseRepeated); i {
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
		file_roles_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRoleRequest); i {
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
		file_roles_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRoleRequest); i {
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
	file_roles_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_roles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_roles_proto_goTypes,
		DependencyIndexes: file_roles_proto_depIdxs,
		MessageInfos:      file_roles_proto_msgTypes,
	}.Build()
	File_roles_proto = out.File
	file_roles_proto_rawDesc = nil
	file_roles_proto_goTypes = nil
	file_roles_proto_depIdxs = nil
}
