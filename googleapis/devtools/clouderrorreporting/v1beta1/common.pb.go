// Copyright 2019 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.2
// source: google/devtools/clouderrorreporting/v1beta1/common.proto

package clouderrorreporting

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Description of a group of similar error events.
type ErrorGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The group resource name.
	// Example: <code>projects/my-project-123/groups/my-groupid</code>
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Group IDs are unique for a given project. If the same kind of error
	// occurs in different service contexts, it will receive the same group ID.
	GroupId string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// Associated tracking issues.
	TrackingIssues []*TrackingIssue `protobuf:"bytes,3,rep,name=tracking_issues,json=trackingIssues,proto3" json:"tracking_issues,omitempty"`
}

func (x *ErrorGroup) Reset() {
	*x = ErrorGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorGroup) ProtoMessage() {}

func (x *ErrorGroup) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorGroup.ProtoReflect.Descriptor instead.
func (*ErrorGroup) Descriptor() ([]byte, []int) {
	return file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ErrorGroup) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *ErrorGroup) GetTrackingIssues() []*TrackingIssue {
	if x != nil {
		return x.TrackingIssues
	}
	return nil
}

// Information related to tracking the progress on resolving the error.
type TrackingIssue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A URL pointing to a related entry in an issue tracking system.
	// Example: `https://github.com/user/project/issues/4`
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *TrackingIssue) Reset() {
	*x = TrackingIssue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackingIssue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackingIssue) ProtoMessage() {}

func (x *TrackingIssue) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackingIssue.ProtoReflect.Descriptor instead.
func (*TrackingIssue) Descriptor() ([]byte, []int) {
	return file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDescGZIP(), []int{1}
}

func (x *TrackingIssue) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// An error event which is returned by the Error Reporting system.
type ErrorEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time when the event occurred as provided in the error report.
	// If the report did not contain a timestamp, the time the error was received
	// by the Error Reporting system is used.
	EventTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// The `ServiceContext` for which this error was reported.
	ServiceContext *ServiceContext `protobuf:"bytes,2,opt,name=service_context,json=serviceContext,proto3" json:"service_context,omitempty"`
	// The stack trace that was reported or logged by the service.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Data about the context in which the error occurred.
	Context *ErrorContext `protobuf:"bytes,5,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *ErrorEvent) Reset() {
	*x = ErrorEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorEvent) ProtoMessage() {}

func (x *ErrorEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorEvent.ProtoReflect.Descriptor instead.
func (*ErrorEvent) Descriptor() ([]byte, []int) {
	return file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorEvent) GetEventTime() *timestamp.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *ErrorEvent) GetServiceContext() *ServiceContext {
	if x != nil {
		return x.ServiceContext
	}
	return nil
}

func (x *ErrorEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorEvent) GetContext() *ErrorContext {
	if x != nil {
		return x.Context
	}
	return nil
}

// Describes a running service that sends errors.
// Its version changes over time and multiple versions can run in parallel.
type ServiceContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An identifier of the service, such as the name of the
	// executable, job, or Google App Engine service name. This field is expected
	// to have a low number of values that are relatively stable over time, as
	// opposed to `version`, which can be changed whenever new code is deployed.
	//
	// Contains the service name for error reports extracted from Google
	// App Engine logs or `default` if the App Engine default service is used.
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// Represents the source code version that the developer provided,
	// which could represent a version label or a Git SHA-1 hash, for example.
	// For App Engine standard environment, the version is set to the version of
	// the app.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Type of the MonitoredResource. List of possible values:
	// https://cloud.google.com/monitoring/api/resources
	//
	// Value is set automatically for incoming errors and must not be set when
	// reporting errors.
	ResourceType string `protobuf:"bytes,4,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
}

func (x *ServiceContext) Reset() {
	*x = ServiceContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceContext) ProtoMessage() {}

func (x *ServiceContext) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceContext.ProtoReflect.Descriptor instead.
func (*ServiceContext) Descriptor() ([]byte, []int) {
	return file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceContext) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ServiceContext) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServiceContext) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

// A description of the context in which an error occurred.
// This data should be provided by the application when reporting an error,
// unless the
// error report has been generated automatically from Google App Engine logs.
type ErrorContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP request which was processed when the error was
	// triggered.
	HttpRequest *HttpRequestContext `protobuf:"bytes,1,opt,name=http_request,json=httpRequest,proto3" json:"http_request,omitempty"`
	// The user who caused or was affected by the crash.
	// This can be a user ID, an email address, or an arbitrary token that
	// uniquely identifies the user.
	// When sending an error report, leave this field empty if the user was not
	// logged in. In this case the
	// Error Reporting system will use other data, such as remote IP address, to
	// distinguish affected users. See `affected_users_count` in
	// `ErrorGroupStats`.
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// The location in the source code where the decision was made to
	// report the error, usually the place where it was logged.
	// For a logged exception this would be the source line where the
	// exception is logged, usually close to the place where it was
	// caught.
	ReportLocation *SourceLocation `protobuf:"bytes,3,opt,name=report_location,json=reportLocation,proto3" json:"report_location,omitempty"`
}

func (x *ErrorContext) Reset() {
	*x = ErrorContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorContext) ProtoMessage() {}

func (x *ErrorContext) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorContext.ProtoReflect.Descriptor instead.
func (*ErrorContext) Descriptor() ([]byte, []int) {
	return file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDescGZIP(), []int{4}
}

func (x *ErrorContext) GetHttpRequest() *HttpRequestContext {
	if x != nil {
		return x.HttpRequest
	}
	return nil
}

func (x *ErrorContext) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ErrorContext) GetReportLocation() *SourceLocation {
	if x != nil {
		return x.ReportLocation
	}
	return nil
}

// HTTP request data that is related to a reported error.
// This data should be provided by the application when reporting an error,
// unless the
// error report has been generated automatically from Google App Engine logs.
type HttpRequestContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of HTTP request, such as `GET`, `POST`, etc.
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// The URL of the request.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// The user agent information that is provided with the request.
	UserAgent string `protobuf:"bytes,3,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// The referrer information that is provided with the request.
	Referrer string `protobuf:"bytes,4,opt,name=referrer,proto3" json:"referrer,omitempty"`
	// The HTTP response status code for the request.
	ResponseStatusCode int32 `protobuf:"varint,5,opt,name=response_status_code,json=responseStatusCode,proto3" json:"response_status_code,omitempty"`
	// The IP address from which the request originated.
	// This can be IPv4, IPv6, or a token which is derived from the
	// IP address, depending on the data that has been provided
	// in the error report.
	RemoteIp string `protobuf:"bytes,6,opt,name=remote_ip,json=remoteIp,proto3" json:"remote_ip,omitempty"`
}

func (x *HttpRequestContext) Reset() {
	*x = HttpRequestContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpRequestContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpRequestContext) ProtoMessage() {}

func (x *HttpRequestContext) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpRequestContext.ProtoReflect.Descriptor instead.
func (*HttpRequestContext) Descriptor() ([]byte, []int) {
	return file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDescGZIP(), []int{5}
}

func (x *HttpRequestContext) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *HttpRequestContext) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HttpRequestContext) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *HttpRequestContext) GetReferrer() string {
	if x != nil {
		return x.Referrer
	}
	return ""
}

func (x *HttpRequestContext) GetResponseStatusCode() int32 {
	if x != nil {
		return x.ResponseStatusCode
	}
	return 0
}

func (x *HttpRequestContext) GetRemoteIp() string {
	if x != nil {
		return x.RemoteIp
	}
	return ""
}

// Indicates a location in the source code of the service for which errors are
// reported. `functionName` must be provided by the application when reporting
// an error, unless the error report contains a `message` with a supported
// exception stack trace. All fields are optional for the later case.
type SourceLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source code filename, which can include a truncated relative
	// path, or a full path from a production machine.
	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	// 1-based. 0 indicates that the line number is unknown.
	LineNumber int32 `protobuf:"varint,2,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
	// Human-readable name of a function or method.
	// The value can include optional context like the class or package name.
	// For example, `my.package.MyClass.method` in case of Java.
	FunctionName string `protobuf:"bytes,4,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
}

func (x *SourceLocation) Reset() {
	*x = SourceLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceLocation) ProtoMessage() {}

func (x *SourceLocation) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceLocation.ProtoReflect.Descriptor instead.
func (*SourceLocation) Descriptor() ([]byte, []int) {
	return file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDescGZIP(), []int{6}
}

func (x *SourceLocation) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *SourceLocation) GetLineNumber() int32 {
	if x != nil {
		return x.LineNumber
	}
	return 0
}

func (x *SourceLocation) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

var File_google_devtools_clouderrorreporting_v1beta1_common_proto protoreflect.FileDescriptor

var file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12,
	0x63, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x73, 0x3a, 0x55, 0xea, 0x41, 0x52, 0x0a, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x21, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x2f, 0x7b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x7d, 0x22, 0x21, 0x0a, 0x0d, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x9c,
	0x02, 0x0a, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x64, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x53, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x69, 0x0a,
	0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0xec, 0x01, 0x0a, 0x0c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x62, 0x0a, 0x0c, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x52, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x64, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc8, 0x01, 0x0a, 0x12, 0x48, 0x74, 0x74, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x12, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f,
	0x69, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x49, 0x70, 0x22, 0x73, 0x0a, 0x0e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0xef, 0x01, 0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x23,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x42, 0x65,
	0x74, 0x61, 0x31, 0xca, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDescOnce sync.Once
	file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDescData = file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDesc
)

func file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDescGZIP() []byte {
	file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDescOnce.Do(func() {
		file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDescData)
	})
	return file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDescData
}

var file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_devtools_clouderrorreporting_v1beta1_common_proto_goTypes = []interface{}{
	(*ErrorGroup)(nil),          // 0: google.devtools.clouderrorreporting.v1beta1.ErrorGroup
	(*TrackingIssue)(nil),       // 1: google.devtools.clouderrorreporting.v1beta1.TrackingIssue
	(*ErrorEvent)(nil),          // 2: google.devtools.clouderrorreporting.v1beta1.ErrorEvent
	(*ServiceContext)(nil),      // 3: google.devtools.clouderrorreporting.v1beta1.ServiceContext
	(*ErrorContext)(nil),        // 4: google.devtools.clouderrorreporting.v1beta1.ErrorContext
	(*HttpRequestContext)(nil),  // 5: google.devtools.clouderrorreporting.v1beta1.HttpRequestContext
	(*SourceLocation)(nil),      // 6: google.devtools.clouderrorreporting.v1beta1.SourceLocation
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_google_devtools_clouderrorreporting_v1beta1_common_proto_depIdxs = []int32{
	1, // 0: google.devtools.clouderrorreporting.v1beta1.ErrorGroup.tracking_issues:type_name -> google.devtools.clouderrorreporting.v1beta1.TrackingIssue
	7, // 1: google.devtools.clouderrorreporting.v1beta1.ErrorEvent.event_time:type_name -> google.protobuf.Timestamp
	3, // 2: google.devtools.clouderrorreporting.v1beta1.ErrorEvent.service_context:type_name -> google.devtools.clouderrorreporting.v1beta1.ServiceContext
	4, // 3: google.devtools.clouderrorreporting.v1beta1.ErrorEvent.context:type_name -> google.devtools.clouderrorreporting.v1beta1.ErrorContext
	5, // 4: google.devtools.clouderrorreporting.v1beta1.ErrorContext.http_request:type_name -> google.devtools.clouderrorreporting.v1beta1.HttpRequestContext
	6, // 5: google.devtools.clouderrorreporting.v1beta1.ErrorContext.report_location:type_name -> google.devtools.clouderrorreporting.v1beta1.SourceLocation
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_devtools_clouderrorreporting_v1beta1_common_proto_init() }
func file_google_devtools_clouderrorreporting_v1beta1_common_proto_init() {
	if File_google_devtools_clouderrorreporting_v1beta1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorGroup); i {
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
		file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackingIssue); i {
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
		file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorEvent); i {
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
		file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceContext); i {
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
		file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorContext); i {
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
		file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpRequestContext); i {
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
		file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceLocation); i {
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
			RawDescriptor: file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_clouderrorreporting_v1beta1_common_proto_goTypes,
		DependencyIndexes: file_google_devtools_clouderrorreporting_v1beta1_common_proto_depIdxs,
		MessageInfos:      file_google_devtools_clouderrorreporting_v1beta1_common_proto_msgTypes,
	}.Build()
	File_google_devtools_clouderrorreporting_v1beta1_common_proto = out.File
	file_google_devtools_clouderrorreporting_v1beta1_common_proto_rawDesc = nil
	file_google_devtools_clouderrorreporting_v1beta1_common_proto_goTypes = nil
	file_google_devtools_clouderrorreporting_v1beta1_common_proto_depIdxs = nil
}
