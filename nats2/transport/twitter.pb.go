// Code generated by protoc-gen-go. DO NOT EDIT.
// source: twitter.proto

/*
Package Transport is a generated protocol buffer package.

It is generated from these files:
	twitter.proto

It has these top-level messages:
	Tweet
*/
package Transport

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Tweet struct {
	Tweet string `protobuf:"bytes,1,opt,name=tweet" json:"tweet,omitempty"`
}

func (m *Tweet) Reset()                    { *m = Tweet{} }
func (m *Tweet) String() string            { return proto.CompactTextString(m) }
func (*Tweet) ProtoMessage()               {}
func (*Tweet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Tweet) GetTweet() string {
	if m != nil {
		return m.Tweet
	}
	return ""
}

func init() {
	proto.RegisterType((*Tweet)(nil), "Transport.Tweet")
}

func init() { proto.RegisterFile("twitter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 80 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x29, 0xcf, 0x2c,
	0x29, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0c, 0x29, 0x4a, 0xcc, 0x2b,
	0x2e, 0xc8, 0x2f, 0x2a, 0x51, 0x92, 0xe5, 0x62, 0x0d, 0x29, 0x4f, 0x4d, 0x2d, 0x11, 0x12, 0xe1,
	0x62, 0x2d, 0x01, 0x31, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x9c, 0x24, 0x36, 0xb0,
	0x06, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x30, 0x4b, 0xfa, 0x41, 0x00, 0x00, 0x00,
}
