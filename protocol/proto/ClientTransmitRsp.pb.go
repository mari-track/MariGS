// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: ClientTransmitRsp.proto

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

type ClientTransmitRsp_CmdId int32

const (
	ClientTransmitRsp_NONE             ClientTransmitRsp_CmdId = 0
	ClientTransmitRsp_CMD_ID           ClientTransmitRsp_CmdId = 241
	ClientTransmitRsp_ENET_CHANNEL_ID  ClientTransmitRsp_CmdId = 0
	ClientTransmitRsp_ENET_IS_RELIABLE ClientTransmitRsp_CmdId = 1
)

// Enum value maps for ClientTransmitRsp_CmdId.
var (
	ClientTransmitRsp_CmdId_name = map[int32]string{
		0:   "NONE",
		241: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	ClientTransmitRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           241,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x ClientTransmitRsp_CmdId) Enum() *ClientTransmitRsp_CmdId {
	p := new(ClientTransmitRsp_CmdId)
	*p = x
	return p
}

func (x ClientTransmitRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientTransmitRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ClientTransmitRsp_proto_enumTypes[0].Descriptor()
}

func (ClientTransmitRsp_CmdId) Type() protoreflect.EnumType {
	return &file_ClientTransmitRsp_proto_enumTypes[0]
}

func (x ClientTransmitRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientTransmitRsp_CmdId.Descriptor instead.
func (ClientTransmitRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ClientTransmitRsp_proto_rawDescGZIP(), []int{0, 0}
}

type ClientTransmitRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32          `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Reason  TransmitReason `protobuf:"varint,2,opt,name=reason,proto3,enum=proto.TransmitReason" json:"reason,omitempty"`
}

func (x *ClientTransmitRsp) Reset() {
	*x = ClientTransmitRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientTransmitRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientTransmitRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientTransmitRsp) ProtoMessage() {}

func (x *ClientTransmitRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ClientTransmitRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientTransmitRsp.ProtoReflect.Descriptor instead.
func (*ClientTransmitRsp) Descriptor() ([]byte, []int) {
	return file_ClientTransmitRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ClientTransmitRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ClientTransmitRsp) GetReason() TransmitReason {
	if x != nil {
		return x.Reason
	}
	return TransmitReason_TRANSMIT_NONE
}

var File_ClientTransmitRsp_proto protoreflect.FileDescriptor

var file_ClientTransmitRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f,
	0x49, 0x44, 0x10, 0xf1, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01,
	0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ClientTransmitRsp_proto_rawDescOnce sync.Once
	file_ClientTransmitRsp_proto_rawDescData = file_ClientTransmitRsp_proto_rawDesc
)

func file_ClientTransmitRsp_proto_rawDescGZIP() []byte {
	file_ClientTransmitRsp_proto_rawDescOnce.Do(func() {
		file_ClientTransmitRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClientTransmitRsp_proto_rawDescData)
	})
	return file_ClientTransmitRsp_proto_rawDescData
}

var file_ClientTransmitRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ClientTransmitRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ClientTransmitRsp_proto_goTypes = []interface{}{
	(ClientTransmitRsp_CmdId)(0), // 0: proto.ClientTransmitRsp.CmdId
	(*ClientTransmitRsp)(nil),    // 1: proto.ClientTransmitRsp
	(TransmitReason)(0),          // 2: proto.TransmitReason
}
var file_ClientTransmitRsp_proto_depIdxs = []int32{
	2, // 0: proto.ClientTransmitRsp.reason:type_name -> proto.TransmitReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ClientTransmitRsp_proto_init() }
func file_ClientTransmitRsp_proto_init() {
	if File_ClientTransmitRsp_proto != nil {
		return
	}
	file_TransmitReason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ClientTransmitRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientTransmitRsp); i {
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
			RawDescriptor: file_ClientTransmitRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ClientTransmitRsp_proto_goTypes,
		DependencyIndexes: file_ClientTransmitRsp_proto_depIdxs,
		EnumInfos:         file_ClientTransmitRsp_proto_enumTypes,
		MessageInfos:      file_ClientTransmitRsp_proto_msgTypes,
	}.Build()
	File_ClientTransmitRsp_proto = out.File
	file_ClientTransmitRsp_proto_rawDesc = nil
	file_ClientTransmitRsp_proto_goTypes = nil
	file_ClientTransmitRsp_proto_depIdxs = nil
}
