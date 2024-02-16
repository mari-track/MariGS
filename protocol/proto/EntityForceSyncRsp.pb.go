// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: EntityForceSyncRsp.proto

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

type EntityForceSyncRsp_CmdId int32

const (
	EntityForceSyncRsp_NONE             EntityForceSyncRsp_CmdId = 0
	EntityForceSyncRsp_CMD_ID           EntityForceSyncRsp_CmdId = 236
	EntityForceSyncRsp_ENET_CHANNEL_ID  EntityForceSyncRsp_CmdId = 0
	EntityForceSyncRsp_ENET_IS_RELIABLE EntityForceSyncRsp_CmdId = 1
)

// Enum value maps for EntityForceSyncRsp_CmdId.
var (
	EntityForceSyncRsp_CmdId_name = map[int32]string{
		0:   "NONE",
		236: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	EntityForceSyncRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           236,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x EntityForceSyncRsp_CmdId) Enum() *EntityForceSyncRsp_CmdId {
	p := new(EntityForceSyncRsp_CmdId)
	*p = x
	return p
}

func (x EntityForceSyncRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntityForceSyncRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_EntityForceSyncRsp_proto_enumTypes[0].Descriptor()
}

func (EntityForceSyncRsp_CmdId) Type() protoreflect.EnumType {
	return &file_EntityForceSyncRsp_proto_enumTypes[0]
}

func (x EntityForceSyncRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntityForceSyncRsp_CmdId.Descriptor instead.
func (EntityForceSyncRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_EntityForceSyncRsp_proto_rawDescGZIP(), []int{0, 0}
}

type EntityForceSyncRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode    int32       `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	EntityId   uint32      `protobuf:"varint,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	FailMotion *MotionInfo `protobuf:"bytes,3,opt,name=fail_motion,json=failMotion,proto3" json:"fail_motion,omitempty"`
	SceneTime  uint32      `protobuf:"varint,4,opt,name=scene_time,json=sceneTime,proto3" json:"scene_time,omitempty"`
}

func (x *EntityForceSyncRsp) Reset() {
	*x = EntityForceSyncRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityForceSyncRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityForceSyncRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityForceSyncRsp) ProtoMessage() {}

func (x *EntityForceSyncRsp) ProtoReflect() protoreflect.Message {
	mi := &file_EntityForceSyncRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityForceSyncRsp.ProtoReflect.Descriptor instead.
func (*EntityForceSyncRsp) Descriptor() ([]byte, []int) {
	return file_EntityForceSyncRsp_proto_rawDescGZIP(), []int{0}
}

func (x *EntityForceSyncRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *EntityForceSyncRsp) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *EntityForceSyncRsp) GetFailMotion() *MotionInfo {
	if x != nil {
		return x.FailMotion
	}
	return nil
}

func (x *EntityForceSyncRsp) GetSceneTime() uint32 {
	if x != nil {
		return x.SceneTime
	}
	return 0
}

var File_EntityForceSyncRsp_proto protoreflect.FileDescriptor

var file_EntityForceSyncRsp_proto_rawDesc = []byte{
	0x0a, 0x18, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x53, 0x79, 0x6e,
	0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x12, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x46, 0x6f,
	0x72, 0x63, 0x65, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x32, 0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x66, 0x61, 0x69, 0x6c, 0x4d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x63, 0x65, 0x6e, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49,
	0x44, 0x10, 0xec, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a,
	0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EntityForceSyncRsp_proto_rawDescOnce sync.Once
	file_EntityForceSyncRsp_proto_rawDescData = file_EntityForceSyncRsp_proto_rawDesc
)

func file_EntityForceSyncRsp_proto_rawDescGZIP() []byte {
	file_EntityForceSyncRsp_proto_rawDescOnce.Do(func() {
		file_EntityForceSyncRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_EntityForceSyncRsp_proto_rawDescData)
	})
	return file_EntityForceSyncRsp_proto_rawDescData
}

var file_EntityForceSyncRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EntityForceSyncRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EntityForceSyncRsp_proto_goTypes = []interface{}{
	(EntityForceSyncRsp_CmdId)(0), // 0: proto.EntityForceSyncRsp.CmdId
	(*EntityForceSyncRsp)(nil),    // 1: proto.EntityForceSyncRsp
	(*MotionInfo)(nil),            // 2: proto.MotionInfo
}
var file_EntityForceSyncRsp_proto_depIdxs = []int32{
	2, // 0: proto.EntityForceSyncRsp.fail_motion:type_name -> proto.MotionInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EntityForceSyncRsp_proto_init() }
func file_EntityForceSyncRsp_proto_init() {
	if File_EntityForceSyncRsp_proto != nil {
		return
	}
	file_MotionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EntityForceSyncRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityForceSyncRsp); i {
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
			RawDescriptor: file_EntityForceSyncRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EntityForceSyncRsp_proto_goTypes,
		DependencyIndexes: file_EntityForceSyncRsp_proto_depIdxs,
		EnumInfos:         file_EntityForceSyncRsp_proto_enumTypes,
		MessageInfos:      file_EntityForceSyncRsp_proto_msgTypes,
	}.Build()
	File_EntityForceSyncRsp_proto = out.File
	file_EntityForceSyncRsp_proto_rawDesc = nil
	file_EntityForceSyncRsp_proto_goTypes = nil
	file_EntityForceSyncRsp_proto_depIdxs = nil
}