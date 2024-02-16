// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: TowerBuffSelectRsp.proto

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

type TowerBuffSelectRsp_CmdId int32

const (
	TowerBuffSelectRsp_NONE             TowerBuffSelectRsp_CmdId = 0
	TowerBuffSelectRsp_CMD_ID           TowerBuffSelectRsp_CmdId = 2414
	TowerBuffSelectRsp_ENET_CHANNEL_ID  TowerBuffSelectRsp_CmdId = 0
	TowerBuffSelectRsp_ENET_IS_RELIABLE TowerBuffSelectRsp_CmdId = 1
)

// Enum value maps for TowerBuffSelectRsp_CmdId.
var (
	TowerBuffSelectRsp_CmdId_name = map[int32]string{
		0:    "NONE",
		2414: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	TowerBuffSelectRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2414,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x TowerBuffSelectRsp_CmdId) Enum() *TowerBuffSelectRsp_CmdId {
	p := new(TowerBuffSelectRsp_CmdId)
	*p = x
	return p
}

func (x TowerBuffSelectRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TowerBuffSelectRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_TowerBuffSelectRsp_proto_enumTypes[0].Descriptor()
}

func (TowerBuffSelectRsp_CmdId) Type() protoreflect.EnumType {
	return &file_TowerBuffSelectRsp_proto_enumTypes[0]
}

func (x TowerBuffSelectRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TowerBuffSelectRsp_CmdId.Descriptor instead.
func (TowerBuffSelectRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_TowerBuffSelectRsp_proto_rawDescGZIP(), []int{0, 0}
}

type TowerBuffSelectRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	TowerBuffId uint32 `protobuf:"varint,2,opt,name=tower_buff_id,json=towerBuffId,proto3" json:"tower_buff_id,omitempty"`
}

func (x *TowerBuffSelectRsp) Reset() {
	*x = TowerBuffSelectRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TowerBuffSelectRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TowerBuffSelectRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TowerBuffSelectRsp) ProtoMessage() {}

func (x *TowerBuffSelectRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TowerBuffSelectRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TowerBuffSelectRsp.ProtoReflect.Descriptor instead.
func (*TowerBuffSelectRsp) Descriptor() ([]byte, []int) {
	return file_TowerBuffSelectRsp_proto_rawDescGZIP(), []int{0}
}

func (x *TowerBuffSelectRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *TowerBuffSelectRsp) GetTowerBuffId() uint32 {
	if x != nil {
		return x.TowerBuffId
	}
	return 0
}

var File_TowerBuffSelectRsp_proto protoreflect.FileDescriptor

var file_TowerBuffSelectRsp_proto_rawDesc = []byte{
	0x0a, 0x18, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x12, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x75, 0x66, 0x66, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x62, 0x75, 0x66, 0x66,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x6f, 0x77, 0x65, 0x72,
	0x42, 0x75, 0x66, 0x66, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44,
	0x5f, 0x49, 0x44, 0x10, 0xee, 0x12, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TowerBuffSelectRsp_proto_rawDescOnce sync.Once
	file_TowerBuffSelectRsp_proto_rawDescData = file_TowerBuffSelectRsp_proto_rawDesc
)

func file_TowerBuffSelectRsp_proto_rawDescGZIP() []byte {
	file_TowerBuffSelectRsp_proto_rawDescOnce.Do(func() {
		file_TowerBuffSelectRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_TowerBuffSelectRsp_proto_rawDescData)
	})
	return file_TowerBuffSelectRsp_proto_rawDescData
}

var file_TowerBuffSelectRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_TowerBuffSelectRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TowerBuffSelectRsp_proto_goTypes = []interface{}{
	(TowerBuffSelectRsp_CmdId)(0), // 0: proto.TowerBuffSelectRsp.CmdId
	(*TowerBuffSelectRsp)(nil),    // 1: proto.TowerBuffSelectRsp
}
var file_TowerBuffSelectRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TowerBuffSelectRsp_proto_init() }
func file_TowerBuffSelectRsp_proto_init() {
	if File_TowerBuffSelectRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TowerBuffSelectRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TowerBuffSelectRsp); i {
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
			RawDescriptor: file_TowerBuffSelectRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TowerBuffSelectRsp_proto_goTypes,
		DependencyIndexes: file_TowerBuffSelectRsp_proto_depIdxs,
		EnumInfos:         file_TowerBuffSelectRsp_proto_enumTypes,
		MessageInfos:      file_TowerBuffSelectRsp_proto_msgTypes,
	}.Build()
	File_TowerBuffSelectRsp_proto = out.File
	file_TowerBuffSelectRsp_proto_rawDesc = nil
	file_TowerBuffSelectRsp_proto_goTypes = nil
	file_TowerBuffSelectRsp_proto_depIdxs = nil
}