// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/service/quest_reward.proto

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

type QuestRewardErrorCode int32

const (
	QuestRewardErrorCode_ErrQuestRewardSystem   QuestRewardErrorCode = 0
	QuestRewardErrorCode_ErrQuestRewardNotData  QuestRewardErrorCode = 1
	QuestRewardErrorCode_ErrQuestRewardConfig   QuestRewardErrorCode = 2
	QuestRewardErrorCode_ErrQuestRewardReceived QuestRewardErrorCode = 3
)

var QuestRewardErrorCode_name = map[int32]string{
	0: "ErrQuestRewardSystem",
	1: "ErrQuestRewardNotData",
	2: "ErrQuestRewardConfig",
	3: "ErrQuestRewardReceived",
}

var QuestRewardErrorCode_value = map[string]int32{
	"ErrQuestRewardSystem":   0,
	"ErrQuestRewardNotData":  1,
	"ErrQuestRewardConfig":   2,
	"ErrQuestRewardReceived": 3,
}

func (QuestRewardErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_501ed2d2d0b38c6c, []int{0}
}

type QuestReward struct {
}

func (m *QuestReward) Reset()      { *m = QuestReward{} }
func (*QuestReward) ProtoMessage() {}
func (*QuestReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_501ed2d2d0b38c6c, []int{0}
}
func (m *QuestReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuestReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuestReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuestReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestReward.Merge(m, src)
}
func (m *QuestReward) XXX_Size() int {
	return m.Size()
}
func (m *QuestReward) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestReward.DiscardUnknown(m)
}

var xxx_messageInfo_QuestReward proto.InternalMessageInfo

func (*QuestReward) XXX_MessageName() string {
	return "service.QuestReward"
}

// 查询情况
type QuestReward_QuestRewardInfoRequest struct {
}

func (m *QuestReward_QuestRewardInfoRequest) Reset()      { *m = QuestReward_QuestRewardInfoRequest{} }
func (*QuestReward_QuestRewardInfoRequest) ProtoMessage() {}
func (*QuestReward_QuestRewardInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_501ed2d2d0b38c6c, []int{0, 0}
}
func (m *QuestReward_QuestRewardInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuestReward_QuestRewardInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuestReward_QuestRewardInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuestReward_QuestRewardInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestReward_QuestRewardInfoRequest.Merge(m, src)
}
func (m *QuestReward_QuestRewardInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuestReward_QuestRewardInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestReward_QuestRewardInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuestReward_QuestRewardInfoRequest proto.InternalMessageInfo

func (*QuestReward_QuestRewardInfoRequest) XXX_MessageName() string {
	return "service.QuestReward.QuestRewardInfoRequest"
}

type QuestReward_QuestRewardInfoResponse struct {
	IsReceived bool `protobuf:"varint,1,opt,name=is_received,json=isReceived,proto3" json:"is_received,omitempty"`
}

func (m *QuestReward_QuestRewardInfoResponse) Reset()      { *m = QuestReward_QuestRewardInfoResponse{} }
func (*QuestReward_QuestRewardInfoResponse) ProtoMessage() {}
func (*QuestReward_QuestRewardInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_501ed2d2d0b38c6c, []int{0, 1}
}
func (m *QuestReward_QuestRewardInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuestReward_QuestRewardInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuestReward_QuestRewardInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuestReward_QuestRewardInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestReward_QuestRewardInfoResponse.Merge(m, src)
}
func (m *QuestReward_QuestRewardInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuestReward_QuestRewardInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestReward_QuestRewardInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuestReward_QuestRewardInfoResponse proto.InternalMessageInfo

func (m *QuestReward_QuestRewardInfoResponse) GetIsReceived() bool {
	if m != nil {
		return m.IsReceived
	}
	return false
}

func (*QuestReward_QuestRewardInfoResponse) XXX_MessageName() string {
	return "service.QuestReward.QuestRewardInfoResponse"
}

type QuestReward_QuestRewardReceiveRewardRequest struct {
}

func (m *QuestReward_QuestRewardReceiveRewardRequest) Reset() {
	*m = QuestReward_QuestRewardReceiveRewardRequest{}
}
func (*QuestReward_QuestRewardReceiveRewardRequest) ProtoMessage() {}
func (*QuestReward_QuestRewardReceiveRewardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_501ed2d2d0b38c6c, []int{0, 2}
}
func (m *QuestReward_QuestRewardReceiveRewardRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuestReward_QuestRewardReceiveRewardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuestReward_QuestRewardReceiveRewardRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuestReward_QuestRewardReceiveRewardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestReward_QuestRewardReceiveRewardRequest.Merge(m, src)
}
func (m *QuestReward_QuestRewardReceiveRewardRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuestReward_QuestRewardReceiveRewardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestReward_QuestRewardReceiveRewardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuestReward_QuestRewardReceiveRewardRequest proto.InternalMessageInfo

func (*QuestReward_QuestRewardReceiveRewardRequest) XXX_MessageName() string {
	return "service.QuestReward.QuestRewardReceiveRewardRequest"
}

type QuestReward_QuestRewardReceiveRewardResponse struct {
	Rewards []*models.Item `protobuf:"bytes,1,rep,name=rewards,proto3" json:"rewards,omitempty"`
}

func (m *QuestReward_QuestRewardReceiveRewardResponse) Reset() {
	*m = QuestReward_QuestRewardReceiveRewardResponse{}
}
func (*QuestReward_QuestRewardReceiveRewardResponse) ProtoMessage() {}
func (*QuestReward_QuestRewardReceiveRewardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_501ed2d2d0b38c6c, []int{0, 3}
}
func (m *QuestReward_QuestRewardReceiveRewardResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuestReward_QuestRewardReceiveRewardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuestReward_QuestRewardReceiveRewardResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuestReward_QuestRewardReceiveRewardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuestReward_QuestRewardReceiveRewardResponse.Merge(m, src)
}
func (m *QuestReward_QuestRewardReceiveRewardResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuestReward_QuestRewardReceiveRewardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuestReward_QuestRewardReceiveRewardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuestReward_QuestRewardReceiveRewardResponse proto.InternalMessageInfo

func (m *QuestReward_QuestRewardReceiveRewardResponse) GetRewards() []*models.Item {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func (*QuestReward_QuestRewardReceiveRewardResponse) XXX_MessageName() string {
	return "service.QuestReward.QuestRewardReceiveRewardResponse"
}
func init() {
	proto.RegisterEnum("service.QuestRewardErrorCode", QuestRewardErrorCode_name, QuestRewardErrorCode_value)
	proto.RegisterType((*QuestReward)(nil), "service.QuestReward")
	proto.RegisterType((*QuestReward_QuestRewardInfoRequest)(nil), "service.QuestReward.QuestRewardInfoRequest")
	proto.RegisterType((*QuestReward_QuestRewardInfoResponse)(nil), "service.QuestReward.QuestRewardInfoResponse")
	proto.RegisterType((*QuestReward_QuestRewardReceiveRewardRequest)(nil), "service.QuestReward.QuestRewardReceiveRewardRequest")
	proto.RegisterType((*QuestReward_QuestRewardReceiveRewardResponse)(nil), "service.QuestReward.QuestRewardReceiveRewardResponse")
}

func init() { proto.RegisterFile("proto/service/quest_reward.proto", fileDescriptor_501ed2d2d0b38c6c) }

var fileDescriptor_501ed2d2d0b38c6c = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0x2c, 0x4d, 0x2d, 0x2e, 0x89,
	0x2f, 0x4a, 0x2d, 0x4f, 0x2c, 0x4a, 0xd1, 0x03, 0x4b, 0x09, 0xb1, 0x43, 0xe5, 0xa4, 0x64, 0x20,
	0x4a, 0x73, 0xf3, 0x53, 0x52, 0x73, 0x8a, 0xf5, 0x93, 0xf2, 0x8b, 0x8b, 0xe3, 0x33, 0x12, 0x73,
	0x72, 0x20, 0xca, 0xa4, 0xc4, 0x50, 0x65, 0x13, 0xd3, 0xa1, 0xe2, 0xe2, 0x68, 0xe2, 0xc5, 0xa9,
	0x50, 0x09, 0x49, 0x34, 0x89, 0x92, 0x92, 0x1c, 0x98, 0x94, 0x34, 0x8a, 0x54, 0x62, 0x72, 0x49,
	0x66, 0x59, 0x66, 0x49, 0x25, 0x44, 0x52, 0xe9, 0x12, 0x23, 0x17, 0x77, 0x20, 0xc8, 0x99, 0x41,
	0x60, 0x57, 0x4a, 0x49, 0x70, 0x89, 0x21, 0x71, 0x3d, 0xf3, 0xd2, 0xf2, 0x83, 0x52, 0xc1, 0xde,
	0x90, 0xb2, 0xe2, 0x12, 0xc7, 0x90, 0x29, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x92, 0xe7, 0xe2,
	0xce, 0x2c, 0x8e, 0x2f, 0x4a, 0x4d, 0x4e, 0xcd, 0x2c, 0x4b, 0x4d, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x08, 0xe2, 0xca, 0x2c, 0x0e, 0x82, 0x8a, 0x48, 0x29, 0x72, 0xc9, 0x23, 0xe9, 0x85, 0x0a,
	0xc3, 0x38, 0x10, 0xe3, 0xbd, 0xb8, 0x14, 0x70, 0x2b, 0x81, 0xda, 0xa3, 0xc6, 0xc5, 0x0e, 0x09,
	0xcc, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x1e, 0x3d, 0x88, 0xaf, 0xf4, 0x3c, 0x4b,
	0x52, 0x73, 0x83, 0x60, 0x92, 0x5a, 0xe7, 0x18, 0xb9, 0x44, 0x90, 0x0c, 0x73, 0x2d, 0x2a, 0xca,
	0x2f, 0x72, 0xce, 0x4f, 0x49, 0x15, 0xd2, 0xe1, 0x12, 0x71, 0x2d, 0x2a, 0x42, 0x92, 0x0a, 0xae,
	0x2c, 0x2e, 0x49, 0xcd, 0x15, 0x60, 0x90, 0x12, 0xea, 0xda, 0x2a, 0xc1, 0x27, 0xc4, 0xf3, 0x7c,
	0xf3, 0xee, 0xe7, 0xbb, 0xe7, 0xbf, 0x9c, 0x32, 0xf3, 0xc5, 0xfa, 0xf5, 0x42, 0xba, 0x5c, 0xa2,
	0xa8, 0xaa, 0xfd, 0xf2, 0x4b, 0x5c, 0x12, 0x4b, 0x12, 0x05, 0x18, 0x61, 0xca, 0x9f, 0x6d, 0x5a,
	0xf8, 0x6c, 0x4e, 0xe7, 0xb3, 0xa9, 0x1b, 0x9e, 0xf5, 0xae, 0xc3, 0x34, 0xdc, 0x39, 0x3f, 0x2f,
	0x2d, 0x33, 0x5d, 0x80, 0x09, 0xa6, 0xfa, 0x65, 0x6b, 0xef, 0xf3, 0xbd, 0xeb, 0xa0, 0x86, 0x6b,
	0x73, 0x89, 0xa1, 0xaa, 0x86, 0x05, 0x96, 0x00, 0xb3, 0x14, 0x7f, 0xd7, 0x56, 0x09, 0x6e, 0x21,
	0xce, 0xa7, 0xdb, 0x37, 0xbd, 0x5c, 0xd4, 0xf6, 0xb4, 0x7f, 0x9a, 0x93, 0xdf, 0x8d, 0x87, 0x72,
	0x0c, 0x2b, 0x1e, 0xc9, 0x31, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x2f, 0x1e, 0xc9, 0x31, 0x7c, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x89,
	0xc7, 0x72, 0x8c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x90, 0x9c, 0x9f,
	0x99, 0xa7, 0x0b, 0x4a, 0x68, 0xa9, 0x45, 0xfa, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0xfa, 0x28,
	0x09, 0x33, 0x89, 0x0d, 0xcc, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x73, 0xb5, 0x6b,
	0xb0, 0x02, 0x00, 0x00,
}

func init() {
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolQuestReward.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolQuestReward_QuestRewardInfoRequest.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolQuestReward_QuestRewardInfoResponse.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolQuestReward_QuestRewardReceiveRewardRequest.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolQuestReward_QuestRewardReceiveRewardResponse.Get().(proto.Message)
	})
}

var poolQuestReward = &sync.Pool{New: func() interface{} { return &QuestReward{} }}

func (m *QuestReward) ReleasePool() { m.Reset(); poolQuestReward.Put(m); m = nil }

var poolQuestReward_QuestRewardInfoRequest = &sync.Pool{New: func() interface{} { return &QuestReward_QuestRewardInfoRequest{} }}

func (m *QuestReward_QuestRewardInfoRequest) ReleasePool() {
	m.Reset()
	poolQuestReward_QuestRewardInfoRequest.Put(m)
	m = nil
}

var poolQuestReward_QuestRewardInfoResponse = &sync.Pool{New: func() interface{} { return &QuestReward_QuestRewardInfoResponse{} }}

func (m *QuestReward_QuestRewardInfoResponse) ReleasePool() {
	m.Reset()
	poolQuestReward_QuestRewardInfoResponse.Put(m)
	m = nil
}

var poolQuestReward_QuestRewardReceiveRewardRequest = &sync.Pool{New: func() interface{} { return &QuestReward_QuestRewardReceiveRewardRequest{} }}

func (m *QuestReward_QuestRewardReceiveRewardRequest) ReleasePool() {
	m.Reset()
	poolQuestReward_QuestRewardReceiveRewardRequest.Put(m)
	m = nil
}

var poolQuestReward_QuestRewardReceiveRewardResponse = &sync.Pool{New: func() interface{} { return &QuestReward_QuestRewardReceiveRewardResponse{} }}

func (m *QuestReward_QuestRewardReceiveRewardResponse) ReleasePool() {
	m.Reset()
	poolQuestReward_QuestRewardReceiveRewardResponse.Put(m)
	m = nil
}
func (x QuestRewardErrorCode) String() string {
	s, ok := QuestRewardErrorCode_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *QuestReward) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QuestReward)
	if !ok {
		that2, ok := that.(QuestReward)
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
func (this *QuestReward_QuestRewardInfoRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QuestReward_QuestRewardInfoRequest)
	if !ok {
		that2, ok := that.(QuestReward_QuestRewardInfoRequest)
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
func (this *QuestReward_QuestRewardInfoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QuestReward_QuestRewardInfoResponse)
	if !ok {
		that2, ok := that.(QuestReward_QuestRewardInfoResponse)
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
func (this *QuestReward_QuestRewardReceiveRewardRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QuestReward_QuestRewardReceiveRewardRequest)
	if !ok {
		that2, ok := that.(QuestReward_QuestRewardReceiveRewardRequest)
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
func (this *QuestReward_QuestRewardReceiveRewardResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QuestReward_QuestRewardReceiveRewardResponse)
	if !ok {
		that2, ok := that.(QuestReward_QuestRewardReceiveRewardResponse)
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
func (m *QuestReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuestReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QuestReward_QuestRewardInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestReward_QuestRewardInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuestReward_QuestRewardInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QuestReward_QuestRewardInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestReward_QuestRewardInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuestReward_QuestRewardInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QuestReward_QuestRewardReceiveRewardRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestReward_QuestRewardReceiveRewardRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuestReward_QuestRewardReceiveRewardRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QuestReward_QuestRewardReceiveRewardResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuestReward_QuestRewardReceiveRewardResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuestReward_QuestRewardReceiveRewardResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintQuestReward(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuestReward(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuestReward(v)
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

func (m *QuestReward) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	w.RawByte('}')

}

func (m *QuestReward_QuestRewardInfoRequest) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	w.RawByte('}')

}

func (m *QuestReward_QuestRewardInfoResponse) JsonBytes(w *coin_server_common_jwriter.Writer) {
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

func (m *QuestReward_QuestRewardReceiveRewardRequest) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	w.RawByte('}')

}

func (m *QuestReward_QuestRewardReceiveRewardResponse) JsonBytes(w *coin_server_common_jwriter.Writer) {
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

func (m *QuestReward) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *QuestReward) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *QuestReward) GoString() string {
	return m.String()
}

func (m *QuestReward_QuestRewardInfoRequest) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *QuestReward_QuestRewardInfoRequest) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *QuestReward_QuestRewardInfoRequest) GoString() string {
	return m.String()
}

func (m *QuestReward_QuestRewardInfoResponse) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *QuestReward_QuestRewardInfoResponse) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *QuestReward_QuestRewardInfoResponse) GoString() string {
	return m.String()
}

func (m *QuestReward_QuestRewardReceiveRewardRequest) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *QuestReward_QuestRewardReceiveRewardRequest) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *QuestReward_QuestRewardReceiveRewardRequest) GoString() string {
	return m.String()
}

func (m *QuestReward_QuestRewardReceiveRewardResponse) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *QuestReward_QuestRewardReceiveRewardResponse) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *QuestReward_QuestRewardReceiveRewardResponse) GoString() string {
	return m.String()
}

func (m *QuestReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QuestReward_QuestRewardInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QuestReward_QuestRewardInfoResponse) Size() (n int) {
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

func (m *QuestReward_QuestRewardReceiveRewardRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QuestReward_QuestRewardReceiveRewardResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovQuestReward(uint64(l))
		}
	}
	return n
}

func sovQuestReward(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuestReward(x uint64) (n int) {
	return sovQuestReward(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuestReward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuestReward
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
			return fmt.Errorf("proto: QuestReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuestReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuestReward
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
func (m *QuestReward_QuestRewardInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuestReward
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
			return fmt.Errorf("proto: QuestRewardInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestRewardInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuestReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuestReward
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
func (m *QuestReward_QuestRewardInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuestReward
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
			return fmt.Errorf("proto: QuestRewardInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestRewardInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsReceived", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuestReward
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
			skippy, err := skipQuestReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuestReward
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
func (m *QuestReward_QuestRewardReceiveRewardRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuestReward
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
			return fmt.Errorf("proto: QuestRewardReceiveRewardRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestRewardReceiveRewardRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuestReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuestReward
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
func (m *QuestReward_QuestRewardReceiveRewardResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuestReward
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
			return fmt.Errorf("proto: QuestRewardReceiveRewardResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuestRewardReceiveRewardResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuestReward
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
				return ErrInvalidLengthQuestReward
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuestReward
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
			skippy, err := skipQuestReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuestReward
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
func skipQuestReward(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuestReward
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
					return 0, ErrIntOverflowQuestReward
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
					return 0, ErrIntOverflowQuestReward
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
				return 0, ErrInvalidLengthQuestReward
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuestReward
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuestReward
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuestReward        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuestReward          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuestReward = fmt.Errorf("proto: unexpected end of group")
)