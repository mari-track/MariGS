// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GadgetPlayInfo.proto

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

type GadgetPlayInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayType          uint32   `protobuf:"varint,1,opt,name=play_type,json=playType,proto3" json:"play_type,omitempty"`
	Duration          uint32   `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	ProgressStageList []uint32 `protobuf:"varint,3,rep,packed,name=progress_stage_list,json=progressStageList,proto3" json:"progress_stage_list,omitempty"`
	StartCd           uint32   `protobuf:"varint,4,opt,name=start_cd,json=startCd,proto3" json:"start_cd,omitempty"`
	StartTime         uint32   `protobuf:"varint,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Progress          uint32   `protobuf:"varint,6,opt,name=progress,proto3" json:"progress,omitempty"`
	// Types that are assignable to PlayInfo:
	//
	//	*GadgetPlayInfo_CrucibleInfo
	PlayInfo isGadgetPlayInfo_PlayInfo `protobuf_oneof:"play_info"`
}

func (x *GadgetPlayInfo) Reset() {
	*x = GadgetPlayInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GadgetPlayInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GadgetPlayInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GadgetPlayInfo) ProtoMessage() {}

func (x *GadgetPlayInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GadgetPlayInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GadgetPlayInfo.ProtoReflect.Descriptor instead.
func (*GadgetPlayInfo) Descriptor() ([]byte, []int) {
	return file_GadgetPlayInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GadgetPlayInfo) GetPlayType() uint32 {
	if x != nil {
		return x.PlayType
	}
	return 0
}

func (x *GadgetPlayInfo) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *GadgetPlayInfo) GetProgressStageList() []uint32 {
	if x != nil {
		return x.ProgressStageList
	}
	return nil
}

func (x *GadgetPlayInfo) GetStartCd() uint32 {
	if x != nil {
		return x.StartCd
	}
	return 0
}

func (x *GadgetPlayInfo) GetStartTime() uint32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *GadgetPlayInfo) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (m *GadgetPlayInfo) GetPlayInfo() isGadgetPlayInfo_PlayInfo {
	if m != nil {
		return m.PlayInfo
	}
	return nil
}

func (x *GadgetPlayInfo) GetCrucibleInfo() *GadgetCrucibleInfo {
	if x, ok := x.GetPlayInfo().(*GadgetPlayInfo_CrucibleInfo); ok {
		return x.CrucibleInfo
	}
	return nil
}

type isGadgetPlayInfo_PlayInfo interface {
	isGadgetPlayInfo_PlayInfo()
}

type GadgetPlayInfo_CrucibleInfo struct {
	CrucibleInfo *GadgetCrucibleInfo `protobuf:"bytes,21,opt,name=crucible_info,json=crucibleInfo,proto3,oneof"`
}

func (*GadgetPlayInfo_CrucibleInfo) isGadgetPlayInfo_PlayInfo() {}

var File_GadgetPlayInfo_proto protoreflect.FileDescriptor

var file_GadgetPlayInfo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x47,
	0x61, 0x64, 0x67, 0x65, 0x74, 0x43, 0x72, 0x75, 0x63, 0x69, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x02, 0x0a, 0x0e, 0x47, 0x61, 0x64, 0x67,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x11, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x67, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x40, 0x0a, 0x0d, 0x63, 0x72, 0x75,
	0x63, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x43,
	0x72, 0x75, 0x63, 0x69, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0c, 0x63,
	0x72, 0x75, 0x63, 0x69, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0b, 0x0a, 0x09, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GadgetPlayInfo_proto_rawDescOnce sync.Once
	file_GadgetPlayInfo_proto_rawDescData = file_GadgetPlayInfo_proto_rawDesc
)

func file_GadgetPlayInfo_proto_rawDescGZIP() []byte {
	file_GadgetPlayInfo_proto_rawDescOnce.Do(func() {
		file_GadgetPlayInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_GadgetPlayInfo_proto_rawDescData)
	})
	return file_GadgetPlayInfo_proto_rawDescData
}

var file_GadgetPlayInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GadgetPlayInfo_proto_goTypes = []interface{}{
	(*GadgetPlayInfo)(nil),     // 0: proto.GadgetPlayInfo
	(*GadgetCrucibleInfo)(nil), // 1: proto.GadgetCrucibleInfo
}
var file_GadgetPlayInfo_proto_depIdxs = []int32{
	1, // 0: proto.GadgetPlayInfo.crucible_info:type_name -> proto.GadgetCrucibleInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GadgetPlayInfo_proto_init() }
func file_GadgetPlayInfo_proto_init() {
	if File_GadgetPlayInfo_proto != nil {
		return
	}
	file_GadgetCrucibleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GadgetPlayInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GadgetPlayInfo); i {
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
	file_GadgetPlayInfo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GadgetPlayInfo_CrucibleInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GadgetPlayInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GadgetPlayInfo_proto_goTypes,
		DependencyIndexes: file_GadgetPlayInfo_proto_depIdxs,
		MessageInfos:      file_GadgetPlayInfo_proto_msgTypes,
	}.Build()
	File_GadgetPlayInfo_proto = out.File
	file_GadgetPlayInfo_proto_rawDesc = nil
	file_GadgetPlayInfo_proto_goTypes = nil
	file_GadgetPlayInfo_proto_depIdxs = nil
}
