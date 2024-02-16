// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: ActivityScheduleInfoNotify.proto

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

type ActivityScheduleInfoNotify_CmdId int32

const (
	ActivityScheduleInfoNotify_NONE             ActivityScheduleInfoNotify_CmdId = 0
	ActivityScheduleInfoNotify_CMD_ID           ActivityScheduleInfoNotify_CmdId = 2007
	ActivityScheduleInfoNotify_ENET_CHANNEL_ID  ActivityScheduleInfoNotify_CmdId = 0
	ActivityScheduleInfoNotify_ENET_IS_RELIABLE ActivityScheduleInfoNotify_CmdId = 1
	ActivityScheduleInfoNotify_IS_ALLOW_CLIENT  ActivityScheduleInfoNotify_CmdId = 1
)

// Enum value maps for ActivityScheduleInfoNotify_CmdId.
var (
	ActivityScheduleInfoNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		2007: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	ActivityScheduleInfoNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           2007,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x ActivityScheduleInfoNotify_CmdId) Enum() *ActivityScheduleInfoNotify_CmdId {
	p := new(ActivityScheduleInfoNotify_CmdId)
	*p = x
	return p
}

func (x ActivityScheduleInfoNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActivityScheduleInfoNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_ActivityScheduleInfoNotify_proto_enumTypes[0].Descriptor()
}

func (ActivityScheduleInfoNotify_CmdId) Type() protoreflect.EnumType {
	return &file_ActivityScheduleInfoNotify_proto_enumTypes[0]
}

func (x ActivityScheduleInfoNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActivityScheduleInfoNotify_CmdId.Descriptor instead.
func (ActivityScheduleInfoNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_ActivityScheduleInfoNotify_proto_rawDescGZIP(), []int{0, 0}
}

type ActivityScheduleInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityScheduleList []*ActivityScheduleInfo `protobuf:"bytes,1,rep,name=activity_schedule_list,json=activityScheduleList,proto3" json:"activity_schedule_list,omitempty"`
	RemainFlySeaLampNum  uint32                  `protobuf:"varint,2,opt,name=remain_fly_sea_lamp_num,json=remainFlySeaLampNum,proto3" json:"remain_fly_sea_lamp_num,omitempty"`
}

func (x *ActivityScheduleInfoNotify) Reset() {
	*x = ActivityScheduleInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActivityScheduleInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityScheduleInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityScheduleInfoNotify) ProtoMessage() {}

func (x *ActivityScheduleInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ActivityScheduleInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityScheduleInfoNotify.ProtoReflect.Descriptor instead.
func (*ActivityScheduleInfoNotify) Descriptor() ([]byte, []int) {
	return file_ActivityScheduleInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityScheduleInfoNotify) GetActivityScheduleList() []*ActivityScheduleInfo {
	if x != nil {
		return x.ActivityScheduleList
	}
	return nil
}

func (x *ActivityScheduleInfoNotify) GetRemainFlySeaLampNum() uint32 {
	if x != nil {
		return x.RemainFlySeaLampNum
	}
	return 0
}

var File_ActivityScheduleInfoNotify_proto protoreflect.FileDescriptor

var file_ActivityScheduleInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x02, 0x0a, 0x1a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x51, 0x0a, 0x16, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x14, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x17, 0x72, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x66, 0x6c, 0x79, 0x5f, 0x73, 0x65, 0x61, 0x5f, 0x6c, 0x61, 0x6d, 0x70, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e,
	0x46, 0x6c, 0x79, 0x53, 0x65, 0x61, 0x4c, 0x61, 0x6d, 0x70, 0x4e, 0x75, 0x6d, 0x22, 0x62, 0x0a,
	0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xd7, 0x0f, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45,
	0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41,
	0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ActivityScheduleInfoNotify_proto_rawDescOnce sync.Once
	file_ActivityScheduleInfoNotify_proto_rawDescData = file_ActivityScheduleInfoNotify_proto_rawDesc
)

func file_ActivityScheduleInfoNotify_proto_rawDescGZIP() []byte {
	file_ActivityScheduleInfoNotify_proto_rawDescOnce.Do(func() {
		file_ActivityScheduleInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ActivityScheduleInfoNotify_proto_rawDescData)
	})
	return file_ActivityScheduleInfoNotify_proto_rawDescData
}

var file_ActivityScheduleInfoNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ActivityScheduleInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ActivityScheduleInfoNotify_proto_goTypes = []interface{}{
	(ActivityScheduleInfoNotify_CmdId)(0), // 0: proto.ActivityScheduleInfoNotify.CmdId
	(*ActivityScheduleInfoNotify)(nil),    // 1: proto.ActivityScheduleInfoNotify
	(*ActivityScheduleInfo)(nil),          // 2: proto.ActivityScheduleInfo
}
var file_ActivityScheduleInfoNotify_proto_depIdxs = []int32{
	2, // 0: proto.ActivityScheduleInfoNotify.activity_schedule_list:type_name -> proto.ActivityScheduleInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ActivityScheduleInfoNotify_proto_init() }
func file_ActivityScheduleInfoNotify_proto_init() {
	if File_ActivityScheduleInfoNotify_proto != nil {
		return
	}
	file_ActivityScheduleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ActivityScheduleInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityScheduleInfoNotify); i {
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
			RawDescriptor: file_ActivityScheduleInfoNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ActivityScheduleInfoNotify_proto_goTypes,
		DependencyIndexes: file_ActivityScheduleInfoNotify_proto_depIdxs,
		EnumInfos:         file_ActivityScheduleInfoNotify_proto_enumTypes,
		MessageInfos:      file_ActivityScheduleInfoNotify_proto_msgTypes,
	}.Build()
	File_ActivityScheduleInfoNotify_proto = out.File
	file_ActivityScheduleInfoNotify_proto_rawDesc = nil
	file_ActivityScheduleInfoNotify_proto_goTypes = nil
	file_ActivityScheduleInfoNotify_proto_depIdxs = nil
}