// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.4.0
// source: WindSeedClientNotify.proto

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

type WindSeedClientNotify_CmdId int32

const (
	WindSeedClientNotify_NONE             WindSeedClientNotify_CmdId = 0
	WindSeedClientNotify_CMD_ID           WindSeedClientNotify_CmdId = 1110
	WindSeedClientNotify_ENET_CHANNEL_ID  WindSeedClientNotify_CmdId = 0
	WindSeedClientNotify_ENET_IS_RELIABLE WindSeedClientNotify_CmdId = 1
)

// Enum value maps for WindSeedClientNotify_CmdId.
var (
	WindSeedClientNotify_CmdId_name = map[int32]string{
		0:    "NONE",
		1110: "CMD_ID",
		// Duplicate value: 0: "ENET_CHANNEL_ID",
		1: "ENET_IS_RELIABLE",
	}
	WindSeedClientNotify_CmdId_value = map[string]int32{
		"NONE":             0,
		"CMD_ID":           1110,
		"ENET_CHANNEL_ID":  0,
		"ENET_IS_RELIABLE": 1,
	}
)

func (x WindSeedClientNotify_CmdId) Enum() *WindSeedClientNotify_CmdId {
	p := new(WindSeedClientNotify_CmdId)
	*p = x
	return p
}

func (x WindSeedClientNotify_CmdId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WindSeedClientNotify_CmdId) Descriptor() protoreflect.EnumDescriptor {
	return file_WindSeedClientNotify_proto_enumTypes[0].Descriptor()
}

func (WindSeedClientNotify_CmdId) Type() protoreflect.EnumType {
	return &file_WindSeedClientNotify_proto_enumTypes[0]
}

func (x WindSeedClientNotify_CmdId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WindSeedClientNotify_CmdId.Descriptor instead.
func (WindSeedClientNotify_CmdId) EnumDescriptor() ([]byte, []int) {
	return file_WindSeedClientNotify_proto_rawDescGZIP(), []int{0, 0}
}

type WindSeedClientNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Notify:
	//
	//	*WindSeedClientNotify_RefreshNotify_
	//	*WindSeedClientNotify_AddWindBulletNotify_
	//	*WindSeedClientNotify_AreaNotify_
	Notify isWindSeedClientNotify_Notify `protobuf_oneof:"notify"`
}

func (x *WindSeedClientNotify) Reset() {
	*x = WindSeedClientNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WindSeedClientNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindSeedClientNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindSeedClientNotify) ProtoMessage() {}

func (x *WindSeedClientNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WindSeedClientNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindSeedClientNotify.ProtoReflect.Descriptor instead.
func (*WindSeedClientNotify) Descriptor() ([]byte, []int) {
	return file_WindSeedClientNotify_proto_rawDescGZIP(), []int{0}
}

func (m *WindSeedClientNotify) GetNotify() isWindSeedClientNotify_Notify {
	if m != nil {
		return m.Notify
	}
	return nil
}

func (x *WindSeedClientNotify) GetRefreshNotify() *WindSeedClientNotify_RefreshNotify {
	if x, ok := x.GetNotify().(*WindSeedClientNotify_RefreshNotify_); ok {
		return x.RefreshNotify
	}
	return nil
}

func (x *WindSeedClientNotify) GetAddWindBulletNotify() *WindSeedClientNotify_AddWindBulletNotify {
	if x, ok := x.GetNotify().(*WindSeedClientNotify_AddWindBulletNotify_); ok {
		return x.AddWindBulletNotify
	}
	return nil
}

func (x *WindSeedClientNotify) GetAreaNotify() *WindSeedClientNotify_AreaNotify {
	if x, ok := x.GetNotify().(*WindSeedClientNotify_AreaNotify_); ok {
		return x.AreaNotify
	}
	return nil
}

type isWindSeedClientNotify_Notify interface {
	isWindSeedClientNotify_Notify()
}

type WindSeedClientNotify_RefreshNotify_ struct {
	RefreshNotify *WindSeedClientNotify_RefreshNotify `protobuf:"bytes,1,opt,name=refresh_notify,json=refreshNotify,proto3,oneof"`
}

type WindSeedClientNotify_AddWindBulletNotify_ struct {
	AddWindBulletNotify *WindSeedClientNotify_AddWindBulletNotify `protobuf:"bytes,2,opt,name=add_wind_bullet_notify,json=addWindBulletNotify,proto3,oneof"`
}

type WindSeedClientNotify_AreaNotify_ struct {
	AreaNotify *WindSeedClientNotify_AreaNotify `protobuf:"bytes,3,opt,name=area_notify,json=areaNotify,proto3,oneof"`
}

func (*WindSeedClientNotify_RefreshNotify_) isWindSeedClientNotify_Notify() {}

func (*WindSeedClientNotify_AddWindBulletNotify_) isWindSeedClientNotify_Notify() {}

func (*WindSeedClientNotify_AreaNotify_) isWindSeedClientNotify_Notify() {}

type WindSeedClientNotify_RefreshNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshNum uint32 `protobuf:"varint,1,opt,name=refresh_num,json=refreshNum,proto3" json:"refresh_num,omitempty"`
}

func (x *WindSeedClientNotify_RefreshNotify) Reset() {
	*x = WindSeedClientNotify_RefreshNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WindSeedClientNotify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindSeedClientNotify_RefreshNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindSeedClientNotify_RefreshNotify) ProtoMessage() {}

func (x *WindSeedClientNotify_RefreshNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WindSeedClientNotify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindSeedClientNotify_RefreshNotify.ProtoReflect.Descriptor instead.
func (*WindSeedClientNotify_RefreshNotify) Descriptor() ([]byte, []int) {
	return file_WindSeedClientNotify_proto_rawDescGZIP(), []int{0, 0}
}

func (x *WindSeedClientNotify_RefreshNotify) GetRefreshNum() uint32 {
	if x != nil {
		return x.RefreshNum
	}
	return 0
}

type WindSeedClientNotify_AddWindBulletNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeedEntityId   uint32  `protobuf:"varint,1,opt,name=seed_entity_id,json=seedEntityId,proto3" json:"seed_entity_id,omitempty"`
	SeedPos        *Vector `protobuf:"bytes,2,opt,name=seed_pos,json=seedPos,proto3" json:"seed_pos,omitempty"`
	CatchPlayerUid uint32  `protobuf:"varint,3,opt,name=catch_player_uid,json=catchPlayerUid,proto3" json:"catch_player_uid,omitempty"`
}

func (x *WindSeedClientNotify_AddWindBulletNotify) Reset() {
	*x = WindSeedClientNotify_AddWindBulletNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WindSeedClientNotify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindSeedClientNotify_AddWindBulletNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindSeedClientNotify_AddWindBulletNotify) ProtoMessage() {}

func (x *WindSeedClientNotify_AddWindBulletNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WindSeedClientNotify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindSeedClientNotify_AddWindBulletNotify.ProtoReflect.Descriptor instead.
func (*WindSeedClientNotify_AddWindBulletNotify) Descriptor() ([]byte, []int) {
	return file_WindSeedClientNotify_proto_rawDescGZIP(), []int{0, 1}
}

func (x *WindSeedClientNotify_AddWindBulletNotify) GetSeedEntityId() uint32 {
	if x != nil {
		return x.SeedEntityId
	}
	return 0
}

func (x *WindSeedClientNotify_AddWindBulletNotify) GetSeedPos() *Vector {
	if x != nil {
		return x.SeedPos
	}
	return nil
}

func (x *WindSeedClientNotify_AddWindBulletNotify) GetCatchPlayerUid() uint32 {
	if x != nil {
		return x.CatchPlayerUid
	}
	return 0
}

type WindSeedClientNotify_AreaNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreaId   uint32 `protobuf:"varint,1,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	AreaCode []byte `protobuf:"bytes,2,opt,name=area_code,json=areaCode,proto3" json:"area_code,omitempty"`
}

func (x *WindSeedClientNotify_AreaNotify) Reset() {
	*x = WindSeedClientNotify_AreaNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WindSeedClientNotify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindSeedClientNotify_AreaNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindSeedClientNotify_AreaNotify) ProtoMessage() {}

func (x *WindSeedClientNotify_AreaNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WindSeedClientNotify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindSeedClientNotify_AreaNotify.ProtoReflect.Descriptor instead.
func (*WindSeedClientNotify_AreaNotify) Descriptor() ([]byte, []int) {
	return file_WindSeedClientNotify_proto_rawDescGZIP(), []int{0, 2}
}

func (x *WindSeedClientNotify_AreaNotify) GetAreaId() uint32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *WindSeedClientNotify_AreaNotify) GetAreaCode() []byte {
	if x != nil {
		return x.AreaCode
	}
	return nil
}

var File_WindSeedClientNotify_proto protoreflect.FileDescriptor

var file_WindSeedClientNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x57, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x04, 0x0a, 0x14, 0x57,
	0x69, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x4c, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x57, 0x69,
	0x6e, 0x64, 0x53, 0x65, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x48, 0x00, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x60, 0x0a, 0x16, 0x61, 0x64, 0x64, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x5f, 0x62, 0x75,
	0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x41, 0x64, 0x64, 0x57, 0x69, 0x6e, 0x64,
	0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x48, 0x00, 0x52, 0x13,
	0x61, 0x64, 0x64, 0x57, 0x69, 0x6e, 0x64, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x43, 0x0a, 0x0b, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x57, 0x69, 0x6e, 0x64, 0x53,
	0x65, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x41, 0x72, 0x65, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x72,
	0x65, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x1a, 0x30, 0x0a, 0x0d, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4e, 0x75, 0x6d, 0x1a, 0x89, 0x01, 0x0a, 0x13, 0x41,
	0x64, 0x64, 0x57, 0x69, 0x6e, 0x64, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x65, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x65, 0x65, 0x64,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x08, 0x73, 0x65, 0x65, 0x64,
	0x5f, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x07, 0x73, 0x65, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x12, 0x28, 0x0a, 0x10,
	0x63, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x55, 0x69, 0x64, 0x1a, 0x42, 0x0a, 0x0a, 0x41, 0x72, 0x65, 0x61, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x4d, 0x0a, 0x05, 0x43, 0x6d,
	0x64, 0x49, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x06, 0x43, 0x4d, 0x44, 0x5f, 0x49, 0x44, 0x10, 0xd6, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e,
	0x45, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x45, 0x4e, 0x45, 0x54, 0x5f, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4c, 0x49, 0x41,
	0x42, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WindSeedClientNotify_proto_rawDescOnce sync.Once
	file_WindSeedClientNotify_proto_rawDescData = file_WindSeedClientNotify_proto_rawDesc
)

func file_WindSeedClientNotify_proto_rawDescGZIP() []byte {
	file_WindSeedClientNotify_proto_rawDescOnce.Do(func() {
		file_WindSeedClientNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_WindSeedClientNotify_proto_rawDescData)
	})
	return file_WindSeedClientNotify_proto_rawDescData
}

var file_WindSeedClientNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_WindSeedClientNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_WindSeedClientNotify_proto_goTypes = []interface{}{
	(WindSeedClientNotify_CmdId)(0),                  // 0: WindSeedClientNotify.CmdId
	(*WindSeedClientNotify)(nil),                     // 1: WindSeedClientNotify
	(*WindSeedClientNotify_RefreshNotify)(nil),       // 2: WindSeedClientNotify.RefreshNotify
	(*WindSeedClientNotify_AddWindBulletNotify)(nil), // 3: WindSeedClientNotify.AddWindBulletNotify
	(*WindSeedClientNotify_AreaNotify)(nil),          // 4: WindSeedClientNotify.AreaNotify
	(*Vector)(nil),                                   // 5: Vector
}
var file_WindSeedClientNotify_proto_depIdxs = []int32{
	2, // 0: WindSeedClientNotify.refresh_notify:type_name -> WindSeedClientNotify.RefreshNotify
	3, // 1: WindSeedClientNotify.add_wind_bullet_notify:type_name -> WindSeedClientNotify.AddWindBulletNotify
	4, // 2: WindSeedClientNotify.area_notify:type_name -> WindSeedClientNotify.AreaNotify
	5, // 3: WindSeedClientNotify.AddWindBulletNotify.seed_pos:type_name -> Vector
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_WindSeedClientNotify_proto_init() }
func file_WindSeedClientNotify_proto_init() {
	if File_WindSeedClientNotify_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WindSeedClientNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindSeedClientNotify); i {
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
		file_WindSeedClientNotify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindSeedClientNotify_RefreshNotify); i {
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
		file_WindSeedClientNotify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindSeedClientNotify_AddWindBulletNotify); i {
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
		file_WindSeedClientNotify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindSeedClientNotify_AreaNotify); i {
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
	file_WindSeedClientNotify_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*WindSeedClientNotify_RefreshNotify_)(nil),
		(*WindSeedClientNotify_AddWindBulletNotify_)(nil),
		(*WindSeedClientNotify_AreaNotify_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_WindSeedClientNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WindSeedClientNotify_proto_goTypes,
		DependencyIndexes: file_WindSeedClientNotify_proto_depIdxs,
		EnumInfos:         file_WindSeedClientNotify_proto_enumTypes,
		MessageInfos:      file_WindSeedClientNotify_proto_msgTypes,
	}.Build()
	File_WindSeedClientNotify_proto = out.File
	file_WindSeedClientNotify_proto_rawDesc = nil
	file_WindSeedClientNotify_proto_goTypes = nil
	file_WindSeedClientNotify_proto_depIdxs = nil
}
