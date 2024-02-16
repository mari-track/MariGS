// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: PlayerInjectFixNotify.proto

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

type PlayerInjectFixNotify_CmdId int32

const (
	PlayerInjectFixNotify_NONE             PlayerInjectFixNotify_CmdId = 0
	PlayerInjectFixNotify_CMD_ID           PlayerInjectFixNotify_CmdId = 166
	PlayerInjectFixNotify_ENET_CHANNEL_ID  PlayerInjectFixNotify_CmdId = 0
	PlayerInjectFixNotify_ENET_IS_RELIABLE PlayerInjectFixNotify_CmdId = 1
)

// Enum value maps for PlayerInjectFixNotify_CmdId.
var (
	PlayerInjectFixNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		166: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	PlayerInjectFixNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           166,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x PlayerInjectFixNotify_CmdId) Enum() *PlayerInjectFixNotify_CmdId {
	p := new(PlayerInjectFixNotify_CmdId)
	*p = x
	return p
}

func (x PlayerInjectFixNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerInjectFixNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerInjectFixNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerInjectFixNotify_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerInjectFixNotify_proto_enumTypes[0]
}

func (x PlayerInjectFixNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerInjectFixNotify_CmdId.Descriptor instead.
func (PlayerInjectFixNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerInjectFixNotify_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerInjectFixNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	InjectFix []byte `protobuf:"bytes,2,opt,name=inject_fix,json=injectFix,proto3" json:"inject_fix,omitempty"`
}

func (x *PlayerInjectFixNotify) Reset() {
	*x = PlayerInjectFixNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerInjectFixNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInjectFixNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInjectFixNotify) ProtoMessage() {}

func (x *PlayerInjectFixNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerInjectFixNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInjectFixNotify.ProtoReflect.Descriptor instead.
func (*PlayerInjectFixNotify) Descriptor() ([]byte, []int) {
	return file_PlayerInjectFixNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerInjectFixNotify) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PlayerInjectFixNotify) GetInjectFix() []byte {
	if x != nil {
		return x.InjectFix
	}
	return nil
}

var File_PlayerInjectFixNotify_proto protoreflect.FileDescriptor

var file_PlayerInjectFixNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x69,
	0x78, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x15, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x69, 0x78, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x69, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x69, 0x78, 0x22, 0x4d, 0x0a,
	0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xa6, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45,
	0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerInjectFixNotify_proto_rawDescOnce sync.Once
	file_PlayerInjectFixNotify_proto_rawDescData = file_PlayerInjectFixNotify_proto_rawDesc
)

func file_PlayerInjectFixNotify_proto_rawDescGZIP() []byte {
	file_PlayerInjectFixNotify_proto_rawDescOnce.Do(func() {
		file_PlayerInjectFixNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerInjectFixNotify_proto_rawDescData)
	})
	return file_PlayerInjectFixNotify_proto_rawDescData
}

var file_PlayerInjectFixNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerInjectFixNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerInjectFixNotify_proto_goTypes = []interface{}{
	(PlayerInjectFixNotify_CmdId)(0), // 0: proto.PlayerInjectFixNotify.CmdId
	(*PlayerInjectFixNotify)(nil),    // 1: proto.PlayerInjectFixNotify
}
var file_PlayerInjectFixNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlayerInjectFixNotify_proto_init() }
func file_PlayerInjectFixNotify_proto_init() {
	if File_PlayerInjectFixNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlayerInjectFixNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInjectFixNotify); i {
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
			RawDescriptor: file_PlayerInjectFixNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerInjectFixNotify_proto_goTypes,
		DependencyIndexes: file_PlayerInjectFixNotify_proto_depIdxs,
		EnumInfos:         file_PlayerInjectFixNotify_proto_enumTypes,
		MessageInfos:      file_PlayerInjectFixNotify_proto_msgTypes,
	}.Build()
	File_PlayerInjectFixNotify_proto = out.File
	file_PlayerInjectFixNotify_proto_rawDesc = nil
	file_PlayerInjectFixNotify_proto_goTypes = nil
	file_PlayerInjectFixNotify_proto_depIdxs = nil
}
