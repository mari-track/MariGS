// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: DropItemReq.proto

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

type DropItemReq_CmdId int32

const (
	DropItemReq_NONE             DropItemReq_CmdId = 0
	DropItemReq_CMD_ID           DropItemReq_CmdId = 610
	DropItemReq_ENET_CHANNEL_ID  DropItemReq_CmdId = 0
	DropItemReq_ENET_IS_RELIABLE DropItemReq_CmdId = 1
	DropItemReq_IS_ALLOW_CLIENT  DropItemReq_CmdId = 1
)

// Enum value maps for DropItemReq_CmdId.
var (
	DropItemReq_CmdId_name = map[int32]string{
		0:   "NONE",
		610: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	DropItemReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           610,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x DropItemReq_CmdId) Enum() *DropItemReq_CmdId {
	p := new(DropItemReq_CmdId)
	*p = x
	return p
}

func (x DropItemReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DropItemReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_DropItemReq_proto_enumTypes[0].Descriptor()
}

func (DropItemReq_CmdId) Type() protoreflect.EnumType {
	return &file_DropItemReq_proto_enumTypes[0]
}

func (x DropItemReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DropItemReq_CmdId.Descriptor instead.
func (DropItemReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_DropItemReq_proto_rawDescGZIP(), []int{0, 0}
}

type DropItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreType StoreType `protobuf:"varint,1,opt,name=store_type,json=storeType,proto3,enum=StoreType" json:"store_type,omitempty"`
	Guid      uint64    `protobuf:"varint,2,opt,name=guid,proto3" json:"guid,omitempty"`
	Count     uint32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Pos       *Vector   `protobuf:"bytes,4,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *DropItemReq) Reset() {
	*x = DropItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DropItemReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropItemReq) ProtoMessage() {}

func (x *DropItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_DropItemReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropItemReq.ProtoReflect.Descriptor instead.
func (*DropItemReq) Descriptor() ([]byte, []int) {
	return file_DropItemReq_proto_rawDescGZIP(), []int{0}
}

func (x *DropItemReq) GetStoreType() StoreType {
	if x != nil {
		return x.StoreType
	}
	return StoreType_STORE_NONE
}

func (x *DropItemReq) GetGuid() uint64 {
	if x != nil {
		return x.Guid
	}
	return 0
}

func (x *DropItemReq) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *DropItemReq) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

var File_DropItemReq_proto protoreflect.FileDescriptor

var file_DropItemReq_proto_rawDesc = []byte{
	0x0a, 0x11, 0x44, 0x72, 0x6f, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x0b, 0x44, 0x72, 0x6f, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x12, 0x29, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x67, 0x75, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x70,
	0x6f, 0x73, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10,
	0xe2, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e,
	0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f,
	0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54,
	0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DropItemReq_proto_rawDescOnce sync.Once
	file_DropItemReq_proto_rawDescData = file_DropItemReq_proto_rawDesc
)

func file_DropItemReq_proto_rawDescGZIP() []byte {
	file_DropItemReq_proto_rawDescOnce.Do(func() {
		file_DropItemReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_DropItemReq_proto_rawDescData)
	})
	return file_DropItemReq_proto_rawDescData
}

var file_DropItemReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DropItemReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DropItemReq_proto_goTypes = []interface{}{
	(DropItemReq_CmdId)(0), // 0: DropItemReq.CmdId
	(*DropItemReq)(nil),    // 1: DropItemReq
	(StoreType)(0),         // 2: StoreType
	(*Vector)(nil),         // 3: Vector
}
var file_DropItemReq_proto_depIdxs = []int32{
	2, // 0: DropItemReq.store_type:type_name -> StoreType
	3, // 1: DropItemReq.pos:type_name -> Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_DropItemReq_proto_init() }
func file_DropItemReq_proto_init() {
	if File_DropItemReq_proto != nil {
		return
	}
	file_Vector_proto_init()
	file_StoreType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DropItemReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropItemReq); i {
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
			RawDescriptor: file_DropItemReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DropItemReq_proto_goTypes,
		DependencyIndexes: file_DropItemReq_proto_depIdxs,
		EnumInfos:         file_DropItemReq_proto_enumTypes,
		MessageInfos:      file_DropItemReq_proto_msgTypes,
	}.Build()
	File_DropItemReq_proto = out.File
	file_DropItemReq_proto_rawDesc = nil
	file_DropItemReq_proto_goTypes = nil
	file_DropItemReq_proto_depIdxs = nil
}
