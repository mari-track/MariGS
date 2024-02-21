// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: AvatarExpeditionStartRsp.proto

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

type AvatarExpeditionStartRsp_CmdId int32

const (
	AvatarExpeditionStartRsp_NONE             AvatarExpeditionStartRsp_CmdId = 0
	AvatarExpeditionStartRsp_CMD_ID           AvatarExpeditionStartRsp_CmdId = 1728
	AvatarExpeditionStartRsp_ENET_CHANNEL_ID  AvatarExpeditionStartRsp_CmdId = 0
	AvatarExpeditionStartRsp_ENET_IS_RELIABLE AvatarExpeditionStartRsp_CmdId = 1
)

// Enum value maps for AvatarExpeditionStartRsp_CmdId.
var (
	AvatarExpeditionStartRsp_CmdId_name = map[int32]string{
		0:    "NONE",
		1728: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	AvatarExpeditionStartRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1728,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x AvatarExpeditionStartRsp_CmdId) Enum() *AvatarExpeditionStartRsp_CmdId {
	p := new(AvatarExpeditionStartRsp_CmdId)
	*p = x
	return p
}

func (x AvatarExpeditionStartRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AvatarExpeditionStartRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_AvatarExpeditionStartRsp_proto_enumTypes[0].Descriptor()
}

func (AvatarExpeditionStartRsp_CmdId) Type() protoreflect.EnumType {
	return &file_AvatarExpeditionStartRsp_proto_enumTypes[0]
}

func (x AvatarExpeditionStartRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AvatarExpeditionStartRsp_CmdId.Descriptor instead.
func (AvatarExpeditionStartRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_AvatarExpeditionStartRsp_proto_rawDescGZIP(), []int{0, 0}
}

type AvatarExpeditionStartRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode           int32                            `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ExpeditionInfoMap map[uint64]*AvatarExpeditionInfo `protobuf:"bytes,2,rep,name=expedition_info_map,json=expeditionInfoMap,proto3" json:"expedition_info_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AvatarExpeditionStartRsp) Reset() {
	*x = AvatarExpeditionStartRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarExpeditionStartRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarExpeditionStartRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarExpeditionStartRsp) ProtoMessage() {}

func (x *AvatarExpeditionStartRsp) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarExpeditionStartRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarExpeditionStartRsp.ProtoReflect.Descriptor instead.
func (*AvatarExpeditionStartRsp) Descriptor() ([]byte, []int) {
	return file_AvatarExpeditionStartRsp_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarExpeditionStartRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *AvatarExpeditionStartRsp) GetExpeditionInfoMap() map[uint64]*AvatarExpeditionInfo {
	if x != nil {
		return x.ExpeditionInfoMap
	}
	return nil
}

var File_AvatarExpeditionStartRsp_proto protoreflect.FileDescriptor

var file_AvatarExpeditionStartRsp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x02, 0x0a,
	0x18, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x60, 0x0a, 0x13, 0x65, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x73, 0x70, 0x2e, 0x45, 0x78, 0x70, 0x65,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x11, 0x65, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x4d, 0x61, 0x70, 0x1a, 0x5b, 0x0a, 0x16, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0xc0, 0x0d, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e,
	0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarExpeditionStartRsp_proto_rawDescOnce sync.Once
	file_AvatarExpeditionStartRsp_proto_rawDescData = file_AvatarExpeditionStartRsp_proto_rawDesc
)

func file_AvatarExpeditionStartRsp_proto_rawDescGZIP() []byte {
	file_AvatarExpeditionStartRsp_proto_rawDescOnce.Do(func() {
		file_AvatarExpeditionStartRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarExpeditionStartRsp_proto_rawDescData)
	})
	return file_AvatarExpeditionStartRsp_proto_rawDescData
}

var file_AvatarExpeditionStartRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AvatarExpeditionStartRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_AvatarExpeditionStartRsp_proto_goTypes = []interface{}{
	(AvatarExpeditionStartRsp_CmdId)(0), // 0: AvatarExpeditionStartRsp.CmdId
	(*AvatarExpeditionStartRsp)(nil),    // 1: AvatarExpeditionStartRsp
	nil,                                 // 2: AvatarExpeditionStartRsp.ExpeditionInfoMapEntry
	(*AvatarExpeditionInfo)(nil),        // 3: AvatarExpeditionInfo
}
var file_AvatarExpeditionStartRsp_proto_depIdxs = []int32{
	2, // 0: AvatarExpeditionStartRsp.expedition_info_map:type_name -> AvatarExpeditionStartRsp.ExpeditionInfoMapEntry
	3, // 1: AvatarExpeditionStartRsp.ExpeditionInfoMapEntry.value:type_name -> AvatarExpeditionInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AvatarExpeditionStartRsp_proto_init() }
func file_AvatarExpeditionStartRsp_proto_init() {
	if File_AvatarExpeditionStartRsp_proto != nil {
		return
	}
	file_AvatarExpeditionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarExpeditionStartRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarExpeditionStartRsp); i {
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
			RawDescriptor: file_AvatarExpeditionStartRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarExpeditionStartRsp_proto_goTypes,
		DependencyIndexes: file_AvatarExpeditionStartRsp_proto_depIdxs,
		EnumInfos:         file_AvatarExpeditionStartRsp_proto_enumTypes,
		MessageInfos:      file_AvatarExpeditionStartRsp_proto_msgTypes,
	}.Build()
	File_AvatarExpeditionStartRsp_proto = out.File
	file_AvatarExpeditionStartRsp_proto_rawDesc = nil
	file_AvatarExpeditionStartRsp_proto_goTypes = nil
	file_AvatarExpeditionStartRsp_proto_depIdxs = nil
}
