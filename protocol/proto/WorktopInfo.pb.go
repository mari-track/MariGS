// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: WorktopInfo.proto

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

type WorktopInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptionList        []uint32 `protobuf:"varint,1,rep,packed,name=option_list,json=optionList,proto3" json:"option_list,omitempty"`
	IsGuestCanOperate bool     `protobuf:"varint,2,opt,name=is_guest_can_operate,json=isGuestCanOperate,proto3" json:"is_guest_can_operate,omitempty"`
}

func (x *WorktopInfo) Reset() {
	*x = WorktopInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorktopInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorktopInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorktopInfo) ProtoMessage() {}

func (x *WorktopInfo) ProtoReflect() protoreflect.Message {
	mi := &file_WorktopInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorktopInfo.ProtoReflect.Descriptor instead.
func (*WorktopInfo) Descriptor() ([]byte, []int) {
	return file_WorktopInfo_proto_rawDescGZIP(), []int{0}
}

func (x *WorktopInfo) GetOptionList() []uint32 {
	if x != nil {
		return x.OptionList
	}
	return nil
}

func (x *WorktopInfo) GetIsGuestCanOperate() bool {
	if x != nil {
		return x.IsGuestCanOperate
	}
	return false
}

var File_WorktopInfo_proto protoreflect.FileDescriptor

var file_WorktopInfo_proto_rawDesc = []byte{
	0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x74, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x0b, 0x57, 0x6f,
	0x72, 0x6b, 0x74, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x69, 0x73,
	0x5f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x47, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WorktopInfo_proto_rawDescOnce sync.Once
	file_WorktopInfo_proto_rawDescData = file_WorktopInfo_proto_rawDesc
)

func file_WorktopInfo_proto_rawDescGZIP() []byte {
	file_WorktopInfo_proto_rawDescOnce.Do(func() {
		file_WorktopInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_WorktopInfo_proto_rawDescData)
	})
	return file_WorktopInfo_proto_rawDescData
}

var file_WorktopInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WorktopInfo_proto_goTypes = []interface{}{
	(*WorktopInfo)(nil), // 0: proto.WorktopInfo
}
var file_WorktopInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_WorktopInfo_proto_init() }
func file_WorktopInfo_proto_init() {
	if File_WorktopInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_WorktopInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorktopInfo); i {
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
			RawDescriptor: file_WorktopInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WorktopInfo_proto_goTypes,
		DependencyIndexes: file_WorktopInfo_proto_depIdxs,
		MessageInfos:      file_WorktopInfo_proto_msgTypes,
	}.Build()
	File_WorktopInfo_proto = out.File
	file_WorktopInfo_proto_rawDesc = nil
	file_WorktopInfo_proto_goTypes = nil
	file_WorktopInfo_proto_depIdxs = nil
}
