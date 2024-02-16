// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: MassiveEntityState.proto

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

type MassiveEntityState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityType   uint32 `protobuf:"varint,1,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	ObjId        int64  `protobuf:"varint,2,opt,name=obj_id,json=objId,proto3" json:"obj_id,omitempty"`
	ElementState uint32 `protobuf:"varint,3,opt,name=element_state,json=elementState,proto3" json:"element_state,omitempty"`
}

func (x *MassiveEntityState) Reset() {
	*x = MassiveEntityState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MassiveEntityState_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MassiveEntityState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MassiveEntityState) ProtoMessage() {}

func (x *MassiveEntityState) ProtoReflect() protoreflect.Message {
	mi := &file_MassiveEntityState_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MassiveEntityState.ProtoReflect.Descriptor instead.
func (*MassiveEntityState) Descriptor() ([]byte, []int) {
	return file_MassiveEntityState_proto_rawDescGZIP(), []int{0}
}

func (x *MassiveEntityState) GetEntityType() uint32 {
	if x != nil {
		return x.EntityType
	}
	return 0
}

func (x *MassiveEntityState) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

func (x *MassiveEntityState) GetElementState() uint32 {
	if x != nil {
		return x.ElementState
	}
	return 0
}

var File_MassiveEntityState_proto protoreflect.FileDescriptor

var file_MassiveEntityState_proto_rawDesc = []byte{
	0x0a, 0x18, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x71, 0x0a, 0x12, 0x4d, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MassiveEntityState_proto_rawDescOnce sync.Once
	file_MassiveEntityState_proto_rawDescData = file_MassiveEntityState_proto_rawDesc
)

func file_MassiveEntityState_proto_rawDescGZIP() []byte {
	file_MassiveEntityState_proto_rawDescOnce.Do(func() {
		file_MassiveEntityState_proto_rawDescData = protoimpl.X.CompressGZIP(file_MassiveEntityState_proto_rawDescData)
	})
	return file_MassiveEntityState_proto_rawDescData
}

var file_MassiveEntityState_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MassiveEntityState_proto_goTypes = []interface{}{
	(*MassiveEntityState)(nil), // 0: proto.MassiveEntityState
}
var file_MassiveEntityState_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MassiveEntityState_proto_init() }
func file_MassiveEntityState_proto_init() {
	if File_MassiveEntityState_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MassiveEntityState_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MassiveEntityState); i {
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
			RawDescriptor: file_MassiveEntityState_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MassiveEntityState_proto_goTypes,
		DependencyIndexes: file_MassiveEntityState_proto_depIdxs,
		MessageInfos:      file_MassiveEntityState_proto_msgTypes,
	}.Build()
	File_MassiveEntityState_proto = out.File
	file_MassiveEntityState_proto_rawDesc = nil
	file_MassiveEntityState_proto_goTypes = nil
	file_MassiveEntityState_proto_depIdxs = nil
}
