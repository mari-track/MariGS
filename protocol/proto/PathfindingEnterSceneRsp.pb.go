// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: PathfindingEnterSceneRsp.proto

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

type PathfindingEnterSceneRsp_CmdId int32

const (
	PathfindingEnterSceneRsp_NONE             PathfindingEnterSceneRsp_CmdId = 0
	PathfindingEnterSceneRsp_CMD_ID           PathfindingEnterSceneRsp_CmdId = 2306
	PathfindingEnterSceneRsp_ENET_CHANNEL_ID  PathfindingEnterSceneRsp_CmdId = 0
	PathfindingEnterSceneRsp_ENET_IS_RELIABLE PathfindingEnterSceneRsp_CmdId = 1
)

// Enum value maps for PathfindingEnterSceneRsp_CmdId.
var (
	PathfindingEnterSceneRsp_CmdId_name = map[int32]string{
		0:    "NONE",
		2306: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	PathfindingEnterSceneRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2306,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x PathfindingEnterSceneRsp_CmdId) Enum() *PathfindingEnterSceneRsp_CmdId {
	p := new(PathfindingEnterSceneRsp_CmdId)
	*p = x
	return p
}

func (x PathfindingEnterSceneRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PathfindingEnterSceneRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_PathfindingEnterSceneRsp_proto_enumTypes[0].Descriptor()
}

func (PathfindingEnterSceneRsp_CmdId) Type() protoreflect.EnumType {
	return &file_PathfindingEnterSceneRsp_proto_enumTypes[0]
}

func (x PathfindingEnterSceneRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PathfindingEnterSceneRsp_CmdId.Descriptor instead.
func (PathfindingEnterSceneRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_PathfindingEnterSceneRsp_proto_rawDescGZIP(), []int{0, 0}
}

type PathfindingEnterSceneRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32 `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *PathfindingEnterSceneRsp) Reset() {
	*x = PathfindingEnterSceneRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PathfindingEnterSceneRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathfindingEnterSceneRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathfindingEnterSceneRsp) ProtoMessage() {}

func (x *PathfindingEnterSceneRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PathfindingEnterSceneRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathfindingEnterSceneRsp.ProtoReflect.Descriptor instead.
func (*PathfindingEnterSceneRsp) Descriptor() ([]byte, []int) {
	return file_PathfindingEnterSceneRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PathfindingEnterSceneRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_PathfindingEnterSceneRsp_proto protoreflect.FileDescriptor

var file_PathfindingEnterSceneRsp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x50, 0x61, 0x74, 0x68, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x83, 0x01, 0x0a, 0x18, 0x50, 0x61, 0x74, 0x68, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d,
	0x44, 0x5f, 0x49, 0x44, 0x10, 0x82, 0x12, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45,
	0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PathfindingEnterSceneRsp_proto_rawDescOnce sync.Once
	file_PathfindingEnterSceneRsp_proto_rawDescData = file_PathfindingEnterSceneRsp_proto_rawDesc
)

func file_PathfindingEnterSceneRsp_proto_rawDescGZIP() []byte {
	file_PathfindingEnterSceneRsp_proto_rawDescOnce.Do(func() {
		file_PathfindingEnterSceneRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PathfindingEnterSceneRsp_proto_rawDescData)
	})
	return file_PathfindingEnterSceneRsp_proto_rawDescData
}

var file_PathfindingEnterSceneRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PathfindingEnterSceneRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PathfindingEnterSceneRsp_proto_goTypes = []interface{}{
	(PathfindingEnterSceneRsp_CmdId)(0), // 0: PathfindingEnterSceneRsp.CmdId
	(*PathfindingEnterSceneRsp)(nil),    // 1: PathfindingEnterSceneRsp
}
var file_PathfindingEnterSceneRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PathfindingEnterSceneRsp_proto_init() }
func file_PathfindingEnterSceneRsp_proto_init() {
	if File_PathfindingEnterSceneRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PathfindingEnterSceneRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathfindingEnterSceneRsp); i {
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
			RawDescriptor: file_PathfindingEnterSceneRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PathfindingEnterSceneRsp_proto_goTypes,
		DependencyIndexes: file_PathfindingEnterSceneRsp_proto_depIdxs,
		EnumInfos:         file_PathfindingEnterSceneRsp_proto_enumTypes,
		MessageInfos:      file_PathfindingEnterSceneRsp_proto_msgTypes,
	}.Build()
	File_PathfindingEnterSceneRsp_proto = out.File
	file_PathfindingEnterSceneRsp_proto_rawDesc = nil
	file_PathfindingEnterSceneRsp_proto_goTypes = nil
	file_PathfindingEnterSceneRsp_proto_depIdxs = nil
}
