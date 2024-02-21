// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: QuestCreateEntityRsp.proto

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

type QuestCreateEntityRsp_CmdId int32

const (
	QuestCreateEntityRsp_NONE             QuestCreateEntityRsp_CmdId = 0
	QuestCreateEntityRsp_CMD_ID           QuestCreateEntityRsp_CmdId = 411
	QuestCreateEntityRsp_ENET_CHANNEL_ID  QuestCreateEntityRsp_CmdId = 0
	QuestCreateEntityRsp_ENET_IS_RELIABLE QuestCreateEntityRsp_CmdId = 1
)

// Enum value maps for QuestCreateEntityRsp_CmdId.
var (
	QuestCreateEntityRsp_CmdId_name = map[int32]string{
		0:   "NONE",
		411: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	QuestCreateEntityRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           411,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x QuestCreateEntityRsp_CmdId) Enum() *QuestCreateEntityRsp_CmdId {
	p := new(QuestCreateEntityRsp_CmdId)
	*p = x
	return p
}

func (x QuestCreateEntityRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuestCreateEntityRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_QuestCreateEntityRsp_proto_enumTypes[0].Descriptor()
}

func (QuestCreateEntityRsp_CmdId) Type() protoreflect.EnumType {
	return &file_QuestCreateEntityRsp_proto_enumTypes[0]
}

func (x QuestCreateEntityRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuestCreateEntityRsp_CmdId.Descriptor instead.
func (QuestCreateEntityRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_QuestCreateEntityRsp_proto_rawDescGZIP(), []int{0, 0}
}

type QuestCreateEntityRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode       int32             `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	EntityId      uint32            `protobuf:"varint,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Entity        *CreateEntityInfo `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	QuestId       uint32            `protobuf:"varint,7,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
	IsRewind      bool              `protobuf:"varint,8,opt,name=is_rewind,json=isRewind,proto3" json:"is_rewind,omitempty"`
	ParentQuestId uint32            `protobuf:"varint,9,opt,name=parent_quest_id,json=parentQuestId,proto3" json:"parent_quest_id,omitempty"`
}

func (x *QuestCreateEntityRsp) Reset() {
	*x = QuestCreateEntityRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QuestCreateEntityRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestCreateEntityRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestCreateEntityRsp) ProtoMessage() {}

func (x *QuestCreateEntityRsp) ProtoReflect() protoreflect.Message {
	mi := &file_QuestCreateEntityRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestCreateEntityRsp.ProtoReflect.Descriptor instead.
func (*QuestCreateEntityRsp) Descriptor() ([]byte, []int) {
	return file_QuestCreateEntityRsp_proto_rawDescGZIP(), []int{0}
}

func (x *QuestCreateEntityRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *QuestCreateEntityRsp) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *QuestCreateEntityRsp) GetEntity() *CreateEntityInfo {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *QuestCreateEntityRsp) GetQuestId() uint32 {
	if x != nil {
		return x.QuestId
	}
	return 0
}

func (x *QuestCreateEntityRsp) GetIsRewind() bool {
	if x != nil {
		return x.IsRewind
	}
	return false
}

func (x *QuestCreateEntityRsp) GetParentQuestId() uint32 {
	if x != nil {
		return x.ParentQuestId
	}
	return 0
}

var File_QuestCreateEntityRsp_proto protoreflect.FileDescriptor

var file_QuestCreateEntityRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x51, 0x75, 0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x73, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x72, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x52, 0x65, 0x77, 0x69, 0x6e, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22,
	0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x9b, 0x03, 0x12,
	0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f,
	0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_QuestCreateEntityRsp_proto_rawDescOnce sync.Once
	file_QuestCreateEntityRsp_proto_rawDescData = file_QuestCreateEntityRsp_proto_rawDesc
)

func file_QuestCreateEntityRsp_proto_rawDescGZIP() []byte {
	file_QuestCreateEntityRsp_proto_rawDescOnce.Do(func() {
		file_QuestCreateEntityRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_QuestCreateEntityRsp_proto_rawDescData)
	})
	return file_QuestCreateEntityRsp_proto_rawDescData
}

var file_QuestCreateEntityRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_QuestCreateEntityRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_QuestCreateEntityRsp_proto_goTypes = []interface{}{
	(QuestCreateEntityRsp_CmdId)(0), // 0: QuestCreateEntityRsp.CmdId
	(*QuestCreateEntityRsp)(nil),    // 1: QuestCreateEntityRsp
	(*CreateEntityInfo)(nil),        // 2: CreateEntityInfo
}
var file_QuestCreateEntityRsp_proto_depIdxs = []int32{
	2, // 0: QuestCreateEntityRsp.entity:type_name -> CreateEntityInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_QuestCreateEntityRsp_proto_init() }
func file_QuestCreateEntityRsp_proto_init() {
	if File_QuestCreateEntityRsp_proto != nil {
		return
	}
	file_CreateEntityInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_QuestCreateEntityRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestCreateEntityRsp); i {
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
			RawDescriptor: file_QuestCreateEntityRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_QuestCreateEntityRsp_proto_goTypes,
		DependencyIndexes: file_QuestCreateEntityRsp_proto_depIdxs,
		EnumInfos:         file_QuestCreateEntityRsp_proto_enumTypes,
		MessageInfos:      file_QuestCreateEntityRsp_proto_msgTypes,
	}.Build()
	File_QuestCreateEntityRsp_proto = out.File
	file_QuestCreateEntityRsp_proto_rawDesc = nil
	file_QuestCreateEntityRsp_proto_goTypes = nil
	file_QuestCreateEntityRsp_proto_depIdxs = nil
}
