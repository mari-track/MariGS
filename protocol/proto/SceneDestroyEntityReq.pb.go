// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SceneDestroyEntityReq.proto

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

type SceneDestroyEntityReq_CmdId int32

const (
	SceneDestroyEntityReq_NONE             SceneDestroyEntityReq_CmdId = 0
	SceneDestroyEntityReq_CMD_ID           SceneDestroyEntityReq_CmdId = 229
	SceneDestroyEntityReq_ENET_CHANNEL_ID  SceneDestroyEntityReq_CmdId = 0
	SceneDestroyEntityReq_ENET_IS_RELIABLE SceneDestroyEntityReq_CmdId = 1
	SceneDestroyEntityReq_IS_ALLOW_CLIENT  SceneDestroyEntityReq_CmdId = 1
)

// Enum value maps for SceneDestroyEntityReq_CmdId.
var (
	SceneDestroyEntityReq_CmdId_name = map[int32]string{
		0:   "NONE",
		229: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	SceneDestroyEntityReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           229,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x SceneDestroyEntityReq_CmdId) Enum() *SceneDestroyEntityReq_CmdId {
	p := new(SceneDestroyEntityReq_CmdId)
	*p = x
	return p
}

func (x SceneDestroyEntityReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SceneDestroyEntityReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_SceneDestroyEntityReq_proto_enumTypes[0].Descriptor()
}

func (SceneDestroyEntityReq_CmdId) Type() protoreflect.EnumType {
	return &file_SceneDestroyEntityReq_proto_enumTypes[0]
}

func (x SceneDestroyEntityReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SceneDestroyEntityReq_CmdId.Descriptor instead.
func (SceneDestroyEntityReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_SceneDestroyEntityReq_proto_rawDescGZIP(), []int{0, 0}
}

type SceneDestroyEntityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId uint32 `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
}

func (x *SceneDestroyEntityReq) Reset() {
	*x = SceneDestroyEntityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneDestroyEntityReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneDestroyEntityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneDestroyEntityReq) ProtoMessage() {}

func (x *SceneDestroyEntityReq) ProtoReflect() protoreflect.Message {
	mi := &file_SceneDestroyEntityReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneDestroyEntityReq.ProtoReflect.Descriptor instead.
func (*SceneDestroyEntityReq) Descriptor() ([]byte, []int) {
	return file_SceneDestroyEntityReq_proto_rawDescGZIP(), []int{0}
}

func (x *SceneDestroyEntityReq) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

var File_SceneDestroyEntityReq_proto protoreflect.FileDescriptor

var file_SceneDestroyEntityReq_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x44, 0x65,
	0x73, 0x74, 0x72, 0x6f, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x1b,
	0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x05, 0x43,
	0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xe5, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c,
	0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_SceneDestroyEntityReq_proto_rawDescOnce sync.Once
	file_SceneDestroyEntityReq_proto_rawDescData = file_SceneDestroyEntityReq_proto_rawDesc
)

func file_SceneDestroyEntityReq_proto_rawDescGZIP() []byte {
	file_SceneDestroyEntityReq_proto_rawDescOnce.Do(func() {
		file_SceneDestroyEntityReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneDestroyEntityReq_proto_rawDescData)
	})
	return file_SceneDestroyEntityReq_proto_rawDescData
}

var file_SceneDestroyEntityReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SceneDestroyEntityReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneDestroyEntityReq_proto_goTypes = []interface{}{
	(SceneDestroyEntityReq_CmdId)(0), // 0: proto.SceneDestroyEntityReq.CmdId
	(*SceneDestroyEntityReq)(nil),    // 1: proto.SceneDestroyEntityReq
}
var file_SceneDestroyEntityReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SceneDestroyEntityReq_proto_init() }
func file_SceneDestroyEntityReq_proto_init() {
	if File_SceneDestroyEntityReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SceneDestroyEntityReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneDestroyEntityReq); i {
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
			RawDescriptor: file_SceneDestroyEntityReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneDestroyEntityReq_proto_goTypes,
		DependencyIndexes: file_SceneDestroyEntityReq_proto_depIdxs,
		EnumInfos:         file_SceneDestroyEntityReq_proto_enumTypes,
		MessageInfos:      file_SceneDestroyEntityReq_proto_msgTypes,
	}.Build()
	File_SceneDestroyEntityReq_proto = out.File
	file_SceneDestroyEntityReq_proto_rawDesc = nil
	file_SceneDestroyEntityReq_proto_goTypes = nil
	file_SceneDestroyEntityReq_proto_depIdxs = nil
}
