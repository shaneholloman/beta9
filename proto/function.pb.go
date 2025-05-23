// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: function.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FunctionInvokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StubId   string `protobuf:"bytes,1,opt,name=stub_id,json=stubId,proto3" json:"stub_id,omitempty"`
	Args     []byte `protobuf:"bytes,2,opt,name=args,proto3" json:"args,omitempty"`
	Headless bool   `protobuf:"varint,3,opt,name=headless,proto3" json:"headless,omitempty"`
}

func (x *FunctionInvokeRequest) Reset() {
	*x = FunctionInvokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionInvokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionInvokeRequest) ProtoMessage() {}

func (x *FunctionInvokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionInvokeRequest.ProtoReflect.Descriptor instead.
func (*FunctionInvokeRequest) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{0}
}

func (x *FunctionInvokeRequest) GetStubId() string {
	if x != nil {
		return x.StubId
	}
	return ""
}

func (x *FunctionInvokeRequest) GetArgs() []byte {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *FunctionInvokeRequest) GetHeadless() bool {
	if x != nil {
		return x.Headless
	}
	return false
}

type FunctionInvokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId   string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Output   string `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Done     bool   `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	ExitCode int32  `protobuf:"varint,4,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	Result   []byte `protobuf:"bytes,5,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *FunctionInvokeResponse) Reset() {
	*x = FunctionInvokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionInvokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionInvokeResponse) ProtoMessage() {}

func (x *FunctionInvokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionInvokeResponse.ProtoReflect.Descriptor instead.
func (*FunctionInvokeResponse) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{1}
}

func (x *FunctionInvokeResponse) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *FunctionInvokeResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *FunctionInvokeResponse) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *FunctionInvokeResponse) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *FunctionInvokeResponse) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

type FunctionGetArgsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *FunctionGetArgsRequest) Reset() {
	*x = FunctionGetArgsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionGetArgsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionGetArgsRequest) ProtoMessage() {}

func (x *FunctionGetArgsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionGetArgsRequest.ProtoReflect.Descriptor instead.
func (*FunctionGetArgsRequest) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{2}
}

func (x *FunctionGetArgsRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type FunctionGetArgsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok   bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Args []byte `protobuf:"bytes,2,opt,name=args,proto3" json:"args,omitempty"`
}

func (x *FunctionGetArgsResponse) Reset() {
	*x = FunctionGetArgsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionGetArgsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionGetArgsResponse) ProtoMessage() {}

func (x *FunctionGetArgsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionGetArgsResponse.ProtoReflect.Descriptor instead.
func (*FunctionGetArgsResponse) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{3}
}

func (x *FunctionGetArgsResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *FunctionGetArgsResponse) GetArgs() []byte {
	if x != nil {
		return x.Args
	}
	return nil
}

type FunctionSetResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Result []byte `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *FunctionSetResultRequest) Reset() {
	*x = FunctionSetResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionSetResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionSetResultRequest) ProtoMessage() {}

func (x *FunctionSetResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionSetResultRequest.ProtoReflect.Descriptor instead.
func (*FunctionSetResultRequest) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{4}
}

func (x *FunctionSetResultRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *FunctionSetResultRequest) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

type FunctionSetResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *FunctionSetResultResponse) Reset() {
	*x = FunctionSetResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionSetResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionSetResultResponse) ProtoMessage() {}

func (x *FunctionSetResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionSetResultResponse.ProtoReflect.Descriptor instead.
func (*FunctionSetResultResponse) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{5}
}

func (x *FunctionSetResultResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type FunctionMonitorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId      string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	StubId      string `protobuf:"bytes,2,opt,name=stub_id,json=stubId,proto3" json:"stub_id,omitempty"`
	ContainerId string `protobuf:"bytes,3,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
}

func (x *FunctionMonitorRequest) Reset() {
	*x = FunctionMonitorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionMonitorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionMonitorRequest) ProtoMessage() {}

func (x *FunctionMonitorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionMonitorRequest.ProtoReflect.Descriptor instead.
func (*FunctionMonitorRequest) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{6}
}

func (x *FunctionMonitorRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *FunctionMonitorRequest) GetStubId() string {
	if x != nil {
		return x.StubId
	}
	return ""
}

func (x *FunctionMonitorRequest) GetContainerId() string {
	if x != nil {
		return x.ContainerId
	}
	return ""
}

type FunctionMonitorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok        bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Cancelled bool `protobuf:"varint,2,opt,name=cancelled,proto3" json:"cancelled,omitempty"`
	Complete  bool `protobuf:"varint,3,opt,name=complete,proto3" json:"complete,omitempty"`
	TimedOut  bool `protobuf:"varint,4,opt,name=timed_out,json=timedOut,proto3" json:"timed_out,omitempty"`
}

func (x *FunctionMonitorResponse) Reset() {
	*x = FunctionMonitorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionMonitorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionMonitorResponse) ProtoMessage() {}

func (x *FunctionMonitorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionMonitorResponse.ProtoReflect.Descriptor instead.
func (*FunctionMonitorResponse) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{7}
}

func (x *FunctionMonitorResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *FunctionMonitorResponse) GetCancelled() bool {
	if x != nil {
		return x.Cancelled
	}
	return false
}

func (x *FunctionMonitorResponse) GetComplete() bool {
	if x != nil {
		return x.Complete
	}
	return false
}

func (x *FunctionMonitorResponse) GetTimedOut() bool {
	if x != nil {
		return x.TimedOut
	}
	return false
}

type FunctionScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StubId       string `protobuf:"bytes,1,opt,name=stub_id,json=stubId,proto3" json:"stub_id,omitempty"`
	When         string `protobuf:"bytes,2,opt,name=when,proto3" json:"when,omitempty"`
	DeploymentId string `protobuf:"bytes,3,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
}

func (x *FunctionScheduleRequest) Reset() {
	*x = FunctionScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionScheduleRequest) ProtoMessage() {}

func (x *FunctionScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionScheduleRequest.ProtoReflect.Descriptor instead.
func (*FunctionScheduleRequest) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{8}
}

func (x *FunctionScheduleRequest) GetStubId() string {
	if x != nil {
		return x.StubId
	}
	return ""
}

func (x *FunctionScheduleRequest) GetWhen() string {
	if x != nil {
		return x.When
	}
	return ""
}

func (x *FunctionScheduleRequest) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

type FunctionScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok             bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	ErrMsg         string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	ScheduledJobId string `protobuf:"bytes,3,opt,name=scheduled_job_id,json=scheduledJobId,proto3" json:"scheduled_job_id,omitempty"`
}

func (x *FunctionScheduleResponse) Reset() {
	*x = FunctionScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_function_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionScheduleResponse) ProtoMessage() {}

func (x *FunctionScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_function_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionScheduleResponse.ProtoReflect.Descriptor instead.
func (*FunctionScheduleResponse) Descriptor() ([]byte, []int) {
	return file_function_proto_rawDescGZIP(), []int{9}
}

func (x *FunctionScheduleResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *FunctionScheduleResponse) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *FunctionScheduleResponse) GetScheduledJobId() string {
	if x != nil {
		return x.ScheduledJobId
	}
	return ""
}

var File_function_proto protoreflect.FileDescriptor

var file_function_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x60, 0x0a, 0x15, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x75, 0x62, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x22, 0x92, 0x01, 0x0a,
	0x16, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x31, 0x0a, 0x16, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74,
	0x41, 0x72, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x17, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x47, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x22, 0x4b, 0x0a, 0x18, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x2b, 0x0a, 0x19, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x6d, 0x0a,
	0x16, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x75, 0x62, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x80, 0x01, 0x0a,
	0x17, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x22,
	0x6b, 0x0a, 0x17, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x74,
	0x75, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x75,
	0x62, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x6d, 0x0a, 0x18,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73,
	0x67, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x6a,
	0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x32, 0xdb, 0x03, 0x0a, 0x0f,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x57, 0x0a, 0x0e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x12, 0x1f, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x58, 0x0a, 0x0f, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x20, 0x2e, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5e, 0x0a, 0x11, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x58, 0x0a, 0x0f, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x5b, 0x0a, 0x10,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x12, 0x21, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x61, 0x6d, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x39, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_function_proto_rawDescOnce sync.Once
	file_function_proto_rawDescData = file_function_proto_rawDesc
)

func file_function_proto_rawDescGZIP() []byte {
	file_function_proto_rawDescOnce.Do(func() {
		file_function_proto_rawDescData = protoimpl.X.CompressGZIP(file_function_proto_rawDescData)
	})
	return file_function_proto_rawDescData
}

var file_function_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_function_proto_goTypes = []interface{}{
	(*FunctionInvokeRequest)(nil),     // 0: function.FunctionInvokeRequest
	(*FunctionInvokeResponse)(nil),    // 1: function.FunctionInvokeResponse
	(*FunctionGetArgsRequest)(nil),    // 2: function.FunctionGetArgsRequest
	(*FunctionGetArgsResponse)(nil),   // 3: function.FunctionGetArgsResponse
	(*FunctionSetResultRequest)(nil),  // 4: function.FunctionSetResultRequest
	(*FunctionSetResultResponse)(nil), // 5: function.FunctionSetResultResponse
	(*FunctionMonitorRequest)(nil),    // 6: function.FunctionMonitorRequest
	(*FunctionMonitorResponse)(nil),   // 7: function.FunctionMonitorResponse
	(*FunctionScheduleRequest)(nil),   // 8: function.FunctionScheduleRequest
	(*FunctionScheduleResponse)(nil),  // 9: function.FunctionScheduleResponse
}
var file_function_proto_depIdxs = []int32{
	0, // 0: function.FunctionService.FunctionInvoke:input_type -> function.FunctionInvokeRequest
	2, // 1: function.FunctionService.FunctionGetArgs:input_type -> function.FunctionGetArgsRequest
	4, // 2: function.FunctionService.FunctionSetResult:input_type -> function.FunctionSetResultRequest
	6, // 3: function.FunctionService.FunctionMonitor:input_type -> function.FunctionMonitorRequest
	8, // 4: function.FunctionService.FunctionSchedule:input_type -> function.FunctionScheduleRequest
	1, // 5: function.FunctionService.FunctionInvoke:output_type -> function.FunctionInvokeResponse
	3, // 6: function.FunctionService.FunctionGetArgs:output_type -> function.FunctionGetArgsResponse
	5, // 7: function.FunctionService.FunctionSetResult:output_type -> function.FunctionSetResultResponse
	7, // 8: function.FunctionService.FunctionMonitor:output_type -> function.FunctionMonitorResponse
	9, // 9: function.FunctionService.FunctionSchedule:output_type -> function.FunctionScheduleResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_function_proto_init() }
func file_function_proto_init() {
	if File_function_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_function_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionInvokeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_function_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionInvokeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_function_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionGetArgsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_function_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionGetArgsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_function_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionSetResultRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_function_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionSetResultResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_function_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionMonitorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_function_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionMonitorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_function_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionScheduleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_function_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionScheduleResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_function_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_function_proto_goTypes,
		DependencyIndexes: file_function_proto_depIdxs,
		MessageInfos:      file_function_proto_msgTypes,
	}.Build()
	File_function_proto = out.File
	file_function_proto_rawDesc = nil
	file_function_proto_goTypes = nil
	file_function_proto_depIdxs = nil
}
