// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: QuestTransmitRsp.proto

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

type QuestTransmitRsp_CmdId int32

const (
	QuestTransmitRsp_NONE             QuestTransmitRsp_CmdId = 0
	QuestTransmitRsp_CMD_ID           QuestTransmitRsp_CmdId = 434
	QuestTransmitRsp_ENET_CHANNEL_ID  QuestTransmitRsp_CmdId = 0
	QuestTransmitRsp_ENET_IS_RELIABLE QuestTransmitRsp_CmdId = 1
)

// Enum value maps for QuestTransmitRsp_CmdId.
var (
	QuestTransmitRsp_CmdId_name = map[int32]string{
		0:   "NONE",
		434: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	QuestTransmitRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           434,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x QuestTransmitRsp_CmdId) Enum() *QuestTransmitRsp_CmdId {
	p := new(QuestTransmitRsp_CmdId)
	*p = x
	return p
}

func (x QuestTransmitRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuestTransmitRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_QuestTransmitRsp_proto_enumTypes[0].Descriptor()
}

func (QuestTransmitRsp_CmdId) Type() protoreflect.EnumType {
	return &file_QuestTransmitRsp_proto_enumTypes[0]
}

func (x QuestTransmitRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuestTransmitRsp_CmdId.Descriptor instead.
func (QuestTransmitRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_QuestTransmitRsp_proto_rawDescGZIP(), []int{0, 0}
}

type QuestTransmitRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	QuestId uint32 `protobuf:"varint,2,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
	PointId uint32 `protobuf:"varint,3,opt,name=point_id,json=pointId,proto3" json:"point_id,omitempty"`
}

func (x *QuestTransmitRsp) Reset() {
	*x = QuestTransmitRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QuestTransmitRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestTransmitRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestTransmitRsp) ProtoMessage() {}

func (x *QuestTransmitRsp) ProtoReflect() protoreflect.Message {
	mi := &file_QuestTransmitRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestTransmitRsp.ProtoReflect.Descriptor instead.
func (*QuestTransmitRsp) Descriptor() ([]byte, []int) {
	return file_QuestTransmitRsp_proto_rawDescGZIP(), []int{0}
}

func (x *QuestTransmitRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *QuestTransmitRsp) GetQuestId() uint32 {
	if x != nil {
		return x.QuestId
	}
	return 0
}

func (x *QuestTransmitRsp) GetPointId() uint32 {
	if x != nil {
		return x.PointId
	}
	return 0
}

var File_QuestTransmitRsp_proto protoreflect.FileDescriptor

var file_QuestTransmitRsp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x51, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x10, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4d, 0x0a,
	0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xb2, 0x03, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45,
	0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_QuestTransmitRsp_proto_rawDescOnce sync.Once
	file_QuestTransmitRsp_proto_rawDescData = file_QuestTransmitRsp_proto_rawDesc
)

func file_QuestTransmitRsp_proto_rawDescGZIP() []byte {
	file_QuestTransmitRsp_proto_rawDescOnce.Do(func() {
		file_QuestTransmitRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_QuestTransmitRsp_proto_rawDescData)
	})
	return file_QuestTransmitRsp_proto_rawDescData
}

var file_QuestTransmitRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_QuestTransmitRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_QuestTransmitRsp_proto_goTypes = []interface{}{
	(QuestTransmitRsp_CmdId)(0), // 0: QuestTransmitRsp.CmdId
	(*QuestTransmitRsp)(nil),    // 1: QuestTransmitRsp
}
var file_QuestTransmitRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_QuestTransmitRsp_proto_init() }
func file_QuestTransmitRsp_proto_init() {
	if File_QuestTransmitRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_QuestTransmitRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestTransmitRsp); i {
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
			RawDescriptor: file_QuestTransmitRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_QuestTransmitRsp_proto_goTypes,
		DependencyIndexes: file_QuestTransmitRsp_proto_depIdxs,
		EnumInfos:         file_QuestTransmitRsp_proto_enumTypes,
		MessageInfos:      file_QuestTransmitRsp_proto_msgTypes,
	}.Build()
	File_QuestTransmitRsp_proto = out.File
	file_QuestTransmitRsp_proto_rawDesc = nil
	file_QuestTransmitRsp_proto_goTypes = nil
	file_QuestTransmitRsp_proto_depIdxs = nil
}
