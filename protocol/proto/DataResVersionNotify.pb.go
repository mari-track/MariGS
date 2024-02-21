// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: DataResVersionNotify.proto

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

type DataResVersionNotify_CmdId int32

const (
	DataResVersionNotify_NONE             DataResVersionNotify_CmdId = 0
	DataResVersionNotify_CMD_ID           DataResVersionNotify_CmdId = 145
	DataResVersionNotify_ENET_CHANNEL_ID  DataResVersionNotify_CmdId = 0
	DataResVersionNotify_ENET_IS_RELIABLE DataResVersionNotify_CmdId = 1
)

// Enum value maps for DataResVersionNotify_CmdId.
var (
	DataResVersionNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		145: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	DataResVersionNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           145,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x DataResVersionNotify_CmdId) Enum() *DataResVersionNotify_CmdId {
	p := new(DataResVersionNotify_CmdId)
	*p = x
	return p
}

func (x DataResVersionNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataResVersionNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_DataResVersionNotify_proto_enumTypes[0].Descriptor()
}

func (DataResVersionNotify_CmdId) Type() protoreflect.EnumType {
	return &file_DataResVersionNotify_proto_enumTypes[0]
}

func (x DataResVersionNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataResVersionNotify_CmdId.Descriptor instead.
func (DataResVersionNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_DataResVersionNotify_proto_rawDescGZIP(), []int{0, 0}
}

type DataResVersionNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientDataVersion          uint32            `protobuf:"varint,1,opt,name=client_data_version,json=clientDataVersion,proto3" json:"client_data_version,omitempty"`
	ClientSilenceDataVersion   uint32            `protobuf:"varint,2,opt,name=client_silence_data_version,json=clientSilenceDataVersion,proto3" json:"client_silence_data_version,omitempty"`
	ClientMd5                  string            `protobuf:"bytes,5,opt,name=client_md5,json=clientMd5,proto3" json:"client_md5,omitempty"`
	ClientSilenceMd5           string            `protobuf:"bytes,6,opt,name=client_silence_md5,json=clientSilenceMd5,proto3" json:"client_silence_md5,omitempty"`
	ResVersionConfig           *ResVersionConfig `protobuf:"bytes,10,opt,name=res_version_config,json=resVersionConfig,proto3" json:"res_version_config,omitempty"`
	ClientVersionSuffix        string            `protobuf:"bytes,11,opt,name=client_version_suffix,json=clientVersionSuffix,proto3" json:"client_version_suffix,omitempty"`
	ClientSilenceVersionSuffix string            `protobuf:"bytes,12,opt,name=client_silence_version_suffix,json=clientSilenceVersionSuffix,proto3" json:"client_silence_version_suffix,omitempty"`
}

func (x *DataResVersionNotify) Reset() {
	*x = DataResVersionNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DataResVersionNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataResVersionNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataResVersionNotify) ProtoMessage() {}

func (x *DataResVersionNotify) ProtoReflect() protoreflect.Message {
	mi := &file_DataResVersionNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataResVersionNotify.ProtoReflect.Descriptor instead.
func (*DataResVersionNotify) Descriptor() ([]byte, []int) {
	return file_DataResVersionNotify_proto_rawDescGZIP(), []int{0}
}

func (x *DataResVersionNotify) GetClientDataVersion() uint32 {
	if x != nil {
		return x.ClientDataVersion
	}
	return 0
}

func (x *DataResVersionNotify) GetClientSilenceDataVersion() uint32 {
	if x != nil {
		return x.ClientSilenceDataVersion
	}
	return 0
}

func (x *DataResVersionNotify) GetClientMd5() string {
	if x != nil {
		return x.ClientMd5
	}
	return ""
}

func (x *DataResVersionNotify) GetClientSilenceMd5() string {
	if x != nil {
		return x.ClientSilenceMd5
	}
	return ""
}

func (x *DataResVersionNotify) GetResVersionConfig() *ResVersionConfig {
	if x != nil {
		return x.ResVersionConfig
	}
	return nil
}

func (x *DataResVersionNotify) GetClientVersionSuffix() string {
	if x != nil {
		return x.ClientVersionSuffix
	}
	return ""
}

func (x *DataResVersionNotify) GetClientSilenceVersionSuffix() string {
	if x != nil {
		return x.ClientSilenceVersionSuffix
	}
	return ""
}

var File_DataResVersionNotify_proto protoreflect.FileDescriptor

var file_DataResVersionNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x52, 0x65,
	0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x03, 0x0a, 0x14, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2e, 0x0a,
	0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a,
	0x1b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x64, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x64, 0x35, 0x12, 0x2c, 0x0a, 0x12, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x64,
	0x35, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x64, 0x35, 0x12, 0x3f, 0x0a, 0x12, 0x72, 0x65, 0x73,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x72, 0x65, 0x73, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x66,
	0x66, 0x69, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x41,
	0x0a, 0x1d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6c,
	0x65, 0x6e, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x66, 0x66, 0x69,
	0x78, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x91,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45,
	0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49,
	0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DataResVersionNotify_proto_rawDescOnce sync.Once
	file_DataResVersionNotify_proto_rawDescData = file_DataResVersionNotify_proto_rawDesc
)

func file_DataResVersionNotify_proto_rawDescGZIP() []byte {
	file_DataResVersionNotify_proto_rawDescOnce.Do(func() {
		file_DataResVersionNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_DataResVersionNotify_proto_rawDescData)
	})
	return file_DataResVersionNotify_proto_rawDescData
}

var file_DataResVersionNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DataResVersionNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DataResVersionNotify_proto_goTypes = []interface{}{
	(DataResVersionNotify_CmdId)(0), // 0: DataResVersionNotify.CmdId
	(*DataResVersionNotify)(nil),    // 1: DataResVersionNotify
	(*ResVersionConfig)(nil),        // 2: ResVersionConfig
}
var file_DataResVersionNotify_proto_depIdxs = []int32{
	2, // 0: DataResVersionNotify.res_version_config:type_name -> ResVersionConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DataResVersionNotify_proto_init() }
func file_DataResVersionNotify_proto_init() {
	if File_DataResVersionNotify_proto != nil {
		return
	}
	file_ResVersionConfig_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DataResVersionNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataResVersionNotify); i {
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
			RawDescriptor: file_DataResVersionNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DataResVersionNotify_proto_goTypes,
		DependencyIndexes: file_DataResVersionNotify_proto_depIdxs,
		EnumInfos:         file_DataResVersionNotify_proto_enumTypes,
		MessageInfos:      file_DataResVersionNotify_proto_msgTypes,
	}.Build()
	File_DataResVersionNotify_proto = out.File
	file_DataResVersionNotify_proto_rawDesc = nil
	file_DataResVersionNotify_proto_goTypes = nil
	file_DataResVersionNotify_proto_depIdxs = nil
}
