// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/broadcast/gateway-tcp.proto

package broadcast

import (
	coin_server_common_buffer "github.com/ywh147906/load-test/common/buffer"
	coin_server_common_jwriter "github.com/ywh147906/load-test/common/jwriter"
	coin_server_common_msgcreate "github.com/ywh147906/load-test/common/msgcreate"
	coin_server_common_proto_jsonany "github.com/ywh147906/load-test/common/proto/jsonany"
	_ "github.com/ywh147906/load-test/common/proto/models"
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

type GatewayStdTcp struct {
}

func (m *GatewayStdTcp) Reset()      { *m = GatewayStdTcp{} }
func (*GatewayStdTcp) ProtoMessage() {}
func (*GatewayStdTcp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1447e8427d2e6e, []int{0}
}
func (m *GatewayStdTcp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GatewayStdTcp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GatewayStdTcp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GatewayStdTcp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayStdTcp.Merge(m, src)
}
func (m *GatewayStdTcp) XXX_Size() int {
	return m.Size()
}
func (m *GatewayStdTcp) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayStdTcp.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayStdTcp proto.InternalMessageInfo

func (*GatewayStdTcp) XXX_MessageName() string {
	return "broadcast.GatewayStdTcp"
}

type GatewayStdTcp_LoginCheck struct {
	RoleId    string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	LoginTime int64  `protobuf:"varint,2,opt,name=login_time,json=loginTime,proto3" json:"login_time,omitempty"`
}

func (m *GatewayStdTcp_LoginCheck) Reset()      { *m = GatewayStdTcp_LoginCheck{} }
func (*GatewayStdTcp_LoginCheck) ProtoMessage() {}
func (*GatewayStdTcp_LoginCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1447e8427d2e6e, []int{0, 0}
}
func (m *GatewayStdTcp_LoginCheck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GatewayStdTcp_LoginCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GatewayStdTcp_LoginCheck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GatewayStdTcp_LoginCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayStdTcp_LoginCheck.Merge(m, src)
}
func (m *GatewayStdTcp_LoginCheck) XXX_Size() int {
	return m.Size()
}
func (m *GatewayStdTcp_LoginCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayStdTcp_LoginCheck.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayStdTcp_LoginCheck proto.InternalMessageInfo

func (m *GatewayStdTcp_LoginCheck) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *GatewayStdTcp_LoginCheck) GetLoginTime() int64 {
	if m != nil {
		return m.LoginTime
	}
	return 0
}

func (*GatewayStdTcp_LoginCheck) XXX_MessageName() string {
	return "broadcast.GatewayStdTcp.LoginCheck"
}

type GatewayStdTcp_KickOffUserPush struct {
	RoleId string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Status int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *GatewayStdTcp_KickOffUserPush) Reset()      { *m = GatewayStdTcp_KickOffUserPush{} }
func (*GatewayStdTcp_KickOffUserPush) ProtoMessage() {}
func (*GatewayStdTcp_KickOffUserPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f1447e8427d2e6e, []int{0, 1}
}
func (m *GatewayStdTcp_KickOffUserPush) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GatewayStdTcp_KickOffUserPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GatewayStdTcp_KickOffUserPush.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GatewayStdTcp_KickOffUserPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayStdTcp_KickOffUserPush.Merge(m, src)
}
func (m *GatewayStdTcp_KickOffUserPush) XXX_Size() int {
	return m.Size()
}
func (m *GatewayStdTcp_KickOffUserPush) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayStdTcp_KickOffUserPush.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayStdTcp_KickOffUserPush proto.InternalMessageInfo

func (m *GatewayStdTcp_KickOffUserPush) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *GatewayStdTcp_KickOffUserPush) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (*GatewayStdTcp_KickOffUserPush) XXX_MessageName() string {
	return "broadcast.GatewayStdTcp.KickOffUserPush"
}
func init() {
	proto.RegisterType((*GatewayStdTcp)(nil), "broadcast.GatewayStdTcp")
	proto.RegisterType((*GatewayStdTcp_LoginCheck)(nil), "broadcast.GatewayStdTcp.LoginCheck")
	proto.RegisterType((*GatewayStdTcp_KickOffUserPush)(nil), "broadcast.GatewayStdTcp.KickOffUserPush")
}

func init() { proto.RegisterFile("proto/broadcast/gateway-tcp.proto", fileDescriptor_9f1447e8427d2e6e) }

var fileDescriptor_9f1447e8427d2e6e = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2a, 0xca, 0x4f, 0x4c, 0x49, 0x4e, 0x2c, 0x2e, 0xd1, 0x4f, 0x4f, 0x2c, 0x49,
	0x2d, 0x4f, 0xac, 0xd4, 0x2d, 0x49, 0x2e, 0xd0, 0x03, 0xcb, 0x09, 0x71, 0xc2, 0x25, 0xa5, 0x24,
	0x21, 0xaa, 0x73, 0xf3, 0x53, 0x52, 0x73, 0x8a, 0xf5, 0x33, 0x52, 0x13, 0x53, 0x52, 0x8b, 0x20,
	0xaa, 0x94, 0x66, 0x32, 0x72, 0xf1, 0xba, 0x43, 0xf4, 0x06, 0x97, 0xa4, 0x84, 0x24, 0x17, 0x48,
	0xb9, 0x70, 0x71, 0xf9, 0xe4, 0xa7, 0x67, 0xe6, 0x39, 0x67, 0xa4, 0x26, 0x67, 0x0b, 0x89, 0x73,
	0xb1, 0x17, 0xe5, 0xe7, 0xa4, 0xc6, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xb1,
	0x81, 0xb8, 0x9e, 0x29, 0x42, 0xb2, 0x5c, 0x5c, 0x39, 0x20, 0x65, 0xf1, 0x25, 0x99, 0xb9, 0xa9,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x9c, 0x60, 0x91, 0x90, 0xcc, 0xdc, 0x54, 0x29, 0x27,
	0x2e, 0x7e, 0xef, 0xcc, 0xe4, 0x6c, 0xff, 0xb4, 0xb4, 0xd0, 0xe2, 0xd4, 0xa2, 0x80, 0xd2, 0xe2,
	0x0c, 0xdc, 0x46, 0x89, 0x71, 0xb1, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16, 0x43, 0x8d, 0x81, 0xf2,
	0x9c, 0x02, 0x6e, 0x3c, 0x94, 0x63, 0x58, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x7c, 0xf1, 0x48, 0x8e, 0xe1, 0xc3, 0x23, 0x39, 0xc6,
	0x09, 0x8f, 0xe5, 0x18, 0x4e, 0x3c, 0x96, 0x63, 0xbc, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x28, 0xa5, 0xe4, 0xfc, 0xcc, 0x3c, 0xdd, 0xe2, 0xd4, 0xa2, 0xb2, 0xd4, 0x22, 0xfd, 0xe4,
	0xfc, 0xdc, 0xdc, 0xfc, 0x3c, 0x7d, 0xb4, 0x50, 0x4a, 0x62, 0x03, 0x0b, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x3f, 0xbe, 0xfb, 0x34, 0x3f, 0x01, 0x00, 0x00,
}

func init() {
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolGatewayStdTcp.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolGatewayStdTcp_LoginCheck.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolGatewayStdTcp_KickOffUserPush.Get().(proto.Message)
	})
}

var poolGatewayStdTcp = &sync.Pool{New: func() interface{} { return &GatewayStdTcp{} }}

func (m *GatewayStdTcp) ReleasePool() { m.Reset(); poolGatewayStdTcp.Put(m); m = nil }

var poolGatewayStdTcp_LoginCheck = &sync.Pool{New: func() interface{} { return &GatewayStdTcp_LoginCheck{} }}

func (m *GatewayStdTcp_LoginCheck) ReleasePool() {
	m.Reset()
	poolGatewayStdTcp_LoginCheck.Put(m)
	m = nil
}

var poolGatewayStdTcp_KickOffUserPush = &sync.Pool{New: func() interface{} { return &GatewayStdTcp_KickOffUserPush{} }}

func (m *GatewayStdTcp_KickOffUserPush) ReleasePool() {
	m.Reset()
	poolGatewayStdTcp_KickOffUserPush.Put(m)
	m = nil
}
func (this *GatewayStdTcp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GatewayStdTcp)
	if !ok {
		that2, ok := that.(GatewayStdTcp)
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
	return true
}
func (this *GatewayStdTcp_LoginCheck) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GatewayStdTcp_LoginCheck)
	if !ok {
		that2, ok := that.(GatewayStdTcp_LoginCheck)
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
	if this.LoginTime != that1.LoginTime {
		return false
	}
	return true
}
func (this *GatewayStdTcp_KickOffUserPush) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GatewayStdTcp_KickOffUserPush)
	if !ok {
		that2, ok := that.(GatewayStdTcp_KickOffUserPush)
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
	if this.Status != that1.Status {
		return false
	}
	return true
}
func (m *GatewayStdTcp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GatewayStdTcp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GatewayStdTcp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GatewayStdTcp_LoginCheck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GatewayStdTcp_LoginCheck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GatewayStdTcp_LoginCheck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LoginTime != 0 {
		i = encodeVarintGatewayTcp(dAtA, i, uint64(m.LoginTime))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RoleId) > 0 {
		i -= len(m.RoleId)
		copy(dAtA[i:], m.RoleId)
		i = encodeVarintGatewayTcp(dAtA, i, uint64(len(m.RoleId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GatewayStdTcp_KickOffUserPush) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GatewayStdTcp_KickOffUserPush) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GatewayStdTcp_KickOffUserPush) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintGatewayTcp(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RoleId) > 0 {
		i -= len(m.RoleId)
		copy(dAtA[i:], m.RoleId)
		i = encodeVarintGatewayTcp(dAtA, i, uint64(len(m.RoleId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGatewayTcp(dAtA []byte, offset int, v uint64) int {
	offset -= sovGatewayTcp(v)
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

func (m *GatewayStdTcp) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	w.RawByte('}')

}

func (m *GatewayStdTcp_LoginCheck) JsonBytes(w *coin_server_common_jwriter.Writer) {
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
	if m.LoginTime != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("login_time")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.LoginTime))
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *GatewayStdTcp_KickOffUserPush) JsonBytes(w *coin_server_common_jwriter.Writer) {
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
	if m.Status != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("status")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.Status))
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *GatewayStdTcp) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *GatewayStdTcp) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *GatewayStdTcp) GoString() string {
	return m.String()
}

func (m *GatewayStdTcp_LoginCheck) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *GatewayStdTcp_LoginCheck) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *GatewayStdTcp_LoginCheck) GoString() string {
	return m.String()
}

func (m *GatewayStdTcp_KickOffUserPush) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *GatewayStdTcp_KickOffUserPush) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *GatewayStdTcp_KickOffUserPush) GoString() string {
	return m.String()
}

func (m *GatewayStdTcp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GatewayStdTcp_LoginCheck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RoleId)
	if l > 0 {
		n += 1 + l + sovGatewayTcp(uint64(l))
	}
	if m.LoginTime != 0 {
		n += 1 + sovGatewayTcp(uint64(m.LoginTime))
	}
	return n
}

func (m *GatewayStdTcp_KickOffUserPush) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RoleId)
	if l > 0 {
		n += 1 + l + sovGatewayTcp(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovGatewayTcp(uint64(m.Status))
	}
	return n
}

func sovGatewayTcp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGatewayTcp(x uint64) (n int) {
	return sovGatewayTcp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GatewayStdTcp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatewayTcp
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
			return fmt.Errorf("proto: GatewayStdTcp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GatewayStdTcp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGatewayTcp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGatewayTcp
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
func (m *GatewayStdTcp_LoginCheck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatewayTcp
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
			return fmt.Errorf("proto: LoginCheck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginCheck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayTcp
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
				return ErrInvalidLengthGatewayTcp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGatewayTcp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoleId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LoginTime", wireType)
			}
			m.LoginTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayTcp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LoginTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGatewayTcp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGatewayTcp
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
func (m *GatewayStdTcp_KickOffUserPush) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGatewayTcp
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
			return fmt.Errorf("proto: KickOffUserPush: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KickOffUserPush: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayTcp
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
				return ErrInvalidLengthGatewayTcp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGatewayTcp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoleId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGatewayTcp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGatewayTcp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGatewayTcp
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
func skipGatewayTcp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGatewayTcp
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
					return 0, ErrIntOverflowGatewayTcp
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
					return 0, ErrIntOverflowGatewayTcp
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
				return 0, ErrInvalidLengthGatewayTcp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGatewayTcp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGatewayTcp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGatewayTcp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGatewayTcp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGatewayTcp = fmt.Errorf("proto: unexpected end of group")
)
