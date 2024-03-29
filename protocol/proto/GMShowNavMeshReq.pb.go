// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GMShowNavMeshReq.proto

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

type GMShowNavMeshReq_CmdId int32

const (
	GMShowNavMeshReq_NONE             GMShowNavMeshReq_CmdId = 0
	GMShowNavMeshReq_CMD_ID           GMShowNavMeshReq_CmdId = 2353
	GMShowNavMeshReq_ENET_CHANNEL_ID  GMShowNavMeshReq_CmdId = 0
	GMShowNavMeshReq_ENET_IS_RELIABLE GMShowNavMeshReq_CmdId = 1
	GMShowNavMeshReq_IS_ALLOW_CLIENT  GMShowNavMeshReq_CmdId = 1
)

// Enum value maps for GMShowNavMeshReq_CmdId.
var (
	GMShowNavMeshReq_CmdId_name = map[int32]string{
		0:    "NONE",
		2353: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	GMShowNavMeshReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2353,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x GMShowNavMeshReq_CmdId) Enum() *GMShowNavMeshReq_CmdId {
	p := new(GMShowNavMeshReq_CmdId)
	*p = x
	return p
}

func (x GMShowNavMeshReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GMShowNavMeshReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GMShowNavMeshReq_proto_enumTypes[0].Descriptor()
}

func (GMShowNavMeshReq_CmdId) Type() protoreflect.EnumType {
	return &file_GMShowNavMeshReq_proto_enumTypes[0]
}

func (x GMShowNavMeshReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GMShowNavMeshReq_CmdId.Descriptor instead.
func (GMShowNavMeshReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GMShowNavMeshReq_proto_rawDescGZIP(), []int{0, 0}
}

type GMShowNavMeshReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Center *Vector `protobuf:"bytes,1,opt,name=center,proto3" json:"center,omitempty"`
	Extent *Vector `protobuf:"bytes,2,opt,name=extent,proto3" json:"extent,omitempty"`
}

func (x *GMShowNavMeshReq) Reset() {
	*x = GMShowNavMeshReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GMShowNavMeshReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GMShowNavMeshReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GMShowNavMeshReq) ProtoMessage() {}

func (x *GMShowNavMeshReq) ProtoReflect() protoreflect.Message {
	mi := &file_GMShowNavMeshReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GMShowNavMeshReq.ProtoReflect.Descriptor instead.
func (*GMShowNavMeshReq) Descriptor() ([]byte, []int) {
	return file_GMShowNavMeshReq_proto_rawDescGZIP(), []int{0}
}

func (x *GMShowNavMeshReq) GetCenter() *Vector {
	if x != nil {
		return x.Center
	}
	return nil
}

func (x *GMShowNavMeshReq) GetExtent() *Vector {
	if x != nil {
		return x.Extent
	}
	return nil
}

var File_GMShowNavMeshReq_proto protoreflect.FileDescriptor

var file_GMShowNavMeshReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x47, 0x4d, 0x53, 0x68, 0x6f, 0x77, 0x4e, 0x61, 0x76, 0x4d, 0x65, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x10, 0x47, 0x4d, 0x53, 0x68, 0x6f,
	0x77, 0x4e, 0x61, 0x76, 0x4d, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x06, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x06,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x62, 0x0a,
	0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xb1, 0x12, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45,
	0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41,
	0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GMShowNavMeshReq_proto_rawDescOnce sync.Once
	file_GMShowNavMeshReq_proto_rawDescData = file_GMShowNavMeshReq_proto_rawDesc
)

func file_GMShowNavMeshReq_proto_rawDescGZIP() []byte {
	file_GMShowNavMeshReq_proto_rawDescOnce.Do(func() {
		file_GMShowNavMeshReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GMShowNavMeshReq_proto_rawDescData)
	})
	return file_GMShowNavMeshReq_proto_rawDescData
}

var file_GMShowNavMeshReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GMShowNavMeshReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GMShowNavMeshReq_proto_goTypes = []interface{}{
	(GMShowNavMeshReq_CmdId)(0), // 0: GMShowNavMeshReq.CmdId
	(*GMShowNavMeshReq)(nil),    // 1: GMShowNavMeshReq
	(*Vector)(nil),              // 2: Vector
}
var file_GMShowNavMeshReq_proto_depIdxs = []int32{
	2, // 0: GMShowNavMeshReq.center:type_name -> Vector
	2, // 1: GMShowNavMeshReq.extent:type_name -> Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GMShowNavMeshReq_proto_init() }
func file_GMShowNavMeshReq_proto_init() {
	if File_GMShowNavMeshReq_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GMShowNavMeshReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GMShowNavMeshReq); i {
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
			RawDescriptor: file_GMShowNavMeshReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GMShowNavMeshReq_proto_goTypes,
		DependencyIndexes: file_GMShowNavMeshReq_proto_depIdxs,
		EnumInfos:         file_GMShowNavMeshReq_proto_enumTypes,
		MessageInfos:      file_GMShowNavMeshReq_proto_msgTypes,
	}.Build()
	File_GMShowNavMeshReq_proto = out.File
	file_GMShowNavMeshReq_proto_rawDesc = nil
	file_GMShowNavMeshReq_proto_goTypes = nil
	file_GMShowNavMeshReq_proto_depIdxs = nil
}
