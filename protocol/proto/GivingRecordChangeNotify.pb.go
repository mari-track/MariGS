// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GivingRecordChangeNotify.proto

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

type GivingRecordChangeNotify_CmdId int32

const (
	GivingRecordChangeNotify_NONE             GivingRecordChangeNotify_CmdId = 0
	GivingRecordChangeNotify_CMD_ID           GivingRecordChangeNotify_CmdId = 156
	GivingRecordChangeNotify_ENET_CHANNEL_ID  GivingRecordChangeNotify_CmdId = 0
	GivingRecordChangeNotify_ENET_IS_RELIABLE GivingRecordChangeNotify_CmdId = 1
)

// Enum value maps for GivingRecordChangeNotify_CmdId.
var (
	GivingRecordChangeNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		156: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	GivingRecordChangeNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           156,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x GivingRecordChangeNotify_CmdId) Enum() *GivingRecordChangeNotify_CmdId {
	p := new(GivingRecordChangeNotify_CmdId)
	*p = x
	return p
}

func (x GivingRecordChangeNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GivingRecordChangeNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GivingRecordChangeNotify_proto_enumTypes[0].Descriptor()
}

func (GivingRecordChangeNotify_CmdId) Type() protoreflect.EnumType {
	return &file_GivingRecordChangeNotify_proto_enumTypes[0]
}

func (x GivingRecordChangeNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GivingRecordChangeNotify_CmdId.Descriptor instead.
func (GivingRecordChangeNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GivingRecordChangeNotify_proto_rawDescGZIP(), []int{0, 0}
}

type GivingRecordChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GivingRecord *GivingRecord `protobuf:"bytes,1,opt,name=giving_record,json=givingRecord,proto3" json:"giving_record,omitempty"`
}

func (x *GivingRecordChangeNotify) Reset() {
	*x = GivingRecordChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GivingRecordChangeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GivingRecordChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GivingRecordChangeNotify) ProtoMessage() {}

func (x *GivingRecordChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_GivingRecordChangeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GivingRecordChangeNotify.ProtoReflect.Descriptor instead.
func (*GivingRecordChangeNotify) Descriptor() ([]byte, []int) {
	return file_GivingRecordChangeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *GivingRecordChangeNotify) GetGivingRecord() *GivingRecord {
	if x != nil {
		return x.GivingRecord
	}
	return nil
}

var File_GivingRecordChangeNotify_proto protoreflect.FileDescriptor

var file_GivingRecordChangeNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x18,
	0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x38, 0x0a, 0x0d, 0x67, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x0c, 0x67, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0x9c, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e,
	0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GivingRecordChangeNotify_proto_rawDescOnce sync.Once
	file_GivingRecordChangeNotify_proto_rawDescData = file_GivingRecordChangeNotify_proto_rawDesc
)

func file_GivingRecordChangeNotify_proto_rawDescGZIP() []byte {
	file_GivingRecordChangeNotify_proto_rawDescOnce.Do(func() {
		file_GivingRecordChangeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_GivingRecordChangeNotify_proto_rawDescData)
	})
	return file_GivingRecordChangeNotify_proto_rawDescData
}

var file_GivingRecordChangeNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GivingRecordChangeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GivingRecordChangeNotify_proto_goTypes = []interface{}{
	(GivingRecordChangeNotify_CmdId)(0), // 0: proto.GivingRecordChangeNotify.CmdId
	(*GivingRecordChangeNotify)(nil),    // 1: proto.GivingRecordChangeNotify
	(*GivingRecord)(nil),                // 2: proto.GivingRecord
}
var file_GivingRecordChangeNotify_proto_depIdxs = []int32{
	2, // 0: proto.GivingRecordChangeNotify.giving_record:type_name -> proto.GivingRecord
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GivingRecordChangeNotify_proto_init() }
func file_GivingRecordChangeNotify_proto_init() {
	if File_GivingRecordChangeNotify_proto != nil {
		return
	}
	file_GivingRecord_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GivingRecordChangeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GivingRecordChangeNotify); i {
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
			RawDescriptor: file_GivingRecordChangeNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GivingRecordChangeNotify_proto_goTypes,
		DependencyIndexes: file_GivingRecordChangeNotify_proto_depIdxs,
		EnumInfos:         file_GivingRecordChangeNotify_proto_enumTypes,
		MessageInfos:      file_GivingRecordChangeNotify_proto_msgTypes,
	}.Build()
	File_GivingRecordChangeNotify_proto = out.File
	file_GivingRecordChangeNotify_proto_rawDesc = nil
	file_GivingRecordChangeNotify_proto_goTypes = nil
	file_GivingRecordChangeNotify_proto_depIdxs = nil
}
