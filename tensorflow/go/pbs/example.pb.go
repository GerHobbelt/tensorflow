// Protocol messages for describing input data Examples for machine learning
// model training or inference.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: tensorflow/core/example/example.proto

package pbs

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

type Example struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Features *Features `protobuf:"bytes,1,opt,name=features,proto3" json:"features,omitempty"`
}

func (x *Example) Reset() {
	*x = Example{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_example_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example) ProtoMessage() {}

func (x *Example) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_example_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example.ProtoReflect.Descriptor instead.
func (*Example) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_example_example_proto_rawDescGZIP(), []int{0}
}

func (x *Example) GetFeatures() *Features {
	if x != nil {
		return x.Features
	}
	return nil
}

type SequenceExample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context      *Features     `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	FeatureLists *FeatureLists `protobuf:"bytes,2,opt,name=feature_lists,json=featureLists,proto3" json:"feature_lists,omitempty"`
}

func (x *SequenceExample) Reset() {
	*x = SequenceExample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_example_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SequenceExample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SequenceExample) ProtoMessage() {}

func (x *SequenceExample) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_example_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SequenceExample.ProtoReflect.Descriptor instead.
func (*SequenceExample) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_example_example_proto_rawDescGZIP(), []int{1}
}

func (x *SequenceExample) GetContext() *Features {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *SequenceExample) GetFeatureLists() *FeatureLists {
	if x != nil {
		return x.FeatureLists
	}
	return nil
}

var File_tensorflow_core_example_example_proto protoreflect.FileDescriptor

var file_tensorflow_core_example_example_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x07, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x08, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3d, 0x0a, 0x0d, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x0c, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x42, 0x81, 0x01, 0x0a, 0x16, 0x6f,
	0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x0d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_example_example_proto_rawDescOnce sync.Once
	file_tensorflow_core_example_example_proto_rawDescData = file_tensorflow_core_example_example_proto_rawDesc
)

func file_tensorflow_core_example_example_proto_rawDescGZIP() []byte {
	file_tensorflow_core_example_example_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_example_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_example_example_proto_rawDescData)
	})
	return file_tensorflow_core_example_example_proto_rawDescData
}

var file_tensorflow_core_example_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_core_example_example_proto_goTypes = []interface{}{
	(*Example)(nil),         // 0: tensorflow.Example
	(*SequenceExample)(nil), // 1: tensorflow.SequenceExample
	(*Features)(nil),        // 2: tensorflow.Features
	(*FeatureLists)(nil),    // 3: tensorflow.FeatureLists
}
var file_tensorflow_core_example_example_proto_depIdxs = []int32{
	2, // 0: tensorflow.Example.features:type_name -> tensorflow.Features
	2, // 1: tensorflow.SequenceExample.context:type_name -> tensorflow.Features
	3, // 2: tensorflow.SequenceExample.feature_lists:type_name -> tensorflow.FeatureLists
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_core_example_example_proto_init() }
func file_tensorflow_core_example_example_proto_init() {
	if File_tensorflow_core_example_example_proto != nil {
		return
	}
	file_tensorflow_core_example_feature_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_example_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Example); i {
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
		file_tensorflow_core_example_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SequenceExample); i {
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
			RawDescriptor: file_tensorflow_core_example_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_example_example_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_example_example_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_example_example_proto_msgTypes,
	}.Build()
	File_tensorflow_core_example_example_proto = out.File
	file_tensorflow_core_example_example_proto_rawDesc = nil
	file_tensorflow_core_example_example_proto_goTypes = nil
	file_tensorflow_core_example_example_proto_depIdxs = nil
}
