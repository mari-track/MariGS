// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: SalesmanActivityDetailInfo.proto

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

type SalesmanActivityDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status              SalesmanStatusType `protobuf:"varint,1,opt,name=status,proto3,enum=proto.SalesmanStatusType" json:"status,omitempty"`
	DayIndex            uint32             `protobuf:"varint,2,opt,name=day_index,json=dayIndex,proto3" json:"day_index,omitempty"`
	LastDeliverTime     uint32             `protobuf:"varint,3,opt,name=last_deliver_time,json=lastDeliverTime,proto3" json:"last_deliver_time,omitempty"`
	DeliverCount        uint32             `protobuf:"varint,4,opt,name=deliver_count,json=deliverCount,proto3" json:"deliver_count,omitempty"`
	SelectedRewardIdMap map[uint32]uint32  `protobuf:"bytes,5,rep,name=selected_reward_id_map,json=selectedRewardIdMap,proto3" json:"selected_reward_id_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *SalesmanActivityDetailInfo) Reset() {
	*x = SalesmanActivityDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SalesmanActivityDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SalesmanActivityDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SalesmanActivityDetailInfo) ProtoMessage() {}

func (x *SalesmanActivityDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SalesmanActivityDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SalesmanActivityDetailInfo.ProtoReflect.Descriptor instead.
func (*SalesmanActivityDetailInfo) Descriptor() ([]byte, []int) {
	return file_SalesmanActivityDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SalesmanActivityDetailInfo) GetStatus() SalesmanStatusType {
	if x != nil {
		return x.Status
	}
	return SalesmanStatusType_SALESMAN_STATUS_NONE
}

func (x *SalesmanActivityDetailInfo) GetDayIndex() uint32 {
	if x != nil {
		return x.DayIndex
	}
	return 0
}

func (x *SalesmanActivityDetailInfo) GetLastDeliverTime() uint32 {
	if x != nil {
		return x.LastDeliverTime
	}
	return 0
}

func (x *SalesmanActivityDetailInfo) GetDeliverCount() uint32 {
	if x != nil {
		return x.DeliverCount
	}
	return 0
}

func (x *SalesmanActivityDetailInfo) GetSelectedRewardIdMap() map[uint32]uint32 {
	if x != nil {
		return x.SelectedRewardIdMap
	}
	return nil
}

var File_SalesmanActivityDetailInfo_proto protoreflect.FileDescriptor

var file_SalesmanActivityDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x20, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x6d, 0x61, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x53, 0x61, 0x6c, 0x65, 0x73,
	0x6d, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x02, 0x0a, 0x1a, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x6d, 0x61, 0x6e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x73,
	0x6d, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x79, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x61, 0x79, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6c,
	0x61, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x6f, 0x0a, 0x16, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x6d, 0x61, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x64, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x13, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49,
	0x64, 0x4d, 0x61, 0x70, 0x1a, 0x46, 0x0a, 0x18, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x64, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SalesmanActivityDetailInfo_proto_rawDescOnce sync.Once
	file_SalesmanActivityDetailInfo_proto_rawDescData = file_SalesmanActivityDetailInfo_proto_rawDesc
)

func file_SalesmanActivityDetailInfo_proto_rawDescGZIP() []byte {
	file_SalesmanActivityDetailInfo_proto_rawDescOnce.Do(func() {
		file_SalesmanActivityDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SalesmanActivityDetailInfo_proto_rawDescData)
	})
	return file_SalesmanActivityDetailInfo_proto_rawDescData
}

var file_SalesmanActivityDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_SalesmanActivityDetailInfo_proto_goTypes = []interface{}{
	(*SalesmanActivityDetailInfo)(nil), // 0: proto.SalesmanActivityDetailInfo
	nil,                                // 1: proto.SalesmanActivityDetailInfo.SelectedRewardIdMapEntry
	(SalesmanStatusType)(0),            // 2: proto.SalesmanStatusType
}
var file_SalesmanActivityDetailInfo_proto_depIdxs = []int32{
	2, // 0: proto.SalesmanActivityDetailInfo.status:type_name -> proto.SalesmanStatusType
	1, // 1: proto.SalesmanActivityDetailInfo.selected_reward_id_map:type_name -> proto.SalesmanActivityDetailInfo.SelectedRewardIdMapEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SalesmanActivityDetailInfo_proto_init() }
func file_SalesmanActivityDetailInfo_proto_init() {
	if File_SalesmanActivityDetailInfo_proto != nil {
		return
	}
	file_SalesmanStatusType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SalesmanActivityDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SalesmanActivityDetailInfo); i {
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
			RawDescriptor: file_SalesmanActivityDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SalesmanActivityDetailInfo_proto_goTypes,
		DependencyIndexes: file_SalesmanActivityDetailInfo_proto_depIdxs,
		MessageInfos:      file_SalesmanActivityDetailInfo_proto_msgTypes,
	}.Build()
	File_SalesmanActivityDetailInfo_proto = out.File
	file_SalesmanActivityDetailInfo_proto_rawDesc = nil
	file_SalesmanActivityDetailInfo_proto_goTypes = nil
	file_SalesmanActivityDetailInfo_proto_depIdxs = nil
}
