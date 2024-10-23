// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: proto/restaurant.proto

package proto

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

// menu items
type Category int32

const (
	Category_UNKNOWN     Category = 0
	Category_APPETIZER   Category = 1
	Category_MAIN_COURSE Category = 2
	Category_DESSERT     Category = 3
	Category_BEVERAGE    Category = 4
	Category_SIDE_DISH   Category = 5
)

// Enum value maps for Category.
var (
	Category_name = map[int32]string{
		0: "UNKNOWN",
		1: "APPETIZER",
		2: "MAIN_COURSE",
		3: "DESSERT",
		4: "BEVERAGE",
		5: "SIDE_DISH",
	}
	Category_value = map[string]int32{
		"UNKNOWN":     0,
		"APPETIZER":   1,
		"MAIN_COURSE": 2,
		"DESSERT":     3,
		"BEVERAGE":    4,
		"SIDE_DISH":   5,
	}
)

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}

func (x Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Category) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_restaurant_proto_enumTypes[0].Descriptor()
}

func (Category) Type() protoreflect.EnumType {
	return &file_proto_restaurant_proto_enumTypes[0]
}

func (x Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Category.Descriptor instead.
func (Category) EnumDescriptor() ([]byte, []int) {
	return file_proto_restaurant_proto_rawDescGZIP(), []int{0}
}

type MenuItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32  `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Category    Category `protobuf:"varint,4,opt,name=category,proto3,enum=Category" json:"category,omitempty"`
	Available   bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *MenuItem) Reset() {
	*x = MenuItem{}
	mi := &file_proto_restaurant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MenuItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuItem) ProtoMessage() {}

func (x *MenuItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuItem.ProtoReflect.Descriptor instead.
func (*MenuItem) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_proto_rawDescGZIP(), []int{0}
}

func (x *MenuItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MenuItem) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *MenuItem) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_UNKNOWN
}

func (x *MenuItem) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

type Restaurant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location      string      `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	ContactNumber string      `protobuf:"bytes,3,opt,name=contact_number,json=contactNumber,proto3" json:"contact_number,omitempty"`
	Menu          []*MenuItem `protobuf:"bytes,4,rep,name=menu,proto3" json:"menu,omitempty"`
}

func (x *Restaurant) Reset() {
	*x = Restaurant{}
	mi := &file_proto_restaurant_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Restaurant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restaurant) ProtoMessage() {}

func (x *Restaurant) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restaurant.ProtoReflect.Descriptor instead.
func (*Restaurant) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_proto_rawDescGZIP(), []int{1}
}

func (x *Restaurant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Restaurant) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Restaurant) GetContactNumber() string {
	if x != nil {
		return x.ContactNumber
	}
	return ""
}

func (x *Restaurant) GetMenu() []*MenuItem {
	if x != nil {
		return x.Menu
	}
	return nil
}

type RestaurantList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restaurants []*Restaurant `protobuf:"bytes,1,rep,name=restaurants,proto3" json:"restaurants,omitempty"` // list of Restu messages
}

func (x *RestaurantList) Reset() {
	*x = RestaurantList{}
	mi := &file_proto_restaurant_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RestaurantList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestaurantList) ProtoMessage() {}

func (x *RestaurantList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestaurantList.ProtoReflect.Descriptor instead.
func (*RestaurantList) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_proto_rawDescGZIP(), []int{2}
}

func (x *RestaurantList) GetRestaurants() []*Restaurant {
	if x != nil {
		return x.Restaurants
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_restaurant_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_proto_rawDescGZIP(), []int{3}
}

type NameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NameRequest) Reset() {
	*x = NameRequest{}
	mi := &file_proto_restaurant_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameRequest) ProtoMessage() {}

func (x *NameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameRequest.ProtoReflect.Descriptor instead.
func (*NameRequest) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_proto_rawDescGZIP(), []int{4}
}

func (x *NameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category Category `protobuf:"varint,1,opt,name=category,proto3,enum=Category" json:"category,omitempty"`
}

func (x *CategoryRequest) Reset() {
	*x = CategoryRequest{}
	mi := &file_proto_restaurant_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryRequest) ProtoMessage() {}

func (x *CategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryRequest.ProtoReflect.Descriptor instead.
func (*CategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_proto_rawDescGZIP(), []int{5}
}

func (x *CategoryRequest) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_UNKNOWN
}

type MenuItemList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*MenuItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *MenuItemList) Reset() {
	*x = MenuItemList{}
	mi := &file_proto_restaurant_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MenuItemList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuItemList) ProtoMessage() {}

func (x *MenuItemList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuItemList.ProtoReflect.Descriptor instead.
func (*MenuItemList) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_proto_rawDescGZIP(), []int{6}
}

func (x *MenuItemList) GetItems() []*MenuItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_proto_restaurant_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_restaurant_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_restaurant_proto_rawDescGZIP(), []int{7}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_restaurant_proto protoreflect.FileDescriptor

var file_proto_restaurant_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x6e,
	0x75, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x04,
	0x6d, 0x65, 0x6e, 0x75, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x65, 0x6e,
	0x75, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x22, 0x3f, 0x0a, 0x0e, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52,
	0x0b, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x21, 0x0a, 0x0b, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x22, 0x2f, 0x0a, 0x0c, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x3e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2a, 0x61, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x41, 0x50, 0x50, 0x45, 0x54, 0x49, 0x5a, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4d,
	0x41, 0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x55, 0x52, 0x53, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x45, 0x53, 0x53, 0x45, 0x52, 0x54, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x45, 0x56,
	0x45, 0x52, 0x41, 0x47, 0x45, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x49, 0x44, 0x45, 0x5f,
	0x44, 0x49, 0x53, 0x48, 0x10, 0x05, 0x32, 0xd7, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0d,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x0b, 0x2e,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x10, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0d, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x1a, 0x5a, 0x18, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_restaurant_proto_rawDescOnce sync.Once
	file_proto_restaurant_proto_rawDescData = file_proto_restaurant_proto_rawDesc
)

func file_proto_restaurant_proto_rawDescGZIP() []byte {
	file_proto_restaurant_proto_rawDescOnce.Do(func() {
		file_proto_restaurant_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_restaurant_proto_rawDescData)
	})
	return file_proto_restaurant_proto_rawDescData
}

var file_proto_restaurant_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_restaurant_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_restaurant_proto_goTypes = []any{
	(Category)(0),           // 0: Category
	(*MenuItem)(nil),        // 1: MenuItem
	(*Restaurant)(nil),      // 2: Restaurant
	(*RestaurantList)(nil),  // 3: RestaurantList
	(*Empty)(nil),           // 4: Empty
	(*NameRequest)(nil),     // 5: NameRequest
	(*CategoryRequest)(nil), // 6: CategoryRequest
	(*MenuItemList)(nil),    // 7: MenuItemList
	(*Response)(nil),        // 8: Response
}
var file_proto_restaurant_proto_depIdxs = []int32{
	0, // 0: MenuItem.category:type_name -> Category
	1, // 1: Restaurant.menu:type_name -> MenuItem
	2, // 2: RestaurantList.restaurants:type_name -> Restaurant
	0, // 3: CategoryRequest.category:type_name -> Category
	1, // 4: MenuItemList.items:type_name -> MenuItem
	2, // 5: RestaurantService.AddRestaurant:input_type -> Restaurant
	4, // 6: RestaurantService.GetAllRestaurants:input_type -> Empty
	5, // 7: RestaurantService.GetRestaurantByName:input_type -> NameRequest
	6, // 8: RestaurantService.GetMenuItemsByCategory:input_type -> CategoryRequest
	8, // 9: RestaurantService.AddRestaurant:output_type -> Response
	3, // 10: RestaurantService.GetAllRestaurants:output_type -> RestaurantList
	2, // 11: RestaurantService.GetRestaurantByName:output_type -> Restaurant
	7, // 12: RestaurantService.GetMenuItemsByCategory:output_type -> MenuItemList
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_restaurant_proto_init() }
func file_proto_restaurant_proto_init() {
	if File_proto_restaurant_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_restaurant_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_restaurant_proto_goTypes,
		DependencyIndexes: file_proto_restaurant_proto_depIdxs,
		EnumInfos:         file_proto_restaurant_proto_enumTypes,
		MessageInfos:      file_proto_restaurant_proto_msgTypes,
	}.Build()
	File_proto_restaurant_proto = out.File
	file_proto_restaurant_proto_rawDesc = nil
	file_proto_restaurant_proto_goTypes = nil
	file_proto_restaurant_proto_depIdxs = nil
}