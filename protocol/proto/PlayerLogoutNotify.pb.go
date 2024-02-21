// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: PlayerLogoutNotify.proto

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

type PlayerLogoutNotify_CmdId int32

const (
	PlayerLogoutNotify_NONE   PlayerLogoutNotify_CmdId = 0
	PlayerLogoutNotify_CMD_ID PlayerLogoutNotify_CmdId = 107
)

// Enum value maps for PlayerLogoutNotify_CmdId.
var (
	PlayerLogoutNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		107: "CMD_ID",
	}
	PlayerLogoutNotify_CmdId_value = map[string]int32{
		"NONE":   0,
		"CMD_ID": 107,
	}
)

func (x PlayerLogoutNotify_CmdId) Enum() *PlayerLogoutNotify_CmdId {
	p := new(PlayerLogoutNotify_CmdId)
	*p = x
	return p
}

func (x PlayerLogoutNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerLogoutNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerLogoutNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerLogoutNotify_CmdId) Type() protoreflect.EnumType {
	return &file_PlayerLogoutNotify_proto_enumTypes[0]
}

func (x PlayerLogoutNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerLogoutNotify_CmdId.Descriptor instead.
func (PlayerLogoutNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PlayerLogoutNotify_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerLogoutNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32 `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *PlayerLogoutNotify) Reset() {
	*x = PlayerLogoutNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerLogoutNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLogoutNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLogoutNotify) ProtoMessage() {}

func (x *PlayerLogoutNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerLogoutNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLogoutNotify.ProtoReflect.Descriptor instead.
func (*PlayerLogoutNotify) Descriptor() ([]byte, []int) {
	return file_PlayerLogoutNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerLogoutNotify) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_PlayerLogoutNotify_proto protoreflect.FileDescriptor

var file_PlayerLogoutNotify_proto_rawDesc = []byte{
	0x0a, 0x18, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x12, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x43, 0x6d,
	0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x6b, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerLogoutNotify_proto_rawDescOnce sync.Once
	file_PlayerLogoutNotify_proto_rawDescData = file_PlayerLogoutNotify_proto_rawDesc
)

func file_PlayerLogoutNotify_proto_rawDescGZIP() []byte {
	file_PlayerLogoutNotify_proto_rawDescOnce.Do(func() {
		file_PlayerLogoutNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerLogoutNotify_proto_rawDescData)
	})
	return file_PlayerLogoutNotify_proto_rawDescData
}

var file_PlayerLogoutNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerLogoutNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerLogoutNotify_proto_goTypes = []interface{}{
	(PlayerLogoutNotify_CmdId)(0), // 0: PlayerLogoutNotify.CmdId
	(*PlayerLogoutNotify)(nil),    // 1: PlayerLogoutNotify
}
var file_PlayerLogoutNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlayerLogoutNotify_proto_init() }
func file_PlayerLogoutNotify_proto_init() {
	if File_PlayerLogoutNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlayerLogoutNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLogoutNotify); i {
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
			RawDescriptor: file_PlayerLogoutNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerLogoutNotify_proto_goTypes,
		DependencyIndexes: file_PlayerLogoutNotify_proto_depIdxs,
		EnumInfos:         file_PlayerLogoutNotify_proto_enumTypes,
		MessageInfos:      file_PlayerLogoutNotify_proto_msgTypes,
	}.Build()
	File_PlayerLogoutNotify_proto = out.File
	file_PlayerLogoutNotify_proto_rawDesc = nil
	file_PlayerLogoutNotify_proto_goTypes = nil
	file_PlayerLogoutNotify_proto_depIdxs = nil
}
