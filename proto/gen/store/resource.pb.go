// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: store/resource.proto

package store

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResourceStorageType int32

const (
	ResourceStorageType_RESOURCE_STORAGE_TYPE_UNSPECIFIED ResourceStorageType = 0
	ResourceStorageType_LOCAL                             ResourceStorageType = 1
	ResourceStorageType_S3                                ResourceStorageType = 2
	ResourceStorageType_EXTERNAL                          ResourceStorageType = 3
)

// Enum value maps for ResourceStorageType.
var (
	ResourceStorageType_name = map[int32]string{
		0: "RESOURCE_STORAGE_TYPE_UNSPECIFIED",
		1: "LOCAL",
		2: "S3",
		3: "EXTERNAL",
	}
	ResourceStorageType_value = map[string]int32{
		"RESOURCE_STORAGE_TYPE_UNSPECIFIED": 0,
		"LOCAL":                             1,
		"S3":                                2,
		"EXTERNAL":                          3,
	}
)

func (x ResourceStorageType) Enum() *ResourceStorageType {
	p := new(ResourceStorageType)
	*p = x
	return p
}

func (x ResourceStorageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceStorageType) Descriptor() protoreflect.EnumDescriptor {
	return file_store_resource_proto_enumTypes[0].Descriptor()
}

func (ResourceStorageType) Type() protoreflect.EnumType {
	return &file_store_resource_proto_enumTypes[0]
}

func (x ResourceStorageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceStorageType.Descriptor instead.
func (ResourceStorageType) EnumDescriptor() ([]byte, []int) {
	return file_store_resource_proto_rawDescGZIP(), []int{0}
}

type ResourcePayload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Payload:
	//
	//	*ResourcePayload_S3Object_
	Payload       isResourcePayload_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourcePayload) Reset() {
	*x = ResourcePayload{}
	mi := &file_store_resource_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourcePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcePayload) ProtoMessage() {}

func (x *ResourcePayload) ProtoReflect() protoreflect.Message {
	mi := &file_store_resource_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcePayload.ProtoReflect.Descriptor instead.
func (*ResourcePayload) Descriptor() ([]byte, []int) {
	return file_store_resource_proto_rawDescGZIP(), []int{0}
}

func (x *ResourcePayload) GetPayload() isResourcePayload_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ResourcePayload) GetS3Object() *ResourcePayload_S3Object {
	if x != nil {
		if x, ok := x.Payload.(*ResourcePayload_S3Object_); ok {
			return x.S3Object
		}
	}
	return nil
}

type isResourcePayload_Payload interface {
	isResourcePayload_Payload()
}

type ResourcePayload_S3Object_ struct {
	S3Object *ResourcePayload_S3Object `protobuf:"bytes,1,opt,name=s3_object,json=s3Object,proto3,oneof"`
}

func (*ResourcePayload_S3Object_) isResourcePayload_Payload() {}

type ResourcePayload_S3Object struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	S3Config *StorageS3Config       `protobuf:"bytes,1,opt,name=s3_config,json=s3Config,proto3" json:"s3_config,omitempty"`
	// key is the S3 object key.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// last_presigned_time is the last time the object was presigned.
	// This is used to determine if the presigned URL is still valid.
	LastPresignedTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_presigned_time,json=lastPresignedTime,proto3" json:"last_presigned_time,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ResourcePayload_S3Object) Reset() {
	*x = ResourcePayload_S3Object{}
	mi := &file_store_resource_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourcePayload_S3Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcePayload_S3Object) ProtoMessage() {}

func (x *ResourcePayload_S3Object) ProtoReflect() protoreflect.Message {
	mi := &file_store_resource_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcePayload_S3Object.ProtoReflect.Descriptor instead.
func (*ResourcePayload_S3Object) Descriptor() ([]byte, []int) {
	return file_store_resource_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ResourcePayload_S3Object) GetS3Config() *StorageS3Config {
	if x != nil {
		return x.S3Config
	}
	return nil
}

func (x *ResourcePayload_S3Object) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ResourcePayload_S3Object) GetLastPresignedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastPresignedTime
	}
	return nil
}

var File_store_resource_proto protoreflect.FileDescriptor

var file_store_resource_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x44, 0x0a, 0x09, 0x73, 0x33, 0x5f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x53, 0x33, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x48, 0x00, 0x52, 0x08, 0x73, 0x33, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0xa3, 0x01,
	0x0a, 0x08, 0x53, 0x33, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x73, 0x33,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x53, 0x33, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x73, 0x33, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4a, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x70, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x5d,
	0x0a, 0x13, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x21, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x33, 0x10, 0x02, 0x12,
	0x0c, 0x0a, 0x08, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x03, 0x42, 0x98, 0x01,
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x42, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75,
	0x73, 0x65, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03,
	0x4d, 0x53, 0x58, 0xaa, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0xca, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xe2,
	0x02, 0x17, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x4d, 0x65, 0x6d, 0x6f,
	0x73, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_resource_proto_rawDescOnce sync.Once
	file_store_resource_proto_rawDescData = file_store_resource_proto_rawDesc
)

func file_store_resource_proto_rawDescGZIP() []byte {
	file_store_resource_proto_rawDescOnce.Do(func() {
		file_store_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_resource_proto_rawDescData)
	})
	return file_store_resource_proto_rawDescData
}

var file_store_resource_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_store_resource_proto_goTypes = []any{
	(ResourceStorageType)(0),         // 0: memos.store.ResourceStorageType
	(*ResourcePayload)(nil),          // 1: memos.store.ResourcePayload
	(*ResourcePayload_S3Object)(nil), // 2: memos.store.ResourcePayload.S3Object
	(*StorageS3Config)(nil),          // 3: memos.store.StorageS3Config
	(*timestamppb.Timestamp)(nil),    // 4: google.protobuf.Timestamp
}
var file_store_resource_proto_depIdxs = []int32{
	2, // 0: memos.store.ResourcePayload.s3_object:type_name -> memos.store.ResourcePayload.S3Object
	3, // 1: memos.store.ResourcePayload.S3Object.s3_config:type_name -> memos.store.StorageS3Config
	4, // 2: memos.store.ResourcePayload.S3Object.last_presigned_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_store_resource_proto_init() }
func file_store_resource_proto_init() {
	if File_store_resource_proto != nil {
		return
	}
	file_store_workspace_setting_proto_init()
	file_store_resource_proto_msgTypes[0].OneofWrappers = []any{
		(*ResourcePayload_S3Object_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_resource_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_resource_proto_goTypes,
		DependencyIndexes: file_store_resource_proto_depIdxs,
		EnumInfos:         file_store_resource_proto_enumTypes,
		MessageInfos:      file_store_resource_proto_msgTypes,
	}.Build()
	File_store_resource_proto = out.File
	file_store_resource_proto_rawDesc = nil
	file_store_resource_proto_goTypes = nil
	file_store_resource_proto_depIdxs = nil
}
