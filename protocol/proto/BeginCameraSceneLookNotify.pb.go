// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: BeginCameraSceneLookNotify.proto

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

type BeginCameraSceneLookNotify_CmdId int32

const (
	BeginCameraSceneLookNotify_NONE             BeginCameraSceneLookNotify_CmdId = 0
	BeginCameraSceneLookNotify_CMD_ID           BeginCameraSceneLookNotify_CmdId = 247
	BeginCameraSceneLookNotify_ENET_CHANNEL_ID  BeginCameraSceneLookNotify_CmdId = 0
	BeginCameraSceneLookNotify_ENET_IS_RELIABLE BeginCameraSceneLookNotify_CmdId = 1
)

// Enum value maps for BeginCameraSceneLookNotify_CmdId.
var (
	BeginCameraSceneLookNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		247: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	BeginCameraSceneLookNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           247,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x BeginCameraSceneLookNotify_CmdId) Enum() *BeginCameraSceneLookNotify_CmdId {
	p := new(BeginCameraSceneLookNotify_CmdId)
	*p = x
	return p
}

func (x BeginCameraSceneLookNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BeginCameraSceneLookNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_BeginCameraSceneLookNotify_proto_enumTypes[0].Descriptor()
}

func (BeginCameraSceneLookNotify_CmdId) Type() protoreflect.EnumType {
	return &file_BeginCameraSceneLookNotify_proto_enumTypes[0]
}

func (x BeginCameraSceneLookNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BeginCameraSceneLookNotify_CmdId.Descriptor instead.
func (BeginCameraSceneLookNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_BeginCameraSceneLookNotify_proto_rawDescGZIP(), []int{0, 0}
}

type BeginCameraSceneLookNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LookPos              *Vector `protobuf:"bytes,1,opt,name=look_pos,json=lookPos,proto3" json:"look_pos,omitempty"`
	Duration             float32 `protobuf:"fixed32,2,opt,name=duration,proto3" json:"duration,omitempty"`
	IsForce              bool    `protobuf:"varint,3,opt,name=is_force,json=isForce,proto3" json:"is_force,omitempty"`
	IsRecoverKeepCurrent bool    `protobuf:"varint,4,opt,name=is_recover_keep_current,json=isRecoverKeepCurrent,proto3" json:"is_recover_keep_current,omitempty"`
	IsAllowInput         bool    `protobuf:"varint,5,opt,name=is_allow_input,json=isAllowInput,proto3" json:"is_allow_input,omitempty"`
}

func (x *BeginCameraSceneLookNotify) Reset() {
	*x = BeginCameraSceneLookNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BeginCameraSceneLookNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeginCameraSceneLookNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeginCameraSceneLookNotify) ProtoMessage() {}

func (x *BeginCameraSceneLookNotify) ProtoReflect() protoreflect.Message {
	mi := &file_BeginCameraSceneLookNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeginCameraSceneLookNotify.ProtoReflect.Descriptor instead.
func (*BeginCameraSceneLookNotify) Descriptor() ([]byte, []int) {
	return file_BeginCameraSceneLookNotify_proto_rawDescGZIP(), []int{0}
}

func (x *BeginCameraSceneLookNotify) GetLookPos() *Vector {
	if x != nil {
		return x.LookPos
	}
	return nil
}

func (x *BeginCameraSceneLookNotify) GetDuration() float32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *BeginCameraSceneLookNotify) GetIsForce() bool {
	if x != nil {
		return x.IsForce
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetIsRecoverKeepCurrent() bool {
	if x != nil {
		return x.IsRecoverKeepCurrent
	}
	return false
}

func (x *BeginCameraSceneLookNotify) GetIsAllowInput() bool {
	if x != nil {
		return x.IsAllowInput
	}
	return false
}

var File_BeginCameraSceneLookNotify_proto protoreflect.FileDescriptor

var file_BeginCameraSceneLookNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x02, 0x0a, 0x1a, 0x42, 0x65, 0x67, 0x69,
	0x6e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4c, 0x6f, 0x6f, 0x6b,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x28, 0x0a, 0x08, 0x6c, 0x6f, 0x6f, 0x6b, 0x5f, 0x70,
	0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x6c, 0x6f, 0x6f, 0x6b, 0x50, 0x6f, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x69, 0x73, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x4b, 0x65, 0x65, 0x70, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x24,
	0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49,
	0x44, 0x10, 0xf7, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45,
	0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a,
	0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BeginCameraSceneLookNotify_proto_rawDescOnce sync.Once
	file_BeginCameraSceneLookNotify_proto_rawDescData = file_BeginCameraSceneLookNotify_proto_rawDesc
)

func file_BeginCameraSceneLookNotify_proto_rawDescGZIP() []byte {
	file_BeginCameraSceneLookNotify_proto_rawDescOnce.Do(func() {
		file_BeginCameraSceneLookNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_BeginCameraSceneLookNotify_proto_rawDescData)
	})
	return file_BeginCameraSceneLookNotify_proto_rawDescData
}

var file_BeginCameraSceneLookNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BeginCameraSceneLookNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BeginCameraSceneLookNotify_proto_goTypes = []interface{}{
	(BeginCameraSceneLookNotify_CmdId)(0), // 0: proto.BeginCameraSceneLookNotify.CmdId
	(*BeginCameraSceneLookNotify)(nil),    // 1: proto.BeginCameraSceneLookNotify
	(*Vector)(nil),                        // 2: proto.Vector
}
var file_BeginCameraSceneLookNotify_proto_depIdxs = []int32{
	2, // 0: proto.BeginCameraSceneLookNotify.look_pos:type_name -> proto.Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BeginCameraSceneLookNotify_proto_init() }
func file_BeginCameraSceneLookNotify_proto_init() {
	if File_BeginCameraSceneLookNotify_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BeginCameraSceneLookNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeginCameraSceneLookNotify); i {
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
			RawDescriptor: file_BeginCameraSceneLookNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BeginCameraSceneLookNotify_proto_goTypes,
		DependencyIndexes: file_BeginCameraSceneLookNotify_proto_depIdxs,
		EnumInfos:         file_BeginCameraSceneLookNotify_proto_enumTypes,
		MessageInfos:      file_BeginCameraSceneLookNotify_proto_msgTypes,
	}.Build()
	File_BeginCameraSceneLookNotify_proto = out.File
	file_BeginCameraSceneLookNotify_proto_rawDesc = nil
	file_BeginCameraSceneLookNotify_proto_goTypes = nil
	file_BeginCameraSceneLookNotify_proto_depIdxs = nil
}