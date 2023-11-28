// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: advent.proto

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

type Solutions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Solutions []*Solution `protobuf:"bytes,1,rep,name=solutions,proto3" json:"solutions,omitempty"`
}

func (x *Solutions) Reset() {
	*x = Solutions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Solutions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Solutions) ProtoMessage() {}

func (x *Solutions) ProtoReflect() protoreflect.Message {
	mi := &file_advent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Solutions.ProtoReflect.Descriptor instead.
func (*Solutions) Descriptor() ([]byte, []int) {
	return file_advent_proto_rawDescGZIP(), []int{0}
}

func (x *Solutions) GetSolutions() []*Solution {
	if x != nil {
		return x.Solutions
	}
	return nil
}

type Solution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year         int32  `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Day          int32  `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	Part         int32  `protobuf:"varint,3,opt,name=part,proto3" json:"part,omitempty"`
	BigAnswer    int64  `protobuf:"varint,4,opt,name=big_answer,json=bigAnswer,proto3" json:"big_answer,omitempty"`
	StringAnswer string `protobuf:"bytes,5,opt,name=string_answer,json=stringAnswer,proto3" json:"string_answer,omitempty"`
	Answer       int32  `protobuf:"varint,6,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *Solution) Reset() {
	*x = Solution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Solution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Solution) ProtoMessage() {}

func (x *Solution) ProtoReflect() protoreflect.Message {
	mi := &file_advent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Solution.ProtoReflect.Descriptor instead.
func (*Solution) Descriptor() ([]byte, []int) {
	return file_advent_proto_rawDescGZIP(), []int{1}
}

func (x *Solution) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Solution) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *Solution) GetPart() int32 {
	if x != nil {
		return x.Part
	}
	return 0
}

func (x *Solution) GetBigAnswer() int64 {
	if x != nil {
		return x.BigAnswer
	}
	return 0
}

func (x *Solution) GetStringAnswer() string {
	if x != nil {
		return x.StringAnswer
	}
	return ""
}

func (x *Solution) GetAnswer() int32 {
	if x != nil {
		return x.Answer
	}
	return 0
}

type SolveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year int32  `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Day  int32  `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	Part int32  `protobuf:"varint,3,opt,name=part,proto3" json:"part,omitempty"`
	Data string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SolveRequest) Reset() {
	*x = SolveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolveRequest) ProtoMessage() {}

func (x *SolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_advent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolveRequest.ProtoReflect.Descriptor instead.
func (*SolveRequest) Descriptor() ([]byte, []int) {
	return file_advent_proto_rawDescGZIP(), []int{2}
}

func (x *SolveRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *SolveRequest) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *SolveRequest) GetPart() int32 {
	if x != nil {
		return x.Part
	}
	return 0
}

func (x *SolveRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type SolveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer       int32  `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
	StringAnswer string `protobuf:"bytes,2,opt,name=string_answer,json=stringAnswer,proto3" json:"string_answer,omitempty"`
	BigAnswer    int64  `protobuf:"varint,3,opt,name=big_answer,json=bigAnswer,proto3" json:"big_answer,omitempty"`
}

func (x *SolveResponse) Reset() {
	*x = SolveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolveResponse) ProtoMessage() {}

func (x *SolveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_advent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolveResponse.ProtoReflect.Descriptor instead.
func (*SolveResponse) Descriptor() ([]byte, []int) {
	return file_advent_proto_rawDescGZIP(), []int{3}
}

func (x *SolveResponse) GetAnswer() int32 {
	if x != nil {
		return x.Answer
	}
	return 0
}

func (x *SolveResponse) GetStringAnswer() string {
	if x != nil {
		return x.StringAnswer
	}
	return ""
}

func (x *SolveResponse) GetBigAnswer() int64 {
	if x != nil {
		return x.BigAnswer
	}
	return 0
}

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year    int32  `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Day     int32  `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	Data    string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	DualDay bool   `protobuf:"varint,4,opt,name=dual_day,json=dualDay,proto3" json:"dual_day,omitempty"`
	Part    int32  `protobuf:"varint,5,opt,name=part,proto3" json:"part,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_advent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_advent_proto_rawDescGZIP(), []int{4}
}

func (x *UploadRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *UploadRequest) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *UploadRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *UploadRequest) GetDualDay() bool {
	if x != nil {
		return x.DualDay
	}
	return false
}

func (x *UploadRequest) GetPart() int32 {
	if x != nil {
		return x.Part
	}
	return 0
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_advent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_advent_proto_rawDescGZIP(), []int{5}
}

type GetDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Day  int32 `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	Part int32 `protobuf:"varint,3,opt,name=part,proto3" json:"part,omitempty"`
}

func (x *GetDataRequest) Reset() {
	*x = GetDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataRequest) ProtoMessage() {}

func (x *GetDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_advent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataRequest.ProtoReflect.Descriptor instead.
func (*GetDataRequest) Descriptor() ([]byte, []int) {
	return file_advent_proto_rawDescGZIP(), []int{6}
}

func (x *GetDataRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *GetDataRequest) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *GetDataRequest) GetPart() int32 {
	if x != nil {
		return x.Part
	}
	return 0
}

type GetDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetDataResponse) Reset() {
	*x = GetDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataResponse) ProtoMessage() {}

func (x *GetDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_advent_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataResponse.ProtoReflect.Descriptor instead.
func (*GetDataResponse) Descriptor() ([]byte, []int) {
	return file_advent_proto_rawDescGZIP(), []int{7}
}

func (x *GetDataResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Callback string `protobuf:"bytes,1,opt,name=callback,proto3" json:"callback,omitempty"`
	Year     int32  `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_advent_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_advent_proto_rawDescGZIP(), []int{8}
}

func (x *RegisterRequest) GetCallback() string {
	if x != nil {
		return x.Callback
	}
	return ""
}

func (x *RegisterRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_advent_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_advent_proto_rawDescGZIP(), []int{9}
}

type AddSolutionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Solution *Solution `protobuf:"bytes,1,opt,name=solution,proto3" json:"solution,omitempty"`
}

func (x *AddSolutionRequest) Reset() {
	*x = AddSolutionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSolutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSolutionRequest) ProtoMessage() {}

func (x *AddSolutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_advent_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSolutionRequest.ProtoReflect.Descriptor instead.
func (*AddSolutionRequest) Descriptor() ([]byte, []int) {
	return file_advent_proto_rawDescGZIP(), []int{10}
}

func (x *AddSolutionRequest) GetSolution() *Solution {
	if x != nil {
		return x.Solution
	}
	return nil
}

type AddSolutionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddSolutionResponse) Reset() {
	*x = AddSolutionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSolutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSolutionResponse) ProtoMessage() {}

func (x *AddSolutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_advent_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSolutionResponse.ProtoReflect.Descriptor instead.
func (*AddSolutionResponse) Descriptor() ([]byte, []int) {
	return file_advent_proto_rawDescGZIP(), []int{11}
}

type GetSolutionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Day  int32 `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	Part int32 `protobuf:"varint,3,opt,name=part,proto3" json:"part,omitempty"`
}

func (x *GetSolutionRequest) Reset() {
	*x = GetSolutionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSolutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSolutionRequest) ProtoMessage() {}

func (x *GetSolutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_advent_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSolutionRequest.ProtoReflect.Descriptor instead.
func (*GetSolutionRequest) Descriptor() ([]byte, []int) {
	return file_advent_proto_rawDescGZIP(), []int{12}
}

func (x *GetSolutionRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *GetSolutionRequest) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *GetSolutionRequest) GetPart() int32 {
	if x != nil {
		return x.Part
	}
	return 0
}

type GetSolutionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Solution *Solution `protobuf:"bytes,1,opt,name=solution,proto3" json:"solution,omitempty"`
}

func (x *GetSolutionResponse) Reset() {
	*x = GetSolutionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advent_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSolutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSolutionResponse) ProtoMessage() {}

func (x *GetSolutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_advent_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSolutionResponse.ProtoReflect.Descriptor instead.
func (*GetSolutionResponse) Descriptor() ([]byte, []int) {
	return file_advent_proto_rawDescGZIP(), []int{13}
}

func (x *GetSolutionResponse) GetSolution() *Solution {
	if x != nil {
		return x.Solution
	}
	return nil
}

var File_advent_proto protoreflect.FileDescriptor

var file_advent_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x41, 0x0a, 0x09,
	0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61,
	0x64, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0xa0, 0x01, 0x0a, 0x08, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64,
	0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x67, 0x5f, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x69, 0x67, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x22, 0x5c, 0x0a, 0x0c, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x6b, 0x0a, 0x0d, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x69, 0x67, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x62, 0x69, 0x67, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x78, 0x0a,
	0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x64, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x75, 0x61, 0x6c,
	0x5f, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x75, 0x61, 0x6c,
	0x44, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x72, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x72, 0x74, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x41, 0x0a, 0x0f,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22,
	0x12, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x48, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x64,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x15, 0x0a,
	0x13, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x72, 0x74, 0x22, 0x49, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0x59, 0x0a, 0x13, 0x41, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x66, 0x43, 0x6f, 0x64, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x12,
	0x1a, 0x2e, 0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x53,
	0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x64,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x6f, 0x6c, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x87, 0x02, 0x0a, 0x1b, 0x41,
	0x64, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x66, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1b, 0x2e, 0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x2e,
	0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61,
	0x64, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54,
	0x0a, 0x0b, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e,
	0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x64, 0x64,
	0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x41,
	0x64, 0x64, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x32, 0x53, 0x0a, 0x0d, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x1a,
	0x2e, 0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x6f,
	0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x64, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x6c,
	0x6f, 0x67, 0x69, 0x63, 0x2f, 0x61, 0x64, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x66, 0x63, 0x6f, 0x64,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_advent_proto_rawDescOnce sync.Once
	file_advent_proto_rawDescData = file_advent_proto_rawDesc
)

func file_advent_proto_rawDescGZIP() []byte {
	file_advent_proto_rawDescOnce.Do(func() {
		file_advent_proto_rawDescData = protoimpl.X.CompressGZIP(file_advent_proto_rawDescData)
	})
	return file_advent_proto_rawDescData
}

var file_advent_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_advent_proto_goTypes = []interface{}{
	(*Solutions)(nil),           // 0: adventofcode.Solutions
	(*Solution)(nil),            // 1: adventofcode.Solution
	(*SolveRequest)(nil),        // 2: adventofcode.SolveRequest
	(*SolveResponse)(nil),       // 3: adventofcode.SolveResponse
	(*UploadRequest)(nil),       // 4: adventofcode.UploadRequest
	(*UploadResponse)(nil),      // 5: adventofcode.UploadResponse
	(*GetDataRequest)(nil),      // 6: adventofcode.GetDataRequest
	(*GetDataResponse)(nil),     // 7: adventofcode.GetDataResponse
	(*RegisterRequest)(nil),     // 8: adventofcode.RegisterRequest
	(*RegisterResponse)(nil),    // 9: adventofcode.RegisterResponse
	(*AddSolutionRequest)(nil),  // 10: adventofcode.AddSolutionRequest
	(*AddSolutionResponse)(nil), // 11: adventofcode.AddSolutionResponse
	(*GetSolutionRequest)(nil),  // 12: adventofcode.GetSolutionRequest
	(*GetSolutionResponse)(nil), // 13: adventofcode.GetSolutionResponse
}
var file_advent_proto_depIdxs = []int32{
	1,  // 0: adventofcode.Solutions.solutions:type_name -> adventofcode.Solution
	1,  // 1: adventofcode.AddSolutionRequest.solution:type_name -> adventofcode.Solution
	1,  // 2: adventofcode.GetSolutionResponse.solution:type_name -> adventofcode.Solution
	2,  // 3: adventofcode.AdventOfCodeService.Solve:input_type -> adventofcode.SolveRequest
	4,  // 4: adventofcode.AdventOfCodeInternalService.Upload:input_type -> adventofcode.UploadRequest
	8,  // 5: adventofcode.AdventOfCodeInternalService.Register:input_type -> adventofcode.RegisterRequest
	10, // 6: adventofcode.AdventOfCodeInternalService.AddSolution:input_type -> adventofcode.AddSolutionRequest
	2,  // 7: adventofcode.SolverService.Solve:input_type -> adventofcode.SolveRequest
	3,  // 8: adventofcode.AdventOfCodeService.Solve:output_type -> adventofcode.SolveResponse
	5,  // 9: adventofcode.AdventOfCodeInternalService.Upload:output_type -> adventofcode.UploadResponse
	9,  // 10: adventofcode.AdventOfCodeInternalService.Register:output_type -> adventofcode.RegisterResponse
	11, // 11: adventofcode.AdventOfCodeInternalService.AddSolution:output_type -> adventofcode.AddSolutionResponse
	3,  // 12: adventofcode.SolverService.Solve:output_type -> adventofcode.SolveResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_advent_proto_init() }
func file_advent_proto_init() {
	if File_advent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_advent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Solutions); i {
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
		file_advent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Solution); i {
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
		file_advent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolveRequest); i {
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
		file_advent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolveResponse); i {
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
		file_advent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_advent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
		file_advent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataRequest); i {
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
		file_advent_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataResponse); i {
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
		file_advent_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_advent_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_advent_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSolutionRequest); i {
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
		file_advent_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSolutionResponse); i {
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
		file_advent_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSolutionRequest); i {
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
		file_advent_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSolutionResponse); i {
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
			RawDescriptor: file_advent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_advent_proto_goTypes,
		DependencyIndexes: file_advent_proto_depIdxs,
		MessageInfos:      file_advent_proto_msgTypes,
	}.Build()
	File_advent_proto = out.File
	file_advent_proto_rawDesc = nil
	file_advent_proto_goTypes = nil
	file_advent_proto_depIdxs = nil
}
