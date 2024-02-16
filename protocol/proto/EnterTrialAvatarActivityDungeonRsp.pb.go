// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: EnterTrialAvatarActivityDungeonRsp.proto

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

type EnterTrialAvatarActivityDungeonRsp_CmdId int32

const (
	EnterTrialAvatarActivityDungeonRsp_NONE             EnterTrialAvatarActivityDungeonRsp_CmdId = 0
	EnterTrialAvatarActivityDungeonRsp_CMD_ID           EnterTrialAvatarActivityDungeonRsp_CmdId = 2042
	EnterTrialAvatarActivityDungeonRsp_ENET_CHANNEL_ID  EnterTrialAvatarActivityDungeonRsp_CmdId = 0
	EnterTrialAvatarActivityDungeonRsp_ENET_IS_RELIABLE EnterTrialAvatarActivityDungeonRsp_CmdId = 1
)

// Enum value maps for EnterTrialAvatarActivityDungeonRsp_CmdId.
var (
	EnterTrialAvatarActivityDungeonRsp_CmdId_name = map[int32]string{
		0:    "NONE",
		2042: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	EnterTrialAvatarActivityDungeonRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2042,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x EnterTrialAvatarActivityDungeonRsp_CmdId) Enum() *EnterTrialAvatarActivityDungeonRsp_CmdId {
	p := new(EnterTrialAvatarActivityDungeonRsp_CmdId)
	*p = x
	return p
}

func (x EnterTrialAvatarActivityDungeonRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnterTrialAvatarActivityDungeonRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_EnterTrialAvatarActivityDungeonRsp_proto_enumTypes[0].Descriptor()
}

func (EnterTrialAvatarActivityDungeonRsp_CmdId) Type() protoreflect.EnumType {
	return &file_EnterTrialAvatarActivityDungeonRsp_proto_enumTypes[0]
}

func (x EnterTrialAvatarActivityDungeonRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnterTrialAvatarActivityDungeonRsp_CmdId.Descriptor instead.
func (EnterTrialAvatarActivityDungeonRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_EnterTrialAvatarActivityDungeonRsp_proto_rawDescGZIP(), []int{0, 0}
}

type EnterTrialAvatarActivityDungeonRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode            int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ActivityId         uint32 `protobuf:"varint,2,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	TrialAvatarIndexId uint32 `protobuf:"varint,3,opt,name=trial_avatar_index_id,json=trialAvatarIndexId,proto3" json:"trial_avatar_index_id,omitempty"`
}

func (x *EnterTrialAvatarActivityDungeonRsp) Reset() {
	*x = EnterTrialAvatarActivityDungeonRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EnterTrialAvatarActivityDungeonRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterTrialAvatarActivityDungeonRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterTrialAvatarActivityDungeonRsp) ProtoMessage() {}

func (x *EnterTrialAvatarActivityDungeonRsp) ProtoReflect() protoreflect.Message {
	mi := &file_EnterTrialAvatarActivityDungeonRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterTrialAvatarActivityDungeonRsp.ProtoReflect.Descriptor instead.
func (*EnterTrialAvatarActivityDungeonRsp) Descriptor() ([]byte, []int) {
	return file_EnterTrialAvatarActivityDungeonRsp_proto_rawDescGZIP(), []int{0}
}

func (x *EnterTrialAvatarActivityDungeonRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *EnterTrialAvatarActivityDungeonRsp) GetActivityId() uint32 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

func (x *EnterTrialAvatarActivityDungeonRsp) GetTrialAvatarIndexId() uint32 {
	if x != nil {
		return x.TrialAvatarIndexId
	}
	return 0
}

var File_EnterTrialAvatarActivityDungeonRsp_proto protoreflect.FileDescriptor

var file_EnterTrialAvatarActivityDungeonRsp_proto_rawDesc = []byte{
	0x0a, 0x28, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x22, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x72, 0x69, 0x61, 0x6c,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x75,
	0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x12, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44,
	0x5f, 0x49, 0x44, 0x10, 0xfa, 0x0f, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45,
	0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EnterTrialAvatarActivityDungeonRsp_proto_rawDescOnce sync.Once
	file_EnterTrialAvatarActivityDungeonRsp_proto_rawDescData = file_EnterTrialAvatarActivityDungeonRsp_proto_rawDesc
)

func file_EnterTrialAvatarActivityDungeonRsp_proto_rawDescGZIP() []byte {
	file_EnterTrialAvatarActivityDungeonRsp_proto_rawDescOnce.Do(func() {
		file_EnterTrialAvatarActivityDungeonRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_EnterTrialAvatarActivityDungeonRsp_proto_rawDescData)
	})
	return file_EnterTrialAvatarActivityDungeonRsp_proto_rawDescData
}

var file_EnterTrialAvatarActivityDungeonRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EnterTrialAvatarActivityDungeonRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EnterTrialAvatarActivityDungeonRsp_proto_goTypes = []interface{}{
	(EnterTrialAvatarActivityDungeonRsp_CmdId)(0), // 0: proto.EnterTrialAvatarActivityDungeonRsp.CmdId
	(*EnterTrialAvatarActivityDungeonRsp)(nil),    // 1: proto.EnterTrialAvatarActivityDungeonRsp
}
var file_EnterTrialAvatarActivityDungeonRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EnterTrialAvatarActivityDungeonRsp_proto_init() }
func file_EnterTrialAvatarActivityDungeonRsp_proto_init() {
	if File_EnterTrialAvatarActivityDungeonRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EnterTrialAvatarActivityDungeonRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterTrialAvatarActivityDungeonRsp); i {
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
			RawDescriptor: file_EnterTrialAvatarActivityDungeonRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EnterTrialAvatarActivityDungeonRsp_proto_goTypes,
		DependencyIndexes: file_EnterTrialAvatarActivityDungeonRsp_proto_depIdxs,
		EnumInfos:         file_EnterTrialAvatarActivityDungeonRsp_proto_enumTypes,
		MessageInfos:      file_EnterTrialAvatarActivityDungeonRsp_proto_msgTypes,
	}.Build()
	File_EnterTrialAvatarActivityDungeonRsp_proto = out.File
	file_EnterTrialAvatarActivityDungeonRsp_proto_rawDesc = nil
	file_EnterTrialAvatarActivityDungeonRsp_proto_goTypes = nil
	file_EnterTrialAvatarActivityDungeonRsp_proto_depIdxs = nil
}