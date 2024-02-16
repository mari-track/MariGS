// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: EvtBulletHitNotify.proto

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

type EvtBulletHitNotify_CmdId int32

const (
	EvtBulletHitNotify_NONE             EvtBulletHitNotify_CmdId = 0
	EvtBulletHitNotify_CMD_ID           EvtBulletHitNotify_CmdId = 313
	EvtBulletHitNotify_ENET_CHANNEL_ID  EvtBulletHitNotify_CmdId = 0
	EvtBulletHitNotify_ENET_IS_RELIABLE EvtBulletHitNotify_CmdId = 1
	EvtBulletHitNotify_IS_ALLOW_CLIENT  EvtBulletHitNotify_CmdId = 1
)

// Enum value maps for EvtBulletHitNotify_CmdId.
var (
	EvtBulletHitNotify_CmdId_name = map[int32]string{
		0:   "NONE",
		313: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
		// Duplicate value: 1: "IS_ALLOW_CLIENT",
	}
	EvtBulletHitNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           313,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
		"IS_ALLOW_CLIENT":  1,
	}
)

func (x EvtBulletHitNotify_CmdId) Enum() *EvtBulletHitNotify_CmdId {
	p := new(EvtBulletHitNotify_CmdId)
	*p = x
	return p
}

func (x EvtBulletHitNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EvtBulletHitNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_EvtBulletHitNotify_proto_enumTypes[0].Descriptor()
}

func (EvtBulletHitNotify_CmdId) Type() protoreflect.EnumType {
	return &file_EvtBulletHitNotify_proto_enumTypes[0]
}

func (x EvtBulletHitNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EvtBulletHitNotify_CmdId.Descriptor instead.
func (EvtBulletHitNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_EvtBulletHitNotify_proto_rawDescGZIP(), []int{0, 0}
}

type EvtBulletHitNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForwardType     ForwardType     `protobuf:"varint,1,opt,name=forward_type,json=forwardType,proto3,enum=proto.ForwardType" json:"forward_type,omitempty"`
	ForwardPeer     uint32          `protobuf:"varint,2,opt,name=forward_peer,json=forwardPeer,proto3" json:"forward_peer,omitempty"`
	EntityId        uint32          `protobuf:"varint,3,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	HitEntityId     uint32          `protobuf:"varint,4,opt,name=hit_entity_id,json=hitEntityId,proto3" json:"hit_entity_id,omitempty"`
	HitColliderType HitColliderType `protobuf:"varint,5,opt,name=hit_collider_type,json=hitColliderType,proto3,enum=proto.HitColliderType" json:"hit_collider_type,omitempty"`
	HitBoxIndex     int32           `protobuf:"varint,6,opt,name=hit_box_index,json=hitBoxIndex,proto3" json:"hit_box_index,omitempty"`
	HitPoint        *Vector         `protobuf:"bytes,7,opt,name=hit_point,json=hitPoint,proto3" json:"hit_point,omitempty"`
	HitNormal       *Vector         `protobuf:"bytes,8,opt,name=hit_normal,json=hitNormal,proto3" json:"hit_normal,omitempty"`
}

func (x *EvtBulletHitNotify) Reset() {
	*x = EvtBulletHitNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtBulletHitNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtBulletHitNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtBulletHitNotify) ProtoMessage() {}

func (x *EvtBulletHitNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtBulletHitNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtBulletHitNotify.ProtoReflect.Descriptor instead.
func (*EvtBulletHitNotify) Descriptor() ([]byte, []int) {
	return file_EvtBulletHitNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtBulletHitNotify) GetForwardType() ForwardType {
	if x != nil {
		return x.ForwardType
	}
	return ForwardType_FORWARD_LOCAL
}

func (x *EvtBulletHitNotify) GetForwardPeer() uint32 {
	if x != nil {
		return x.ForwardPeer
	}
	return 0
}

func (x *EvtBulletHitNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *EvtBulletHitNotify) GetHitEntityId() uint32 {
	if x != nil {
		return x.HitEntityId
	}
	return 0
}

func (x *EvtBulletHitNotify) GetHitColliderType() HitColliderType {
	if x != nil {
		return x.HitColliderType
	}
	return HitColliderType_HIT_COLLIDER_INVALID
}

func (x *EvtBulletHitNotify) GetHitBoxIndex() int32 {
	if x != nil {
		return x.HitBoxIndex
	}
	return 0
}

func (x *EvtBulletHitNotify) GetHitPoint() *Vector {
	if x != nil {
		return x.HitPoint
	}
	return nil
}

func (x *EvtBulletHitNotify) GetHitNormal() *Vector {
	if x != nil {
		return x.HitNormal
	}
	return nil
}

var File_EvtBulletHitNotify_proto protoreflect.FileDescriptor

var file_EvtBulletHitNotify_proto_rawDesc = []byte{
	0x0a, 0x18, 0x45, 0x76, 0x74, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x48, 0x69, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x48, 0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x03, 0x0a, 0x12, 0x45, 0x76,
	0x74, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x48, 0x69, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x35, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x50, 0x65, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x68, 0x69, 0x74, 0x5f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x68, 0x69, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x11, 0x68,
	0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48,
	0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f,
	0x68, 0x69, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x68, 0x69, 0x74, 0x5f, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x68, 0x69, 0x74, 0x42, 0x6f, 0x78, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x2a, 0x0a, 0x09, 0x68, 0x69, 0x74, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x68, 0x69, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x2c, 0x0a, 0x0a, 0x68, 0x69, 0x74, 0x5f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x09, 0x68, 0x69, 0x74, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x22, 0x62, 0x0a,
	0x05, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xb9, 0x02, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45,
	0x4c, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x53, 0x5f, 0x41,
	0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x02, 0x10,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvtBulletHitNotify_proto_rawDescOnce sync.Once
	file_EvtBulletHitNotify_proto_rawDescData = file_EvtBulletHitNotify_proto_rawDesc
)

func file_EvtBulletHitNotify_proto_rawDescGZIP() []byte {
	file_EvtBulletHitNotify_proto_rawDescOnce.Do(func() {
		file_EvtBulletHitNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtBulletHitNotify_proto_rawDescData)
	})
	return file_EvtBulletHitNotify_proto_rawDescData
}

var file_EvtBulletHitNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_EvtBulletHitNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtBulletHitNotify_proto_goTypes = []interface{}{
	(EvtBulletHitNotify_CmdId)(0), // 0: proto.EvtBulletHitNotify.CmdId
	(*EvtBulletHitNotify)(nil),    // 1: proto.EvtBulletHitNotify
	(ForwardType)(0),              // 2: proto.ForwardType
	(HitColliderType)(0),          // 3: proto.HitColliderType
	(*Vector)(nil),                // 4: proto.Vector
}
var file_EvtBulletHitNotify_proto_depIdxs = []int32{
	2, // 0: proto.EvtBulletHitNotify.forward_type:type_name -> proto.ForwardType
	3, // 1: proto.EvtBulletHitNotify.hit_collider_type:type_name -> proto.HitColliderType
	4, // 2: proto.EvtBulletHitNotify.hit_point:type_name -> proto.Vector
	4, // 3: proto.EvtBulletHitNotify.hit_normal:type_name -> proto.Vector
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_EvtBulletHitNotify_proto_init() }
func file_EvtBulletHitNotify_proto_init() {
	if File_EvtBulletHitNotify_proto != nil {
		return
	}
	file_Vector_proto_init()
	file_ForwardType_proto_init()
	file_HitColliderType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtBulletHitNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtBulletHitNotify); i {
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
			RawDescriptor: file_EvtBulletHitNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtBulletHitNotify_proto_goTypes,
		DependencyIndexes: file_EvtBulletHitNotify_proto_depIdxs,
		EnumInfos:         file_EvtBulletHitNotify_proto_enumTypes,
		MessageInfos:      file_EvtBulletHitNotify_proto_msgTypes,
	}.Build()
	File_EvtBulletHitNotify_proto = out.File
	file_EvtBulletHitNotify_proto_rawDesc = nil
	file_EvtBulletHitNotify_proto_goTypes = nil
	file_EvtBulletHitNotify_proto_depIdxs = nil
}
