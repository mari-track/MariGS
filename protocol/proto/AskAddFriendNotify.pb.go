// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: AskAddFriendNotify.proto

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

type AskAddFriendNotify_CmdId int32

const (
	AskAddFriendNotify_NONE             AskAddFriendNotify_CmdId = 0
	AskAddFriendNotify_CMD_ID           AskAddFriendNotify_CmdId = 4022
	AskAddFriendNotify_ENET_CHANNEL_ID  AskAddFriendNotify_CmdId = 0
	AskAddFriendNotify_ENET_IS_RELIABLE AskAddFriendNotify_CmdId = 1
)

// Enum value maps for AskAddFriendNotify_CmdId.
var (
	AskAddFriendNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		4022: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	AskAddFriendNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           4022,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x AskAddFriendNotify_CmdId) Enum() *AskAddFriendNotify_CmdId {
	p := new(AskAddFriendNotify_CmdId)
	*p = x
	return p
}

func (x AskAddFriendNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AskAddFriendNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_AskAddFriendNotify_proto_enumTypes[0].Descriptor()
}

func (AskAddFriendNotify_CmdId) Type() protoreflect.EnumType {
	return &file_AskAddFriendNotify_proto_enumTypes[0]
}

func (x AskAddFriendNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AskAddFriendNotify_CmdId.Descriptor instead.
func (AskAddFriendNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_AskAddFriendNotify_proto_rawDescGZIP(), []int{0, 0}
}

type AskAddFriendNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetUid         uint32       `protobuf:"varint,1,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
	TargetFriendBrief *FriendBrief `protobuf:"bytes,2,opt,name=target_friend_brief,json=targetFriendBrief,proto3" json:"target_friend_brief,omitempty"`
}

func (x *AskAddFriendNotify) Reset() {
	*x = AskAddFriendNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AskAddFriendNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskAddFriendNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskAddFriendNotify) ProtoMessage() {}

func (x *AskAddFriendNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AskAddFriendNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskAddFriendNotify.ProtoReflect.Descriptor instead.
func (*AskAddFriendNotify) Descriptor() ([]byte, []int) {
	return file_AskAddFriendNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AskAddFriendNotify) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

func (x *AskAddFriendNotify) GetTargetFriendBrief() *FriendBrief {
	if x != nil {
		return x.TargetFriendBrief
	}
	return nil
}

var File_AskAddFriendNotify_proto protoreflect.FileDescriptor

var file_AskAddFriendNotify_proto_rawDesc = []byte{
	0x0a, 0x18, 0x41, 0x73, 0x6b, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x42, 0x72, 0x69, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01,
	0x0a, 0x12, 0x41, 0x73, 0x6b, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x55, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x42, 0x72, 0x69, 0x65, 0x66, 0x52, 0x11,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x42, 0x72, 0x69, 0x65,
	0x66, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xb6,
	0x1f, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45,
	0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49,
	0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AskAddFriendNotify_proto_rawDescOnce sync.Once
	file_AskAddFriendNotify_proto_rawDescData = file_AskAddFriendNotify_proto_rawDesc
)

func file_AskAddFriendNotify_proto_rawDescGZIP() []byte {
	file_AskAddFriendNotify_proto_rawDescOnce.Do(func() {
		file_AskAddFriendNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AskAddFriendNotify_proto_rawDescData)
	})
	return file_AskAddFriendNotify_proto_rawDescData
}

var file_AskAddFriendNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AskAddFriendNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AskAddFriendNotify_proto_goTypes = []interface{}{
	(AskAddFriendNotify_CmdId)(0), // 0: AskAddFriendNotify.CmdId
	(*AskAddFriendNotify)(nil),    // 1: AskAddFriendNotify
	(*FriendBrief)(nil),           // 2: FriendBrief
}
var file_AskAddFriendNotify_proto_depIdxs = []int32{
	2, // 0: AskAddFriendNotify.target_friend_brief:type_name -> FriendBrief
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AskAddFriendNotify_proto_init() }
func file_AskAddFriendNotify_proto_init() {
	if File_AskAddFriendNotify_proto != nil {
		return
	}
	file_FriendBrief_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AskAddFriendNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskAddFriendNotify); i {
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
			RawDescriptor: file_AskAddFriendNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AskAddFriendNotify_proto_goTypes,
		DependencyIndexes: file_AskAddFriendNotify_proto_depIdxs,
		EnumInfos:         file_AskAddFriendNotify_proto_enumTypes,
		MessageInfos:      file_AskAddFriendNotify_proto_msgTypes,
	}.Build()
	File_AskAddFriendNotify_proto = out.File
	file_AskAddFriendNotify_proto_rawDesc = nil
	file_AskAddFriendNotify_proto_goTypes = nil
	file_AskAddFriendNotify_proto_depIdxs = nil
}
