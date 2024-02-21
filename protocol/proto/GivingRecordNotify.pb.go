// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GivingRecordNotify.proto

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

type GivingRecordNotify_CmdId int32

const (
	GivingRecordNotify_NONE             GivingRecordNotify_CmdId = 0
	GivingRecordNotify_CMD_ID           GivingRecordNotify_CmdId = 155
	GivingRecordNotify_ENET_CHANNEL_ID  GivingRecordNotify_CmdId = 0
	GivingRecordNotify_ENET_IS_RELIABLE GivingRecordNotify_CmdId = 1
)

// Enum value maps for GivingRecordNotify_CmdId.
var (
	GivingRecordNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		155: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	GivingRecordNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           155,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x GivingRecordNotify_CmdId) Enum() *GivingRecordNotify_CmdId {
	p := new(GivingRecordNotify_CmdId)
	*p = x
	return p
}

func (x GivingRecordNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GivingRecordNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GivingRecordNotify_proto_enumTypes[0].Descriptor()
}

func (GivingRecordNotify_CmdId) Type() protoreflect.EnumType {
	return &file_GivingRecordNotify_proto_enumTypes[0]
}

func (x GivingRecordNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GivingRecordNotify_CmdId.Descriptor instead.
func (GivingRecordNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GivingRecordNotify_proto_rawDescGZIP(), []int{0, 0}
}

type GivingRecordNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GivingRecordList []*GivingRecord `protobuf:"bytes,1,rep,name=giving_record_list,json=givingRecordList,proto3" json:"giving_record_list,omitempty"`
}

func (x *GivingRecordNotify) Reset() {
	*x = GivingRecordNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GivingRecordNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GivingRecordNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GivingRecordNotify) ProtoMessage() {}

func (x *GivingRecordNotify) ProtoReflect() protoreflect.Message {
	mi := &file_GivingRecordNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GivingRecordNotify.ProtoReflect.Descriptor instead.
func (*GivingRecordNotify) Descriptor() ([]byte, []int) {
	return file_GivingRecordNotify_proto_rawDescGZIP(), []int{0}
}

func (x *GivingRecordNotify) GetGivingRecordList() []*GivingRecord {
	if x != nil {
		return x.GivingRecordList
	}
	return nil
}

var File_GivingRecordNotify_proto protoreflect.FileDescriptor

var file_GivingRecordNotify_proto_rawDesc = []byte{
	0x0a, 0x18, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x47, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0,
	0x01, 0x0a, 0x12, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x3b, 0x0a, 0x12, 0x67, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x10, 0x67, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0x9b, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e,
	0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GivingRecordNotify_proto_rawDescOnce sync.Once
	file_GivingRecordNotify_proto_rawDescData = file_GivingRecordNotify_proto_rawDesc
)

func file_GivingRecordNotify_proto_rawDescGZIP() []byte {
	file_GivingRecordNotify_proto_rawDescOnce.Do(func() {
		file_GivingRecordNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_GivingRecordNotify_proto_rawDescData)
	})
	return file_GivingRecordNotify_proto_rawDescData
}

var file_GivingRecordNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GivingRecordNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GivingRecordNotify_proto_goTypes = []interface{}{
	(GivingRecordNotify_CmdId)(0), // 0: GivingRecordNotify.CmdId
	(*GivingRecordNotify)(nil),    // 1: GivingRecordNotify
	(*GivingRecord)(nil),          // 2: GivingRecord
}
var file_GivingRecordNotify_proto_depIdxs = []int32{
	2, // 0: GivingRecordNotify.giving_record_list:type_name -> GivingRecord
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GivingRecordNotify_proto_init() }
func file_GivingRecordNotify_proto_init() {
	if File_GivingRecordNotify_proto != nil {
		return
	}
	file_GivingRecord_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GivingRecordNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GivingRecordNotify); i {
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
			RawDescriptor: file_GivingRecordNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GivingRecordNotify_proto_goTypes,
		DependencyIndexes: file_GivingRecordNotify_proto_depIdxs,
		EnumInfos:         file_GivingRecordNotify_proto_enumTypes,
		MessageInfos:      file_GivingRecordNotify_proto_msgTypes,
	}.Build()
	File_GivingRecordNotify_proto = out.File
	file_GivingRecordNotify_proto_rawDesc = nil
	file_GivingRecordNotify_proto_goTypes = nil
	file_GivingRecordNotify_proto_depIdxs = nil
}
