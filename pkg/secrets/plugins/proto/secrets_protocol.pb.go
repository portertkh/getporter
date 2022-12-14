// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: pkg/secrets/plugins/proto/secrets_protocol.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResolveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyName  string `protobuf:"bytes,1,opt,name=KeyName,proto3" json:"KeyName,omitempty"`
	KeyValue string `protobuf:"bytes,2,opt,name=KeyValue,proto3" json:"KeyValue,omitempty"`
}

func (x *ResolveRequest) Reset() {
	*x = ResolveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_plugins_proto_secrets_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveRequest) ProtoMessage() {}

func (x *ResolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_plugins_proto_secrets_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveRequest.ProtoReflect.Descriptor instead.
func (*ResolveRequest) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *ResolveRequest) GetKeyName() string {
	if x != nil {
		return x.KeyName
	}
	return ""
}

func (x *ResolveRequest) GetKeyValue() string {
	if x != nil {
		return x.KeyValue
	}
	return ""
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyName  string `protobuf:"bytes,1,opt,name=KeyName,proto3" json:"KeyName,omitempty"`
	KeyValue string `protobuf:"bytes,2,opt,name=KeyValue,proto3" json:"KeyValue,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_plugins_proto_secrets_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_plugins_proto_secrets_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetKeyName() string {
	if x != nil {
		return x.KeyName
	}
	return ""
}

func (x *CreateRequest) GetKeyValue() string {
	if x != nil {
		return x.KeyValue
	}
	return ""
}

func (x *CreateRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ResolveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *ResolveResponse) Reset() {
	*x = ResolveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_plugins_proto_secrets_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveResponse) ProtoMessage() {}

func (x *ResolveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_plugins_proto_secrets_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveResponse.ProtoReflect.Descriptor instead.
func (*ResolveResponse) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *ResolveResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_secrets_plugins_proto_secrets_protocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_secrets_plugins_proto_secrets_protocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDescGZIP(), []int{3}
}

var File_pkg_secrets_plugins_proto_secrets_protocol_proto protoreflect.FileDescriptor

var file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDesc = []byte{
	0x0a, 0x30, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x22, 0x46, 0x0a, 0x0e, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x5b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x27, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8a, 0x01, 0x0a, 0x0f,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x3c, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x65, 0x74, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x68, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDescOnce sync.Once
	file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDescData = file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDesc
)

func file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDescGZIP() []byte {
	file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDescOnce.Do(func() {
		file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDescData)
	})
	return file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDescData
}

var file_pkg_secrets_plugins_proto_secrets_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_secrets_plugins_proto_secrets_protocol_proto_goTypes = []interface{}{
	(*ResolveRequest)(nil),  // 0: plugins.ResolveRequest
	(*CreateRequest)(nil),   // 1: plugins.CreateRequest
	(*ResolveResponse)(nil), // 2: plugins.ResolveResponse
	(*CreateResponse)(nil),  // 3: plugins.CreateResponse
}
var file_pkg_secrets_plugins_proto_secrets_protocol_proto_depIdxs = []int32{
	0, // 0: plugins.SecretsProtocol.Resolve:input_type -> plugins.ResolveRequest
	1, // 1: plugins.SecretsProtocol.Create:input_type -> plugins.CreateRequest
	2, // 2: plugins.SecretsProtocol.Resolve:output_type -> plugins.ResolveResponse
	3, // 3: plugins.SecretsProtocol.Create:output_type -> plugins.CreateResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_secrets_plugins_proto_secrets_protocol_proto_init() }
func file_pkg_secrets_plugins_proto_secrets_protocol_proto_init() {
	if File_pkg_secrets_plugins_proto_secrets_protocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_secrets_plugins_proto_secrets_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveRequest); i {
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
		file_pkg_secrets_plugins_proto_secrets_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_pkg_secrets_plugins_proto_secrets_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveResponse); i {
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
		file_pkg_secrets_plugins_proto_secrets_protocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
			RawDescriptor: file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_secrets_plugins_proto_secrets_protocol_proto_goTypes,
		DependencyIndexes: file_pkg_secrets_plugins_proto_secrets_protocol_proto_depIdxs,
		MessageInfos:      file_pkg_secrets_plugins_proto_secrets_protocol_proto_msgTypes,
	}.Build()
	File_pkg_secrets_plugins_proto_secrets_protocol_proto = out.File
	file_pkg_secrets_plugins_proto_secrets_protocol_proto_rawDesc = nil
	file_pkg_secrets_plugins_proto_secrets_protocol_proto_goTypes = nil
	file_pkg_secrets_plugins_proto_secrets_protocol_proto_depIdxs = nil
}
