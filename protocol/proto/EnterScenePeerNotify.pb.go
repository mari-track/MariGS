// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: EnterScenePeerNotify.proto

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

type EnterScenePeerNotify_CmdId int32

const (
	EnterScenePeerNotify_NONE             EnterScenePeerNotify_CmdId = 0
	EnterScenePeerNotify_CMD_ID           EnterScenePeerNotify_CmdId = 284
	EnterScenePeerNotify_ENET_CHANNEL_ID  EnterScenePeerNotify_CmdId = 0
	EnterScenePeerNotify_ENET_IS_RELIABLE EnterScenePeerNotify_CmdId = 1
)

// Enum value maps for EnterScenePeerNotify_CmdId.
var (
	EnterScenePeerNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		284: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	EnterScenePeerNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           284,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x EnterScenePeerNotify_CmdId) Enum() *EnterScenePeerNotify_CmdId {
	p := new(EnterScenePeerNotify_CmdId)
	*p = x
	return p
}

func (x EnterScenePeerNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnterScenePeerNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_EnterScenePeerNotify_proto_enumTypes[0].Descriptor()
}

func (EnterScenePeerNotify_CmdId) Type() protoreflect.EnumType {
	return &file_EnterScenePeerNotify_proto_enumTypes[0]
}

func (x EnterScenePeerNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnterScenePeerNotify_CmdId.Descriptor instead.
func (EnterScenePeerNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_EnterScenePeerNotify_proto_rawDescGZIP(), []int{0, 0}
}

type EnterScenePeerNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestSceneId     uint32 `protobuf:"varint,1,opt,name=dest_scene_id,json=destSceneId,proto3" json:"dest_scene_id,omitempty"`
	PeerId          uint32 `protobuf:"varint,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	HostPeerId      uint32 `protobuf:"varint,3,opt,name=host_peer_id,json=hostPeerId,proto3" json:"host_peer_id,omitempty"`
	EnterSceneToken uint32 `protobuf:"varint,4,opt,name=enter_scene_token,json=enterSceneToken,proto3" json:"enter_scene_token,omitempty"`
}

func (x *EnterScenePeerNotify) Reset() {
	*x = EnterScenePeerNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EnterScenePeerNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterScenePeerNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterScenePeerNotify) ProtoMessage() {}

func (x *EnterScenePeerNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EnterScenePeerNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterScenePeerNotify.ProtoReflect.Descriptor instead.
func (*EnterScenePeerNotify) Descriptor() ([]byte, []int) {
	return file_EnterScenePeerNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EnterScenePeerNotify) GetDestSceneId() uint32 {
	if x != nil {
		return x.DestSceneId
	}
	return 0
}

func (x *EnterScenePeerNotify) GetPeerId() uint32 {
	if x != nil {
		return x.PeerId
	}
	return 0
}

func (x *EnterScenePeerNotify) GetHostPeerId() uint32 {
	if x != nil {
		return x.HostPeerId
	}
	return 0
}

func (x *EnterScenePeerNotify) GetEnterSceneToken() uint32 {
	if x != nil {
		return x.EnterSceneToken
	}
	return 0
}

var File_EnterScenePeerNotify_proto protoreflect.FileDescriptor

var file_EnterScenePeerNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x65, 0x65, 0x72,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a,
	0x14, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x65, 0x65, 0x72, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x68, 0x6f, 0x73, 0x74, 0x50, 0x65,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x9c, 0x02,
	0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53,
	0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_EnterScenePeerNotify_proto_rawDescOnce sync.Once
	file_EnterScenePeerNotify_proto_rawDescData = file_EnterScenePeerNotify_proto_rawDesc
)

func file_EnterScenePeerNotify_proto_rawDescGZIP() []byte {
	file_EnterScenePeerNotify_proto_rawDescOnce.Do(func() {
		file_EnterScenePeerNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EnterScenePeerNotify_proto_rawDescData)
	})
	return file_EnterScenePeerNotify_proto_rawDescData
}

var file_EnterScenePeerNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EnterScenePeerNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EnterScenePeerNotify_proto_goTypes = []interface{}{
	(EnterScenePeerNotify_CmdId)(0), // 0: EnterScenePeerNotify.CmdId
	(*EnterScenePeerNotify)(nil),    // 1: EnterScenePeerNotify
}
var file_EnterScenePeerNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EnterScenePeerNotify_proto_init() }
func file_EnterScenePeerNotify_proto_init() {
	if File_EnterScenePeerNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EnterScenePeerNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterScenePeerNotify); i {
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
			RawDescriptor: file_EnterScenePeerNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EnterScenePeerNotify_proto_goTypes,
		DependencyIndexes: file_EnterScenePeerNotify_proto_depIdxs,
		EnumInfos:         file_EnterScenePeerNotify_proto_enumTypes,
		MessageInfos:      file_EnterScenePeerNotify_proto_msgTypes,
	}.Build()
	File_EnterScenePeerNotify_proto = out.File
	file_EnterScenePeerNotify_proto_rawDesc = nil
	file_EnterScenePeerNotify_proto_goTypes = nil
	file_EnterScenePeerNotify_proto_depIdxs = nil
}
