// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: pkg/proto/controller.proto

package grpc

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

// Connect Request from client
type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID             string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	ConnectClientAddress string `protobuf:"bytes,2,opt,name=connectClientAddress,proto3" json:"connectClientAddress,omitempty"`
	Username             string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password             string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	EnableRTSPRelay      bool   `protobuf:"varint,5,opt,name=enableRTSPRelay,proto3" json:"enableRTSPRelay,omitempty"`
	EnableRecord         bool   `protobuf:"varint,6,opt,name=enableRecord,proto3" json:"enableRecord,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_controller_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *ConnectRequest) GetConnectClientAddress() string {
	if x != nil {
		return x.ConnectClientAddress
	}
	return ""
}

func (x *ConnectRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ConnectRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ConnectRequest) GetEnableRTSPRelay() bool {
	if x != nil {
		return x.EnableRTSPRelay
	}
	return false
}

func (x *ConnectRequest) GetEnableRecord() bool {
	if x != nil {
		return x.EnableRecord
	}
	return false
}

type ConnectReplyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelayAddress string `protobuf:"bytes,1,opt,name=relayAddress,proto3" json:"relayAddress,omitempty"`
}

func (x *ConnectReplyData) Reset() {
	*x = ConnectReplyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectReplyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectReplyData) ProtoMessage() {}

func (x *ConnectReplyData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_controller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectReplyData.ProtoReflect.Descriptor instead.
func (*ConnectReplyData) Descriptor() ([]byte, []int) {
	return file_pkg_proto_controller_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectReplyData) GetRelayAddress() string {
	if x != nil {
		return x.RelayAddress
	}
	return ""
}

// Connect Reply from server
type ConnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *ConnectReplyData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ConnectReply) Reset() {
	*x = ConnectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_controller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectReply) ProtoMessage() {}

func (x *ConnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_controller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectReply.ProtoReflect.Descriptor instead.
func (*ConnectReply) Descriptor() ([]byte, []int) {
	return file_pkg_proto_controller_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ConnectReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConnectReply) GetData() *ConnectReplyData {
	if x != nil {
		return x.Data
	}
	return nil
}

// Disconnect Request from client
type DisconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID             string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	ConnectClientAddress string `protobuf:"bytes,2,opt,name=connectClientAddress,proto3" json:"connectClientAddress,omitempty"`
}

func (x *DisconnectRequest) Reset() {
	*x = DisconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_controller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectRequest) ProtoMessage() {}

func (x *DisconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_controller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectRequest.ProtoReflect.Descriptor instead.
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_controller_proto_rawDescGZIP(), []int{3}
}

func (x *DisconnectRequest) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *DisconnectRequest) GetConnectClientAddress() string {
	if x != nil {
		return x.ConnectClientAddress
	}
	return ""
}

// Disconnect Reply from server
type DisconnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DisconnectReply) Reset() {
	*x = DisconnectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_controller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectReply) ProtoMessage() {}

func (x *DisconnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_controller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectReply.ProtoReflect.Descriptor instead.
func (*DisconnectReply) Descriptor() ([]byte, []int) {
	return file_pkg_proto_controller_proto_rawDescGZIP(), []int{4}
}

func (x *DisconnectReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DisconnectReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// GetRecordFile Request from client
type GetRecordFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID  string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GetRecordFileRequest) Reset() {
	*x = GetRecordFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_controller_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordFileRequest) ProtoMessage() {}

func (x *GetRecordFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_controller_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordFileRequest.ProtoReflect.Descriptor instead.
func (*GetRecordFileRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_controller_proto_rawDescGZIP(), []int{5}
}

func (x *GetRecordFileRequest) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *GetRecordFileRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// GetRecordFile Reply from server
type GetRecordFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message     string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FileAddress string `protobuf:"bytes,3,opt,name=fileAddress,proto3" json:"fileAddress,omitempty"`
	StartTime   int64  `protobuf:"varint,4,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime     int64  `protobuf:"varint,5,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *GetRecordFileReply) Reset() {
	*x = GetRecordFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_controller_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordFileReply) ProtoMessage() {}

func (x *GetRecordFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_controller_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordFileReply.ProtoReflect.Descriptor instead.
func (*GetRecordFileReply) Descriptor() ([]byte, []int) {
	return file_pkg_proto_controller_proto_rawDescGZIP(), []int{6}
}

func (x *GetRecordFileReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetRecordFileReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetRecordFileReply) GetFileAddress() string {
	if x != nil {
		return x.FileAddress
	}
	return ""
}

func (x *GetRecordFileReply) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *GetRecordFileReply) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

var File_pkg_proto_controller_proto protoreflect.FileDescriptor

var file_pkg_proto_controller_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a,
	0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x14, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x54, 0x53, 0x50, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x54, 0x53, 0x50, 0x52, 0x65, 0x6c, 0x61,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x36, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x63, 0x0a,
	0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x63, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3f, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x9c, 0x01, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xae, 0x01, 0x0a, 0x0a, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x0f, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x68, 0x75, 0x79, 0x6e,
	0x68, 0x31, 0x31, 0x30, 0x38, 0x2f, 0x68, 0x63, 0x6d, 0x75, 0x74, 0x2d, 0x74, 0x68, 0x65, 0x73,
	0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_controller_proto_rawDescOnce sync.Once
	file_pkg_proto_controller_proto_rawDescData = file_pkg_proto_controller_proto_rawDesc
)

func file_pkg_proto_controller_proto_rawDescGZIP() []byte {
	file_pkg_proto_controller_proto_rawDescOnce.Do(func() {
		file_pkg_proto_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_controller_proto_rawDescData)
	})
	return file_pkg_proto_controller_proto_rawDescData
}

var file_pkg_proto_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_proto_controller_proto_goTypes = []interface{}{
	(*ConnectRequest)(nil),       // 0: ConnectRequest
	(*ConnectReplyData)(nil),     // 1: ConnectReplyData
	(*ConnectReply)(nil),         // 2: ConnectReply
	(*DisconnectRequest)(nil),    // 3: DisconnectRequest
	(*DisconnectReply)(nil),      // 4: DisconnectReply
	(*GetRecordFileRequest)(nil), // 5: GetRecordFileRequest
	(*GetRecordFileReply)(nil),   // 6: GetRecordFileReply
}
var file_pkg_proto_controller_proto_depIdxs = []int32{
	1, // 0: ConnectReply.data:type_name -> ConnectReplyData
	0, // 1: Controller.Connect:input_type -> ConnectRequest
	3, // 2: Controller.Disconnect:input_type -> DisconnectRequest
	5, // 3: Controller.GetRecordFile:input_type -> GetRecordFileRequest
	2, // 4: Controller.Connect:output_type -> ConnectReply
	4, // 5: Controller.Disconnect:output_type -> DisconnectReply
	6, // 6: Controller.GetRecordFile:output_type -> GetRecordFileReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_controller_proto_init() }
func file_pkg_proto_controller_proto_init() {
	if File_pkg_proto_controller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_pkg_proto_controller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectReplyData); i {
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
		file_pkg_proto_controller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectReply); i {
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
		file_pkg_proto_controller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectRequest); i {
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
		file_pkg_proto_controller_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectReply); i {
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
		file_pkg_proto_controller_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordFileRequest); i {
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
		file_pkg_proto_controller_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordFileReply); i {
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
			RawDescriptor: file_pkg_proto_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_controller_proto_goTypes,
		DependencyIndexes: file_pkg_proto_controller_proto_depIdxs,
		MessageInfos:      file_pkg_proto_controller_proto_msgTypes,
	}.Build()
	File_pkg_proto_controller_proto = out.File
	file_pkg_proto_controller_proto_rawDesc = nil
	file_pkg_proto_controller_proto_goTypes = nil
	file_pkg_proto_controller_proto_depIdxs = nil
}