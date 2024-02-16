// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: HitClientTrivialNotify.proto

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

type HitClientTrivialNotify_CmdId int32

const (
	HitClientTrivialNotify_NONE             HitClientTrivialNotify_CmdId = 0
	HitClientTrivialNotify_CMD_ID           HitClientTrivialNotify_CmdId = 274
	HitClientTrivialNotify_ENET_CHANNEL_ID  HitClientTrivialNotify_CmdId = 0
	HitClientTrivialNotify_ENET_IS_RELIABLE HitClientTrivialNotify_CmdId = 1
	HitClientTrivialNotify_IS_ALLOW_CLIENT  HitClientTrivialNotify_CmdId = 1
)

// Enum value maps for HitClientTrivialNotify_CmdId.
var (
	HitClientTrivialNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		274: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	HitClientTrivialNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           274,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x HitClientTrivialNotify_CmdId) Enum() *HitClientTrivialNotify_CmdId {
	p := new(HitClientTrivialNotify_CmdId)
	*p = x
	return p
}

func (x HitClientTrivialNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HitClientTrivialNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_HitClientTrivialNotify_proto_enumTypes[0].Descriptor()
}

func (HitClientTrivialNotify_CmdId) Type() protoreflect.EnumType {
	return &file_HitClientTrivialNotify_proto_enumTypes[0]
}

func (x HitClientTrivialNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HitClientTrivialNotify_CmdId.Descriptor instead.
func (HitClientTrivialNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_HitClientTrivialNotify_proto_rawDescGZIP(), []int{0, 0}
}

type HitClientTrivialNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position      *Vector `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	OwnerEntityId uint32  `protobuf:"varint,2,opt,name=owner_entity_id,json=ownerEntityId,proto3" json:"owner_entity_id,omitempty"`
}

func (x *HitClientTrivialNotify) Reset() {
	*x = HitClientTrivialNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HitClientTrivialNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HitClientTrivialNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HitClientTrivialNotify) ProtoMessage() {}

func (x *HitClientTrivialNotify) ProtoReflect() protoreflect.Message {
	mi := &file_HitClientTrivialNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HitClientTrivialNotify.ProtoReflect.Descriptor instead.
func (*HitClientTrivialNotify) Descriptor() ([]byte, []int) {
	return file_HitClientTrivialNotify_proto_rawDescGZIP(), []int{0}
}

func (x *HitClientTrivialNotify) GetPosition() *Vector {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *HitClientTrivialNotify) GetOwnerEntityId() uint32 {
	if x != nil {
		return x.OwnerEntityId
	}
	return 0
}

var File_HitClientTrivialNotify_proto protoreflect.FileDescriptor

var file_HitClientTrivialNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x48, 0x69, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x69, 0x76, 0x69,
	0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x16, 0x48, 0x69, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x72, 0x69, 0x76, 0x69, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x29,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x92,
	0x02, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45,
	0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49,
	0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10,
	0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HitClientTrivialNotify_proto_rawDescOnce sync.Once
	file_HitClientTrivialNotify_proto_rawDescData = file_HitClientTrivialNotify_proto_rawDesc
)

func file_HitClientTrivialNotify_proto_rawDescGZIP() []byte {
	file_HitClientTrivialNotify_proto_rawDescOnce.Do(func() {
		file_HitClientTrivialNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_HitClientTrivialNotify_proto_rawDescData)
	})
	return file_HitClientTrivialNotify_proto_rawDescData
}

var file_HitClientTrivialNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_HitClientTrivialNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HitClientTrivialNotify_proto_goTypes = []interface{}{
	(HitClientTrivialNotify_CmdId)(0), // 0: proto.HitClientTrivialNotify.CmdId
	(*HitClientTrivialNotify)(nil),    // 1: proto.HitClientTrivialNotify
	(*Vector)(nil),                    // 2: proto.Vector
}
var file_HitClientTrivialNotify_proto_depIdxs = []int32{
	2, // 0: proto.HitClientTrivialNotify.position:type_name -> proto.Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HitClientTrivialNotify_proto_init() }
func file_HitClientTrivialNotify_proto_init() {
	if File_HitClientTrivialNotify_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HitClientTrivialNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HitClientTrivialNotify); i {
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
			RawDescriptor: file_HitClientTrivialNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HitClientTrivialNotify_proto_goTypes,
		DependencyIndexes: file_HitClientTrivialNotify_proto_depIdxs,
		EnumInfos:         file_HitClientTrivialNotify_proto_enumTypes,
		MessageInfos:      file_HitClientTrivialNotify_proto_msgTypes,
	}.Build()
	File_HitClientTrivialNotify_proto = out.File
	file_HitClientTrivialNotify_proto_rawDesc = nil
	file_HitClientTrivialNotify_proto_goTypes = nil
	file_HitClientTrivialNotify_proto_depIdxs = nil
}