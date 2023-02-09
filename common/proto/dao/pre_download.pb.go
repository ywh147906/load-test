// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/dao/pre_download.proto

package dao

import (
	coin_server_common_buffer "github.com/ywh147906/load-test/common/buffer"
	coin_server_common_bytespool "github.com/ywh147906/load-test/common/bytespool"
	coin_server_common_jwriter "github.com/ywh147906/load-test/common/jwriter"
	coin_server_common_msgcreate "github.com/ywh147906/load-test/common/msgcreate"
	coin_server_common_proto_jsonany "github.com/ywh147906/load-test/common/proto/jsonany"
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

type PreDownloadData struct {
	RoleId     string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty" pk`
	IsReceived bool   `protobuf:"varint,2,opt,name=is_received,json=isReceived,proto3" json:"is_received,omitempty"`
}

func (m *PreDownloadData) Reset()      { *m = PreDownloadData{} }
func (*PreDownloadData) ProtoMessage() {}
func (*PreDownloadData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3cba655d454c391, []int{0}
}
func (m *PreDownloadData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PreDownloadData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PreDownloadData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PreDownloadData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreDownloadData.Merge(m, src)
}
func (m *PreDownloadData) XXX_Size() int {
	return m.Size()
}
func (m *PreDownloadData) XXX_DiscardUnknown() {
	xxx_messageInfo_PreDownloadData.DiscardUnknown(m)
}

var xxx_messageInfo_PreDownloadData proto.InternalMessageInfo

func (m *PreDownloadData) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *PreDownloadData) GetIsReceived() bool {
	if m != nil {
		return m.IsReceived
	}
	return false
}

func (*PreDownloadData) XXX_MessageName() string {
	return "dao.PreDownloadData"
}
func init() {
	proto.RegisterType((*PreDownloadData)(nil), "dao.PreDownloadData")
}

func init() { proto.RegisterFile("proto/dao/pre_download.proto", fileDescriptor_f3cba655d454c391) }

var fileDescriptor_f3cba655d454c391 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x49, 0xcc, 0xd7, 0x2f, 0x28, 0x4a, 0x8d, 0x4f, 0xc9, 0x2f, 0xcf, 0xcb, 0xc9,
	0x4f, 0x4c, 0xd1, 0x03, 0x0b, 0x0b, 0x31, 0xa7, 0x24, 0xe6, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7,
	0x43, 0x94, 0x81, 0x58, 0x10, 0x29, 0xa5, 0x60, 0x2e, 0xfe, 0x80, 0xa2, 0x54, 0x17, 0xa8, 0x7a,
	0x97, 0xc4, 0x92, 0x44, 0x21, 0x79, 0x2e, 0xf6, 0xa2, 0xfc, 0x9c, 0xd4, 0xf8, 0xcc, 0x14, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0xb6, 0x4f, 0xf7, 0xe4, 0x99, 0x0a, 0xb2, 0x83, 0xd8, 0x40,
	0xc2, 0x9e, 0x29, 0x42, 0xf2, 0x5c, 0xdc, 0x99, 0xc5, 0xf1, 0x45, 0xa9, 0xc9, 0xa9, 0x99, 0x65,
	0xa9, 0x29, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x5c, 0x99, 0xc5, 0x41, 0x50, 0x11, 0x27,
	0xaf, 0x1b, 0x0f, 0xe5, 0x18, 0x56, 0x3c, 0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23,
	0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x5f, 0x3c, 0x92, 0x63, 0xf8, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x13, 0x8f, 0xe5, 0x18, 0x2f, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21,
	0x4a, 0x26, 0x39, 0x3f, 0x33, 0x4f, 0xb7, 0x38, 0xb5, 0xa8, 0x2c, 0xb5, 0x48, 0x3f, 0x39, 0x3f,
	0x37, 0x37, 0x3f, 0x4f, 0x1f, 0xee, 0x9d, 0x24, 0x36, 0x30, 0xd3, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x3c, 0x80, 0x4c, 0x06, 0xe2, 0x00, 0x00, 0x00,
}

func init() {
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolPreDownloadData.Get().(proto.Message)
	})
}

var poolPreDownloadData = &sync.Pool{New: func() interface{} { return &PreDownloadData{} }}

func (m *PreDownloadData) ReleasePool() { m.Reset(); poolPreDownloadData.Put(m); m = nil }

func (m *PreDownloadData) PK() string {
	if m == nil {
		return ""
	}
	return m.RoleId
}

func (m *PreDownloadData) PKAppendTo(d []byte) []byte {
	if m == nil {
		return d
	}
	return append(d, m.RoleId...)
}

func (m *PreDownloadData) ToKVSave() ([]byte, []byte) {
	msgName := m.XXX_MessageName()
	dk := coin_server_common_bytespool.GetSample(64)
	dk = dk[:0]
	dk = append(dk, msgName...)
	dk = append(dk, ':', 'k', ':')
	dk = m.PKAppendTo(dk)
	return dk, m.ToSave()
}

func (m *PreDownloadData) ToSave() []byte {
	msgName := m.XXX_MessageName()
	ml := len(msgName)
	d := coin_server_common_bytespool.GetSample(1 + ml + m.Size())
	d[0] = uint8(ml)
	copy(d[1:], msgName)
	_, _ = m.MarshalToSizedBuffer(d[1+ml:])
	return d
}

func (m *PreDownloadData) KVKey() string {
	return m.XXX_MessageName() + ":k:" + m.PK()
}

func (this *PreDownloadData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PreDownloadData)
	if !ok {
		that2, ok := that.(PreDownloadData)
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
	if this.IsReceived != that1.IsReceived {
		return false
	}
	return true
}
func (m *PreDownloadData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PreDownloadData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PreDownloadData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0x10
	}
	if len(m.RoleId) > 0 {
		i -= len(m.RoleId)
		copy(dAtA[i:], m.RoleId)
		i = encodeVarintPreDownload(dAtA, i, uint64(len(m.RoleId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPreDownload(dAtA []byte, offset int, v uint64) int {
	offset -= sovPreDownload(v)
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

func (m *PreDownloadData) JsonBytes(w *coin_server_common_jwriter.Writer) {
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
	if m.IsReceived {
		if needWriteComma {
			w.RawByte(',')
		}
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

func (m *PreDownloadData) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *PreDownloadData) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *PreDownloadData) GoString() string {
	return m.String()
}

func (m *PreDownloadData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RoleId)
	if l > 0 {
		n += 1 + l + sovPreDownload(uint64(l))
	}
	if m.IsReceived {
		n += 2
	}
	return n
}

func sovPreDownload(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPreDownload(x uint64) (n int) {
	return sovPreDownload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PreDownloadData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPreDownload
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
			return fmt.Errorf("proto: PreDownloadData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PreDownloadData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPreDownload
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
				return ErrInvalidLengthPreDownload
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPreDownload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoleId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsReceived", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPreDownload
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
			skippy, err := skipPreDownload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPreDownload
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
func skipPreDownload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPreDownload
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
					return 0, ErrIntOverflowPreDownload
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
					return 0, ErrIntOverflowPreDownload
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
				return 0, ErrInvalidLengthPreDownload
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPreDownload
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPreDownload
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPreDownload        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPreDownload          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPreDownload = fmt.Errorf("proto: unexpected end of group")
)