// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: BargainTerminateNotify.proto

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

type BargainTerminateNotify_CmdId int32

const (
	BargainTerminateNotify_NONE             BargainTerminateNotify_CmdId = 0
	BargainTerminateNotify_CMD_ID           BargainTerminateNotify_CmdId = 426
	BargainTerminateNotify_ENET_CHANNEL_ID  BargainTerminateNotify_CmdId = 0
	BargainTerminateNotify_ENET_IS_RELIABLE BargainTerminateNotify_CmdId = 1
)

// Enum value maps for BargainTerminateNotify_CmdId.
var (
	BargainTerminateNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		426: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	BargainTerminateNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           426,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x BargainTerminateNotify_CmdId) Enum() *BargainTerminateNotify_CmdId {
	p := new(BargainTerminateNotify_CmdId)
	*p = x
	return p
}

func (x BargainTerminateNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BargainTerminateNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_BargainTerminateNotify_proto_enumTypes[0].Descriptor()
}

func (BargainTerminateNotify_CmdId) Type() protoreflect.EnumType {
	return &file_BargainTerminateNotify_proto_enumTypes[0]
}

func (x BargainTerminateNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BargainTerminateNotify_CmdId.Descriptor instead.
func (BargainTerminateNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_BargainTerminateNotify_proto_rawDescGZIP(), []int{0, 0}
}

type BargainTerminateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BargainId uint32 `protobuf:"varint,1,opt,name=bargain_id,json=bargainId,proto3" json:"bargain_id,omitempty"`
}

func (x *BargainTerminateNotify) Reset() {
	*x = BargainTerminateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BargainTerminateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BargainTerminateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BargainTerminateNotify) ProtoMessage() {}

func (x *BargainTerminateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_BargainTerminateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BargainTerminateNotify.ProtoReflect.Descriptor instead.
func (*BargainTerminateNotify) Descriptor() ([]byte, []int) {
	return file_BargainTerminateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *BargainTerminateNotify) GetBargainId() uint32 {
	if x != nil {
		return x.BargainId
	}
	return 0
}

var File_BargainTerminateNotify_proto protoreflect.FileDescriptor

var file_BargainTerminateNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x42, 0x61, 0x72, 0x67, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86,
	0x01, 0x0a, 0x16, 0x42, 0x61, 0x72, 0x67, 0x61, 0x69, 0x6e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x72,
	0x67, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62,
	0x61, 0x72, 0x67, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49,
	0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43,
	0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xaa, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BargainTerminateNotify_proto_rawDescOnce sync.Once
	file_BargainTerminateNotify_proto_rawDescData = file_BargainTerminateNotify_proto_rawDesc
)

func file_BargainTerminateNotify_proto_rawDescGZIP() []byte {
	file_BargainTerminateNotify_proto_rawDescOnce.Do(func() {
		file_BargainTerminateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_BargainTerminateNotify_proto_rawDescData)
	})
	return file_BargainTerminateNotify_proto_rawDescData
}

var file_BargainTerminateNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BargainTerminateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BargainTerminateNotify_proto_goTypes = []interface{}{
	(BargainTerminateNotify_CmdId)(0), // 0: BargainTerminateNotify.CmdId
	(*BargainTerminateNotify)(nil),    // 1: BargainTerminateNotify
}
var file_BargainTerminateNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BargainTerminateNotify_proto_init() }
func file_BargainTerminateNotify_proto_init() {
	if File_BargainTerminateNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BargainTerminateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BargainTerminateNotify); i {
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
			RawDescriptor: file_BargainTerminateNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BargainTerminateNotify_proto_goTypes,
		DependencyIndexes: file_BargainTerminateNotify_proto_depIdxs,
		EnumInfos:         file_BargainTerminateNotify_proto_enumTypes,
		MessageInfos:      file_BargainTerminateNotify_proto_msgTypes,
	}.Build()
	File_BargainTerminateNotify_proto = out.File
	file_BargainTerminateNotify_proto_rawDesc = nil
	file_BargainTerminateNotify_proto_goTypes = nil
	file_BargainTerminateNotify_proto_depIdxs = nil
}
