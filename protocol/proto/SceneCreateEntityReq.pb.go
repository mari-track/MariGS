// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SceneCreateEntityReq.proto

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

type SceneCreateEntityReq_CmdId int32

const (
	SceneCreateEntityReq_NONE             SceneCreateEntityReq_CmdId = 0
	SceneCreateEntityReq_CMD_ID           SceneCreateEntityReq_CmdId = 227
	SceneCreateEntityReq_ENET_CHANNEL_ID  SceneCreateEntityReq_CmdId = 0
	SceneCreateEntityReq_ENET_IS_RELIABLE SceneCreateEntityReq_CmdId = 1
	SceneCreateEntityReq_IS_ALLOW_CLIENT  SceneCreateEntityReq_CmdId = 1
)

// Enum value maps for SceneCreateEntityReq_CmdId.
var (
	SceneCreateEntityReq_CmdId_name = map[int32]string{
		0:   "NONE",
		227: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	SceneCreateEntityReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           227,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x SceneCreateEntityReq_CmdId) Enum() *SceneCreateEntityReq_CmdId {
	p := new(SceneCreateEntityReq_CmdId)
	*p = x
	return p
}

func (x SceneCreateEntityReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SceneCreateEntityReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_SceneCreateEntityReq_proto_enumTypes[0].Descriptor()
}

func (SceneCreateEntityReq_CmdId) Type() protoreflect.EnumType {
	return &file_SceneCreateEntityReq_proto_enumTypes[0]
}

func (x SceneCreateEntityReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SceneCreateEntityReq_CmdId.Descriptor instead.
func (SceneCreateEntityReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_SceneCreateEntityReq_proto_rawDescGZIP(), []int{0, 0}
}

type SceneCreateEntityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity                  *CreateEntityInfo `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Reason                  CreateReason      `protobuf:"varint,2,opt,name=reason,proto3,enum=CreateReason" json:"reason,omitempty"`
	IsDestroyWhenDisconnect bool              `protobuf:"varint,3,opt,name=is_destroy_when_disconnect,json=isDestroyWhenDisconnect,proto3" json:"is_destroy_when_disconnect,omitempty"`
}

func (x *SceneCreateEntityReq) Reset() {
	*x = SceneCreateEntityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneCreateEntityReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneCreateEntityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneCreateEntityReq) ProtoMessage() {}

func (x *SceneCreateEntityReq) ProtoReflect() protoreflect.Message {
	mi := &file_SceneCreateEntityReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneCreateEntityReq.ProtoReflect.Descriptor instead.
func (*SceneCreateEntityReq) Descriptor() ([]byte, []int) {
	return file_SceneCreateEntityReq_proto_rawDescGZIP(), []int{0}
}

func (x *SceneCreateEntityReq) GetEntity() *CreateEntityInfo {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *SceneCreateEntityReq) GetReason() CreateReason {
	if x != nil {
		return x.Reason
	}
	return CreateReason_CREATE_NONE
}

func (x *SceneCreateEntityReq) GetIsDestroyWhenDisconnect() bool {
	if x != nil {
		return x.IsDestroyWhenDisconnect
	}
	return false
}

var File_SceneCreateEntityReq_proto protoreflect.FileDescriptor

var file_SceneCreateEntityReq_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x02, 0x0a, 0x14, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x12, 0x29, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x25, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x1a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x72, 0x6f,
	0x79, 0x5f, 0x77, 0x68, 0x65, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x69, 0x73, 0x44, 0x65, 0x73, 0x74, 0x72,
	0x6f, 0x79, 0x57, 0x68, 0x65, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xe3, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53,
	0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49,
	0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01,
	0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneCreateEntityReq_proto_rawDescOnce sync.Once
	file_SceneCreateEntityReq_proto_rawDescData = file_SceneCreateEntityReq_proto_rawDesc
)

func file_SceneCreateEntityReq_proto_rawDescGZIP() []byte {
	file_SceneCreateEntityReq_proto_rawDescOnce.Do(func() {
		file_SceneCreateEntityReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneCreateEntityReq_proto_rawDescData)
	})
	return file_SceneCreateEntityReq_proto_rawDescData
}

var file_SceneCreateEntityReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SceneCreateEntityReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneCreateEntityReq_proto_goTypes = []interface{}{
	(SceneCreateEntityReq_CmdId)(0), // 0: SceneCreateEntityReq.CmdId
	(*SceneCreateEntityReq)(nil),    // 1: SceneCreateEntityReq
	(*CreateEntityInfo)(nil),        // 2: CreateEntityInfo
	(CreateReason)(0),               // 3: CreateReason
}
var file_SceneCreateEntityReq_proto_depIdxs = []int32{
	2, // 0: SceneCreateEntityReq.entity:type_name -> CreateEntityInfo
	3, // 1: SceneCreateEntityReq.reason:type_name -> CreateReason
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SceneCreateEntityReq_proto_init() }
func file_SceneCreateEntityReq_proto_init() {
	if File_SceneCreateEntityReq_proto != nil {
		return
	}
	file_CreateEntityInfo_proto_init()
	file_CreateReason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneCreateEntityReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneCreateEntityReq); i {
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
			RawDescriptor: file_SceneCreateEntityReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneCreateEntityReq_proto_goTypes,
		DependencyIndexes: file_SceneCreateEntityReq_proto_depIdxs,
		EnumInfos:         file_SceneCreateEntityReq_proto_enumTypes,
		MessageInfos:      file_SceneCreateEntityReq_proto_msgTypes,
	}.Build()
	File_SceneCreateEntityReq_proto = out.File
	file_SceneCreateEntityReq_proto_rawDesc = nil
	file_SceneCreateEntityReq_proto_goTypes = nil
	file_SceneCreateEntityReq_proto_depIdxs = nil
}
