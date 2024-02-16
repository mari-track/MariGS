// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: EvtDoSkillSuccNotify.proto

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

type EvtDoSkillSuccNotify_CmdId int32

const (
	EvtDoSkillSuccNotify_NONE             EvtDoSkillSuccNotify_CmdId = 0
	EvtDoSkillSuccNotify_CMD_ID           EvtDoSkillSuccNotify_CmdId = 304
	EvtDoSkillSuccNotify_ENET_CHANNEL_ID  EvtDoSkillSuccNotify_CmdId = 0
	EvtDoSkillSuccNotify_ENET_IS_RELIABLE EvtDoSkillSuccNotify_CmdId = 1
	EvtDoSkillSuccNotify_IS_ALLOW_CLIENT  EvtDoSkillSuccNotify_CmdId = 1
)

// Enum value maps for EvtDoSkillSuccNotify_CmdId.
var (
	EvtDoSkillSuccNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		304: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	EvtDoSkillSuccNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           304,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x EvtDoSkillSuccNotify_CmdId) Enum() *EvtDoSkillSuccNotify_CmdId {
	p := new(EvtDoSkillSuccNotify_CmdId)
	*p = x
	return p
}

func (x EvtDoSkillSuccNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EvtDoSkillSuccNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_EvtDoSkillSuccNotify_proto_enumTypes[0].Descriptor()
}

func (EvtDoSkillSuccNotify_CmdId) Type() protoreflect.EnumType {
	return &file_EvtDoSkillSuccNotify_proto_enumTypes[0]
}

func (x EvtDoSkillSuccNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EvtDoSkillSuccNotify_CmdId.Descriptor instead.
func (EvtDoSkillSuccNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_EvtDoSkillSuccNotify_proto_rawDescGZIP(), []int{0, 0}
}

type EvtDoSkillSuccNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForwardType ForwardType `protobuf:"varint,1,opt,name=forward_type,json=forwardType,proto3,enum=proto.ForwardType" json:"forward_type,omitempty"`
	CasterId    uint32      `protobuf:"varint,2,opt,name=caster_id,json=casterId,proto3" json:"caster_id,omitempty"`
	SkillId     uint32      `protobuf:"varint,3,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	Forward     *Vector     `protobuf:"bytes,4,opt,name=forward,proto3" json:"forward,omitempty"`
}

func (x *EvtDoSkillSuccNotify) Reset() {
	*x = EvtDoSkillSuccNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtDoSkillSuccNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtDoSkillSuccNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtDoSkillSuccNotify) ProtoMessage() {}

func (x *EvtDoSkillSuccNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtDoSkillSuccNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtDoSkillSuccNotify.ProtoReflect.Descriptor instead.
func (*EvtDoSkillSuccNotify) Descriptor() ([]byte, []int) {
	return file_EvtDoSkillSuccNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtDoSkillSuccNotify) GetForwardType() ForwardType {
	if x != nil {
		return x.ForwardType
	}
	return ForwardType_FORWARD_LOCAL
}

func (x *EvtDoSkillSuccNotify) GetCasterId() uint32 {
	if x != nil {
		return x.CasterId
	}
	return 0
}

func (x *EvtDoSkillSuccNotify) GetSkillId() uint32 {
	if x != nil {
		return x.SkillId
	}
	return 0
}

func (x *EvtDoSkillSuccNotify) GetForward() *Vector {
	if x != nil {
		return x.Forward
	}
	return nil
}

var File_EvtDoSkillSuccNotify_proto protoreflect.FileDescriptor

var file_EvtDoSkillSuccNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x45, 0x76, 0x74, 0x44, 0x6f, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x53, 0x75, 0x63, 0x63,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x02, 0x0a, 0x14, 0x45, 0x76, 0x74, 0x44, 0x6f, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x53, 0x75, 0x63, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x35, 0x0a,
	0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x07,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f,
	0x49, 0x44, 0x10, 0xb0, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtDoSkillSuccNotify_proto_rawDescOnce sync.Once
	file_EvtDoSkillSuccNotify_proto_rawDescData = file_EvtDoSkillSuccNotify_proto_rawDesc
)

func file_EvtDoSkillSuccNotify_proto_rawDescGZIP() []byte {
	file_EvtDoSkillSuccNotify_proto_rawDescOnce.Do(func() {
		file_EvtDoSkillSuccNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtDoSkillSuccNotify_proto_rawDescData)
	})
	return file_EvtDoSkillSuccNotify_proto_rawDescData
}

var file_EvtDoSkillSuccNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EvtDoSkillSuccNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtDoSkillSuccNotify_proto_goTypes = []interface{}{
	(EvtDoSkillSuccNotify_CmdId)(0), // 0: proto.EvtDoSkillSuccNotify.CmdId
	(*EvtDoSkillSuccNotify)(nil),    // 1: proto.EvtDoSkillSuccNotify
	(ForwardType)(0),                // 2: proto.ForwardType
	(*Vector)(nil),                  // 3: proto.Vector
}
var file_EvtDoSkillSuccNotify_proto_depIdxs = []int32{
	2, // 0: proto.EvtDoSkillSuccNotify.forward_type:type_name -> proto.ForwardType
	3, // 1: proto.EvtDoSkillSuccNotify.forward:type_name -> proto.Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EvtDoSkillSuccNotify_proto_init() }
func file_EvtDoSkillSuccNotify_proto_init() {
	if File_EvtDoSkillSuccNotify_proto != nil {
		return
	}
	file_Vector_proto_init()
	file_ForwardType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtDoSkillSuccNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtDoSkillSuccNotify); i {
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
			RawDescriptor: file_EvtDoSkillSuccNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtDoSkillSuccNotify_proto_goTypes,
		DependencyIndexes: file_EvtDoSkillSuccNotify_proto_depIdxs,
		EnumInfos:         file_EvtDoSkillSuccNotify_proto_enumTypes,
		MessageInfos:      file_EvtDoSkillSuccNotify_proto_msgTypes,
	}.Build()
	File_EvtDoSkillSuccNotify_proto = out.File
	file_EvtDoSkillSuccNotify_proto_rawDesc = nil
	file_EvtDoSkillSuccNotify_proto_goTypes = nil
	file_EvtDoSkillSuccNotify_proto_depIdxs = nil
}
