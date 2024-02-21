// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SceneEntityMoveNotify.proto

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

type SceneEntityMoveNotify_CmdId int32

const (
	SceneEntityMoveNotify_NONE             SceneEntityMoveNotify_CmdId = 0
	SceneEntityMoveNotify_CMD_ID           SceneEntityMoveNotify_CmdId = 212
	SceneEntityMoveNotify_ENET_CHANNEL_ID  SceneEntityMoveNotify_CmdId = 1
	SceneEntityMoveNotify_ENET_IS_RELIABLE SceneEntityMoveNotify_CmdId = 0
)

// Enum value maps for SceneEntityMoveNotify_CmdId.
var (
	SceneEntityMoveNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		212: "CMD_ID",
		1:   "ENET_CHANNEL_ID",
		// Duplicate value: 0: "ENET_IS_RELIABLE",
	}
	SceneEntityMoveNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           212,
		"ENET_CHANNEL_ID":  1,
		"ENET_IS_RELIABLE": 0,
	}
)

func (x SceneEntityMoveNotify_CmdId) Enum() *SceneEntityMoveNotify_CmdId {
	p := new(SceneEntityMoveNotify_CmdId)
	*p = x
	return p
}

func (x SceneEntityMoveNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SceneEntityMoveNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_SceneEntityMoveNotify_proto_enumTypes[0].Descriptor()
}

func (SceneEntityMoveNotify_CmdId) Type() protoreflect.EnumType {
	return &file_SceneEntityMoveNotify_proto_enumTypes[0]
}

func (x SceneEntityMoveNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SceneEntityMoveNotify_CmdId.Descriptor instead.
func (SceneEntityMoveNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_SceneEntityMoveNotify_proto_rawDescGZIP(), []int{0, 0}
}

type SceneEntityMoveNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId    uint32      `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	MotionInfo  *MotionInfo `protobuf:"bytes,2,opt,name=motion_info,json=motionInfo,proto3" json:"motion_info,omitempty"`
	SceneTime   uint32      `protobuf:"varint,3,opt,name=scene_time,json=sceneTime,proto3" json:"scene_time,omitempty"`
	ReliableSeq uint32      `protobuf:"varint,4,opt,name=reliable_seq,json=reliableSeq,proto3" json:"reliable_seq,omitempty"`
}

func (x *SceneEntityMoveNotify) Reset() {
	*x = SceneEntityMoveNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneEntityMoveNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneEntityMoveNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneEntityMoveNotify) ProtoMessage() {}

func (x *SceneEntityMoveNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SceneEntityMoveNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneEntityMoveNotify.ProtoReflect.Descriptor instead.
func (*SceneEntityMoveNotify) Descriptor() ([]byte, []int) {
	return file_SceneEntityMoveNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SceneEntityMoveNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *SceneEntityMoveNotify) GetMotionInfo() *MotionInfo {
	if x != nil {
		return x.MotionInfo
	}
	return nil
}

func (x *SceneEntityMoveNotify) GetSceneTime() uint32 {
	if x != nil {
		return x.SceneTime
	}
	return 0
}

func (x *SceneEntityMoveNotify) GetReliableSeq() uint32 {
	if x != nil {
		return x.ReliableSeq
	}
	return 0
}

var File_SceneEntityMoveNotify_proto protoreflect.FileDescriptor

var file_SceneEntityMoveNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x76,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf3, 0x01, 0x0a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d,
	0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x0b, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x65, 0x71, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44,
	0x5f, 0x49, 0x44, 0x10, 0xd4, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x00, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneEntityMoveNotify_proto_rawDescOnce sync.Once
	file_SceneEntityMoveNotify_proto_rawDescData = file_SceneEntityMoveNotify_proto_rawDesc
)

func file_SceneEntityMoveNotify_proto_rawDescGZIP() []byte {
	file_SceneEntityMoveNotify_proto_rawDescOnce.Do(func() {
		file_SceneEntityMoveNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneEntityMoveNotify_proto_rawDescData)
	})
	return file_SceneEntityMoveNotify_proto_rawDescData
}

var file_SceneEntityMoveNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SceneEntityMoveNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneEntityMoveNotify_proto_goTypes = []interface{}{
	(SceneEntityMoveNotify_CmdId)(0), // 0: SceneEntityMoveNotify.CmdId
	(*SceneEntityMoveNotify)(nil),    // 1: SceneEntityMoveNotify
	(*MotionInfo)(nil),               // 2: MotionInfo
}
var file_SceneEntityMoveNotify_proto_depIdxs = []int32{
	2, // 0: SceneEntityMoveNotify.motion_info:type_name -> MotionInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SceneEntityMoveNotify_proto_init() }
func file_SceneEntityMoveNotify_proto_init() {
	if File_SceneEntityMoveNotify_proto != nil {
		return
	}
	file_MotionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneEntityMoveNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneEntityMoveNotify); i {
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
			RawDescriptor: file_SceneEntityMoveNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneEntityMoveNotify_proto_goTypes,
		DependencyIndexes: file_SceneEntityMoveNotify_proto_depIdxs,
		EnumInfos:         file_SceneEntityMoveNotify_proto_enumTypes,
		MessageInfos:      file_SceneEntityMoveNotify_proto_msgTypes,
	}.Build()
	File_SceneEntityMoveNotify_proto = out.File
	file_SceneEntityMoveNotify_proto_rawDesc = nil
	file_SceneEntityMoveNotify_proto_goTypes = nil
	file_SceneEntityMoveNotify_proto_depIdxs = nil
}
