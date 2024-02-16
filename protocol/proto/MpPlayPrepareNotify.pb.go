// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: MpPlayPrepareNotify.proto

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

type MpPlayPrepareNotify_CmdId int32

const (
	MpPlayPrepareNotify_NONE             MpPlayPrepareNotify_CmdId = 0
	MpPlayPrepareNotify_CMD_ID           MpPlayPrepareNotify_CmdId = 1823
	MpPlayPrepareNotify_ENET_CHANNEL_ID  MpPlayPrepareNotify_CmdId = 0
	MpPlayPrepareNotify_ENET_IS_RELIABLE MpPlayPrepareNotify_CmdId = 1
)

// Enum value maps for MpPlayPrepareNotify_CmdId.
var (
	MpPlayPrepareNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		1823: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	MpPlayPrepareNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1823,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x MpPlayPrepareNotify_CmdId) Enum() *MpPlayPrepareNotify_CmdId {
	p := new(MpPlayPrepareNotify_CmdId)
	*p = x
	return p
}

func (x MpPlayPrepareNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MpPlayPrepareNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_MpPlayPrepareNotify_proto_enumTypes[0].Descriptor()
}

func (MpPlayPrepareNotify_CmdId) Type() protoreflect.EnumType {
	return &file_MpPlayPrepareNotify_proto_enumTypes[0]
}

func (x MpPlayPrepareNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MpPlayPrepareNotify_CmdId.Descriptor instead.
func (MpPlayPrepareNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_MpPlayPrepareNotify_proto_rawDescGZIP(), []int{0, 0}
}

type MpPlayPrepareNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MpPlayId       uint32 `protobuf:"varint,1,opt,name=mp_play_id,json=mpPlayId,proto3" json:"mp_play_id,omitempty"`
	PrepareEndTime uint32 `protobuf:"varint,2,opt,name=prepare_end_time,json=prepareEndTime,proto3" json:"prepare_end_time,omitempty"`
}

func (x *MpPlayPrepareNotify) Reset() {
	*x = MpPlayPrepareNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MpPlayPrepareNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpPlayPrepareNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpPlayPrepareNotify) ProtoMessage() {}

func (x *MpPlayPrepareNotify) ProtoReflect() protoreflect.Message {
	mi := &file_MpPlayPrepareNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpPlayPrepareNotify.ProtoReflect.Descriptor instead.
func (*MpPlayPrepareNotify) Descriptor() ([]byte, []int) {
	return file_MpPlayPrepareNotify_proto_rawDescGZIP(), []int{0}
}

func (x *MpPlayPrepareNotify) GetMpPlayId() uint32 {
	if x != nil {
		return x.MpPlayId
	}
	return 0
}

func (x *MpPlayPrepareNotify) GetPrepareEndTime() uint32 {
	if x != nil {
		return x.PrepareEndTime
	}
	return 0
}

var File_MpPlayPrepareNotify_proto protoreflect.FileDescriptor

var file_MpPlayPrepareNotify_proto_rawDesc = []byte{
	0x0a, 0x19, 0x4d, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x13, 0x4d, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x50, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1c, 0x0a, 0x0a, 0x6d, 0x70,
	0x5f, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x6d, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x70,
	0x61, 0x72, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x45, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0x9f, 0x0e, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e,
	0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MpPlayPrepareNotify_proto_rawDescOnce sync.Once
	file_MpPlayPrepareNotify_proto_rawDescData = file_MpPlayPrepareNotify_proto_rawDesc
)

func file_MpPlayPrepareNotify_proto_rawDescGZIP() []byte {
	file_MpPlayPrepareNotify_proto_rawDescOnce.Do(func() {
		file_MpPlayPrepareNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_MpPlayPrepareNotify_proto_rawDescData)
	})
	return file_MpPlayPrepareNotify_proto_rawDescData
}

var file_MpPlayPrepareNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MpPlayPrepareNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MpPlayPrepareNotify_proto_goTypes = []interface{}{
	(MpPlayPrepareNotify_CmdId)(0), // 0: proto.MpPlayPrepareNotify.CmdId
	(*MpPlayPrepareNotify)(nil),    // 1: proto.MpPlayPrepareNotify
}
var file_MpPlayPrepareNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MpPlayPrepareNotify_proto_init() }
func file_MpPlayPrepareNotify_proto_init() {
	if File_MpPlayPrepareNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MpPlayPrepareNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpPlayPrepareNotify); i {
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
			RawDescriptor: file_MpPlayPrepareNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MpPlayPrepareNotify_proto_goTypes,
		DependencyIndexes: file_MpPlayPrepareNotify_proto_depIdxs,
		EnumInfos:         file_MpPlayPrepareNotify_proto_enumTypes,
		MessageInfos:      file_MpPlayPrepareNotify_proto_msgTypes,
	}.Build()
	File_MpPlayPrepareNotify_proto = out.File
	file_MpPlayPrepareNotify_proto_rawDesc = nil
	file_MpPlayPrepareNotify_proto_goTypes = nil
	file_MpPlayPrepareNotify_proto_depIdxs = nil
}
