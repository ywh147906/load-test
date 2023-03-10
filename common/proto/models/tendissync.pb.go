// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/models/tendissync.proto

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

type TendisCmd struct {
	Cmd string `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
}

func (m *TendisCmd) Reset()      { *m = TendisCmd{} }
func (*TendisCmd) ProtoMessage() {}
func (*TendisCmd) Descriptor() ([]byte, []int) {
	return fileDescriptor_08a4faf5b265a199, []int{0}
}
func (m *TendisCmd) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TendisCmd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TendisCmd.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TendisCmd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TendisCmd.Merge(m, src)
}
func (m *TendisCmd) XXX_Size() int {
	return m.Size()
}
func (m *TendisCmd) XXX_DiscardUnknown() {
	xxx_messageInfo_TendisCmd.DiscardUnknown(m)
}

var xxx_messageInfo_TendisCmd proto.InternalMessageInfo

func (m *TendisCmd) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (*TendisCmd) XXX_MessageName() string {
	return "models.TendisCmd"
}
func init() {
	proto.RegisterType((*TendisCmd)(nil), "models.TendisCmd")
}

func init() { proto.RegisterFile("proto/models/tendissync.proto", fileDescriptor_08a4faf5b265a199) }

var fileDescriptor_08a4faf5b265a199 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0xd6, 0x2f, 0x49, 0xcd, 0x4b, 0xc9, 0x2c, 0x2e,
	0xae, 0xcc, 0x4b, 0xd6, 0x03, 0x8b, 0x0b, 0xb1, 0x41, 0x24, 0x94, 0x64, 0xb9, 0x38, 0x43, 0xc0,
	0x72, 0xce, 0xb9, 0x29, 0x42, 0x02, 0x5c, 0xcc, 0xc9, 0xb9, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x20, 0xa6, 0x93, 0xef, 0x8d, 0x87, 0x72, 0x0c, 0x2b, 0x1e, 0xc9, 0x31, 0x9e, 0x78,
	0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x2f, 0x1e, 0xc9, 0x31, 0x7c,
	0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x89, 0xc7, 0x72, 0x8c, 0x17, 0x1e, 0xcb, 0x31,
	0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x9f, 0x9c, 0x9f, 0x99, 0xa7, 0x5b, 0x9c, 0x5a, 0x54, 0x96,
	0x5a, 0xa4, 0x9f, 0x9c, 0x9f, 0x9b, 0x9b, 0x9f, 0xa7, 0x8f, 0xec, 0x8c, 0x24, 0x36, 0x30, 0xcf,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x92, 0x18, 0xdf, 0x45, 0x9d, 0x00, 0x00, 0x00,
}

func init() {
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolTendisCmd.Get().(proto.Message)
	})
}

var poolTendisCmd = &sync.Pool{New: func() interface{} { return &TendisCmd{} }}

func (m *TendisCmd) ReleasePool() { m.Reset(); poolTendisCmd.Put(m); m = nil }
func (this *TendisCmd) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TendisCmd)
	if !ok {
		that2, ok := that.(TendisCmd)
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
	if this.Cmd != that1.Cmd {
		return false
	}
	return true
}
func (m *TendisCmd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TendisCmd) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TendisCmd) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cmd) > 0 {
		i -= len(m.Cmd)
		copy(dAtA[i:], m.Cmd)
		i = encodeVarintTendissync(dAtA, i, uint64(len(m.Cmd)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTendissync(dAtA []byte, offset int, v uint64) int {
	offset -= sovTendissync(v)
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

func (m *TendisCmd) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.Cmd != "" {
		w.RawByte('"')
		w.RawString("cmd")
		w.RawByte('"')
		w.RawByte(':')
		w.String(m.Cmd)
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *TendisCmd) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *TendisCmd) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *TendisCmd) GoString() string {
	return m.String()
}

func (m *TendisCmd) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cmd)
	if l > 0 {
		n += 1 + l + sovTendissync(uint64(l))
	}
	return n
}

func sovTendissync(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTendissync(x uint64) (n int) {
	return sovTendissync(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TendisCmd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTendissync
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
			return fmt.Errorf("proto: TendisCmd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TendisCmd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cmd", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTendissync
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
				return ErrInvalidLengthTendissync
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTendissync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cmd = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTendissync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTendissync
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
func skipTendissync(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTendissync
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
					return 0, ErrIntOverflowTendissync
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
					return 0, ErrIntOverflowTendissync
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
				return 0, ErrInvalidLengthTendissync
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTendissync
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTendissync
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTendissync        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTendissync          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTendissync = fmt.Errorf("proto: unexpected end of group")
)
