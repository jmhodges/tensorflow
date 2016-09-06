// Code generated by protoc-gen-go.
// source: tensorflow/core/framework/allocation_description.proto
// DO NOT EDIT!

/*
Package tensorflow is a generated protocol buffer package.

It is generated from these files:
	tensorflow/core/framework/allocation_description.proto
	tensorflow/core/framework/attr_value.proto
	tensorflow/core/framework/device_attributes.proto
	tensorflow/core/framework/function.proto
	tensorflow/core/framework/graph.proto
	tensorflow/core/framework/kernel_def.proto
	tensorflow/core/framework/log_memory.proto
	tensorflow/core/framework/op_def.proto
	tensorflow/core/framework/step_stats.proto
	tensorflow/core/framework/summary.proto
	tensorflow/core/framework/tensor.proto
	tensorflow/core/framework/tensor_description.proto
	tensorflow/core/framework/tensor_shape.proto
	tensorflow/core/framework/tensor_slice.proto
	tensorflow/core/framework/types.proto
	tensorflow/core/framework/versions.proto
	tensorflow/core/protobuf/config.proto
	tensorflow/core/protobuf/saver.proto
	tensorflow/core/util/saved_tensor_slice.proto

It has these top-level messages:
	AllocationDescription
	AttrValue
	NameAttrList
	DeviceAttributes
	FunctionDefLibrary
	FunctionDef
	GradientDef
	GraphDef
	KernelDef
	MemoryLogStep
	MemoryLogTensorAllocation
	MemoryLogTensorDeallocation
	MemoryLogTensorOutput
	MemoryLogRawAllocation
	MemoryLogRawDeallocation
	OpDef
	OpDeprecation
	OpList
	AllocatorMemoryUsed
	NodeOutput
	NodeExecStats
	DeviceStepStats
	StepStats
	HistogramProto
	Summary
	TensorProto
	TensorDescription
	TensorShapeProto
	TensorSliceProto
	VersionDef
	GPUOptions
	OptimizerOptions
	GraphOptions
	ThreadPoolOptionProto
	ConfigProto
	DebugTensorWatch
	RunOptions
	RunMetadata
	SaverDef
	SavedSliceMeta
	SavedTensorSliceMeta
	SavedSlice
	SavedTensorSlices
*/
package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type AllocationDescription struct {
	// Total number of bytes requested
	RequestedBytes int64 `protobuf:"varint,1,opt,name=requested_bytes,json=requestedBytes" json:"requested_bytes,omitempty"`
	// Total number of bytes allocated if known
	AllocatedBytes int64 `protobuf:"varint,2,opt,name=allocated_bytes,json=allocatedBytes" json:"allocated_bytes,omitempty"`
	// Name of the allocator used
	AllocatorName string `protobuf:"bytes,3,opt,name=allocator_name,json=allocatorName" json:"allocator_name,omitempty"`
	// Identifier of the allocated buffer if known
	AllocationId int64 `protobuf:"varint,4,opt,name=allocation_id,json=allocationId" json:"allocation_id,omitempty"`
	// Set if this tensor only has one remaining reference
	HasSingleReference bool `protobuf:"varint,5,opt,name=has_single_reference,json=hasSingleReference" json:"has_single_reference,omitempty"`
	// Address of the allocation.
	Ptr uint64 `protobuf:"varint,6,opt,name=ptr" json:"ptr,omitempty"`
}

func (m *AllocationDescription) Reset()                    { *m = AllocationDescription{} }
func (m *AllocationDescription) String() string            { return proto.CompactTextString(m) }
func (*AllocationDescription) ProtoMessage()               {}
func (*AllocationDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*AllocationDescription)(nil), "tensorflow.AllocationDescription")
}

var fileDescriptor0 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0x5d, 0x17, 0x1d, 0x5c, 0x95, 0xa0, 0x10, 0xf0, 0x22, 0x8a, 0xe8, 0xa9, 0x15,
	0x04, 0x4f, 0x5e, 0x2c, 0x5e, 0xbc, 0xc8, 0x52, 0x1f, 0x20, 0x74, 0xdb, 0xe9, 0x5a, 0xac, 0x99,
	0x3a, 0x89, 0x2c, 0xbe, 0xb9, 0xde, 0x4c, 0xba, 0x6e, 0xea, 0xc1, 0xdb, 0xf0, 0xfd, 0x5f, 0x32,
	0xc3, 0x0f, 0xb7, 0x0e, 0x8d, 0x25, 0x6e, 0x3a, 0x5a, 0x65, 0x15, 0x31, 0x66, 0x0d, 0x97, 0x6f,
	0xb8, 0x22, 0x7e, 0xcd, 0xca, 0xae, 0xa3, 0xaa, 0x74, 0x2d, 0x19, 0x5d, 0xa3, 0xad, 0xb8, 0xed,
	0xc3, 0x9c, 0xf6, 0x4c, 0x8e, 0x24, 0x8c, 0xef, 0xce, 0xbe, 0x05, 0x1c, 0xdf, 0x47, 0xf9, 0x61,
	0x74, 0xe5, 0x25, 0x1c, 0x30, 0xbe, 0x7f, 0xa0, 0x75, 0x58, 0xeb, 0xc5, 0xa7, 0x43, 0xab, 0xc4,
	0xa9, 0xb8, 0x4a, 0x8a, 0xfd, 0x88, 0xf3, 0x40, 0x83, 0xf8, 0xbb, 0x2e, 0x8a, 0x5b, 0x6b, 0x31,
	0xe2, 0xb5, 0x78, 0x01, 0x1b, 0x42, 0xac, 0x8d, 0x3f, 0x55, 0x25, 0xde, 0xdb, 0x2d, 0x66, 0x91,
	0x3e, 0x79, 0x28, 0xcf, 0x61, 0xf6, 0xe7, 0xfc, 0xb6, 0x56, 0x93, 0xe1, 0xb7, 0xbd, 0x11, 0x3e,
	0xd6, 0xf2, 0x1a, 0x8e, 0x5e, 0x4a, 0xab, 0x6d, 0x6b, 0x96, 0x1d, 0x6a, 0xc6, 0x06, 0x19, 0x4d,
	0x85, 0x6a, 0xdb, 0xbb, 0x3b, 0x85, 0xf4, 0xd9, 0xf3, 0x10, 0x15, 0x9b, 0x44, 0x1e, 0x42, 0xd2,
	0x3b, 0x56, 0x53, 0x2f, 0x4c, 0x8a, 0x30, 0xe6, 0x77, 0xa0, 0x88, 0x97, 0xe9, 0xd8, 0x46, 0x1a,
	0x0b, 0xcc, 0x4f, 0xfe, 0x2d, 0x65, 0x1e, 0xfa, 0xb3, 0x73, 0xf1, 0x25, 0xc4, 0x62, 0x3a, 0x94,
	0x79, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xb7, 0xe1, 0x85, 0x86, 0x01, 0x00, 0x00,
}
