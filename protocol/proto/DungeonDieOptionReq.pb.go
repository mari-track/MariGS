// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: DungeonDieOptionReq.proto

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

type DungeonDieOptionReq_CmdId int32

const (
	DungeonDieOptionReq_NONE             DungeonDieOptionReq_CmdId = 0
	DungeonDieOptionReq_CMD_ID           DungeonDieOptionReq_CmdId = 912
	DungeonDieOptionReq_ENET_CHANNEL_ID  DungeonDieOptionReq_CmdId = 0
	DungeonDieOptionReq_ENET_IS_RELIABLE DungeonDieOptionReq_CmdId = 1
	DungeonDieOptionReq_IS_ALLOW_CLIENT  DungeonDieOptionReq_CmdId = 1
)

// Enum value maps for DungeonDieOptionReq_CmdId.
var (
	DungeonDieOptionReq_CmdId_name = map[int32]string{
		0:   "NONE",
		912: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	DungeonDieOptionReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           912,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x DungeonDieOptionReq_CmdId) Enum() *DungeonDieOptionReq_CmdId {
	p := new(DungeonDieOptionReq_CmdId)
	*p = x
	return p
}

func (x DungeonDieOptionReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DungeonDieOptionReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_DungeonDieOptionReq_proto_enumTypes[0].Descriptor()
}

func (DungeonDieOptionReq_CmdId) Type() protoreflect.EnumType {
	return &file_DungeonDieOptionReq_proto_enumTypes[0]
}

func (x DungeonDieOptionReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DungeonDieOptionReq_CmdId.Descriptor instead.
func (DungeonDieOptionReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_DungeonDieOptionReq_proto_rawDescGZIP(), []int{0, 0}
}

type DungeonDieOptionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DieOption PlayerDieOption `protobuf:"varint,1,opt,name=die_option,json=dieOption,proto3,enum=proto.PlayerDieOption" json:"die_option,omitempty"`
}

func (x *DungeonDieOptionReq) Reset() {
	*x = DungeonDieOptionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DungeonDieOptionReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DungeonDieOptionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DungeonDieOptionReq) ProtoMessage() {}

func (x *DungeonDieOptionReq) ProtoReflect() protoreflect.Message {
	mi := &file_DungeonDieOptionReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DungeonDieOptionReq.ProtoReflect.Descriptor instead.
func (*DungeonDieOptionReq) Descriptor() ([]byte, []int) {
	return file_DungeonDieOptionReq_proto_rawDescGZIP(), []int{0}
}

func (x *DungeonDieOptionReq) GetDieOption() PlayerDieOption {
	if x != nil {
		return x.DieOption
	}
	return PlayerDieOption_DIE_OPT_NONE
}

var File_DungeonDieOptionReq_proto protoreflect.FileDescriptor

var file_DungeonDieOptionReq_proto_rawDesc = []byte{
	0x0a, 0x19, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x44, 0x69, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x13, 0x44, 0x75,
	0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x44, 0x69, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x35, 0x0a, 0x0a, 0x64, 0x69, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64,
	0x69, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49,
	0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43,
	0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x90, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f,
	0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DungeonDieOptionReq_proto_rawDescOnce sync.Once
	file_DungeonDieOptionReq_proto_rawDescData = file_DungeonDieOptionReq_proto_rawDesc
)

func file_DungeonDieOptionReq_proto_rawDescGZIP() []byte {
	file_DungeonDieOptionReq_proto_rawDescOnce.Do(func() {
		file_DungeonDieOptionReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_DungeonDieOptionReq_proto_rawDescData)
	})
	return file_DungeonDieOptionReq_proto_rawDescData
}

var file_DungeonDieOptionReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DungeonDieOptionReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DungeonDieOptionReq_proto_goTypes = []interface{}{
	(DungeonDieOptionReq_CmdId)(0), // 0: proto.DungeonDieOptionReq.CmdId
	(*DungeonDieOptionReq)(nil),    // 1: proto.DungeonDieOptionReq
	(PlayerDieOption)(0),           // 2: proto.PlayerDieOption
}
var file_DungeonDieOptionReq_proto_depIdxs = []int32{
	2, // 0: proto.DungeonDieOptionReq.die_option:type_name -> proto.PlayerDieOption
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DungeonDieOptionReq_proto_init() }
func file_DungeonDieOptionReq_proto_init() {
	if File_DungeonDieOptionReq_proto != nil {
		return
	}
	file_PlayerDieOption_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DungeonDieOptionReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DungeonDieOptionReq); i {
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
			RawDescriptor: file_DungeonDieOptionReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DungeonDieOptionReq_proto_goTypes,
		DependencyIndexes: file_DungeonDieOptionReq_proto_depIdxs,
		EnumInfos:         file_DungeonDieOptionReq_proto_enumTypes,
		MessageInfos:      file_DungeonDieOptionReq_proto_msgTypes,
	}.Build()
	File_DungeonDieOptionReq_proto = out.File
	file_DungeonDieOptionReq_proto_rawDesc = nil
	file_DungeonDieOptionReq_proto_goTypes = nil
	file_DungeonDieOptionReq_proto_depIdxs = nil
}
