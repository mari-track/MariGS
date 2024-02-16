// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: ReceivedTrialAvatarActivityRewardReq.proto

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

type ReceivedTrialAvatarActivityRewardReq_CmdId int32

const (
	ReceivedTrialAvatarActivityRewardReq_NONE             ReceivedTrialAvatarActivityRewardReq_CmdId = 0
	ReceivedTrialAvatarActivityRewardReq_CMD_ID           ReceivedTrialAvatarActivityRewardReq_CmdId = 2043
	ReceivedTrialAvatarActivityRewardReq_ENET_CHANNEL_ID  ReceivedTrialAvatarActivityRewardReq_CmdId = 0
	ReceivedTrialAvatarActivityRewardReq_ENET_IS_RELIABLE ReceivedTrialAvatarActivityRewardReq_CmdId = 1
	ReceivedTrialAvatarActivityRewardReq_IS_ALLOW_CLIENT  ReceivedTrialAvatarActivityRewardReq_CmdId = 1
)

// Enum value maps for ReceivedTrialAvatarActivityRewardReq_CmdId.
var (
	ReceivedTrialAvatarActivityRewardReq_CmdId_name = map[int32]string{
		0:    "NONE",
		2043: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	ReceivedTrialAvatarActivityRewardReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2043,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x ReceivedTrialAvatarActivityRewardReq_CmdId) Enum() *ReceivedTrialAvatarActivityRewardReq_CmdId {
	p := new(ReceivedTrialAvatarActivityRewardReq_CmdId)
	*p = x
	return p
}

func (x ReceivedTrialAvatarActivityRewardReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReceivedTrialAvatarActivityRewardReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ReceivedTrialAvatarActivityRewardReq_proto_enumTypes[0].Descriptor()
}

func (ReceivedTrialAvatarActivityRewardReq_CmdId) Type() protoreflect.EnumType {
	return &file_ReceivedTrialAvatarActivityRewardReq_proto_enumTypes[0]
}

func (x ReceivedTrialAvatarActivityRewardReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReceivedTrialAvatarActivityRewardReq_CmdId.Descriptor instead.
func (ReceivedTrialAvatarActivityRewardReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ReceivedTrialAvatarActivityRewardReq_proto_rawDescGZIP(), []int{0, 0}
}

type ReceivedTrialAvatarActivityRewardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrialAvatarIndexId uint32 `protobuf:"varint,1,opt,name=trial_avatar_index_id,json=trialAvatarIndexId,proto3" json:"trial_avatar_index_id,omitempty"`
}

func (x *ReceivedTrialAvatarActivityRewardReq) Reset() {
	*x = ReceivedTrialAvatarActivityRewardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ReceivedTrialAvatarActivityRewardReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceivedTrialAvatarActivityRewardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceivedTrialAvatarActivityRewardReq) ProtoMessage() {}

func (x *ReceivedTrialAvatarActivityRewardReq) ProtoReflect() protoreflect.Message {
	mi := &file_ReceivedTrialAvatarActivityRewardReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceivedTrialAvatarActivityRewardReq.ProtoReflect.Descriptor instead.
func (*ReceivedTrialAvatarActivityRewardReq) Descriptor() ([]byte, []int) {
	return file_ReceivedTrialAvatarActivityRewardReq_proto_rawDescGZIP(), []int{0}
}

func (x *ReceivedTrialAvatarActivityRewardReq) GetTrialAvatarIndexId() uint32 {
	if x != nil {
		return x.TrialAvatarIndexId
	}
	return 0
}

var File_ReceivedTrialAvatarActivityRewardReq_proto protoreflect.FileDescriptor

var file_ReceivedTrialAvatarActivityRewardReq_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x24, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x31, 0x0a, 0x15,
	0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x74, 0x72, 0x69,
	0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x64, 0x22,
	0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xfb, 0x0f, 0x12,
	0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f,
	0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53,
	0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a,
	0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ReceivedTrialAvatarActivityRewardReq_proto_rawDescOnce sync.Once
	file_ReceivedTrialAvatarActivityRewardReq_proto_rawDescData = file_ReceivedTrialAvatarActivityRewardReq_proto_rawDesc
)

func file_ReceivedTrialAvatarActivityRewardReq_proto_rawDescGZIP() []byte {
	file_ReceivedTrialAvatarActivityRewardReq_proto_rawDescOnce.Do(func() {
		file_ReceivedTrialAvatarActivityRewardReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ReceivedTrialAvatarActivityRewardReq_proto_rawDescData)
	})
	return file_ReceivedTrialAvatarActivityRewardReq_proto_rawDescData
}

var file_ReceivedTrialAvatarActivityRewardReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ReceivedTrialAvatarActivityRewardReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ReceivedTrialAvatarActivityRewardReq_proto_goTypes = []interface{}{
	(ReceivedTrialAvatarActivityRewardReq_CmdId)(0), // 0: proto.ReceivedTrialAvatarActivityRewardReq.CmdId
	(*ReceivedTrialAvatarActivityRewardReq)(nil),    // 1: proto.ReceivedTrialAvatarActivityRewardReq
}
var file_ReceivedTrialAvatarActivityRewardReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ReceivedTrialAvatarActivityRewardReq_proto_init() }
func file_ReceivedTrialAvatarActivityRewardReq_proto_init() {
	if File_ReceivedTrialAvatarActivityRewardReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ReceivedTrialAvatarActivityRewardReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceivedTrialAvatarActivityRewardReq); i {
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
			RawDescriptor: file_ReceivedTrialAvatarActivityRewardReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ReceivedTrialAvatarActivityRewardReq_proto_goTypes,
		DependencyIndexes: file_ReceivedTrialAvatarActivityRewardReq_proto_depIdxs,
		EnumInfos:         file_ReceivedTrialAvatarActivityRewardReq_proto_enumTypes,
		MessageInfos:      file_ReceivedTrialAvatarActivityRewardReq_proto_msgTypes,
	}.Build()
	File_ReceivedTrialAvatarActivityRewardReq_proto = out.File
	file_ReceivedTrialAvatarActivityRewardReq_proto_rawDesc = nil
	file_ReceivedTrialAvatarActivityRewardReq_proto_goTypes = nil
	file_ReceivedTrialAvatarActivityRewardReq_proto_depIdxs = nil
}
