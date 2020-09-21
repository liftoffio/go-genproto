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
// source: google/cloud/aiplatform/v1beta1/prediction_service.proto

package aiplatform

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [PredictionService.Predict][google.cloud.aiplatform.v1beta1.PredictionService.Predict].
type PredictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the Endpoint requested to serve the prediction.
	// Format:
	// `projects/{project}/locations/{location}/endpoints/{endpoint}`
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Required. The instances that are the input to the prediction call.
	// A DeployedModel may have an upper limit on the number of instances it
	// supports per request, and when it is exceeded the prediction call errors
	// in case of AutoML Models, or, in case of customer created Models, the
	// behaviour is as documented by that Model.
	// The schema of any single instance may be specified via Endpoint's
	// DeployedModels' [Model's][google.cloud.aiplatform.v1beta1.DeployedModel.model]
	// [PredictSchemata's][google.cloud.aiplatform.v1beta1.Model.predict_schemata]
	// [instance_schema_uri][google.cloud.aiplatform.v1beta1.PredictSchemata.instance_schema_uri].
	Instances []*structpb.Value `protobuf:"bytes,2,rep,name=instances,proto3" json:"instances,omitempty"`
	// The parameters that govern the prediction. The schema of the parameters may
	// be specified via Endpoint's DeployedModels' [Model's ][google.cloud.aiplatform.v1beta1.DeployedModel.model]
	// [PredictSchemata's][google.cloud.aiplatform.v1beta1.Model.predict_schemata]
	// [parameters_schema_uri][google.cloud.aiplatform.v1beta1.PredictSchemata.parameters_schema_uri].
	Parameters *structpb.Value `protobuf:"bytes,3,opt,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *PredictRequest) Reset() {
	*x = PredictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_prediction_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictRequest) ProtoMessage() {}

func (x *PredictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_prediction_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictRequest.ProtoReflect.Descriptor instead.
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDescGZIP(), []int{0}
}

func (x *PredictRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *PredictRequest) GetInstances() []*structpb.Value {
	if x != nil {
		return x.Instances
	}
	return nil
}

func (x *PredictRequest) GetParameters() *structpb.Value {
	if x != nil {
		return x.Parameters
	}
	return nil
}

// Response message for [PredictionService.Predict][google.cloud.aiplatform.v1beta1.PredictionService.Predict].
type PredictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The predictions that are the output of the predictions call.
	// The schema of any single prediction may be specified via Endpoint's
	// DeployedModels' [Model's ][google.cloud.aiplatform.v1beta1.DeployedModel.model]
	// [PredictSchemata's][google.cloud.aiplatform.v1beta1.Model.predict_schemata]
	// [prediction_schema_uri][google.cloud.aiplatform.v1beta1.PredictSchemata.prediction_schema_uri].
	Predictions []*structpb.Value `protobuf:"bytes,1,rep,name=predictions,proto3" json:"predictions,omitempty"`
	// ID of the Endpoint's DeployedModel that served this prediction.
	DeployedModelId string `protobuf:"bytes,2,opt,name=deployed_model_id,json=deployedModelId,proto3" json:"deployed_model_id,omitempty"`
}

func (x *PredictResponse) Reset() {
	*x = PredictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_prediction_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictResponse) ProtoMessage() {}

func (x *PredictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_prediction_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictResponse.ProtoReflect.Descriptor instead.
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDescGZIP(), []int{1}
}

func (x *PredictResponse) GetPredictions() []*structpb.Value {
	if x != nil {
		return x.Predictions
	}
	return nil
}

func (x *PredictResponse) GetDeployedModelId() string {
	if x != nil {
		return x.DeployedModelId
	}
	return ""
}

// Request message for [PredictionService.Explain][google.cloud.aiplatform.v1beta1.PredictionService.Explain].
type ExplainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the Endpoint requested to serve the explanation.
	// Format:
	// `projects/{project}/locations/{location}/endpoints/{endpoint}`
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Required. The instances that are the input to the explanation call.
	// A DeployedModel may have an upper limit on the number of instances it
	// supports per request, and when it is exceeded the explanation call errors
	// in case of AutoML Models, or, in case of customer created Models, the
	// behaviour is as documented by that Model.
	// The schema of any single instance may be specified via Endpoint's
	// DeployedModels' [Model's][google.cloud.aiplatform.v1beta1.DeployedModel.model]
	// [PredictSchemata's][google.cloud.aiplatform.v1beta1.Model.predict_schemata]
	// [instance_schema_uri][google.cloud.aiplatform.v1beta1.PredictSchemata.instance_schema_uri].
	Instances []*structpb.Value `protobuf:"bytes,2,rep,name=instances,proto3" json:"instances,omitempty"`
	// The parameters that govern the prediction. The schema of the parameters may
	// be specified via Endpoint's DeployedModels' [Model's ][google.cloud.aiplatform.v1beta1.DeployedModel.model]
	// [PredictSchemata's][google.cloud.aiplatform.v1beta1.Model.predict_schemata]
	// [parameters_schema_uri][google.cloud.aiplatform.v1beta1.PredictSchemata.parameters_schema_uri].
	Parameters *structpb.Value `protobuf:"bytes,4,opt,name=parameters,proto3" json:"parameters,omitempty"`
	// If specified, this ExplainRequest will be served by the chosen
	// DeployedModel, overriding [Endpoint.traffic_split][google.cloud.aiplatform.v1beta1.Endpoint.traffic_split].
	DeployedModelId string `protobuf:"bytes,3,opt,name=deployed_model_id,json=deployedModelId,proto3" json:"deployed_model_id,omitempty"`
}

func (x *ExplainRequest) Reset() {
	*x = ExplainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_prediction_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplainRequest) ProtoMessage() {}

func (x *ExplainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_prediction_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplainRequest.ProtoReflect.Descriptor instead.
func (*ExplainRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDescGZIP(), []int{2}
}

func (x *ExplainRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *ExplainRequest) GetInstances() []*structpb.Value {
	if x != nil {
		return x.Instances
	}
	return nil
}

func (x *ExplainRequest) GetParameters() *structpb.Value {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *ExplainRequest) GetDeployedModelId() string {
	if x != nil {
		return x.DeployedModelId
	}
	return ""
}

// Response message for [PredictionService.Explain][google.cloud.aiplatform.v1beta1.PredictionService.Explain].
type ExplainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The explanations of the [Model's
	// predictions][PredictionResponse.predictions][].
	//
	// It has the same number of elements as [instances][google.cloud.aiplatform.v1beta1.ExplainRequest.instances]
	// to be explained.
	Explanations []*Explanation `protobuf:"bytes,1,rep,name=explanations,proto3" json:"explanations,omitempty"`
	// ID of the Endpoint's DeployedModel that served this explanation.
	DeployedModelId string `protobuf:"bytes,2,opt,name=deployed_model_id,json=deployedModelId,proto3" json:"deployed_model_id,omitempty"`
}

func (x *ExplainResponse) Reset() {
	*x = ExplainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_prediction_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplainResponse) ProtoMessage() {}

func (x *ExplainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_prediction_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplainResponse.ProtoReflect.Descriptor instead.
func (*ExplainResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDescGZIP(), []int{3}
}

func (x *ExplainResponse) GetExplanations() []*Explanation {
	if x != nil {
		return x.Explanations
	}
	return nil
}

func (x *ExplainResponse) GetDeployedModelId() string {
	if x != nil {
		return x.DeployedModelId
	}
	return ""
}

var File_google_cloud_aiplatform_v1beta1_prediction_service_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcb, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x24, 0x0a, 0x22, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22, 0x77, 0x0a,
	0x0f, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x22, 0xf7, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x24, 0x0a, 0x22, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x39, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64,
	0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0c, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x70,
	0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x49, 0x64, 0x32, 0xa8, 0x04, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xd7, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x12, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x43, 0x22,
	0x3e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x3a,
	0x01, 0x2a, 0xda, 0x41, 0x1d, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2c, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2c, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x12, 0xe9, 0x01, 0x0a, 0x07, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x12, 0x2f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x7b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x43, 0x22, 0x3e, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x7b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x2a,
	0x7d, 0x3a, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x2c, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x2c, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x1a, 0x4d,
	0xca, 0x41, 0x19, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x8a, 0x01,
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x16, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_prediction_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_aiplatform_v1beta1_prediction_service_proto_goTypes = []interface{}{
	(*PredictRequest)(nil),  // 0: google.cloud.aiplatform.v1beta1.PredictRequest
	(*PredictResponse)(nil), // 1: google.cloud.aiplatform.v1beta1.PredictResponse
	(*ExplainRequest)(nil),  // 2: google.cloud.aiplatform.v1beta1.ExplainRequest
	(*ExplainResponse)(nil), // 3: google.cloud.aiplatform.v1beta1.ExplainResponse
	(*structpb.Value)(nil),  // 4: google.protobuf.Value
	(*Explanation)(nil),     // 5: google.cloud.aiplatform.v1beta1.Explanation
}
var file_google_cloud_aiplatform_v1beta1_prediction_service_proto_depIdxs = []int32{
	4, // 0: google.cloud.aiplatform.v1beta1.PredictRequest.instances:type_name -> google.protobuf.Value
	4, // 1: google.cloud.aiplatform.v1beta1.PredictRequest.parameters:type_name -> google.protobuf.Value
	4, // 2: google.cloud.aiplatform.v1beta1.PredictResponse.predictions:type_name -> google.protobuf.Value
	4, // 3: google.cloud.aiplatform.v1beta1.ExplainRequest.instances:type_name -> google.protobuf.Value
	4, // 4: google.cloud.aiplatform.v1beta1.ExplainRequest.parameters:type_name -> google.protobuf.Value
	5, // 5: google.cloud.aiplatform.v1beta1.ExplainResponse.explanations:type_name -> google.cloud.aiplatform.v1beta1.Explanation
	0, // 6: google.cloud.aiplatform.v1beta1.PredictionService.Predict:input_type -> google.cloud.aiplatform.v1beta1.PredictRequest
	2, // 7: google.cloud.aiplatform.v1beta1.PredictionService.Explain:input_type -> google.cloud.aiplatform.v1beta1.ExplainRequest
	1, // 8: google.cloud.aiplatform.v1beta1.PredictionService.Predict:output_type -> google.cloud.aiplatform.v1beta1.PredictResponse
	3, // 9: google.cloud.aiplatform.v1beta1.PredictionService.Explain:output_type -> google.cloud.aiplatform.v1beta1.ExplainResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_prediction_service_proto_init() }
func file_google_cloud_aiplatform_v1beta1_prediction_service_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_prediction_service_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1beta1_explanation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_prediction_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictRequest); i {
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
		file_google_cloud_aiplatform_v1beta1_prediction_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictResponse); i {
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
		file_google_cloud_aiplatform_v1beta1_prediction_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExplainRequest); i {
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
		file_google_cloud_aiplatform_v1beta1_prediction_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExplainResponse); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_prediction_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_prediction_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_prediction_service_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_prediction_service_proto = out.File
	file_google_cloud_aiplatform_v1beta1_prediction_service_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_prediction_service_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_prediction_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PredictionServiceClient is the client API for PredictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictionServiceClient interface {
	// Perform an online prediction.
	Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error)
	// Perform an online explanation.
	//
	// If [ExplainRequest.deployed_model_id] is specified, the corresponding
	// DeployModel must have [explanation_spec][google.cloud.aiplatform.v1beta1.DeployedModel.explanation_spec]
	// populated. If [ExplainRequest.deployed_model_id] is not specified, all
	// DeployedModels must have [explanation_spec][google.cloud.aiplatform.v1beta1.DeployedModel.explanation_spec]
	// populated. Only deployed AutoML tabular Models have
	// explanation_spec.
	Explain(ctx context.Context, in *ExplainRequest, opts ...grpc.CallOption) (*ExplainResponse, error)
}

type predictionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPredictionServiceClient(cc grpc.ClientConnInterface) PredictionServiceClient {
	return &predictionServiceClient{cc}
}

func (c *predictionServiceClient) Predict(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictResponse, error) {
	out := new(PredictResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.aiplatform.v1beta1.PredictionService/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictionServiceClient) Explain(ctx context.Context, in *ExplainRequest, opts ...grpc.CallOption) (*ExplainResponse, error) {
	out := new(ExplainResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.aiplatform.v1beta1.PredictionService/Explain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictionServiceServer is the server API for PredictionService service.
type PredictionServiceServer interface {
	// Perform an online prediction.
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
	// Perform an online explanation.
	//
	// If [ExplainRequest.deployed_model_id] is specified, the corresponding
	// DeployModel must have [explanation_spec][google.cloud.aiplatform.v1beta1.DeployedModel.explanation_spec]
	// populated. If [ExplainRequest.deployed_model_id] is not specified, all
	// DeployedModels must have [explanation_spec][google.cloud.aiplatform.v1beta1.DeployedModel.explanation_spec]
	// populated. Only deployed AutoML tabular Models have
	// explanation_spec.
	Explain(context.Context, *ExplainRequest) (*ExplainResponse, error)
}

// UnimplementedPredictionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPredictionServiceServer struct {
}

func (*UnimplementedPredictionServiceServer) Predict(context.Context, *PredictRequest) (*PredictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}
func (*UnimplementedPredictionServiceServer) Explain(context.Context, *ExplainRequest) (*ExplainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Explain not implemented")
}

func RegisterPredictionServiceServer(s *grpc.Server, srv PredictionServiceServer) {
	s.RegisterService(&_PredictionService_serviceDesc, srv)
}

func _PredictionService_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.aiplatform.v1beta1.PredictionService/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Predict(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PredictionService_Explain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExplainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictionServiceServer).Explain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.aiplatform.v1beta1.PredictionService/Explain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictionServiceServer).Explain(ctx, req.(*ExplainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.aiplatform.v1beta1.PredictionService",
	HandlerType: (*PredictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _PredictionService_Predict_Handler,
		},
		{
			MethodName: "Explain",
			Handler:    _PredictionService_Explain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/aiplatform/v1beta1/prediction_service.proto",
}
