// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: ShapeBox.proto

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

type ShapeBox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Center  *Vector `protobuf:"bytes,1,opt,name=center,proto3" json:"center,omitempty"`
	Axis0   *Vector `protobuf:"bytes,2,opt,name=axis0,proto3" json:"axis0,omitempty"`
	Axis1   *Vector `protobuf:"bytes,3,opt,name=axis1,proto3" json:"axis1,omitempty"`
	Axis2   *Vector `protobuf:"bytes,4,opt,name=axis2,proto3" json:"axis2,omitempty"`
	Extents *Vector `protobuf:"bytes,5,opt,name=extents,proto3" json:"extents,omitempty"`
}

func (x *ShapeBox) Reset() {
	*x = ShapeBox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ShapeBox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShapeBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShapeBox) ProtoMessage() {}

func (x *ShapeBox) ProtoReflect() protoreflect.Message {
	mi := &file_ShapeBox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShapeBox.ProtoReflect.Descriptor instead.
func (*ShapeBox) Descriptor() ([]byte, []int) {
	return file_ShapeBox_proto_rawDescGZIP(), []int{0}
}

func (x *ShapeBox) GetCenter() *Vector {
	if x != nil {
		return x.Center
	}
	return nil
}

func (x *ShapeBox) GetAxis0() *Vector {
	if x != nil {
		return x.Axis0
	}
	return nil
}

func (x *ShapeBox) GetAxis1() *Vector {
	if x != nil {
		return x.Axis1
	}
	return nil
}

func (x *ShapeBox) GetAxis2() *Vector {
	if x != nil {
		return x.Axis2
	}
	return nil
}

func (x *ShapeBox) GetExtents() *Vector {
	if x != nil {
		return x.Extents
	}
	return nil
}

var File_ShapeBox_proto protoreflect.FileDescriptor

var file_ShapeBox_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x68, 0x61, 0x70, 0x65, 0x42, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x08, 0x53, 0x68, 0x61, 0x70, 0x65, 0x42,
	0x6f, 0x78, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x61, 0x78, 0x69,
	0x73, 0x30, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x61, 0x78, 0x69, 0x73, 0x30, 0x12, 0x23,
	0x0a, 0x05, 0x61, 0x78, 0x69, 0x73, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x61, 0x78,
	0x69, 0x73, 0x31, 0x12, 0x23, 0x0a, 0x05, 0x61, 0x78, 0x69, 0x73, 0x32, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x05, 0x61, 0x78, 0x69, 0x73, 0x32, 0x12, 0x27, 0x0a, 0x07, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ShapeBox_proto_rawDescOnce sync.Once
	file_ShapeBox_proto_rawDescData = file_ShapeBox_proto_rawDesc
)

func file_ShapeBox_proto_rawDescGZIP() []byte {
	file_ShapeBox_proto_rawDescOnce.Do(func() {
		file_ShapeBox_proto_rawDescData = protoimpl.X.CompressGZIP(file_ShapeBox_proto_rawDescData)
	})
	return file_ShapeBox_proto_rawDescData
}

var file_ShapeBox_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ShapeBox_proto_goTypes = []interface{}{
	(*ShapeBox)(nil), // 0: proto.ShapeBox
	(*Vector)(nil),   // 1: proto.Vector
}
var file_ShapeBox_proto_depIdxs = []int32{
	1, // 0: proto.ShapeBox.center:type_name -> proto.Vector
	1, // 1: proto.ShapeBox.axis0:type_name -> proto.Vector
	1, // 2: proto.ShapeBox.axis1:type_name -> proto.Vector
	1, // 3: proto.ShapeBox.axis2:type_name -> proto.Vector
	1, // 4: proto.ShapeBox.extents:type_name -> proto.Vector
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ShapeBox_proto_init() }
func file_ShapeBox_proto_init() {
	if File_ShapeBox_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ShapeBox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShapeBox); i {
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
			RawDescriptor: file_ShapeBox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ShapeBox_proto_goTypes,
		DependencyIndexes: file_ShapeBox_proto_depIdxs,
		MessageInfos:      file_ShapeBox_proto_msgTypes,
	}.Build()
	File_ShapeBox_proto = out.File
	file_ShapeBox_proto_rawDesc = nil
	file_ShapeBox_proto_goTypes = nil
	file_ShapeBox_proto_depIdxs = nil
}