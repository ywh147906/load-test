// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/models/divination.proto

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

// 能量球
type EnergyBall struct {
	Quality int64 `protobuf:"varint,1,opt,name=quality,proto3" json:"quality,omitempty"`
	Exp     int64 `protobuf:"varint,2,opt,name=exp,proto3" json:"exp,omitempty"`
}

func (m *EnergyBall) Reset()      { *m = EnergyBall{} }
func (*EnergyBall) ProtoMessage() {}
func (*EnergyBall) Descriptor() ([]byte, []int) {
	return fileDescriptor_1af6c9e893686a2f, []int{0}
}
func (m *EnergyBall) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EnergyBall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EnergyBall.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EnergyBall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnergyBall.Merge(m, src)
}
func (m *EnergyBall) XXX_Size() int {
	return m.Size()
}
func (m *EnergyBall) XXX_DiscardUnknown() {
	xxx_messageInfo_EnergyBall.DiscardUnknown(m)
}

var xxx_messageInfo_EnergyBall proto.InternalMessageInfo

func (m *EnergyBall) GetQuality() int64 {
	if m != nil {
		return m.Quality
	}
	return 0
}

func (m *EnergyBall) GetExp() int64 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (*EnergyBall) XXX_MessageName() string {
	return "models.EnergyBall"
}

type Divination struct {
	TotalCount     int64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	AvailableCount int64 `protobuf:"varint,2,opt,name=available_count,json=availableCount,proto3" json:"available_count,omitempty"`
	ResetAt        int64 `protobuf:"varint,3,opt,name=reset_at,json=resetAt,proto3" json:"reset_at,omitempty"`
}

func (m *Divination) Reset()      { *m = Divination{} }
func (*Divination) ProtoMessage() {}
func (*Divination) Descriptor() ([]byte, []int) {
	return fileDescriptor_1af6c9e893686a2f, []int{1}
}
func (m *Divination) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Divination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Divination.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Divination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Divination.Merge(m, src)
}
func (m *Divination) XXX_Size() int {
	return m.Size()
}
func (m *Divination) XXX_DiscardUnknown() {
	xxx_messageInfo_Divination.DiscardUnknown(m)
}

var xxx_messageInfo_Divination proto.InternalMessageInfo

func (m *Divination) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *Divination) GetAvailableCount() int64 {
	if m != nil {
		return m.AvailableCount
	}
	return 0
}

func (m *Divination) GetResetAt() int64 {
	if m != nil {
		return m.ResetAt
	}
	return 0
}

func (*Divination) XXX_MessageName() string {
	return "models.Divination"
}
func init() {
	proto.RegisterType((*EnergyBall)(nil), "models.EnergyBall")
	proto.RegisterType((*Divination)(nil), "models.Divination")
}

func init() { proto.RegisterFile("proto/models/divination.proto", fileDescriptor_1af6c9e893686a2f) }

var fileDescriptor_1af6c9e893686a2f = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x6d, 0x22, 0x15, 0x74, 0x48, 0x80, 0x3c, 0x85, 0x81, 0x2b, 0xea, 0x02, 0x0b, 0x64,
	0x60, 0x61, 0xa5, 0xc0, 0xc8, 0xc2, 0xc8, 0x52, 0xb9, 0xa9, 0x85, 0x2c, 0x39, 0xbe, 0xd6, 0xb9,
	0x46, 0xf4, 0x2d, 0x78, 0x0c, 0x1e, 0xa5, 0x63, 0xc7, 0x8e, 0xe0, 0x2c, 0x8c, 0x3c, 0x02, 0xaa,
	0x53, 0x2a, 0x36, 0x7f, 0x9f, 0xff, 0xfb, 0x75, 0x3a, 0x38, 0x9b, 0x06, 0x62, 0x2a, 0x2a, 0x9a,
	0x18, 0x57, 0x17, 0x13, 0xdb, 0x58, 0xaf, 0xd9, 0x92, 0xbf, 0x4e, 0x5e, 0xf5, 0xba, 0x8f, 0xc1,
	0x2d, 0xc0, 0xa3, 0x37, 0xe1, 0x75, 0x31, 0xd4, 0xce, 0xa9, 0x1c, 0xf6, 0x67, 0x73, 0xed, 0x2c,
	0x2f, 0x72, 0x79, 0x2e, 0x2f, 0xb3, 0xe7, 0x3f, 0x54, 0x27, 0x90, 0x99, 0xb7, 0x69, 0xbe, 0x97,
	0xec, 0xe6, 0x39, 0x98, 0x01, 0x3c, 0xec, 0x5a, 0x55, 0x1f, 0x0e, 0x99, 0x58, 0xbb, 0x51, 0x49,
	0x73, 0xcf, 0xdb, 0x69, 0x48, 0xea, 0x7e, 0x63, 0xd4, 0x05, 0x1c, 0xeb, 0x46, 0x5b, 0xa7, 0xc7,
	0xce, 0x6c, 0x43, 0x5d, 0xd9, 0xd1, 0x4e, 0x77, 0xc1, 0x53, 0x38, 0x08, 0xa6, 0x36, 0x3c, 0xd2,
	0x9c, 0x67, 0xdd, 0x12, 0x89, 0xef, 0x78, 0xf8, 0xb4, 0xfe, 0x42, 0xf1, 0x11, 0x51, 0x2e, 0x23,
	0xca, 0x55, 0x44, 0xf9, 0x19, 0x51, 0x7e, 0x47, 0x14, 0x3f, 0x11, 0xe5, 0x7b, 0x8b, 0x62, 0xd9,
	0xa2, 0x5c, 0xb5, 0x28, 0xd6, 0x2d, 0x8a, 0x97, 0x7e, 0x49, 0xd6, 0x5f, 0xd5, 0x26, 0x34, 0x26,
	0x14, 0x25, 0x55, 0x15, 0xf9, 0xe2, 0xff, 0x51, 0xc6, 0xbd, 0x44, 0x37, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x7e, 0x99, 0x34, 0x1d, 0x2b, 0x01, 0x00, 0x00,
}

func init() {
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolEnergyBall.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolDivination.Get().(proto.Message)
	})
}

var poolEnergyBall = &sync.Pool{New: func() interface{} { return &EnergyBall{} }}

func (m *EnergyBall) ReleasePool() { m.Reset(); poolEnergyBall.Put(m); m = nil }

var poolDivination = &sync.Pool{New: func() interface{} { return &Divination{} }}

func (m *Divination) ReleasePool() { m.Reset(); poolDivination.Put(m); m = nil }
func (this *EnergyBall) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EnergyBall)
	if !ok {
		that2, ok := that.(EnergyBall)
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
	if this.Quality != that1.Quality {
		return false
	}
	if this.Exp != that1.Exp {
		return false
	}
	return true
}
func (this *Divination) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Divination)
	if !ok {
		that2, ok := that.(Divination)
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
	if this.TotalCount != that1.TotalCount {
		return false
	}
	if this.AvailableCount != that1.AvailableCount {
		return false
	}
	if this.ResetAt != that1.ResetAt {
		return false
	}
	return true
}
func (m *EnergyBall) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnergyBall) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EnergyBall) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Exp != 0 {
		i = encodeVarintDivination(dAtA, i, uint64(m.Exp))
		i--
		dAtA[i] = 0x10
	}
	if m.Quality != 0 {
		i = encodeVarintDivination(dAtA, i, uint64(m.Quality))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Divination) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Divination) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Divination) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ResetAt != 0 {
		i = encodeVarintDivination(dAtA, i, uint64(m.ResetAt))
		i--
		dAtA[i] = 0x18
	}
	if m.AvailableCount != 0 {
		i = encodeVarintDivination(dAtA, i, uint64(m.AvailableCount))
		i--
		dAtA[i] = 0x10
	}
	if m.TotalCount != 0 {
		i = encodeVarintDivination(dAtA, i, uint64(m.TotalCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDivination(dAtA []byte, offset int, v uint64) int {
	offset -= sovDivination(v)
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

func (m *EnergyBall) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.Quality != 0 {
		w.RawByte('"')
		w.RawString("quality")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.Quality))
		needWriteComma = true
	}
	if m.Exp != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("exp")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.Exp))
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *Divination) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.TotalCount != 0 {
		w.RawByte('"')
		w.RawString("total_count")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.TotalCount))
		needWriteComma = true
	}
	if m.AvailableCount != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("available_count")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.AvailableCount))
		needWriteComma = true
	}
	if m.ResetAt != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("reset_at")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.ResetAt))
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *EnergyBall) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *EnergyBall) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *EnergyBall) GoString() string {
	return m.String()
}

func (m *Divination) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *Divination) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *Divination) GoString() string {
	return m.String()
}

func (m *EnergyBall) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Quality != 0 {
		n += 1 + sovDivination(uint64(m.Quality))
	}
	if m.Exp != 0 {
		n += 1 + sovDivination(uint64(m.Exp))
	}
	return n
}

func (m *Divination) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalCount != 0 {
		n += 1 + sovDivination(uint64(m.TotalCount))
	}
	if m.AvailableCount != 0 {
		n += 1 + sovDivination(uint64(m.AvailableCount))
	}
	if m.ResetAt != 0 {
		n += 1 + sovDivination(uint64(m.ResetAt))
	}
	return n
}

func sovDivination(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDivination(x uint64) (n int) {
	return sovDivination(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EnergyBall) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDivination
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
			return fmt.Errorf("proto: EnergyBall: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnergyBall: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quality", wireType)
			}
			m.Quality = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDivination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Quality |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exp", wireType)
			}
			m.Exp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDivination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Exp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDivination(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDivination
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
func (m *Divination) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDivination
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
			return fmt.Errorf("proto: Divination: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Divination: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalCount", wireType)
			}
			m.TotalCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDivination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvailableCount", wireType)
			}
			m.AvailableCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDivination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AvailableCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResetAt", wireType)
			}
			m.ResetAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDivination
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResetAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDivination(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDivination
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
func skipDivination(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDivination
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
					return 0, ErrIntOverflowDivination
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
					return 0, ErrIntOverflowDivination
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
				return 0, ErrInvalidLengthDivination
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDivination
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDivination
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDivination        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDivination          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDivination = fmt.Errorf("proto: unexpected end of group")
)