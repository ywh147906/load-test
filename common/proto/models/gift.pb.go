// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/models/gift.proto

package models

import (
	coin_server_common_buffer "github.com/ywh147906/load-test/common/buffer"
	coin_server_common_jwriter "github.com/ywh147906/load-test/common/jwriter"
	coin_server_common_msgcreate "github.com/ywh147906/load-test/common/msgcreate"
	coin_server_common_proto_jsonany "github.com/ywh147906/load-test/common/proto/jsonany"
	fmt "fmt"
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

type Gift struct {
	GiftId    string        `protobuf:"bytes,1,opt,name=gift_id,json=giftId,proto3" json:"gift_id,omitempty"`
	GiftNo    int64         `protobuf:"zigzag64,2,opt,name=gift_no,json=giftNo,proto3" json:"gift_no,omitempty"`
	RoleId    string        `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	DrawCount int64         `protobuf:"zigzag64,4,opt,name=draw_count,json=drawCount,proto3" json:"draw_count,omitempty"`
	Records   []*DrawRecord `protobuf:"bytes,5,rep,name=records,proto3" json:"records,omitempty"`
	CreateAt  int64         `protobuf:"zigzag64,6,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (m *Gift) Reset()      { *m = Gift{} }
func (*Gift) ProtoMessage() {}
func (*Gift) Descriptor() ([]byte, []int) {
	return fileDescriptor_687fd520f89ec088, []int{0}
}
func (m *Gift) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Gift) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Gift.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Gift) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gift.Merge(m, src)
}
func (m *Gift) XXX_Size() int {
	return m.Size()
}
func (m *Gift) XXX_DiscardUnknown() {
	xxx_messageInfo_Gift.DiscardUnknown(m)
}

var xxx_messageInfo_Gift proto.InternalMessageInfo

func (m *Gift) GetGiftId() string {
	if m != nil {
		return m.GiftId
	}
	return ""
}

func (m *Gift) GetGiftNo() int64 {
	if m != nil {
		return m.GiftNo
	}
	return 0
}

func (m *Gift) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *Gift) GetDrawCount() int64 {
	if m != nil {
		return m.DrawCount
	}
	return 0
}

func (m *Gift) GetRecords() []*DrawRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *Gift) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (*Gift) XXX_MessageName() string {
	return "models.Gift"
}

type DrawRecord struct {
	RoleId   string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Item     *Item  `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	DrawTime int64  `protobuf:"zigzag64,3,opt,name=draw_time,json=drawTime,proto3" json:"draw_time,omitempty"`
	Role     *Role  `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
}

func (m *DrawRecord) Reset()      { *m = DrawRecord{} }
func (*DrawRecord) ProtoMessage() {}
func (*DrawRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_687fd520f89ec088, []int{1}
}
func (m *DrawRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DrawRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DrawRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DrawRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DrawRecord.Merge(m, src)
}
func (m *DrawRecord) XXX_Size() int {
	return m.Size()
}
func (m *DrawRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DrawRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DrawRecord proto.InternalMessageInfo

func (m *DrawRecord) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *DrawRecord) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *DrawRecord) GetDrawTime() int64 {
	if m != nil {
		return m.DrawTime
	}
	return 0
}

func (m *DrawRecord) GetRole() *Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (*DrawRecord) XXX_MessageName() string {
	return "models.DrawRecord"
}
func init() {
	proto.RegisterType((*Gift)(nil), "models.Gift")
	proto.RegisterType((*DrawRecord)(nil), "models.DrawRecord")
}

func init() { proto.RegisterFile("proto/models/gift.proto", fileDescriptor_687fd520f89ec088) }

var fileDescriptor_687fd520f89ec088 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbf, 0x4e, 0x02, 0x31,
	0x1c, 0xc7, 0xaf, 0x82, 0x07, 0x14, 0xa7, 0x0e, 0x72, 0xc1, 0x58, 0x2f, 0x4c, 0x0c, 0xca, 0x25,
	0xf8, 0x04, 0xfe, 0x49, 0x0c, 0x83, 0x0e, 0x8d, 0x93, 0x0b, 0x39, 0xee, 0x7e, 0x90, 0x26, 0xf4,
	0x6a, 0x7a, 0x45, 0xde, 0xc0, 0xd9, 0xc7, 0xf0, 0x1d, 0x7c, 0x01, 0x46, 0x46, 0x46, 0xed, 0x2d,
	0x8e, 0x3e, 0x82, 0x69, 0x1b, 0x14, 0xc6, 0xdf, 0xf7, 0xd3, 0xcf, 0xef, 0x4f, 0x8a, 0x3b, 0xcf,
	0x4a, 0x6a, 0x99, 0x08, 0x99, 0xc3, 0xbc, 0x4c, 0x66, 0x7c, 0xaa, 0x07, 0x2e, 0x21, 0xa1, 0x8f,
	0xba, 0xc7, 0x7b, 0x0f, 0x26, 0xe9, 0xcc, 0xf3, 0xee, 0xbe, 0xb8, 0x28, 0x41, 0x79, 0xd0, 0xfb,
	0x40, 0xb8, 0x7e, 0xc7, 0xa7, 0x9a, 0x74, 0x70, 0xc3, 0xf6, 0x1b, 0xf3, 0x3c, 0x42, 0x31, 0xea,
	0xb7, 0x58, 0x68, 0xcb, 0x51, 0xfe, 0x07, 0x0a, 0x19, 0x1d, 0xc4, 0xa8, 0x4f, 0x3c, 0x78, 0x90,
	0x16, 0x28, 0x39, 0x07, 0x6b, 0xd4, 0xbc, 0x61, 0xcb, 0x51, 0x4e, 0x4e, 0x31, 0xce, 0x55, 0xba,
	0x1c, 0x67, 0x72, 0x51, 0xe8, 0xa8, 0xee, 0xa4, 0x96, 0x4d, 0x6e, 0x6c, 0x40, 0xce, 0x71, 0x43,
	0x41, 0x26, 0x55, 0x5e, 0x46, 0x87, 0x71, 0xad, 0xdf, 0x1e, 0x92, 0x81, 0xdf, 0x6b, 0x70, 0xab,
	0xd2, 0x25, 0x73, 0x88, 0x6d, 0x9f, 0x90, 0x13, 0xdc, 0xca, 0x14, 0xa4, 0x1a, 0xc6, 0xa9, 0x8e,
	0x42, 0xd7, 0xab, 0xe9, 0x83, 0x2b, 0xdd, 0x7b, 0x45, 0x18, 0xff, 0x4b, 0xbb, 0x1b, 0xa1, 0xbd,
	0x8d, 0x62, 0x5c, 0xe7, 0x1a, 0x84, 0x3b, 0xa0, 0x3d, 0x3c, 0xda, 0xce, 0x1b, 0x69, 0x10, 0xcc,
	0x11, 0x3b, 0xc6, 0xed, 0xac, 0xb9, 0x00, 0x77, 0x0e, 0x61, 0x4d, 0x1b, 0x3c, 0x72, 0x01, 0x56,
	0xb7, 0x8d, 0xdc, 0x29, 0x3b, 0x3a, 0x93, 0x73, 0x60, 0x8e, 0x5c, 0xdf, 0x6f, 0xbe, 0x68, 0xf0,
	0x6e, 0x28, 0x5a, 0x19, 0x8a, 0xd6, 0x86, 0xa2, 0x4f, 0x43, 0xd1, 0xb7, 0xa1, 0xc1, 0x8f, 0xa1,
	0xe8, 0xad, 0xa2, 0xc1, 0xaa, 0xa2, 0x68, 0x5d, 0xd1, 0x60, 0x53, 0xd1, 0xe0, 0xe9, 0x2c, 0x93,
	0xbc, 0xb8, 0x28, 0x41, 0xbd, 0x80, 0x4a, 0x32, 0x29, 0x84, 0x2c, 0x92, 0xdd, 0x0f, 0x9a, 0x84,
	0xae, 0xba, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x01, 0x74, 0xd4, 0xf0, 0x01, 0x00, 0x00,
}

func init() {
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolGift.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolDrawRecord.Get().(proto.Message)
	})
}

var poolGift = &sync.Pool{New: func() interface{} { return &Gift{} }}

func (m *Gift) ReleasePool() { m.Reset(); poolGift.Put(m); m = nil }

var poolDrawRecord = &sync.Pool{New: func() interface{} { return &DrawRecord{} }}

func (m *DrawRecord) ReleasePool() { m.Reset(); poolDrawRecord.Put(m); m = nil }
func (this *Gift) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Gift)
	if !ok {
		that2, ok := that.(Gift)
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
	if this.GiftId != that1.GiftId {
		return false
	}
	if this.GiftNo != that1.GiftNo {
		return false
	}
	if this.RoleId != that1.RoleId {
		return false
	}
	if this.DrawCount != that1.DrawCount {
		return false
	}
	if len(this.Records) != len(that1.Records) {
		return false
	}
	for i := range this.Records {
		if !this.Records[i].Equal(that1.Records[i]) {
			return false
		}
	}
	if this.CreateAt != that1.CreateAt {
		return false
	}
	return true
}
func (this *DrawRecord) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DrawRecord)
	if !ok {
		that2, ok := that.(DrawRecord)
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
	if !this.Item.Equal(that1.Item) {
		return false
	}
	if this.DrawTime != that1.DrawTime {
		return false
	}
	if !this.Role.Equal(that1.Role) {
		return false
	}
	return true
}
func (m *Gift) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Gift) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Gift) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreateAt != 0 {
		i = encodeVarintGift(dAtA, i, uint64((uint64(m.CreateAt)<<1)^uint64((m.CreateAt>>63))))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGift(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.DrawCount != 0 {
		i = encodeVarintGift(dAtA, i, uint64((uint64(m.DrawCount)<<1)^uint64((m.DrawCount>>63))))
		i--
		dAtA[i] = 0x20
	}
	if len(m.RoleId) > 0 {
		i -= len(m.RoleId)
		copy(dAtA[i:], m.RoleId)
		i = encodeVarintGift(dAtA, i, uint64(len(m.RoleId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.GiftNo != 0 {
		i = encodeVarintGift(dAtA, i, uint64((uint64(m.GiftNo)<<1)^uint64((m.GiftNo>>63))))
		i--
		dAtA[i] = 0x10
	}
	if len(m.GiftId) > 0 {
		i -= len(m.GiftId)
		copy(dAtA[i:], m.GiftId)
		i = encodeVarintGift(dAtA, i, uint64(len(m.GiftId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DrawRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DrawRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DrawRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Role != nil {
		{
			size, err := m.Role.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGift(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.DrawTime != 0 {
		i = encodeVarintGift(dAtA, i, uint64((uint64(m.DrawTime)<<1)^uint64((m.DrawTime>>63))))
		i--
		dAtA[i] = 0x18
	}
	if m.Item != nil {
		{
			size, err := m.Item.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGift(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RoleId) > 0 {
		i -= len(m.RoleId)
		copy(dAtA[i:], m.RoleId)
		i = encodeVarintGift(dAtA, i, uint64(len(m.RoleId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGift(dAtA []byte, offset int, v uint64) int {
	offset -= sovGift(v)
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

func (m *Gift) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.GiftId != "" {
		w.RawByte('"')
		w.RawString("gift_id")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.GiftId)
		needWriteComma = true
	}
	if m.GiftNo != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("gift_no")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.GiftNo))
		needWriteComma = true
	}
	if m.RoleId != "" {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("role_id")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.RoleId)
		needWriteComma = true
	}
	if m.DrawCount != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("draw_count")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.DrawCount))
		needWriteComma = true
	}
	if needWriteComma {
		w.RawByte(',')
	}
	w.RawByte('"')
	w.RawString("records")
	w.RawByte('"')
	w.RawByte(':')
	if m.Records == nil {
		w.RawString("null")
	} else if len(m.Records) == 0 {
		w.RawString("[]")
	} else {
		w.RawByte('[')
		for i, v := range m.Records {
			v.JsonBytes(w)
			if i != len(m.Records)-1 {
				w.RawByte(',')
			}
		}
		w.RawByte(']')
	}
	needWriteComma = true
	if m.CreateAt != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("create_at")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.CreateAt))
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *DrawRecord) JsonBytes(w *coin_server_common_jwriter.Writer) {
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
	w.RawString("item")
	w.RawByte('"')
	w.RawByte(':')
	m.Item.JsonBytes(w)
	needWriteComma = true
	if m.DrawTime != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("draw_time")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.DrawTime))
		needWriteComma = true
	}
	if needWriteComma {
		w.RawByte(',')
	}
	w.RawByte('"')
	w.RawString("role")
	w.RawByte('"')
	w.RawByte(':')
	m.Role.JsonBytes(w)
	needWriteComma = true
	_ = needWriteComma
	w.RawByte('}')

}

func (m *Gift) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *Gift) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *Gift) GoString() string {
	return m.String()
}

func (m *DrawRecord) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *DrawRecord) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *DrawRecord) GoString() string {
	return m.String()
}

func (m *Gift) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GiftId)
	if l > 0 {
		n += 1 + l + sovGift(uint64(l))
	}
	if m.GiftNo != 0 {
		n += 1 + sozGift(uint64(m.GiftNo))
	}
	l = len(m.RoleId)
	if l > 0 {
		n += 1 + l + sovGift(uint64(l))
	}
	if m.DrawCount != 0 {
		n += 1 + sozGift(uint64(m.DrawCount))
	}
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovGift(uint64(l))
		}
	}
	if m.CreateAt != 0 {
		n += 1 + sozGift(uint64(m.CreateAt))
	}
	return n
}

func (m *DrawRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RoleId)
	if l > 0 {
		n += 1 + l + sovGift(uint64(l))
	}
	if m.Item != nil {
		l = m.Item.Size()
		n += 1 + l + sovGift(uint64(l))
	}
	if m.DrawTime != 0 {
		n += 1 + sozGift(uint64(m.DrawTime))
	}
	if m.Role != nil {
		l = m.Role.Size()
		n += 1 + l + sovGift(uint64(l))
	}
	return n
}

func sovGift(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGift(x uint64) (n int) {
	return sovGift(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Gift) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGift
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
			return fmt.Errorf("proto: Gift: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Gift: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GiftId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGift
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
				return ErrInvalidLengthGift
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGift
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GiftId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GiftNo", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGift
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.GiftNo = int64(v)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGift
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
				return ErrInvalidLengthGift
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGift
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoleId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DrawCount", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGift
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.DrawCount = int64(v)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGift
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
				return ErrInvalidLengthGift
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGift
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, &DrawRecord{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateAt", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGift
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.CreateAt = int64(v)
		default:
			iNdEx = preIndex
			skippy, err := skipGift(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGift
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
func (m *DrawRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGift
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
			return fmt.Errorf("proto: DrawRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DrawRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGift
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
				return ErrInvalidLengthGift
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGift
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoleId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Item", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGift
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
				return ErrInvalidLengthGift
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGift
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Item == nil {
				m.Item = &Item{}
			}
			if err := m.Item.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DrawTime", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGift
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.DrawTime = int64(v)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGift
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
				return ErrInvalidLengthGift
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGift
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Role == nil {
				m.Role = &Role{}
			}
			if err := m.Role.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGift(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGift
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
func skipGift(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGift
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
					return 0, ErrIntOverflowGift
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
					return 0, ErrIntOverflowGift
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
				return 0, ErrInvalidLengthGift
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGift
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGift
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGift        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGift          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGift = fmt.Errorf("proto: unexpected end of group")
)
