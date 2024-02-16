// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: QuestDestroyEntityRsp.proto

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

type QuestDestroyEntityRsp_CmdId int32

const (
	QuestDestroyEntityRsp_NONE             QuestDestroyEntityRsp_CmdId = 0
	QuestDestroyEntityRsp_CMD_ID           QuestDestroyEntityRsp_CmdId = 413
	QuestDestroyEntityRsp_ENET_CHANNEL_ID  QuestDestroyEntityRsp_CmdId = 0
	QuestDestroyEntityRsp_ENET_IS_RELIABLE QuestDestroyEntityRsp_CmdId = 1
)

// Enum value maps for QuestDestroyEntityRsp_CmdId.
var (
	QuestDestroyEntityRsp_CmdId_name = map[int32]string{
		0:   "NONE",
		413: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	QuestDestroyEntityRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           413,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x QuestDestroyEntityRsp_CmdId) Enum() *QuestDestroyEntityRsp_CmdId {
	p := new(QuestDestroyEntityRsp_CmdId)
	*p = x
	return p
}

func (x QuestDestroyEntityRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuestDestroyEntityRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_QuestDestroyEntityRsp_proto_enumTypes[0].Descriptor()
}

func (QuestDestroyEntityRsp_CmdId) Type() protoreflect.EnumType {
	return &file_QuestDestroyEntityRsp_proto_enumTypes[0]
}

func (x QuestDestroyEntityRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuestDestroyEntityRsp_CmdId.Descriptor instead.
func (QuestDestroyEntityRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_QuestDestroyEntityRsp_proto_rawDescGZIP(), []int{0, 0}
}

type QuestDestroyEntityRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode  int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	EntityId uint32 `protobuf:"varint,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	QuestId  uint32 `protobuf:"varint,3,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
	SceneId  uint32 `protobuf:"varint,4,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
}

func (x *QuestDestroyEntityRsp) Reset() {
	*x = QuestDestroyEntityRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QuestDestroyEntityRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestDestroyEntityRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestDestroyEntityRsp) ProtoMessage() {}

func (x *QuestDestroyEntityRsp) ProtoReflect() protoreflect.Message {
	mi := &file_QuestDestroyEntityRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestDestroyEntityRsp.ProtoReflect.Descriptor instead.
func (*QuestDestroyEntityRsp) Descriptor() ([]byte, []int) {
	return file_QuestDestroyEntityRsp_proto_rawDescGZIP(), []int{0}
}

func (x *QuestDestroyEntityRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *QuestDestroyEntityRsp) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *QuestDestroyEntityRsp) GetQuestId() uint32 {
	if x != nil {
		return x.QuestId
	}
	return 0
}

func (x *QuestDestroyEntityRsp) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

var File_QuestDestroyEntityRsp_proto protoreflect.FileDescriptor

var file_QuestDestroyEntityRsp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x51, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65,
	0x73, 0x74, 0x72, 0x6f, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43,
	0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x9d, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_QuestDestroyEntityRsp_proto_rawDescOnce sync.Once
	file_QuestDestroyEntityRsp_proto_rawDescData = file_QuestDestroyEntityRsp_proto_rawDesc
)

func file_QuestDestroyEntityRsp_proto_rawDescGZIP() []byte {
	file_QuestDestroyEntityRsp_proto_rawDescOnce.Do(func() {
		file_QuestDestroyEntityRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_QuestDestroyEntityRsp_proto_rawDescData)
	})
	return file_QuestDestroyEntityRsp_proto_rawDescData
}

var file_QuestDestroyEntityRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_QuestDestroyEntityRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_QuestDestroyEntityRsp_proto_goTypes = []interface{}{
	(QuestDestroyEntityRsp_CmdId)(0), // 0: proto.QuestDestroyEntityRsp.CmdId
	(*QuestDestroyEntityRsp)(nil),    // 1: proto.QuestDestroyEntityRsp
}
var file_QuestDestroyEntityRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_QuestDestroyEntityRsp_proto_init() }
func file_QuestDestroyEntityRsp_proto_init() {
	if File_QuestDestroyEntityRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_QuestDestroyEntityRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestDestroyEntityRsp); i {
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
			RawDescriptor: file_QuestDestroyEntityRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_QuestDestroyEntityRsp_proto_goTypes,
		DependencyIndexes: file_QuestDestroyEntityRsp_proto_depIdxs,
		EnumInfos:         file_QuestDestroyEntityRsp_proto_enumTypes,
		MessageInfos:      file_QuestDestroyEntityRsp_proto_msgTypes,
	}.Build()
	File_QuestDestroyEntityRsp_proto = out.File
	file_QuestDestroyEntityRsp_proto_rawDesc = nil
	file_QuestDestroyEntityRsp_proto_goTypes = nil
	file_QuestDestroyEntityRsp_proto_depIdxs = nil
}
