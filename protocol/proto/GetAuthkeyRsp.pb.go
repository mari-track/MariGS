// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GetAuthkeyRsp.proto

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

type GetAuthkeyRsp_CmdId int32

const (
	GetAuthkeyRsp_NONE             GetAuthkeyRsp_CmdId = 0
	GetAuthkeyRsp_CMD_ID           GetAuthkeyRsp_CmdId = 1409
	GetAuthkeyRsp_ENET_CHANNEL_ID  GetAuthkeyRsp_CmdId = 0
	GetAuthkeyRsp_ENET_IS_RELIABLE GetAuthkeyRsp_CmdId = 1
)

// Enum value maps for GetAuthkeyRsp_CmdId.
var (
	GetAuthkeyRsp_CmdId_name = map[int32]string{
		0:    "NONE",
		1409: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	GetAuthkeyRsp_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1409,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x GetAuthkeyRsp_CmdId) Enum() *GetAuthkeyRsp_CmdId {
	p := new(GetAuthkeyRsp_CmdId)
	*p = x
	return p
}

func (x GetAuthkeyRsp_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetAuthkeyRsp_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_GetAuthkeyRsp_proto_enumTypes[0].Descriptor()
}

func (GetAuthkeyRsp_CmdId) Type() protoreflect.EnumType {
	return &file_GetAuthkeyRsp_proto_enumTypes[0]
}

func (x GetAuthkeyRsp_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetAuthkeyRsp_CmdId.Descriptor instead.
func (GetAuthkeyRsp_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_GetAuthkeyRsp_proto_rawDescGZIP(), []int{0, 0}
}

type GetAuthkeyRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode    int32  `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Authkey    string `protobuf:"bytes,2,opt,name=authkey,proto3" json:"authkey,omitempty"`
	AuthAppid  string `protobuf:"bytes,3,opt,name=auth_appid,json=authAppid,proto3" json:"auth_appid,omitempty"`
	SignType   uint32 `protobuf:"varint,4,opt,name=sign_type,json=signType,proto3" json:"sign_type,omitempty"`
	AuthkeyVer uint32 `protobuf:"varint,5,opt,name=authkey_ver,json=authkeyVer,proto3" json:"authkey_ver,omitempty"`
	GameBiz    string `protobuf:"bytes,6,opt,name=game_biz,json=gameBiz,proto3" json:"game_biz,omitempty"`
}

func (x *GetAuthkeyRsp) Reset() {
	*x = GetAuthkeyRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetAuthkeyRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthkeyRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthkeyRsp) ProtoMessage() {}

func (x *GetAuthkeyRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetAuthkeyRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthkeyRsp.ProtoReflect.Descriptor instead.
func (*GetAuthkeyRsp) Descriptor() ([]byte, []int) {
	return file_GetAuthkeyRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetAuthkeyRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetAuthkeyRsp) GetAuthkey() string {
	if x != nil {
		return x.Authkey
	}
	return ""
}

func (x *GetAuthkeyRsp) GetAuthAppid() string {
	if x != nil {
		return x.AuthAppid
	}
	return ""
}

func (x *GetAuthkeyRsp) GetSignType() uint32 {
	if x != nil {
		return x.SignType
	}
	return 0
}

func (x *GetAuthkeyRsp) GetAuthkeyVer() uint32 {
	if x != nil {
		return x.AuthkeyVer
	}
	return 0
}

func (x *GetAuthkeyRsp) GetGameBiz() string {
	if x != nil {
		return x.GameBiz
	}
	return ""
}

var File_GetAuthkeyRsp_proto protoreflect.FileDescriptor

var file_GetAuthkeyRsp_proto_rawDesc = []byte{
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x52, 0x73, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x6b, 0x65, 0x79, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x75, 0x74, 0x68, 0x41, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73,
	0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6b,
	0x65, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x75,
	0x74, 0x68, 0x6b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x62, 0x69, 0x7a, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x61, 0x6d, 0x65,
	0x42, 0x69, 0x7a, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44,
	0x10, 0x81, 0x0b, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54,
	0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02,
	0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetAuthkeyRsp_proto_rawDescOnce sync.Once
	file_GetAuthkeyRsp_proto_rawDescData = file_GetAuthkeyRsp_proto_rawDesc
)

func file_GetAuthkeyRsp_proto_rawDescGZIP() []byte {
	file_GetAuthkeyRsp_proto_rawDescOnce.Do(func() {
		file_GetAuthkeyRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetAuthkeyRsp_proto_rawDescData)
	})
	return file_GetAuthkeyRsp_proto_rawDescData
}

var file_GetAuthkeyRsp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GetAuthkeyRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetAuthkeyRsp_proto_goTypes = []interface{}{
	(GetAuthkeyRsp_CmdId)(0), // 0: GetAuthkeyRsp.CmdId
	(*GetAuthkeyRsp)(nil),    // 1: GetAuthkeyRsp
}
var file_GetAuthkeyRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetAuthkeyRsp_proto_init() }
func file_GetAuthkeyRsp_proto_init() {
	if File_GetAuthkeyRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetAuthkeyRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthkeyRsp); i {
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
			RawDescriptor: file_GetAuthkeyRsp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetAuthkeyRsp_proto_goTypes,
		DependencyIndexes: file_GetAuthkeyRsp_proto_depIdxs,
		EnumInfos:         file_GetAuthkeyRsp_proto_enumTypes,
		MessageInfos:      file_GetAuthkeyRsp_proto_msgTypes,
	}.Build()
	File_GetAuthkeyRsp_proto = out.File
	file_GetAuthkeyRsp_proto_rawDesc = nil
	file_GetAuthkeyRsp_proto_goTypes = nil
	file_GetAuthkeyRsp_proto_depIdxs = nil
}
