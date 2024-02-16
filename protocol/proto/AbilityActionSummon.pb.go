// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: AbilityActionSummon.proto

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

type AbilityActionSummon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos *Vector `protobuf:"bytes,1,opt,name=pos,proto3" json:"pos,omitempty"`
	Rot *Vector `protobuf:"bytes,2,opt,name=rot,proto3" json:"rot,omitempty"`
}

func (x *AbilityActionSummon) Reset() {
	*x = AbilityActionSummon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AbilityActionSummon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityActionSummon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityActionSummon) ProtoMessage() {}

func (x *AbilityActionSummon) ProtoReflect() protoreflect.Message {
	mi := &file_AbilityActionSummon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityActionSummon.ProtoReflect.Descriptor instead.
func (*AbilityActionSummon) Descriptor() ([]byte, []int) {
	return file_AbilityActionSummon_proto_rawDescGZIP(), []int{0}
}

func (x *AbilityActionSummon) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *AbilityActionSummon) GetRot() *Vector {
	if x != nil {
		return x.Rot
	}
	return nil
}

var File_AbilityActionSummon_proto protoreflect.FileDescriptor

var file_AbilityActionSummon_proto_rawDesc = []byte{
	0x0a, 0x19, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x57, 0x0a, 0x13, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x03, 0x72, 0x6f, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x72, 0x6f, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AbilityActionSummon_proto_rawDescOnce sync.Once
	file_AbilityActionSummon_proto_rawDescData = file_AbilityActionSummon_proto_rawDesc
)

func file_AbilityActionSummon_proto_rawDescGZIP() []byte {
	file_AbilityActionSummon_proto_rawDescOnce.Do(func() {
		file_AbilityActionSummon_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilityActionSummon_proto_rawDescData)
	})
	return file_AbilityActionSummon_proto_rawDescData
}

var file_AbilityActionSummon_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AbilityActionSummon_proto_goTypes = []interface{}{
	(*AbilityActionSummon)(nil), // 0: proto.AbilityActionSummon
	(*Vector)(nil),              // 1: proto.Vector
}
var file_AbilityActionSummon_proto_depIdxs = []int32{
	1, // 0: proto.AbilityActionSummon.pos:type_name -> proto.Vector
	1, // 1: proto.AbilityActionSummon.rot:type_name -> proto.Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AbilityActionSummon_proto_init() }
func file_AbilityActionSummon_proto_init() {
	if File_AbilityActionSummon_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AbilityActionSummon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityActionSummon); i {
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
			RawDescriptor: file_AbilityActionSummon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilityActionSummon_proto_goTypes,
		DependencyIndexes: file_AbilityActionSummon_proto_depIdxs,
		MessageInfos:      file_AbilityActionSummon_proto_msgTypes,
	}.Build()
	File_AbilityActionSummon_proto = out.File
	file_AbilityActionSummon_proto_rawDesc = nil
	file_AbilityActionSummon_proto_goTypes = nil
	file_AbilityActionSummon_proto_depIdxs = nil
}
