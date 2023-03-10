// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/dao/equipment.proto

package dao

import (
	coin_server_common_buffer "github.com/ywh147906/load-test/common/buffer"
	coin_server_common_bytespool "github.com/ywh147906/load-test/common/bytespool"
	coin_server_common_jwriter "github.com/ywh147906/load-test/common/jwriter"
	coin_server_common_msgcreate "github.com/ywh147906/load-test/common/msgcreate"
	coin_server_common_proto_jsonany "github.com/ywh147906/load-test/common/proto/jsonany"
	models "github.com/ywh147906/load-test/common/proto/models"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	sync "sync"
	unsafe "unsafe"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// 装备信息
type Equipment struct {
	// 装备ID
	EquipId string `protobuf:"bytes,1,opt,name=equip_id,json=equipId,proto3" json:"equip_id,omitempty" pk`
	// 装备模板ID
	ItemId int64 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	// 物等
	Level int64 `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	// 词缀效果
	Affix []*models.Affix `protobuf:"bytes,4,rep,name=affix,proto3" json:"affix,omitempty"`
	// 装备在哪个英雄身上（0表示没有装备）
	HeroId      int64  `protobuf:"varint,5,opt,name=hero_id,json=heroId,proto3" json:"hero_id,omitempty"`
	ForgeName   string `protobuf:"bytes,6,opt,name=forge_name,json=forgeName,proto3" json:"forge_name,omitempty"`
	LightEffect int64  `protobuf:"varint,7,opt,name=light_effect,json=lightEffect,proto3" json:"light_effect,omitempty"`
}

func (m *Equipment) Reset()      { *m = Equipment{} }
func (*Equipment) ProtoMessage() {}
func (*Equipment) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5be86170c41e6a8, []int{0}
}
func (m *Equipment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Equipment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Equipment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Equipment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Equipment.Merge(m, src)
}
func (m *Equipment) XXX_Size() int {
	return m.Size()
}
func (m *Equipment) XXX_DiscardUnknown() {
	xxx_messageInfo_Equipment.DiscardUnknown(m)
}

var xxx_messageInfo_Equipment proto.InternalMessageInfo

func (m *Equipment) GetEquipId() string {
	if m != nil {
		return m.EquipId
	}
	return ""
}

func (m *Equipment) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *Equipment) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Equipment) GetAffix() []*models.Affix {
	if m != nil {
		return m.Affix
	}
	return nil
}

func (m *Equipment) GetHeroId() int64 {
	if m != nil {
		return m.HeroId
	}
	return 0
}

func (m *Equipment) GetForgeName() string {
	if m != nil {
		return m.ForgeName
	}
	return ""
}

func (m *Equipment) GetLightEffect() int64 {
	if m != nil {
		return m.LightEffect
	}
	return 0
}

func (*Equipment) XXX_MessageName() string {
	return "dao.Equipment"
}

// 装备摘要信息
type EquipmentBrief struct {
	// 装备id
	EquipId string `protobuf:"bytes,1,opt,name=equip_id,json=equipId,proto3" json:"equip_id,omitempty" pk`
	// 装备item id
	ItemId    int64 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	BaseScore int64 `protobuf:"varint,3,opt,name=base_score,json=baseScore,proto3" json:"base_score,omitempty"`
	Level     int64 `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	// 装备在哪个英雄身上（0表示没有装备）
	HeroId int64 `protobuf:"varint,5,opt,name=hero_id,json=heroId,proto3" json:"hero_id,omitempty"`
}

func (m *EquipmentBrief) Reset()      { *m = EquipmentBrief{} }
func (*EquipmentBrief) ProtoMessage() {}
func (*EquipmentBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5be86170c41e6a8, []int{1}
}
func (m *EquipmentBrief) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EquipmentBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EquipmentBrief.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EquipmentBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquipmentBrief.Merge(m, src)
}
func (m *EquipmentBrief) XXX_Size() int {
	return m.Size()
}
func (m *EquipmentBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_EquipmentBrief.DiscardUnknown(m)
}

var xxx_messageInfo_EquipmentBrief proto.InternalMessageInfo

func (m *EquipmentBrief) GetEquipId() string {
	if m != nil {
		return m.EquipId
	}
	return ""
}

func (m *EquipmentBrief) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *EquipmentBrief) GetBaseScore() int64 {
	if m != nil {
		return m.BaseScore
	}
	return 0
}

func (m *EquipmentBrief) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *EquipmentBrief) GetHeroId() int64 {
	if m != nil {
		return m.HeroId
	}
	return 0
}

func (*EquipmentBrief) XXX_MessageName() string {
	return "dao.EquipmentBrief"
}

type EquipId struct {
	RoleId  string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty" pk`
	EquipId int64  `protobuf:"varint,2,opt,name=equip_id,json=equipId,proto3" json:"equip_id,omitempty"`
}

func (m *EquipId) Reset()      { *m = EquipId{} }
func (*EquipId) ProtoMessage() {}
func (*EquipId) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5be86170c41e6a8, []int{2}
}
func (m *EquipId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EquipId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EquipId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EquipId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquipId.Merge(m, src)
}
func (m *EquipId) XXX_Size() int {
	return m.Size()
}
func (m *EquipId) XXX_DiscardUnknown() {
	xxx_messageInfo_EquipId.DiscardUnknown(m)
}

var xxx_messageInfo_EquipId proto.InternalMessageInfo

func (m *EquipId) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *EquipId) GetEquipId() int64 {
	if m != nil {
		return m.EquipId
	}
	return 0
}

func (*EquipId) XXX_MessageName() string {
	return "dao.EquipId"
}
func init() {
	proto.RegisterType((*Equipment)(nil), "dao.Equipment")
	proto.RegisterType((*EquipmentBrief)(nil), "dao.EquipmentBrief")
	proto.RegisterType((*EquipId)(nil), "dao.EquipId")
}

func init() { proto.RegisterFile("proto/dao/equipment.proto", fileDescriptor_a5be86170c41e6a8) }

var fileDescriptor_a5be86170c41e6a8 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbd, 0xae, 0xd3, 0x30,
	0x14, 0xc7, 0xe3, 0x9b, 0xdb, 0x84, 0xf8, 0x02, 0x83, 0x55, 0x89, 0xb4, 0x6a, 0xdd, 0x0f, 0x96,
	0x2e, 0x34, 0x12, 0x3c, 0x01, 0x95, 0x32, 0x84, 0x81, 0x21, 0x6c, 0x2c, 0x51, 0x1a, 0x9f, 0xa4,
	0x11, 0x49, 0x5c, 0x9c, 0x50, 0xf1, 0x18, 0xcc, 0x3c, 0x01, 0x8f, 0xd2, 0xb1, 0x63, 0x17, 0x10,
	0x24, 0x0b, 0x23, 0xe2, 0x09, 0x90, 0xed, 0xb6, 0x54, 0x20, 0x26, 0xb6, 0x73, 0x7e, 0xe7, 0xc3,
	0xff, 0xff, 0x91, 0xf1, 0x60, 0x2b, 0x78, 0xc3, 0x3d, 0x16, 0x73, 0x0f, 0xde, 0xbe, 0xcb, 0xb7,
	0x25, 0x54, 0xcd, 0x52, 0x31, 0x62, 0xb2, 0x98, 0x0f, 0xfb, 0x19, 0xcf, 0xb8, 0xee, 0x91, 0x91,
	0x2e, 0x0d, 0x47, 0x9a, 0x94, 0x9c, 0x41, 0x51, 0xff, 0x39, 0x38, 0xff, 0x8c, 0xb0, 0xe3, 0x9f,
	0x19, 0x99, 0xe1, 0x7b, 0xaa, 0x21, 0xca, 0x99, 0x8b, 0xa6, 0x68, 0xe1, 0xac, 0xac, 0x9f, 0x5f,
	0x26, 0x37, 0xdb, 0x37, 0xa1, 0xad, 0x78, 0xc0, 0xc8, 0x23, 0x6c, 0xe7, 0x0d, 0x94, 0xb2, 0xe3,
	0x66, 0x8a, 0x16, 0x66, 0x68, 0xc9, 0x34, 0x60, 0xa4, 0x8f, 0x7b, 0x05, 0xec, 0xa0, 0x70, 0x4d,
	0x85, 0x75, 0x42, 0x1e, 0xe3, 0x5e, 0x9c, 0xa6, 0xf9, 0x7b, 0xf7, 0x76, 0x6a, 0x2e, 0xee, 0x9e,
	0x3e, 0x58, 0x6a, 0x1d, 0xcb, 0xe7, 0x12, 0x86, 0xba, 0x26, 0x77, 0x6e, 0x40, 0x70, 0xb9, 0xb3,
	0xa7, 0x77, 0xca, 0x34, 0x60, 0x64, 0x8c, 0x71, 0xca, 0x45, 0x06, 0x51, 0x15, 0x97, 0xe0, 0x5a,
	0x52, 0x51, 0xe8, 0x28, 0xf2, 0x32, 0x2e, 0x81, 0xcc, 0xf0, 0xfd, 0x22, 0xcf, 0x36, 0x4d, 0x04,
	0x69, 0x0a, 0x49, 0xe3, 0xda, 0x6a, 0xf8, 0x4e, 0x31, 0x5f, 0xa1, 0xf9, 0x47, 0x84, 0x1f, 0x5e,
	0xfc, 0xad, 0x44, 0x0e, 0xe9, 0x7f, 0x99, 0x1c, 0x63, 0xbc, 0x8e, 0x6b, 0x88, 0xea, 0x84, 0x0b,
	0x38, 0x39, 0x75, 0x24, 0x79, 0x25, 0xc1, 0xef, 0x1b, 0xdc, 0x5e, 0xdf, 0xe0, 0x5f, 0xf6, 0xe6,
	0x3e, 0xb6, 0xfd, 0xd3, 0x8b, 0x13, 0x6c, 0x0b, 0x5e, 0xc0, 0xdf, 0x9a, 0x2c, 0x89, 0x03, 0x46,
	0x06, 0x57, 0xaa, 0xb5, 0xa6, 0xb3, 0xda, 0xd5, 0x8b, 0xe3, 0x37, 0x6a, 0x7c, 0x6a, 0x29, 0xda,
	0xb7, 0x14, 0x1d, 0x5a, 0x8a, 0xbe, 0xb6, 0x14, 0x7d, 0x6f, 0xa9, 0xf1, 0xa3, 0xa5, 0xe8, 0x43,
	0x47, 0x8d, 0x7d, 0x47, 0xd1, 0xa1, 0xa3, 0xc6, 0xb1, 0xa3, 0xc6, 0xeb, 0x51, 0xc2, 0xf3, 0xea,
	0x49, 0x0d, 0x62, 0x07, 0xc2, 0x4b, 0x78, 0x59, 0xf2, 0xca, 0xbb, 0x7c, 0xac, 0xb5, 0xa5, 0xc2,
	0x67, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x1c, 0x19, 0x33, 0x6c, 0x02, 0x00, 0x00,
}

func init() {
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolEquipment.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolEquipmentBrief.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolEquipId.Get().(proto.Message)
	})
}

var poolEquipment = &sync.Pool{New: func() interface{} { return &Equipment{} }}

func (m *Equipment) ReleasePool() { m.Reset(); poolEquipment.Put(m); m = nil }

var poolEquipmentBrief = &sync.Pool{New: func() interface{} { return &EquipmentBrief{} }}

func (m *EquipmentBrief) ReleasePool() { m.Reset(); poolEquipmentBrief.Put(m); m = nil }

var poolEquipId = &sync.Pool{New: func() interface{} { return &EquipId{} }}

func (m *EquipId) ReleasePool() { m.Reset(); poolEquipId.Put(m); m = nil }

func (m *Equipment) PK() string {
	if m == nil {
		return ""
	}
	return m.EquipId
}

func (m *Equipment) PKAppendTo(d []byte) []byte {
	if m == nil {
		return d
	}
	return append(d, m.EquipId...)
}

func (m *Equipment) ToKVSave() ([]byte, []byte) {
	msgName := m.XXX_MessageName()
	dk := coin_server_common_bytespool.GetSample(64)
	dk = dk[:0]
	dk = append(dk, msgName...)
	dk = append(dk, ':', 'k', ':')
	dk = m.PKAppendTo(dk)
	return dk, m.ToSave()
}

func (m *Equipment) ToSave() []byte {
	msgName := m.XXX_MessageName()
	ml := len(msgName)
	d := coin_server_common_bytespool.GetSample(1 + ml + m.Size())
	d[0] = uint8(ml)
	copy(d[1:], msgName)
	_, _ = m.MarshalToSizedBuffer(d[1+ml:])
	return d
}

func (m *Equipment) KVKey() string {
	return m.XXX_MessageName() + ":k:" + m.PK()
}

func (m *EquipmentBrief) PK() string {
	if m == nil {
		return ""
	}
	return m.EquipId
}

func (m *EquipmentBrief) PKAppendTo(d []byte) []byte {
	if m == nil {
		return d
	}
	return append(d, m.EquipId...)
}

func (m *EquipmentBrief) ToKVSave() ([]byte, []byte) {
	msgName := m.XXX_MessageName()
	dk := coin_server_common_bytespool.GetSample(64)
	dk = dk[:0]
	dk = append(dk, msgName...)
	dk = append(dk, ':', 'k', ':')
	dk = m.PKAppendTo(dk)
	return dk, m.ToSave()
}

func (m *EquipmentBrief) ToSave() []byte {
	msgName := m.XXX_MessageName()
	ml := len(msgName)
	d := coin_server_common_bytespool.GetSample(1 + ml + m.Size())
	d[0] = uint8(ml)
	copy(d[1:], msgName)
	_, _ = m.MarshalToSizedBuffer(d[1+ml:])
	return d
}

func (m *EquipmentBrief) KVKey() string {
	return m.XXX_MessageName() + ":k:" + m.PK()
}

func (m *EquipId) PK() string {
	if m == nil {
		return ""
	}
	return m.RoleId
}

func (m *EquipId) PKAppendTo(d []byte) []byte {
	if m == nil {
		return d
	}
	return append(d, m.RoleId...)
}

func (m *EquipId) ToKVSave() ([]byte, []byte) {
	msgName := m.XXX_MessageName()
	dk := coin_server_common_bytespool.GetSample(64)
	dk = dk[:0]
	dk = append(dk, msgName...)
	dk = append(dk, ':', 'k', ':')
	dk = m.PKAppendTo(dk)
	return dk, m.ToSave()
}

func (m *EquipId) ToSave() []byte {
	msgName := m.XXX_MessageName()
	ml := len(msgName)
	d := coin_server_common_bytespool.GetSample(1 + ml + m.Size())
	d[0] = uint8(ml)
	copy(d[1:], msgName)
	_, _ = m.MarshalToSizedBuffer(d[1+ml:])
	return d
}

func (m *EquipId) KVKey() string {
	return m.XXX_MessageName() + ":k:" + m.PK()
}

func (this *Equipment) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Equipment)
	if !ok {
		that2, ok := that.(Equipment)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.EquipId != that1.EquipId {
		return false
	}
	if this.ItemId != that1.ItemId {
		return false
	}
	if this.Level != that1.Level {
		return false
	}
	if len(this.Affix) != len(that1.Affix) {
		return false
	}
	for i := range this.Affix {
		if !this.Affix[i].Equal(that1.Affix[i]) {
			return false
		}
	}
	if this.HeroId != that1.HeroId {
		return false
	}
	if this.ForgeName != that1.ForgeName {
		return false
	}
	if this.LightEffect != that1.LightEffect {
		return false
	}
	return true
}
func (this *EquipmentBrief) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EquipmentBrief)
	if !ok {
		that2, ok := that.(EquipmentBrief)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.EquipId != that1.EquipId {
		return false
	}
	if this.ItemId != that1.ItemId {
		return false
	}
	if this.BaseScore != that1.BaseScore {
		return false
	}
	if this.Level != that1.Level {
		return false
	}
	if this.HeroId != that1.HeroId {
		return false
	}
	return true
}
func (this *EquipId) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EquipId)
	if !ok {
		that2, ok := that.(EquipId)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.RoleId != that1.RoleId {
		return false
	}
	if this.EquipId != that1.EquipId {
		return false
	}
	return true
}
func (m *Equipment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Equipment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Equipment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LightEffect != 0 {
		i = encodeVarintEquipment(dAtA, i, uint64(m.LightEffect))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ForgeName) > 0 {
		i -= len(m.ForgeName)
		copy(dAtA[i:], m.ForgeName)
		i = encodeVarintEquipment(dAtA, i, uint64(len(m.ForgeName)))
		i--
		dAtA[i] = 0x32
	}
	if m.HeroId != 0 {
		i = encodeVarintEquipment(dAtA, i, uint64(m.HeroId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Affix) > 0 {
		for iNdEx := len(m.Affix) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Affix[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEquipment(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Level != 0 {
		i = encodeVarintEquipment(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x18
	}
	if m.ItemId != 0 {
		i = encodeVarintEquipment(dAtA, i, uint64(m.ItemId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.EquipId) > 0 {
		i -= len(m.EquipId)
		copy(dAtA[i:], m.EquipId)
		i = encodeVarintEquipment(dAtA, i, uint64(len(m.EquipId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EquipmentBrief) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EquipmentBrief) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EquipmentBrief) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HeroId != 0 {
		i = encodeVarintEquipment(dAtA, i, uint64(m.HeroId))
		i--
		dAtA[i] = 0x28
	}
	if m.Level != 0 {
		i = encodeVarintEquipment(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x20
	}
	if m.BaseScore != 0 {
		i = encodeVarintEquipment(dAtA, i, uint64(m.BaseScore))
		i--
		dAtA[i] = 0x18
	}
	if m.ItemId != 0 {
		i = encodeVarintEquipment(dAtA, i, uint64(m.ItemId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.EquipId) > 0 {
		i -= len(m.EquipId)
		copy(dAtA[i:], m.EquipId)
		i = encodeVarintEquipment(dAtA, i, uint64(len(m.EquipId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EquipId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EquipId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EquipId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EquipId != 0 {
		i = encodeVarintEquipment(dAtA, i, uint64(m.EquipId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RoleId) > 0 {
		i -= len(m.RoleId)
		copy(dAtA[i:], m.RoleId)
		i = encodeVarintEquipment(dAtA, i, uint64(len(m.RoleId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEquipment(dAtA []byte, offset int, v uint64) int {
	offset -= sovEquipment(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

var _ = coin_server_common_proto_jsonany.Any{}

func (m *Equipment) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.EquipId != "" {
		w.RawByte('"')
		w.RawString("equip_id")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.EquipId)
		needWriteComma = true
	}
	if m.ItemId != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("item_id")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.ItemId))
		needWriteComma = true
	}
	if m.Level != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("level")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.Level))
		needWriteComma = true
	}
	if needWriteComma {
		w.RawByte(',')
	}
	w.RawByte('"')
	w.RawString("affix")
	w.RawByte('"')
	w.RawByte(':')
	if m.Affix == nil {
		w.RawString("null")
	} else if len(m.Affix) == 0 {
		w.RawString("[]")
	} else {
		w.RawByte('[')
		for i, v := range m.Affix {
			v.JsonBytes(w)
			if i != len(m.Affix)-1 {
				w.RawByte(',')
			}
		}
		w.RawByte(']')
	}
	needWriteComma = true
	if m.HeroId != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("hero_id")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.HeroId))
		needWriteComma = true
	}
	if m.ForgeName != "" {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("forge_name")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.ForgeName)
		needWriteComma = true
	}
	if m.LightEffect != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("light_effect")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.LightEffect))
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *EquipmentBrief) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.EquipId != "" {
		w.RawByte('"')
		w.RawString("equip_id")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.EquipId)
		needWriteComma = true
	}
	if m.ItemId != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("item_id")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.ItemId))
		needWriteComma = true
	}
	if m.BaseScore != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("base_score")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.BaseScore))
		needWriteComma = true
	}
	if m.Level != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("level")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.Level))
		needWriteComma = true
	}
	if m.HeroId != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("hero_id")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.HeroId))
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *EquipId) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.RoleId != "" {
		w.RawByte('"')
		w.RawString("role_id")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.RoleId)
		needWriteComma = true
	}
	if m.EquipId != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("equip_id")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.EquipId))
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *Equipment) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *Equipment) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *Equipment) GoString() string {
	return m.String()
}

func (m *EquipmentBrief) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *EquipmentBrief) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *EquipmentBrief) GoString() string {
	return m.String()
}

func (m *EquipId) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *EquipId) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *EquipId) GoString() string {
	return m.String()
}

func (m *Equipment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EquipId)
	if l > 0 {
		n += 1 + l + sovEquipment(uint64(l))
	}
	if m.ItemId != 0 {
		n += 1 + sovEquipment(uint64(m.ItemId))
	}
	if m.Level != 0 {
		n += 1 + sovEquipment(uint64(m.Level))
	}
	if len(m.Affix) > 0 {
		for _, e := range m.Affix {
			l = e.Size()
			n += 1 + l + sovEquipment(uint64(l))
		}
	}
	if m.HeroId != 0 {
		n += 1 + sovEquipment(uint64(m.HeroId))
	}
	l = len(m.ForgeName)
	if l > 0 {
		n += 1 + l + sovEquipment(uint64(l))
	}
	if m.LightEffect != 0 {
		n += 1 + sovEquipment(uint64(m.LightEffect))
	}
	return n
}

func (m *EquipmentBrief) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EquipId)
	if l > 0 {
		n += 1 + l + sovEquipment(uint64(l))
	}
	if m.ItemId != 0 {
		n += 1 + sovEquipment(uint64(m.ItemId))
	}
	if m.BaseScore != 0 {
		n += 1 + sovEquipment(uint64(m.BaseScore))
	}
	if m.Level != 0 {
		n += 1 + sovEquipment(uint64(m.Level))
	}
	if m.HeroId != 0 {
		n += 1 + sovEquipment(uint64(m.HeroId))
	}
	return n
}

func (m *EquipId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RoleId)
	if l > 0 {
		n += 1 + l + sovEquipment(uint64(l))
	}
	if m.EquipId != 0 {
		n += 1 + sovEquipment(uint64(m.EquipId))
	}
	return n
}

func sovEquipment(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEquipment(x uint64) (n int) {
	return sovEquipment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Equipment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEquipment
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Equipment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Equipment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EquipId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEquipment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEquipment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EquipId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemId", wireType)
			}
			m.ItemId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ItemId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Affix", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEquipment
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEquipment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Affix = append(m.Affix, &models.Affix{})
			if err := m.Affix[len(m.Affix)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeroId", wireType)
			}
			m.HeroId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeroId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForgeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEquipment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEquipment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ForgeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LightEffect", wireType)
			}
			m.LightEffect = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LightEffect |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEquipment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEquipment
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EquipmentBrief) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEquipment
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EquipmentBrief: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EquipmentBrief: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EquipId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEquipment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEquipment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EquipId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemId", wireType)
			}
			m.ItemId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ItemId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseScore", wireType)
			}
			m.BaseScore = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BaseScore |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeroId", wireType)
			}
			m.HeroId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeroId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEquipment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEquipment
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EquipId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEquipment
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EquipId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EquipId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEquipment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEquipment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoleId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EquipId", wireType)
			}
			m.EquipId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EquipId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEquipment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEquipment
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEquipment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEquipment
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEquipment
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEquipment
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEquipment
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEquipment
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEquipment        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEquipment          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEquipment = fmt.Errorf("proto: unexpected end of group")
)
