// This file holds the proto message formats for interacting with the core API endpoints

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: mcs.proto

package mcs

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

type CTP_Flag int32

const (
	CTP_SESSION_INIT CTP_Flag = 0 // used by client to propose session ID
	CTP_ACK          CTP_Flag = 1 // generic acknowledge, multiple uses
	CTP_DATA         CTP_Flag = 2 // data transfer
	CTP_FIN          CTP_Flag = 3 // transfer is complete
	CTP_RETRANS      CTP_Flag = 4 // request retransmission of prev packet
)

// Enum value maps for CTP_Flag.
var (
	CTP_Flag_name = map[int32]string{
		0: "SESSION_INIT",
		1: "ACK",
		2: "DATA",
		3: "FIN",
		4: "RETRANS",
	}
	CTP_Flag_value = map[string]int32{
		"SESSION_INIT": 0,
		"ACK":          1,
		"DATA":         2,
		"FIN":          3,
		"RETRANS":      4,
	}
)

func (x CTP_Flag) Enum() *CTP_Flag {
	p := new(CTP_Flag)
	*p = x
	return p
}

func (x CTP_Flag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CTP_Flag) Descriptor() protoreflect.EnumDescriptor {
	return file_mcs_proto_enumTypes[0].Descriptor()
}

func (CTP_Flag) Type() protoreflect.EnumType {
	return &file_mcs_proto_enumTypes[0]
}

func (x CTP_Flag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CTP_Flag.Descriptor instead.
func (CTP_Flag) EnumDescriptor() ([]byte, []int) {
	return file_mcs_proto_rawDescGZIP(), []int{3, 0}
}

// Meteor Communication Standard
type MCS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    int32     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`      // HTTP Status Code
	Uuid      string    `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`           // UUID for Bot or Action UUI
	Hostname  string    `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`   // hostname
	Groupname string    `protobuf:"bytes,4,opt,name=groupname,proto3" json:"groupname,omitempty"` // group name
	Mode      string    `protobuf:"bytes,5,opt,name=mode,proto3" json:"mode,omitempty"`           // the action mode (shell, scrshot, etc)
	Interface string    `protobuf:"bytes,6,opt,name=interface,proto3" json:"interface,omitempty"` // the network interface of the host
	Desc      string    `protobuf:"bytes,7,opt,name=desc,proto3" json:"desc,omitempty"`           // description for group or API response
	Args      string    `protobuf:"bytes,8,opt,name=args,proto3" json:"args,omitempty"`           // the data of the action (command line args, bot uuid, etc)
	Result    string    `protobuf:"bytes,9,opt,name=result,proto3" json:"result,omitempty"`       // the result of an action (usually command output)
	Interval  int32     `protobuf:"varint,10,opt,name=interval,proto3" json:"interval,omitempty"` // callback interval
	Delta     int32     `protobuf:"varint,11,opt,name=delta,proto3" json:"delta,omitempty"`       // jitter for interval
	Actions   []*Action `protobuf:"bytes,12,rep,name=actions,proto3" json:"actions,omitempty"`    // array of Actions
	AuthDat   *AuthDat  `protobuf:"bytes,13,opt,name=authDat,proto3" json:"authDat,omitempty"`    // User authentication data
}

func (x *MCS) Reset() {
	*x = MCS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MCS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MCS) ProtoMessage() {}

func (x *MCS) ProtoReflect() protoreflect.Message {
	mi := &file_mcs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MCS.ProtoReflect.Descriptor instead.
func (*MCS) Descriptor() ([]byte, []int) {
	return file_mcs_proto_rawDescGZIP(), []int{0}
}

func (x *MCS) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MCS) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *MCS) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *MCS) GetGroupname() string {
	if x != nil {
		return x.Groupname
	}
	return ""
}

func (x *MCS) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *MCS) GetInterface() string {
	if x != nil {
		return x.Interface
	}
	return ""
}

func (x *MCS) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *MCS) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

func (x *MCS) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *MCS) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *MCS) GetDelta() int32 {
	if x != nil {
		return x.Delta
	}
	return 0
}

func (x *MCS) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *MCS) GetAuthDat() *AuthDat {
	if x != nil {
		return x.AuthDat
	}
	return nil
}

// Actions to be run by the bot
type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // the UUID of the action
	Mode string `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"` // type of data that will be in "args" or is being requested
	Args string `protobuf:"bytes,3,opt,name=args,proto3" json:"args,omitempty"` // the data of the action (command line args, bot uuid, etc)
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_mcs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_mcs_proto_rawDescGZIP(), []int{1}
}

func (x *Action) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Action) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *Action) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

// User authentication data
type AuthDat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"` // username
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"` // password
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`       // login session token
}

func (x *AuthDat) Reset() {
	*x = AuthDat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthDat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthDat) ProtoMessage() {}

func (x *AuthDat) ProtoReflect() protoreflect.Message {
	mi := &file_mcs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthDat.ProtoReflect.Descriptor instead.
func (*AuthDat) Descriptor() ([]byte, []int) {
	return file_mcs_proto_rawDescGZIP(), []int{2}
}

func (x *AuthDat) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthDat) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AuthDat) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

//Proto for Cera Transfer Protocol
type CTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId int32    `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`            // unique int will be generated for each session
	Payload   []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`                                  // the actual data being sent will go here
	Checksum  []byte   `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty"`                                // last 8 digits of payload SHA1 will go here
	TypeFlag  CTP_Flag `protobuf:"varint,4,opt,name=type_flag,json=typeFlag,proto3,enum=CTP_Flag" json:"type_flag,omitempty"` // flag to say what type of data is in the payload
}

func (x *CTP) Reset() {
	*x = CTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CTP) ProtoMessage() {}

func (x *CTP) ProtoReflect() protoreflect.Message {
	mi := &file_mcs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CTP.ProtoReflect.Descriptor instead.
func (*CTP) Descriptor() ([]byte, []int) {
	return file_mcs_proto_rawDescGZIP(), []int{3}
}

func (x *CTP) GetSessionId() int32 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *CTP) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *CTP) GetChecksum() []byte {
	if x != nil {
		return x.Checksum
	}
	return nil
}

func (x *CTP) GetTypeFlag() CTP_Flag {
	if x != nil {
		return x.TypeFlag
	}
	return CTP_SESSION_INIT
}

var File_mcs_proto protoreflect.FileDescriptor

var file_mcs_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x02, 0x0a, 0x03,
	0x4d, 0x43, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x21, 0x0a,
	0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x22, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x52, 0x07, 0x61, 0x75, 0x74,
	0x68, 0x44, 0x61, 0x74, 0x22, 0x44, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x57, 0x0a, 0x07, 0x41, 0x75,
	0x74, 0x68, 0x44, 0x61, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xc5, 0x01, 0x0a, 0x03, 0x43, 0x54, 0x50, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d,
	0x12, 0x26, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x43, 0x54, 0x50, 0x2e, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x08,
	0x74, 0x79, 0x70, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x22, 0x41, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x49, 0x54,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44,
	0x41, 0x54, 0x41, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x0b,
	0x0a, 0x07, 0x52, 0x45, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x10, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mcs_proto_rawDescOnce sync.Once
	file_mcs_proto_rawDescData = file_mcs_proto_rawDesc
)

func file_mcs_proto_rawDescGZIP() []byte {
	file_mcs_proto_rawDescOnce.Do(func() {
		file_mcs_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcs_proto_rawDescData)
	})
	return file_mcs_proto_rawDescData
}

var file_mcs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mcs_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mcs_proto_goTypes = []interface{}{
	(CTP_Flag)(0),   // 0: CTP.Flag
	(*MCS)(nil),     // 1: MCS
	(*Action)(nil),  // 2: Action
	(*AuthDat)(nil), // 3: AuthDat
	(*CTP)(nil),     // 4: CTP
}
var file_mcs_proto_depIdxs = []int32{
	2, // 0: MCS.actions:type_name -> Action
	3, // 1: MCS.authDat:type_name -> AuthDat
	0, // 2: CTP.type_flag:type_name -> CTP.Flag
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mcs_proto_init() }
func file_mcs_proto_init() {
	if File_mcs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mcs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MCS); i {
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
		file_mcs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_mcs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthDat); i {
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
		file_mcs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CTP); i {
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
			RawDescriptor: file_mcs_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcs_proto_goTypes,
		DependencyIndexes: file_mcs_proto_depIdxs,
		EnumInfos:         file_mcs_proto_enumTypes,
		MessageInfos:      file_mcs_proto_msgTypes,
	}.Build()
	File_mcs_proto = out.File
	file_mcs_proto_rawDesc = nil
	file_mcs_proto_goTypes = nil
	file_mcs_proto_depIdxs = nil
}
