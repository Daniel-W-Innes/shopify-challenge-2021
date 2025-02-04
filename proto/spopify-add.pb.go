// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: proto/spopify-add.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId int32  `protobuf:"varint,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Error         string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spopify_add_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spopify_add_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_proto_spopify_add_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetTransactionId() int32 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

func (x *Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ImageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tags      []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	IsPrivate bool     `protobuf:"varint,3,opt,name=is_private,json=isPrivate,proto3" json:"is_private,omitempty"`
}

func (x *ImageInfo) Reset() {
	*x = ImageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spopify_add_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInfo) ProtoMessage() {}

func (x *ImageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spopify_add_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInfo.ProtoReflect.Descriptor instead.
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return file_proto_spopify_add_proto_rawDescGZIP(), []int{1}
}

func (x *ImageInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImageInfo) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ImageInfo) GetIsPrivate() bool {
	if x != nil {
		return x.IsPrivate
	}
	return false
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId int32      `protobuf:"varint,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	ImageData     []byte     `protobuf:"bytes,2,opt,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
	Info          *ImageInfo `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spopify_add_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spopify_add_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_proto_spopify_add_proto_rawDescGZIP(), []int{2}
}

func (x *Image) GetTransactionId() int32 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

func (x *Image) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

func (x *Image) GetInfo() *ImageInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type ImageTile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId int32  `protobuf:"varint,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	ImageData     []byte `protobuf:"bytes,2,opt,name=image_data,json=imageData,proto3" json:"image_data,omitempty"`
	Number        int32  `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *ImageTile) Reset() {
	*x = ImageTile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spopify_add_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageTile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageTile) ProtoMessage() {}

func (x *ImageTile) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spopify_add_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageTile.ProtoReflect.Descriptor instead.
func (*ImageTile) Descriptor() ([]byte, []int) {
	return file_proto_spopify_add_proto_rawDescGZIP(), []int{3}
}

func (x *ImageTile) GetTransactionId() int32 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

func (x *ImageTile) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

func (x *ImageTile) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ImageTileRegistration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId int32      `protobuf:"varint,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Info          *ImageInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	NumX          int32      `protobuf:"varint,3,opt,name=num_x,json=numX,proto3" json:"num_x,omitempty"`
	NumY          int32      `protobuf:"varint,4,opt,name=num_y,json=numY,proto3" json:"num_y,omitempty"`
}

func (x *ImageTileRegistration) Reset() {
	*x = ImageTileRegistration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_spopify_add_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageTileRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageTileRegistration) ProtoMessage() {}

func (x *ImageTileRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_proto_spopify_add_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageTileRegistration.ProtoReflect.Descriptor instead.
func (*ImageTileRegistration) Descriptor() ([]byte, []int) {
	return file_proto_spopify_add_proto_rawDescGZIP(), []int{4}
}

func (x *ImageTileRegistration) GetTransactionId() int32 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

func (x *ImageTileRegistration) GetInfo() *ImageInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ImageTileRegistration) GetNumX() int32 {
	if x != nil {
		return x.NumX
	}
	return 0
}

func (x *ImageTileRegistration) GetNumY() int32 {
	if x != nil {
		return x.NumY
	}
	return 0
}

var File_proto_spopify_add_proto protoreflect.FileDescriptor

var file_proto_spopify_add_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2d,
	0x61, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x52, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x6d, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x69, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x69,
	0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x88, 0x01, 0x0a, 0x15, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x13, 0x0a, 0x05, 0x6e, 0x75, 0x6d, 0x5f, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x6e, 0x75, 0x6d, 0x58, 0x12, 0x13, 0x0a, 0x05, 0x6e, 0x75, 0x6d, 0x5f, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x59, 0x32, 0xfa, 0x01, 0x0a, 0x0a,
	0x53, 0x68, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x41, 0x64, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x41, 0x64, 0x64, 0x12, 0x06, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x09,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x09, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x41, 0x64, 0x64, 0x12, 0x06, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x12, 0x39, 0x0a, 0x12, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x54, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x11,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6c,
	0x65, 0x12, 0x16, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x2c, 0x0a, 0x0d, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x41, 0x64, 0x64, 0x54, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x0a, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x54, 0x69, 0x6c, 0x65, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x6e, 0x69, 0x65, 0x6c, 0x2d, 0x57, 0x2d,
	0x49, 0x6e, 0x6e, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x69, 0x66, 0x79, 0x2d, 0x61, 0x64,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_spopify_add_proto_rawDescOnce sync.Once
	file_proto_spopify_add_proto_rawDescData = file_proto_spopify_add_proto_rawDesc
)

func file_proto_spopify_add_proto_rawDescGZIP() []byte {
	file_proto_spopify_add_proto_rawDescOnce.Do(func() {
		file_proto_spopify_add_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_spopify_add_proto_rawDescData)
	})
	return file_proto_spopify_add_proto_rawDescData
}

var file_proto_spopify_add_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_spopify_add_proto_goTypes = []interface{}{
	(*Response)(nil),              // 0: Response
	(*ImageInfo)(nil),             // 1: ImageInfo
	(*Image)(nil),                 // 2: Image
	(*ImageTile)(nil),             // 3: ImageTile
	(*ImageTileRegistration)(nil), // 4: ImageTileRegistration
}
var file_proto_spopify_add_proto_depIdxs = []int32{
	1, // 0: Image.info:type_name -> ImageInfo
	1, // 1: ImageTileRegistration.info:type_name -> ImageInfo
	2, // 2: ShopifyAdd.ImageAdd:input_type -> Image
	2, // 3: ShopifyAdd.ImagesAdd:input_type -> Image
	4, // 4: ShopifyAdd.ImageRegisterTiles:input_type -> ImageTileRegistration
	4, // 5: ShopifyAdd.ImageRegisterTile:input_type -> ImageTileRegistration
	3, // 6: ShopifyAdd.ImageAddTiles:input_type -> ImageTile
	0, // 7: ShopifyAdd.ImageAdd:output_type -> Response
	0, // 8: ShopifyAdd.ImagesAdd:output_type -> Response
	0, // 9: ShopifyAdd.ImageRegisterTiles:output_type -> Response
	0, // 10: ShopifyAdd.ImageRegisterTile:output_type -> Response
	0, // 11: ShopifyAdd.ImageAddTiles:output_type -> Response
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_spopify_add_proto_init() }
func file_proto_spopify_add_proto_init() {
	if File_proto_spopify_add_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_spopify_add_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_spopify_add_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageInfo); i {
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
		file_proto_spopify_add_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_proto_spopify_add_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageTile); i {
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
		file_proto_spopify_add_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageTileRegistration); i {
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
			RawDescriptor: file_proto_spopify_add_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_spopify_add_proto_goTypes,
		DependencyIndexes: file_proto_spopify_add_proto_depIdxs,
		MessageInfos:      file_proto_spopify_add_proto_msgTypes,
	}.Build()
	File_proto_spopify_add_proto = out.File
	file_proto_spopify_add_proto_rawDesc = nil
	file_proto_spopify_add_proto_goTypes = nil
	file_proto_spopify_add_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShopifyAddClient is the client API for ShopifyAdd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShopifyAddClient interface {
	ImageAdd(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Response, error)
	ImagesAdd(ctx context.Context, opts ...grpc.CallOption) (ShopifyAdd_ImagesAddClient, error)
	ImageRegisterTiles(ctx context.Context, in *ImageTileRegistration, opts ...grpc.CallOption) (*Response, error)
	ImageRegisterTile(ctx context.Context, opts ...grpc.CallOption) (ShopifyAdd_ImageRegisterTileClient, error)
	ImageAddTiles(ctx context.Context, opts ...grpc.CallOption) (ShopifyAdd_ImageAddTilesClient, error)
}

type shopifyAddClient struct {
	cc grpc.ClientConnInterface
}

func NewShopifyAddClient(cc grpc.ClientConnInterface) ShopifyAddClient {
	return &shopifyAddClient{cc}
}

func (c *shopifyAddClient) ImageAdd(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ShopifyAdd/ImageAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopifyAddClient) ImagesAdd(ctx context.Context, opts ...grpc.CallOption) (ShopifyAdd_ImagesAddClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ShopifyAdd_serviceDesc.Streams[0], "/ShopifyAdd/ImagesAdd", opts...)
	if err != nil {
		return nil, err
	}
	x := &shopifyAddImagesAddClient{stream}
	return x, nil
}

type ShopifyAdd_ImagesAddClient interface {
	Send(*Image) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type shopifyAddImagesAddClient struct {
	grpc.ClientStream
}

func (x *shopifyAddImagesAddClient) Send(m *Image) error {
	return x.ClientStream.SendMsg(m)
}

func (x *shopifyAddImagesAddClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *shopifyAddClient) ImageRegisterTiles(ctx context.Context, in *ImageTileRegistration, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ShopifyAdd/ImageRegisterTiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopifyAddClient) ImageRegisterTile(ctx context.Context, opts ...grpc.CallOption) (ShopifyAdd_ImageRegisterTileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ShopifyAdd_serviceDesc.Streams[1], "/ShopifyAdd/ImageRegisterTile", opts...)
	if err != nil {
		return nil, err
	}
	x := &shopifyAddImageRegisterTileClient{stream}
	return x, nil
}

type ShopifyAdd_ImageRegisterTileClient interface {
	Send(*ImageTileRegistration) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type shopifyAddImageRegisterTileClient struct {
	grpc.ClientStream
}

func (x *shopifyAddImageRegisterTileClient) Send(m *ImageTileRegistration) error {
	return x.ClientStream.SendMsg(m)
}

func (x *shopifyAddImageRegisterTileClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *shopifyAddClient) ImageAddTiles(ctx context.Context, opts ...grpc.CallOption) (ShopifyAdd_ImageAddTilesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ShopifyAdd_serviceDesc.Streams[2], "/ShopifyAdd/ImageAddTiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &shopifyAddImageAddTilesClient{stream}
	return x, nil
}

type ShopifyAdd_ImageAddTilesClient interface {
	Send(*ImageTile) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type shopifyAddImageAddTilesClient struct {
	grpc.ClientStream
}

func (x *shopifyAddImageAddTilesClient) Send(m *ImageTile) error {
	return x.ClientStream.SendMsg(m)
}

func (x *shopifyAddImageAddTilesClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ShopifyAddServer is the server API for ShopifyAdd service.
type ShopifyAddServer interface {
	ImageAdd(context.Context, *Image) (*Response, error)
	ImagesAdd(ShopifyAdd_ImagesAddServer) error
	ImageRegisterTiles(context.Context, *ImageTileRegistration) (*Response, error)
	ImageRegisterTile(ShopifyAdd_ImageRegisterTileServer) error
	ImageAddTiles(ShopifyAdd_ImageAddTilesServer) error
}

// UnimplementedShopifyAddServer can be embedded to have forward compatible implementations.
type UnimplementedShopifyAddServer struct {
}

func (*UnimplementedShopifyAddServer) ImageAdd(context.Context, *Image) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImageAdd not implemented")
}
func (*UnimplementedShopifyAddServer) ImagesAdd(ShopifyAdd_ImagesAddServer) error {
	return status.Errorf(codes.Unimplemented, "method ImagesAdd not implemented")
}
func (*UnimplementedShopifyAddServer) ImageRegisterTiles(context.Context, *ImageTileRegistration) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImageRegisterTiles not implemented")
}
func (*UnimplementedShopifyAddServer) ImageRegisterTile(ShopifyAdd_ImageRegisterTileServer) error {
	return status.Errorf(codes.Unimplemented, "method ImageRegisterTile not implemented")
}
func (*UnimplementedShopifyAddServer) ImageAddTiles(ShopifyAdd_ImageAddTilesServer) error {
	return status.Errorf(codes.Unimplemented, "method ImageAddTiles not implemented")
}

func RegisterShopifyAddServer(s *grpc.Server, srv ShopifyAddServer) {
	s.RegisterService(&_ShopifyAdd_serviceDesc, srv)
}

func _ShopifyAdd_ImageAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Image)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopifyAddServer).ImageAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopifyAdd/ImageAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopifyAddServer).ImageAdd(ctx, req.(*Image))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopifyAdd_ImagesAdd_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ShopifyAddServer).ImagesAdd(&shopifyAddImagesAddServer{stream})
}

type ShopifyAdd_ImagesAddServer interface {
	Send(*Response) error
	Recv() (*Image, error)
	grpc.ServerStream
}

type shopifyAddImagesAddServer struct {
	grpc.ServerStream
}

func (x *shopifyAddImagesAddServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *shopifyAddImagesAddServer) Recv() (*Image, error) {
	m := new(Image)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ShopifyAdd_ImageRegisterTiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageTileRegistration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopifyAddServer).ImageRegisterTiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopifyAdd/ImageRegisterTiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopifyAddServer).ImageRegisterTiles(ctx, req.(*ImageTileRegistration))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopifyAdd_ImageRegisterTile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ShopifyAddServer).ImageRegisterTile(&shopifyAddImageRegisterTileServer{stream})
}

type ShopifyAdd_ImageRegisterTileServer interface {
	Send(*Response) error
	Recv() (*ImageTileRegistration, error)
	grpc.ServerStream
}

type shopifyAddImageRegisterTileServer struct {
	grpc.ServerStream
}

func (x *shopifyAddImageRegisterTileServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *shopifyAddImageRegisterTileServer) Recv() (*ImageTileRegistration, error) {
	m := new(ImageTileRegistration)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ShopifyAdd_ImageAddTiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ShopifyAddServer).ImageAddTiles(&shopifyAddImageAddTilesServer{stream})
}

type ShopifyAdd_ImageAddTilesServer interface {
	Send(*Response) error
	Recv() (*ImageTile, error)
	grpc.ServerStream
}

type shopifyAddImageAddTilesServer struct {
	grpc.ServerStream
}

func (x *shopifyAddImageAddTilesServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *shopifyAddImageAddTilesServer) Recv() (*ImageTile, error) {
	m := new(ImageTile)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ShopifyAdd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShopifyAdd",
	HandlerType: (*ShopifyAddServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImageAdd",
			Handler:    _ShopifyAdd_ImageAdd_Handler,
		},
		{
			MethodName: "ImageRegisterTiles",
			Handler:    _ShopifyAdd_ImageRegisterTiles_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ImagesAdd",
			Handler:       _ShopifyAdd_ImagesAdd_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ImageRegisterTile",
			Handler:       _ShopifyAdd_ImageRegisterTile_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ImageAddTiles",
			Handler:       _ShopifyAdd_ImageAddTiles_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/spopify-add.proto",
}
