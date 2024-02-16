// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: ShapeSphere.proto

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

type ShapeSphere struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Center *Vector `protobuf:"bytes,1,opt,name=center,proto3" json:"center,omitempty"`
	Radius float32 `protobuf:"fixed32,2,opt,name=radius,proto3" json:"radius,omitempty"`
}

func (x *ShapeSphere) Reset() {
	*x = ShapeSphere{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ShapeSphere_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShapeSphere) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShapeSphere) ProtoMessage() {}

func (x *ShapeSphere) ProtoReflect() protoreflect.Message {
	mi := &file_ShapeSphere_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShapeSphere.ProtoReflect.Descriptor instead.
func (*ShapeSphere) Descriptor() ([]byte, []int) {
	return file_ShapeSphere_proto_rawDescGZIP(), []int{0}
}

func (x *ShapeSphere) GetCenter() *Vector {
	if x != nil {
		return x.Center
	}
	return nil
}

func (x *ShapeSphere) GetRadius() float32 {
	if x != nil {
		return x.Radius
	}
	return 0
}

var File_ShapeSphere_proto protoreflect.FileDescriptor

var file_ShapeSphere_proto_rawDesc = []byte{
	0x0a, 0x11, 0x53, 0x68, 0x61, 0x70, 0x65, 0x53, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x0b, 0x53, 0x68, 0x61, 0x70,
	0x65, 0x53, 0x70, 0x68, 0x65, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ShapeSphere_proto_rawDescOnce sync.Once
	file_ShapeSphere_proto_rawDescData = file_ShapeSphere_proto_rawDesc
)

func file_ShapeSphere_proto_rawDescGZIP() []byte {
	file_ShapeSphere_proto_rawDescOnce.Do(func() {
		file_ShapeSphere_proto_rawDescData = protoimpl.X.CompressGZIP(file_ShapeSphere_proto_rawDescData)
	})
	return file_ShapeSphere_proto_rawDescData
}

var file_ShapeSphere_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ShapeSphere_proto_goTypes = []interface{}{
	(*ShapeSphere)(nil), // 0: proto.ShapeSphere
	(*Vector)(nil),      // 1: proto.Vector
}
var file_ShapeSphere_proto_depIdxs = []int32{
	1, // 0: proto.ShapeSphere.center:type_name -> proto.Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ShapeSphere_proto_init() }
func file_ShapeSphere_proto_init() {
	if File_ShapeSphere_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ShapeSphere_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShapeSphere); i {
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
			RawDescriptor: file_ShapeSphere_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ShapeSphere_proto_goTypes,
		DependencyIndexes: file_ShapeSphere_proto_depIdxs,
		MessageInfos:      file_ShapeSphere_proto_msgTypes,
	}.Build()
	File_ShapeSphere_proto = out.File
	file_ShapeSphere_proto_rawDesc = nil
	file_ShapeSphere_proto_goTypes = nil
	file_ShapeSphere_proto_depIdxs = nil
}
