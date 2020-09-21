// Copyright 2020 Google LLC
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
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/cloud/aiplatform/v1beta1/explanation_metadata.proto

package aiplatform

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

// Metadata describing the Model's input and output for explanation.
type ExplanationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Map from feature names to feature input metadata. Keys are the name of the
	// features. Values are the specification of the feature.
	//
	// An empty InputMetadata is valid. It describes a text feature which has the
	// name specified as the key in [ExplanationMetadata.inputs][google.cloud.aiplatform.v1beta1.ExplanationMetadata.inputs]. The baseline
	// of the empty feature is chosen by AI Platform.
	//
	Inputs map[string]*ExplanationMetadata_InputMetadata `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Required. Map from output names to output metadata.
	//
	// Keys are the name of the output field in the prediction to be explained.
	// Currently only one key is allowed.
	//
	Outputs map[string]*ExplanationMetadata_OutputMetadata `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Points to a YAML file stored on Google Cloud Storage describing the format
	// of the [feature attributions][google.cloud.aiplatform.v1beta1.Attribution.feature_attributions].
	// The schema is defined as an OpenAPI 3.0.2
	// [Schema Object](https://tinyurl.com/y538mdwt#schema-object).
	// AutoML tabular Models always have this field populated by AI Platform.
	// Note: The URI given on output may be different, including the URI scheme,
	// than the one given on input. The output URI will point to a location where
	// the user only has a read access.
	FeatureAttributionsSchemaUri string `protobuf:"bytes,3,opt,name=feature_attributions_schema_uri,json=featureAttributionsSchemaUri,proto3" json:"feature_attributions_schema_uri,omitempty"`
}

func (x *ExplanationMetadata) Reset() {
	*x = ExplanationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplanationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplanationMetadata) ProtoMessage() {}

func (x *ExplanationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplanationMetadata.ProtoReflect.Descriptor instead.
func (*ExplanationMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *ExplanationMetadata) GetInputs() map[string]*ExplanationMetadata_InputMetadata {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *ExplanationMetadata) GetOutputs() map[string]*ExplanationMetadata_OutputMetadata {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *ExplanationMetadata) GetFeatureAttributionsSchemaUri() string {
	if x != nil {
		return x.FeatureAttributionsSchemaUri
	}
	return ""
}

// Metadata of the input of a feature.
type ExplanationMetadata_InputMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Baseline inputs for this feature.
	//
	// If no baseline is specified, AI Platform chooses the baseline for this
	// feature. If multiple baselines are specified, AI Platform returns the
	// average attributions across them in
	// [Attributions.baseline_attribution][].
	//
	// The element of the baselines must be in the same format as the feature's
	// input in the [instance][google.cloud.aiplatform.v1beta1.ExplainRequest.instances][]. The schema of any
	// single instance may be specified via Endpoint's DeployedModels'
	// [Model's][google.cloud.aiplatform.v1beta1.DeployedModel.model]
	// [PredictSchemata's][google.cloud.aiplatform.v1beta1.Model.predict_schemata]
	// [instance_schema_uri][google.cloud.aiplatform.v1beta1.PredictSchemata.instance_schema_uri].
	InputBaselines []*structpb.Value `protobuf:"bytes,1,rep,name=input_baselines,json=inputBaselines,proto3" json:"input_baselines,omitempty"`
}

func (x *ExplanationMetadata_InputMetadata) Reset() {
	*x = ExplanationMetadata_InputMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplanationMetadata_InputMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplanationMetadata_InputMetadata) ProtoMessage() {}

func (x *ExplanationMetadata_InputMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplanationMetadata_InputMetadata.ProtoReflect.Descriptor instead.
func (*ExplanationMetadata_InputMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ExplanationMetadata_InputMetadata) GetInputBaselines() []*structpb.Value {
	if x != nil {
		return x.InputBaselines
	}
	return nil
}

// Metadata of the prediction output to be explained.
type ExplanationMetadata_OutputMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines how to map [Attribution.output_index][google.cloud.aiplatform.v1beta1.Attribution.output_index] to
	// [Attribution.output_display_name][google.cloud.aiplatform.v1beta1.Attribution.output_display_name].
	//
	// If neither of the fields are specified,
	// [Attribution.output_display_name][google.cloud.aiplatform.v1beta1.Attribution.output_display_name] will not be populated.
	//
	// Types that are assignable to DisplayNameMapping:
	//	*ExplanationMetadata_OutputMetadata_IndexDisplayNameMapping
	//	*ExplanationMetadata_OutputMetadata_DisplayNameMappingKey
	DisplayNameMapping isExplanationMetadata_OutputMetadata_DisplayNameMapping `protobuf_oneof:"display_name_mapping"`
}

func (x *ExplanationMetadata_OutputMetadata) Reset() {
	*x = ExplanationMetadata_OutputMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplanationMetadata_OutputMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplanationMetadata_OutputMetadata) ProtoMessage() {}

func (x *ExplanationMetadata_OutputMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplanationMetadata_OutputMetadata.ProtoReflect.Descriptor instead.
func (*ExplanationMetadata_OutputMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_rawDescGZIP(), []int{0, 1}
}

func (m *ExplanationMetadata_OutputMetadata) GetDisplayNameMapping() isExplanationMetadata_OutputMetadata_DisplayNameMapping {
	if m != nil {
		return m.DisplayNameMapping
	}
	return nil
}

func (x *ExplanationMetadata_OutputMetadata) GetIndexDisplayNameMapping() *structpb.Value {
	if x, ok := x.GetDisplayNameMapping().(*ExplanationMetadata_OutputMetadata_IndexDisplayNameMapping); ok {
		return x.IndexDisplayNameMapping
	}
	return nil
}

func (x *ExplanationMetadata_OutputMetadata) GetDisplayNameMappingKey() string {
	if x, ok := x.GetDisplayNameMapping().(*ExplanationMetadata_OutputMetadata_DisplayNameMappingKey); ok {
		return x.DisplayNameMappingKey
	}
	return ""
}

type isExplanationMetadata_OutputMetadata_DisplayNameMapping interface {
	isExplanationMetadata_OutputMetadata_DisplayNameMapping()
}

type ExplanationMetadata_OutputMetadata_IndexDisplayNameMapping struct {
	// Static mapping between the index and display name.
	//
	// Use this if the outputs are a deterministic n-dimensional array, e.g. a
	// list of scores of all the classes in a pre-defined order for a
	// multi-classification Model. It's not feasible if the outputs are
	// non-deterministic, e.g. the Model produces top-k classes or sort the
	// outputs by their values.
	//
	// The shape of the value must be an n-dimensional array of strings. The
	// number of dimentions must match that of the outputs to be explained.
	// The [Attribution.output_display_name][google.cloud.aiplatform.v1beta1.Attribution.output_display_name] is populated by locating in the
	// mapping with [Attribution.output_index][google.cloud.aiplatform.v1beta1.Attribution.output_index].
	//
	IndexDisplayNameMapping *structpb.Value `protobuf:"bytes,1,opt,name=index_display_name_mapping,json=indexDisplayNameMapping,proto3,oneof"`
}

type ExplanationMetadata_OutputMetadata_DisplayNameMappingKey struct {
	// Specify a field name in the prediction to look for the display name.
	//
	// Use this if the prediction contains the display names for the outputs.
	//
	// The display names in the prediction must have the same shape of the
	// outputs, so that it can be located by [Attribution.output_index][google.cloud.aiplatform.v1beta1.Attribution.output_index] for
	// a specific output.
	DisplayNameMappingKey string `protobuf:"bytes,2,opt,name=display_name_mapping_key,json=displayNameMappingKey,proto3,oneof"`
}

func (*ExplanationMetadata_OutputMetadata_IndexDisplayNameMapping) isExplanationMetadata_OutputMetadata_DisplayNameMapping() {
}

func (*ExplanationMetadata_OutputMetadata_DisplayNameMappingKey) isExplanationMetadata_OutputMetadata_DisplayNameMapping() {
}

var File_google_cloud_aiplatform_v1beta1_explanation_metadata_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x06, 0x0a, 0x13, 0x45,
	0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x5d, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x73, 0x12, 0x60, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x1f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x55, 0x72, 0x69, 0x1a, 0x50, 0x0a, 0x0d, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x0f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x42, 0x61, 0x73, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x1a, 0xba, 0x01, 0x0a,
	0x0e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x55, 0x0a, 0x1a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x17, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x18, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x15, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4b, 0x65,
	0x79, 0x42, 0x16, 0x0a, 0x14, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x1a, 0x7d, 0x0a, 0x0b, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x58, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6c,
	0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x7f, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x59, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6c,
	0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x8c, 0x01, 0x0a, 0x23, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x42, 0x18, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_goTypes = []interface{}{
	(*ExplanationMetadata)(nil),                // 0: google.cloud.aiplatform.v1beta1.ExplanationMetadata
	(*ExplanationMetadata_InputMetadata)(nil),  // 1: google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata
	(*ExplanationMetadata_OutputMetadata)(nil), // 2: google.cloud.aiplatform.v1beta1.ExplanationMetadata.OutputMetadata
	nil,                    // 3: google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputsEntry
	nil,                    // 4: google.cloud.aiplatform.v1beta1.ExplanationMetadata.OutputsEntry
	(*structpb.Value)(nil), // 5: google.protobuf.Value
}
var file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_depIdxs = []int32{
	3, // 0: google.cloud.aiplatform.v1beta1.ExplanationMetadata.inputs:type_name -> google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputsEntry
	4, // 1: google.cloud.aiplatform.v1beta1.ExplanationMetadata.outputs:type_name -> google.cloud.aiplatform.v1beta1.ExplanationMetadata.OutputsEntry
	5, // 2: google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata.input_baselines:type_name -> google.protobuf.Value
	5, // 3: google.cloud.aiplatform.v1beta1.ExplanationMetadata.OutputMetadata.index_display_name_mapping:type_name -> google.protobuf.Value
	1, // 4: google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputsEntry.value:type_name -> google.cloud.aiplatform.v1beta1.ExplanationMetadata.InputMetadata
	2, // 5: google.cloud.aiplatform.v1beta1.ExplanationMetadata.OutputsEntry.value:type_name -> google.cloud.aiplatform.v1beta1.ExplanationMetadata.OutputMetadata
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_init() }
func file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_explanation_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExplanationMetadata); i {
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
		file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExplanationMetadata_InputMetadata); i {
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
		file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExplanationMetadata_OutputMetadata); i {
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
	file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ExplanationMetadata_OutputMetadata_IndexDisplayNameMapping)(nil),
		(*ExplanationMetadata_OutputMetadata_DisplayNameMappingKey)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_explanation_metadata_proto = out.File
	file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_explanation_metadata_proto_depIdxs = nil
}
