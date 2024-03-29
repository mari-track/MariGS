// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: WearEquipReq.proto

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

type WearEquipReq_CmdId int32

const (
	WearEquipReq_NONE             WearEquipReq_CmdId = 0
	WearEquipReq_CMD_ID           WearEquipReq_CmdId = 614
	WearEquipReq_ENET_CHANNEL_ID  WearEquipReq_CmdId = 0
	WearEquipReq_ENET_IS_RELIABLE WearEquipReq_CmdId = 1
	WearEquipReq_IS_ALLOW_CLIENT  WearEquipReq_CmdId = 1
)

// Enum value maps for WearEquipReq_CmdId.
var (
	WearEquipReq_CmdId_name = map[int32]string{
		0:   "NONE",
		614: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	WearEquipReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           614,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x WearEquipReq_CmdId) Enum() *WearEquipReq_CmdId {
	p := new(WearEquipReq_CmdId)
	*p = x
	return p
}

func (x WearEquipReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WearEquipReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_WearEquipReq_proto_enumTypes[0].Descriptor()
}

func (WearEquipReq_CmdId) Type() protoreflect.EnumType {
	return &file_WearEquipReq_proto_enumTypes[0]
}

func (x WearEquipReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WearEquipReq_CmdId.Descriptor instead.
func (WearEquipReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_WearEquipReq_proto_rawDescGZIP(), []int{0, 0}
}

type WearEquipReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarGuid uint64 `protobuf:"varint,1,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	EquipGuid  uint64 `protobuf:"varint,2,opt,name=equip_guid,json=equipGuid,proto3" json:"equip_guid,omitempty"`
}

func (x *WearEquipReq) Reset() {
	*x = WearEquipReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WearEquipReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WearEquipReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WearEquipReq) ProtoMessage() {}

func (x *WearEquipReq) ProtoReflect() protoreflect.Message {
	mi := &file_WearEquipReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WearEquipReq.ProtoReflect.Descriptor instead.
func (*WearEquipReq) Descriptor() ([]byte, []int) {
	return file_WearEquipReq_proto_rawDescGZIP(), []int{0}
}

func (x *WearEquipReq) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *WearEquipReq) GetEquipGuid() uint64 {
	if x != nil {
		return x.EquipGuid
	}
	return 0
}

var File_WearEquipReq_proto protoreflect.FileDescriptor

var file_WearEquipReq_proto_rawDesc = []byte{
	0x0a, 0x12, 0x57, 0x65, 0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x52, 0x65, 0x71, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x0c, 0x57, 0x65, 0x61, 0x72, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x71, 0x75, 0x69, 0x70, 0x5f,
	0x67, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x65, 0x71, 0x75, 0x69,
	0x70, 0x47, 0x75, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f,
	0x49, 0x44, 0x10, 0xe6, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WearEquipReq_proto_rawDescOnce sync.Once
	file_WearEquipReq_proto_rawDescData = file_WearEquipReq_proto_rawDesc
)

func file_WearEquipReq_proto_rawDescGZIP() []byte {
	file_WearEquipReq_proto_rawDescOnce.Do(func() {
		file_WearEquipReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_WearEquipReq_proto_rawDescData)
	})
	return file_WearEquipReq_proto_rawDescData
}

var file_WearEquipReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_WearEquipReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WearEquipReq_proto_goTypes = []interface{}{
	(WearEquipReq_CmdId)(0), // 0: WearEquipReq.CmdId
	(*WearEquipReq)(nil),    // 1: WearEquipReq
}
var file_WearEquipReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_WearEquipReq_proto_init() }
func file_WearEquipReq_proto_init() {
	if File_WearEquipReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_WearEquipReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WearEquipReq); i {
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
			RawDescriptor: file_WearEquipReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WearEquipReq_proto_goTypes,
		DependencyIndexes: file_WearEquipReq_proto_depIdxs,
		EnumInfos:         file_WearEquipReq_proto_enumTypes,
		MessageInfos:      file_WearEquipReq_proto_msgTypes,
	}.Build()
	File_WearEquipReq_proto = out.File
	file_WearEquipReq_proto_rawDesc = nil
	file_WearEquipReq_proto_goTypes = nil
	file_WearEquipReq_proto_depIdxs = nil
}
