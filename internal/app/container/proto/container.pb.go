// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: container.proto

package container

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

type SignatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignatureRequest) Reset() {
	*x = SignatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureRequest) ProtoMessage() {}

func (x *SignatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_container_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignatureRequest.ProtoReflect.Descriptor instead.
func (*SignatureRequest) Descriptor() ([]byte, []int) {
	return file_container_proto_rawDescGZIP(), []int{0}
}

func (x *SignatureRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type FunctionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Functions []*Function `protobuf:"bytes,1,rep,name=functions,proto3" json:"functions,omitempty"`
}

func (x *FunctionsResponse) Reset() {
	*x = FunctionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionsResponse) ProtoMessage() {}

func (x *FunctionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_container_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionsResponse.ProtoReflect.Descriptor instead.
func (*FunctionsResponse) Descriptor() ([]byte, []int) {
	return file_container_proto_rawDescGZIP(), []int{1}
}

func (x *FunctionsResponse) GetFunctions() []*Function {
	if x != nil {
		return x.Functions
	}
	return nil
}

type Function struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FunctionName      string `protobuf:"bytes,1,opt,name=functionName,proto3" json:"functionName,omitempty"`
	FunctionSignature string `protobuf:"bytes,2,opt,name=functionSignature,proto3" json:"functionSignature,omitempty"`
	FunctionComment   string `protobuf:"bytes,3,opt,name=functionComment,proto3" json:"functionComment,omitempty"`
	FileName          string `protobuf:"bytes,4,opt,name=fileName,proto3" json:"fileName,omitempty"`
	PackageName       string `protobuf:"bytes,5,opt,name=packageName,proto3" json:"packageName,omitempty"`
	PackageLink       string `protobuf:"bytes,6,opt,name=packageLink,proto3" json:"packageLink,omitempty"`
}

func (x *Function) Reset() {
	*x = Function{}
	if protoimpl.UnsafeEnabled {
		mi := &file_container_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Function) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Function) ProtoMessage() {}

func (x *Function) ProtoReflect() protoreflect.Message {
	mi := &file_container_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Function.ProtoReflect.Descriptor instead.
func (*Function) Descriptor() ([]byte, []int) {
	return file_container_proto_rawDescGZIP(), []int{2}
}

func (x *Function) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *Function) GetFunctionSignature() string {
	if x != nil {
		return x.FunctionSignature
	}
	return ""
}

func (x *Function) GetFunctionComment() string {
	if x != nil {
		return x.FunctionComment
	}
	return ""
}

func (x *Function) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Function) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *Function) GetPackageLink() string {
	if x != nil {
		return x.PackageLink
	}
	return ""
}

var File_container_proto protoreflect.FileDescriptor

var file_container_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x10,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x46,
	0x0a, 0x11, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xe6, 0x01, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x32,
	0x50, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x04,
	0x46, 0x69, 0x6e, 0x64, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_container_proto_rawDescOnce sync.Once
	file_container_proto_rawDescData = file_container_proto_rawDesc
)

func file_container_proto_rawDescGZIP() []byte {
	file_container_proto_rawDescOnce.Do(func() {
		file_container_proto_rawDescData = protoimpl.X.CompressGZIP(file_container_proto_rawDescData)
	})
	return file_container_proto_rawDescData
}

var file_container_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_container_proto_goTypes = []any{
	(*SignatureRequest)(nil),  // 0: container.SignatureRequest
	(*FunctionsResponse)(nil), // 1: container.FunctionsResponse
	(*Function)(nil),          // 2: container.Function
}
var file_container_proto_depIdxs = []int32{
	2, // 0: container.FunctionsResponse.functions:type_name -> container.Function
	0, // 1: container.Container.Find:input_type -> container.SignatureRequest
	1, // 2: container.Container.Find:output_type -> container.FunctionsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_container_proto_init() }
func file_container_proto_init() {
	if File_container_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_container_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SignatureRequest); i {
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
		file_container_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FunctionsResponse); i {
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
		file_container_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Function); i {
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
			RawDescriptor: file_container_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_container_proto_goTypes,
		DependencyIndexes: file_container_proto_depIdxs,
		MessageInfos:      file_container_proto_msgTypes,
	}.Build()
	File_container_proto = out.File
	file_container_proto_rawDesc = nil
	file_container_proto_goTypes = nil
	file_container_proto_depIdxs = nil
}
