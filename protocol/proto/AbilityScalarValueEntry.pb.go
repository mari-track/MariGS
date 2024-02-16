// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: AbilityScalarValueEntry.proto

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

type AbilityScalarValueEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       *AbilityString    `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ValueType AbilityScalarType `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=proto.AbilityScalarType" json:"value_type,omitempty"`
	// Types that are assignable to Value:
	//
	//	*AbilityScalarValueEntry_FloatValue
	//	*AbilityScalarValueEntry_StringValue
	//	*AbilityScalarValueEntry_IntValue
	//	*AbilityScalarValueEntry_UintValue
	Value isAbilityScalarValueEntry_Value `protobuf_oneof:"value"`
}

func (x *AbilityScalarValueEntry) Reset() {
	*x = AbilityScalarValueEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AbilityScalarValueEntry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityScalarValueEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityScalarValueEntry) ProtoMessage() {}

func (x *AbilityScalarValueEntry) ProtoReflect() protoreflect.Message {
	mi := &file_AbilityScalarValueEntry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityScalarValueEntry.ProtoReflect.Descriptor instead.
func (*AbilityScalarValueEntry) Descriptor() ([]byte, []int) {
	return file_AbilityScalarValueEntry_proto_rawDescGZIP(), []int{0}
}

func (x *AbilityScalarValueEntry) GetKey() *AbilityString {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *AbilityScalarValueEntry) GetValueType() AbilityScalarType {
	if x != nil {
		return x.ValueType
	}
	return AbilityScalarType_ABILITY_SCALAR_TYPE_UNKNOW
}

func (m *AbilityScalarValueEntry) GetValue() isAbilityScalarValueEntry_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *AbilityScalarValueEntry) GetFloatValue() float32 {
	if x, ok := x.GetValue().(*AbilityScalarValueEntry_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (x *AbilityScalarValueEntry) GetStringValue() string {
	if x, ok := x.GetValue().(*AbilityScalarValueEntry_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *AbilityScalarValueEntry) GetIntValue() int32 {
	if x, ok := x.GetValue().(*AbilityScalarValueEntry_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *AbilityScalarValueEntry) GetUintValue() uint32 {
	if x, ok := x.GetValue().(*AbilityScalarValueEntry_UintValue); ok {
		return x.UintValue
	}
	return 0
}

type isAbilityScalarValueEntry_Value interface {
	isAbilityScalarValueEntry_Value()
}

type AbilityScalarValueEntry_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,3,opt,name=float_value,json=floatValue,proto3,oneof"`
}

type AbilityScalarValueEntry_StringValue struct {
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type AbilityScalarValueEntry_IntValue struct {
	IntValue int32 `protobuf:"varint,5,opt,name=int_value,json=intValue,proto3,oneof"`
}

type AbilityScalarValueEntry_UintValue struct {
	UintValue uint32 `protobuf:"varint,6,opt,name=uint_value,json=uintValue,proto3,oneof"`
}

func (*AbilityScalarValueEntry_FloatValue) isAbilityScalarValueEntry_Value() {}

func (*AbilityScalarValueEntry_StringValue) isAbilityScalarValueEntry_Value() {}

func (*AbilityScalarValueEntry_IntValue) isAbilityScalarValueEntry_Value() {}

func (*AbilityScalarValueEntry_UintValue) isAbilityScalarValueEntry_Value() {}

var File_AbilityScalarValueEntry_proto protoreflect.FileDescriptor

var file_AbilityScalarValueEntry_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x41, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x02, 0x0a, 0x17, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x26, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x63, 0x61, 0x6c,
	0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x21, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x75, 0x69, 0x6e, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x09,
	0x75, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AbilityScalarValueEntry_proto_rawDescOnce sync.Once
	file_AbilityScalarValueEntry_proto_rawDescData = file_AbilityScalarValueEntry_proto_rawDesc
)

func file_AbilityScalarValueEntry_proto_rawDescGZIP() []byte {
	file_AbilityScalarValueEntry_proto_rawDescOnce.Do(func() {
		file_AbilityScalarValueEntry_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilityScalarValueEntry_proto_rawDescData)
	})
	return file_AbilityScalarValueEntry_proto_rawDescData
}

var file_AbilityScalarValueEntry_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AbilityScalarValueEntry_proto_goTypes = []interface{}{
	(*AbilityScalarValueEntry)(nil), // 0: proto.AbilityScalarValueEntry
	(*AbilityString)(nil),           // 1: proto.AbilityString
	(AbilityScalarType)(0),          // 2: proto.AbilityScalarType
}
var file_AbilityScalarValueEntry_proto_depIdxs = []int32{
	1, // 0: proto.AbilityScalarValueEntry.key:type_name -> proto.AbilityString
	2, // 1: proto.AbilityScalarValueEntry.value_type:type_name -> proto.AbilityScalarType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_AbilityScalarValueEntry_proto_init() }
func file_AbilityScalarValueEntry_proto_init() {
	if File_AbilityScalarValueEntry_proto != nil {
		return
	}
	file_AbilityString_proto_init()
	file_AbilityScalarType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AbilityScalarValueEntry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityScalarValueEntry); i {
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
	file_AbilityScalarValueEntry_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AbilityScalarValueEntry_FloatValue)(nil),
		(*AbilityScalarValueEntry_StringValue)(nil),
		(*AbilityScalarValueEntry_IntValue)(nil),
		(*AbilityScalarValueEntry_UintValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_AbilityScalarValueEntry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilityScalarValueEntry_proto_goTypes,
		DependencyIndexes: file_AbilityScalarValueEntry_proto_depIdxs,
		MessageInfos:      file_AbilityScalarValueEntry_proto_msgTypes,
	}.Build()
	File_AbilityScalarValueEntry_proto = out.File
	file_AbilityScalarValueEntry_proto_rawDesc = nil
	file_AbilityScalarValueEntry_proto_goTypes = nil
	file_AbilityScalarValueEntry_proto_depIdxs = nil
}
