// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GetCompoundDataRsp.proto

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

type GetCompoundDataRsp_CmdId int32

const (
	GetCompoundDataRsp_NONE             GetCompoundDataRsp_CmdId = 0
	GetCompoundDataRsp_CMD_ID           GetCompoundDataRsp_CmdId = 139
	GetCompoundDataRsp_ENET_CHANNEL_ID  GetCompoundDataRsp_CmdId = 0
	GetCompoundDataRsp_ENET_IS_RELIABLE GetCompoundDataRsp_CmdId = 1
)

// Enum value maps for GetCompoundDataRsp_CmdId.
var (
	GetCompoundDataRsp_CmdId_name = map[int32]string{
		0:   "NONE",
		139: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	GetCompoundDataRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           139,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x GetCompoundDataRsp_CmdId) Enum() *GetCompoundDataRsp_CmdId {
	p := new(GetCompoundDataRsp_CmdId)
	*p = x
	return p
}

func (x GetCompoundDataRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetCompoundDataRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetCompoundDataRsp_proto_enumTypes[0].Descriptor()
}

func (GetCompoundDataRsp_CmdId) Type() protoreflect.EnumType {
	return &file_GetCompoundDataRsp_proto_enumTypes[0]
}

func (x GetCompoundDataRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetCompoundDataRsp_CmdId.Descriptor instead.
func (GetCompoundDataRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetCompoundDataRsp_proto_rawDescGZIP(), []int{0, 0}
}

type GetCompoundDataRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode             uint32               `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	UnlockCompoundList  []uint32             `protobuf:"varint,2,rep,packed,name=unlock_compound_list,json=unlockCompoundList,proto3" json:"unlock_compound_list,omitempty"`
	CompoundQueDataList []*CompoundQueueData `protobuf:"bytes,3,rep,name=compound_que_data_list,json=compoundQueDataList,proto3" json:"compound_que_data_list,omitempty"`
}

func (x *GetCompoundDataRsp) Reset() {
	*x = GetCompoundDataRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetCompoundDataRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompoundDataRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompoundDataRsp) ProtoMessage() {}

func (x *GetCompoundDataRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetCompoundDataRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompoundDataRsp.ProtoReflect.Descriptor instead.
func (*GetCompoundDataRsp) Descriptor() ([]byte, []int) {
	return file_GetCompoundDataRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetCompoundDataRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetCompoundDataRsp) GetUnlockCompoundList() []uint32 {
	if x != nil {
		return x.UnlockCompoundList
	}
	return nil
}

func (x *GetCompoundDataRsp) GetCompoundQueDataList() []*CompoundQueueData {
	if x != nil {
		return x.CompoundQueDataList
	}
	return nil
}

var File_GetCompoundDataRsp_proto protoreflect.FileDescriptor

var file_GetCompoundDataRsp_proto_rawDesc = []byte{
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x75,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x75, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4d, 0x0a,
	0x16, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e,
	0x64, 0x51, 0x75, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x05,
	0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0x8b, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c,
	0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetCompoundDataRsp_proto_rawDescOnce sync.Once
	file_GetCompoundDataRsp_proto_rawDescData = file_GetCompoundDataRsp_proto_rawDesc
)

func file_GetCompoundDataRsp_proto_rawDescGZIP() []byte {
	file_GetCompoundDataRsp_proto_rawDescOnce.Do(func() {
		file_GetCompoundDataRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetCompoundDataRsp_proto_rawDescData)
	})
	return file_GetCompoundDataRsp_proto_rawDescData
}

var file_GetCompoundDataRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetCompoundDataRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetCompoundDataRsp_proto_goTypes = []interface{}{
	(GetCompoundDataRsp_CmdId)(0), // 0: proto.GetCompoundDataRsp.CmdId
	(*GetCompoundDataRsp)(nil),    // 1: proto.GetCompoundDataRsp
	(*CompoundQueueData)(nil),     // 2: proto.CompoundQueueData
}
var file_GetCompoundDataRsp_proto_depIdxs = []int32{
	2, // 0: proto.GetCompoundDataRsp.compound_que_data_list:type_name -> proto.CompoundQueueData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetCompoundDataRsp_proto_init() }
func file_GetCompoundDataRsp_proto_init() {
	if File_GetCompoundDataRsp_proto != nil {
		return
	}
	file_CompoundQueueData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetCompoundDataRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompoundDataRsp); i {
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
			RawDescriptor: file_GetCompoundDataRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetCompoundDataRsp_proto_goTypes,
		DependencyIndexes: file_GetCompoundDataRsp_proto_depIdxs,
		EnumInfos:         file_GetCompoundDataRsp_proto_enumTypes,
		MessageInfos:      file_GetCompoundDataRsp_proto_msgTypes,
	}.Build()
	File_GetCompoundDataRsp_proto = out.File
	file_GetCompoundDataRsp_proto_rawDesc = nil
	file_GetCompoundDataRsp_proto_goTypes = nil
	file_GetCompoundDataRsp_proto_depIdxs = nil
}
