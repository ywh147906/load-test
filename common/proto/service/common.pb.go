// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/service/common.proto

package service

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

type CommonErrorCode int32

const (
	CommonErrorCode_ErrInvalidRequestParam CommonErrorCode = 0
	CommonErrorCode_ErrTaskNotCompleted    CommonErrorCode = 10
	CommonErrorCode_ErrTaskInvalidType     CommonErrorCode = 11
	CommonErrorCode_ErrTaskNotExist        CommonErrorCode = 12
	CommonErrorCode_ErrTaskAlreadyFinish   CommonErrorCode = 13
	CommonErrorCode_ErrTaskCfgNotExist     CommonErrorCode = 14
	CommonErrorCode_ErrTaskChoiceNotExist  CommonErrorCode = 15
	CommonErrorCode_ErrSensitive           CommonErrorCode = 16
)

var CommonErrorCode_name = map[int32]string{
	0:  "ErrInvalidRequestParam",
	10: "ErrTaskNotCompleted",
	11: "ErrTaskInvalidType",
	12: "ErrTaskNotExist",
	13: "ErrTaskAlreadyFinish",
	14: "ErrTaskCfgNotExist",
	15: "ErrTaskChoiceNotExist",
	16: "ErrSensitive",
}

var CommonErrorCode_value = map[string]int32{
	"ErrInvalidRequestParam": 0,
	"ErrTaskNotCompleted":    10,
	"ErrTaskInvalidType":     11,
	"ErrTaskNotExist":        12,
	"ErrTaskAlreadyFinish":   13,
	"ErrTaskCfgNotExist":     14,
	"ErrTaskChoiceNotExist":  15,
	"ErrSensitive":           16,
}

func (CommonErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ebc88e943de6af78, []int{0}
}

type Empty struct {
}

func (m *Empty) Reset()      { *m = Empty{} }
func (*Empty) ProtoMessage() {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebc88e943de6af78, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return m.Size()
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func (*Empty) XXX_MessageName() string {
	return "service.Empty"
}

type Common struct {
}

func (m *Common) Reset()      { *m = Common{} }
func (*Common) ProtoMessage() {}
func (*Common) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebc88e943de6af78, []int{1}
}
func (m *Common) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Common) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Common.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Common) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Common.Merge(m, src)
}
func (m *Common) XXX_Size() int {
	return m.Size()
}
func (m *Common) XXX_DiscardUnknown() {
	xxx_messageInfo_Common.DiscardUnknown(m)
}

var xxx_messageInfo_Common proto.InternalMessageInfo

func (*Common) XXX_MessageName() string {
	return "service.Common"
}

type Common_CheckSensitiveTextRequest struct {
	Txt string `protobuf:"bytes,1,opt,name=txt,proto3" json:"txt,omitempty"`
}

func (m *Common_CheckSensitiveTextRequest) Reset()      { *m = Common_CheckSensitiveTextRequest{} }
func (*Common_CheckSensitiveTextRequest) ProtoMessage() {}
func (*Common_CheckSensitiveTextRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebc88e943de6af78, []int{1, 0}
}
func (m *Common_CheckSensitiveTextRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Common_CheckSensitiveTextRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Common_CheckSensitiveTextRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Common_CheckSensitiveTextRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Common_CheckSensitiveTextRequest.Merge(m, src)
}
func (m *Common_CheckSensitiveTextRequest) XXX_Size() int {
	return m.Size()
}
func (m *Common_CheckSensitiveTextRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Common_CheckSensitiveTextRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Common_CheckSensitiveTextRequest proto.InternalMessageInfo

func (m *Common_CheckSensitiveTextRequest) GetTxt() string {
	if m != nil {
		return m.Txt
	}
	return ""
}

func (*Common_CheckSensitiveTextRequest) XXX_MessageName() string {
	return "service.Common.CheckSensitiveTextRequest"
}

type Common_CheckSensitiveTextResponse struct {
	IsPass bool `protobuf:"varint,1,opt,name=is_pass,json=isPass,proto3" json:"is_pass,omitempty"`
}

func (m *Common_CheckSensitiveTextResponse) Reset()      { *m = Common_CheckSensitiveTextResponse{} }
func (*Common_CheckSensitiveTextResponse) ProtoMessage() {}
func (*Common_CheckSensitiveTextResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebc88e943de6af78, []int{1, 1}
}
func (m *Common_CheckSensitiveTextResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Common_CheckSensitiveTextResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Common_CheckSensitiveTextResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Common_CheckSensitiveTextResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Common_CheckSensitiveTextResponse.Merge(m, src)
}
func (m *Common_CheckSensitiveTextResponse) XXX_Size() int {
	return m.Size()
}
func (m *Common_CheckSensitiveTextResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Common_CheckSensitiveTextResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Common_CheckSensitiveTextResponse proto.InternalMessageInfo

func (m *Common_CheckSensitiveTextResponse) GetIsPass() bool {
	if m != nil {
		return m.IsPass
	}
	return false
}

func (*Common_CheckSensitiveTextResponse) XXX_MessageName() string {
	return "service.Common.CheckSensitiveTextResponse"
}
func init() {
	proto.RegisterEnum("service.CommonErrorCode", CommonErrorCode_name, CommonErrorCode_value)
	proto.RegisterType((*Empty)(nil), "service.Empty")
	proto.RegisterType((*Common)(nil), "service.Common")
	proto.RegisterType((*Common_CheckSensitiveTextRequest)(nil), "service.Common.CheckSensitiveTextRequest")
	proto.RegisterType((*Common_CheckSensitiveTextResponse)(nil), "service.Common.CheckSensitiveTextResponse")
}

func init() { proto.RegisterFile("proto/service/common.proto", fileDescriptor_ebc88e943de6af78) }

var fileDescriptor_ebc88e943de6af78 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6b, 0x13, 0x4d,
	0x18, 0x87, 0x77, 0x29, 0x5f, 0xf3, 0x39, 0x56, 0xb3, 0x8e, 0x8d, 0x59, 0xb7, 0x38, 0xd4, 0x9e,
	0x24, 0xd0, 0xe6, 0x50, 0xc4, 0x8b, 0x17, 0x0d, 0x2b, 0x78, 0x29, 0x25, 0xe6, 0xe4, 0x45, 0xb6,
	0xc9, 0x68, 0x86, 0x66, 0x77, 0xd6, 0x99, 0x35, 0x24, 0x37, 0x29, 0xa9, 0xc6, 0xd8, 0xd0, 0x1c,
	0x0c, 0x4a, 0x28, 0x6a, 0xa1, 0x87, 0x62, 0x69, 0x63, 0x84, 0x50, 0x25, 0x78, 0xef, 0xb1, 0xc7,
	0x1e, 0x75, 0xd3, 0xd9, 0x24, 0x37, 0xff, 0x04, 0x69, 0xb2, 0x4d, 0x0d, 0x7a, 0x9b, 0x79, 0xf9,
	0x3d, 0xf3, 0xbc, 0x33, 0xf3, 0x02, 0xcd, 0x66, 0xd4, 0xa1, 0x51, 0x8e, 0x59, 0x96, 0x24, 0x71,
	0x34, 0x49, 0x4d, 0x93, 0x5a, 0x73, 0xfd, 0x22, 0x0c, 0xf8, 0x55, 0x2d, 0x3c, 0x08, 0x99, 0x34,
	0x85, 0x33, 0x3c, 0xba, 0x64, 0x70, 0x3c, 0x48, 0xcc, 0x04, 0xc0, 0x7f, 0xba, 0x69, 0x3b, 0xf9,
	0x19, 0x0b, 0x8c, 0xc7, 0xfa, 0xa8, 0x36, 0x0b, 0xae, 0xc6, 0xd2, 0x38, 0xb9, 0xfc, 0x00, 0x5b,
	0x9c, 0x38, 0x24, 0x8b, 0x13, 0x38, 0xe7, 0xc4, 0xf1, 0xd3, 0x67, 0x98, 0x3b, 0x50, 0x01, 0x63,
	0x4e, 0xce, 0x51, 0xe5, 0x69, 0xf9, 0xc6, 0xb9, 0xf8, 0xc9, 0x52, 0xbb, 0x09, 0xb4, 0x7f, 0xc5,
	0xb9, 0x4d, 0x2d, 0x8e, 0x61, 0x18, 0x04, 0x08, 0x7f, 0x64, 0x1b, 0x9c, 0xf7, 0x99, 0xff, 0xe3,
	0xe3, 0x84, 0x2f, 0x1a, 0x9c, 0x47, 0x76, 0xc6, 0x40, 0x70, 0x20, 0xd4, 0x19, 0xa3, 0x2c, 0x46,
	0x53, 0x18, 0xde, 0x02, 0x57, 0x74, 0xc6, 0xee, 0x5b, 0x59, 0x23, 0x43, 0x52, 0xbe, 0x71, 0xd1,
	0x60, 0x86, 0xa9, 0x48, 0xda, 0x54, 0xa9, 0xa9, 0x86, 0x61, 0xc8, 0x7b, 0xfb, 0xd1, 0xab, 0x16,
	0x3b, 0xb5, 0x42, 0xb7, 0xbe, 0xef, 0x35, 0x56, 0xc4, 0xfa, 0x8a, 0x57, 0xfd, 0x0c, 0xe7, 0xc0,
	0x65, 0x9d, 0xb1, 0x84, 0xc1, 0x97, 0x17, 0xa8, 0x13, 0xa3, 0xa6, 0x9d, 0xc1, 0x0e, 0x4e, 0x29,
	0x40, 0x0b, 0x95, 0x9a, 0xea, 0x25, 0x18, 0x3c, 0x6e, 0xb5, 0x44, 0x69, 0xd3, 0xdb, 0xf8, 0x24,
	0xf6, 0xd6, 0xbc, 0x62, 0x05, 0xce, 0x03, 0xe8, 0xe7, 0x7d, 0x59, 0x22, 0x6f, 0x63, 0xe5, 0xfc,
	0x9f, 0x92, 0xdd, 0xad, 0x4e, 0xa3, 0x25, 0xde, 0xbf, 0xee, 0xd4, 0x0a, 0x03, 0x1a, 0x46, 0x40,
	0xf0, 0x4c, 0xa2, 0xe7, 0x08, 0x77, 0x94, 0x89, 0x51, 0xc1, 0xf1, 0xd7, 0xb2, 0xd8, 0x7d, 0x27,
	0x36, 0xb6, 0x61, 0x14, 0x4c, 0xfa, 0xd9, 0x3b, 0x19, 0x86, 0x8d, 0x54, 0xfe, 0x1e, 0xb1, 0x08,
	0x4f, 0x2b, 0x17, 0x46, 0x01, 0xb1, 0xff, 0xe5, 0xaf, 0x8e, 0x62, 0x8f, 0x9f, 0x0c, 0xcf, 0xbf,
	0x78, 0xda, 0xd1, 0x20, 0xde, 0x5b, 0x2d, 0x77, 0xbe, 0xef, 0x9d, 0x59, 0x6e, 0x83, 0xd0, 0x29,
	0x94, 0xa6, 0x24, 0x89, 0x87, 0x5c, 0x50, 0xbb, 0x5e, 0x6a, 0xaa, 0xd7, 0xe0, 0x94, 0xd8, 0xaa,
	0x89, 0xe2, 0x0b, 0xef, 0x4d, 0xdd, 0x7f, 0x81, 0xfe, 0xc5, 0x7a, 0xcf, 0x5f, 0xf5, 0x36, 0xbf,
	0xc1, 0x08, 0x98, 0xd0, 0x19, 0x1b, 0x7e, 0x9b, 0xa2, 0x68, 0x6a, 0xa9, 0xa9, 0x4e, 0x42, 0x28,
	0xd6, 0x56, 0x45, 0x65, 0xc7, 0xab, 0xae, 0x7b, 0x85, 0x0f, 0xdd, 0x7a, 0xd9, 0x6b, 0xbc, 0xbc,
	0xbb, 0x70, 0xf4, 0x13, 0x49, 0xdb, 0x2e, 0x92, 0x0f, 0x5c, 0x24, 0x1f, 0xba, 0x48, 0xfe, 0xe1,
	0x22, 0xb9, 0xeb, 0x22, 0xe9, 0x97, 0x8b, 0xe4, 0x4a, 0x1b, 0x49, 0x07, 0x6d, 0x24, 0x1f, 0xb6,
	0x91, 0x74, 0xd4, 0x46, 0xd2, 0xc3, 0xe9, 0x24, 0x25, 0xd6, 0xec, 0xc9, 0xec, 0x61, 0xe6, 0x0f,
	0x64, 0x74, 0x64, 0x4a, 0x97, 0xc6, 0xfb, 0xdb, 0xf9, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x99,
	0x58, 0xec, 0x00, 0xbd, 0x02, 0x00, 0x00,
}

func init() {
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolEmpty.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolCommon.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolCommon_CheckSensitiveTextRequest.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolCommon_CheckSensitiveTextResponse.Get().(proto.Message)
	})
}

var poolEmpty = &sync.Pool{New: func() interface{} { return &Empty{} }}

func (m *Empty) ReleasePool() { m.Reset(); poolEmpty.Put(m); m = nil }

var poolCommon = &sync.Pool{New: func() interface{} { return &Common{} }}

func (m *Common) ReleasePool() { m.Reset(); poolCommon.Put(m); m = nil }

var poolCommon_CheckSensitiveTextRequest = &sync.Pool{New: func() interface{} { return &Common_CheckSensitiveTextRequest{} }}

func (m *Common_CheckSensitiveTextRequest) ReleasePool() {
	m.Reset()
	poolCommon_CheckSensitiveTextRequest.Put(m)
	m = nil
}

var poolCommon_CheckSensitiveTextResponse = &sync.Pool{New: func() interface{} { return &Common_CheckSensitiveTextResponse{} }}

func (m *Common_CheckSensitiveTextResponse) ReleasePool() {
	m.Reset()
	poolCommon_CheckSensitiveTextResponse.Put(m)
	m = nil
}
func (x CommonErrorCode) String() string {
	s, ok := CommonErrorCode_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Empty) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Empty)
	if !ok {
		that2, ok := that.(Empty)
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
func (this *Common) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Common)
	if !ok {
		that2, ok := that.(Common)
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
func (this *Common_CheckSensitiveTextRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Common_CheckSensitiveTextRequest)
	if !ok {
		that2, ok := that.(Common_CheckSensitiveTextRequest)
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
	if this.Txt != that1.Txt {
		return false
	}
	return true
}
func (this *Common_CheckSensitiveTextResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Common_CheckSensitiveTextResponse)
	if !ok {
		that2, ok := that.(Common_CheckSensitiveTextResponse)
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
	if this.IsPass != that1.IsPass {
		return false
	}
	return true
}
func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Empty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Common) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Common) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Common) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Common_CheckSensitiveTextRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Common_CheckSensitiveTextRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Common_CheckSensitiveTextRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Txt) > 0 {
		i -= len(m.Txt)
		copy(dAtA[i:], m.Txt)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Txt)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Common_CheckSensitiveTextResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Common_CheckSensitiveTextResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Common_CheckSensitiveTextResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsPass {
		i--
		if m.IsPass {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
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

func (m *Empty) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	w.RawByte('}')

}

func (m *Common) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	w.RawByte('}')

}

func (m *Common_CheckSensitiveTextRequest) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.Txt != "" {
		w.RawByte('"')
		w.RawString("txt")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.Txt)
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *Common_CheckSensitiveTextResponse) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.IsPass {
		w.RawByte('"')
		w.RawString("is_pass")
		w.RawByte('"')
		w.RawByte(':')
		w.Bool(m.IsPass)
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *Empty) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *Empty) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *Empty) GoString() string {
	return m.String()
}

func (m *Common) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *Common) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *Common) GoString() string {
	return m.String()
}

func (m *Common_CheckSensitiveTextRequest) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *Common_CheckSensitiveTextRequest) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *Common_CheckSensitiveTextRequest) GoString() string {
	return m.String()
}

func (m *Common_CheckSensitiveTextResponse) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *Common_CheckSensitiveTextResponse) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *Common_CheckSensitiveTextResponse) GoString() string {
	return m.String()
}

func (m *Empty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Common) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Common_CheckSensitiveTextRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Txt)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *Common_CheckSensitiveTextResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsPass {
		n += 2
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *Common) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Common: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Common: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *Common_CheckSensitiveTextRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: CheckSensitiveTextRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckSensitiveTextRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *Common_CheckSensitiveTextResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: CheckSensitiveTextResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckSensitiveTextResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPass", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
			m.IsPass = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
