// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Resource struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string            `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Operation            string            `protobuf:"bytes,3,opt,name=operation,proto3" json:"operation,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Resource) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *Resource) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Event struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StartTime            float64    `protobuf:"fixed64,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Resource             *Resource  `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	Origin               string     `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin,omitempty"`
	Duration             float64    `protobuf:"fixed64,5,opt,name=duration,proto3" json:"duration,omitempty"`
	ErrorCode            ErrorCode  `protobuf:"varint,6,opt,name=error_code,json=errorCode,proto3,enum=protocol.ErrorCode" json:"error_code,omitempty"`
	Exception            *Exception `protobuf:"bytes,7,opt,name=exception,proto3" json:"exception,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetStartTime() float64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Event) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Event) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Event) GetDuration() float64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Event) GetErrorCode() ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return ErrorCode_OK
}

func (m *Event) GetException() *Exception {
	if m != nil {
		return m.Exception
	}
	return nil
}

func init() {
	proto.RegisterType((*Resource)(nil), "protocol.Resource")
	proto.RegisterMapType((map[string]string)(nil), "protocol.Resource.MetadataEntry")
	proto.RegisterType((*Event)(nil), "protocol.Event")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x4f, 0x4b, 0xfc, 0x30,
	0x14, 0xa4, 0xdd, 0x3f, 0xbf, 0xf6, 0x2d, 0xbf, 0x75, 0x79, 0x8a, 0x84, 0xa2, 0x50, 0xf6, 0xb4,
	0xa7, 0x82, 0xf5, 0x22, 0xea, 0x4d, 0xf6, 0xe8, 0x25, 0x78, 0x5f, 0x62, 0xfb, 0x90, 0xe0, 0xb6,
	0x29, 0xd9, 0xec, 0x62, 0x8f, 0x7e, 0x3e, 0xbf, 0x94, 0x24, 0x4d, 0x5b, 0xc5, 0x53, 0xe7, 0xcd,
	0xbc, 0xe9, 0x4c, 0x12, 0x58, 0xd0, 0x89, 0x6a, 0x93, 0x35, 0x5a, 0x19, 0x85, 0x91, 0xfb, 0x14,
	0x6a, 0x9f, 0xac, 0x48, 0x6b, 0xa5, 0x77, 0x85, 0x2a, 0xa9, 0xd3, 0x92, 0x33, 0xfa, 0x28, 0xa8,
	0x31, 0x52, 0xd5, 0x1d, 0xb1, 0xfe, 0x0a, 0x20, 0xe2, 0x74, 0x50, 0x47, 0x5d, 0x10, 0x22, 0x4c,
	0x6b, 0x51, 0x11, 0x0b, 0xd2, 0x60, 0x13, 0x73, 0x87, 0x2d, 0x67, 0xda, 0x86, 0x58, 0xd8, 0x71,
	0x16, 0xe3, 0x15, 0xc4, 0xaa, 0x21, 0x2d, 0xec, 0x7f, 0xd8, 0xc4, 0x09, 0x23, 0x81, 0x8f, 0x10,
	0x55, 0x64, 0x44, 0x29, 0x8c, 0x60, 0xd3, 0x74, 0xb2, 0x59, 0xe4, 0x69, 0xd6, 0x57, 0xca, 0xfa,
	0xac, 0xec, 0xd9, 0xaf, 0x6c, 0x6b, 0xa3, 0x5b, 0x3e, 0x38, 0x92, 0x07, 0xf8, 0xff, 0x4b, 0xc2,
	0x15, 0x4c, 0xde, 0xa9, 0xf5, 0x9d, 0x2c, 0xc4, 0x0b, 0x98, 0x9d, 0xc4, 0xfe, 0xd8, 0x77, 0xea,
	0x86, 0xfb, 0xf0, 0x2e, 0x58, 0x7f, 0x86, 0x30, 0xdb, 0xda, 0xab, 0xc0, 0x25, 0x84, 0xb2, 0xf4,
	0xa6, 0x50, 0x96, 0x78, 0x0d, 0x70, 0x30, 0x42, 0x9b, 0x9d, 0x91, 0x55, 0x67, 0x0c, 0x78, 0xec,
	0x98, 0x17, 0x59, 0x11, 0x66, 0x10, 0x69, 0xdf, 0xcc, 0x1d, 0x68, 0x91, 0xe3, 0xdf, 0xce, 0x7c,
	0xd8, 0xc1, 0x4b, 0x98, 0x2b, 0x2d, 0xdf, 0x64, 0xcd, 0xa6, 0x2e, 0xc2, 0x4f, 0x98, 0x40, 0x54,
	0x1e, 0xfd, 0xc5, 0xcc, 0x5c, 0xc8, 0x30, 0x63, 0x0e, 0x30, 0xbe, 0x07, 0x9b, 0xa7, 0xc1, 0x66,
	0x99, 0x9f, 0x8f, 0x29, 0x5b, 0xab, 0x3d, 0xa9, 0x92, 0x78, 0x4c, 0x3d, 0xc4, 0x1b, 0x88, 0x87,
	0x17, 0x63, 0xff, 0x5c, 0xb1, 0x9f, 0x96, 0x5e, 0xe2, 0xe3, 0xd6, 0xeb, 0xdc, 0xc9, 0xb7, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xad, 0xfc, 0xcc, 0x8a, 0x14, 0x02, 0x00, 0x00,
}