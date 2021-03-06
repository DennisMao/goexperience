// Code generated by protoc-gen-go. DO NOT EDIT.
// source: runtime/internal/stream_chunk.proto

package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// StreamError is a response type which is returned when
// streaming rpc returns an error.
type StreamError struct {
	GrpcCode             int32      `protobuf:"varint,1,opt,name=grpc_code,json=grpcCode" json:"grpc_code,omitempty"`
	HttpCode             int32      `protobuf:"varint,2,opt,name=http_code,json=httpCode" json:"http_code,omitempty"`
	Message              string     `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	HttpStatus           string     `protobuf:"bytes,4,opt,name=http_status,json=httpStatus" json:"http_status,omitempty"`
	Details              []*any.Any `protobuf:"bytes,5,rep,name=details" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamError) Reset()         { *m = StreamError{} }
func (m *StreamError) String() string { return proto.CompactTextString(m) }
func (*StreamError) ProtoMessage()    {}
func (*StreamError) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_chunk_6adb10b40700458b, []int{0}
}
func (m *StreamError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamError.Unmarshal(m, b)
}
func (m *StreamError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamError.Marshal(b, m, deterministic)
}
func (dst *StreamError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamError.Merge(dst, src)
}
func (m *StreamError) XXX_Size() int {
	return xxx_messageInfo_StreamError.Size(m)
}
func (m *StreamError) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamError.DiscardUnknown(m)
}

var xxx_messageInfo_StreamError proto.InternalMessageInfo

func (m *StreamError) GetGrpcCode() int32 {
	if m != nil {
		return m.GrpcCode
	}
	return 0
}

func (m *StreamError) GetHttpCode() int32 {
	if m != nil {
		return m.HttpCode
	}
	return 0
}

func (m *StreamError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StreamError) GetHttpStatus() string {
	if m != nil {
		return m.HttpStatus
	}
	return ""
}

func (m *StreamError) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamError)(nil), "grpc.gateway.runtime.StreamError")
}

func init() {
	proto.RegisterFile("runtime/internal/stream_chunk.proto", fileDescriptor_stream_chunk_6adb10b40700458b)
}

var fileDescriptor_stream_chunk_6adb10b40700458b = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0xc1, 0x4e, 0x84, 0x30,
	0x10, 0x86, 0x83, 0xeb, 0xba, 0xbb, 0xc3, 0x8d, 0xec, 0xa1, 0xea, 0x41, 0xa2, 0x17, 0x4e, 0x25,
	0xd1, 0x27, 0x50, 0xe3, 0x0b, 0xb0, 0x37, 0x2f, 0x9b, 0x59, 0x18, 0x0b, 0x11, 0x5a, 0x32, 0x1d,
	0x62, 0x78, 0x2d, 0x9f, 0xd0, 0xb4, 0xc8, 0xb1, 0xdf, 0xf7, 0xff, 0x93, 0x3f, 0x85, 0x27, 0x9e,
	0xac, 0x74, 0x03, 0x95, 0x9d, 0x15, 0x62, 0x8b, 0x7d, 0xe9, 0x85, 0x09, 0x87, 0x73, 0xdd, 0x4e,
	0xf6, 0x5b, 0x8f, 0xec, 0xc4, 0x65, 0x47, 0xc3, 0x63, 0xad, 0x0d, 0x0a, 0xfd, 0xe0, 0xac, 0xff,
	0x1b, 0x77, 0xb7, 0xc6, 0x39, 0xd3, 0x53, 0x19, 0x33, 0x97, 0xe9, 0xab, 0x44, 0x3b, 0x2f, 0x85,
	0xc7, 0xdf, 0x04, 0xd2, 0x53, 0xbc, 0xf3, 0xc1, 0xec, 0x38, 0xbb, 0x87, 0x43, 0x38, 0x71, 0xae,
	0x5d, 0x43, 0x2a, 0xc9, 0x93, 0x62, 0x5b, 0xed, 0x03, 0x78, 0x77, 0x0d, 0x05, 0xd9, 0x8a, 0x8c,
	0x8b, 0xbc, 0x5a, 0x64, 0x00, 0x51, 0x2a, 0xd8, 0x0d, 0xe4, 0x3d, 0x1a, 0x52, 0x9b, 0x3c, 0x29,
	0x0e, 0xd5, 0xfa, 0xcc, 0x1e, 0x20, 0x8d, 0x35, 0x2f, 0x28, 0x93, 0x57, 0xd7, 0xd1, 0x42, 0x40,
	0xa7, 0x48, 0x32, 0x0d, 0xbb, 0x86, 0x04, 0xbb, 0xde, 0xab, 0x6d, 0xbe, 0x29, 0xd2, 0xe7, 0xa3,
	0x5e, 0x16, 0xeb, 0x75, 0xb1, 0x7e, 0xb5, 0x73, 0xb5, 0x86, 0xde, 0xe0, 0x73, 0xbf, 0x7e, 0xc2,
	0xe5, 0x26, 0x46, 0x5e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x16, 0x75, 0x92, 0x08, 0x1f, 0x01,
	0x00, 0x00,
}
