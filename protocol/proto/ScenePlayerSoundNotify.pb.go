// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: ScenePlayerSoundNotify.proto

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

type ScenePlayerSoundNotify_PlaySoundType int32

const (
	ScenePlayerSoundNotify_PLAY_SOUND_NONE  ScenePlayerSoundNotify_PlaySoundType = 0
	ScenePlayerSoundNotify_PLAY_SOUND_START ScenePlayerSoundNotify_PlaySoundType = 1
	ScenePlayerSoundNotify_PLAY_SOUND_STOP  ScenePlayerSoundNotify_PlaySoundType = 2
)

// Enum value maps for ScenePlayerSoundNotify_PlaySoundType.
var (
	ScenePlayerSoundNotify_PlaySoundType_name = map[int32]string{
		0: "PLAY_SOUND_NONE",
		1: "PLAY_SOUND_START",
		2: "PLAY_SOUND_STOP",
	}
	ScenePlayerSoundNotify_PlaySoundType_value = map[string]int32{
		"PLAY_SOUND_NONE":  0,
		"PLAY_SOUND_START": 1,
		"PLAY_SOUND_STOP":  2,
	}
)

func (x ScenePlayerSoundNotify_PlaySoundType) Enum() *ScenePlayerSoundNotify_PlaySoundType {
	p := new(ScenePlayerSoundNotify_PlaySoundType)
	*p = x
	return p
}

func (x ScenePlayerSoundNotify_PlaySoundType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScenePlayerSoundNotify_PlaySoundType) Descriptor() protoreflect.EnumDescriptor {
	return file_ScenePlayerSoundNotify_proto_enumTypes[0].Descriptor()
}

func (ScenePlayerSoundNotify_PlaySoundType) Type() protoreflect.EnumType {
	return &file_ScenePlayerSoundNotify_proto_enumTypes[0]
}

func (x ScenePlayerSoundNotify_PlaySoundType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScenePlayerSoundNotify_PlaySoundType.Descriptor instead.
func (ScenePlayerSoundNotify_PlaySoundType) EnumDescriptor() ([]byte, []int) {
	return file_ScenePlayerSoundNotify_proto_rawDescGZIP(), []int{0, 0}
}

type ScenePlayerSoundNotify_CmdId int32

const (
	ScenePlayerSoundNotify_NONE             ScenePlayerSoundNotify_CmdId = 0
	ScenePlayerSoundNotify_CMD_ID           ScenePlayerSoundNotify_CmdId = 261
	ScenePlayerSoundNotify_ENET_CHANNEL_ID  ScenePlayerSoundNotify_CmdId = 0
	ScenePlayerSoundNotify_ENET_IS_RELIABLE ScenePlayerSoundNotify_CmdId = 1
)

// Enum value maps for ScenePlayerSoundNotify_CmdId.
var (
	ScenePlayerSoundNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		261: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	ScenePlayerSoundNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           261,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x ScenePlayerSoundNotify_CmdId) Enum() *ScenePlayerSoundNotify_CmdId {
	p := new(ScenePlayerSoundNotify_CmdId)
	*p = x
	return p
}

func (x ScenePlayerSoundNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScenePlayerSoundNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ScenePlayerSoundNotify_proto_enumTypes[1].Descriptor()
}

func (ScenePlayerSoundNotify_CmdId) Type() protoreflect.EnumType {
	return &file_ScenePlayerSoundNotify_proto_enumTypes[1]
}

func (x ScenePlayerSoundNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScenePlayerSoundNotify_CmdId.Descriptor instead.
func (ScenePlayerSoundNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ScenePlayerSoundNotify_proto_rawDescGZIP(), []int{0, 1}
}

type ScenePlayerSoundNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SoundName string                               `protobuf:"bytes,1,opt,name=sound_name,json=soundName,proto3" json:"sound_name,omitempty"`
	PlayPos   *Vector                              `protobuf:"bytes,2,opt,name=play_pos,json=playPos,proto3" json:"play_pos,omitempty"`
	PlayType  ScenePlayerSoundNotify_PlaySoundType `protobuf:"varint,3,opt,name=play_type,json=playType,proto3,enum=proto.ScenePlayerSoundNotify_PlaySoundType" json:"play_type,omitempty"`
}

func (x *ScenePlayerSoundNotify) Reset() {
	*x = ScenePlayerSoundNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ScenePlayerSoundNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScenePlayerSoundNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScenePlayerSoundNotify) ProtoMessage() {}

func (x *ScenePlayerSoundNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ScenePlayerSoundNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScenePlayerSoundNotify.ProtoReflect.Descriptor instead.
func (*ScenePlayerSoundNotify) Descriptor() ([]byte, []int) {
	return file_ScenePlayerSoundNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ScenePlayerSoundNotify) GetSoundName() string {
	if x != nil {
		return x.SoundName
	}
	return ""
}

func (x *ScenePlayerSoundNotify) GetPlayPos() *Vector {
	if x != nil {
		return x.PlayPos
	}
	return nil
}

func (x *ScenePlayerSoundNotify) GetPlayType() ScenePlayerSoundNotify_PlaySoundType {
	if x != nil {
		return x.PlayType
	}
	return ScenePlayerSoundNotify_PLAY_SOUND_NONE
}

var File_ScenePlayerSoundNotify_proto protoreflect.FileDescriptor

var file_ScenePlayerSoundNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6f, 0x75,
	0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x02, 0x0a, 0x16, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07,
	0x70, 0x6c, 0x61, 0x79, 0x50, 0x6f, 0x73, 0x12, 0x48, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6f,
	0x75, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x53, 0x6f,
	0x75, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x4f, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x53, 0x4f, 0x55, 0x4e, 0x44,
	0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4c, 0x41, 0x59, 0x5f,
	0x53, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x53, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x4f, 0x50,
	0x10, 0x02, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0x85, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e,
	0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ScenePlayerSoundNotify_proto_rawDescOnce sync.Once
	file_ScenePlayerSoundNotify_proto_rawDescData = file_ScenePlayerSoundNotify_proto_rawDesc
)

func file_ScenePlayerSoundNotify_proto_rawDescGZIP() []byte {
	file_ScenePlayerSoundNotify_proto_rawDescOnce.Do(func() {
		file_ScenePlayerSoundNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ScenePlayerSoundNotify_proto_rawDescData)
	})
	return file_ScenePlayerSoundNotify_proto_rawDescData
}

var file_ScenePlayerSoundNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ScenePlayerSoundNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ScenePlayerSoundNotify_proto_goTypes = []interface{}{
	(ScenePlayerSoundNotify_PlaySoundType)(0), // 0: proto.ScenePlayerSoundNotify.PlaySoundType
	(ScenePlayerSoundNotify_CmdId)(0),         // 1: proto.ScenePlayerSoundNotify.CmdId
	(*ScenePlayerSoundNotify)(nil),            // 2: proto.ScenePlayerSoundNotify
	(*Vector)(nil),                            // 3: proto.Vector
}
var file_ScenePlayerSoundNotify_proto_depIdxs = []int32{
	3, // 0: proto.ScenePlayerSoundNotify.play_pos:type_name -> proto.Vector
	0, // 1: proto.ScenePlayerSoundNotify.play_type:type_name -> proto.ScenePlayerSoundNotify.PlaySoundType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ScenePlayerSoundNotify_proto_init() }
func file_ScenePlayerSoundNotify_proto_init() {
	if File_ScenePlayerSoundNotify_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ScenePlayerSoundNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScenePlayerSoundNotify); i {
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
			RawDescriptor: file_ScenePlayerSoundNotify_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ScenePlayerSoundNotify_proto_goTypes,
		DependencyIndexes: file_ScenePlayerSoundNotify_proto_depIdxs,
		EnumInfos:         file_ScenePlayerSoundNotify_proto_enumTypes,
		MessageInfos:      file_ScenePlayerSoundNotify_proto_msgTypes,
	}.Build()
	File_ScenePlayerSoundNotify_proto = out.File
	file_ScenePlayerSoundNotify_proto_rawDesc = nil
	file_ScenePlayerSoundNotify_proto_goTypes = nil
	file_ScenePlayerSoundNotify_proto_depIdxs = nil
}