// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: Investigation.proto

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

type Investigation_State int32

const (
	Investigation_INVALID      Investigation_State = 0
	Investigation_IN_PROGRESS  Investigation_State = 1
	Investigation_COMPLETE     Investigation_State = 2
	Investigation_REWARD_TAKEN Investigation_State = 3
)

// Enum value maps for Investigation_State.
var (
	Investigation_State_name = map[int32]string{
		0: "INVALID",
		1: "IN_PROGRESS",
		2: "COMPLETE",
		3: "REWARD_TAKEN",
	}
	Investigation_State_value = map[string]int32{
		"INVALID":      0,
		"IN_PROGRESS":  1,
		"COMPLETE":     2,
		"REWARD_TAKEN": 3,
	}
)

func (x Investigation_State) Enum() *Investigation_State {
	p := new(Investigation_State)
	*p = x
	return p
}

func (x Investigation_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Investigation_State) Descriptor() protoreflect.EnumDescriptor {
	return file_Investigation_proto_enumTypes[0].Descriptor()
}

func (Investigation_State) Type() protoreflect.EnumType {
	return &file_Investigation_proto_enumTypes[0]
}

func (x Investigation_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Investigation_State.Descriptor instead.
func (Investigation_State) EnumDescriptor() ([]byte, []int) {
	return file_Investigation_proto_rawDescGZIP(), []int{0, 0}
}

type Investigation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint32              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Progress      uint32              `protobuf:"varint,2,opt,name=progress,proto3" json:"progress,omitempty"`
	TotalProgress uint32              `protobuf:"varint,3,opt,name=total_progress,json=totalProgress,proto3" json:"total_progress,omitempty"`
	State         Investigation_State `protobuf:"varint,4,opt,name=state,proto3,enum=proto.Investigation_State" json:"state,omitempty"`
}

func (x *Investigation) Reset() {
	*x = Investigation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Investigation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Investigation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Investigation) ProtoMessage() {}

func (x *Investigation) ProtoReflect() protoreflect.Message {
	mi := &file_Investigation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Investigation.ProtoReflect.Descriptor instead.
func (*Investigation) Descriptor() ([]byte, []int) {
	return file_Investigation_proto_rawDescGZIP(), []int{0}
}

func (x *Investigation) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Investigation) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *Investigation) GetTotalProgress() uint32 {
	if x != nil {
		return x.TotalProgress
	}
	return 0
}

func (x *Investigation) GetState() Investigation_State {
	if x != nil {
		return x.State
	}
	return Investigation_INVALID
}

var File_Investigation_proto protoreflect.FileDescriptor

var file_Investigation_proto_rawDesc = []byte{
	0x0a, 0x13, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a,
	0x0d, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x22, 0x45, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f,
	0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x57, 0x41,
	0x52, 0x44, 0x5f, 0x54, 0x41, 0x4b, 0x45, 0x4e, 0x10, 0x03, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Investigation_proto_rawDescOnce sync.Once
	file_Investigation_proto_rawDescData = file_Investigation_proto_rawDesc
)

func file_Investigation_proto_rawDescGZIP() []byte {
	file_Investigation_proto_rawDescOnce.Do(func() {
		file_Investigation_proto_rawDescData = protoimpl.X.CompressGZIP(file_Investigation_proto_rawDescData)
	})
	return file_Investigation_proto_rawDescData
}

var file_Investigation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Investigation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Investigation_proto_goTypes = []interface{}{
	(Investigation_State)(0), // 0: proto.Investigation.State
	(*Investigation)(nil),    // 1: proto.Investigation
}
var file_Investigation_proto_depIdxs = []int32{
	0, // 0: proto.Investigation.state:type_name -> proto.Investigation.State
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Investigation_proto_init() }
func file_Investigation_proto_init() {
	if File_Investigation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Investigation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Investigation); i {
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
			RawDescriptor: file_Investigation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Investigation_proto_goTypes,
		DependencyIndexes: file_Investigation_proto_depIdxs,
		EnumInfos:         file_Investigation_proto_enumTypes,
		MessageInfos:      file_Investigation_proto_msgTypes,
	}.Build()
	File_Investigation_proto = out.File
	file_Investigation_proto_rawDesc = nil
	file_Investigation_proto_goTypes = nil
	file_Investigation_proto_depIdxs = nil
}
