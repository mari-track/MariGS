// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: LevelupCityReq.proto

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

type LevelupCityReq_CmdId int32

const (
	LevelupCityReq_NONE             LevelupCityReq_CmdId = 0
	LevelupCityReq_CMD_ID           LevelupCityReq_CmdId = 255
	LevelupCityReq_ENET_CHANNEL_ID  LevelupCityReq_CmdId = 0
	LevelupCityReq_ENET_IS_RELIABLE LevelupCityReq_CmdId = 1
	LevelupCityReq_IS_ALLOW_CLIENT  LevelupCityReq_CmdId = 1
)

// Enum value maps for LevelupCityReq_CmdId.
var (
	LevelupCityReq_CmdId_name = map[int32]string{
		0:   "NONE",
		255: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	LevelupCityReq_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           255,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x LevelupCityReq_CmdId) Enum() *LevelupCityReq_CmdId {
	p := new(LevelupCityReq_CmdId)
	*p = x
	return p
}

func (x LevelupCityReq_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LevelupCityReq_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_LevelupCityReq_proto_enumTypes[0].Descriptor()
}

func (LevelupCityReq_CmdId) Type() protoreflect.EnumType {
	return &file_LevelupCityReq_proto_enumTypes[0]
}

func (x LevelupCityReq_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LevelupCityReq_CmdId.Descriptor instead.
func (LevelupCityReq_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_LevelupCityReq_proto_rawDescGZIP(), []int{0, 0}
}

type LevelupCityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId uint32 `protobuf:"varint,1,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	AreaId  uint32 `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	ItemNum uint32 `protobuf:"varint,3,opt,name=item_num,json=itemNum,proto3" json:"item_num,omitempty"`
}

func (x *LevelupCityReq) Reset() {
	*x = LevelupCityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LevelupCityReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelupCityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelupCityReq) ProtoMessage() {}

func (x *LevelupCityReq) ProtoReflect() protoreflect.Message {
	mi := &file_LevelupCityReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelupCityReq.ProtoReflect.Descriptor instead.
func (*LevelupCityReq) Descriptor() ([]byte, []int) {
	return file_LevelupCityReq_proto_rawDescGZIP(), []int{0}
}

func (x *LevelupCityReq) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *LevelupCityReq) GetAreaId() uint32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *LevelupCityReq) GetItemNum() uint32 {
	if x != nil {
		return x.ItemNum
	}
	return 0
}

var File_LevelupCityReq_proto protoreflect.FileDescriptor

var file_LevelupCityReq_proto_rawDesc = []byte{
	0x0a, 0x14, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x75, 0x70, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x0e, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x75, 0x70, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x75, 0x6d, 0x22, 0x62, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49,
	0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43,
	0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xff, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f,
	0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LevelupCityReq_proto_rawDescOnce sync.Once
	file_LevelupCityReq_proto_rawDescData = file_LevelupCityReq_proto_rawDesc
)

func file_LevelupCityReq_proto_rawDescGZIP() []byte {
	file_LevelupCityReq_proto_rawDescOnce.Do(func() {
		file_LevelupCityReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_LevelupCityReq_proto_rawDescData)
	})
	return file_LevelupCityReq_proto_rawDescData
}

var file_LevelupCityReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_LevelupCityReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LevelupCityReq_proto_goTypes = []interface{}{
	(LevelupCityReq_CmdId)(0), // 0: LevelupCityReq.CmdId
	(*LevelupCityReq)(nil),    // 1: LevelupCityReq
}
var file_LevelupCityReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LevelupCityReq_proto_init() }
func file_LevelupCityReq_proto_init() {
	if File_LevelupCityReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_LevelupCityReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelupCityReq); i {
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
			RawDescriptor: file_LevelupCityReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LevelupCityReq_proto_goTypes,
		DependencyIndexes: file_LevelupCityReq_proto_depIdxs,
		EnumInfos:         file_LevelupCityReq_proto_enumTypes,
		MessageInfos:      file_LevelupCityReq_proto_msgTypes,
	}.Build()
	File_LevelupCityReq_proto = out.File
	file_LevelupCityReq_proto_rawDesc = nil
	file_LevelupCityReq_proto_goTypes = nil
	file_LevelupCityReq_proto_depIdxs = nil
}
