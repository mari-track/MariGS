// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: GachaInfo.proto

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

type GachaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GachaType              uint32 `protobuf:"varint,1,opt,name=gacha_type,json=gachaType,proto3" json:"gacha_type,omitempty"`
	ScheduleId             uint32 `protobuf:"varint,2,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	BeginTime              uint32 `protobuf:"varint,3,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	EndTime                uint32 `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	CostItemId             uint32 `protobuf:"varint,5,opt,name=cost_item_id,json=costItemId,proto3" json:"cost_item_id,omitempty"`
	CostItemNum            uint32 `protobuf:"varint,6,opt,name=cost_item_num,json=costItemNum,proto3" json:"cost_item_num,omitempty"`
	GachaPrefabPath        string `protobuf:"bytes,7,opt,name=gacha_prefab_path,json=gachaPrefabPath,proto3" json:"gacha_prefab_path,omitempty"`
	GachaProbUrl           string `protobuf:"bytes,8,opt,name=gacha_prob_url,json=gachaProbUrl,proto3" json:"gacha_prob_url,omitempty"`
	GachaRecordUrl         string `protobuf:"bytes,9,opt,name=gacha_record_url,json=gachaRecordUrl,proto3" json:"gacha_record_url,omitempty"`
	GachaPreviewPrefabPath string `protobuf:"bytes,10,opt,name=gacha_preview_prefab_path,json=gachaPreviewPrefabPath,proto3" json:"gacha_preview_prefab_path,omitempty"`
	TenCostItemId          uint32 `protobuf:"varint,11,opt,name=ten_cost_item_id,json=tenCostItemId,proto3" json:"ten_cost_item_id,omitempty"`
	TenCostItemNum         uint32 `protobuf:"varint,12,opt,name=ten_cost_item_num,json=tenCostItemNum,proto3" json:"ten_cost_item_num,omitempty"`
	LeftGachaTimes         uint32 `protobuf:"varint,13,opt,name=left_gacha_times,json=leftGachaTimes,proto3" json:"left_gacha_times,omitempty"`
	GachaTimesLimit        uint32 `protobuf:"varint,14,opt,name=gacha_times_limit,json=gachaTimesLimit,proto3" json:"gacha_times_limit,omitempty"`
	GachaSortId            uint32 `protobuf:"varint,15,opt,name=gacha_sort_id,json=gachaSortId,proto3" json:"gacha_sort_id,omitempty"`
}

func (x *GachaInfo) Reset() {
	*x = GachaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GachaInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GachaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GachaInfo) ProtoMessage() {}

func (x *GachaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GachaInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GachaInfo.ProtoReflect.Descriptor instead.
func (*GachaInfo) Descriptor() ([]byte, []int) {
	return file_GachaInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GachaInfo) GetGachaType() uint32 {
	if x != nil {
		return x.GachaType
	}
	return 0
}

func (x *GachaInfo) GetScheduleId() uint32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

func (x *GachaInfo) GetBeginTime() uint32 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *GachaInfo) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *GachaInfo) GetCostItemId() uint32 {
	if x != nil {
		return x.CostItemId
	}
	return 0
}

func (x *GachaInfo) GetCostItemNum() uint32 {
	if x != nil {
		return x.CostItemNum
	}
	return 0
}

func (x *GachaInfo) GetGachaPrefabPath() string {
	if x != nil {
		return x.GachaPrefabPath
	}
	return ""
}

func (x *GachaInfo) GetGachaProbUrl() string {
	if x != nil {
		return x.GachaProbUrl
	}
	return ""
}

func (x *GachaInfo) GetGachaRecordUrl() string {
	if x != nil {
		return x.GachaRecordUrl
	}
	return ""
}

func (x *GachaInfo) GetGachaPreviewPrefabPath() string {
	if x != nil {
		return x.GachaPreviewPrefabPath
	}
	return ""
}

func (x *GachaInfo) GetTenCostItemId() uint32 {
	if x != nil {
		return x.TenCostItemId
	}
	return 0
}

func (x *GachaInfo) GetTenCostItemNum() uint32 {
	if x != nil {
		return x.TenCostItemNum
	}
	return 0
}

func (x *GachaInfo) GetLeftGachaTimes() uint32 {
	if x != nil {
		return x.LeftGachaTimes
	}
	return 0
}

func (x *GachaInfo) GetGachaTimesLimit() uint32 {
	if x != nil {
		return x.GachaTimesLimit
	}
	return 0
}

func (x *GachaInfo) GetGachaSortId() uint32 {
	if x != nil {
		return x.GachaSortId
	}
	return 0
}

var File_GachaInfo_proto protoreflect.FileDescriptor

var file_GachaInfo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x04, 0x0a, 0x09, 0x47, 0x61, 0x63,
	0x68, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0c, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x4e, 0x75, 0x6d, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f,
	0x70, 0x72, 0x65, 0x66, 0x61, 0x62, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x50, 0x72, 0x65, 0x66, 0x61, 0x62, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x62,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x50, 0x72, 0x6f, 0x62, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x67, 0x61, 0x63, 0x68, 0x61, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55,
	0x72, 0x6c, 0x12, 0x39, 0x0a, 0x19, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x61, 0x62, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x67, 0x61, 0x63, 0x68, 0x61, 0x50, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x50, 0x72, 0x65, 0x66, 0x61, 0x62, 0x50, 0x61, 0x74, 0x68, 0x12, 0x27, 0x0a,
	0x10, 0x74, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x65, 0x6e, 0x43, 0x6f, 0x73, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x11, 0x74, 0x65, 0x6e, 0x5f, 0x63, 0x6f,
	0x73, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x74, 0x65, 0x6e, 0x43, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x75,
	0x6d, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6c, 0x65, 0x66,
	0x74, 0x47, 0x61, 0x63, 0x68, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x67,
	0x61, 0x63, 0x68, 0x61, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x67, 0x61, 0x63, 0x68, 0x61,
	0x5f, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x67, 0x61, 0x63, 0x68, 0x61, 0x53, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GachaInfo_proto_rawDescOnce sync.Once
	file_GachaInfo_proto_rawDescData = file_GachaInfo_proto_rawDesc
)

func file_GachaInfo_proto_rawDescGZIP() []byte {
	file_GachaInfo_proto_rawDescOnce.Do(func() {
		file_GachaInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_GachaInfo_proto_rawDescData)
	})
	return file_GachaInfo_proto_rawDescData
}

var file_GachaInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GachaInfo_proto_goTypes = []interface{}{
	(*GachaInfo)(nil), // 0: proto.GachaInfo
}
var file_GachaInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GachaInfo_proto_init() }
func file_GachaInfo_proto_init() {
	if File_GachaInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GachaInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GachaInfo); i {
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
			RawDescriptor: file_GachaInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GachaInfo_proto_goTypes,
		DependencyIndexes: file_GachaInfo_proto_depIdxs,
		MessageInfos:      file_GachaInfo_proto_msgTypes,
	}.Build()
	File_GachaInfo_proto = out.File
	file_GachaInfo_proto_rawDesc = nil
	file_GachaInfo_proto_goTypes = nil
	file_GachaInfo_proto_depIdxs = nil
}
