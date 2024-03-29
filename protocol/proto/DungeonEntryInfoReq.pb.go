// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: DungeonEntryInfoReq.proto

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

type DungeonEntryInfoReq_CmdId int32

const (
	DungeonEntryInfoReq_NONE             DungeonEntryInfoReq_CmdId = 0
	DungeonEntryInfoReq_CMD_ID           DungeonEntryInfoReq_CmdId = 901
	DungeonEntryInfoReq_ENET_CHANNEL_ID  DungeonEntryInfoReq_CmdId = 0
	DungeonEntryInfoReq_ENET_IS_RELIABLE DungeonEntryInfoReq_CmdId = 1
	DungeonEntryInfoReq_IS_ALLOW_CLIENT  DungeonEntryInfoReq_CmdId = 1
)

// Enum value maps for DungeonEntryInfoReq_CmdId.
var (
	DungeonEntryInfoReq_CmdId_name = map[int32]string{
		0:   "NONE",
		901: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	DungeonEntryInfoReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           901,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x DungeonEntryInfoReq_CmdId) Enum() *DungeonEntryInfoReq_CmdId {
	p := new(DungeonEntryInfoReq_CmdId)
	*p = x
	return p
}

func (x DungeonEntryInfoReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DungeonEntryInfoReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_DungeonEntryInfoReq_proto_enumTypes[0].Descriptor()
}

func (DungeonEntryInfoReq_CmdId) Type() protoreflect.EnumType {
	return &file_DungeonEntryInfoReq_proto_enumTypes[0]
}

func (x DungeonEntryInfoReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DungeonEntryInfoReq_CmdId.Descriptor instead.
func (DungeonEntryInfoReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_DungeonEntryInfoReq_proto_rawDescGZIP(), []int{0, 0}
}

type DungeonEntryInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PointId uint32 `protobuf:"varint,1,opt,name=point_id,json=pointId,proto3" json:"point_id,omitempty"`
}

func (x *DungeonEntryInfoReq) Reset() {
	*x = DungeonEntryInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DungeonEntryInfoReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DungeonEntryInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DungeonEntryInfoReq) ProtoMessage() {}

func (x *DungeonEntryInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_DungeonEntryInfoReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DungeonEntryInfoReq.ProtoReflect.Descriptor instead.
func (*DungeonEntryInfoReq) Descriptor() ([]byte, []int) {
	return file_DungeonEntryInfoReq_proto_rawDescGZIP(), []int{0}
}

func (x *DungeonEntryInfoReq) GetPointId() uint32 {
	if x != nil {
		return x.PointId
	}
	return 0
}

var File_DungeonEntryInfoReq_proto protoreflect.FileDescriptor

var file_DungeonEntryInfoReq_proto_rawDesc = []byte{
	0x0a, 0x19, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x13,
	0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x62,
	0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x85, 0x07, 0x12, 0x13,
	0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52,
	0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f,
	0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02,
	0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DungeonEntryInfoReq_proto_rawDescOnce sync.Once
	file_DungeonEntryInfoReq_proto_rawDescData = file_DungeonEntryInfoReq_proto_rawDesc
)

func file_DungeonEntryInfoReq_proto_rawDescGZIP() []byte {
	file_DungeonEntryInfoReq_proto_rawDescOnce.Do(func() {
		file_DungeonEntryInfoReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_DungeonEntryInfoReq_proto_rawDescData)
	})
	return file_DungeonEntryInfoReq_proto_rawDescData
}

var file_DungeonEntryInfoReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DungeonEntryInfoReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DungeonEntryInfoReq_proto_goTypes = []interface{}{
	(DungeonEntryInfoReq_CmdId)(0), // 0: DungeonEntryInfoReq.CmdId
	(*DungeonEntryInfoReq)(nil),    // 1: DungeonEntryInfoReq
}
var file_DungeonEntryInfoReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DungeonEntryInfoReq_proto_init() }
func file_DungeonEntryInfoReq_proto_init() {
	if File_DungeonEntryInfoReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DungeonEntryInfoReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DungeonEntryInfoReq); i {
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
			RawDescriptor: file_DungeonEntryInfoReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DungeonEntryInfoReq_proto_goTypes,
		DependencyIndexes: file_DungeonEntryInfoReq_proto_depIdxs,
		EnumInfos:         file_DungeonEntryInfoReq_proto_enumTypes,
		MessageInfos:      file_DungeonEntryInfoReq_proto_msgTypes,
	}.Build()
	File_DungeonEntryInfoReq_proto = out.File
	file_DungeonEntryInfoReq_proto_rawDesc = nil
	file_DungeonEntryInfoReq_proto_goTypes = nil
	file_DungeonEntryInfoReq_proto_depIdxs = nil
}
