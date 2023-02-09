// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/dao/formation.proto

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

type Formation struct {
	RoleId         string             `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty" pk`
	Assembles      []*models.Assemble `protobuf:"bytes,2,rep,name=assembles,proto3" json:"assembles,omitempty"`
	DefaultIndex   int64              `protobuf:"varint,3,opt,name=default_index,json=defaultIndex,proto3" json:"default_index,omitempty"`
	SetDefaultTime int64              `protobuf:"varint,4,opt,name=set_default_time,json=setDefaultTime,proto3" json:"set_default_time,omitempty"`
}

func (m *Formation) Reset()      { *m = Formation{} }
func (*Formation) ProtoMessage() {}
func (*Formation) Descriptor() ([]byte, []int) {
	return fileDescriptor_a58cdfc82c20bb14, []int{0}
}
func (m *Formation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Formation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Formation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Formation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Formation.Merge(m, src)
}
func (m *Formation) XXX_Size() int {
	return m.Size()
}
func (m *Formation) XXX_DiscardUnknown() {
	xxx_messageInfo_Formation.DiscardUnknown(m)
}

var xxx_messageInfo_Formation proto.InternalMessageInfo

func (m *Formation) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *Formation) GetAssembles() []*models.Assemble {
	if m != nil {
		return m.Assembles
	}
	return nil
}

func (m *Formation) GetDefaultIndex() int64 {
	if m != nil {
		return m.DefaultIndex
	}
	return 0
}

func (m *Formation) GetSetDefaultTime() int64 {
	if m != nil {
		return m.SetDefaultTime
	}
	return 0
}

func (*Formation) XXX_MessageName() string {
	return "dao.Formation"
}
func init() {
	proto.RegisterType((*Formation)(nil), "dao.Formation")
}

func init() { proto.RegisterFile("proto/dao/formation.proto", fileDescriptor_a58cdfc82c20bb14) }

var fileDescriptor_a58cdfc82c20bb14 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd0, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc0, 0xf1, 0x6c, 0x23, 0x95, 0xae, 0x1f, 0x94, 0xe0, 0x21, 0x96, 0x32, 0x2d, 0x7a, 0xc9,
	0xc5, 0x04, 0xf4, 0x09, 0x2c, 0x22, 0xd4, 0x63, 0xf0, 0xe4, 0x25, 0xa4, 0xdd, 0x69, 0x59, 0xcc,
	0x66, 0x4a, 0x76, 0x15, 0x1f, 0xc3, 0xc7, 0x10, 0x7c, 0x91, 0x1e, 0x7b, 0xec, 0x49, 0x74, 0x73,
	0xf1, 0x28, 0x3e, 0x81, 0x24, 0x69, 0x15, 0xbc, 0x0d, 0xbf, 0xf9, 0xef, 0x1e, 0x86, 0x1f, 0x2f,
	0x0a, 0x32, 0x14, 0x89, 0x94, 0xa2, 0x19, 0x15, 0x2a, 0x35, 0x92, 0xf2, 0xb0, 0x36, 0xcf, 0x15,
	0x29, 0xf5, 0x8e, 0xe6, 0x34, 0xa7, 0xa6, 0xa9, 0xa6, 0x66, 0xd5, 0xeb, 0x37, 0xa2, 0x48, 0x60,
	0xa6, 0xff, 0x3f, 0x3c, 0x79, 0x65, 0xbc, 0x73, 0xbd, 0x35, 0x6f, 0xc0, 0x77, 0x0b, 0xca, 0x30,
	0x91, 0xc2, 0x67, 0x43, 0x16, 0x74, 0x46, 0xed, 0xef, 0xb7, 0x41, 0x6b, 0x71, 0x1f, 0xb7, 0x2b,
	0x1e, 0x0b, 0x2f, 0xe4, 0x9d, 0x54, 0x6b, 0x54, 0x93, 0x0c, 0xb5, 0xdf, 0x1a, 0xba, 0xc1, 0xde,
	0x79, 0x37, 0x6c, 0xbe, 0x0e, 0x2f, 0x37, 0x8b, 0xf8, 0x2f, 0xf1, 0x4e, 0xf9, 0x81, 0xc0, 0x59,
	0xfa, 0x90, 0x99, 0x44, 0xe6, 0x02, 0x9f, 0x7c, 0x77, 0xc8, 0x02, 0x37, 0xde, 0xdf, 0xe0, 0xb8,
	0x32, 0x2f, 0xe0, 0x5d, 0x8d, 0x26, 0xd9, 0x86, 0x46, 0x2a, 0xf4, 0x77, 0xea, 0xee, 0x50, 0xa3,
	0xb9, 0x6a, 0xf8, 0x56, 0x2a, 0x1c, 0xdd, 0xac, 0x3f, 0xc0, 0x79, 0xb1, 0xc0, 0x96, 0x16, 0xd8,
	0xca, 0x02, 0x7b, 0xb7, 0xc0, 0x3e, 0x2d, 0x38, 0x5f, 0x16, 0xd8, 0x73, 0x09, 0xce, 0xb2, 0x04,
	0xb6, 0x2a, 0xc1, 0x59, 0x97, 0xe0, 0xdc, 0xf5, 0xa7, 0x24, 0xf3, 0x33, 0x8d, 0xc5, 0x23, 0x16,
	0xd1, 0x94, 0x94, 0xa2, 0x3c, 0xfa, 0x3d, 0xe1, 0xa4, 0x5d, 0x8f, 0x17, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xa4, 0xa0, 0x97, 0x02, 0x56, 0x01, 0x00, 0x00,
}

func init() {
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolFormation.Get().(proto.Message)
	})
}

var poolFormation = &sync.Pool{New: func() interface{} { return &Formation{} }}

func (m *Formation) ReleasePool() { m.Reset(); poolFormation.Put(m); m = nil }

func (m *Formation) PK() string {
	if m == nil {
		return ""
	}
	return m.RoleId
}

func (m *Formation) PKAppendTo(d []byte) []byte {
	if m == nil {
		return d
	}
	return append(d, m.RoleId...)
}

func (m *Formation) ToKVSave() ([]byte, []byte) {
	msgName := m.XXX_MessageName()
	dk := coin_server_common_bytespool.GetSample(64)
	dk = dk[:0]
	dk = append(dk, msgName...)
	dk = append(dk, ':', 'k', ':')
	dk = m.PKAppendTo(dk)
	return dk, m.ToSave()
}

func (m *Formation) ToSave() []byte {
	msgName := m.XXX_MessageName()
	ml := len(msgName)
	d := coin_server_common_bytespool.GetSample(1 + ml + m.Size())
	d[0] = uint8(ml)
	copy(d[1:], msgName)
	_, _ = m.MarshalToSizedBuffer(d[1+ml:])
	return d
}

func (m *Formation) KVKey() string {
	return m.XXX_MessageName() + ":k:" + m.PK()
}

func (this *Formation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Formation)
	if !ok {
		that2, ok := that.(Formation)
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
	if len(this.Assembles) != len(that1.Assembles) {
		return false
	}
	for i := range this.Assembles {
		if !this.Assembles[i].Equal(that1.Assembles[i]) {
			return false
		}
	}
	if this.DefaultIndex != that1.DefaultIndex {
		return false
	}
	if this.SetDefaultTime != that1.SetDefaultTime {
		return false
	}
	return true
}
func (m *Formation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Formation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Formation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SetDefaultTime != 0 {
		i = encodeVarintFormation(dAtA, i, uint64(m.SetDefaultTime))
		i--
		dAtA[i] = 0x20
	}
	if m.DefaultIndex != 0 {
		i = encodeVarintFormation(dAtA, i, uint64(m.DefaultIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Assembles) > 0 {
		for iNdEx := len(m.Assembles) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Assembles[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFormation(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.RoleId) > 0 {
		i -= len(m.RoleId)
		copy(dAtA[i:], m.RoleId)
		i = encodeVarintFormation(dAtA, i, uint64(len(m.RoleId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFormation(dAtA []byte, offset int, v uint64) int {
	offset -= sovFormation(v)
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

func (m *Formation) JsonBytes(w *coin_server_common_jwriter.Writer) {
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
	w.RawString("assembles")
	w.RawByte('"')
	w.RawByte(':')
	if m.Assembles == nil {
		w.RawString("null")
	} else if len(m.Assembles) == 0 {
		w.RawString("[]")
	} else {
		w.RawByte('[')
		for i, v := range m.Assembles {
			v.JsonBytes(w)
			if i != len(m.Assembles)-1 {
				w.RawByte(',')
			}
		}
		w.RawByte(']')
	}
	needWriteComma = true
	if m.DefaultIndex != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("default_index")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.DefaultIndex))
		needWriteComma = true
	}
	if m.SetDefaultTime != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("set_default_time")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.SetDefaultTime))
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *Formation) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *Formation) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *Formation) GoString() string {
	return m.String()
}

func (m *Formation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RoleId)
	if l > 0 {
		n += 1 + l + sovFormation(uint64(l))
	}
	if len(m.Assembles) > 0 {
		for _, e := range m.Assembles {
			l = e.Size()
			n += 1 + l + sovFormation(uint64(l))
		}
	}
	if m.DefaultIndex != 0 {
		n += 1 + sovFormation(uint64(m.DefaultIndex))
	}
	if m.SetDefaultTime != 0 {
		n += 1 + sovFormation(uint64(m.SetDefaultTime))
	}
	return n
}

func sovFormation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFormation(x uint64) (n int) {
	return sovFormation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Formation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFormation
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
			return fmt.Errorf("proto: Formation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Formation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFormation
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
				return ErrInvalidLengthFormation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFormation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoleId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assembles", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFormation
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
				return ErrInvalidLengthFormation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFormation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Assembles = append(m.Assembles, &models.Assemble{})
			if err := m.Assembles[len(m.Assembles)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultIndex", wireType)
			}
			m.DefaultIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFormation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DefaultIndex |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetDefaultTime", wireType)
			}
			m.SetDefaultTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFormation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SetDefaultTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFormation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFormation
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
func skipFormation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFormation
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
					return 0, ErrIntOverflowFormation
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
					return 0, ErrIntOverflowFormation
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
				return 0, ErrInvalidLengthFormation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFormation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFormation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFormation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFormation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFormation = fmt.Errorf("proto: unexpected end of group")
)
