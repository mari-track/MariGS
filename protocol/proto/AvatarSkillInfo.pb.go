// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: AvatarSkillInfo.proto

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

type AvatarSkillInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PassCdTime     uint32   `protobuf:"varint,1,opt,name=pass_cd_time,json=passCdTime,proto3" json:"pass_cd_time,omitempty"`
	FullCdTimeList []uint32 `protobuf:"varint,2,rep,packed,name=full_cd_time_list,json=fullCdTimeList,proto3" json:"full_cd_time_list,omitempty"`
	MaxChargeCount uint32   `protobuf:"varint,3,opt,name=max_charge_count,json=maxChargeCount,proto3" json:"max_charge_count,omitempty"`
}

func (x *AvatarSkillInfo) Reset() {
	*x = AvatarSkillInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarSkillInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarSkillInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarSkillInfo) ProtoMessage() {}

func (x *AvatarSkillInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarSkillInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarSkillInfo.ProtoReflect.Descriptor instead.
func (*AvatarSkillInfo) Descriptor() ([]byte, []int) {
	return file_AvatarSkillInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarSkillInfo) GetPassCdTime() uint32 {
	if x != nil {
		return x.PassCdTime
	}
	return 0
}

func (x *AvatarSkillInfo) GetFullCdTimeList() []uint32 {
	if x != nil {
		return x.FullCdTimeList
	}
	return nil
}

func (x *AvatarSkillInfo) GetMaxChargeCount() uint32 {
	if x != nil {
		return x.MaxChargeCount
	}
	return 0
}

var File_AvatarSkillInfo_proto protoreflect.FileDescriptor

var file_AvatarSkillInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88,
	0x01, 0x0a, 0x0f, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x63, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x73, 0x73, 0x43, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x11, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x63, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x0e, 0x66, 0x75, 0x6c, 0x6c, 0x43, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x43, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarSkillInfo_proto_rawDescOnce sync.Once
	file_AvatarSkillInfo_proto_rawDescData = file_AvatarSkillInfo_proto_rawDesc
)

func file_AvatarSkillInfo_proto_rawDescGZIP() []byte {
	file_AvatarSkillInfo_proto_rawDescOnce.Do(func() {
		file_AvatarSkillInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarSkillInfo_proto_rawDescData)
	})
	return file_AvatarSkillInfo_proto_rawDescData
}

var file_AvatarSkillInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarSkillInfo_proto_goTypes = []interface{}{
	(*AvatarSkillInfo)(nil), // 0: proto.AvatarSkillInfo
}
var file_AvatarSkillInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarSkillInfo_proto_init() }
func file_AvatarSkillInfo_proto_init() {
	if File_AvatarSkillInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarSkillInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarSkillInfo); i {
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
			RawDescriptor: file_AvatarSkillInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarSkillInfo_proto_goTypes,
		DependencyIndexes: file_AvatarSkillInfo_proto_depIdxs,
		MessageInfos:      file_AvatarSkillInfo_proto_msgTypes,
	}.Build()
	File_AvatarSkillInfo_proto = out.File
	file_AvatarSkillInfo_proto_rawDesc = nil
	file_AvatarSkillInfo_proto_goTypes = nil
	file_AvatarSkillInfo_proto_depIdxs = nil
}
