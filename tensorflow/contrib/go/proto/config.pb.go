// Code generated by protoc-gen-go.
// source: tensorflow/core/protobuf/config.proto
// DO NOT EDIT!

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow2 "tensorflow/core/framework"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Optimization level
type OptimizerOptions_Level int32

const (
	// L1 is the default level.
	// Optimization performed at L1 :
	// 1. Common subexpression elimination
	// 2. Constant folding
	OptimizerOptions_L1 OptimizerOptions_Level = 0
	// No optimizations
	OptimizerOptions_L0 OptimizerOptions_Level = -1
)

var OptimizerOptions_Level_name = map[int32]string{
	0:  "L1",
	-1: "L0",
}
var OptimizerOptions_Level_value = map[string]int32{
	"L1": 0,
	"L0": -1,
}

func (x OptimizerOptions_Level) String() string {
	return proto.EnumName(OptimizerOptions_Level_name, int32(x))
}
func (OptimizerOptions_Level) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{1, 0} }

// TODO(pbar) Turn this into a TraceOptions proto which allows
// tracing to be controlled in a more orthogonal manner?
type RunOptions_TraceLevel int32

const (
	RunOptions_NO_TRACE       RunOptions_TraceLevel = 0
	RunOptions_SOFTWARE_TRACE RunOptions_TraceLevel = 1
	RunOptions_HARDWARE_TRACE RunOptions_TraceLevel = 2
	RunOptions_FULL_TRACE     RunOptions_TraceLevel = 3
)

var RunOptions_TraceLevel_name = map[int32]string{
	0: "NO_TRACE",
	1: "SOFTWARE_TRACE",
	2: "HARDWARE_TRACE",
	3: "FULL_TRACE",
}
var RunOptions_TraceLevel_value = map[string]int32{
	"NO_TRACE":       0,
	"SOFTWARE_TRACE": 1,
	"HARDWARE_TRACE": 2,
	"FULL_TRACE":     3,
}

func (x RunOptions_TraceLevel) String() string {
	return proto.EnumName(RunOptions_TraceLevel_name, int32(x))
}
func (RunOptions_TraceLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor16, []int{6, 0} }

type GPUOptions struct {
	// A value between 0 and 1 that indicates what fraction of the
	// available GPU memory to pre-allocate for each process.  1 means
	// to pre-allocate all of the GPU memory, 0.5 means the process
	// allocates ~50% of the available GPU memory.
	PerProcessGpuMemoryFraction float64 `protobuf:"fixed64,1,opt,name=per_process_gpu_memory_fraction,json=perProcessGpuMemoryFraction" json:"per_process_gpu_memory_fraction,omitempty"`
	// The type of GPU allocation strategy to use.
	//
	// Allowed values:
	// "": The empty string (default) uses a system-chosen default
	//     which may change over time.
	//
	// "BFC": A "Best-fit with coalescing" algorithm, simplified from a
	//        version of dlmalloc.
	AllocatorType string `protobuf:"bytes,2,opt,name=allocator_type,json=allocatorType" json:"allocator_type,omitempty"`
	// Delay deletion of up to this many bytes to reduce the number of
	// interactions with gpu driver code.  If 0, the system chooses
	// a reasonable default (several MBs).
	DeferredDeletionBytes int64 `protobuf:"varint,3,opt,name=deferred_deletion_bytes,json=deferredDeletionBytes" json:"deferred_deletion_bytes,omitempty"`
	// If true, the allocator does not pre-allocate the entire specified
	// GPU memory region, instead starting small and growing as needed.
	AllowGrowth bool `protobuf:"varint,4,opt,name=allow_growth,json=allowGrowth" json:"allow_growth,omitempty"`
	// A comma-separated list of GPU ids that determines the 'visible'
	// to 'virtual' mapping of GPU devices.  For example, if TensorFlow
	// can see 8 GPU devices in the process, and one wanted to map
	// visible GPU devices 5 and 3 as "/gpu:0", and "/gpu:1", then one
	// would specify this field as "5,3".  This field is similar in
	// spirit to the CUDA_VISIBLE_DEVICES environment variable, except
	// it applies to the visible GPU devices in the process.
	//
	// NOTE: The GPU driver provides the process with the visible GPUs
	// in an order which is not guaranteed to have any correlation to
	// the *physical* GPU id in the machine.  This field is used for
	// remapping "visible" to "virtual", which means this operates only
	// after the process starts.  Users are required to use vendor
	// specific mechanisms (e.g., CUDA_VISIBLE_DEVICES) to control the
	// physical to visible device mapping prior to invoking TensorFlow.
	VisibleDeviceList string `protobuf:"bytes,5,opt,name=visible_device_list,json=visibleDeviceList" json:"visible_device_list,omitempty"`
}

func (m *GPUOptions) Reset()                    { *m = GPUOptions{} }
func (m *GPUOptions) String() string            { return proto.CompactTextString(m) }
func (*GPUOptions) ProtoMessage()               {}
func (*GPUOptions) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

// Options passed to the graph optimizer
type OptimizerOptions struct {
	// If true, optimize the graph using common subexpression elimination.
	DoCommonSubexpressionElimination bool `protobuf:"varint,1,opt,name=do_common_subexpression_elimination,json=doCommonSubexpressionElimination" json:"do_common_subexpression_elimination,omitempty"`
	// If true, perform constant folding optimization on the graph.
	DoConstantFolding bool `protobuf:"varint,2,opt,name=do_constant_folding,json=doConstantFolding" json:"do_constant_folding,omitempty"`
	// If true, perform function inlining on the graph.
	DoFunctionInlining bool                   `protobuf:"varint,4,opt,name=do_function_inlining,json=doFunctionInlining" json:"do_function_inlining,omitempty"`
	OptLevel           OptimizerOptions_Level `protobuf:"varint,3,opt,name=opt_level,json=optLevel,enum=tensorflow.OptimizerOptions_Level" json:"opt_level,omitempty"`
}

func (m *OptimizerOptions) Reset()                    { *m = OptimizerOptions{} }
func (m *OptimizerOptions) String() string            { return proto.CompactTextString(m) }
func (*OptimizerOptions) ProtoMessage()               {}
func (*OptimizerOptions) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{1} }

type GraphOptions struct {
	// If true, use control flow to schedule the activation of Recv nodes.
	// (Currently ignored.)
	EnableRecvScheduling bool `protobuf:"varint,2,opt,name=enable_recv_scheduling,json=enableRecvScheduling" json:"enable_recv_scheduling,omitempty"`
	// Options controlling how graph is optimized.
	OptimizerOptions *OptimizerOptions `protobuf:"bytes,3,opt,name=optimizer_options,json=optimizerOptions" json:"optimizer_options,omitempty"`
	// The number of steps to run before returning a cost model detailing
	// the memory usage and performance of each node of the graph. 0 means
	// no cost model.
	BuildCostModel int64 `protobuf:"varint,4,opt,name=build_cost_model,json=buildCostModel" json:"build_cost_model,omitempty"`
	// Annotate each Node with Op output shape data, to the extent it can
	// be statically inferred.
	InferShapes bool `protobuf:"varint,5,opt,name=infer_shapes,json=inferShapes" json:"infer_shapes,omitempty"`
	// Only place the subgraphs that are run, rather than the entire graph.
	//
	// This is useful for interactive graph building, where one might
	// produce graphs that cannot be placed during the debugging
	// process.  In particular, it allows the client to continue work in
	// a session after adding a node to a graph whose placement
	// constraints are unsatisfiable.
	PlacePrunedGraph bool `protobuf:"varint,6,opt,name=place_pruned_graph,json=placePrunedGraph" json:"place_pruned_graph,omitempty"`
}

func (m *GraphOptions) Reset()                    { *m = GraphOptions{} }
func (m *GraphOptions) String() string            { return proto.CompactTextString(m) }
func (*GraphOptions) ProtoMessage()               {}
func (*GraphOptions) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{2} }

func (m *GraphOptions) GetOptimizerOptions() *OptimizerOptions {
	if m != nil {
		return m.OptimizerOptions
	}
	return nil
}

type ThreadPoolOptionProto struct {
	// The number of threads in the pool.
	//
	// 0 means the system picks a value based on where this option proto is used
	// (see the declaration of the specific field for more info).
	NumThreads int32 `protobuf:"varint,1,opt,name=num_threads,json=numThreads" json:"num_threads,omitempty"`
}

func (m *ThreadPoolOptionProto) Reset()                    { *m = ThreadPoolOptionProto{} }
func (m *ThreadPoolOptionProto) String() string            { return proto.CompactTextString(m) }
func (*ThreadPoolOptionProto) ProtoMessage()               {}
func (*ThreadPoolOptionProto) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{3} }

// Session configuration parameters.
// The system picks an appropriate values for fields that are not set.
type ConfigProto struct {
	// Map from device type name (e.g., "CPU" or "GPU" ) to maximum
	// number of devices of that type to use.  If a particular device
	// type is not found in the map, the system picks an appropriate
	// number.
	DeviceCount map[string]int32 `protobuf:"bytes,1,rep,name=device_count,json=deviceCount" json:"device_count,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// The execution of an individual op (for some op types) can be
	// parallelized on a pool of intra_op_parallelism_threads.
	// 0 means the system picks an appropriate number.
	IntraOpParallelismThreads int32 `protobuf:"varint,2,opt,name=intra_op_parallelism_threads,json=intraOpParallelismThreads" json:"intra_op_parallelism_threads,omitempty"`
	// Nodes that perform blocking operations are enqueued on a pool of
	// inter_op_parallelism_threads available in each process.
	//
	// 0 means the system picks an appropriate number.
	//
	// Note that the first Session created in the process sets the
	// number of threads for all future sessions unless use_per_session_threads is
	// true or session_inter_op_thread_pool is configured.
	InterOpParallelismThreads int32 `protobuf:"varint,5,opt,name=inter_op_parallelism_threads,json=interOpParallelismThreads" json:"inter_op_parallelism_threads,omitempty"`
	// If true, use a new set of threads for this session rather than the global
	// pool of threads. Only supported by direct sessions.
	//
	// If false, use the global threads created by the first session, or the
	// per-session thread pools configured by session_inter_op_thread_pool.
	//
	// This option is deprecated. The same effect can be achieved by setting
	// session_inter_op_thread_pool to have one element, whose num_threads equals
	// inter_op_parallelism_threads.
	UsePerSessionThreads bool `protobuf:"varint,9,opt,name=use_per_session_threads,json=usePerSessionThreads" json:"use_per_session_threads,omitempty"`
	// This option is experimental - it may be replaced with a different mechanism
	// in the future. The intended use is for when some session invocations need
	// to run in a background pool limited to a small number of threads.
	//
	// Configures session thread pools. If this is configured, then RunOptions for
	// a Run call can select the thread pool to use.
	//
	// If a pool's num_threads is 0, then inter_op_parallelism_threads is used.
	SessionInterOpThreadPool []*ThreadPoolOptionProto `protobuf:"bytes,12,rep,name=session_inter_op_thread_pool,json=sessionInterOpThreadPool" json:"session_inter_op_thread_pool,omitempty"`
	// Assignment of Nodes to Devices is recomputed every placement_period
	// steps until the system warms up (at which point the recomputation
	// typically slows down automatically).
	PlacementPeriod int32 `protobuf:"varint,3,opt,name=placement_period,json=placementPeriod" json:"placement_period,omitempty"`
	// When any filters are present sessions will ignore all devices which do not
	// match the filters. Each filter can be partially specified, e.g. "/job:ps"
	// "/job:worker/replica:3", etc.
	DeviceFilters []string `protobuf:"bytes,4,rep,name=device_filters,json=deviceFilters" json:"device_filters,omitempty"`
	// Options that apply to all GPUs.
	GpuOptions *GPUOptions `protobuf:"bytes,6,opt,name=gpu_options,json=gpuOptions" json:"gpu_options,omitempty"`
	// Whether soft placement is allowed. If allow_soft_placement is true,
	// an op will be placed on CPU if
	//   1. there's no GPU implementation for the OP
	// or
	//   2. no GPU devices are known or registered
	// or
	//   3. need to co-locate with reftype input(s) which are from CPU.
	AllowSoftPlacement bool `protobuf:"varint,7,opt,name=allow_soft_placement,json=allowSoftPlacement" json:"allow_soft_placement,omitempty"`
	// Whether device placements should be logged.
	LogDevicePlacement bool `protobuf:"varint,8,opt,name=log_device_placement,json=logDevicePlacement" json:"log_device_placement,omitempty"`
	// Options that apply to all graphs.
	GraphOptions *GraphOptions `protobuf:"bytes,10,opt,name=graph_options,json=graphOptions" json:"graph_options,omitempty"`
	// Global timeout for all blocking operations in this session.  If non-zero,
	// and not overridden on a per-operation basis, this value will be used as the
	// deadline for all blocking operations.
	OperationTimeoutInMs int64 `protobuf:"varint,11,opt,name=operation_timeout_in_ms,json=operationTimeoutInMs" json:"operation_timeout_in_ms,omitempty"`
}

func (m *ConfigProto) Reset()                    { *m = ConfigProto{} }
func (m *ConfigProto) String() string            { return proto.CompactTextString(m) }
func (*ConfigProto) ProtoMessage()               {}
func (*ConfigProto) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{4} }

func (m *ConfigProto) GetDeviceCount() map[string]int32 {
	if m != nil {
		return m.DeviceCount
	}
	return nil
}

func (m *ConfigProto) GetSessionInterOpThreadPool() []*ThreadPoolOptionProto {
	if m != nil {
		return m.SessionInterOpThreadPool
	}
	return nil
}

func (m *ConfigProto) GetGpuOptions() *GPUOptions {
	if m != nil {
		return m.GpuOptions
	}
	return nil
}

func (m *ConfigProto) GetGraphOptions() *GraphOptions {
	if m != nil {
		return m.GraphOptions
	}
	return nil
}

// EXPERIMENTAL. Option for watching a node.
type DebugTensorWatch struct {
	// Name of the node to watch.
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	// Output slot to watch.
	// The semantics of output_slot == -1 is that the node is only watched for
	// completion, but not for any output tensors. See NodeCompletionCallback
	// in debug_gateway.h.
	// TODO(cais): Implement this semantics.
	OutputSlot int32 `protobuf:"varint,2,opt,name=output_slot,json=outputSlot" json:"output_slot,omitempty"`
	// Name(s) of the debugging op(s).
	// One or more than one probes on a tensor.
	// e.g., {"DebugIdentity", "DebugNanCount"}
	DebugOps []string `protobuf:"bytes,3,rep,name=debug_ops,json=debugOps" json:"debug_ops,omitempty"`
	// URL(s) for debug targets(s).
	//   E.g., "file:///foo/tfdbg_dump", "grpc://localhost:11011"
	// Each debug op listed in debug_ops will publish its output tensor (debug
	// signal) to all URLs in debug_urls.
	DebugUrls []string `protobuf:"bytes,4,rep,name=debug_urls,json=debugUrls" json:"debug_urls,omitempty"`
}

func (m *DebugTensorWatch) Reset()                    { *m = DebugTensorWatch{} }
func (m *DebugTensorWatch) String() string            { return proto.CompactTextString(m) }
func (*DebugTensorWatch) ProtoMessage()               {}
func (*DebugTensorWatch) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{5} }

// EXPERIMENTAL. Options for a single Run() call.
type RunOptions struct {
	TraceLevel RunOptions_TraceLevel `protobuf:"varint,1,opt,name=trace_level,json=traceLevel,enum=tensorflow.RunOptions_TraceLevel" json:"trace_level,omitempty"`
	// Time to wait for operation to complete in milliseconds.
	TimeoutInMs int64 `protobuf:"varint,2,opt,name=timeout_in_ms,json=timeoutInMs" json:"timeout_in_ms,omitempty"`
	// The thread pool to use, if session_inter_op_thread_pool is configured.
	InterOpThreadPool int32 `protobuf:"varint,3,opt,name=inter_op_thread_pool,json=interOpThreadPool" json:"inter_op_thread_pool,omitempty"`
	// Debugging options
	DebugTensorWatchOpts []*DebugTensorWatch `protobuf:"bytes,4,rep,name=debug_tensor_watch_opts,json=debugTensorWatchOpts" json:"debug_tensor_watch_opts,omitempty"`
	// Whether the partition graph(s) executed by the executor(s) should be
	// outputted via RunMetadata.
	OutputPartitionGraphs bool `protobuf:"varint,5,opt,name=output_partition_graphs,json=outputPartitionGraphs" json:"output_partition_graphs,omitempty"`
}

func (m *RunOptions) Reset()                    { *m = RunOptions{} }
func (m *RunOptions) String() string            { return proto.CompactTextString(m) }
func (*RunOptions) ProtoMessage()               {}
func (*RunOptions) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{6} }

func (m *RunOptions) GetDebugTensorWatchOpts() []*DebugTensorWatch {
	if m != nil {
		return m.DebugTensorWatchOpts
	}
	return nil
}

// EXPERIMENTAL. Metadata output (i.e., non-Tensor) for a single Run() call.
type RunMetadata struct {
	// Statistics traced for this step. Populated if tracing is turned on via the
	// "RunOptions" proto.
	// EXPERIMENTAL: The format and set of events may change in future versions.
	StepStats *StepStats `protobuf:"bytes,1,opt,name=step_stats,json=stepStats" json:"step_stats,omitempty"`
	// The cost graph for the computation defined by the run call.
	CostGraph *tensorflow2.CostGraphDef `protobuf:"bytes,2,opt,name=cost_graph,json=costGraph" json:"cost_graph,omitempty"`
	// Graphs of the partitions executed by executors.
	PartitionGraphs []*GraphDef `protobuf:"bytes,3,rep,name=partition_graphs,json=partitionGraphs" json:"partition_graphs,omitempty"`
}

func (m *RunMetadata) Reset()                    { *m = RunMetadata{} }
func (m *RunMetadata) String() string            { return proto.CompactTextString(m) }
func (*RunMetadata) ProtoMessage()               {}
func (*RunMetadata) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{7} }

func (m *RunMetadata) GetStepStats() *StepStats {
	if m != nil {
		return m.StepStats
	}
	return nil
}

func (m *RunMetadata) GetCostGraph() *tensorflow2.CostGraphDef {
	if m != nil {
		return m.CostGraph
	}
	return nil
}

func (m *RunMetadata) GetPartitionGraphs() []*GraphDef {
	if m != nil {
		return m.PartitionGraphs
	}
	return nil
}

func init() {
	proto.RegisterType((*GPUOptions)(nil), "tensorflow.GPUOptions")
	proto.RegisterType((*OptimizerOptions)(nil), "tensorflow.OptimizerOptions")
	proto.RegisterType((*GraphOptions)(nil), "tensorflow.GraphOptions")
	proto.RegisterType((*ThreadPoolOptionProto)(nil), "tensorflow.ThreadPoolOptionProto")
	proto.RegisterType((*ConfigProto)(nil), "tensorflow.ConfigProto")
	proto.RegisterType((*DebugTensorWatch)(nil), "tensorflow.DebugTensorWatch")
	proto.RegisterType((*RunOptions)(nil), "tensorflow.RunOptions")
	proto.RegisterType((*RunMetadata)(nil), "tensorflow.RunMetadata")
	proto.RegisterEnum("tensorflow.OptimizerOptions_Level", OptimizerOptions_Level_name, OptimizerOptions_Level_value)
	proto.RegisterEnum("tensorflow.RunOptions_TraceLevel", RunOptions_TraceLevel_name, RunOptions_TraceLevel_value)
}

var fileDescriptor16 = []byte{
	// 1269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x56, 0x5f, 0x73, 0xdb, 0xc4,
	0x17, 0xfd, 0xc9, 0xae, 0xf3, 0xb3, 0xaf, 0xdd, 0xd4, 0x59, 0x9c, 0xd6, 0xb4, 0x65, 0xda, 0x9a,
	0xc9, 0x4c, 0x61, 0xc0, 0x2d, 0xe5, 0x5f, 0x87, 0x19, 0x60, 0x9a, 0xbf, 0x04, 0x92, 0xc6, 0xb3,
	0x76, 0xa7, 0x8f, 0x3b, 0xb2, 0xb5, 0x76, 0x34, 0x95, 0xb5, 0x1a, 0xed, 0x2a, 0x21, 0x3c, 0x33,
	0x3c, 0xf1, 0x0d, 0x78, 0xe7, 0x2b, 0xf0, 0xc4, 0xf7, 0xe2, 0x0d, 0xee, 0xde, 0x95, 0x2c, 0xc7,
	0x09, 0x25, 0x2f, 0x91, 0xee, 0x39, 0xf7, 0xfa, 0xee, 0xd9, 0xb3, 0x77, 0x05, 0x5b, 0x46, 0xc6,
	0x5a, 0xa5, 0xd3, 0x48, 0x9d, 0x3f, 0x99, 0xa8, 0x54, 0x3e, 0x49, 0x52, 0x65, 0xd4, 0x38, 0x9b,
	0xe2, 0x5b, 0x3c, 0x0d, 0x67, 0x7d, 0x7a, 0x67, 0x50, 0xd2, 0xee, 0x7e, 0xb8, 0x9a, 0x32, 0x4d,
	0xfd, 0xb9, 0x3c, 0x57, 0xe9, 0x1b, 0x7c, 0xd5, 0x46, 0xcc, 0x52, 0x3f, 0x39, 0x75, 0x79, 0x77,
	0xb7, 0xfe, 0x9d, 0xbb, 0x4c, 0x7b, 0x4b, 0x49, 0x6d, 0x64, 0x22, 0xb4, 0xf1, 0x8d, 0x76, 0xdc,
	0xde, 0xcf, 0x15, 0x80, 0x83, 0xc1, 0xab, 0x93, 0xc4, 0x84, 0x2a, 0xd6, 0x6c, 0x17, 0x1e, 0x24,
	0x32, 0x15, 0x88, 0x4d, 0xa4, 0xd6, 0x62, 0x96, 0x64, 0x62, 0x2e, 0xe7, 0x2a, 0xbd, 0x10, 0x58,
	0x63, 0x62, 0x39, 0x5d, 0xef, 0xa1, 0xf7, 0xd8, 0xe3, 0xf7, 0x90, 0x36, 0x70, 0xac, 0x83, 0x24,
	0x3b, 0x26, 0xce, 0x7e, 0x4e, 0x61, 0x5b, 0xb0, 0xee, 0x47, 0x91, 0x9a, 0xf8, 0x46, 0xa5, 0xc2,
	0x5c, 0x24, 0xb2, 0x5b, 0xc1, 0xa4, 0x06, 0xbf, 0xb9, 0x88, 0x8e, 0x30, 0xc8, 0xbe, 0x80, 0x3b,
	0x81, 0x9c, 0xca, 0x34, 0x95, 0x81, 0x08, 0x64, 0x24, 0x6d, 0xae, 0x18, 0x5f, 0x18, 0xa9, 0xbb,
	0x55, 0xe4, 0x57, 0xf9, 0x66, 0x01, 0xef, 0xe6, 0xe8, 0xb6, 0x05, 0xd9, 0x23, 0x68, 0xd9, 0x42,
	0xe7, 0xa8, 0x8d, 0x3a, 0x37, 0xa7, 0xdd, 0x1b, 0x48, 0xae, 0xf3, 0x26, 0xc5, 0x0e, 0x28, 0xc4,
	0xfa, 0xf0, 0xce, 0x59, 0xa8, 0xc3, 0x71, 0x24, 0xb1, 0xf2, 0x59, 0x38, 0x91, 0x22, 0x0a, 0xb5,
	0xe9, 0xd6, 0xa8, 0x8d, 0x8d, 0x1c, 0xda, 0x25, 0xe4, 0x08, 0x81, 0xde, 0xef, 0x15, 0x68, 0x5b,
	0x0d, 0xe6, 0xe1, 0x4f, 0x32, 0x2d, 0xc4, 0x38, 0x86, 0xf7, 0x03, 0x25, 0x26, 0x6a, 0x3e, 0xc7,
	0xbe, 0x74, 0x36, 0x96, 0x3f, 0x26, 0x29, 0xae, 0xd7, 0x76, 0x29, 0x23, 0xa4, 0xc6, 0xfe, 0x42,
	0x90, 0x3a, 0x7f, 0x18, 0xa8, 0x1d, 0x62, 0x0e, 0x97, 0x89, 0x7b, 0x25, 0xcf, 0xf6, 0x44, 0xe5,
	0x62, 0xd4, 0x3f, 0x36, 0x62, 0xaa, 0xa2, 0x20, 0x8c, 0x67, 0x24, 0x4d, 0x9d, 0x6f, 0xd8, 0x74,
	0x87, 0xec, 0x3b, 0x80, 0x3d, 0x85, 0x0e, 0xf2, 0xa7, 0x59, 0x4c, 0xa2, 0x8a, 0x30, 0x8e, 0xc2,
	0xd8, 0x26, 0xb8, 0xe5, 0xb2, 0x40, 0xed, 0xe7, 0xd0, 0x61, 0x8e, 0xb0, 0x6f, 0xa1, 0xa1, 0x12,
	0x23, 0x22, 0x79, 0x26, 0x23, 0x92, 0x70, 0xfd, 0x59, 0xaf, 0x5f, 0x9a, 0xa1, 0xbf, 0xba, 0xc2,
	0xfe, 0x91, 0x65, 0xf2, 0x3a, 0x26, 0xd1, 0x53, 0xef, 0x21, 0xd4, 0xe8, 0x81, 0xad, 0x41, 0xe5,
	0xe8, 0x93, 0xf6, 0xff, 0xd8, 0x2d, 0xfc, 0xff, 0xb4, 0xfd, 0x77, 0xf1, 0xe7, 0xf5, 0xfe, 0xa8,
	0x40, 0xeb, 0xc0, 0x7a, 0xad, 0x10, 0xe9, 0x33, 0xb8, 0x2d, 0x63, 0xdf, 0x0a, 0x9d, 0xca, 0xc9,
	0x99, 0xd0, 0x93, 0x53, 0x19, 0x64, 0x51, 0xb9, 0xb0, 0x8e, 0x43, 0x39, 0x82, 0xc3, 0x05, 0xc6,
	0x0e, 0x61, 0x43, 0x15, 0xcd, 0x08, 0xe5, 0x4a, 0x51, 0xc7, 0xcd, 0x67, 0xf7, 0xdf, 0xd6, 0x31,
	0x6f, 0xab, 0xd5, 0x5d, 0x7a, 0x0c, 0xed, 0x71, 0x16, 0x46, 0x81, 0xa0, 0xe3, 0x32, 0x57, 0xe8,
	0x24, 0x92, 0xa8, 0xca, 0xd7, 0x29, 0xbe, 0x83, 0xe1, 0x63, 0x1b, 0xb5, 0xbe, 0x09, 0x63, 0x34,
	0x94, 0xd0, 0xa7, 0x7e, 0x82, 0x26, 0xab, 0x39, 0xdf, 0x50, 0x6c, 0x48, 0x21, 0xf6, 0x11, 0xb0,
	0x24, 0xf2, 0xd1, 0x2e, 0x49, 0x9a, 0xc5, 0x68, 0x4b, 0x3a, 0x56, 0xdd, 0x35, 0x22, 0xb6, 0x09,
	0x19, 0x10, 0x40, 0x12, 0x7c, 0x7f, 0xa3, 0xee, 0xb5, 0x2b, 0x7c, 0x4b, 0xbf, 0x09, 0x93, 0xff,
	0xb4, 0x49, 0xef, 0x39, 0x6c, 0x8e, 0x4e, 0x53, 0xe9, 0x07, 0x03, 0xa5, 0x22, 0xd7, 0xfc, 0x80,
	0xa6, 0xc1, 0x03, 0x68, 0xc6, 0xd9, 0x5c, 0x18, 0x02, 0x35, 0xd9, 0xa9, 0xc6, 0x01, 0x43, 0x8e,
	0xae, 0x7b, 0xbf, 0xad, 0x41, 0x73, 0x87, 0xe6, 0x87, 0x4b, 0xf8, 0x01, 0x5a, 0xb9, 0xa9, 0x27,
	0x2a, 0x8b, 0x0d, 0x66, 0x54, 0x51, 0xb7, 0xc7, 0xcb, 0xba, 0x2d, 0xd1, 0xfb, 0xce, 0xe6, 0x3b,
	0x96, 0xba, 0x17, 0x9b, 0xf4, 0x82, 0x37, 0x83, 0x32, 0x82, 0x9e, 0xb9, 0x1f, 0x62, 0xd4, 0xc7,
	0x5d, 0x10, 0x89, 0x9f, 0xe2, 0x21, 0xc2, 0xa6, 0x75, 0xd9, 0x4e, 0x85, 0xda, 0x79, 0x97, 0x38,
	0x27, 0xc9, 0xa0, 0x64, 0xe4, 0xdd, 0xe5, 0x05, 0x68, 0x1b, 0xaf, 0x2d, 0x50, 0x5b, 0x14, 0xb0,
	0x7b, 0x76, 0x4d, 0x81, 0xcf, 0xe1, 0x4e, 0xa6, 0x51, 0x71, 0xbb, 0x31, 0xb9, 0x6e, 0x45, 0x6e,
	0xc3, 0x59, 0x08, 0xe1, 0x01, 0x6e, 0x91, 0x03, 0x8b, 0x34, 0x1f, 0xee, 0x17, 0xf4, 0xc5, 0xef,
	0xbb, 0x3c, 0x91, 0xa0, 0xc2, 0xdd, 0x16, 0xa9, 0xf2, 0x68, 0x59, 0x95, 0x6b, 0xf5, 0xe7, 0xdd,
	0xbc, 0xcc, 0xa1, 0xeb, 0xb0, 0x24, 0xb1, 0x0f, 0xc0, 0xed, 0xf9, 0x5c, 0xe2, 0x79, 0xc5, 0xfe,
	0x42, 0x15, 0x90, 0x49, 0x6b, 0xfc, 0xd6, 0x22, 0x3e, 0xa0, 0xb0, 0x1d, 0x79, 0xf9, 0x9e, 0x4c,
	0xc3, 0x08, 0xeb, 0x68, 0xf4, 0x60, 0xd5, 0x8e, 0x3c, 0x17, 0xdd, 0x77, 0x41, 0xf6, 0x25, 0x34,
	0xed, 0x4c, 0x2d, 0x1c, 0xbf, 0x46, 0x8e, 0xbf, 0xbd, 0xdc, 0x63, 0x39, 0x8c, 0x39, 0x20, 0xb5,
	0x70, 0x39, 0x0e, 0x03, 0x37, 0xf3, 0xb4, 0x9a, 0x62, 0x2f, 0xc5, 0xaf, 0x77, 0xff, 0xef, 0x86,
	0x01, 0x61, 0x43, 0x84, 0x06, 0x05, 0x62, 0x33, 0x22, 0x35, 0x2b, 0xc6, 0x5f, 0x99, 0x51, 0x77,
	0x19, 0x88, 0x39, 0x63, 0x94, 0x19, 0x5f, 0xc3, 0x4d, 0xf2, 0xfb, 0xa2, 0x3d, 0xa0, 0xf6, 0xba,
	0x97, 0xda, 0x5b, 0x3a, 0xfb, 0xbc, 0x35, 0x5b, 0x9e, 0x04, 0xb8, 0x8f, 0x0a, 0x45, 0x22, 0xb7,
	0x0b, 0x3c, 0xa4, 0x52, 0x65, 0x06, 0xb7, 0x46, 0xcc, 0x75, 0xb7, 0x49, 0xe7, 0xb1, 0xb3, 0x80,
	0x47, 0x0e, 0x3d, 0x8c, 0x8f, 0xf5, 0xdd, 0x6f, 0xa0, 0xbd, 0xea, 0x50, 0xd6, 0x86, 0xea, 0x1b,
	0x79, 0x41, 0x47, 0xa1, 0xc1, 0xed, 0x23, 0xeb, 0x40, 0xed, 0xcc, 0x8f, 0x32, 0x99, 0xfb, 0xd1,
	0xbd, 0x7c, 0x55, 0x79, 0xee, 0xf5, 0x7e, 0xf5, 0x6c, 0x81, 0x71, 0x36, 0x1b, 0x51, 0x97, 0xaf,
	0x7d, 0x33, 0x39, 0x65, 0xf7, 0xa0, 0x11, 0xe3, 0x99, 0x17, 0x31, 0x5e, 0x7c, 0x79, 0x99, 0xba,
	0x0d, 0xbc, 0xc4, 0x77, 0x7b, 0xe0, 0xf0, 0xc7, 0x13, 0xec, 0x4e, 0x47, 0xca, 0xe4, 0x15, 0xc1,
	0x85, 0x86, 0x18, 0xb1, 0xd9, 0x81, 0xad, 0x88, 0x42, 0xd8, 0xa9, 0x64, 0xf7, 0xb1, 0x4e, 0x81,
	0x93, 0x44, 0xb3, 0xf7, 0x00, 0x1c, 0x98, 0xa5, 0x51, 0xb1, 0xcb, 0x8e, 0xfe, 0x0a, 0x03, 0xbd,
	0x5f, 0xaa, 0x00, 0x3c, 0x8b, 0x0b, 0x51, 0xb6, 0xa1, 0x89, 0x07, 0xc7, 0xde, 0x3f, 0x34, 0x94,
	0x3d, 0x1a, 0xca, 0x97, 0x4c, 0x59, 0x92, 0xfb, 0x23, 0xcb, 0x74, 0x33, 0x19, 0xcc, 0xe2, 0x99,
	0xf5, 0xe0, 0xe6, 0x65, 0x39, 0x2b, 0x24, 0x67, 0xd3, 0x94, 0x2a, 0xb2, 0x27, 0xd0, 0xb9, 0xf6,
	0x14, 0x38, 0xbb, 0x6e, 0x84, 0x57, 0xbc, 0x3d, 0xb4, 0x97, 0xaf, 0x5d, 0x86, 0x6b, 0x45, 0x9c,
	0x5b, 0xdd, 0xec, 0xce, 0xbb, 0x35, 0xad, 0xcc, 0xe1, 0x55, 0x81, 0x79, 0x27, 0x58, 0x89, 0xe0,
	0x1a, 0xb4, 0xbd, 0xd1, 0x73, 0x65, 0x71, 0x12, 0x98, 0x90, 0x9c, 0x40, 0x1e, 0x29, 0x86, 0xed,
	0xa6, 0x83, 0x07, 0x05, 0x4a, 0x76, 0xd2, 0x3d, 0x0e, 0x50, 0xae, 0x9d, 0xb5, 0xa0, 0xfe, 0xf2,
	0x44, 0x8c, 0xf8, 0x8b, 0x9d, 0x3d, 0xbc, 0x82, 0x18, 0xac, 0x0f, 0x4f, 0xf6, 0x47, 0xaf, 0x5f,
	0xf0, 0xbd, 0x3c, 0xe6, 0xd9, 0xd8, 0x77, 0x2f, 0xf8, 0xee, 0x52, 0xac, 0xc2, 0xd6, 0x01, 0xf6,
	0x5f, 0x1d, 0x1d, 0xe5, 0xef, 0xd5, 0xde, 0x9f, 0x1e, 0x34, 0x51, 0xdb, 0x63, 0x69, 0xfc, 0xc0,
	0x37, 0x3e, 0x5e, 0x54, 0x50, 0x7e, 0xfd, 0xd0, 0x46, 0x34, 0x9f, 0x6d, 0x2e, 0xaf, 0x71, 0x88,
	0xe8, 0xd0, 0x82, 0xbc, 0xa1, 0x8b, 0x47, 0x3c, 0xb0, 0x50, 0x7e, 0x86, 0x91, 0xf0, 0x2b, 0x07,
	0xc2, 0x5e, 0x2f, 0xb4, 0x8a, 0x5d, 0x39, 0xe5, 0x8d, 0x49, 0xf1, 0x86, 0x63, 0xb1, 0x7d, 0x45,
	0x83, 0x2a, 0x09, 0xdb, 0xb9, 0x72, 0x9e, 0x6c, 0xea, 0xad, 0xe4, 0xb2, 0x26, 0xdb, 0x1f, 0x43,
	0x57, 0xa5, 0xb3, 0x65, 0xee, 0xe2, 0x33, 0x6e, 0xbb, 0xb5, 0x34, 0xdf, 0xf5, 0xc0, 0xfb, 0xcb,
	0xf3, 0xc6, 0x6b, 0xf4, 0x3d, 0xf7, 0xe9, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xbb, 0x15,
	0x96, 0x83, 0x0a, 0x00, 0x00,
}
