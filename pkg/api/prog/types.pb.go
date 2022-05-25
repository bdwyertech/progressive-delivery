// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api/prog/types.proto

package api

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

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  int32  `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prog_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_api_prog_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_api_prog_types_proto_rawDescGZIP(), []int{0}
}

func (x *Pagination) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Pagination) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Namespace   string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Message     string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ListError) Reset() {
	*x = ListError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prog_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListError) ProtoMessage() {}

func (x *ListError) ProtoReflect() protoreflect.Message {
	mi := &file_api_prog_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListError.ProtoReflect.Descriptor instead.
func (*ListError) Descriptor() ([]byte, []int) {
	return file_api_prog_types_proto_rawDescGZIP(), []int{1}
}

func (x *ListError) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ListError) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Canary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ClusterName      string                  `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Provider         string                  `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	TargetReference  *CanaryTargetReference  `protobuf:"bytes,4,opt,name=target_reference,json=targetReference,proto3" json:"target_reference,omitempty"`
	TargetDeployment *CanaryTargetDeployment `protobuf:"bytes,5,opt,name=target_deployment,json=targetDeployment,proto3" json:"target_deployment,omitempty"`
	Status           *CanaryStatus           `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Canary) Reset() {
	*x = Canary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prog_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Canary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Canary) ProtoMessage() {}

func (x *Canary) ProtoReflect() protoreflect.Message {
	mi := &file_api_prog_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Canary.ProtoReflect.Descriptor instead.
func (*Canary) Descriptor() ([]byte, []int) {
	return file_api_prog_types_proto_rawDescGZIP(), []int{2}
}

func (x *Canary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Canary) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *Canary) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Canary) GetTargetReference() *CanaryTargetReference {
	if x != nil {
		return x.TargetReference
	}
	return nil
}

func (x *Canary) GetTargetDeployment() *CanaryTargetDeployment {
	if x != nil {
		return x.TargetDeployment
	}
	return nil
}

func (x *Canary) GetStatus() *CanaryStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type CanaryTargetReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CanaryTargetReference) Reset() {
	*x = CanaryTargetReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prog_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanaryTargetReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanaryTargetReference) ProtoMessage() {}

func (x *CanaryTargetReference) ProtoReflect() protoreflect.Message {
	mi := &file_api_prog_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanaryTargetReference.ProtoReflect.Descriptor instead.
func (*CanaryTargetReference) Descriptor() ([]byte, []int) {
	return file_api_prog_types_proto_rawDescGZIP(), []int{3}
}

func (x *CanaryTargetReference) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *CanaryTargetReference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CanaryStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phase              string             `protobuf:"bytes,1,opt,name=phase,proto3" json:"phase,omitempty"`
	FailedChecks       int32              `protobuf:"varint,2,opt,name=failed_checks,json=failedChecks,proto3" json:"failed_checks,omitempty"`
	CanaryWeight       int32              `protobuf:"varint,3,opt,name=canary_weight,json=canaryWeight,proto3" json:"canary_weight,omitempty"`
	Iterations         int32              `protobuf:"varint,4,opt,name=iterations,proto3" json:"iterations,omitempty"`
	LastTransitionTime string             `protobuf:"bytes,5,opt,name=lastTransition_time,json=lastTransitionTime,proto3" json:"lastTransition_time,omitempty"`
	Conditions         []*CanaryCondition `protobuf:"bytes,6,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *CanaryStatus) Reset() {
	*x = CanaryStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prog_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanaryStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanaryStatus) ProtoMessage() {}

func (x *CanaryStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_prog_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanaryStatus.ProtoReflect.Descriptor instead.
func (*CanaryStatus) Descriptor() ([]byte, []int) {
	return file_api_prog_types_proto_rawDescGZIP(), []int{4}
}

func (x *CanaryStatus) GetPhase() string {
	if x != nil {
		return x.Phase
	}
	return ""
}

func (x *CanaryStatus) GetFailedChecks() int32 {
	if x != nil {
		return x.FailedChecks
	}
	return 0
}

func (x *CanaryStatus) GetCanaryWeight() int32 {
	if x != nil {
		return x.CanaryWeight
	}
	return 0
}

func (x *CanaryStatus) GetIterations() int32 {
	if x != nil {
		return x.Iterations
	}
	return 0
}

func (x *CanaryStatus) GetLastTransitionTime() string {
	if x != nil {
		return x.LastTransitionTime
	}
	return ""
}

func (x *CanaryStatus) GetConditions() []*CanaryCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

type CanaryCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type               string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Status             string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	LastUpdateTime     string `protobuf:"bytes,3,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
	LastTransitionTime string `protobuf:"bytes,4,opt,name=last_transition_time,json=lastTransitionTime,proto3" json:"last_transition_time,omitempty"`
	Reason             string `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	Message            string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CanaryCondition) Reset() {
	*x = CanaryCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prog_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanaryCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanaryCondition) ProtoMessage() {}

func (x *CanaryCondition) ProtoReflect() protoreflect.Message {
	mi := &file_api_prog_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanaryCondition.ProtoReflect.Descriptor instead.
func (*CanaryCondition) Descriptor() ([]byte, []int) {
	return file_api_prog_types_proto_rawDescGZIP(), []int{5}
}

func (x *CanaryCondition) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CanaryCondition) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CanaryCondition) GetLastUpdateTime() string {
	if x != nil {
		return x.LastUpdateTime
	}
	return ""
}

func (x *CanaryCondition) GetLastTransitionTime() string {
	if x != nil {
		return x.LastTransitionTime
	}
	return ""
}

func (x *CanaryCondition) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *CanaryCondition) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CanaryTargetDeployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid             string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ResourceVersion string `protobuf:"bytes,2,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty"`
}

func (x *CanaryTargetDeployment) Reset() {
	*x = CanaryTargetDeployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prog_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanaryTargetDeployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanaryTargetDeployment) ProtoMessage() {}

func (x *CanaryTargetDeployment) ProtoReflect() protoreflect.Message {
	mi := &file_api_prog_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanaryTargetDeployment.ProtoReflect.Descriptor instead.
func (*CanaryTargetDeployment) Descriptor() ([]byte, []int) {
	return file_api_prog_types_proto_rawDescGZIP(), []int{6}
}

func (x *CanaryTargetDeployment) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *CanaryTargetDeployment) GetResourceVersion() string {
	if x != nil {
		return x.ResourceVersion
	}
	return ""
}

var File_api_prog_types_proto protoreflect.FileDescriptor

var file_api_prog_types_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x66, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8b, 0x02, 0x0a, 0x06, 0x43, 0x61, 0x6e,
	0x61, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x11, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x10, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3f, 0x0a, 0x15, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x79,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf1, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x61,
	0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x61, 0x6e, 0x61,
	0x72, 0x79, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x74,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x43, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xcb, 0x01, 0x0a, 0x0f,
	0x43, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x55, 0x0a, 0x16, 0x43, 0x61, 0x6e,
	0x61, 0x72, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x65, 0x61, 0x76, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x76, 0x65, 0x2d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_prog_types_proto_rawDescOnce sync.Once
	file_api_prog_types_proto_rawDescData = file_api_prog_types_proto_rawDesc
)

func file_api_prog_types_proto_rawDescGZIP() []byte {
	file_api_prog_types_proto_rawDescOnce.Do(func() {
		file_api_prog_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_prog_types_proto_rawDescData)
	})
	return file_api_prog_types_proto_rawDescData
}

var file_api_prog_types_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_prog_types_proto_goTypes = []interface{}{
	(*Pagination)(nil),             // 0: Pagination
	(*ListError)(nil),              // 1: ListError
	(*Canary)(nil),                 // 2: Canary
	(*CanaryTargetReference)(nil),  // 3: CanaryTargetReference
	(*CanaryStatus)(nil),           // 4: CanaryStatus
	(*CanaryCondition)(nil),        // 5: CanaryCondition
	(*CanaryTargetDeployment)(nil), // 6: CanaryTargetDeployment
}
var file_api_prog_types_proto_depIdxs = []int32{
	3, // 0: Canary.target_reference:type_name -> CanaryTargetReference
	6, // 1: Canary.target_deployment:type_name -> CanaryTargetDeployment
	4, // 2: Canary.status:type_name -> CanaryStatus
	5, // 3: CanaryStatus.conditions:type_name -> CanaryCondition
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_prog_types_proto_init() }
func file_api_prog_types_proto_init() {
	if File_api_prog_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_prog_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_api_prog_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListError); i {
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
		file_api_prog_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Canary); i {
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
		file_api_prog_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CanaryTargetReference); i {
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
		file_api_prog_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CanaryStatus); i {
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
		file_api_prog_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CanaryCondition); i {
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
		file_api_prog_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CanaryTargetDeployment); i {
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
			RawDescriptor: file_api_prog_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_prog_types_proto_goTypes,
		DependencyIndexes: file_api_prog_types_proto_depIdxs,
		MessageInfos:      file_api_prog_types_proto_msgTypes,
	}.Build()
	File_api_prog_types_proto = out.File
	file_api_prog_types_proto_rawDesc = nil
	file_api_prog_types_proto_goTypes = nil
	file_api_prog_types_proto_depIdxs = nil
}
