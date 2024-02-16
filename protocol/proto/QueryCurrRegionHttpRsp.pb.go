// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: QueryCurrRegionHttpRsp.proto

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

type QueryCurrRegionHttpRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode                           int32       `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Msg                               string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	RegionInfo                        *RegionInfo `protobuf:"bytes,3,opt,name=region_info,json=regionInfo,proto3" json:"region_info,omitempty"`
	ClientSecretKey                   []byte      `protobuf:"bytes,11,opt,name=client_secret_key,json=clientSecretKey,proto3" json:"client_secret_key,omitempty"`
	RegionCustomConfigEncrypted       []byte      `protobuf:"bytes,12,opt,name=region_custom_config_encrypted,json=regionCustomConfigEncrypted,proto3" json:"region_custom_config_encrypted,omitempty"`
	ClientRegionCustomConfigEncrypted []byte      `protobuf:"bytes,13,opt,name=client_region_custom_config_encrypted,json=clientRegionCustomConfigEncrypted,proto3" json:"client_region_custom_config_encrypted,omitempty"`
	// Types that are assignable to Detail:
	//
	//	*QueryCurrRegionHttpRsp_ForceUdpate
	//	*QueryCurrRegionHttpRsp_StopServer
	Detail isQueryCurrRegionHttpRsp_Detail `protobuf_oneof:"detail"`
}

func (x *QueryCurrRegionHttpRsp) Reset() {
	*x = QueryCurrRegionHttpRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QueryCurrRegionHttpRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCurrRegionHttpRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCurrRegionHttpRsp) ProtoMessage() {}

func (x *QueryCurrRegionHttpRsp) ProtoReflect() protoreflect.Message {
	mi := &file_QueryCurrRegionHttpRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryCurrRegionHttpRsp.ProtoReflect.Descriptor instead.
func (*QueryCurrRegionHttpRsp) Descriptor() ([]byte, []int) {
	return file_QueryCurrRegionHttpRsp_proto_rawDescGZIP(), []int{0}
}

func (x *QueryCurrRegionHttpRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *QueryCurrRegionHttpRsp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *QueryCurrRegionHttpRsp) GetRegionInfo() *RegionInfo {
	if x != nil {
		return x.RegionInfo
	}
	return nil
}

func (x *QueryCurrRegionHttpRsp) GetClientSecretKey() []byte {
	if x != nil {
		return x.ClientSecretKey
	}
	return nil
}

func (x *QueryCurrRegionHttpRsp) GetRegionCustomConfigEncrypted() []byte {
	if x != nil {
		return x.RegionCustomConfigEncrypted
	}
	return nil
}

func (x *QueryCurrRegionHttpRsp) GetClientRegionCustomConfigEncrypted() []byte {
	if x != nil {
		return x.ClientRegionCustomConfigEncrypted
	}
	return nil
}

func (m *QueryCurrRegionHttpRsp) GetDetail() isQueryCurrRegionHttpRsp_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (x *QueryCurrRegionHttpRsp) GetForceUdpate() *ForceUpdateInfo {
	if x, ok := x.GetDetail().(*QueryCurrRegionHttpRsp_ForceUdpate); ok {
		return x.ForceUdpate
	}
	return nil
}

func (x *QueryCurrRegionHttpRsp) GetStopServer() *StopServerInfo {
	if x, ok := x.GetDetail().(*QueryCurrRegionHttpRsp_StopServer); ok {
		return x.StopServer
	}
	return nil
}

type isQueryCurrRegionHttpRsp_Detail interface {
	isQueryCurrRegionHttpRsp_Detail()
}

type QueryCurrRegionHttpRsp_ForceUdpate struct {
	ForceUdpate *ForceUpdateInfo `protobuf:"bytes,4,opt,name=force_udpate,json=forceUdpate,proto3,oneof"`
}

type QueryCurrRegionHttpRsp_StopServer struct {
	StopServer *StopServerInfo `protobuf:"bytes,5,opt,name=stop_server,json=stopServer,proto3,oneof"`
}

func (*QueryCurrRegionHttpRsp_ForceUdpate) isQueryCurrRegionHttpRsp_Detail() {}

func (*QueryCurrRegionHttpRsp_StopServer) isQueryCurrRegionHttpRsp_Detail() {}

var File_QueryCurrRegionHttpRsp_proto protoreflect.FileDescriptor

var file_QueryCurrRegionHttpRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x75, 0x72, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x46, 0x6f, 0x72,
	0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x03, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x75,
	0x72, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x32, 0x0a, 0x0b, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x2a, 0x0a, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x43, 0x0a, 0x1e, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x1b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x12, 0x50, 0x0a, 0x25, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x21, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x12, 0x3b, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x64, 0x70, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x48, 0x00, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x55, 0x64, 0x70, 0x61, 0x74, 0x65, 0x12,
	0x38, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x6f,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0a, 0x73,
	0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_QueryCurrRegionHttpRsp_proto_rawDescOnce sync.Once
	file_QueryCurrRegionHttpRsp_proto_rawDescData = file_QueryCurrRegionHttpRsp_proto_rawDesc
)

func file_QueryCurrRegionHttpRsp_proto_rawDescGZIP() []byte {
	file_QueryCurrRegionHttpRsp_proto_rawDescOnce.Do(func() {
		file_QueryCurrRegionHttpRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_QueryCurrRegionHttpRsp_proto_rawDescData)
	})
	return file_QueryCurrRegionHttpRsp_proto_rawDescData
}

var file_QueryCurrRegionHttpRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_QueryCurrRegionHttpRsp_proto_goTypes = []interface{}{
	(*QueryCurrRegionHttpRsp)(nil), // 0: proto.QueryCurrRegionHttpRsp
	(*RegionInfo)(nil),             // 1: proto.RegionInfo
	(*ForceUpdateInfo)(nil),        // 2: proto.ForceUpdateInfo
	(*StopServerInfo)(nil),         // 3: proto.StopServerInfo
}
var file_QueryCurrRegionHttpRsp_proto_depIdxs = []int32{
	1, // 0: proto.QueryCurrRegionHttpRsp.region_info:type_name -> proto.RegionInfo
	2, // 1: proto.QueryCurrRegionHttpRsp.force_udpate:type_name -> proto.ForceUpdateInfo
	3, // 2: proto.QueryCurrRegionHttpRsp.stop_server:type_name -> proto.StopServerInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_QueryCurrRegionHttpRsp_proto_init() }
func file_QueryCurrRegionHttpRsp_proto_init() {
	if File_QueryCurrRegionHttpRsp_proto != nil {
		return
	}
	file_StopServerInfo_proto_init()
	file_ForceUpdateInfo_proto_init()
	file_RegionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_QueryCurrRegionHttpRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCurrRegionHttpRsp); i {
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
	file_QueryCurrRegionHttpRsp_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*QueryCurrRegionHttpRsp_ForceUdpate)(nil),
		(*QueryCurrRegionHttpRsp_StopServer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_QueryCurrRegionHttpRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_QueryCurrRegionHttpRsp_proto_goTypes,
		DependencyIndexes: file_QueryCurrRegionHttpRsp_proto_depIdxs,
		MessageInfos:      file_QueryCurrRegionHttpRsp_proto_msgTypes,
	}.Build()
	File_QueryCurrRegionHttpRsp_proto = out.File
	file_QueryCurrRegionHttpRsp_proto_rawDesc = nil
	file_QueryCurrRegionHttpRsp_proto_goTypes = nil
	file_QueryCurrRegionHttpRsp_proto_depIdxs = nil
}
