// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/api/cloud/usage/v1/message.proto

package usage

import (
	reflect "reflect"
	"strconv"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RecordType int32

const (
	RECORD_TYPE_UNSPECIFIED      RecordType = 0
	RECORD_TYPE_ACTIONS          RecordType = 1
	RECORD_TYPE_ACTIVE_STORAGE   RecordType = 2
	RECORD_TYPE_RETAINED_STORAGE RecordType = 3
)

// Enum value maps for RecordType.
var (
	RecordType_name = map[int32]string{
		0: "RECORD_TYPE_UNSPECIFIED",
		1: "RECORD_TYPE_ACTIONS",
		2: "RECORD_TYPE_ACTIVE_STORAGE",
		3: "RECORD_TYPE_RETAINED_STORAGE",
	}
	RecordType_value = map[string]int32{
		"RECORD_TYPE_UNSPECIFIED":      0,
		"RECORD_TYPE_ACTIONS":          1,
		"RECORD_TYPE_ACTIVE_STORAGE":   2,
		"RECORD_TYPE_RETAINED_STORAGE": 3,
	}
)

func (x RecordType) Enum() *RecordType {
	p := new(RecordType)
	*p = x
	return p
}

func (x RecordType) String() string {
	switch x {
	case RECORD_TYPE_UNSPECIFIED:
		return "Unspecified"
	case RECORD_TYPE_ACTIONS:
		return "Actions"
	case RECORD_TYPE_ACTIVE_STORAGE:
		return "ActiveStorage"
	case RECORD_TYPE_RETAINED_STORAGE:
		return "RetainedStorage"
	default:
		return strconv.Itoa(int(x))
	}

}

func (RecordType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_cloud_usage_v1_message_proto_enumTypes[0].Descriptor()
}

func (RecordType) Type() protoreflect.EnumType {
	return &file_temporal_api_cloud_usage_v1_message_proto_enumTypes[0]
}

func (x RecordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecordType.Descriptor instead.
func (RecordType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_cloud_usage_v1_message_proto_rawDescGZIP(), []int{0}
}

type RecordUnit int32

const (
	RECORD_UNIT_UNSPECIFIED  RecordUnit = 0
	RECORD_UNIT_NUMBER       RecordUnit = 1
	RECORD_UNIT_BYTE_SECONDS RecordUnit = 2
)

// Enum value maps for RecordUnit.
var (
	RecordUnit_name = map[int32]string{
		0: "RECORD_UNIT_UNSPECIFIED",
		1: "RECORD_UNIT_NUMBER",
		2: "RECORD_UNIT_BYTE_SECONDS",
	}
	RecordUnit_value = map[string]int32{
		"RECORD_UNIT_UNSPECIFIED":  0,
		"RECORD_UNIT_NUMBER":       1,
		"RECORD_UNIT_BYTE_SECONDS": 2,
	}
)

func (x RecordUnit) Enum() *RecordUnit {
	p := new(RecordUnit)
	*p = x
	return p
}

func (x RecordUnit) String() string {
	switch x {
	case RECORD_UNIT_UNSPECIFIED:
		return "Unspecified"
	case RECORD_UNIT_NUMBER:
		return "Number"
	case RECORD_UNIT_BYTE_SECONDS:
		return "ByteSeconds"
	default:
		return strconv.Itoa(int(x))
	}

}

func (RecordUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_cloud_usage_v1_message_proto_enumTypes[1].Descriptor()
}

func (RecordUnit) Type() protoreflect.EnumType {
	return &file_temporal_api_cloud_usage_v1_message_proto_enumTypes[1]
}

func (x RecordUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecordUnit.Descriptor instead.
func (RecordUnit) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_cloud_usage_v1_message_proto_rawDescGZIP(), []int{1}
}

type GroupByKey int32

const (
	GROUP_BY_KEY_UNSPECIFIED GroupByKey = 0
	GROUP_BY_KEY_NAMESPACE   GroupByKey = 1
)

// Enum value maps for GroupByKey.
var (
	GroupByKey_name = map[int32]string{
		0: "GROUP_BY_KEY_UNSPECIFIED",
		1: "GROUP_BY_KEY_NAMESPACE",
	}
	GroupByKey_value = map[string]int32{
		"GROUP_BY_KEY_UNSPECIFIED": 0,
		"GROUP_BY_KEY_NAMESPACE":   1,
	}
)

func (x GroupByKey) Enum() *GroupByKey {
	p := new(GroupByKey)
	*p = x
	return p
}

func (x GroupByKey) String() string {
	switch x {
	case GROUP_BY_KEY_UNSPECIFIED:
		return "Unspecified"
	case GROUP_BY_KEY_NAMESPACE:
		return "Namespace"
	default:
		return strconv.Itoa(int(x))
	}

}

func (GroupByKey) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_cloud_usage_v1_message_proto_enumTypes[2].Descriptor()
}

func (GroupByKey) Type() protoreflect.EnumType {
	return &file_temporal_api_cloud_usage_v1_message_proto_enumTypes[2]
}

func (x GroupByKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupByKey.Descriptor instead.
func (GroupByKey) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_cloud_usage_v1_message_proto_rawDescGZIP(), []int{2}
}

type Summary struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Start of UTC day for now (inclusive)
	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// End of UTC day for now (exclusive)
	EndTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Records grouped by namespace
	RecordGroups []*RecordGroup `protobuf:"bytes,3,rep,name=record_groups,json=recordGroups,proto3" json:"record_groups,omitempty"`
	// True if data for given time window is not fully available yet (e.g. delays)
	// When true, records for the given time range could still be added/updated in the future (until false)
	Incomplete    bool `protobuf:"varint,4,opt,name=incomplete,proto3" json:"incomplete,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Summary) Reset() {
	*x = Summary{}
	mi := &file_temporal_api_cloud_usage_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Summary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary) ProtoMessage() {}

func (x *Summary) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_usage_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary.ProtoReflect.Descriptor instead.
func (*Summary) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_usage_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *Summary) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Summary) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Summary) GetRecordGroups() []*RecordGroup {
	if x != nil {
		return x.RecordGroups
	}
	return nil
}

func (x *Summary) GetIncomplete() bool {
	if x != nil {
		return x.Incomplete
	}
	return false
}

type RecordGroup struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// GroupBy keys and their values for this record group. Multiple fields are combined with logical AND.
	GroupBys      []*GroupBy `protobuf:"bytes,1,rep,name=group_bys,json=groupBys,proto3" json:"group_bys,omitempty"`
	Records       []*Record  `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecordGroup) Reset() {
	*x = RecordGroup{}
	mi := &file_temporal_api_cloud_usage_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecordGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordGroup) ProtoMessage() {}

func (x *RecordGroup) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_usage_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordGroup.ProtoReflect.Descriptor instead.
func (*RecordGroup) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_usage_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *RecordGroup) GetGroupBys() []*GroupBy {
	if x != nil {
		return x.GroupBys
	}
	return nil
}

func (x *RecordGroup) GetRecords() []*Record {
	if x != nil {
		return x.Records
	}
	return nil
}

type GroupBy struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           GroupByKey             `protobuf:"varint,1,opt,name=key,proto3,enum=temporal.api.cloud.usage.v1.GroupByKey" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GroupBy) Reset() {
	*x = GroupBy{}
	mi := &file_temporal_api_cloud_usage_v1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupBy) ProtoMessage() {}

func (x *GroupBy) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_usage_v1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupBy.ProtoReflect.Descriptor instead.
func (*GroupBy) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_usage_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *GroupBy) GetKey() GroupByKey {
	if x != nil {
		return x.Key
	}
	return GROUP_BY_KEY_UNSPECIFIED
}

func (x *GroupBy) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Record struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          RecordType             `protobuf:"varint,1,opt,name=type,proto3,enum=temporal.api.cloud.usage.v1.RecordType" json:"type,omitempty"`
	Unit          RecordUnit             `protobuf:"varint,2,opt,name=unit,proto3,enum=temporal.api.cloud.usage.v1.RecordUnit" json:"unit,omitempty"`
	Value         float64                `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Record) Reset() {
	*x = Record{}
	mi := &file_temporal_api_cloud_usage_v1_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_usage_v1_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_usage_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *Record) GetType() RecordType {
	if x != nil {
		return x.Type
	}
	return RECORD_TYPE_UNSPECIFIED
}

func (x *Record) GetUnit() RecordUnit {
	if x != nil {
		return x.Unit
	}
	return RECORD_UNIT_UNSPECIFIED
}

func (x *Record) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_temporal_api_cloud_usage_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_cloud_usage_v1_message_proto_rawDesc = []byte{
	0x0a, 0x29, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x3d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x39, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x51, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x42, 0x02, 0x68, 0x00, 0x12, 0x22, 0x0a, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x02, 0x68, 0x00, 0x22, 0x97, 0x01, 0x0a, 0x0b, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x45, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x62, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x52,
	0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x73, 0x42, 0x02, 0x68, 0x00, 0x12, 0x41, 0x0a,
	0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x02, 0x68, 0x00, 0x22, 0x62, 0x0a,
	0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x12, 0x3d, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x42, 0x02, 0x68, 0x00, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x68, 0x00, 0x22, 0xa4,
	0x01, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x3f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x3f, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e,
	0x69, 0x74, 0x42, 0x02, 0x68, 0x00, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x68, 0x00, 0x2a,
	0x84, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x17, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x43,
	0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x53,
	0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x10,
	0x02, 0x12, 0x20, 0x0a, 0x1c, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x52, 0x45, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x10,
	0x03, 0x2a, 0x5f, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x12,
	0x1b, 0x0a, 0x17, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x52,
	0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52,
	0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x55, 0x4e, 0x49, 0x54,
	0x5f, 0x42, 0x59, 0x54, 0x45, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x53, 0x10, 0x02, 0x2a,
	0x46, 0x0a, 0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x18,
	0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x42, 0x59, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x5f, 0x42, 0x59, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50,
	0x41, 0x43, 0x45, 0x10, 0x01, 0x42, 0x9d, 0x01, 0x0a, 0x1e, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x61, 0x67, 0x65, 0xaa, 0x02, 0x1d,
	0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x21, 0x54,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x55, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_cloud_usage_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_cloud_usage_v1_message_proto_rawDescData = file_temporal_api_cloud_usage_v1_message_proto_rawDesc
)

func file_temporal_api_cloud_usage_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_cloud_usage_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_cloud_usage_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_cloud_usage_v1_message_proto_rawDescData)
	})
	return file_temporal_api_cloud_usage_v1_message_proto_rawDescData
}

var file_temporal_api_cloud_usage_v1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_temporal_api_cloud_usage_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_temporal_api_cloud_usage_v1_message_proto_goTypes = []any{
	(RecordType)(0),               // 0: temporal.api.cloud.usage.v1.RecordType
	(RecordUnit)(0),               // 1: temporal.api.cloud.usage.v1.RecordUnit
	(GroupByKey)(0),               // 2: temporal.api.cloud.usage.v1.GroupByKey
	(*Summary)(nil),               // 3: temporal.api.cloud.usage.v1.Summary
	(*RecordGroup)(nil),           // 4: temporal.api.cloud.usage.v1.RecordGroup
	(*GroupBy)(nil),               // 5: temporal.api.cloud.usage.v1.GroupBy
	(*Record)(nil),                // 6: temporal.api.cloud.usage.v1.Record
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_temporal_api_cloud_usage_v1_message_proto_depIdxs = []int32{
	7, // 0: temporal.api.cloud.usage.v1.Summary.start_time:type_name -> google.protobuf.Timestamp
	7, // 1: temporal.api.cloud.usage.v1.Summary.end_time:type_name -> google.protobuf.Timestamp
	4, // 2: temporal.api.cloud.usage.v1.Summary.record_groups:type_name -> temporal.api.cloud.usage.v1.RecordGroup
	5, // 3: temporal.api.cloud.usage.v1.RecordGroup.group_bys:type_name -> temporal.api.cloud.usage.v1.GroupBy
	6, // 4: temporal.api.cloud.usage.v1.RecordGroup.records:type_name -> temporal.api.cloud.usage.v1.Record
	2, // 5: temporal.api.cloud.usage.v1.GroupBy.key:type_name -> temporal.api.cloud.usage.v1.GroupByKey
	0, // 6: temporal.api.cloud.usage.v1.Record.type:type_name -> temporal.api.cloud.usage.v1.RecordType
	1, // 7: temporal.api.cloud.usage.v1.Record.unit:type_name -> temporal.api.cloud.usage.v1.RecordUnit
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_temporal_api_cloud_usage_v1_message_proto_init() }
func file_temporal_api_cloud_usage_v1_message_proto_init() {
	if File_temporal_api_cloud_usage_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_cloud_usage_v1_message_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_cloud_usage_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_cloud_usage_v1_message_proto_depIdxs,
		EnumInfos:         file_temporal_api_cloud_usage_v1_message_proto_enumTypes,
		MessageInfos:      file_temporal_api_cloud_usage_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_cloud_usage_v1_message_proto = out.File
	file_temporal_api_cloud_usage_v1_message_proto_rawDesc = nil
	file_temporal_api_cloud_usage_v1_message_proto_goTypes = nil
	file_temporal_api_cloud_usage_v1_message_proto_depIdxs = nil
}
