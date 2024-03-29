// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: ResinChangeNotify.proto

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

type ResinChangeNotify_CmdId int32

const (
	ResinChangeNotify_NONE             ResinChangeNotify_CmdId = 0
	ResinChangeNotify_CMD_ID           ResinChangeNotify_CmdId = 643
	ResinChangeNotify_ENET_CHANNEL_ID  ResinChangeNotify_CmdId = 0
	ResinChangeNotify_ENET_IS_RELIABLE ResinChangeNotify_CmdId = 1
)

// Enum value maps for ResinChangeNotify_CmdId.
var (
	ResinChangeNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		643: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	ResinChangeNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           643,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x ResinChangeNotify_CmdId) Enum() *ResinChangeNotify_CmdId {
	p := new(ResinChangeNotify_CmdId)
	*p = x
	return p
}

func (x ResinChangeNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResinChangeNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ResinChangeNotify_proto_enumTypes[0].Descriptor()
}

func (ResinChangeNotify_CmdId) Type() protoreflect.EnumType {
	return &file_ResinChangeNotify_proto_enumTypes[0]
}

func (x ResinChangeNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResinChangeNotify_CmdId.Descriptor instead.
func (ResinChangeNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ResinChangeNotify_proto_rawDescGZIP(), []int{0, 0}
}

type ResinChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurValue         uint32 `protobuf:"varint,1,opt,name=cur_value,json=curValue,proto3" json:"cur_value,omitempty"`
	NextAddTimestamp uint32 `protobuf:"varint,2,opt,name=next_add_timestamp,json=nextAddTimestamp,proto3" json:"next_add_timestamp,omitempty"`
	CurBuyCount      uint32 `protobuf:"varint,3,opt,name=cur_buy_count,json=curBuyCount,proto3" json:"cur_buy_count,omitempty"`
}

func (x *ResinChangeNotify) Reset() {
	*x = ResinChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ResinChangeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResinChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResinChangeNotify) ProtoMessage() {}

func (x *ResinChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ResinChangeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResinChangeNotify.ProtoReflect.Descriptor instead.
func (*ResinChangeNotify) Descriptor() ([]byte, []int) {
	return file_ResinChangeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ResinChangeNotify) GetCurValue() uint32 {
	if x != nil {
		return x.CurValue
	}
	return 0
}

func (x *ResinChangeNotify) GetNextAddTimestamp() uint32 {
	if x != nil {
		return x.NextAddTimestamp
	}
	return 0
}

func (x *ResinChangeNotify) GetCurBuyCount() uint32 {
	if x != nil {
		return x.CurBuyCount
	}
	return 0
}

var File_ResinChangeNotify_proto protoreflect.FileDescriptor

var file_ResinChangeNotify_proto_rawDesc = []byte{
	0x0a, 0x17, 0x52, 0x65, 0x73, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x11, 0x52, 0x65,
	0x73, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x63, 0x75, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a, 0x12,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6e, 0x65, 0x78, 0x74, 0x41, 0x64,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x75,
	0x72, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x42, 0x75, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4d,
	0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x83, 0x05, 0x12, 0x13,
	0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52,
	0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ResinChangeNotify_proto_rawDescOnce sync.Once
	file_ResinChangeNotify_proto_rawDescData = file_ResinChangeNotify_proto_rawDesc
)

func file_ResinChangeNotify_proto_rawDescGZIP() []byte {
	file_ResinChangeNotify_proto_rawDescOnce.Do(func() {
		file_ResinChangeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ResinChangeNotify_proto_rawDescData)
	})
	return file_ResinChangeNotify_proto_rawDescData
}

var file_ResinChangeNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ResinChangeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ResinChangeNotify_proto_goTypes = []interface{}{
	(ResinChangeNotify_CmdId)(0), // 0: ResinChangeNotify.CmdId
	(*ResinChangeNotify)(nil),    // 1: ResinChangeNotify
}
var file_ResinChangeNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ResinChangeNotify_proto_init() }
func file_ResinChangeNotify_proto_init() {
	if File_ResinChangeNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ResinChangeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResinChangeNotify); i {
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
			RawDescriptor: file_ResinChangeNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ResinChangeNotify_proto_goTypes,
		DependencyIndexes: file_ResinChangeNotify_proto_depIdxs,
		EnumInfos:         file_ResinChangeNotify_proto_enumTypes,
		MessageInfos:      file_ResinChangeNotify_proto_msgTypes,
	}.Build()
	File_ResinChangeNotify_proto = out.File
	file_ResinChangeNotify_proto_rawDesc = nil
	file_ResinChangeNotify_proto_goTypes = nil
	file_ResinChangeNotify_proto_depIdxs = nil
}
