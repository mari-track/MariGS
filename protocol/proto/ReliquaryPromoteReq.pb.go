// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: ReliquaryPromoteReq.proto

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

type ReliquaryPromoteReq_CmdId int32

const (
	ReliquaryPromoteReq_NONE             ReliquaryPromoteReq_CmdId = 0
	ReliquaryPromoteReq_CMD_ID           ReliquaryPromoteReq_CmdId = 625
	ReliquaryPromoteReq_ENET_CHANNEL_ID  ReliquaryPromoteReq_CmdId = 0
	ReliquaryPromoteReq_ENET_IS_RELIABLE ReliquaryPromoteReq_CmdId = 1
	ReliquaryPromoteReq_IS_ALLOW_CLIENT  ReliquaryPromoteReq_CmdId = 1
)

// Enum value maps for ReliquaryPromoteReq_CmdId.
var (
	ReliquaryPromoteReq_CmdId_name = map[int32]string{
		0:   "NONE",
		625: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	ReliquaryPromoteReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           625,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x ReliquaryPromoteReq_CmdId) Enum() *ReliquaryPromoteReq_CmdId {
	p := new(ReliquaryPromoteReq_CmdId)
	*p = x
	return p
}

func (x ReliquaryPromoteReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReliquaryPromoteReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ReliquaryPromoteReq_proto_enumTypes[0].Descriptor()
}

func (ReliquaryPromoteReq_CmdId) Type() protoreflect.EnumType {
	return &file_ReliquaryPromoteReq_proto_enumTypes[0]
}

func (x ReliquaryPromoteReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReliquaryPromoteReq_CmdId.Descriptor instead.
func (ReliquaryPromoteReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ReliquaryPromoteReq_proto_rawDescGZIP(), []int{0, 0}
}

type ReliquaryPromoteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetGuid uint64 `protobuf:"varint,1,opt,name=target_guid,json=targetGuid,proto3" json:"target_guid,omitempty"`
	ItemGuid   uint64 `protobuf:"varint,2,opt,name=item_guid,json=itemGuid,proto3" json:"item_guid,omitempty"`
}

func (x *ReliquaryPromoteReq) Reset() {
	*x = ReliquaryPromoteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ReliquaryPromoteReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReliquaryPromoteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReliquaryPromoteReq) ProtoMessage() {}

func (x *ReliquaryPromoteReq) ProtoReflect() protoreflect.Message {
	mi := &file_ReliquaryPromoteReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReliquaryPromoteReq.ProtoReflect.Descriptor instead.
func (*ReliquaryPromoteReq) Descriptor() ([]byte, []int) {
	return file_ReliquaryPromoteReq_proto_rawDescGZIP(), []int{0}
}

func (x *ReliquaryPromoteReq) GetTargetGuid() uint64 {
	if x != nil {
		return x.TargetGuid
	}
	return 0
}

func (x *ReliquaryPromoteReq) GetItemGuid() uint64 {
	if x != nil {
		return x.ItemGuid
	}
	return 0
}

var File_ReliquaryPromoteReq_proto protoreflect.FileDescriptor

var file_ReliquaryPromoteReq_proto_rawDesc = []byte{
	0x0a, 0x19, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x69, 0x71, 0x75, 0x61, 0x72, 0x79,
	0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x47, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x69, 0x74, 0x65, 0x6d, 0x47, 0x75, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49,
	0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43,
	0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xf1, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f,
	0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ReliquaryPromoteReq_proto_rawDescOnce sync.Once
	file_ReliquaryPromoteReq_proto_rawDescData = file_ReliquaryPromoteReq_proto_rawDesc
)

func file_ReliquaryPromoteReq_proto_rawDescGZIP() []byte {
	file_ReliquaryPromoteReq_proto_rawDescOnce.Do(func() {
		file_ReliquaryPromoteReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ReliquaryPromoteReq_proto_rawDescData)
	})
	return file_ReliquaryPromoteReq_proto_rawDescData
}

var file_ReliquaryPromoteReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ReliquaryPromoteReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ReliquaryPromoteReq_proto_goTypes = []interface{}{
	(ReliquaryPromoteReq_CmdId)(0), // 0: proto.ReliquaryPromoteReq.CmdId
	(*ReliquaryPromoteReq)(nil),    // 1: proto.ReliquaryPromoteReq
}
var file_ReliquaryPromoteReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ReliquaryPromoteReq_proto_init() }
func file_ReliquaryPromoteReq_proto_init() {
	if File_ReliquaryPromoteReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ReliquaryPromoteReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReliquaryPromoteReq); i {
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
			RawDescriptor: file_ReliquaryPromoteReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ReliquaryPromoteReq_proto_goTypes,
		DependencyIndexes: file_ReliquaryPromoteReq_proto_depIdxs,
		EnumInfos:         file_ReliquaryPromoteReq_proto_enumTypes,
		MessageInfos:      file_ReliquaryPromoteReq_proto_msgTypes,
	}.Build()
	File_ReliquaryPromoteReq_proto = out.File
	file_ReliquaryPromoteReq_proto_rawDesc = nil
	file_ReliquaryPromoteReq_proto_goTypes = nil
	file_ReliquaryPromoteReq_proto_depIdxs = nil
}