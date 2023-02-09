// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/dao/activity_weekly.proto

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

type ActivityWeeklyData struct {
	RoleId             string                           `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty" pk`
	ActivityWeeklyData map[int64]*models.ActivityWeekly `protobuf:"bytes,2,rep,name=activity_weekly_data,json=activityWeeklyData,proto3" json:"activity_weekly_data,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ActivityWeeklyData) Reset()      { *m = ActivityWeeklyData{} }
func (*ActivityWeeklyData) ProtoMessage() {}
func (*ActivityWeeklyData) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfc9028d5e8caa2f, []int{0}
}
func (m *ActivityWeeklyData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActivityWeeklyData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActivityWeeklyData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActivityWeeklyData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityWeeklyData.Merge(m, src)
}
func (m *ActivityWeeklyData) XXX_Size() int {
	return m.Size()
}
func (m *ActivityWeeklyData) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityWeeklyData.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityWeeklyData proto.InternalMessageInfo

func (m *ActivityWeeklyData) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *ActivityWeeklyData) GetActivityWeeklyData() map[int64]*models.ActivityWeekly {
	if m != nil {
		return m.ActivityWeeklyData
	}
	return nil
}

func (*ActivityWeeklyData) XXX_MessageName() string {
	return "dao.ActivityWeeklyData"
}

type ActivityWeeklyRankingData struct {
	ActivityId string `protobuf:"bytes,1,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty" pk`
	Version    int64  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *ActivityWeeklyRankingData) Reset()      { *m = ActivityWeeklyRankingData{} }
func (*ActivityWeeklyRankingData) ProtoMessage() {}
func (*ActivityWeeklyRankingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfc9028d5e8caa2f, []int{1}
}
func (m *ActivityWeeklyRankingData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActivityWeeklyRankingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActivityWeeklyRankingData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActivityWeeklyRankingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityWeeklyRankingData.Merge(m, src)
}
func (m *ActivityWeeklyRankingData) XXX_Size() int {
	return m.Size()
}
func (m *ActivityWeeklyRankingData) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityWeeklyRankingData.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityWeeklyRankingData proto.InternalMessageInfo

func (m *ActivityWeeklyRankingData) GetActivityId() string {
	if m != nil {
		return m.ActivityId
	}
	return ""
}

func (m *ActivityWeeklyRankingData) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (*ActivityWeeklyRankingData) XXX_MessageName() string {
	return "dao.ActivityWeeklyRankingData"
}

type ActivityWeeklyGuildData struct {
	GuildKey string `protobuf:"bytes,1,opt,name=guild_key,json=guildKey,proto3" json:"guild_key,omitempty" pk`
	Version  int64  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Score    int64  `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (m *ActivityWeeklyGuildData) Reset()      { *m = ActivityWeeklyGuildData{} }
func (*ActivityWeeklyGuildData) ProtoMessage() {}
func (*ActivityWeeklyGuildData) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfc9028d5e8caa2f, []int{2}
}
func (m *ActivityWeeklyGuildData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActivityWeeklyGuildData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActivityWeeklyGuildData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActivityWeeklyGuildData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityWeeklyGuildData.Merge(m, src)
}
func (m *ActivityWeeklyGuildData) XXX_Size() int {
	return m.Size()
}
func (m *ActivityWeeklyGuildData) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityWeeklyGuildData.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityWeeklyGuildData proto.InternalMessageInfo

func (m *ActivityWeeklyGuildData) GetGuildKey() string {
	if m != nil {
		return m.GuildKey
	}
	return ""
}

func (m *ActivityWeeklyGuildData) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ActivityWeeklyGuildData) GetScore() int64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (*ActivityWeeklyGuildData) XXX_MessageName() string {
	return "dao.ActivityWeeklyGuildData"
}
func init() {
	proto.RegisterType((*ActivityWeeklyData)(nil), "dao.ActivityWeeklyData")
	proto.RegisterMapType((map[int64]*models.ActivityWeekly)(nil), "dao.ActivityWeeklyData.ActivityWeeklyDataEntry")
	proto.RegisterType((*ActivityWeeklyRankingData)(nil), "dao.ActivityWeeklyRankingData")
	proto.RegisterType((*ActivityWeeklyGuildData)(nil), "dao.ActivityWeeklyGuildData")
}

func init() { proto.RegisterFile("proto/dao/activity_weekly.proto", fileDescriptor_dfc9028d5e8caa2f) }

var fileDescriptor_dfc9028d5e8caa2f = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x73, 0x09, 0x6d, 0xff, 0xbd, 0x2e, 0x7f, 0x8e, 0xa2, 0xb1, 0xc8, 0xb5, 0xc4, 0xc1,
	0x0e, 0x9a, 0x40, 0x5d, 0xc4, 0xcd, 0xa2, 0x48, 0x75, 0xcb, 0x22, 0x08, 0x1a, 0xce, 0xdc, 0x11,
	0x42, 0xd2, 0x5c, 0x49, 0xd2, 0x48, 0xbe, 0x85, 0x1f, 0xc3, 0x8f, 0xd2, 0xb1, 0x63, 0x27, 0xd1,
	0x64, 0x71, 0x14, 0x17, 0x57, 0xc9, 0xc5, 0x56, 0x48, 0xdb, 0xed, 0x7d, 0x9f, 0xdf, 0x9b, 0xe7,
	0x7d, 0x9f, 0x70, 0xb0, 0x3b, 0x09, 0x79, 0xcc, 0x0d, 0x4a, 0xb8, 0x41, 0xec, 0xd8, 0x4d, 0xdc,
	0x38, 0xb5, 0x9e, 0x18, 0xf3, 0xfc, 0x54, 0x17, 0x04, 0x29, 0x94, 0xf0, 0x4e, 0xdb, 0xe1, 0x0e,
	0x2f, 0x27, 0x8b, 0xaa, 0x44, 0x1d, 0xad, 0x54, 0xc6, 0x9c, 0x32, 0x3f, 0xda, 0xfc, 0xb9, 0xf6,
	0x0d, 0x20, 0x3a, 0xff, 0x25, 0xb7, 0x02, 0x5c, 0x90, 0x98, 0xa0, 0x2e, 0x6c, 0x84, 0xdc, 0x67,
	0x96, 0x4b, 0x55, 0xd0, 0x03, 0xfd, 0xe6, 0xb0, 0xfe, 0xf5, 0xda, 0x95, 0x27, 0x9e, 0x59, 0x2f,
	0xe4, 0x11, 0x45, 0x04, 0xb6, 0x2b, 0x86, 0x16, 0x25, 0x31, 0x51, 0xe5, 0x9e, 0xd2, 0x6f, 0x0d,
	0x0c, 0x9d, 0x12, 0xae, 0xaf, 0xfb, 0x6e, 0x90, 0x2e, 0x83, 0x38, 0x4c, 0x4d, 0x44, 0xd6, 0x40,
	0xe7, 0x1e, 0xee, 0x6e, 0x19, 0x47, 0xff, 0xa1, 0xe2, 0xb1, 0x54, 0x9c, 0xa6, 0x98, 0x45, 0x89,
	0x8e, 0x60, 0x2d, 0x21, 0xfe, 0x94, 0xa9, 0x72, 0x0f, 0xf4, 0x5b, 0x83, 0x1d, 0xbd, 0x4c, 0x5d,
	0x59, 0x68, 0x96, 0x43, 0x67, 0xf2, 0x29, 0xd0, 0x1e, 0xe0, 0x5e, 0x05, 0x92, 0xc0, 0x73, 0x03,
	0x47, 0xe4, 0x3f, 0x84, 0xad, 0x55, 0xbc, 0xb5, 0x7f, 0x00, 0x97, 0x68, 0x44, 0x91, 0x0a, 0x1b,
	0x09, 0x0b, 0x23, 0x97, 0x07, 0x62, 0xb3, 0x62, 0x2e, 0x5b, 0x2d, 0xa8, 0x9e, 0x7f, 0x35, 0x75,
	0x7d, 0x2a, 0xdc, 0x0f, 0x60, 0xd3, 0x29, 0x1a, 0x6b, 0x19, 0xe2, 0xcf, 0xfb, 0x9f, 0x00, 0x37,
	0x2c, 0xdd, 0xee, 0x8c, 0xda, 0xb0, 0x16, 0xd9, 0x3c, 0x64, 0xaa, 0x22, 0xf4, 0xb2, 0x19, 0x5e,
	0x2f, 0xde, 0xb1, 0xf4, 0x92, 0x61, 0x30, 0xcb, 0x30, 0x98, 0x67, 0x18, 0xbc, 0x65, 0x18, 0x7c,
	0x64, 0x58, 0xfa, 0xcc, 0x30, 0x78, 0xce, 0xb1, 0x34, 0xcb, 0x31, 0x98, 0xe7, 0x58, 0x5a, 0xe4,
	0x58, 0xba, 0xdb, 0xb7, 0xb9, 0x1b, 0x1c, 0x47, 0x2c, 0x4c, 0x58, 0x68, 0xd8, 0x7c, 0x3c, 0xe6,
	0x81, 0xb1, 0x7a, 0x6a, 0x8f, 0x75, 0x51, 0x9e, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xce, 0xd4,
	0xee, 0x3f, 0x7e, 0x02, 0x00, 0x00,
}

func init() {
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolActivityWeeklyData.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolActivityWeeklyRankingData.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolActivityWeeklyGuildData.Get().(proto.Message)
	})
}

var poolActivityWeeklyData = &sync.Pool{New: func() interface{} { return &ActivityWeeklyData{} }}

func (m *ActivityWeeklyData) ReleasePool() { m.Reset(); poolActivityWeeklyData.Put(m); m = nil }

var poolActivityWeeklyRankingData = &sync.Pool{New: func() interface{} { return &ActivityWeeklyRankingData{} }}

func (m *ActivityWeeklyRankingData) ReleasePool() {
	m.Reset()
	poolActivityWeeklyRankingData.Put(m)
	m = nil
}

var poolActivityWeeklyGuildData = &sync.Pool{New: func() interface{} { return &ActivityWeeklyGuildData{} }}

func (m *ActivityWeeklyGuildData) ReleasePool() {
	m.Reset()
	poolActivityWeeklyGuildData.Put(m)
	m = nil
}

func (m *ActivityWeeklyData) PK() string {
	if m == nil {
		return ""
	}
	return m.RoleId
}

func (m *ActivityWeeklyData) PKAppendTo(d []byte) []byte {
	if m == nil {
		return d
	}
	return append(d, m.RoleId...)
}

func (m *ActivityWeeklyData) ToKVSave() ([]byte, []byte) {
	msgName := m.XXX_MessageName()
	dk := coin_server_common_bytespool.GetSample(64)
	dk = dk[:0]
	dk = append(dk, msgName...)
	dk = append(dk, ':', 'k', ':')
	dk = m.PKAppendTo(dk)
	return dk, m.ToSave()
}

func (m *ActivityWeeklyData) ToSave() []byte {
	msgName := m.XXX_MessageName()
	ml := len(msgName)
	d := coin_server_common_bytespool.GetSample(1 + ml + m.Size())
	d[0] = uint8(ml)
	copy(d[1:], msgName)
	_, _ = m.MarshalToSizedBuffer(d[1+ml:])
	return d
}

func (m *ActivityWeeklyData) KVKey() string {
	return m.XXX_MessageName() + ":k:" + m.PK()
}

func (m *ActivityWeeklyRankingData) PK() string {
	if m == nil {
		return ""
	}
	return m.ActivityId
}

func (m *ActivityWeeklyRankingData) PKAppendTo(d []byte) []byte {
	if m == nil {
		return d
	}
	return append(d, m.ActivityId...)
}

func (m *ActivityWeeklyRankingData) ToKVSave() ([]byte, []byte) {
	msgName := m.XXX_MessageName()
	dk := coin_server_common_bytespool.GetSample(64)
	dk = dk[:0]
	dk = append(dk, msgName...)
	dk = append(dk, ':', 'k', ':')
	dk = m.PKAppendTo(dk)
	return dk, m.ToSave()
}

func (m *ActivityWeeklyRankingData) ToSave() []byte {
	msgName := m.XXX_MessageName()
	ml := len(msgName)
	d := coin_server_common_bytespool.GetSample(1 + ml + m.Size())
	d[0] = uint8(ml)
	copy(d[1:], msgName)
	_, _ = m.MarshalToSizedBuffer(d[1+ml:])
	return d
}

func (m *ActivityWeeklyRankingData) KVKey() string {
	return m.XXX_MessageName() + ":k:" + m.PK()
}

func (m *ActivityWeeklyGuildData) PK() string {
	if m == nil {
		return ""
	}
	return m.GuildKey
}

func (m *ActivityWeeklyGuildData) PKAppendTo(d []byte) []byte {
	if m == nil {
		return d
	}
	return append(d, m.GuildKey...)
}

func (m *ActivityWeeklyGuildData) ToKVSave() ([]byte, []byte) {
	msgName := m.XXX_MessageName()
	dk := coin_server_common_bytespool.GetSample(64)
	dk = dk[:0]
	dk = append(dk, msgName...)
	dk = append(dk, ':', 'k', ':')
	dk = m.PKAppendTo(dk)
	return dk, m.ToSave()
}

func (m *ActivityWeeklyGuildData) ToSave() []byte {
	msgName := m.XXX_MessageName()
	ml := len(msgName)
	d := coin_server_common_bytespool.GetSample(1 + ml + m.Size())
	d[0] = uint8(ml)
	copy(d[1:], msgName)
	_, _ = m.MarshalToSizedBuffer(d[1+ml:])
	return d
}

func (m *ActivityWeeklyGuildData) KVKey() string {
	return m.XXX_MessageName() + ":k:" + m.PK()
}

func (this *ActivityWeeklyData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ActivityWeeklyData)
	if !ok {
		that2, ok := that.(ActivityWeeklyData)
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
	if len(this.ActivityWeeklyData) != len(that1.ActivityWeeklyData) {
		return false
	}
	for i := range this.ActivityWeeklyData {
		if !this.ActivityWeeklyData[i].Equal(that1.ActivityWeeklyData[i]) {
			return false
		}
	}
	return true
}
func (this *ActivityWeeklyRankingData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ActivityWeeklyRankingData)
	if !ok {
		that2, ok := that.(ActivityWeeklyRankingData)
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
	if this.ActivityId != that1.ActivityId {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	return true
}
func (this *ActivityWeeklyGuildData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ActivityWeeklyGuildData)
	if !ok {
		that2, ok := that.(ActivityWeeklyGuildData)
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
	if this.GuildKey != that1.GuildKey {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if this.Score != that1.Score {
		return false
	}
	return true
}
func (m *ActivityWeeklyData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActivityWeeklyData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActivityWeeklyData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActivityWeeklyData) > 0 {
		for k := range m.ActivityWeeklyData {
			v := m.ActivityWeeklyData[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintActivityWeekly(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintActivityWeekly(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintActivityWeekly(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.RoleId) > 0 {
		i -= len(m.RoleId)
		copy(dAtA[i:], m.RoleId)
		i = encodeVarintActivityWeekly(dAtA, i, uint64(len(m.RoleId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ActivityWeeklyRankingData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActivityWeeklyRankingData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActivityWeeklyRankingData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintActivityWeekly(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ActivityId) > 0 {
		i -= len(m.ActivityId)
		copy(dAtA[i:], m.ActivityId)
		i = encodeVarintActivityWeekly(dAtA, i, uint64(len(m.ActivityId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ActivityWeeklyGuildData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActivityWeeklyGuildData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActivityWeeklyGuildData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Score != 0 {
		i = encodeVarintActivityWeekly(dAtA, i, uint64(m.Score))
		i--
		dAtA[i] = 0x18
	}
	if m.Version != 0 {
		i = encodeVarintActivityWeekly(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x10
	}
	if len(m.GuildKey) > 0 {
		i -= len(m.GuildKey)
		copy(dAtA[i:], m.GuildKey)
		i = encodeVarintActivityWeekly(dAtA, i, uint64(len(m.GuildKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintActivityWeekly(dAtA []byte, offset int, v uint64) int {
	offset -= sovActivityWeekly(v)
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

func (m *ActivityWeeklyData) JsonBytes(w *coin_server_common_jwriter.Writer) {
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
	if needWriteComma {
		w.RawByte(',')
	}
	w.RawByte('"')
	w.RawString("activity_weekly_data")
	w.RawByte('"')
	w.RawByte(':')
	if m.ActivityWeeklyData == nil {
		w.RawString("null")
	} else if len(m.ActivityWeeklyData) == 0 {
		w.RawString("{}")
	} else {
		w.RawByte('{')
		mlActivityWeeklyData := len(m.ActivityWeeklyData)
		for k, v := range m.ActivityWeeklyData {
			w.RawByte('"')
			w.Int64(int64(k))
			w.RawByte('"')
			w.RawByte(':')
			v.JsonBytes(w)
			mlActivityWeeklyData--
			if mlActivityWeeklyData != 0 {
				w.RawByte(',')
			}
		}
		w.RawByte('}')
	}
	needWriteComma = true
	_ = needWriteComma
	w.RawByte('}')

}

func (m *ActivityWeeklyRankingData) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.ActivityId != "" {
		w.RawByte('"')
		w.RawString("activity_id")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.ActivityId)
		needWriteComma = true
	}
	if m.Version != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("version")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.Version))
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *ActivityWeeklyGuildData) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.GuildKey != "" {
		w.RawByte('"')
		w.RawString("guild_key")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.GuildKey)
		needWriteComma = true
	}
	if m.Version != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("version")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.Version))
		needWriteComma = true
	}
	if m.Score != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("score")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.Score))
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *ActivityWeeklyData) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *ActivityWeeklyData) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *ActivityWeeklyData) GoString() string {
	return m.String()
}

func (m *ActivityWeeklyRankingData) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *ActivityWeeklyRankingData) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *ActivityWeeklyRankingData) GoString() string {
	return m.String()
}

func (m *ActivityWeeklyGuildData) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *ActivityWeeklyGuildData) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *ActivityWeeklyGuildData) GoString() string {
	return m.String()
}

func (m *ActivityWeeklyData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RoleId)
	if l > 0 {
		n += 1 + l + sovActivityWeekly(uint64(l))
	}
	if len(m.ActivityWeeklyData) > 0 {
		for k, v := range m.ActivityWeeklyData {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovActivityWeekly(uint64(l))
			}
			mapEntrySize := 1 + sovActivityWeekly(uint64(k)) + l
			n += mapEntrySize + 1 + sovActivityWeekly(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *ActivityWeeklyRankingData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ActivityId)
	if l > 0 {
		n += 1 + l + sovActivityWeekly(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovActivityWeekly(uint64(m.Version))
	}
	return n
}

func (m *ActivityWeeklyGuildData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GuildKey)
	if l > 0 {
		n += 1 + l + sovActivityWeekly(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovActivityWeekly(uint64(m.Version))
	}
	if m.Score != 0 {
		n += 1 + sovActivityWeekly(uint64(m.Score))
	}
	return n
}

func sovActivityWeekly(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozActivityWeekly(x uint64) (n int) {
	return sovActivityWeekly(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ActivityWeeklyData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActivityWeekly
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
			return fmt.Errorf("proto: ActivityWeeklyData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivityWeeklyData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityWeekly
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
				return ErrInvalidLengthActivityWeekly
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActivityWeekly
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoleId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivityWeeklyData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityWeekly
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
				return ErrInvalidLengthActivityWeekly
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthActivityWeekly
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActivityWeeklyData == nil {
				m.ActivityWeeklyData = make(map[int64]*models.ActivityWeekly)
			}
			var mapkey int64
			var mapvalue *models.ActivityWeekly
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowActivityWeekly
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowActivityWeekly
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowActivityWeekly
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthActivityWeekly
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthActivityWeekly
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &models.ActivityWeekly{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipActivityWeekly(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthActivityWeekly
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ActivityWeeklyData[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipActivityWeekly(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActivityWeekly
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
func (m *ActivityWeeklyRankingData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActivityWeekly
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
			return fmt.Errorf("proto: ActivityWeeklyRankingData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivityWeeklyRankingData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivityId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityWeekly
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
				return ErrInvalidLengthActivityWeekly
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActivityWeekly
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActivityId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityWeekly
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipActivityWeekly(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActivityWeekly
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
func (m *ActivityWeeklyGuildData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActivityWeekly
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
			return fmt.Errorf("proto: ActivityWeeklyGuildData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivityWeeklyGuildData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuildKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityWeekly
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
				return ErrInvalidLengthActivityWeekly
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthActivityWeekly
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GuildKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityWeekly
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActivityWeekly
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Score |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipActivityWeekly(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActivityWeekly
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
func skipActivityWeekly(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActivityWeekly
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
					return 0, ErrIntOverflowActivityWeekly
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
					return 0, ErrIntOverflowActivityWeekly
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
				return 0, ErrInvalidLengthActivityWeekly
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupActivityWeekly
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthActivityWeekly
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthActivityWeekly        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActivityWeekly          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupActivityWeekly = fmt.Errorf("proto: unexpected end of group")
)
