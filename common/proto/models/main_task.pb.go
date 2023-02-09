// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/models/main_task.proto

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

type MainTask struct {
	TaskId     int64           `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Progress   map[int64]int64 `protobuf:"bytes,2,rep,name=progress,proto3" json:"progress,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Finish     map[int64]bool  `protobuf:"bytes,3,rep,name=finish,proto3" json:"finish,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Status     TaskStatus      `protobuf:"varint,4,opt,name=status,proto3,enum=models.TaskStatus" json:"status,omitempty"`
	Last       int64           `protobuf:"varint,5,opt,name=last,proto3" json:"last,omitempty"`
	AcceptTime int64           `protobuf:"varint,6,opt,name=accept_time,json=acceptTime,proto3" json:"accept_time,omitempty"`
}

func (m *MainTask) Reset()      { *m = MainTask{} }
func (*MainTask) ProtoMessage() {}
func (*MainTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_9685d68c9be6ab3f, []int{0}
}
func (m *MainTask) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MainTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MainTask.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MainTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MainTask.Merge(m, src)
}
func (m *MainTask) XXX_Size() int {
	return m.Size()
}
func (m *MainTask) XXX_DiscardUnknown() {
	xxx_messageInfo_MainTask.DiscardUnknown(m)
}

var xxx_messageInfo_MainTask proto.InternalMessageInfo

func (m *MainTask) GetTaskId() int64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *MainTask) GetProgress() map[int64]int64 {
	if m != nil {
		return m.Progress
	}
	return nil
}

func (m *MainTask) GetFinish() map[int64]bool {
	if m != nil {
		return m.Finish
	}
	return nil
}

func (m *MainTask) GetStatus() TaskStatus {
	if m != nil {
		return m.Status
	}
	return TaskStatus_NotStarted
}

func (m *MainTask) GetLast() int64 {
	if m != nil {
		return m.Last
	}
	return 0
}

func (m *MainTask) GetAcceptTime() int64 {
	if m != nil {
		return m.AcceptTime
	}
	return 0
}

func (*MainTask) XXX_MessageName() string {
	return "models.MainTask"
}

type MainTaskChapterFinish struct {
	Finish map[int64]RewardStatus `protobuf:"bytes,1,rep,name=finish,proto3" json:"finish,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=models.RewardStatus"`
}

func (m *MainTaskChapterFinish) Reset()      { *m = MainTaskChapterFinish{} }
func (*MainTaskChapterFinish) ProtoMessage() {}
func (*MainTaskChapterFinish) Descriptor() ([]byte, []int) {
	return fileDescriptor_9685d68c9be6ab3f, []int{1}
}
func (m *MainTaskChapterFinish) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MainTaskChapterFinish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MainTaskChapterFinish.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MainTaskChapterFinish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MainTaskChapterFinish.Merge(m, src)
}
func (m *MainTaskChapterFinish) XXX_Size() int {
	return m.Size()
}
func (m *MainTaskChapterFinish) XXX_DiscardUnknown() {
	xxx_messageInfo_MainTaskChapterFinish.DiscardUnknown(m)
}

var xxx_messageInfo_MainTaskChapterFinish proto.InternalMessageInfo

func (m *MainTaskChapterFinish) GetFinish() map[int64]RewardStatus {
	if m != nil {
		return m.Finish
	}
	return nil
}

func (*MainTaskChapterFinish) XXX_MessageName() string {
	return "models.MainTaskChapterFinish"
}

type MainTaskFinish struct {
	ChapterFinish map[int64]*MainTaskChapterFinish `protobuf:"bytes,1,rep,name=chapter_finish,json=chapterFinish,proto3" json:"chapter_finish,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *MainTaskFinish) Reset()      { *m = MainTaskFinish{} }
func (*MainTaskFinish) ProtoMessage() {}
func (*MainTaskFinish) Descriptor() ([]byte, []int) {
	return fileDescriptor_9685d68c9be6ab3f, []int{2}
}
func (m *MainTaskFinish) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MainTaskFinish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MainTaskFinish.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MainTaskFinish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MainTaskFinish.Merge(m, src)
}
func (m *MainTaskFinish) XXX_Size() int {
	return m.Size()
}
func (m *MainTaskFinish) XXX_DiscardUnknown() {
	xxx_messageInfo_MainTaskFinish.DiscardUnknown(m)
}

var xxx_messageInfo_MainTaskFinish proto.InternalMessageInfo

func (m *MainTaskFinish) GetChapterFinish() map[int64]*MainTaskChapterFinish {
	if m != nil {
		return m.ChapterFinish
	}
	return nil
}

func (*MainTaskFinish) XXX_MessageName() string {
	return "models.MainTaskFinish"
}
func init() {
	proto.RegisterType((*MainTask)(nil), "models.MainTask")
	proto.RegisterMapType((map[int64]bool)(nil), "models.MainTask.FinishEntry")
	proto.RegisterMapType((map[int64]int64)(nil), "models.MainTask.ProgressEntry")
	proto.RegisterType((*MainTaskChapterFinish)(nil), "models.MainTaskChapterFinish")
	proto.RegisterMapType((map[int64]RewardStatus)(nil), "models.MainTaskChapterFinish.FinishEntry")
	proto.RegisterType((*MainTaskFinish)(nil), "models.MainTaskFinish")
	proto.RegisterMapType((map[int64]*MainTaskChapterFinish)(nil), "models.MainTaskFinish.ChapterFinishEntry")
}

func init() { proto.RegisterFile("proto/models/main_task.proto", fileDescriptor_9685d68c9be6ab3f) }

var fileDescriptor_9685d68c9be6ab3f = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x4f, 0x6b, 0x14, 0x31,
	0x18, 0xc6, 0x27, 0xbb, 0xed, 0xb8, 0xbc, 0x4b, 0x17, 0x09, 0x15, 0xc3, 0x50, 0xd3, 0xa5, 0xa7,
	0xb5, 0xe0, 0x2e, 0x6c, 0x3d, 0x68, 0x3d, 0xa9, 0x28, 0x78, 0x28, 0x96, 0xb1, 0x27, 0x2f, 0x43,
	0x9c, 0x89, 0x36, 0xec, 0xce, 0x1f, 0x92, 0x6c, 0xa5, 0xdf, 0xc2, 0x8f, 0x21, 0xf8, 0x31, 0xbc,
	0xf4, 0xd8, 0x63, 0x8f, 0x3a, 0x73, 0xf1, 0x24, 0x7e, 0x04, 0x99, 0x64, 0xa6, 0x64, 0xda, 0xd2,
	0xde, 0xf2, 0xbe, 0xef, 0xf3, 0x24, 0xbf, 0xf7, 0x81, 0xc0, 0x56, 0x21, 0x73, 0x9d, 0xcf, 0xd2,
	0x3c, 0xe1, 0x4b, 0x35, 0x4b, 0x99, 0xc8, 0x22, 0xcd, 0xd4, 0x62, 0x6a, 0xda, 0xd8, 0xb7, 0xfd,
	0x80, 0x74, 0x54, 0x3c, 0x5b, 0xa5, 0xca, 0x2a, 0x76, 0xfe, 0xf6, 0x60, 0x70, 0xc0, 0x44, 0x76,
	0xc4, 0xd4, 0x02, 0x3f, 0x84, 0x7b, 0xb5, 0x39, 0x12, 0x09, 0x41, 0x63, 0x34, 0xe9, 0x87, 0x7e,
	0x5d, 0xbe, 0x4b, 0xf0, 0x3e, 0x0c, 0x0a, 0x99, 0x7f, 0x91, 0x5c, 0x29, 0xd2, 0x1b, 0xf7, 0x27,
	0xc3, 0x39, 0x9d, 0xda, 0xcb, 0xa6, 0xad, 0x79, 0x7a, 0xd8, 0x08, 0xde, 0x64, 0x5a, 0x9e, 0x86,
	0x97, 0x7a, 0xfc, 0x14, 0xfc, 0xcf, 0x22, 0x13, 0xea, 0x98, 0xf4, 0x8d, 0x73, 0xeb, 0x9a, 0xf3,
	0xad, 0x19, 0x5b, 0x5f, 0xa3, 0xc5, 0xbb, 0xe0, 0x2b, 0xcd, 0xf4, 0x4a, 0x91, 0xb5, 0x31, 0x9a,
	0x8c, 0xe6, 0xb8, 0x75, 0xd5, 0x8e, 0x0f, 0x66, 0x12, 0x36, 0x0a, 0x8c, 0x61, 0x6d, 0xc9, 0x94,
	0x26, 0xeb, 0x86, 0xd9, 0x9c, 0xf1, 0x36, 0x0c, 0x59, 0x1c, 0xf3, 0x42, 0x47, 0x5a, 0xa4, 0x9c,
	0xf8, 0x66, 0x04, 0xb6, 0x75, 0x24, 0x52, 0x1e, 0xbc, 0x80, 0x8d, 0x0e, 0x31, 0xbe, 0x0f, 0xfd,
	0x05, 0x3f, 0x6d, 0x16, 0xaf, 0x8f, 0x78, 0x13, 0xd6, 0x4f, 0xd8, 0x72, 0xc5, 0x49, 0xcf, 0xf4,
	0x6c, 0xb1, 0xdf, 0x7b, 0x86, 0x82, 0xe7, 0x30, 0x74, 0xa0, 0xef, 0xb2, 0x0e, 0x1c, 0xeb, 0xce,
	0x0f, 0x04, 0x0f, 0xda, 0xcd, 0x5f, 0x1f, 0xb3, 0x42, 0x73, 0x69, 0xaf, 0xc2, 0x2f, 0x2f, 0x83,
	0x42, 0x26, 0xa8, 0xc7, 0x57, 0x83, 0xea, 0xc8, 0x6f, 0x4a, 0x2d, 0x78, 0x7f, 0x17, 0xd7, 0xae,
	0xcb, 0x35, 0x9a, 0x6f, 0xb6, 0x4f, 0x84, 0xfc, 0x2b, 0x93, 0x49, 0x93, 0xab, 0x43, 0xfb, 0x13,
	0xc1, 0xa8, 0x7d, 0xbe, 0xc1, 0x3c, 0x84, 0x51, 0x6c, 0x41, 0xa2, 0xdb, 0x71, 0x1b, 0xce, 0x0e,
	0xb5, 0xc5, 0xdd, 0x88, 0xdd, 0x5e, 0x10, 0x01, 0xbe, 0x2e, 0xba, 0x01, 0x7e, 0xcf, 0x85, 0x1f,
	0xce, 0x1f, 0xdd, 0x9a, 0x8f, 0xb3, 0xc5, 0xab, 0x83, 0x8b, 0xdf, 0xd4, 0xfb, 0x5e, 0x52, 0x74,
	0x56, 0x52, 0x74, 0x5e, 0x52, 0xf4, 0xab, 0xa4, 0xe8, 0x4f, 0x49, 0xbd, 0x7f, 0x25, 0x45, 0xdf,
	0x2a, 0xea, 0x9d, 0x55, 0x14, 0x9d, 0x57, 0xd4, 0xbb, 0xa8, 0xa8, 0xf7, 0x71, 0x3b, 0xce, 0x45,
	0xf6, 0x44, 0x71, 0x79, 0xc2, 0xe5, 0x2c, 0xce, 0xd3, 0x34, 0xcf, 0x66, 0xee, 0xff, 0xf9, 0xe4,
	0x9b, 0x6a, 0xef, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0x86, 0x8c, 0xbf, 0x7c, 0x03, 0x00,
	0x00,
}

func init() {
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolMainTask.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolMainTaskChapterFinish.Get().(proto.Message)
	})
	coin_server_common_msgcreate.RegisterNewMessage(func() proto.Message {
		return poolMainTaskFinish.Get().(proto.Message)
	})
}

var poolMainTask = &sync.Pool{New: func() interface{} { return &MainTask{} }}

func (m *MainTask) ReleasePool() { m.Reset(); poolMainTask.Put(m); m = nil }

var poolMainTaskChapterFinish = &sync.Pool{New: func() interface{} { return &MainTaskChapterFinish{} }}

func (m *MainTaskChapterFinish) ReleasePool() { m.Reset(); poolMainTaskChapterFinish.Put(m); m = nil }

var poolMainTaskFinish = &sync.Pool{New: func() interface{} { return &MainTaskFinish{} }}

func (m *MainTaskFinish) ReleasePool() { m.Reset(); poolMainTaskFinish.Put(m); m = nil }
func (this *MainTask) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MainTask)
	if !ok {
		that2, ok := that.(MainTask)
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
	if this.TaskId != that1.TaskId {
		return false
	}
	if len(this.Progress) != len(that1.Progress) {
		return false
	}
	for i := range this.Progress {
		if this.Progress[i] != that1.Progress[i] {
			return false
		}
	}
	if len(this.Finish) != len(that1.Finish) {
		return false
	}
	for i := range this.Finish {
		if this.Finish[i] != that1.Finish[i] {
			return false
		}
	}
	if this.Status != that1.Status {
		return false
	}
	if this.Last != that1.Last {
		return false
	}
	if this.AcceptTime != that1.AcceptTime {
		return false
	}
	return true
}
func (this *MainTaskChapterFinish) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MainTaskChapterFinish)
	if !ok {
		that2, ok := that.(MainTaskChapterFinish)
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
	if len(this.Finish) != len(that1.Finish) {
		return false
	}
	for i := range this.Finish {
		if this.Finish[i] != that1.Finish[i] {
			return false
		}
	}
	return true
}
func (this *MainTaskFinish) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MainTaskFinish)
	if !ok {
		that2, ok := that.(MainTaskFinish)
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
	if len(this.ChapterFinish) != len(that1.ChapterFinish) {
		return false
	}
	for i := range this.ChapterFinish {
		if !this.ChapterFinish[i].Equal(that1.ChapterFinish[i]) {
			return false
		}
	}
	return true
}
func (m *MainTask) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MainTask) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MainTask) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AcceptTime != 0 {
		i = encodeVarintMainTask(dAtA, i, uint64(m.AcceptTime))
		i--
		dAtA[i] = 0x30
	}
	if m.Last != 0 {
		i = encodeVarintMainTask(dAtA, i, uint64(m.Last))
		i--
		dAtA[i] = 0x28
	}
	if m.Status != 0 {
		i = encodeVarintMainTask(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Finish) > 0 {
		for k := range m.Finish {
			v := m.Finish[k]
			baseI := i
			i--
			if v {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
			i = encodeVarintMainTask(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintMainTask(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Progress) > 0 {
		for k := range m.Progress {
			v := m.Progress[k]
			baseI := i
			i = encodeVarintMainTask(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i = encodeVarintMainTask(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintMainTask(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.TaskId != 0 {
		i = encodeVarintMainTask(dAtA, i, uint64(m.TaskId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MainTaskChapterFinish) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MainTaskChapterFinish) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MainTaskChapterFinish) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Finish) > 0 {
		for k := range m.Finish {
			v := m.Finish[k]
			baseI := i
			i = encodeVarintMainTask(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i = encodeVarintMainTask(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintMainTask(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MainTaskFinish) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MainTaskFinish) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MainTaskFinish) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChapterFinish) > 0 {
		for k := range m.ChapterFinish {
			v := m.ChapterFinish[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintMainTask(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintMainTask(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintMainTask(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintMainTask(dAtA []byte, offset int, v uint64) int {
	offset -= sovMainTask(v)
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

func (m *MainTask) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	if m.TaskId != 0 {
		w.RawByte('"')
		w.RawString("task_id")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.TaskId))
		needWriteComma = true
	}
	if needWriteComma {
		w.RawByte(',')
	}
	w.RawByte('"')
	w.RawString("progress")
	w.RawByte('"')
	w.RawByte(':')
	if m.Progress == nil {
		w.RawString("null")
	} else if len(m.Progress) == 0 {
		w.RawString("{}")
	} else {
		w.RawByte('{')
		mlProgress := len(m.Progress)
		for k, v := range m.Progress {
			w.RawByte('"')
			w.Int64(int64(k))
			w.RawByte('"')
			w.RawByte(':')
			w.Int64(int64(v))
			mlProgress--
			if mlProgress != 0 {
				w.RawByte(',')
			}
		}
		w.RawByte('}')
	}
	needWriteComma = true
	if needWriteComma {
		w.RawByte(',')
	}
	w.RawByte('"')
	w.RawString("finish")
	w.RawByte('"')
	w.RawByte(':')
	if m.Finish == nil {
		w.RawString("null")
	} else if len(m.Finish) == 0 {
		w.RawString("{}")
	} else {
		w.RawByte('{')
		mlFinish := len(m.Finish)
		for k, v := range m.Finish {
			w.RawByte('"')
			w.Int64(int64(k))
			w.RawByte('"')
			w.RawByte(':')
			w.Bool(v)
			mlFinish--
			if mlFinish != 0 {
				w.RawByte(',')
			}
		}
		w.RawByte('}')
	}
	needWriteComma = true
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
	if m.Last != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("last")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.Last))
		needWriteComma = true
	}
	if m.AcceptTime != 0 {
		if needWriteComma {
			w.RawByte(',')
		}
		w.RawByte('"')
		w.RawString("accept_time")
		w.RawByte('"')
		w.RawByte(':')
		w.Int64(int64(m.AcceptTime))
		needWriteComma = true
	}
	_ = needWriteComma
	w.RawByte('}')

}

func (m *MainTaskChapterFinish) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	w.RawByte('"')
	w.RawString("finish")
	w.RawByte('"')
	w.RawByte(':')
	if m.Finish == nil {
		w.RawString("null")
	} else if len(m.Finish) == 0 {
		w.RawString("{}")
	} else {
		w.RawByte('{')
		mlFinish := len(m.Finish)
		for k, v := range m.Finish {
			w.RawByte('"')
			w.Int64(int64(k))
			w.RawByte('"')
			w.RawByte(':')
			w.Int64(int64(v))
			mlFinish--
			if mlFinish != 0 {
				w.RawByte(',')
			}
		}
		w.RawByte('}')
	}
	needWriteComma = true
	_ = needWriteComma
	w.RawByte('}')

}

func (m *MainTaskFinish) JsonBytes(w *coin_server_common_jwriter.Writer) {
	if m == nil {
		w.RawString("null")
		return
	}

	w.RawByte('{')
	needWriteComma := false
	w.RawByte('"')
	w.RawString("chapter_finish")
	w.RawByte('"')
	w.RawByte(':')
	if m.ChapterFinish == nil {
		w.RawString("null")
	} else if len(m.ChapterFinish) == 0 {
		w.RawString("{}")
	} else {
		w.RawByte('{')
		mlChapterFinish := len(m.ChapterFinish)
		for k, v := range m.ChapterFinish {
			w.RawByte('"')
			w.Int64(int64(k))
			w.RawByte('"')
			w.RawByte(':')
			v.JsonBytes(w)
			mlChapterFinish--
			if mlChapterFinish != 0 {
				w.RawByte(',')
			}
		}
		w.RawByte('}')
	}
	needWriteComma = true
	_ = needWriteComma
	w.RawByte('}')

}

func (m *MainTask) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *MainTask) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *MainTask) GoString() string {
	return m.String()
}

func (m *MainTaskChapterFinish) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *MainTaskChapterFinish) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *MainTaskChapterFinish) GoString() string {
	return m.String()
}

func (m *MainTaskFinish) MarshalJSON() ([]byte, error) {
	w := coin_server_common_jwriter.Writer{Buffer: coin_server_common_buffer.Buffer{Buf: make([]byte, 0, 2048)}}
	m.JsonBytes(&w)
	return w.BuildBytes()
}
func (m *MainTaskFinish) String() string {
	d, _ := m.MarshalJSON()
	return *(*string)(unsafe.Pointer(&d))
}
func (m *MainTaskFinish) GoString() string {
	return m.String()
}

func (m *MainTask) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TaskId != 0 {
		n += 1 + sovMainTask(uint64(m.TaskId))
	}
	if len(m.Progress) > 0 {
		for k, v := range m.Progress {
			_ = k
			_ = v
			mapEntrySize := 1 + sovMainTask(uint64(k)) + 1 + sovMainTask(uint64(v))
			n += mapEntrySize + 1 + sovMainTask(uint64(mapEntrySize))
		}
	}
	if len(m.Finish) > 0 {
		for k, v := range m.Finish {
			_ = k
			_ = v
			mapEntrySize := 1 + sovMainTask(uint64(k)) + 1 + 1
			n += mapEntrySize + 1 + sovMainTask(uint64(mapEntrySize))
		}
	}
	if m.Status != 0 {
		n += 1 + sovMainTask(uint64(m.Status))
	}
	if m.Last != 0 {
		n += 1 + sovMainTask(uint64(m.Last))
	}
	if m.AcceptTime != 0 {
		n += 1 + sovMainTask(uint64(m.AcceptTime))
	}
	return n
}

func (m *MainTaskChapterFinish) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Finish) > 0 {
		for k, v := range m.Finish {
			_ = k
			_ = v
			mapEntrySize := 1 + sovMainTask(uint64(k)) + 1 + sovMainTask(uint64(v))
			n += mapEntrySize + 1 + sovMainTask(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *MainTaskFinish) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChapterFinish) > 0 {
		for k, v := range m.ChapterFinish {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovMainTask(uint64(l))
			}
			mapEntrySize := 1 + sovMainTask(uint64(k)) + l
			n += mapEntrySize + 1 + sovMainTask(uint64(mapEntrySize))
		}
	}
	return n
}

func sovMainTask(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMainTask(x uint64) (n int) {
	return sovMainTask(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MainTask) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMainTask
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
			return fmt.Errorf("proto: MainTask: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MainTask: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskId", wireType)
			}
			m.TaskId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMainTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TaskId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Progress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMainTask
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
				return ErrInvalidLengthMainTask
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMainTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Progress == nil {
				m.Progress = make(map[int64]int64)
			}
			var mapkey int64
			var mapvalue int64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMainTask
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMainTask
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMainTask
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMainTask(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthMainTask
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Progress[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Finish", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMainTask
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
				return ErrInvalidLengthMainTask
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMainTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Finish == nil {
				m.Finish = make(map[int64]bool)
			}
			var mapkey int64
			var mapvalue bool
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMainTask
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMainTask
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapvaluetemp int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMainTask
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvaluetemp |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					mapvalue = bool(mapvaluetemp != 0)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMainTask(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthMainTask
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Finish[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMainTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= TaskStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Last", wireType)
			}
			m.Last = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMainTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Last |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcceptTime", wireType)
			}
			m.AcceptTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMainTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AcceptTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMainTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMainTask
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
func (m *MainTaskChapterFinish) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMainTask
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
			return fmt.Errorf("proto: MainTaskChapterFinish: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MainTaskChapterFinish: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Finish", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMainTask
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
				return ErrInvalidLengthMainTask
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMainTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Finish == nil {
				m.Finish = make(map[int64]RewardStatus)
			}
			var mapkey int64
			var mapvalue RewardStatus
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMainTask
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMainTask
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMainTask
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= RewardStatus(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMainTask(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthMainTask
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Finish[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMainTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMainTask
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
func (m *MainTaskFinish) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMainTask
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
			return fmt.Errorf("proto: MainTaskFinish: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MainTaskFinish: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChapterFinish", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMainTask
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
				return ErrInvalidLengthMainTask
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMainTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ChapterFinish == nil {
				m.ChapterFinish = make(map[int64]*MainTaskChapterFinish)
			}
			var mapkey int64
			var mapvalue *MainTaskChapterFinish
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMainTask
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMainTask
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMainTask
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthMainTask
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthMainTask
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &MainTaskChapterFinish{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMainTask(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthMainTask
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ChapterFinish[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMainTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMainTask
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
func skipMainTask(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMainTask
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
					return 0, ErrIntOverflowMainTask
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
					return 0, ErrIntOverflowMainTask
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
				return 0, ErrInvalidLengthMainTask
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMainTask
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMainTask
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMainTask        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMainTask          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMainTask = fmt.Errorf("proto: unexpected end of group")
)
