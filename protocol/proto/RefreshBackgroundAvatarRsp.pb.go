// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: RefreshBackgroundAvatarRsp.proto

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

type RefreshBackgroundAvatarRsp_CmdId int32

const (
	RefreshBackgroundAvatarRsp_NONE             RefreshBackgroundAvatarRsp_CmdId = 0
	RefreshBackgroundAvatarRsp_CMD_ID           RefreshBackgroundAvatarRsp_CmdId = 1714
	RefreshBackgroundAvatarRsp_ENET_CHANNEL_ID  RefreshBackgroundAvatarRsp_CmdId = 0
	RefreshBackgroundAvatarRsp_ENET_IS_RELIABLE RefreshBackgroundAvatarRsp_CmdId = 1
)

// Enum value maps for RefreshBackgroundAvatarRsp_CmdId.
var (
	RefreshBackgroundAvatarRsp_CmdId_name = map[int32]string{
		0:    "NONE",
		1714: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	RefreshBackgroundAvatarRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1714,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x RefreshBackgroundAvatarRsp_CmdId) Enum() *RefreshBackgroundAvatarRsp_CmdId {
	p := new(RefreshBackgroundAvatarRsp_CmdId)
	*p = x
	return p
}

func (x RefreshBackgroundAvatarRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RefreshBackgroundAvatarRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_RefreshBackgroundAvatarRsp_proto_enumTypes[0].Descriptor()
}

func (RefreshBackgroundAvatarRsp_CmdId) Type() protoreflect.EnumType {
	return &file_RefreshBackgroundAvatarRsp_proto_enumTypes[0]
}

func (x RefreshBackgroundAvatarRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RefreshBackgroundAvatarRsp_CmdId.Descriptor instead.
func (RefreshBackgroundAvatarRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_RefreshBackgroundAvatarRsp_proto_rawDescGZIP(), []int{0, 0}
}

type RefreshBackgroundAvatarRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode       int32             `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	HpFullTimeMap map[uint64]uint32 `protobuf:"bytes,2,rep,name=hp_full_time_map,json=hpFullTimeMap,proto3" json:"hp_full_time_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *RefreshBackgroundAvatarRsp) Reset() {
	*x = RefreshBackgroundAvatarRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RefreshBackgroundAvatarRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshBackgroundAvatarRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshBackgroundAvatarRsp) ProtoMessage() {}

func (x *RefreshBackgroundAvatarRsp) ProtoReflect() protoreflect.Message {
	mi := &file_RefreshBackgroundAvatarRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshBackgroundAvatarRsp.ProtoReflect.Descriptor instead.
func (*RefreshBackgroundAvatarRsp) Descriptor() ([]byte, []int) {
	return file_RefreshBackgroundAvatarRsp_proto_rawDescGZIP(), []int{0}
}

func (x *RefreshBackgroundAvatarRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *RefreshBackgroundAvatarRsp) GetHpFullTimeMap() map[uint64]uint32 {
	if x != nil {
		return x.HpFullTimeMap
	}
	return nil
}

var File_RefreshBackgroundAvatarRsp_proto protoreflect.FileDescriptor

var file_RefreshBackgroundAvatarRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa0, 0x02, 0x0a, 0x1a, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x42, 0x61,
	0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x68,
	0x70, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x42,
	0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52,
	0x73, 0x70, 0x2e, 0x48, 0x70, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x68, 0x70, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x4d, 0x61, 0x70, 0x1a, 0x40, 0x0a, 0x12, 0x48, 0x70, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44,
	0x5f, 0x49, 0x44, 0x10, 0xb2, 0x0d, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RefreshBackgroundAvatarRsp_proto_rawDescOnce sync.Once
	file_RefreshBackgroundAvatarRsp_proto_rawDescData = file_RefreshBackgroundAvatarRsp_proto_rawDesc
)

func file_RefreshBackgroundAvatarRsp_proto_rawDescGZIP() []byte {
	file_RefreshBackgroundAvatarRsp_proto_rawDescOnce.Do(func() {
		file_RefreshBackgroundAvatarRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_RefreshBackgroundAvatarRsp_proto_rawDescData)
	})
	return file_RefreshBackgroundAvatarRsp_proto_rawDescData
}

var file_RefreshBackgroundAvatarRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RefreshBackgroundAvatarRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_RefreshBackgroundAvatarRsp_proto_goTypes = []interface{}{
	(RefreshBackgroundAvatarRsp_CmdId)(0), // 0: RefreshBackgroundAvatarRsp.CmdId
	(*RefreshBackgroundAvatarRsp)(nil),    // 1: RefreshBackgroundAvatarRsp
	nil,                                   // 2: RefreshBackgroundAvatarRsp.HpFullTimeMapEntry
}
var file_RefreshBackgroundAvatarRsp_proto_depIdxs = []int32{
	2, // 0: RefreshBackgroundAvatarRsp.hp_full_time_map:type_name -> RefreshBackgroundAvatarRsp.HpFullTimeMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RefreshBackgroundAvatarRsp_proto_init() }
func file_RefreshBackgroundAvatarRsp_proto_init() {
	if File_RefreshBackgroundAvatarRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RefreshBackgroundAvatarRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshBackgroundAvatarRsp); i {
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
			RawDescriptor: file_RefreshBackgroundAvatarRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RefreshBackgroundAvatarRsp_proto_goTypes,
		DependencyIndexes: file_RefreshBackgroundAvatarRsp_proto_depIdxs,
		EnumInfos:         file_RefreshBackgroundAvatarRsp_proto_enumTypes,
		MessageInfos:      file_RefreshBackgroundAvatarRsp_proto_msgTypes,
	}.Build()
	File_RefreshBackgroundAvatarRsp_proto = out.File
	file_RefreshBackgroundAvatarRsp_proto_rawDesc = nil
	file_RefreshBackgroundAvatarRsp_proto_goTypes = nil
	file_RefreshBackgroundAvatarRsp_proto_depIdxs = nil
}
