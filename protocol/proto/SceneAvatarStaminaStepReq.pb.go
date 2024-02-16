// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SceneAvatarStaminaStepReq.proto

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

type SceneAvatarStaminaStepReq_CmdId int32

const (
	SceneAvatarStaminaStepReq_NONE             SceneAvatarStaminaStepReq_CmdId = 0
	SceneAvatarStaminaStepReq_CMD_ID           SceneAvatarStaminaStepReq_CmdId = 210
	SceneAvatarStaminaStepReq_ENET_CHANNEL_ID  SceneAvatarStaminaStepReq_CmdId = 1
	SceneAvatarStaminaStepReq_ENET_IS_RELIABLE SceneAvatarStaminaStepReq_CmdId = 1
	SceneAvatarStaminaStepReq_IS_ALLOW_CLIENT  SceneAvatarStaminaStepReq_CmdId = 1
)

// Enum value maps for SceneAvatarStaminaStepReq_CmdId.
var (
	SceneAvatarStaminaStepReq_CmdId_name = map[int32]string{
		0:   "NONE",
		210: "CMD_ID",
		1:   "ENET_CHANNEL_ID",
		// Duplicate value: 1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	SceneAvatarStaminaStepReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           210,
		"ENET_CHANNEL_ID":  1,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x SceneAvatarStaminaStepReq_CmdId) Enum() *SceneAvatarStaminaStepReq_CmdId {
	p := new(SceneAvatarStaminaStepReq_CmdId)
	*p = x
	return p
}

func (x SceneAvatarStaminaStepReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SceneAvatarStaminaStepReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_SceneAvatarStaminaStepReq_proto_enumTypes[0].Descriptor()
}

func (SceneAvatarStaminaStepReq_CmdId) Type() protoreflect.EnumType {
	return &file_SceneAvatarStaminaStepReq_proto_enumTypes[0]
}

func (x SceneAvatarStaminaStepReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SceneAvatarStaminaStepReq_CmdId.Descriptor instead.
func (SceneAvatarStaminaStepReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_SceneAvatarStaminaStepReq_proto_rawDescGZIP(), []int{0, 0}
}

type SceneAvatarStaminaStepReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UseClientRot bool    `protobuf:"varint,1,opt,name=use_client_rot,json=useClientRot,proto3" json:"use_client_rot,omitempty"`
	Rot          *Vector `protobuf:"bytes,2,opt,name=rot,proto3" json:"rot,omitempty"`
}

func (x *SceneAvatarStaminaStepReq) Reset() {
	*x = SceneAvatarStaminaStepReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneAvatarStaminaStepReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneAvatarStaminaStepReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneAvatarStaminaStepReq) ProtoMessage() {}

func (x *SceneAvatarStaminaStepReq) ProtoReflect() protoreflect.Message {
	mi := &file_SceneAvatarStaminaStepReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneAvatarStaminaStepReq.ProtoReflect.Descriptor instead.
func (*SceneAvatarStaminaStepReq) Descriptor() ([]byte, []int) {
	return file_SceneAvatarStaminaStepReq_proto_rawDescGZIP(), []int{0}
}

func (x *SceneAvatarStaminaStepReq) GetUseClientRot() bool {
	if x != nil {
		return x.UseClientRot
	}
	return false
}

func (x *SceneAvatarStaminaStepReq) GetRot() *Vector {
	if x != nil {
		return x.Rot
	}
	return nil
}

var File_SceneAvatarStaminaStepReq_proto protoreflect.FileDescriptor

var file_SceneAvatarStaminaStepReq_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x74, 0x61,
	0x6d, 0x69, 0x6e, 0x61, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x19, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x53, 0x74, 0x65,
	0x70, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x75, 0x73,
	0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x74, 0x12, 0x1f, 0x0a, 0x03, 0x72, 0x6f,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x72, 0x6f, 0x74, 0x22, 0x62, 0x0a, 0x05, 0x43,
	0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xd2, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x01,
	0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c,
	0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_SceneAvatarStaminaStepReq_proto_rawDescOnce sync.Once
	file_SceneAvatarStaminaStepReq_proto_rawDescData = file_SceneAvatarStaminaStepReq_proto_rawDesc
)

func file_SceneAvatarStaminaStepReq_proto_rawDescGZIP() []byte {
	file_SceneAvatarStaminaStepReq_proto_rawDescOnce.Do(func() {
		file_SceneAvatarStaminaStepReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneAvatarStaminaStepReq_proto_rawDescData)
	})
	return file_SceneAvatarStaminaStepReq_proto_rawDescData
}

var file_SceneAvatarStaminaStepReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SceneAvatarStaminaStepReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneAvatarStaminaStepReq_proto_goTypes = []interface{}{
	(SceneAvatarStaminaStepReq_CmdId)(0), // 0: proto.SceneAvatarStaminaStepReq.CmdId
	(*SceneAvatarStaminaStepReq)(nil),    // 1: proto.SceneAvatarStaminaStepReq
	(*Vector)(nil),                       // 2: proto.Vector
}
var file_SceneAvatarStaminaStepReq_proto_depIdxs = []int32{
	2, // 0: proto.SceneAvatarStaminaStepReq.rot:type_name -> proto.Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SceneAvatarStaminaStepReq_proto_init() }
func file_SceneAvatarStaminaStepReq_proto_init() {
	if File_SceneAvatarStaminaStepReq_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneAvatarStaminaStepReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneAvatarStaminaStepReq); i {
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
			RawDescriptor: file_SceneAvatarStaminaStepReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneAvatarStaminaStepReq_proto_goTypes,
		DependencyIndexes: file_SceneAvatarStaminaStepReq_proto_depIdxs,
		EnumInfos:         file_SceneAvatarStaminaStepReq_proto_enumTypes,
		MessageInfos:      file_SceneAvatarStaminaStepReq_proto_msgTypes,
	}.Build()
	File_SceneAvatarStaminaStepReq_proto = out.File
	file_SceneAvatarStaminaStepReq_proto_rawDesc = nil
	file_SceneAvatarStaminaStepReq_proto_goTypes = nil
	file_SceneAvatarStaminaStepReq_proto_depIdxs = nil
}