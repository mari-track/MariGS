// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: AbilityInvocationFixedNotify.proto

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

type AbilityInvocationFixedNotify_CmdId int32

const (
	AbilityInvocationFixedNotify_NONE             AbilityInvocationFixedNotify_CmdId = 0
	AbilityInvocationFixedNotify_CMD_ID           AbilityInvocationFixedNotify_CmdId = 1101
	AbilityInvocationFixedNotify_ENET_CHANNEL_ID  AbilityInvocationFixedNotify_CmdId = 0
	AbilityInvocationFixedNotify_ENET_IS_RELIABLE AbilityInvocationFixedNotify_CmdId = 1
	AbilityInvocationFixedNotify_IS_ALLOW_CLIENT  AbilityInvocationFixedNotify_CmdId = 1
)

// Enum value maps for AbilityInvocationFixedNotify_CmdId.
var (
	AbilityInvocationFixedNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		1101: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	AbilityInvocationFixedNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1101,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x AbilityInvocationFixedNotify_CmdId) Enum() *AbilityInvocationFixedNotify_CmdId {
	p := new(AbilityInvocationFixedNotify_CmdId)
	*p = x
	return p
}

func (x AbilityInvocationFixedNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AbilityInvocationFixedNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_AbilityInvocationFixedNotify_proto_enumTypes[0].Descriptor()
}

func (AbilityInvocationFixedNotify_CmdId) Type() protoreflect.EnumType {
	return &file_AbilityInvocationFixedNotify_proto_enumTypes[0]
}

func (x AbilityInvocationFixedNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AbilityInvocationFixedNotify_CmdId.Descriptor instead.
func (AbilityInvocationFixedNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_AbilityInvocationFixedNotify_proto_rawDescGZIP(), []int{0, 0}
}

type AbilityInvocationFixedNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invoke1St *AbilityInvokeEntry `protobuf:"bytes,1,opt,name=invoke1st,proto3" json:"invoke1st,omitempty"`
	Invoke2Nd *AbilityInvokeEntry `protobuf:"bytes,2,opt,name=invoke2nd,proto3" json:"invoke2nd,omitempty"`
	Invoke3Rd *AbilityInvokeEntry `protobuf:"bytes,3,opt,name=invoke3rd,proto3" json:"invoke3rd,omitempty"`
	Invoke4Th *AbilityInvokeEntry `protobuf:"bytes,4,opt,name=invoke4th,proto3" json:"invoke4th,omitempty"`
	Invoke5Th *AbilityInvokeEntry `protobuf:"bytes,5,opt,name=invoke5th,proto3" json:"invoke5th,omitempty"`
	Invoke6Th *AbilityInvokeEntry `protobuf:"bytes,6,opt,name=invoke6th,proto3" json:"invoke6th,omitempty"`
}

func (x *AbilityInvocationFixedNotify) Reset() {
	*x = AbilityInvocationFixedNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AbilityInvocationFixedNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityInvocationFixedNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityInvocationFixedNotify) ProtoMessage() {}

func (x *AbilityInvocationFixedNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AbilityInvocationFixedNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityInvocationFixedNotify.ProtoReflect.Descriptor instead.
func (*AbilityInvocationFixedNotify) Descriptor() ([]byte, []int) {
	return file_AbilityInvocationFixedNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AbilityInvocationFixedNotify) GetInvoke1St() *AbilityInvokeEntry {
	if x != nil {
		return x.Invoke1St
	}
	return nil
}

func (x *AbilityInvocationFixedNotify) GetInvoke2Nd() *AbilityInvokeEntry {
	if x != nil {
		return x.Invoke2Nd
	}
	return nil
}

func (x *AbilityInvocationFixedNotify) GetInvoke3Rd() *AbilityInvokeEntry {
	if x != nil {
		return x.Invoke3Rd
	}
	return nil
}

func (x *AbilityInvocationFixedNotify) GetInvoke4Th() *AbilityInvokeEntry {
	if x != nil {
		return x.Invoke4Th
	}
	return nil
}

func (x *AbilityInvocationFixedNotify) GetInvoke5Th() *AbilityInvokeEntry {
	if x != nil {
		return x.Invoke5Th
	}
	return nil
}

func (x *AbilityInvocationFixedNotify) GetInvoke6Th() *AbilityInvokeEntry {
	if x != nil {
		return x.Invoke6Th
	}
	return nil
}

var File_AbilityInvocationFixedNotify_proto protoreflect.FileDescriptor

var file_AbilityInvocationFixedNotify_proto_rawDesc = []byte{
	0x0a, 0x22, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x69, 0x78, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76,
	0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4,
	0x03, 0x0a, 0x1c, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x78, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x31, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x31, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x31,
	0x73, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x32, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49,
	0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x32, 0x6e, 0x64, 0x12, 0x31, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x33,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x69,
	0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x33, 0x72, 0x64, 0x12, 0x31, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x34, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x34, 0x74, 0x68, 0x12, 0x31, 0x0a, 0x09, 0x69,
	0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x35, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x35, 0x74, 0x68, 0x12, 0x31,
	0x0a, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x36, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x36, 0x74,
	0x68, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xcd,
	0x08, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45,
	0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49,
	0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f,
	0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10,
	0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AbilityInvocationFixedNotify_proto_rawDescOnce sync.Once
	file_AbilityInvocationFixedNotify_proto_rawDescData = file_AbilityInvocationFixedNotify_proto_rawDesc
)

func file_AbilityInvocationFixedNotify_proto_rawDescGZIP() []byte {
	file_AbilityInvocationFixedNotify_proto_rawDescOnce.Do(func() {
		file_AbilityInvocationFixedNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilityInvocationFixedNotify_proto_rawDescData)
	})
	return file_AbilityInvocationFixedNotify_proto_rawDescData
}

var file_AbilityInvocationFixedNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AbilityInvocationFixedNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AbilityInvocationFixedNotify_proto_goTypes = []interface{}{
	(AbilityInvocationFixedNotify_CmdId)(0), // 0: AbilityInvocationFixedNotify.CmdId
	(*AbilityInvocationFixedNotify)(nil),    // 1: AbilityInvocationFixedNotify
	(*AbilityInvokeEntry)(nil),              // 2: AbilityInvokeEntry
}
var file_AbilityInvocationFixedNotify_proto_depIdxs = []int32{
	2, // 0: AbilityInvocationFixedNotify.invoke1st:type_name -> AbilityInvokeEntry
	2, // 1: AbilityInvocationFixedNotify.invoke2nd:type_name -> AbilityInvokeEntry
	2, // 2: AbilityInvocationFixedNotify.invoke3rd:type_name -> AbilityInvokeEntry
	2, // 3: AbilityInvocationFixedNotify.invoke4th:type_name -> AbilityInvokeEntry
	2, // 4: AbilityInvocationFixedNotify.invoke5th:type_name -> AbilityInvokeEntry
	2, // 5: AbilityInvocationFixedNotify.invoke6th:type_name -> AbilityInvokeEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_AbilityInvocationFixedNotify_proto_init() }
func file_AbilityInvocationFixedNotify_proto_init() {
	if File_AbilityInvocationFixedNotify_proto != nil {
		return
	}
	file_AbilityInvokeEntry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AbilityInvocationFixedNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityInvocationFixedNotify); i {
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
			RawDescriptor: file_AbilityInvocationFixedNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilityInvocationFixedNotify_proto_goTypes,
		DependencyIndexes: file_AbilityInvocationFixedNotify_proto_depIdxs,
		EnumInfos:         file_AbilityInvocationFixedNotify_proto_enumTypes,
		MessageInfos:      file_AbilityInvocationFixedNotify_proto_msgTypes,
	}.Build()
	File_AbilityInvocationFixedNotify_proto = out.File
	file_AbilityInvocationFixedNotify_proto_rawDesc = nil
	file_AbilityInvocationFixedNotify_proto_goTypes = nil
	file_AbilityInvocationFixedNotify_proto_depIdxs = nil
}
