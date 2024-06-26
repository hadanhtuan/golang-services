// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: proto/user/user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	sdk "property-service/proto/sdk"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MsgUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,20,opt,name=id,proto3" json:"id,omitempty"`
	Role      *string `protobuf:"bytes,1,opt,name=role,proto3,oneof" json:"role,omitempty"`
	Username  string  `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName string  `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string  `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email     string  `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string  `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Password  string  `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	Gender    string  `protobuf:"bytes,8,opt,name=gender,proto3" json:"gender,omitempty"`
	Dob       int64   `protobuf:"varint,9,opt,name=dob,proto3" json:"dob,omitempty"`
	IsActive  *bool   `protobuf:"varint,10,opt,name=isActive,proto3,oneof" json:"isActive,omitempty"`
	UserAgent string  `protobuf:"bytes,11,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
	IpAddress string  `protobuf:"bytes,12,opt,name=ipAddress,proto3" json:"ipAddress,omitempty"`
	DeviceId  string  `protobuf:"bytes,13,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
}

func (x *MsgUser) Reset() {
	*x = MsgUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgUser) ProtoMessage() {}

func (x *MsgUser) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgUser.ProtoReflect.Descriptor instead.
func (*MsgUser) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *MsgUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MsgUser) GetRole() string {
	if x != nil && x.Role != nil {
		return *x.Role
	}
	return ""
}

func (x *MsgUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MsgUser) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *MsgUser) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *MsgUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *MsgUser) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *MsgUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *MsgUser) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *MsgUser) GetDob() int64 {
	if x != nil {
		return x.Dob
	}
	return 0
}

func (x *MsgUser) GetIsActive() bool {
	if x != nil && x.IsActive != nil {
		return *x.IsActive
	}
	return false
}

func (x *MsgUser) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *MsgUser) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *MsgUser) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type MsgToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	DeviceId     string `protobuf:"bytes,3,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	UserId       string `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *MsgToken) Reset() {
	*x = MsgToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgToken) ProtoMessage() {}

func (x *MsgToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgToken.ProtoReflect.Descriptor instead.
func (*MsgToken) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *MsgToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *MsgToken) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *MsgToken) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *MsgToken) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type MsgID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MsgID) Reset() {
	*x = MsgID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgID) ProtoMessage() {}

func (x *MsgID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgID.ProtoReflect.Descriptor instead.
func (*MsgID) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *MsgID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MsgQueryUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paginate    *sdk.Pagination `protobuf:"bytes,1,opt,name=paginate,proto3" json:"paginate,omitempty"`
	TimeQuery   *sdk.TimeQuery  `protobuf:"bytes,2,opt,name=timeQuery,proto3,oneof" json:"timeQuery,omitempty"`
	QueryFields *MsgUser        `protobuf:"bytes,3,opt,name=queryFields,proto3,oneof" json:"queryFields,omitempty"`
	OrderFields *OrderUser      `protobuf:"bytes,4,opt,name=orderFields,proto3,oneof" json:"orderFields,omitempty"`
}

func (x *MsgQueryUser) Reset() {
	*x = MsgQueryUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgQueryUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgQueryUser) ProtoMessage() {}

func (x *MsgQueryUser) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgQueryUser.ProtoReflect.Descriptor instead.
func (*MsgQueryUser) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *MsgQueryUser) GetPaginate() *sdk.Pagination {
	if x != nil {
		return x.Paginate
	}
	return nil
}

func (x *MsgQueryUser) GetTimeQuery() *sdk.TimeQuery {
	if x != nil {
		return x.TimeQuery
	}
	return nil
}

func (x *MsgQueryUser) GetQueryFields() *MsgUser {
	if x != nil {
		return x.QueryFields
	}
	return nil
}

func (x *MsgQueryUser) GetOrderFields() *OrderUser {
	if x != nil {
		return x.OrderFields
	}
	return nil
}

type MsgQueryUserByIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paginate *sdk.Pagination `protobuf:"bytes,1,opt,name=paginate,proto3" json:"paginate,omitempty"`
	Ids      []string        `protobuf:"bytes,3,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *MsgQueryUserByIds) Reset() {
	*x = MsgQueryUserByIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgQueryUserByIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgQueryUserByIds) ProtoMessage() {}

func (x *MsgQueryUserByIds) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgQueryUserByIds.ProtoReflect.Descriptor instead.
func (*MsgQueryUserByIds) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *MsgQueryUserByIds) GetPaginate() *sdk.Pagination {
	if x != nil {
		return x.Paginate
	}
	return nil
}

func (x *MsgQueryUserByIds) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type OrderUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt *sdk.Sort `protobuf:"varint,1,opt,name=createdAt,proto3,enum=sdk.Sort,oneof" json:"createdAt,omitempty"`
}

func (x *OrderUser) Reset() {
	*x = OrderUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUser) ProtoMessage() {}

func (x *OrderUser) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderUser.ProtoReflect.Descriptor instead.
func (*OrderUser) Descriptor() ([]byte, []int) {
	return file_proto_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *OrderUser) GetCreatedAt() sdk.Sort {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return sdk.Sort(0)
}

var File_proto_user_user_proto protoreflect.FileDescriptor

var file_proto_user_user_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x64, 0x6b, 0x2f,
	0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x03, 0x0a, 0x07, 0x4d, 0x73,
	0x67, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x1f, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x05,
	0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x98, 0x02, 0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65,
	0x72, 0x48, 0x01, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x48, 0x02, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x88,
	0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x22, 0x52, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0x47, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x2c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x48,
	0x00, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x84, 0x04,
	0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4d, 0x73, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e,
	0x73, 0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x42, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x49, 0x64,
	0x73, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4d, 0x73, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x73, 0x1a, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d,
	0x73, 0x67, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x1a, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x14, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x1a, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_user_proto_rawDescOnce sync.Once
	file_proto_user_user_proto_rawDescData = file_proto_user_user_proto_rawDesc
)

func file_proto_user_user_proto_rawDescGZIP() []byte {
	file_proto_user_user_proto_rawDescOnce.Do(func() {
		file_proto_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_user_proto_rawDescData)
	})
	return file_proto_user_user_proto_rawDescData
}

var file_proto_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_user_user_proto_goTypes = []interface{}{
	(*MsgUser)(nil),           // 0: userService.MsgUser
	(*MsgToken)(nil),          // 1: userService.MsgToken
	(*MsgID)(nil),             // 2: userService.MsgID
	(*MsgQueryUser)(nil),      // 3: userService.MsgQueryUser
	(*MsgQueryUserByIds)(nil), // 4: userService.MsgQueryUserByIds
	(*OrderUser)(nil),         // 5: userService.OrderUser
	(*sdk.Pagination)(nil),    // 6: sdk.Pagination
	(*sdk.TimeQuery)(nil),     // 7: sdk.TimeQuery
	(sdk.Sort)(0),             // 8: sdk.Sort
	(*sdk.BaseResponse)(nil),  // 9: sdk.BaseResponse
}
var file_proto_user_user_proto_depIdxs = []int32{
	6,  // 0: userService.MsgQueryUser.paginate:type_name -> sdk.Pagination
	7,  // 1: userService.MsgQueryUser.timeQuery:type_name -> sdk.TimeQuery
	0,  // 2: userService.MsgQueryUser.queryFields:type_name -> userService.MsgUser
	5,  // 3: userService.MsgQueryUser.orderFields:type_name -> userService.OrderUser
	6,  // 4: userService.MsgQueryUserByIds.paginate:type_name -> sdk.Pagination
	8,  // 5: userService.OrderUser.createdAt:type_name -> sdk.Sort
	0,  // 6: userService.userService.login:input_type -> userService.MsgUser
	0,  // 7: userService.userService.updateUser:input_type -> userService.MsgUser
	3,  // 8: userService.userService.getUsers:input_type -> userService.MsgQueryUser
	4,  // 9: userService.userService.getUsersByIds:input_type -> userService.MsgQueryUserByIds
	0,  // 10: userService.userService.register:input_type -> userService.MsgUser
	1,  // 11: userService.userService.refreshToken:input_type -> userService.MsgToken
	0,  // 12: userService.userService.logout:input_type -> userService.MsgUser
	2,  // 13: userService.userService.getProfile:input_type -> userService.MsgID
	1,  // 14: userService.userService.verifyToken:input_type -> userService.MsgToken
	9,  // 15: userService.userService.login:output_type -> sdk.BaseResponse
	9,  // 16: userService.userService.updateUser:output_type -> sdk.BaseResponse
	9,  // 17: userService.userService.getUsers:output_type -> sdk.BaseResponse
	9,  // 18: userService.userService.getUsersByIds:output_type -> sdk.BaseResponse
	9,  // 19: userService.userService.register:output_type -> sdk.BaseResponse
	9,  // 20: userService.userService.refreshToken:output_type -> sdk.BaseResponse
	9,  // 21: userService.userService.logout:output_type -> sdk.BaseResponse
	9,  // 22: userService.userService.getProfile:output_type -> sdk.BaseResponse
	9,  // 23: userService.userService.verifyToken:output_type -> sdk.BaseResponse
	15, // [15:24] is the sub-list for method output_type
	6,  // [6:15] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_user_user_proto_init() }
func file_proto_user_user_proto_init() {
	if File_proto_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgUser); i {
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
		file_proto_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgToken); i {
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
		file_proto_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgID); i {
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
		file_proto_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgQueryUser); i {
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
		file_proto_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgQueryUserByIds); i {
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
		file_proto_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderUser); i {
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
	file_proto_user_user_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_user_user_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_proto_user_user_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_user_proto_goTypes,
		DependencyIndexes: file_proto_user_user_proto_depIdxs,
		MessageInfos:      file_proto_user_user_proto_msgTypes,
	}.Build()
	File_proto_user_user_proto = out.File
	file_proto_user_user_proto_rawDesc = nil
	file_proto_user_user_proto_goTypes = nil
	file_proto_user_user_proto_depIdxs = nil
}
