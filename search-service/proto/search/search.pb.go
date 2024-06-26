// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: proto/search/search.proto

package search

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sdk "search-service/proto/sdk"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageSearchPrefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paginate   *sdk.Pagination `protobuf:"bytes,1,opt,name=paginate,proto3" json:"paginate,omitempty"`
	SearchText string          `protobuf:"bytes,2,opt,name=searchText,proto3" json:"searchText,omitempty"`
}

func (x *MessageSearchPrefix) Reset() {
	*x = MessageSearchPrefix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageSearchPrefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSearchPrefix) ProtoMessage() {}

func (x *MessageSearchPrefix) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSearchPrefix.ProtoReflect.Descriptor instead.
func (*MessageSearchPrefix) Descriptor() ([]byte, []int) {
	return file_proto_search_search_proto_rawDescGZIP(), []int{0}
}

func (x *MessageSearchPrefix) GetPaginate() *sdk.Pagination {
	if x != nil {
		return x.Paginate
	}
	return nil
}

func (x *MessageSearchPrefix) GetSearchText() string {
	if x != nil {
		return x.SearchText
	}
	return ""
}

type MsgIP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paginate  *sdk.Pagination `protobuf:"bytes,1,opt,name=paginate,proto3" json:"paginate,omitempty"`
	IpAddress string          `protobuf:"bytes,2,opt,name=ipAddress,proto3" json:"ipAddress,omitempty"`
}

func (x *MsgIP) Reset() {
	*x = MsgIP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgIP) ProtoMessage() {}

func (x *MsgIP) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgIP.ProtoReflect.Descriptor instead.
func (*MsgIP) Descriptor() ([]byte, []int) {
	return file_proto_search_search_proto_rawDescGZIP(), []int{1}
}

func (x *MsgIP) GetPaginate() *sdk.Pagination {
	if x != nil {
		return x.Paginate
	}
	return nil
}

func (x *MsgIP) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

type MsgSuggestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paginate *sdk.Pagination `protobuf:"bytes,1,opt,name=paginate,proto3" json:"paginate,omitempty"`
	UserId   *string         `protobuf:"bytes,2,opt,name=userId,proto3,oneof" json:"userId,omitempty"`
}

func (x *MsgSuggestion) Reset() {
	*x = MsgSuggestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSuggestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSuggestion) ProtoMessage() {}

func (x *MsgSuggestion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSuggestion.ProtoReflect.Descriptor instead.
func (*MsgSuggestion) Descriptor() ([]byte, []int) {
	return file_proto_search_search_proto_rawDescGZIP(), []int{2}
}

func (x *MsgSuggestion) GetPaginate() *sdk.Pagination {
	if x != nil {
		return x.Paginate
	}
	return nil
}

func (x *MsgSuggestion) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

type MsgSearchProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paginate    *sdk.Pagination `protobuf:"bytes,1,opt,name=paginate,proto3" json:"paginate,omitempty"`
	QueryFields *MsgProperty    `protobuf:"bytes,2,opt,name=queryFields,proto3,oneof" json:"queryFields,omitempty"`
}

func (x *MsgSearchProperty) Reset() {
	*x = MsgSearchProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSearchProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSearchProperty) ProtoMessage() {}

func (x *MsgSearchProperty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSearchProperty.ProtoReflect.Descriptor instead.
func (*MsgSearchProperty) Descriptor() ([]byte, []int) {
	return file_proto_search_search_proto_rawDescGZIP(), []int{3}
}

func (x *MsgSearchProperty) GetPaginate() *sdk.Pagination {
	if x != nil {
		return x.Paginate
	}
	return nil
}

func (x *MsgSearchProperty) GetQueryFields() *MsgProperty {
	if x != nil {
		return x.QueryFields
	}
	return nil
}

type MsgProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               *string       `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	HostId           *string       `protobuf:"bytes,2,opt,name=hostId,proto3,oneof" json:"hostId,omitempty"`
	HostName         *string       `protobuf:"bytes,44,opt,name=hostName,proto3,oneof" json:"hostName,omitempty"`
	HostAvatar       *string       `protobuf:"bytes,46,opt,name=hostAvatar,proto3,oneof" json:"hostAvatar,omitempty"`
	PropertyType     *string       `protobuf:"bytes,3,opt,name=propertyType,proto3,oneof" json:"propertyType,omitempty"`
	Status           *string       `protobuf:"bytes,45,opt,name=status,proto3,oneof" json:"status,omitempty"`
	OverallRate      *float32      `protobuf:"fixed32,4,opt,name=overallRate,proto3,oneof" json:"overallRate,omitempty"`
	NextCheckInDate  int64         `protobuf:"varint,50,opt,name=nextCheckInDate,proto3" json:"nextCheckInDate,omitempty"`
	NextCheckoutDate int64         `protobuf:"varint,51,opt,name=nextCheckoutDate,proto3" json:"nextCheckoutDate,omitempty"`
	MaxGuests        *int32        `protobuf:"varint,5,opt,name=maxGuests,proto3,oneof" json:"maxGuests,omitempty"`
	MaxPets          *int32        `protobuf:"varint,24,opt,name=maxPets,proto3,oneof" json:"maxPets,omitempty"`
	NumBeds          *int32        `protobuf:"varint,6,opt,name=numBeds,proto3,oneof" json:"numBeds,omitempty"`
	NumBedrooms      *int32        `protobuf:"varint,7,opt,name=numBedrooms,proto3,oneof" json:"numBedrooms,omitempty"`
	NumBathrooms     *int32        `protobuf:"varint,8,opt,name=numBathrooms,proto3,oneof" json:"numBathrooms,omitempty"`
	IsGuestFavor     *bool         `protobuf:"varint,10,opt,name=isGuestFavor,proto3,oneof" json:"isGuestFavor,omitempty"`
	IsAllowPet       *bool         `protobuf:"varint,11,opt,name=isAllowPet,proto3,oneof" json:"isAllowPet,omitempty"`
	IsSelfCheckIn    *bool         `protobuf:"varint,12,opt,name=isSelfCheckIn,proto3,oneof" json:"isSelfCheckIn,omitempty"`
	IsInstantBook    *bool         `protobuf:"varint,13,opt,name=isInstantBook,proto3,oneof" json:"isInstantBook,omitempty"`
	Title            *string       `protobuf:"bytes,14,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Body             *string       `protobuf:"bytes,15,opt,name=body,proto3,oneof" json:"body,omitempty"`
	Address          *string       `protobuf:"bytes,41,opt,name=address,proto3,oneof" json:"address,omitempty"`
	CityCode         *string       `protobuf:"bytes,16,opt,name=cityCode,proto3,oneof" json:"cityCode,omitempty"`
	NationCode       *string       `protobuf:"bytes,17,opt,name=nationCode,proto3,oneof" json:"nationCode,omitempty"`
	Lat              *string       `protobuf:"bytes,18,opt,name=lat,proto3,oneof" json:"lat,omitempty"`
	Long             *string       `protobuf:"bytes,19,opt,name=long,proto3,oneof" json:"long,omitempty"`
	ServiceFee       *float32      `protobuf:"fixed32,20,opt,name=serviceFee,proto3,oneof" json:"serviceFee,omitempty"`
	NightPriceMin    *float64      `protobuf:"fixed64,21,opt,name=nightPriceMin,proto3,oneof" json:"nightPriceMin,omitempty"`
	NightPriceMax    *float64      `protobuf:"fixed64,22,opt,name=nightPriceMax,proto3,oneof" json:"nightPriceMax,omitempty"`
	Amenities        []*MsgAmenity `protobuf:"bytes,23,rep,name=amenities,proto3" json:"amenities,omitempty"`
	UserId           *string       `protobuf:"bytes,27,opt,name=userId,proto3,oneof" json:"userId,omitempty"`
}

func (x *MsgProperty) Reset() {
	*x = MsgProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgProperty) ProtoMessage() {}

func (x *MsgProperty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgProperty.ProtoReflect.Descriptor instead.
func (*MsgProperty) Descriptor() ([]byte, []int) {
	return file_proto_search_search_proto_rawDescGZIP(), []int{4}
}

func (x *MsgProperty) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *MsgProperty) GetHostId() string {
	if x != nil && x.HostId != nil {
		return *x.HostId
	}
	return ""
}

func (x *MsgProperty) GetHostName() string {
	if x != nil && x.HostName != nil {
		return *x.HostName
	}
	return ""
}

func (x *MsgProperty) GetHostAvatar() string {
	if x != nil && x.HostAvatar != nil {
		return *x.HostAvatar
	}
	return ""
}

func (x *MsgProperty) GetPropertyType() string {
	if x != nil && x.PropertyType != nil {
		return *x.PropertyType
	}
	return ""
}

func (x *MsgProperty) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *MsgProperty) GetOverallRate() float32 {
	if x != nil && x.OverallRate != nil {
		return *x.OverallRate
	}
	return 0
}

func (x *MsgProperty) GetNextCheckInDate() int64 {
	if x != nil {
		return x.NextCheckInDate
	}
	return 0
}

func (x *MsgProperty) GetNextCheckoutDate() int64 {
	if x != nil {
		return x.NextCheckoutDate
	}
	return 0
}

func (x *MsgProperty) GetMaxGuests() int32 {
	if x != nil && x.MaxGuests != nil {
		return *x.MaxGuests
	}
	return 0
}

func (x *MsgProperty) GetMaxPets() int32 {
	if x != nil && x.MaxPets != nil {
		return *x.MaxPets
	}
	return 0
}

func (x *MsgProperty) GetNumBeds() int32 {
	if x != nil && x.NumBeds != nil {
		return *x.NumBeds
	}
	return 0
}

func (x *MsgProperty) GetNumBedrooms() int32 {
	if x != nil && x.NumBedrooms != nil {
		return *x.NumBedrooms
	}
	return 0
}

func (x *MsgProperty) GetNumBathrooms() int32 {
	if x != nil && x.NumBathrooms != nil {
		return *x.NumBathrooms
	}
	return 0
}

func (x *MsgProperty) GetIsGuestFavor() bool {
	if x != nil && x.IsGuestFavor != nil {
		return *x.IsGuestFavor
	}
	return false
}

func (x *MsgProperty) GetIsAllowPet() bool {
	if x != nil && x.IsAllowPet != nil {
		return *x.IsAllowPet
	}
	return false
}

func (x *MsgProperty) GetIsSelfCheckIn() bool {
	if x != nil && x.IsSelfCheckIn != nil {
		return *x.IsSelfCheckIn
	}
	return false
}

func (x *MsgProperty) GetIsInstantBook() bool {
	if x != nil && x.IsInstantBook != nil {
		return *x.IsInstantBook
	}
	return false
}

func (x *MsgProperty) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *MsgProperty) GetBody() string {
	if x != nil && x.Body != nil {
		return *x.Body
	}
	return ""
}

func (x *MsgProperty) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *MsgProperty) GetCityCode() string {
	if x != nil && x.CityCode != nil {
		return *x.CityCode
	}
	return ""
}

func (x *MsgProperty) GetNationCode() string {
	if x != nil && x.NationCode != nil {
		return *x.NationCode
	}
	return ""
}

func (x *MsgProperty) GetLat() string {
	if x != nil && x.Lat != nil {
		return *x.Lat
	}
	return ""
}

func (x *MsgProperty) GetLong() string {
	if x != nil && x.Long != nil {
		return *x.Long
	}
	return ""
}

func (x *MsgProperty) GetServiceFee() float32 {
	if x != nil && x.ServiceFee != nil {
		return *x.ServiceFee
	}
	return 0
}

func (x *MsgProperty) GetNightPriceMin() float64 {
	if x != nil && x.NightPriceMin != nil {
		return *x.NightPriceMin
	}
	return 0
}

func (x *MsgProperty) GetNightPriceMax() float64 {
	if x != nil && x.NightPriceMax != nil {
		return *x.NightPriceMax
	}
	return 0
}

func (x *MsgProperty) GetAmenities() []*MsgAmenity {
	if x != nil {
		return x.Amenities
	}
	return nil
}

func (x *MsgProperty) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

type MsgAmenity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name        *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Description *string `protobuf:"bytes,3,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Icon        *string `protobuf:"bytes,4,opt,name=icon,proto3,oneof" json:"icon,omitempty"`
}

func (x *MsgAmenity) Reset() {
	*x = MsgAmenity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_search_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAmenity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAmenity) ProtoMessage() {}

func (x *MsgAmenity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_search_search_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgAmenity.ProtoReflect.Descriptor instead.
func (*MsgAmenity) Descriptor() ([]byte, []int) {
	return file_proto_search_search_proto_rawDescGZIP(), []int{5}
}

func (x *MsgAmenity) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *MsgAmenity) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *MsgAmenity) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *MsgAmenity) GetIcon() string {
	if x != nil && x.Icon != nil {
		return *x.Icon
	}
	return ""
}

var File_proto_search_search_proto protoreflect.FileDescriptor

var file_proto_search_search_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x62, 0x0a, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54,
	0x65, 0x78, 0x74, 0x22, 0x52, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x50, 0x12, 0x2b, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x70, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x64, 0x0a, 0x0d, 0x4d, 0x73, 0x67, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x93, 0x01,
	0x0a, 0x11, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x12, 0x41, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x48, 0x00, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x22, 0xa6, 0x0b, 0x0a, 0x0b, 0x4d, 0x73, 0x67, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x68, 0x6f, 0x73, 0x74, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0a, 0x68, 0x6f,
	0x73, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x04, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x2d,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01,
	0x01, 0x12, 0x25, 0x0a, 0x0b, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x48, 0x06, 0x52, 0x0b, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c,
	0x6c, 0x52, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x33, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6e, 0x65,
	0x78, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x21,
	0x0a, 0x09, 0x6d, 0x61, 0x78, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x07, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x50, 0x65, 0x74, 0x73, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x08, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x50, 0x65, 0x74, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x42, 0x65, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x09, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x42, 0x65, 0x64, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x25, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x42, 0x65, 0x64, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x0a, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x42, 0x65, 0x64, 0x72, 0x6f,
	0x6f, 0x6d, 0x73, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x42, 0x61, 0x74,
	0x68, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0b, 0x52, 0x0c,
	0x6e, 0x75, 0x6d, 0x42, 0x61, 0x74, 0x68, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x27, 0x0a, 0x0c, 0x69, 0x73, 0x47, 0x75, 0x65, 0x73, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x48, 0x0c, 0x52, 0x0c, 0x69, 0x73, 0x47, 0x75, 0x65, 0x73, 0x74,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x69, 0x73, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x50, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x48, 0x0d, 0x52, 0x0a,
	0x69, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a,
	0x0d, 0x69, 0x73, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x0e, 0x52, 0x0d, 0x69, 0x73, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x49, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x69, 0x73, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x0f, 0x52, 0x0d, 0x69, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x10, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x48, 0x11, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x48, 0x12, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x69, 0x74, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x48, 0x13, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x14, 0x52, 0x0a, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03,
	0x6c, 0x61, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x48, 0x15, 0x52, 0x03, 0x6c, 0x61, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x16, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02,
	0x48, 0x17, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x29, 0x0a, 0x0d, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4d,
	0x69, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x48, 0x18, 0x52, 0x0d, 0x6e, 0x69, 0x67, 0x68,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4d, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d,
	0x6e, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x78, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x19, 0x52, 0x0d, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x4d, 0x61, 0x78, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x09, 0x61, 0x6d, 0x65, 0x6e, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x41, 0x6d,
	0x65, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x09, 0x61, 0x6d, 0x65, 0x6e, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x1b, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x1a, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a,
	0x03, 0x5f, 0x69, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x68, 0x6f, 0x73, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x0f, 0x0a, 0x0d, 0x5f,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x76, 0x65, 0x72,
	0x61, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6d, 0x61, 0x78, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x61, 0x78, 0x50, 0x65, 0x74,
	0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6e, 0x75, 0x6d, 0x42, 0x65, 0x64, 0x73, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x6e, 0x75, 0x6d, 0x42, 0x65, 0x64, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x42, 0x0f, 0x0a,
	0x0d, 0x5f, 0x6e, 0x75, 0x6d, 0x42, 0x61, 0x74, 0x68, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x69, 0x73, 0x47, 0x75, 0x65, 0x73, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x65, 0x74, 0x42, 0x10,
	0x0a, 0x0e, 0x5f, 0x69, 0x73, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e,
	0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x69, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x6c, 0x61, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4d, 0x69, 0x6e, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4d, 0x61,
	0x78, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa3, 0x01, 0x0a,
	0x0a, 0x4d, 0x73, 0x67, 0x41, 0x6d, 0x65, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x13, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x69, 0x63,
	0x6f, 0x6e, 0x32, 0x9d, 0x02, 0x0a, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x10, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x11, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x22,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x1a, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x49, 0x50, 0x1a, 0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x20, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x73,
	0x67, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x1a,
	0x11, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_search_search_proto_rawDescOnce sync.Once
	file_proto_search_search_proto_rawDescData = file_proto_search_search_proto_rawDesc
)

func file_proto_search_search_proto_rawDescGZIP() []byte {
	file_proto_search_search_proto_rawDescOnce.Do(func() {
		file_proto_search_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_search_search_proto_rawDescData)
	})
	return file_proto_search_search_proto_rawDescData
}

var file_proto_search_search_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_search_search_proto_goTypes = []interface{}{
	(*MessageSearchPrefix)(nil), // 0: searchService.MessageSearchPrefix
	(*MsgIP)(nil),               // 1: searchService.MsgIP
	(*MsgSuggestion)(nil),       // 2: searchService.MsgSuggestion
	(*MsgSearchProperty)(nil),   // 3: searchService.MsgSearchProperty
	(*MsgProperty)(nil),         // 4: searchService.MsgProperty
	(*MsgAmenity)(nil),          // 5: searchService.MsgAmenity
	(*sdk.Pagination)(nil),      // 6: sdk.Pagination
	(*sdk.BaseResponse)(nil),    // 7: sdk.BaseResponse
}
var file_proto_search_search_proto_depIdxs = []int32{
	6,  // 0: searchService.MessageSearchPrefix.paginate:type_name -> sdk.Pagination
	6,  // 1: searchService.MsgIP.paginate:type_name -> sdk.Pagination
	6,  // 2: searchService.MsgSuggestion.paginate:type_name -> sdk.Pagination
	6,  // 3: searchService.MsgSearchProperty.paginate:type_name -> sdk.Pagination
	4,  // 4: searchService.MsgSearchProperty.queryFields:type_name -> searchService.MsgProperty
	5,  // 5: searchService.MsgProperty.amenities:type_name -> searchService.MsgAmenity
	2,  // 6: searchService.searchService.RenderSuggestion:input_type -> searchService.MsgSuggestion
	0,  // 7: searchService.searchService.searchTitlePrefix:input_type -> searchService.MessageSearchPrefix
	1,  // 8: searchService.searchService.getNation:input_type -> searchService.MsgIP
	3,  // 9: searchService.searchService.searchProperty:input_type -> searchService.MsgSearchProperty
	7,  // 10: searchService.searchService.RenderSuggestion:output_type -> sdk.BaseResponse
	7,  // 11: searchService.searchService.searchTitlePrefix:output_type -> sdk.BaseResponse
	7,  // 12: searchService.searchService.getNation:output_type -> sdk.BaseResponse
	7,  // 13: searchService.searchService.searchProperty:output_type -> sdk.BaseResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_search_search_proto_init() }
func file_proto_search_search_proto_init() {
	if File_proto_search_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_search_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageSearchPrefix); i {
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
		file_proto_search_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgIP); i {
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
		file_proto_search_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSuggestion); i {
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
		file_proto_search_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSearchProperty); i {
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
		file_proto_search_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgProperty); i {
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
		file_proto_search_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAmenity); i {
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
	file_proto_search_search_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_proto_search_search_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_proto_search_search_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_proto_search_search_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_search_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_search_search_proto_goTypes,
		DependencyIndexes: file_proto_search_search_proto_depIdxs,
		MessageInfos:      file_proto_search_search_proto_msgTypes,
	}.Build()
	File_proto_search_search_proto = out.File
	file_proto_search_search_proto_rawDesc = nil
	file_proto_search_search_proto_goTypes = nil
	file_proto_search_search_proto_depIdxs = nil
}
