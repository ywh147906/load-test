// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/cppcenter/cpp_center.proto

package cppcenter

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

type CppCenter struct {
}

func (m *CppCenter) Reset()      { *m = CppCenter{} }
func (*CppCenter) ProtoMessage() {}
func (*CppCenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4ae9adf11ed8ffa, []int{0}
}
func (m *CppCenter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CppCenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CppCenter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CppCenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CppCenter.Merge(m, src)
}
func (m *CppCenter) XXX_Size() int {
	return m.Size()
}
func (m *CppCenter) XXX_DiscardUnknown() {
	xxx_messageInfo_CppCenter.DiscardUnknown(m)
}

var xxx_messageInfo_CppCenter proto.InternalMessageInfo

func (*CppCenter) XXX_MessageName() string {
	return "cppcenter.CppCenter"
}

// 中心服启动，要求所有战斗服务推送自己状态
type CppCenter_CenterStartPush struct {
	ServerId int64 `protobuf:"varint,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (m *CppCenter_CenterStartPush) Reset()      { *m = CppCenter_CenterStartPush{} }
func (*CppCenter_CenterStartPush) ProtoMessage() {}
func (*CppCenter_CenterStartPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4ae9adf11ed8ffa, []int{0, 0}
}
func (m *CppCenter_CenterStartPush) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CppCenter_CenterStartPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CppCenter_CenterStartPush.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CppCenter_CenterStartPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CppCenter_CenterStartPush.Merge(m, src)
}
func (m *CppCenter_CenterStartPush) XXX_Size() int {
	return m.Size()
}
func (m *CppCenter_CenterStartPush) XXX_DiscardUnknown() {
	xxx_messageInfo_CppCenter_CenterStartPush.DiscardUnknown(m)
}

var xxx_messageInfo_CppCenter_CenterStartPush proto.InternalMessageInfo

func (m *CppCenter_CenterStartPush) GetServerId() int64 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (*CppCenter_CenterStartPush) XXX_MessageName() string {
	return "cppcenter.CppCenter.CenterStartPush"
}
func init() {
	proto.RegisterType((*CppCenter)(nil), "cppcenter.CppCenter")
	proto.RegisterType((*CppCenter_CenterStartPush)(nil), "cppcenter.CppCenter.CenterStartPush")
}

func init() { proto.RegisterFile("proto/cppcenter/cpp_center.proto", fileDescriptor_f4ae9adf11ed8ffa) }

var fileDescriptor_f4ae9adf11ed8ffa = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2e, 0x28, 0x48, 0x4e, 0xcd, 0x2b, 0x49, 0x2d, 0x02, 0xb1, 0xe2, 0x21, 0x4c,
	0x3d, 0xb0, 0x94, 0x10, 0x27, 0x5c, 0x4e, 0xc9, 0x9a, 0x8b, 0xd3, 0xb9, 0xa0, 0xc0, 0x19, 0xcc,
	0x91, 0xd2, 0xe3, 0xe2, 0x87, 0xb0, 0x82, 0x4b, 0x12, 0x8b, 0x4a, 0x02, 0x4a, 0x8b, 0x33, 0x84,
	0xa4, 0xb9, 0x38, 0x8b, 0x53, 0x8b, 0xca, 0x52, 0x8b, 0xe2, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x98, 0x83, 0x38, 0x20, 0x02, 0x9e, 0x29, 0x4e, 0x01, 0x37, 0x1e, 0xca, 0x31, 0xac, 0x78,
	0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xbe,
	0x78, 0x24, 0xc7, 0xf0, 0xe1, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x27, 0x1e, 0xcb, 0x31,
	0x5e, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x52, 0x72, 0x7e, 0x66, 0x9e, 0x2e,
	0x44, 0xb3, 0x7e, 0x72, 0x7e, 0x6e, 0x6e, 0x7e, 0x9e, 0x3e, 0x9a, 0x53, 0x93, 0xd8, 0xc0, 0x02,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xaf, 0x14, 0x73, 0xc4, 0x00, 0x00, 0x00,
}

func init() {
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolCppCenter.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolCppCenter_CenterStartPush.Get().(proto.Message)
	})
}

var poolCppCenter = &sync.Pool{New: func() interface{} { return &CppCenter{} }}

func (m *CppCenter) ReleasePool() { m.Reset(); poolCppCenter.Put(m); m = nil }

var poolCppCenter_CenterStartPush = &sync.Pool{New: func() interface{} { return &CppCenter_CenterStartPush{} }}

func (m *CppCenter_CenterStartPush) ReleasePool() {
	m.Reset()
	poolCppCenter_CenterStartPush.Put(m)
	m = nil
}
func (this *CppCenter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CppCenter)
	if !ok {
		that2, ok := that.(CppCenter)
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
func (this *CppCenter_CenterStartPush) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CppCenter_CenterStartPush)
	if !ok {
		that2, ok := that.(CppCenter_CenterStartPush)
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
	if this.ServerId != that1.ServerId {
		return false
	}
	return true
}
func (m *CppCenter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CppCenter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CppCenter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *CppCenter_CenterStartPush) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CppCenter_CenterStartPush) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CppCenter_CenterStartPush) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ServerId != 0 {
		i = encodeVarintCppCenter(dAtA, i, uint64(m.ServerId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCppCenter(dAtA []byte, offset int, v uint64) int {
	offset -= sovCppCenter(v)
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

func (m *CppCenter) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	w.RawByte('}')

}

func (m *CppCenter_CenterStartPush) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.ServerId != 0 {
		w.RawByte('"')
		w.RawString("server_id")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.ServerId))
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *CppCenter) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *CppCenter) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *CppCenter) GoString() string {
	return m.String()
}

func (m *CppCenter_CenterStartPush) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *CppCenter_CenterStartPush) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *CppCenter_CenterStartPush) GoString() string {
	return m.String()
}

func (m *CppCenter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *CppCenter_CenterStartPush) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ServerId != 0 {
		n += 1 + sovCppCenter(uint64(m.ServerId))
	}
	return n
}

func sovCppCenter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCppCenter(x uint64) (n int) {
	return sovCppCenter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CppCenter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCppCenter
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
			return fmt.Errorf("proto: CppCenter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CppCenter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCppCenter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCppCenter
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
func (m *CppCenter_CenterStartPush) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCppCenter
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
			return fmt.Errorf("proto: CenterStartPush: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CenterStartPush: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerId", wireType)
			}
			m.ServerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCppCenter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServerId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCppCenter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCppCenter
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
func skipCppCenter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCppCenter
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
					return 0, ErrIntOverflowCppCenter
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
					return 0, ErrIntOverflowCppCenter
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
				return 0, ErrInvalidLengthCppCenter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCppCenter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCppCenter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCppCenter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCppCenter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCppCenter = fmt.Errorf("proto: unexpected end of group")
)
