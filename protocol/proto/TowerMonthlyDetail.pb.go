// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: TowerMonthlyDetail.proto

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

type TowerMonthlyDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MonthlyBrief        *TowerMonthlyBrief        `protobuf:"bytes,1,opt,name=monthly_brief,json=monthlyBrief,proto3" json:"monthly_brief,omitempty"`
	MonthlyCombatRecord *TowerMonthlyCombatRecord `protobuf:"bytes,2,opt,name=monthly_combat_record,json=monthlyCombatRecord,proto3" json:"monthly_combat_record,omitempty"`
}

func (x *TowerMonthlyDetail) Reset() {
	*x = TowerMonthlyDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TowerMonthlyDetail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TowerMonthlyDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TowerMonthlyDetail) ProtoMessage() {}

func (x *TowerMonthlyDetail) ProtoReflect() protoreflect.Message {
	mi := &file_TowerMonthlyDetail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TowerMonthlyDetail.ProtoReflect.Descriptor instead.
func (*TowerMonthlyDetail) Descriptor() ([]byte, []int) {
	return file_TowerMonthlyDetail_proto_rawDescGZIP(), []int{0}
}

func (x *TowerMonthlyDetail) GetMonthlyBrief() *TowerMonthlyBrief {
	if x != nil {
		return x.MonthlyBrief
	}
	return nil
}

func (x *TowerMonthlyDetail) GetMonthlyCombatRecord() *TowerMonthlyCombatRecord {
	if x != nil {
		return x.MonthlyCombatRecord
	}
	return nil
}

var File_TowerMonthlyDetail_proto protoreflect.FileDescriptor

var file_TowerMonthlyDetail_proto_rawDesc = []byte{
	0x0a, 0x18, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42,
	0x72, 0x69, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x54, 0x6f, 0x77, 0x65,
	0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x12, 0x54,
	0x6f, 0x77, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x3d, 0x0a, 0x0d, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x5f, 0x62, 0x72, 0x69,
	0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x72, 0x69,
	0x65, 0x66, 0x52, 0x0c, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x72, 0x69, 0x65, 0x66,
	0x12, 0x53, 0x0a, 0x15, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x62,
	0x61, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x13, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TowerMonthlyDetail_proto_rawDescOnce sync.Once
	file_TowerMonthlyDetail_proto_rawDescData = file_TowerMonthlyDetail_proto_rawDesc
)

func file_TowerMonthlyDetail_proto_rawDescGZIP() []byte {
	file_TowerMonthlyDetail_proto_rawDescOnce.Do(func() {
		file_TowerMonthlyDetail_proto_rawDescData = protoimpl.X.CompressGZIP(file_TowerMonthlyDetail_proto_rawDescData)
	})
	return file_TowerMonthlyDetail_proto_rawDescData
}

var file_TowerMonthlyDetail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TowerMonthlyDetail_proto_goTypes = []interface{}{
	(*TowerMonthlyDetail)(nil),       // 0: proto.TowerMonthlyDetail
	(*TowerMonthlyBrief)(nil),        // 1: proto.TowerMonthlyBrief
	(*TowerMonthlyCombatRecord)(nil), // 2: proto.TowerMonthlyCombatRecord
}
var file_TowerMonthlyDetail_proto_depIdxs = []int32{
	1, // 0: proto.TowerMonthlyDetail.monthly_brief:type_name -> proto.TowerMonthlyBrief
	2, // 1: proto.TowerMonthlyDetail.monthly_combat_record:type_name -> proto.TowerMonthlyCombatRecord
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_TowerMonthlyDetail_proto_init() }
func file_TowerMonthlyDetail_proto_init() {
	if File_TowerMonthlyDetail_proto != nil {
		return
	}
	file_TowerMonthlyBrief_proto_init()
	file_TowerMonthlyCombatRecord_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TowerMonthlyDetail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TowerMonthlyDetail); i {
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
			RawDescriptor: file_TowerMonthlyDetail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TowerMonthlyDetail_proto_goTypes,
		DependencyIndexes: file_TowerMonthlyDetail_proto_depIdxs,
		MessageInfos:      file_TowerMonthlyDetail_proto_msgTypes,
	}.Build()
	File_TowerMonthlyDetail_proto = out.File
	file_TowerMonthlyDetail_proto_rawDesc = nil
	file_TowerMonthlyDetail_proto_goTypes = nil
	file_TowerMonthlyDetail_proto_depIdxs = nil
}
