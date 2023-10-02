// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: tensorflow/core/protobuf/critical_section.proto

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

// Protocol buffer representing a CriticalSection.
type CriticalSectionDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the critical section handle.
	CriticalSectionName string `protobuf:"bytes,1,opt,name=critical_section_name,json=criticalSectionName,proto3" json:"critical_section_name,omitempty"`
}

func (x *CriticalSectionDef) Reset() {
	*x = CriticalSectionDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_critical_section_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CriticalSectionDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CriticalSectionDef) ProtoMessage() {}

func (x *CriticalSectionDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_critical_section_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CriticalSectionDef.ProtoReflect.Descriptor instead.
func (*CriticalSectionDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_critical_section_proto_rawDescGZIP(), []int{0}
}

func (x *CriticalSectionDef) GetCriticalSectionName() string {
	if x != nil {
		return x.CriticalSectionName
	}
	return ""
}

// Protocol buffer representing a CriticalSection execution.
type CriticalSectionExecutionDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the critical section handle.
	ExecuteInCriticalSectionName string `protobuf:"bytes,1,opt,name=execute_in_critical_section_name,json=executeInCriticalSectionName,proto3" json:"execute_in_critical_section_name,omitempty"`
	// Whether this operation requires exclusive access to its resources,
	// (i.e., no other CriticalSections may request the same resources).
	ExclusiveResourceAccess bool `protobuf:"varint,2,opt,name=exclusive_resource_access,json=exclusiveResourceAccess,proto3" json:"exclusive_resource_access,omitempty"`
}

func (x *CriticalSectionExecutionDef) Reset() {
	*x = CriticalSectionExecutionDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_critical_section_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CriticalSectionExecutionDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CriticalSectionExecutionDef) ProtoMessage() {}

func (x *CriticalSectionExecutionDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_critical_section_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CriticalSectionExecutionDef.ProtoReflect.Descriptor instead.
func (*CriticalSectionExecutionDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_critical_section_proto_rawDescGZIP(), []int{1}
}

func (x *CriticalSectionExecutionDef) GetExecuteInCriticalSectionName() string {
	if x != nil {
		return x.ExecuteInCriticalSectionName
	}
	return ""
}

func (x *CriticalSectionExecutionDef) GetExclusiveResourceAccess() bool {
	if x != nil {
		return x.ExclusiveResourceAccess
	}
	return false
}

var File_tensorflow_core_protobuf_critical_section_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_critical_section_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x72, 0x69, 0x74, 0x69,
	0x63, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x48, 0x0a,
	0x12, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x66, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x69, 0x74,
	0x69, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x12, 0x46, 0x0a, 0x20, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x1c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x43, 0x72, 0x69, 0x74,
	0x69, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x3a, 0x0a, 0x19, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x17, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x8d, 0x01, 0x0a, 0x18,
	0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x15, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50,
	0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67,
	0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f,
	0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_critical_section_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_critical_section_proto_rawDescData = file_tensorflow_core_protobuf_critical_section_proto_rawDesc
)

func file_tensorflow_core_protobuf_critical_section_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_critical_section_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_critical_section_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_critical_section_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_critical_section_proto_rawDescData
}

var file_tensorflow_core_protobuf_critical_section_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_core_protobuf_critical_section_proto_goTypes = []interface{}{
	(*CriticalSectionDef)(nil),          // 0: tensorflow.CriticalSectionDef
	(*CriticalSectionExecutionDef)(nil), // 1: tensorflow.CriticalSectionExecutionDef
}
var file_tensorflow_core_protobuf_critical_section_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_critical_section_proto_init() }
func file_tensorflow_core_protobuf_critical_section_proto_init() {
	if File_tensorflow_core_protobuf_critical_section_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_critical_section_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CriticalSectionDef); i {
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
		file_tensorflow_core_protobuf_critical_section_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CriticalSectionExecutionDef); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_critical_section_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_critical_section_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_critical_section_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_critical_section_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_critical_section_proto = out.File
	file_tensorflow_core_protobuf_critical_section_proto_rawDesc = nil
	file_tensorflow_core_protobuf_critical_section_proto_goTypes = nil
	file_tensorflow_core_protobuf_critical_section_proto_depIdxs = nil
}
