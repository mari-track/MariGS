// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SceneKickPlayerNotify.proto

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

type SceneKickPlayerNotify_CmdId int32

const (
	SceneKickPlayerNotify_NONE             SceneKickPlayerNotify_CmdId = 0
	SceneKickPlayerNotify_CMD_ID           SceneKickPlayerNotify_CmdId = 273
	SceneKickPlayerNotify_ENET_CHANNEL_ID  SceneKickPlayerNotify_CmdId = 0
	SceneKickPlayerNotify_ENET_IS_RELIABLE SceneKickPlayerNotify_CmdId = 1
	SceneKickPlayerNotify_IS_ALLOW_CLIENT  SceneKickPlayerNotify_CmdId = 1
)

// Enum value maps for SceneKickPlayerNotify_CmdId.
var (
	SceneKickPlayerNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		273: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	SceneKickPlayerNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           273,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x SceneKickPlayerNotify_CmdId) Enum() *SceneKickPlayerNotify_CmdId {
	p := new(SceneKickPlayerNotify_CmdId)
	*p = x
	return p
}

func (x SceneKickPlayerNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SceneKickPlayerNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_SceneKickPlayerNotify_proto_enumTypes[0].Descriptor()
}

func (SceneKickPlayerNotify_CmdId) Type() protoreflect.EnumType {
	return &file_SceneKickPlayerNotify_proto_enumTypes[0]
}

func (x SceneKickPlayerNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SceneKickPlayerNotify_CmdId.Descriptor instead.
func (SceneKickPlayerNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_SceneKickPlayerNotify_proto_rawDescGZIP(), []int{0, 0}
}

type SceneKickPlayerNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KickerUid uint32 `protobuf:"varint,1,opt,name=kicker_uid,json=kickerUid,proto3" json:"kicker_uid,omitempty"`
	TargetUid uint32 `protobuf:"varint,2,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
}

func (x *SceneKickPlayerNotify) Reset() {
	*x = SceneKickPlayerNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneKickPlayerNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneKickPlayerNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneKickPlayerNotify) ProtoMessage() {}

func (x *SceneKickPlayerNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SceneKickPlayerNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneKickPlayerNotify.ProtoReflect.Descriptor instead.
func (*SceneKickPlayerNotify) Descriptor() ([]byte, []int) {
	return file_SceneKickPlayerNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SceneKickPlayerNotify) GetKickerUid() uint32 {
	if x != nil {
		return x.KickerUid
	}
	return 0
}

func (x *SceneKickPlayerNotify) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

var File_SceneKickPlayerNotify_proto protoreflect.FileDescriptor

var file_SceneKickPlayerNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4b, 0x69, 0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4b, 0x69,
	0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x6b, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x6b, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x05,
	0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x91, 0x02, 0x12, 0x13, 0x0a, 0x0f,
	0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c,
	0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c,
	0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneKickPlayerNotify_proto_rawDescOnce sync.Once
	file_SceneKickPlayerNotify_proto_rawDescData = file_SceneKickPlayerNotify_proto_rawDesc
)

func file_SceneKickPlayerNotify_proto_rawDescGZIP() []byte {
	file_SceneKickPlayerNotify_proto_rawDescOnce.Do(func() {
		file_SceneKickPlayerNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneKickPlayerNotify_proto_rawDescData)
	})
	return file_SceneKickPlayerNotify_proto_rawDescData
}

var file_SceneKickPlayerNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SceneKickPlayerNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneKickPlayerNotify_proto_goTypes = []interface{}{
	(SceneKickPlayerNotify_CmdId)(0), // 0: proto.SceneKickPlayerNotify.CmdId
	(*SceneKickPlayerNotify)(nil),    // 1: proto.SceneKickPlayerNotify
}
var file_SceneKickPlayerNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SceneKickPlayerNotify_proto_init() }
func file_SceneKickPlayerNotify_proto_init() {
	if File_SceneKickPlayerNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SceneKickPlayerNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneKickPlayerNotify); i {
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
			RawDescriptor: file_SceneKickPlayerNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneKickPlayerNotify_proto_goTypes,
		DependencyIndexes: file_SceneKickPlayerNotify_proto_depIdxs,
		EnumInfos:         file_SceneKickPlayerNotify_proto_enumTypes,
		MessageInfos:      file_SceneKickPlayerNotify_proto_msgTypes,
	}.Build()
	File_SceneKickPlayerNotify_proto = out.File
	file_SceneKickPlayerNotify_proto_rawDesc = nil
	file_SceneKickPlayerNotify_proto_goTypes = nil
	file_SceneKickPlayerNotify_proto_depIdxs = nil
}
