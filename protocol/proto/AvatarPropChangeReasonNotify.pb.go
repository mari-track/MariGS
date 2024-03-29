// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: AvatarPropChangeReasonNotify.proto

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

type AvatarPropChangeReasonNotify_CmdId int32

const (
	AvatarPropChangeReasonNotify_NONE             AvatarPropChangeReasonNotify_CmdId = 0
	AvatarPropChangeReasonNotify_CMD_ID           AvatarPropChangeReasonNotify_CmdId = 1209
	AvatarPropChangeReasonNotify_ENET_CHANNEL_ID  AvatarPropChangeReasonNotify_CmdId = 0
	AvatarPropChangeReasonNotify_ENET_IS_RELIABLE AvatarPropChangeReasonNotify_CmdId = 1
)

// Enum value maps for AvatarPropChangeReasonNotify_CmdId.
var (
	AvatarPropChangeReasonNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		1209: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	AvatarPropChangeReasonNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1209,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x AvatarPropChangeReasonNotify_CmdId) Enum() *AvatarPropChangeReasonNotify_CmdId {
	p := new(AvatarPropChangeReasonNotify_CmdId)
	*p = x
	return p
}

func (x AvatarPropChangeReasonNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AvatarPropChangeReasonNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_AvatarPropChangeReasonNotify_proto_enumTypes[0].Descriptor()
}

func (AvatarPropChangeReasonNotify_CmdId) Type() protoreflect.EnumType {
	return &file_AvatarPropChangeReasonNotify_proto_enumTypes[0]
}

func (x AvatarPropChangeReasonNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AvatarPropChangeReasonNotify_CmdId.Descriptor instead.
func (AvatarPropChangeReasonNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_AvatarPropChangeReasonNotify_proto_rawDescGZIP(), []int{0, 0}
}

type AvatarPropChangeReasonNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarGuid uint64           `protobuf:"varint,1,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
	PropType   uint32           `protobuf:"varint,2,opt,name=prop_type,json=propType,proto3" json:"prop_type,omitempty"`
	OldValue   float32          `protobuf:"fixed32,3,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
	CurValue   float32          `protobuf:"fixed32,4,opt,name=cur_value,json=curValue,proto3" json:"cur_value,omitempty"`
	Reason     PropChangeReason `protobuf:"varint,5,opt,name=reason,proto3,enum=PropChangeReason" json:"reason,omitempty"`
}

func (x *AvatarPropChangeReasonNotify) Reset() {
	*x = AvatarPropChangeReasonNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarPropChangeReasonNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarPropChangeReasonNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarPropChangeReasonNotify) ProtoMessage() {}

func (x *AvatarPropChangeReasonNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarPropChangeReasonNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarPropChangeReasonNotify.ProtoReflect.Descriptor instead.
func (*AvatarPropChangeReasonNotify) Descriptor() ([]byte, []int) {
	return file_AvatarPropChangeReasonNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarPropChangeReasonNotify) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

func (x *AvatarPropChangeReasonNotify) GetPropType() uint32 {
	if x != nil {
		return x.PropType
	}
	return 0
}

func (x *AvatarPropChangeReasonNotify) GetOldValue() float32 {
	if x != nil {
		return x.OldValue
	}
	return 0
}

func (x *AvatarPropChangeReasonNotify) GetCurValue() float32 {
	if x != nil {
		return x.CurValue
	}
	return 0
}

func (x *AvatarPropChangeReasonNotify) GetReason() PropChangeReason {
	if x != nil {
		return x.Reason
	}
	return PropChangeReason_PROP_CHANGE_NONE
}

var File_AvatarPropChangeReasonNotify_proto protoreflect.FileDescriptor

var file_AvatarPropChangeReasonNotify_proto_rawDesc = []byte{
	0x0a, 0x22, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x50, 0x72, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x02, 0x0a,
	0x1c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f,
	0x6c, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x6f, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xb9, 0x09,
	0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53,
	0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_AvatarPropChangeReasonNotify_proto_rawDescOnce sync.Once
	file_AvatarPropChangeReasonNotify_proto_rawDescData = file_AvatarPropChangeReasonNotify_proto_rawDesc
)

func file_AvatarPropChangeReasonNotify_proto_rawDescGZIP() []byte {
	file_AvatarPropChangeReasonNotify_proto_rawDescOnce.Do(func() {
		file_AvatarPropChangeReasonNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarPropChangeReasonNotify_proto_rawDescData)
	})
	return file_AvatarPropChangeReasonNotify_proto_rawDescData
}

var file_AvatarPropChangeReasonNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AvatarPropChangeReasonNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarPropChangeReasonNotify_proto_goTypes = []interface{}{
	(AvatarPropChangeReasonNotify_CmdId)(0), // 0: AvatarPropChangeReasonNotify.CmdId
	(*AvatarPropChangeReasonNotify)(nil),    // 1: AvatarPropChangeReasonNotify
	(PropChangeReason)(0),                   // 2: PropChangeReason
}
var file_AvatarPropChangeReasonNotify_proto_depIdxs = []int32{
	2, // 0: AvatarPropChangeReasonNotify.reason:type_name -> PropChangeReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AvatarPropChangeReasonNotify_proto_init() }
func file_AvatarPropChangeReasonNotify_proto_init() {
	if File_AvatarPropChangeReasonNotify_proto != nil {
		return
	}
	file_PropChangeReason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarPropChangeReasonNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarPropChangeReasonNotify); i {
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
			RawDescriptor: file_AvatarPropChangeReasonNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarPropChangeReasonNotify_proto_goTypes,
		DependencyIndexes: file_AvatarPropChangeReasonNotify_proto_depIdxs,
		EnumInfos:         file_AvatarPropChangeReasonNotify_proto_enumTypes,
		MessageInfos:      file_AvatarPropChangeReasonNotify_proto_msgTypes,
	}.Build()
	File_AvatarPropChangeReasonNotify_proto = out.File
	file_AvatarPropChangeReasonNotify_proto_rawDesc = nil
	file_AvatarPropChangeReasonNotify_proto_goTypes = nil
	file_AvatarPropChangeReasonNotify_proto_depIdxs = nil
}
