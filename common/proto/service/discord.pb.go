// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/service/discord.proto

package service

import (
	coin_server_common_buffer "github.com/ywh147906/load-test/common/buffer"
	coin_server_common_jwriter "github.com/ywh147906/load-test/common/jwriter"
	coin_server_common_msgcreate "github.com/ywh147906/load-test/common/msgcreate"
	coin_server_common_proto_jsonany "github.com/ywh147906/load-test/common/proto/jsonany"
	models "github.com/ywh147906/load-test/common/proto/models"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	strconv "strconv"
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

type DiscordErrorCode int32

const (
	DiscordErrorCode_ErrDiscordSystem   DiscordErrorCode = 0
	DiscordErrorCode_ErrDiscordNotData  DiscordErrorCode = 1
	DiscordErrorCode_ErrDiscordConfig   DiscordErrorCode = 2
	DiscordErrorCode_ErrDiscordReceived DiscordErrorCode = 3
)

var DiscordErrorCode_name = map[int32]string{
	0: "ErrDiscordSystem",
	1: "ErrDiscordNotData",
	2: "ErrDiscordConfig",
	3: "ErrDiscordReceived",
}

var DiscordErrorCode_value = map[string]int32{
	"ErrDiscordSystem":   0,
	"ErrDiscordNotData":  1,
	"ErrDiscordConfig":   2,
	"ErrDiscordReceived": 3,
}

func (DiscordErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b9a2d93676bfa42e, []int{0}
}

type Discord struct {
}

func (m *Discord) Reset()      { *m = Discord{} }
func (*Discord) ProtoMessage() {}
func (*Discord) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a2d93676bfa42e, []int{0}
}
func (m *Discord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Discord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Discord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Discord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Discord.Merge(m, src)
}
func (m *Discord) XXX_Size() int {
	return m.Size()
}
func (m *Discord) XXX_DiscardUnknown() {
	xxx_messageInfo_Discord.DiscardUnknown(m)
}

var xxx_messageInfo_Discord proto.InternalMessageInfo

func (*Discord) XXX_MessageName() string {
	return "service.Discord"
}

// 查询情况
type Discord_DiscordInfoRequest struct {
}

func (m *Discord_DiscordInfoRequest) Reset()      { *m = Discord_DiscordInfoRequest{} }
func (*Discord_DiscordInfoRequest) ProtoMessage() {}
func (*Discord_DiscordInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a2d93676bfa42e, []int{0, 0}
}
func (m *Discord_DiscordInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Discord_DiscordInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Discord_DiscordInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Discord_DiscordInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Discord_DiscordInfoRequest.Merge(m, src)
}
func (m *Discord_DiscordInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *Discord_DiscordInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Discord_DiscordInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Discord_DiscordInfoRequest proto.InternalMessageInfo

func (*Discord_DiscordInfoRequest) XXX_MessageName() string {
	return "service.Discord.DiscordInfoRequest"
}

type Discord_DiscordInfoResponse struct {
	IsReceived bool `protobuf:"varint,1,opt,name=is_received,json=isReceived,proto3" json:"is_received,omitempty"`
}

func (m *Discord_DiscordInfoResponse) Reset()      { *m = Discord_DiscordInfoResponse{} }
func (*Discord_DiscordInfoResponse) ProtoMessage() {}
func (*Discord_DiscordInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a2d93676bfa42e, []int{0, 1}
}
func (m *Discord_DiscordInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Discord_DiscordInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Discord_DiscordInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Discord_DiscordInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Discord_DiscordInfoResponse.Merge(m, src)
}
func (m *Discord_DiscordInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *Discord_DiscordInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Discord_DiscordInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Discord_DiscordInfoResponse proto.InternalMessageInfo

func (m *Discord_DiscordInfoResponse) GetIsReceived() bool {
	if m != nil {
		return m.IsReceived
	}
	return false
}

func (*Discord_DiscordInfoResponse) XXX_MessageName() string {
	return "service.Discord.DiscordInfoResponse"
}

type Discord_DiscordReceiveRewardRequest struct {
}

func (m *Discord_DiscordReceiveRewardRequest) Reset()      { *m = Discord_DiscordReceiveRewardRequest{} }
func (*Discord_DiscordReceiveRewardRequest) ProtoMessage() {}
func (*Discord_DiscordReceiveRewardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a2d93676bfa42e, []int{0, 2}
}
func (m *Discord_DiscordReceiveRewardRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Discord_DiscordReceiveRewardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Discord_DiscordReceiveRewardRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Discord_DiscordReceiveRewardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Discord_DiscordReceiveRewardRequest.Merge(m, src)
}
func (m *Discord_DiscordReceiveRewardRequest) XXX_Size() int {
	return m.Size()
}
func (m *Discord_DiscordReceiveRewardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Discord_DiscordReceiveRewardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Discord_DiscordReceiveRewardRequest proto.InternalMessageInfo

func (*Discord_DiscordReceiveRewardRequest) XXX_MessageName() string {
	return "service.Discord.DiscordReceiveRewardRequest"
}

type Discord_DiscordReceiveRewardResponse struct {
	Rewards []*models.Item `protobuf:"bytes,1,rep,name=rewards,proto3" json:"rewards,omitempty"`
}

func (m *Discord_DiscordReceiveRewardResponse) Reset()      { *m = Discord_DiscordReceiveRewardResponse{} }
func (*Discord_DiscordReceiveRewardResponse) ProtoMessage() {}
func (*Discord_DiscordReceiveRewardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a2d93676bfa42e, []int{0, 3}
}
func (m *Discord_DiscordReceiveRewardResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Discord_DiscordReceiveRewardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Discord_DiscordReceiveRewardResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Discord_DiscordReceiveRewardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Discord_DiscordReceiveRewardResponse.Merge(m, src)
}
func (m *Discord_DiscordReceiveRewardResponse) XXX_Size() int {
	return m.Size()
}
func (m *Discord_DiscordReceiveRewardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Discord_DiscordReceiveRewardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Discord_DiscordReceiveRewardResponse proto.InternalMessageInfo

func (m *Discord_DiscordReceiveRewardResponse) GetRewards() []*models.Item {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func (*Discord_DiscordReceiveRewardResponse) XXX_MessageName() string {
	return "service.Discord.DiscordReceiveRewardResponse"
}
func init() {
	proto.RegisterEnum("service.DiscordErrorCode", DiscordErrorCode_name, DiscordErrorCode_value)
	proto.RegisterType((*Discord)(nil), "service.Discord")
	proto.RegisterType((*Discord_DiscordInfoRequest)(nil), "service.Discord.DiscordInfoRequest")
	proto.RegisterType((*Discord_DiscordInfoResponse)(nil), "service.Discord.DiscordInfoResponse")
	proto.RegisterType((*Discord_DiscordReceiveRewardRequest)(nil), "service.Discord.DiscordReceiveRewardRequest")
	proto.RegisterType((*Discord_DiscordReceiveRewardResponse)(nil), "service.Discord.DiscordReceiveRewardResponse")
}

func init() { proto.RegisterFile("proto/service/discord.proto", fileDescriptor_b9a2d93676bfa42e) }

var fileDescriptor_b9a2d93676bfa42e = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x86, 0x33, 0x2e, 0x58, 0x9d, 0x2e, 0x1a, 0x47, 0xd1, 0x98, 0xae, 0x63, 0xf1, 0xa0, 0x55,
	0xb0, 0x01, 0x05, 0x7f, 0x80, 0xbb, 0x2b, 0xec, 0x65, 0x0f, 0xf1, 0xe6, 0x65, 0xc9, 0x26, 0xd3,
	0x3a, 0x90, 0x64, 0xea, 0xcc, 0x58, 0xe9, 0x0f, 0x50, 0xb0, 0x50, 0xf4, 0xa0, 0x82, 0x20, 0x88,
	0x9e, 0x3c, 0xb5, 0x52, 0x29, 0x82, 0x45, 0xcf, 0x3d, 0xf6, 0xd8, 0xa3, 0x26, 0x34, 0x6d, 0x6f,
	0xfe, 0x04, 0x69, 0x33, 0xa1, 0xa4, 0x78, 0xcb, 0xf7, 0x3c, 0x6f, 0x5e, 0x66, 0x86, 0x0f, 0x96,
	0x1a, 0x9c, 0x49, 0x66, 0x09, 0xc2, 0x9b, 0xd4, 0x25, 0x96, 0x47, 0x85, 0xcb, 0xb8, 0x57, 0x5d,
	0x51, 0x54, 0x50, 0xd8, 0xdc, 0x49, 0x53, 0x01, 0xf3, 0x88, 0x2f, 0xac, 0x63, 0x26, 0xc4, 0xd1,
	0x63, 0xc7, 0xf7, 0xd3, 0x98, 0x79, 0x31, 0x6f, 0x9d, 0xba, 0xe2, 0x97, 0x36, 0xb8, 0x20, 0x4a,
	0x5c, 0xde, 0x10, 0x52, 0xfa, 0x99, 0x2a, 0xe5, 0x94, 0xe3, 0x4a, 0xda, 0xa4, 0xb2, 0x95, 0xca,
	0x6b, 0xbf, 0x00, 0x2c, 0xec, 0xa5, 0x27, 0x34, 0x2f, 0x40, 0xa4, 0x3e, 0x0f, 0xc2, 0x1a, 0xb3,
	0xc9, 0x93, 0xa7, 0x44, 0x48, 0xf3, 0x1e, 0x3c, 0x9f, 0xa3, 0xa2, 0xc1, 0x42, 0x41, 0xd0, 0x55,
	0x58, 0xa4, 0xe2, 0x88, 0x13, 0x97, 0xd0, 0x26, 0xf1, 0x0c, 0x50, 0x06, 0x95, 0x53, 0x36, 0xa4,
	0xc2, 0x56, 0xc4, 0xbc, 0x02, 0x4b, 0xea, 0x3f, 0x85, 0x6c, 0xf2, 0xcc, 0x59, 0x0e, 0x69, 0xed,
	0x03, 0xb8, 0xf3, 0x7f, 0xad, 0xfa, 0xaf, 0xc3, 0x02, 0x5f, 0x11, 0x61, 0x80, 0xf2, 0x56, 0xa5,
	0x78, 0x67, 0xbb, 0x9a, 0xde, 0xa0, 0x7a, 0x20, 0x49, 0x60, 0x67, 0xf2, 0xd6, 0x0f, 0x00, 0x75,
	0x55, 0xb4, 0xcf, 0x39, 0xe3, 0xbb, 0xcc, 0x23, 0xa8, 0x02, 0xf5, 0x7d, 0xce, 0x15, 0x7e, 0xd8,
	0x12, 0x92, 0x04, 0xba, 0x66, 0xa2, 0xf6, 0xc0, 0x38, 0x83, 0xb6, 0x67, 0xdf, 0x86, 0xb3, 0xe1,
	0xc7, 0xc5, 0x9b, 0xf7, 0xf3, 0x5e, 0x0f, 0xdd, 0x84, 0xe7, 0xd6, 0xc9, 0x43, 0x26, 0xf7, 0x1c,
	0xe9, 0xe8, 0x20, 0x8b, 0x26, 0xfd, 0x4f, 0xc9, 0x87, 0x97, 0xc9, 0xdb, 0xaf, 0x49, 0xa7, 0x9b,
	0x2f, 0xdd, 0x65, 0x61, 0x8d, 0xd6, 0xf5, 0x13, 0x59, 0x72, 0xf1, 0xbc, 0x33, 0xfb, 0xd9, 0x55,
	0xa5, 0x37, 0x20, 0x5a, 0x27, 0xb3, 0x07, 0xd1, 0xb7, 0xcc, 0xb3, 0xed, 0x81, 0x51, 0x44, 0xa7,
	0xa7, 0xdf, 0xfb, 0x8b, 0xcf, 0x2f, 0xa6, 0xaf, 0xde, 0xdd, 0x3f, 0x9c, 0xfc, 0xc1, 0xda, 0x97,
	0x08, 0x83, 0x51, 0x84, 0xc1, 0x38, 0xc2, 0xe0, 0x77, 0x84, 0xc1, 0x3c, 0xc2, 0xda, 0xdf, 0x08,
	0x83, 0xd7, 0x31, 0xd6, 0x46, 0x31, 0x06, 0xe3, 0x18, 0x6b, 0x93, 0x18, 0x6b, 0x8f, 0xca, 0x2e,
	0xa3, 0xe1, 0xed, 0xe5, 0x02, 0x11, 0x6e, 0xb9, 0x2c, 0x08, 0x58, 0x68, 0xe5, 0x76, 0xed, 0xf8,
	0xe4, 0x6a, 0xbc, 0xfb, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xbc, 0xbb, 0xc7, 0x83, 0x02, 0x00,
	0x00,
}

func init() {
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolDiscord.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolDiscord_DiscordInfoRequest.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolDiscord_DiscordInfoResponse.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolDiscord_DiscordReceiveRewardRequest.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolDiscord_DiscordReceiveRewardResponse.Get().(proto.Message)
	})
}

var poolDiscord = &sync.Pool{New: func() interface{} { return &Discord{} }}

func (m *Discord) ReleasePool() { m.Reset(); poolDiscord.Put(m); m = nil }

var poolDiscord_DiscordInfoRequest = &sync.Pool{New: func() interface{} { return &Discord_DiscordInfoRequest{} }}

func (m *Discord_DiscordInfoRequest) ReleasePool() {
	m.Reset()
	poolDiscord_DiscordInfoRequest.Put(m)
	m = nil
}

var poolDiscord_DiscordInfoResponse = &sync.Pool{New: func() interface{} { return &Discord_DiscordInfoResponse{} }}

func (m *Discord_DiscordInfoResponse) ReleasePool() {
	m.Reset()
	poolDiscord_DiscordInfoResponse.Put(m)
	m = nil
}

var poolDiscord_DiscordReceiveRewardRequest = &sync.Pool{New: func() interface{} { return &Discord_DiscordReceiveRewardRequest{} }}

func (m *Discord_DiscordReceiveRewardRequest) ReleasePool() {
	m.Reset()
	poolDiscord_DiscordReceiveRewardRequest.Put(m)
	m = nil
}

var poolDiscord_DiscordReceiveRewardResponse = &sync.Pool{New: func() interface{} { return &Discord_DiscordReceiveRewardResponse{} }}

func (m *Discord_DiscordReceiveRewardResponse) ReleasePool() {
	m.Reset()
	poolDiscord_DiscordReceiveRewardResponse.Put(m)
	m = nil
}
func (x DiscordErrorCode) String() string {
	s, ok := DiscordErrorCode_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Discord) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Discord)
	if !ok {
		that2, ok := that.(Discord)
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
func (this *Discord_DiscordInfoRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Discord_DiscordInfoRequest)
	if !ok {
		that2, ok := that.(Discord_DiscordInfoRequest)
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
func (this *Discord_DiscordInfoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Discord_DiscordInfoResponse)
	if !ok {
		that2, ok := that.(Discord_DiscordInfoResponse)
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
	if this.IsReceived != that1.IsReceived {
		return false
	}
	return true
}
func (this *Discord_DiscordReceiveRewardRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Discord_DiscordReceiveRewardRequest)
	if !ok {
		that2, ok := that.(Discord_DiscordReceiveRewardRequest)
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
func (this *Discord_DiscordReceiveRewardResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Discord_DiscordReceiveRewardResponse)
	if !ok {
		that2, ok := that.(Discord_DiscordReceiveRewardResponse)
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
	if len(this.Rewards) != len(that1.Rewards) {
		return false
	}
	for i := range this.Rewards {
		if !this.Rewards[i].Equal(that1.Rewards[i]) {
			return false
		}
	}
	return true
}
func (m *Discord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Discord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Discord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Discord_DiscordInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Discord_DiscordInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Discord_DiscordInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Discord_DiscordInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Discord_DiscordInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Discord_DiscordInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsReceived {
		i--
		if m.IsReceived {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Discord_DiscordReceiveRewardRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Discord_DiscordReceiveRewardRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Discord_DiscordReceiveRewardRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Discord_DiscordReceiveRewardResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Discord_DiscordReceiveRewardResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Discord_DiscordReceiveRewardResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDiscord(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDiscord(dAtA []byte, offset int, v uint64) int {
	offset -= sovDiscord(v)
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

func (m *Discord) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	w.RawByte('}')

}

func (m *Discord_DiscordInfoRequest) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	w.RawByte('}')

}

func (m *Discord_DiscordInfoResponse) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.IsReceived {
		w.RawByte('"')
		w.RawString("is_received")
		w.RawByte('"')
		w.RawByte(':')
		w.Bool(m.IsReceived)
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *Discord_DiscordReceiveRewardRequest) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	w.RawByte('}')

}

func (m *Discord_DiscordReceiveRewardResponse) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	w.RawByte('"')
	w.RawString("rewards")
	w.RawByte('"')
	w.RawByte(':')
	if m.Rewards == nil {
		w.RawString("null")
	} else if len(m.Rewards) == 0 {
		w.RawString("[]")
	} else {
		w.RawByte('[')
		for i, v := range m.Rewards {
			v.JsonBytes(w)
			if i != len(m.Rewards)-1 {
				w.RawByte(',')
			}
		}
		w.RawByte(']')
	}
	needWriteComma = true
	_ = needWriteComma
	w.RawByte('}')

}

func (m *Discord) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *Discord) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *Discord) GoString() string {
	return m.String()
}

func (m *Discord_DiscordInfoRequest) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *Discord_DiscordInfoRequest) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *Discord_DiscordInfoRequest) GoString() string {
	return m.String()
}

func (m *Discord_DiscordInfoResponse) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *Discord_DiscordInfoResponse) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *Discord_DiscordInfoResponse) GoString() string {
	return m.String()
}

func (m *Discord_DiscordReceiveRewardRequest) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *Discord_DiscordReceiveRewardRequest) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *Discord_DiscordReceiveRewardRequest) GoString() string {
	return m.String()
}

func (m *Discord_DiscordReceiveRewardResponse) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *Discord_DiscordReceiveRewardResponse) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *Discord_DiscordReceiveRewardResponse) GoString() string {
	return m.String()
}

func (m *Discord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Discord_DiscordInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Discord_DiscordInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsReceived {
		n += 2
	}
	return n
}

func (m *Discord_DiscordReceiveRewardRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Discord_DiscordReceiveRewardResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovDiscord(uint64(l))
		}
	}
	return n
}

func sovDiscord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDiscord(x uint64) (n int) {
	return sovDiscord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Discord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscord
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
			return fmt.Errorf("proto: Discord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Discord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDiscord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDiscord
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
func (m *Discord_DiscordInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscord
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
			return fmt.Errorf("proto: DiscordInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DiscordInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDiscord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDiscord
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
func (m *Discord_DiscordInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscord
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
			return fmt.Errorf("proto: DiscordInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DiscordInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsReceived", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsReceived = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDiscord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDiscord
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
func (m *Discord_DiscordReceiveRewardRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscord
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
			return fmt.Errorf("proto: DiscordReceiveRewardRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DiscordReceiveRewardRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDiscord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDiscord
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
func (m *Discord_DiscordReceiveRewardResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscord
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
			return fmt.Errorf("proto: DiscordReceiveRewardResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DiscordReceiveRewardResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscord
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
				return ErrInvalidLengthDiscord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDiscord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, &models.Item{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiscord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDiscord
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
func skipDiscord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDiscord
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
					return 0, ErrIntOverflowDiscord
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
					return 0, ErrIntOverflowDiscord
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
				return 0, ErrInvalidLengthDiscord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDiscord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDiscord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDiscord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDiscord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDiscord = fmt.Errorf("proto: unexpected end of group")
)
