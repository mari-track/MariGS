// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GetSceneAreaReq.proto

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

type GetSceneAreaReq_CmdId int32

const (
	GetSceneAreaReq_NONE             GetSceneAreaReq_CmdId = 0
	GetSceneAreaReq_CMD_ID           GetSceneAreaReq_CmdId = 222
	GetSceneAreaReq_ENET_CHANNEL_ID  GetSceneAreaReq_CmdId = 0
	GetSceneAreaReq_ENET_IS_RELIABLE GetSceneAreaReq_CmdId = 1
	GetSceneAreaReq_IS_ALLOW_CLIENT  GetSceneAreaReq_CmdId = 1
)

// Enum value maps for GetSceneAreaReq_CmdId.
var (
	GetSceneAreaReq_CmdId_name = map[int32]string{
		0:   "NONE",
		222: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	GetSceneAreaReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           222,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x GetSceneAreaReq_CmdId) Enum() *GetSceneAreaReq_CmdId {
	p := new(GetSceneAreaReq_CmdId)
	*p = x
	return p
}

func (x GetSceneAreaReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetSceneAreaReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetSceneAreaReq_proto_enumTypes[0].Descriptor()
}

func (GetSceneAreaReq_CmdId) Type() protoreflect.EnumType {
	return &file_GetSceneAreaReq_proto_enumTypes[0]
}

func (x GetSceneAreaReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetSceneAreaReq_CmdId.Descriptor instead.
func (GetSceneAreaReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetSceneAreaReq_proto_rawDescGZIP(), []int{0, 0}
}

type GetSceneAreaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId   uint32 `protobuf:"varint,1,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	BelongUid uint32 `protobuf:"varint,2,opt,name=belong_uid,json=belongUid,proto3" json:"belong_uid,omitempty"`
}

func (x *GetSceneAreaReq) Reset() {
	*x = GetSceneAreaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetSceneAreaReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSceneAreaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSceneAreaReq) ProtoMessage() {}

func (x *GetSceneAreaReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetSceneAreaReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSceneAreaReq.ProtoReflect.Descriptor instead.
func (*GetSceneAreaReq) Descriptor() ([]byte, []int) {
	return file_GetSceneAreaReq_proto_rawDescGZIP(), []int{0}
}

func (x *GetSceneAreaReq) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *GetSceneAreaReq) GetBelongUid() uint32 {
	if x != nil {
		return x.BelongUid
	}
	return 0
}

var File_GetSceneAreaReq_proto protoreflect.FileDescriptor

var file_GetSceneAreaReq_proto_rawDesc = []byte{
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf,
	0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52,
	0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x55, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x05,
	0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xde, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c,
	0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c,
	0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetSceneAreaReq_proto_rawDescOnce sync.Once
	file_GetSceneAreaReq_proto_rawDescData = file_GetSceneAreaReq_proto_rawDesc
)

func file_GetSceneAreaReq_proto_rawDescGZIP() []byte {
	file_GetSceneAreaReq_proto_rawDescOnce.Do(func() {
		file_GetSceneAreaReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetSceneAreaReq_proto_rawDescData)
	})
	return file_GetSceneAreaReq_proto_rawDescData
}

var file_GetSceneAreaReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetSceneAreaReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetSceneAreaReq_proto_goTypes = []interface{}{
	(GetSceneAreaReq_CmdId)(0), // 0: proto.GetSceneAreaReq.CmdId
	(*GetSceneAreaReq)(nil),    // 1: proto.GetSceneAreaReq
}
var file_GetSceneAreaReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetSceneAreaReq_proto_init() }
func file_GetSceneAreaReq_proto_init() {
	if File_GetSceneAreaReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetSceneAreaReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSceneAreaReq); i {
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
			RawDescriptor: file_GetSceneAreaReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetSceneAreaReq_proto_goTypes,
		DependencyIndexes: file_GetSceneAreaReq_proto_depIdxs,
		EnumInfos:         file_GetSceneAreaReq_proto_enumTypes,
		MessageInfos:      file_GetSceneAreaReq_proto_msgTypes,
	}.Build()
	File_GetSceneAreaReq_proto = out.File
	file_GetSceneAreaReq_proto_rawDesc = nil
	file_GetSceneAreaReq_proto_goTypes = nil
	file_GetSceneAreaReq_proto_depIdxs = nil
}