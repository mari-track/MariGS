// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GetBargainDataReq.proto

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

type GetBargainDataReq_CmdId int32

const (
	GetBargainDataReq_NONE             GetBargainDataReq_CmdId = 0
	GetBargainDataReq_CMD_ID           GetBargainDataReq_CmdId = 427
	GetBargainDataReq_ENET_CHANNEL_ID  GetBargainDataReq_CmdId = 0
	GetBargainDataReq_ENET_IS_RELIABLE GetBargainDataReq_CmdId = 1
	GetBargainDataReq_IS_ALLOW_CLIENT  GetBargainDataReq_CmdId = 1
)

// Enum value maps for GetBargainDataReq_CmdId.
var (
	GetBargainDataReq_CmdId_name = map[int32]string{
		0:   "NONE",
		427: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	GetBargainDataReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           427,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x GetBargainDataReq_CmdId) Enum() *GetBargainDataReq_CmdId {
	p := new(GetBargainDataReq_CmdId)
	*p = x
	return p
}

func (x GetBargainDataReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetBargainDataReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetBargainDataReq_proto_enumTypes[0].Descriptor()
}

func (GetBargainDataReq_CmdId) Type() protoreflect.EnumType {
	return &file_GetBargainDataReq_proto_enumTypes[0]
}

func (x GetBargainDataReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetBargainDataReq_CmdId.Descriptor instead.
func (GetBargainDataReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetBargainDataReq_proto_rawDescGZIP(), []int{0, 0}
}

type GetBargainDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BargainId uint32 `protobuf:"varint,1,opt,name=bargain_id,json=bargainId,proto3" json:"bargain_id,omitempty"`
}

func (x *GetBargainDataReq) Reset() {
	*x = GetBargainDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetBargainDataReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBargainDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBargainDataReq) ProtoMessage() {}

func (x *GetBargainDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetBargainDataReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBargainDataReq.ProtoReflect.Descriptor instead.
func (*GetBargainDataReq) Descriptor() ([]byte, []int) {
	return file_GetBargainDataReq_proto_rawDescGZIP(), []int{0}
}

func (x *GetBargainDataReq) GetBargainId() uint32 {
	if x != nil {
		return x.BargainId
	}
	return 0
}

var File_GetBargainDataReq_proto protoreflect.FileDescriptor

var file_GetBargainDataReq_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x42, 0x61, 0x72, 0x67, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x72, 0x67, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x72, 0x67, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x61, 0x72, 0x67, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x62,
	0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xab, 0x03, 0x12, 0x13,
	0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52,
	0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f,
	0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02,
	0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetBargainDataReq_proto_rawDescOnce sync.Once
	file_GetBargainDataReq_proto_rawDescData = file_GetBargainDataReq_proto_rawDesc
)

func file_GetBargainDataReq_proto_rawDescGZIP() []byte {
	file_GetBargainDataReq_proto_rawDescOnce.Do(func() {
		file_GetBargainDataReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetBargainDataReq_proto_rawDescData)
	})
	return file_GetBargainDataReq_proto_rawDescData
}

var file_GetBargainDataReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetBargainDataReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetBargainDataReq_proto_goTypes = []interface{}{
	(GetBargainDataReq_CmdId)(0), // 0: GetBargainDataReq.CmdId
	(*GetBargainDataReq)(nil),    // 1: GetBargainDataReq
}
var file_GetBargainDataReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetBargainDataReq_proto_init() }
func file_GetBargainDataReq_proto_init() {
	if File_GetBargainDataReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetBargainDataReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBargainDataReq); i {
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
			RawDescriptor: file_GetBargainDataReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetBargainDataReq_proto_goTypes,
		DependencyIndexes: file_GetBargainDataReq_proto_depIdxs,
		EnumInfos:         file_GetBargainDataReq_proto_enumTypes,
		MessageInfos:      file_GetBargainDataReq_proto_msgTypes,
	}.Build()
	File_GetBargainDataReq_proto = out.File
	file_GetBargainDataReq_proto_rawDesc = nil
	file_GetBargainDataReq_proto_goTypes = nil
	file_GetBargainDataReq_proto_depIdxs = nil
}
