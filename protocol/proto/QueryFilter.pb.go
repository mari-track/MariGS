// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: QueryFilter.proto

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

type QueryFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeId   int32 `protobuf:"varint,1,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	AreaMask int32 `protobuf:"varint,2,opt,name=area_mask,json=areaMask,proto3" json:"area_mask,omitempty"`
}

func (x *QueryFilter) Reset() {
	*x = QueryFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QueryFilter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryFilter) ProtoMessage() {}

func (x *QueryFilter) ProtoReflect() protoreflect.Message {
	mi := &file_QueryFilter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryFilter.ProtoReflect.Descriptor instead.
func (*QueryFilter) Descriptor() ([]byte, []int) {
	return file_QueryFilter_proto_rawDescGZIP(), []int{0}
}

func (x *QueryFilter) GetTypeId() int32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *QueryFilter) GetAreaMask() int32 {
	if x != nil {
		return x.AreaMask
	}
	return 0
}

var File_QueryFilter_proto protoreflect.FileDescriptor

var file_QueryFilter_proto_rawDesc = []byte{
	0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0b, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x4d, 0x61, 0x73, 0x6b, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_QueryFilter_proto_rawDescOnce sync.Once
	file_QueryFilter_proto_rawDescData = file_QueryFilter_proto_rawDesc
)

func file_QueryFilter_proto_rawDescGZIP() []byte {
	file_QueryFilter_proto_rawDescOnce.Do(func() {
		file_QueryFilter_proto_rawDescData = protoimpl.X.CompressGZIP(file_QueryFilter_proto_rawDescData)
	})
	return file_QueryFilter_proto_rawDescData
}

var file_QueryFilter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_QueryFilter_proto_goTypes = []interface{}{
	(*QueryFilter)(nil), // 0: proto.QueryFilter
}
var file_QueryFilter_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_QueryFilter_proto_init() }
func file_QueryFilter_proto_init() {
	if File_QueryFilter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_QueryFilter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryFilter); i {
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
			RawDescriptor: file_QueryFilter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_QueryFilter_proto_goTypes,
		DependencyIndexes: file_QueryFilter_proto_depIdxs,
		MessageInfos:      file_QueryFilter_proto_msgTypes,
	}.Build()
	File_QueryFilter_proto = out.File
	file_QueryFilter_proto_rawDesc = nil
	file_QueryFilter_proto_goTypes = nil
	file_QueryFilter_proto_depIdxs = nil
}
