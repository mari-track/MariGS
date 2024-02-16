// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: MassiveBoxInfo.proto

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

type MassiveBoxInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ConfigId uint32  `protobuf:"varint,2,opt,name=configId,proto3" json:"configId,omitempty"`
	Center   *Vector `protobuf:"bytes,3,opt,name=center,proto3" json:"center,omitempty"`
	Extents  *Vector `protobuf:"bytes,4,opt,name=extents,proto3" json:"extents,omitempty"`
	Up       *Vector `protobuf:"bytes,5,opt,name=up,proto3" json:"up,omitempty"`
	Forward  *Vector `protobuf:"bytes,6,opt,name=forward,proto3" json:"forward,omitempty"`
	Right    *Vector `protobuf:"bytes,7,opt,name=right,proto3" json:"right,omitempty"`
}

func (x *MassiveBoxInfo) Reset() {
	*x = MassiveBoxInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MassiveBoxInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MassiveBoxInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MassiveBoxInfo) ProtoMessage() {}

func (x *MassiveBoxInfo) ProtoReflect() protoreflect.Message {
	mi := &file_MassiveBoxInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MassiveBoxInfo.ProtoReflect.Descriptor instead.
func (*MassiveBoxInfo) Descriptor() ([]byte, []int) {
	return file_MassiveBoxInfo_proto_rawDescGZIP(), []int{0}
}

func (x *MassiveBoxInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MassiveBoxInfo) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *MassiveBoxInfo) GetCenter() *Vector {
	if x != nil {
		return x.Center
	}
	return nil
}

func (x *MassiveBoxInfo) GetExtents() *Vector {
	if x != nil {
		return x.Extents
	}
	return nil
}

func (x *MassiveBoxInfo) GetUp() *Vector {
	if x != nil {
		return x.Up
	}
	return nil
}

func (x *MassiveBoxInfo) GetForward() *Vector {
	if x != nil {
		return x.Forward
	}
	return nil
}

func (x *MassiveBoxInfo) GetRight() *Vector {
	if x != nil {
		return x.Right
	}
	return nil
}

var File_MassiveBoxInfo_proto protoreflect.FileDescriptor

var file_MassiveBoxInfo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x42, 0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x01, 0x0a, 0x0e,
	0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x42, 0x6f, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x12, 0x27, 0x0a, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x02, 0x75, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x02, 0x75, 0x70, 0x12, 0x27, 0x0a, 0x07, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x12, 0x23, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MassiveBoxInfo_proto_rawDescOnce sync.Once
	file_MassiveBoxInfo_proto_rawDescData = file_MassiveBoxInfo_proto_rawDesc
)

func file_MassiveBoxInfo_proto_rawDescGZIP() []byte {
	file_MassiveBoxInfo_proto_rawDescOnce.Do(func() {
		file_MassiveBoxInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_MassiveBoxInfo_proto_rawDescData)
	})
	return file_MassiveBoxInfo_proto_rawDescData
}

var file_MassiveBoxInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MassiveBoxInfo_proto_goTypes = []interface{}{
	(*MassiveBoxInfo)(nil), // 0: proto.MassiveBoxInfo
	(*Vector)(nil),         // 1: proto.Vector
}
var file_MassiveBoxInfo_proto_depIdxs = []int32{
	1, // 0: proto.MassiveBoxInfo.center:type_name -> proto.Vector
	1, // 1: proto.MassiveBoxInfo.extents:type_name -> proto.Vector
	1, // 2: proto.MassiveBoxInfo.up:type_name -> proto.Vector
	1, // 3: proto.MassiveBoxInfo.forward:type_name -> proto.Vector
	1, // 4: proto.MassiveBoxInfo.right:type_name -> proto.Vector
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_MassiveBoxInfo_proto_init() }
func file_MassiveBoxInfo_proto_init() {
	if File_MassiveBoxInfo_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MassiveBoxInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MassiveBoxInfo); i {
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
			RawDescriptor: file_MassiveBoxInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MassiveBoxInfo_proto_goTypes,
		DependencyIndexes: file_MassiveBoxInfo_proto_depIdxs,
		MessageInfos:      file_MassiveBoxInfo_proto_msgTypes,
	}.Build()
	File_MassiveBoxInfo_proto = out.File
	file_MassiveBoxInfo_proto_rawDesc = nil
	file_MassiveBoxInfo_proto_goTypes = nil
	file_MassiveBoxInfo_proto_depIdxs = nil
}
