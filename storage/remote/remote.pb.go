// Code generated by protoc-gen-go.
// source: remote.proto
// DO NOT EDIT!

/*
Package remote is a generated protocol buffer package.

It is generated from these files:
	remote.proto

It has these top-level messages:
	Sample
	LabelPair
	TimeSeries
	WriteRequest
	ReadRequest
	ReadResponse
	LabelMatchers
	LabelMatcher
*/
package remote

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

type MatchType int32

const (
	MatchType_EQUAL          MatchType = 0
	MatchType_NOT_EQUAL      MatchType = 1
	MatchType_REGEX_MATCH    MatchType = 2
	MatchType_REGEX_NO_MATCH MatchType = 3
)

var MatchType_name = map[int32]string{
	0: "EQUAL",
	1: "NOT_EQUAL",
	2: "REGEX_MATCH",
	3: "REGEX_NO_MATCH",
}
var MatchType_value = map[string]int32{
	"EQUAL":          0,
	"NOT_EQUAL":      1,
	"REGEX_MATCH":    2,
	"REGEX_NO_MATCH": 3,
}

func (x MatchType) String() string {
	return proto.EnumName(MatchType_name, int32(x))
}
func (MatchType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Sample struct {
	Value       float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	TimestampMs int64   `protobuf:"varint,2,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
}

func (m *Sample) Reset()                    { *m = Sample{} }
func (m *Sample) String() string            { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()               {}
func (*Sample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LabelPair struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *LabelPair) Reset()                    { *m = LabelPair{} }
func (m *LabelPair) String() string            { return proto.CompactTextString(m) }
func (*LabelPair) ProtoMessage()               {}
func (*LabelPair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type TimeSeries struct {
	Labels []*LabelPair `protobuf:"bytes,1,rep,name=labels" json:"labels,omitempty"`
	// Sorted by time, oldest sample first.
	Samples []*Sample `protobuf:"bytes,2,rep,name=samples" json:"samples,omitempty"`
}

func (m *TimeSeries) Reset()                    { *m = TimeSeries{} }
func (m *TimeSeries) String() string            { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()               {}
func (*TimeSeries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TimeSeries) GetLabels() []*LabelPair {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TimeSeries) GetSamples() []*Sample {
	if m != nil {
		return m.Samples
	}
	return nil
}

type WriteRequest struct {
	Timeseries []*TimeSeries `protobuf:"bytes,1,rep,name=timeseries" json:"timeseries,omitempty"`
}

func (m *WriteRequest) Reset()                    { *m = WriteRequest{} }
func (m *WriteRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()               {}
func (*WriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WriteRequest) GetTimeseries() []*TimeSeries {
	if m != nil {
		return m.Timeseries
	}
	return nil
}

type ReadRequest struct {
	StartTimestampMs int64            `protobuf:"varint,1,opt,name=start_timestamp_ms,json=startTimestampMs" json:"start_timestamp_ms,omitempty"`
	EndTimestampMs   int64            `protobuf:"varint,2,opt,name=end_timestamp_ms,json=endTimestampMs" json:"end_timestamp_ms,omitempty"`
	MatchersSet      []*LabelMatchers `protobuf:"bytes,3,rep,name=matchers_set,json=matchersSet" json:"matchers_set,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadRequest) GetMatchersSet() []*LabelMatchers {
	if m != nil {
		return m.MatchersSet
	}
	return nil
}

type ReadResponse struct {
	Timeseries []*TimeSeries `protobuf:"bytes,1,rep,name=timeseries" json:"timeseries,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReadResponse) GetTimeseries() []*TimeSeries {
	if m != nil {
		return m.Timeseries
	}
	return nil
}

type LabelMatchers struct {
	Matchers []*LabelMatcher `protobuf:"bytes,1,rep,name=matchers" json:"matchers,omitempty"`
}

func (m *LabelMatchers) Reset()                    { *m = LabelMatchers{} }
func (m *LabelMatchers) String() string            { return proto.CompactTextString(m) }
func (*LabelMatchers) ProtoMessage()               {}
func (*LabelMatchers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LabelMatchers) GetMatchers() []*LabelMatcher {
	if m != nil {
		return m.Matchers
	}
	return nil
}

type LabelMatcher struct {
	Type  MatchType `protobuf:"varint,1,opt,name=type,enum=remote.MatchType" json:"type,omitempty"`
	Name  string    `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Value string    `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *LabelMatcher) Reset()                    { *m = LabelMatcher{} }
func (m *LabelMatcher) String() string            { return proto.CompactTextString(m) }
func (*LabelMatcher) ProtoMessage()               {}
func (*LabelMatcher) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*Sample)(nil), "remote.Sample")
	proto.RegisterType((*LabelPair)(nil), "remote.LabelPair")
	proto.RegisterType((*TimeSeries)(nil), "remote.TimeSeries")
	proto.RegisterType((*WriteRequest)(nil), "remote.WriteRequest")
	proto.RegisterType((*ReadRequest)(nil), "remote.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "remote.ReadResponse")
	proto.RegisterType((*LabelMatchers)(nil), "remote.LabelMatchers")
	proto.RegisterType((*LabelMatcher)(nil), "remote.LabelMatcher")
	proto.RegisterEnum("remote.MatchType", MatchType_name, MatchType_value)
}

func init() { proto.RegisterFile("remote.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0xab, 0xd3, 0x40,
	0x10, 0x35, 0x49, 0x1b, 0xcd, 0x24, 0x8d, 0x71, 0xa9, 0xd0, 0x47, 0x0d, 0x08, 0x55, 0xa4, 0x48,
	0x45, 0xf0, 0x35, 0x4a, 0x51, 0xa4, 0x1f, 0xba, 0x8d, 0xe8, 0x5b, 0xd8, 0xda, 0x01, 0x03, 0x49,
	0x93, 0x9b, 0xdd, 0x5e, 0xb8, 0xbf, 0xe6, 0xfe, 0xd5, 0xbb, 0xd9, 0xcd, 0x57, 0xa1, 0x4f, 0xf7,
	0x2d, 0x33, 0xe7, 0xcc, 0x99, 0x33, 0xb3, 0x13, 0xf0, 0x2a, 0xcc, 0x0b, 0x81, 0x8b, 0xb2, 0x2a,
	0x44, 0x41, 0x6c, 0x1d, 0x85, 0x11, 0xd8, 0x7b, 0x96, 0x97, 0x19, 0x92, 0x29, 0x8c, 0x6f, 0x59,
	0x76, 0xc6, 0x99, 0xf1, 0xca, 0x98, 0x1b, 0x54, 0x07, 0xe4, 0x35, 0x78, 0x22, 0xcd, 0x91, 0x0b,
	0x49, 0x4a, 0x72, 0x3e, 0x33, 0x25, 0x68, 0x51, 0xb7, 0xcb, 0x6d, 0x78, 0xf8, 0x09, 0x9c, 0x35,
	0x3b, 0x60, 0xf6, 0x93, 0xa5, 0x15, 0x21, 0x30, 0x3a, 0xb1, 0x5c, 0x8b, 0x38, 0x54, 0x7d, 0xf7,
	0xca, 0xa6, 0x4a, 0xea, 0x20, 0x64, 0x00, 0xb1, 0x54, 0xd9, 0x63, 0x95, 0x22, 0x27, 0x6f, 0xc1,
	0xce, 0x6a, 0x11, 0x2e, 0x2b, 0xad, 0xb9, 0xbb, 0x7c, 0xb1, 0x68, 0xec, 0x76, 0xd2, 0xb4, 0x21,
	0x90, 0x39, 0x3c, 0xe5, 0xca, 0x72, 0xed, 0xa6, 0xe6, 0xfa, 0x2d, 0x57, 0x4f, 0x42, 0x5b, 0x38,
	0xfc, 0x02, 0xde, 0x9f, 0x2a, 0x15, 0x48, 0xf1, 0xe6, 0x2c, 0xed, 0x92, 0x25, 0x80, 0x32, 0xae,
	0x5a, 0x36, 0x8d, 0x48, 0x5b, 0xdc, 0x9b, 0xa1, 0x03, 0x56, 0x78, 0x6f, 0x80, 0x4b, 0x91, 0x1d,
	0x5b, 0x8d, 0xf7, 0x40, 0xe4, 0xe0, 0x95, 0x48, 0x2e, 0xd6, 0x62, 0xa8, 0xb5, 0x04, 0x0a, 0x89,
	0xfb, 0xdd, 0x48, 0xaf, 0x01, 0x9e, 0x8e, 0xc9, 0x95, 0x15, 0xfa, 0x32, 0x3f, 0x64, 0x7e, 0x06,
	0x2f, 0x67, 0xe2, 0xdf, 0x7f, 0xac, 0x78, 0xc2, 0x51, 0xcc, 0x2c, 0xe5, 0xee, 0xe5, 0xc5, 0x1a,
	0x36, 0x0d, 0x81, 0xba, 0x2d, 0x75, 0x8f, 0xa2, 0x9e, 0x52, 0x1b, 0xe4, 0x65, 0x71, 0xe2, 0xf8,
	0xa8, 0x29, 0x23, 0x98, 0x5c, 0x74, 0x20, 0x1f, 0xe0, 0x59, 0xdb, 0xa3, 0x91, 0x98, 0x5e, 0xb3,
	0x42, 0x3b, 0x56, 0x98, 0x80, 0x37, 0x44, 0xc8, 0x1b, 0x18, 0x89, 0xbb, 0x52, 0x5f, 0x82, 0xdf,
	0xbf, 0xa7, 0x82, 0x63, 0x09, 0x50, 0x05, 0x77, 0x07, 0x63, 0x5e, 0x3b, 0x18, 0x6b, 0x70, 0x30,
	0xef, 0x7e, 0x80, 0xd3, 0x15, 0x13, 0x07, 0xc6, 0xab, 0x5f, 0xbf, 0xa3, 0x75, 0xf0, 0x84, 0x4c,
	0xc0, 0xd9, 0xee, 0xe2, 0x44, 0x87, 0x06, 0x79, 0x2e, 0xdf, 0x6b, 0xf5, 0x6d, 0xf5, 0x37, 0xd9,
	0x44, 0xf1, 0xd7, 0xef, 0x81, 0x29, 0x3b, 0xf8, 0x3a, 0xb1, 0xdd, 0x35, 0x39, 0xeb, 0x60, 0xab,
	0xbf, 0xe0, 0xe3, 0x43, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x8e, 0x81, 0xfb, 0x15, 0x03, 0x00,
	0x00,
}
