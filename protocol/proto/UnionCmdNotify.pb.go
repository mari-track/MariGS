// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: UnionCmdNotify.proto

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

type UnionCmdNotify_CmdId int32

const (
	UnionCmdNotify_NONE             UnionCmdNotify_CmdId = 0
	UnionCmdNotify_CMD_ID           UnionCmdNotify_CmdId = 16
	UnionCmdNotify_ENET_CHANNEL_ID  UnionCmdNotify_CmdId = 0
	UnionCmdNotify_ENET_IS_RELIABLE UnionCmdNotify_CmdId = 1
	UnionCmdNotify_IS_ALLOW_CLIENT  UnionCmdNotify_CmdId = 1
)

// Enum value maps for UnionCmdNotify_CmdId.
var (
	UnionCmdNotify_CmdId_name = map[int32]string{
		0:  "NONE",
		16: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	UnionCmdNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           16,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x UnionCmdNotify_CmdId) Enum() *UnionCmdNotify_CmdId {
	p := new(UnionCmdNotify_CmdId)
	*p = x
	return p
}

func (x UnionCmdNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UnionCmdNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_UnionCmdNotify_proto_enumTypes[0].Descriptor()
}

func (UnionCmdNotify_CmdId) Type() protoreflect.EnumType {
	return &file_UnionCmdNotify_proto_enumTypes[0]
}

func (x UnionCmdNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UnionCmdNotify_CmdId.Descriptor instead.
func (UnionCmdNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_UnionCmdNotify_proto_rawDescGZIP(), []int{0, 0}
}

type UnionCmdNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CmdList []*UnionCmd `protobuf:"bytes,1,rep,name=cmd_list,json=cmdList,proto3" json:"cmd_list,omitempty"`
}

func (x *UnionCmdNotify) Reset() {
	*x = UnionCmdNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UnionCmdNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnionCmdNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnionCmdNotify) ProtoMessage() {}

func (x *UnionCmdNotify) ProtoReflect() protoreflect.Message {
	mi := &file_UnionCmdNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnionCmdNotify.ProtoReflect.Descriptor instead.
func (*UnionCmdNotify) Descriptor() ([]byte, []int) {
	return file_UnionCmdNotify_proto_rawDescGZIP(), []int{0}
}

func (x *UnionCmdNotify) GetCmdList() []*UnionCmd {
	if x != nil {
		return x.CmdList
	}
	return nil
}

var File_UnionCmdNotify_proto protoreflect.FileDescriptor

var file_UnionCmdNotify_proto_rawDesc = []byte{
	0x0a, 0x14, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x43, 0x6d, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x43, 0x6d, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x0e, 0x55, 0x6e, 0x69, 0x6f, 0x6e,
	0x43, 0x6d, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x6d, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x55, 0x6e,
	0x69, 0x6f, 0x6e, 0x43, 0x6d, 0x64, 0x52, 0x07, 0x63, 0x6d, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x61, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x10, 0x12, 0x13,
	0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52,
	0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f,
	0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02,
	0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UnionCmdNotify_proto_rawDescOnce sync.Once
	file_UnionCmdNotify_proto_rawDescData = file_UnionCmdNotify_proto_rawDesc
)

func file_UnionCmdNotify_proto_rawDescGZIP() []byte {
	file_UnionCmdNotify_proto_rawDescOnce.Do(func() {
		file_UnionCmdNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_UnionCmdNotify_proto_rawDescData)
	})
	return file_UnionCmdNotify_proto_rawDescData
}

var file_UnionCmdNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_UnionCmdNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_UnionCmdNotify_proto_goTypes = []interface{}{
	(UnionCmdNotify_CmdId)(0), // 0: UnionCmdNotify.CmdId
	(*UnionCmdNotify)(nil),    // 1: UnionCmdNotify
	(*UnionCmd)(nil),          // 2: UnionCmd
}
var file_UnionCmdNotify_proto_depIdxs = []int32{
	2, // 0: UnionCmdNotify.cmd_list:type_name -> UnionCmd
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_UnionCmdNotify_proto_init() }
func file_UnionCmdNotify_proto_init() {
	if File_UnionCmdNotify_proto != nil {
		return
	}
	file_UnionCmd_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_UnionCmdNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnionCmdNotify); i {
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
			RawDescriptor: file_UnionCmdNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UnionCmdNotify_proto_goTypes,
		DependencyIndexes: file_UnionCmdNotify_proto_depIdxs,
		EnumInfos:         file_UnionCmdNotify_proto_enumTypes,
		MessageInfos:      file_UnionCmdNotify_proto_msgTypes,
	}.Build()
	File_UnionCmdNotify_proto = out.File
	file_UnionCmdNotify_proto_rawDesc = nil
	file_UnionCmdNotify_proto_goTypes = nil
	file_UnionCmdNotify_proto_depIdxs = nil
}
