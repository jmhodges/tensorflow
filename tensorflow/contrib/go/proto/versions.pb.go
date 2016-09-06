// Code generated by protoc-gen-go.
// source: tensorflow/core/framework/versions.proto
// DO NOT EDIT!

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Version information for a piece of serialized data
//
// There are different types of versions for each type of data
// (GraphDef, etc.), but they all have the same common shape
// described here.
//
// Each consumer has "consumer" and "min_producer" versions (specified
// elsewhere).  A consumer is allowed to consume this data if
//
//   producer >= min_producer
//   consumer >= min_consumer
//   consumer not in bad_consumers
//
type VersionDef struct {
	// The version of the code that produced this data.
	Producer int32 `protobuf:"varint,1,opt,name=producer" json:"producer,omitempty"`
	// Any consumer below this version is not allowed to consume this data.
	MinConsumer int32 `protobuf:"varint,2,opt,name=min_consumer,json=minConsumer" json:"min_consumer,omitempty"`
	// Specific consumer versions which are disallowed (e.g. due to bugs).
	BadConsumers []int32 `protobuf:"varint,3,rep,name=bad_consumers,json=badConsumers" json:"bad_consumers,omitempty"`
}

func (m *VersionDef) Reset()                    { *m = VersionDef{} }
func (m *VersionDef) String() string            { return proto.CompactTextString(m) }
func (*VersionDef) ProtoMessage()               {}
func (*VersionDef) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func init() {
	proto.RegisterType((*VersionDef)(nil), "tensorflow.VersionDef")
}

var fileDescriptor15 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x8e, 0x3d, 0x0e, 0x82, 0x40,
	0x10, 0x46, 0xb3, 0x12, 0x8d, 0x19, 0xd1, 0x62, 0x2b, 0x62, 0x85, 0xda, 0x50, 0xb1, 0x85, 0x37,
	0x40, 0x0f, 0x40, 0x2c, 0x6c, 0x0d, 0x3f, 0x8b, 0x21, 0x0a, 0x43, 0x66, 0x40, 0xae, 0x6e, 0xe9,
	0x8a, 0xb8, 0x94, 0xf3, 0xbd, 0x97, 0xc9, 0x83, 0xa0, 0xd5, 0x35, 0x23, 0x15, 0x4f, 0xec, 0x55,
	0x86, 0xa4, 0x55, 0x41, 0x49, 0xa5, 0x7b, 0xa4, 0x87, 0x7a, 0x69, 0xe2, 0x12, 0x6b, 0x0e, 0x1b,
	0xc2, 0x16, 0x25, 0x4c, 0xe6, 0xbe, 0x01, 0xb8, 0xfe, 0xe8, 0x59, 0x17, 0x72, 0x0b, 0x4b, 0xa3,
	0xe4, 0x5d, 0xa6, 0xc9, 0x13, 0xbe, 0x08, 0xe6, 0x17, 0x7b, 0xcb, 0x1d, 0xb8, 0x55, 0x59, 0xdf,
	0x32, 0xf3, 0xa7, 0xab, 0x0c, 0x9f, 0x0d, 0x7c, 0x65, 0xb6, 0xd3, 0x38, 0xc9, 0x03, 0xac, 0xd3,
	0x24, 0xb7, 0x0a, 0x7b, 0x8e, 0xef, 0x18, 0xc7, 0x35, 0xe3, 0xdf, 0xe1, 0x48, 0x81, 0x87, 0x74,
	0x0f, 0xa7, 0x86, 0xd0, 0x86, 0x46, 0x9b, 0xb1, 0x85, 0xe3, 0x6f, 0x28, 0xc7, 0xe2, 0x2d, 0x44,
	0xba, 0x18, 0xaa, 0x8f, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xca, 0xc7, 0x87, 0xe1, 0x00,
	0x00, 0x00,
}