// Code generated by protoc-gen-go.
// source: task_worker.proto
// DO NOT EDIT!

/*
Package ga4gh_task_ref is a generated protocol buffer package.

It is generated from these files:
	task_worker.proto

It has these top-level messages:
	Worker
	Assignment
	UpdateWorkerRequest
	UpdateWorkerResponse
	UpdateJobLogsRequest
	UpdateJobLogsResponse
	QueuedTaskInfoRequest
	QueuedTaskInfo
*/
package ga4gh_task_ref

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ga4gh_task_exec "tes/ga4gh"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Worker struct {
	Id        string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Active    map[string]bool            `protobuf:"bytes,2,rep,name=active" json:"active,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Assigned  map[string]bool            `protobuf:"bytes,4,rep,name=assigned" json:"assigned,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Resources *ga4gh_task_exec.Resources `protobuf:"bytes,5,opt,name=resources" json:"resources,omitempty"`
	LastPing  int64                      `protobuf:"varint,6,opt,name=last_ping,json=lastPing" json:"last_ping,omitempty"`
}

func (m *Worker) Reset()                    { *m = Worker{} }
func (m *Worker) String() string            { return proto.CompactTextString(m) }
func (*Worker) ProtoMessage()               {}
func (*Worker) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Worker) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Worker) GetActive() map[string]bool {
	if m != nil {
		return m.Active
	}
	return nil
}

func (m *Worker) GetAssigned() map[string]bool {
	if m != nil {
		return m.Assigned
	}
	return nil
}

func (m *Worker) GetResources() *ga4gh_task_exec.Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Worker) GetLastPing() int64 {
	if m != nil {
		return m.LastPing
	}
	return 0
}

type Assignment struct {
	Job  *ga4gh_task_exec.Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	Auth string               `protobuf:"bytes,2,opt,name=auth" json:"auth,omitempty"`
}

func (m *Assignment) Reset()                    { *m = Assignment{} }
func (m *Assignment) String() string            { return proto.CompactTextString(m) }
func (*Assignment) ProtoMessage()               {}
func (*Assignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Assignment) GetJob() *ga4gh_task_exec.Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *Assignment) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

type UpdateWorkerRequest struct {
	// ID of the worker
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Description of CPU, RAM, etc.
	Resources *ga4gh_task_exec.Resources `protobuf:"bytes,2,opt,name=resources" json:"resources,omitempty"`
	// Hostname of the worker host.
	Hostname string `protobuf:"bytes,3,opt,name=hostname" json:"hostname,omitempty"`
	// States of all jobs in the worker: job ID -> State
	States map[string]ga4gh_task_exec.State `protobuf:"bytes,4,rep,name=states" json:"states,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=ga4gh_task_exec.State"`
}

func (m *UpdateWorkerRequest) Reset()                    { *m = UpdateWorkerRequest{} }
func (m *UpdateWorkerRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateWorkerRequest) ProtoMessage()               {}
func (*UpdateWorkerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateWorkerRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateWorkerRequest) GetResources() *ga4gh_task_exec.Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *UpdateWorkerRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *UpdateWorkerRequest) GetStates() map[string]ga4gh_task_exec.State {
	if m != nil {
		return m.States
	}
	return nil
}

type UpdateWorkerResponse struct {
	// New jobs which have been assigned to this worker.
	Assigned []*Assignment `protobuf:"bytes,3,rep,name=assigned" json:"assigned,omitempty"`
	// IDs of jobs assigned to this worker that were canceled.
	Canceled []string `protobuf:"bytes,4,rep,name=canceled" json:"canceled,omitempty"`
}

func (m *UpdateWorkerResponse) Reset()                    { *m = UpdateWorkerResponse{} }
func (m *UpdateWorkerResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateWorkerResponse) ProtoMessage()               {}
func (*UpdateWorkerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateWorkerResponse) GetAssigned() []*Assignment {
	if m != nil {
		return m.Assigned
	}
	return nil
}

func (m *UpdateWorkerResponse) GetCanceled() []string {
	if m != nil {
		return m.Canceled
	}
	return nil
}

type UpdateJobLogsRequest struct {
	Id       string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Step     int64                   `protobuf:"varint,2,opt,name=step" json:"step,omitempty"`
	Log      *ga4gh_task_exec.JobLog `protobuf:"bytes,4,opt,name=log" json:"log,omitempty"`
	WorkerId string                  `protobuf:"bytes,5,opt,name=worker_id,json=workerId" json:"worker_id,omitempty"`
}

func (m *UpdateJobLogsRequest) Reset()                    { *m = UpdateJobLogsRequest{} }
func (m *UpdateJobLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateJobLogsRequest) ProtoMessage()               {}
func (*UpdateJobLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateJobLogsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateJobLogsRequest) GetStep() int64 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *UpdateJobLogsRequest) GetLog() *ga4gh_task_exec.JobLog {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *UpdateJobLogsRequest) GetWorkerId() string {
	if m != nil {
		return m.WorkerId
	}
	return ""
}

type UpdateJobLogsResponse struct {
}

func (m *UpdateJobLogsResponse) Reset()                    { *m = UpdateJobLogsResponse{} }
func (m *UpdateJobLogsResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateJobLogsResponse) ProtoMessage()               {}
func (*UpdateJobLogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type QueuedTaskInfoRequest struct {
	MaxTasks int32 `protobuf:"varint,1,opt,name=max_tasks,json=maxTasks" json:"max_tasks,omitempty"`
}

func (m *QueuedTaskInfoRequest) Reset()                    { *m = QueuedTaskInfoRequest{} }
func (m *QueuedTaskInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*QueuedTaskInfoRequest) ProtoMessage()               {}
func (*QueuedTaskInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *QueuedTaskInfoRequest) GetMaxTasks() int32 {
	if m != nil {
		return m.MaxTasks
	}
	return 0
}

type QueuedTaskInfo struct {
	Inputs    []string                   `protobuf:"bytes,1,rep,name=inputs" json:"inputs,omitempty"`
	Resources *ga4gh_task_exec.Resources `protobuf:"bytes,2,opt,name=resources" json:"resources,omitempty"`
}

func (m *QueuedTaskInfo) Reset()                    { *m = QueuedTaskInfo{} }
func (m *QueuedTaskInfo) String() string            { return proto.CompactTextString(m) }
func (*QueuedTaskInfo) ProtoMessage()               {}
func (*QueuedTaskInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *QueuedTaskInfo) GetInputs() []string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *QueuedTaskInfo) GetResources() *ga4gh_task_exec.Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func init() {
	proto.RegisterType((*Worker)(nil), "ga4gh_task_ref.Worker")
	proto.RegisterType((*Assignment)(nil), "ga4gh_task_ref.Assignment")
	proto.RegisterType((*UpdateWorkerRequest)(nil), "ga4gh_task_ref.UpdateWorkerRequest")
	proto.RegisterType((*UpdateWorkerResponse)(nil), "ga4gh_task_ref.UpdateWorkerResponse")
	proto.RegisterType((*UpdateJobLogsRequest)(nil), "ga4gh_task_ref.UpdateJobLogsRequest")
	proto.RegisterType((*UpdateJobLogsResponse)(nil), "ga4gh_task_ref.UpdateJobLogsResponse")
	proto.RegisterType((*QueuedTaskInfoRequest)(nil), "ga4gh_task_ref.QueuedTaskInfoRequest")
	proto.RegisterType((*QueuedTaskInfo)(nil), "ga4gh_task_ref.QueuedTaskInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Scheduler service

type SchedulerClient interface {
	GetQueueInfo(ctx context.Context, in *QueuedTaskInfoRequest, opts ...grpc.CallOption) (Scheduler_GetQueueInfoClient, error)
	UpdateJobLogs(ctx context.Context, in *UpdateJobLogsRequest, opts ...grpc.CallOption) (*UpdateJobLogsResponse, error)
	UpdateWorker(ctx context.Context, in *UpdateWorkerRequest, opts ...grpc.CallOption) (*UpdateWorkerResponse, error)
}

type schedulerClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerClient(cc *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) GetQueueInfo(ctx context.Context, in *QueuedTaskInfoRequest, opts ...grpc.CallOption) (Scheduler_GetQueueInfoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Scheduler_serviceDesc.Streams[0], c.cc, "/ga4gh_task_ref.Scheduler/GetQueueInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &schedulerGetQueueInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Scheduler_GetQueueInfoClient interface {
	Recv() (*QueuedTaskInfo, error)
	grpc.ClientStream
}

type schedulerGetQueueInfoClient struct {
	grpc.ClientStream
}

func (x *schedulerGetQueueInfoClient) Recv() (*QueuedTaskInfo, error) {
	m := new(QueuedTaskInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *schedulerClient) UpdateJobLogs(ctx context.Context, in *UpdateJobLogsRequest, opts ...grpc.CallOption) (*UpdateJobLogsResponse, error) {
	out := new(UpdateJobLogsResponse)
	err := grpc.Invoke(ctx, "/ga4gh_task_ref.Scheduler/UpdateJobLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) UpdateWorker(ctx context.Context, in *UpdateWorkerRequest, opts ...grpc.CallOption) (*UpdateWorkerResponse, error) {
	out := new(UpdateWorkerResponse)
	err := grpc.Invoke(ctx, "/ga4gh_task_ref.Scheduler/UpdateWorker", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Scheduler service

type SchedulerServer interface {
	GetQueueInfo(*QueuedTaskInfoRequest, Scheduler_GetQueueInfoServer) error
	UpdateJobLogs(context.Context, *UpdateJobLogsRequest) (*UpdateJobLogsResponse, error)
	UpdateWorker(context.Context, *UpdateWorkerRequest) (*UpdateWorkerResponse, error)
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_GetQueueInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueuedTaskInfoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SchedulerServer).GetQueueInfo(m, &schedulerGetQueueInfoServer{stream})
}

type Scheduler_GetQueueInfoServer interface {
	Send(*QueuedTaskInfo) error
	grpc.ServerStream
}

type schedulerGetQueueInfoServer struct {
	grpc.ServerStream
}

func (x *schedulerGetQueueInfoServer) Send(m *QueuedTaskInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _Scheduler_UpdateJobLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJobLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).UpdateJobLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_ref.Scheduler/UpdateJobLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).UpdateJobLogs(ctx, req.(*UpdateJobLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_UpdateWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).UpdateWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ga4gh_task_ref.Scheduler/UpdateWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).UpdateWorker(ctx, req.(*UpdateWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ga4gh_task_ref.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateJobLogs",
			Handler:    _Scheduler_UpdateJobLogs_Handler,
		},
		{
			MethodName: "UpdateWorker",
			Handler:    _Scheduler_UpdateWorker_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetQueueInfo",
			Handler:       _Scheduler_GetQueueInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "task_worker.proto",
}

func init() { proto.RegisterFile("task_worker.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 617 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xd1, 0x6e, 0xd3, 0x4a,
	0x10, 0xad, 0xed, 0xc4, 0xb2, 0x27, 0x6d, 0x74, 0xef, 0x92, 0xb6, 0x96, 0x2b, 0xa1, 0xc8, 0xa4,
	0x28, 0x48, 0xc8, 0xa0, 0x50, 0xa1, 0x52, 0x5e, 0xe0, 0x01, 0x95, 0x56, 0x7d, 0xa0, 0x5b, 0x10,
	0x42, 0x48, 0x44, 0x1b, 0x7b, 0xeb, 0xb8, 0x49, 0xbc, 0xc6, 0xbb, 0x2e, 0xe9, 0x2b, 0x3f, 0xc1,
	0x07, 0xf0, 0x6d, 0xfc, 0x07, 0xf2, 0xda, 0x49, 0xed, 0x24, 0x2d, 0x94, 0xb7, 0x9d, 0xf5, 0x39,
	0x33, 0x67, 0xe6, 0xcc, 0x1a, 0xfe, 0x17, 0x84, 0x8f, 0xfa, 0xdf, 0x58, 0x32, 0xa2, 0x89, 0x1b,
	0x27, 0x4c, 0x30, 0xd4, 0x0c, 0xc8, 0x5e, 0x30, 0xec, 0xcb, 0x0f, 0x09, 0x3d, 0xb7, 0x5b, 0xf2,
	0x44, 0xa7, 0xd4, 0x4b, 0x45, 0xc8, 0xa2, 0x1c, 0xe5, 0xfc, 0x52, 0x41, 0xff, 0x28, 0x69, 0xa8,
	0x09, 0x6a, 0xe8, 0x5b, 0x4a, 0x5b, 0xe9, 0x9a, 0x58, 0x0d, 0x7d, 0x74, 0x00, 0x3a, 0xf1, 0x44,
	0x78, 0x49, 0x2d, 0xb5, 0xad, 0x75, 0x1b, 0x3d, 0xc7, 0xad, 0x66, 0x74, 0x73, 0x9e, 0xfb, 0x5a,
	0x82, 0xde, 0x44, 0x22, 0xb9, 0xc2, 0x05, 0x03, 0xbd, 0x02, 0x83, 0x70, 0x1e, 0x06, 0x11, 0xf5,
	0xad, 0x9a, 0x64, 0x77, 0x6e, 0x62, 0x17, 0xb0, 0x9c, 0x3f, 0x67, 0xa1, 0x7d, 0x30, 0x13, 0xca,
	0x59, 0x9a, 0x78, 0x94, 0x5b, 0xf5, 0xb6, 0xd2, 0x6d, 0xf4, 0xec, 0x72, 0x8a, 0xac, 0x11, 0x17,
	0xcf, 0x10, 0xf8, 0x1a, 0x8c, 0x76, 0xc0, 0x1c, 0x13, 0x2e, 0xfa, 0x71, 0x18, 0x05, 0x96, 0xde,
	0x56, 0xba, 0x1a, 0x36, 0xb2, 0x8b, 0x77, 0x61, 0x14, 0xd8, 0x2f, 0xa0, 0x51, 0xd2, 0x8b, 0xfe,
	0x03, 0x6d, 0x44, 0xaf, 0x8a, 0xa6, 0xb3, 0x23, 0x6a, 0x41, 0xfd, 0x92, 0x8c, 0xd3, 0xac, 0x69,
	0xa5, 0x6b, 0xe0, 0x3c, 0x38, 0x50, 0xf7, 0x15, 0xfb, 0x25, 0x6c, 0x54, 0xc4, 0xde, 0x85, 0xec,
	0xbc, 0x05, 0xc8, 0xc9, 0x13, 0x1a, 0x09, 0xf4, 0x10, 0xb4, 0x0b, 0x36, 0x90, 0xcc, 0x46, 0xaf,
	0xb5, 0xd4, 0xd6, 0x31, 0x1b, 0xe0, 0x0c, 0x80, 0x10, 0xd4, 0x48, 0x2a, 0x86, 0x32, 0x9d, 0x89,
	0xe5, 0xd9, 0xf9, 0xa1, 0xc2, 0xbd, 0x0f, 0xb1, 0x4f, 0x04, 0xcd, 0x27, 0x88, 0xe9, 0xd7, 0x94,
	0x72, 0xb1, 0x64, 0x5f, 0x65, 0x80, 0xea, 0x5d, 0x06, 0x68, 0x83, 0x31, 0x64, 0x5c, 0x44, 0x64,
	0x42, 0x2d, 0x4d, 0xe6, 0x9b, 0xc7, 0xe8, 0x10, 0x74, 0x2e, 0x88, 0xa0, 0xbc, 0xb0, 0xf5, 0xc9,
	0xa2, 0xad, 0x2b, 0xa4, 0xb9, 0x67, 0x92, 0x51, 0x6c, 0x48, 0x4e, 0xb7, 0x4f, 0xa1, 0x51, 0xba,
	0x5e, 0x31, 0xcb, 0xc7, 0xe5, 0x59, 0x36, 0x7b, 0x5b, 0x4b, 0xda, 0x25, 0xbd, 0x3c, 0xe3, 0x0b,
	0x68, 0x55, 0xab, 0xf3, 0x98, 0x45, 0x9c, 0xa2, 0xe7, 0xa5, 0x65, 0xd4, 0xa4, 0x6a, 0x7b, 0x51,
	0xf5, 0xb5, 0x37, 0xa5, 0x15, 0xb4, 0xc1, 0xf0, 0x48, 0xe4, 0xd1, 0x71, 0xb1, 0xc4, 0x26, 0x9e,
	0xc7, 0xce, 0x77, 0x65, 0x56, 0xec, 0x98, 0x0d, 0x4e, 0x58, 0xc0, 0x6f, 0xb2, 0x01, 0x41, 0x8d,
	0x0b, 0x1a, 0xcb, 0x2e, 0x34, 0x2c, 0xcf, 0xe8, 0x11, 0x68, 0x63, 0x16, 0x58, 0x35, 0x69, 0xca,
	0xf6, 0x2a, 0xfb, 0x4f, 0x58, 0x80, 0x33, 0x4c, 0xb6, 0xcc, 0xf9, 0xab, 0xee, 0x87, 0xbe, 0x7c,
	0x06, 0x26, 0x36, 0xf2, 0x8b, 0x23, 0xdf, 0xd9, 0x86, 0xcd, 0x05, 0x0d, 0x79, 0xc7, 0xce, 0x1e,
	0x6c, 0x9e, 0xa6, 0x34, 0xa5, 0xfe, 0x7b, 0xc2, 0x47, 0x47, 0xd1, 0x39, 0x9b, 0xa9, 0xdb, 0x01,
	0x73, 0x42, 0xa6, 0xb2, 0x16, 0x97, 0x22, 0xeb, 0xd8, 0x98, 0x90, 0x69, 0x06, 0xe3, 0xce, 0x00,
	0x9a, 0x55, 0x16, 0xda, 0x02, 0x3d, 0x8c, 0xe2, 0x54, 0x64, 0xd8, 0xac, 0xff, 0x22, 0xfa, 0xf7,
	0xdd, 0xea, 0xfd, 0x54, 0xc1, 0x3c, 0xf3, 0x86, 0xd4, 0x4f, 0xc7, 0x34, 0x41, 0x9f, 0x60, 0xfd,
	0x90, 0x0a, 0x59, 0x54, 0xd6, 0xdb, 0x5d, 0xf4, 0x65, 0x65, 0x17, 0xf6, 0xfd, 0xdb, 0x61, 0xce,
	0xda, 0x53, 0x05, 0x7d, 0x81, 0x8d, 0xca, 0x6c, 0x50, 0x67, 0xf5, 0xa6, 0x56, 0xed, 0xb3, 0x77,
	0xff, 0x80, 0x2a, 0x06, 0xbc, 0x86, 0x3e, 0xc3, 0x7a, 0x79, 0xd9, 0xd0, 0x83, 0xbf, 0x78, 0x08,
	0x76, 0xe7, 0x76, 0xd0, 0x2c, 0xf9, 0x40, 0x97, 0x3f, 0xe7, 0x67, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xd5, 0x08, 0xa3, 0xe5, 0xd7, 0x05, 0x00, 0x00,
}
