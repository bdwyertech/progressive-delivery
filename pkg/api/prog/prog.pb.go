// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api/prog/prog.proto

package api

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetVersionRequest) Reset() {
	*x = GetVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prog_prog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionRequest) ProtoMessage() {}

func (x *GetVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_prog_prog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionRequest.ProtoReflect.Descriptor instead.
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return file_api_prog_prog_proto_rawDescGZIP(), []int{0}
}

type GetVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetVersionResponse) Reset() {
	*x = GetVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prog_prog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionResponse) ProtoMessage() {}

func (x *GetVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_prog_prog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionResponse.ProtoReflect.Descriptor instead.
func (*GetVersionResponse) Descriptor() ([]byte, []int) {
	return file_api_prog_prog_proto_rawDescGZIP(), []int{1}
}

func (x *GetVersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ListCanariesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string      `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Pagination  *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListCanariesRequest) Reset() {
	*x = ListCanariesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prog_prog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCanariesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCanariesRequest) ProtoMessage() {}

func (x *ListCanariesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_prog_prog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCanariesRequest.ProtoReflect.Descriptor instead.
func (*ListCanariesRequest) Descriptor() ([]byte, []int) {
	return file_api_prog_prog_proto_rawDescGZIP(), []int{2}
}

func (x *ListCanariesRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ListCanariesRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListCanariesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Canaries      []*Canary    `protobuf:"bytes,1,rep,name=canaries,proto3" json:"canaries,omitempty"`
	NextPageToken string       `protobuf:"bytes,2,opt,name=nextPage_token,json=nextPageToken,proto3" json:"nextPage_token,omitempty"`
	Errors        []*ListError `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ListCanariesResponse) Reset() {
	*x = ListCanariesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_prog_prog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCanariesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCanariesResponse) ProtoMessage() {}

func (x *ListCanariesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_prog_prog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCanariesResponse.ProtoReflect.Descriptor instead.
func (*ListCanariesResponse) Descriptor() ([]byte, []int) {
	return file_api_prog_prog_proto_rawDescGZIP(), []int{3}
}

func (x *ListCanariesResponse) GetCanaries() []*Canary {
	if x != nil {
		return x.Canaries
	}
	return nil
}

func (x *ListCanariesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListCanariesResponse) GetErrors() []*ListError {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_api_prog_prog_proto protoreflect.FileDescriptor

var file_api_prog_prog_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2e,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x65,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x08, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x6e, 0x61, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x32, 0xbb,
	0x01, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x76, 0x65, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x47, 0x65,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x61, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x42, 0x30, 0x5a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x76,
	0x65, 0x2d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_prog_prog_proto_rawDescOnce sync.Once
	file_api_prog_prog_proto_rawDescData = file_api_prog_prog_proto_rawDesc
)

func file_api_prog_prog_proto_rawDescGZIP() []byte {
	file_api_prog_prog_proto_rawDescOnce.Do(func() {
		file_api_prog_prog_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_prog_prog_proto_rawDescData)
	})
	return file_api_prog_prog_proto_rawDescData
}

var file_api_prog_prog_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_prog_prog_proto_goTypes = []interface{}{
	(*GetVersionRequest)(nil),    // 0: GetVersionRequest
	(*GetVersionResponse)(nil),   // 1: GetVersionResponse
	(*ListCanariesRequest)(nil),  // 2: ListCanariesRequest
	(*ListCanariesResponse)(nil), // 3: ListCanariesResponse
	(*Pagination)(nil),           // 4: Pagination
	(*Canary)(nil),               // 5: Canary
	(*ListError)(nil),            // 6: ListError
}
var file_api_prog_prog_proto_depIdxs = []int32{
	4, // 0: ListCanariesRequest.pagination:type_name -> Pagination
	5, // 1: ListCanariesResponse.canaries:type_name -> Canary
	6, // 2: ListCanariesResponse.errors:type_name -> ListError
	0, // 3: ProgressiveDeliveryService.GetVersion:input_type -> GetVersionRequest
	2, // 4: ProgressiveDeliveryService.ListCanaries:input_type -> ListCanariesRequest
	1, // 5: ProgressiveDeliveryService.GetVersion:output_type -> GetVersionResponse
	3, // 6: ProgressiveDeliveryService.ListCanaries:output_type -> ListCanariesResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_prog_prog_proto_init() }
func file_api_prog_prog_proto_init() {
	if File_api_prog_prog_proto != nil {
		return
	}
	file_api_prog_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_prog_prog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVersionRequest); i {
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
		file_api_prog_prog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVersionResponse); i {
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
		file_api_prog_prog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCanariesRequest); i {
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
		file_api_prog_prog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCanariesResponse); i {
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
			RawDescriptor: file_api_prog_prog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_prog_prog_proto_goTypes,
		DependencyIndexes: file_api_prog_prog_proto_depIdxs,
		MessageInfos:      file_api_prog_prog_proto_msgTypes,
	}.Build()
	File_api_prog_prog_proto = out.File
	file_api_prog_prog_proto_rawDesc = nil
	file_api_prog_prog_proto_goTypes = nil
	file_api_prog_prog_proto_depIdxs = nil
}
