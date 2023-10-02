// This is used for convolution logging. Also see
// tensorflow/core/protobuf/autotuing.h

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: tensorflow/core/protobuf/conv_autotuning.proto

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

// A convolution. Currently it's only used for logging. In the future, we may
// want to use it in the API as well.
type ConvolutionProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind     ConvolutionKind             `protobuf:"varint,1,opt,name=kind,proto3,enum=dnn.ConvolutionKind" json:"kind,omitempty"`
	Input    *TensorDescriptorProto      `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	Filter   *TensorDescriptorProto      `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	Output   *TensorDescriptorProto      `protobuf:"bytes,4,opt,name=output,proto3" json:"output,omitempty"`
	ConvDesc *ConvolutionDescriptorProto `protobuf:"bytes,5,opt,name=conv_desc,json=convDesc,proto3" json:"conv_desc,omitempty"`
	// result = conv_scale * conv(...) + side_value_scale * side_value.
	// side_value is an arbitrary buffer if activation is not none. Otherwise, it
	// has to be the result buffer (using its old values).
	ConvScale        float64                        `protobuf:"fixed64,6,opt,name=conv_scale,json=convScale,proto3" json:"conv_scale,omitempty"`
	SideValueScale   float64                        `protobuf:"fixed64,7,opt,name=side_value_scale,json=sideValueScale,proto3" json:"side_value_scale,omitempty"`
	Activation       ActivationMode `protobuf:"varint,8,opt,name=activation,proto3,enum=dnn.ActivationMode" json:"activation,omitempty"`
	InputAddress     int64                          `protobuf:"varint,9,opt,name=input_address,json=inputAddress,proto3" json:"input_address,omitempty"`
	FilterAddress    int64                          `protobuf:"varint,10,opt,name=filter_address,json=filterAddress,proto3" json:"filter_address,omitempty"`
	OutputAddress    int64                          `protobuf:"varint,11,opt,name=output_address,json=outputAddress,proto3" json:"output_address,omitempty"`
	BiasAddress      int64                          `protobuf:"varint,12,opt,name=bias_address,json=biasAddress,proto3" json:"bias_address,omitempty"`
	SideInputAddress int64                          `protobuf:"varint,13,opt,name=side_input_address,json=sideInputAddress,proto3" json:"side_input_address,omitempty"`
}

func (x *ConvolutionProto) Reset() {
	*x = ConvolutionProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_conv_autotuning_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvolutionProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvolutionProto) ProtoMessage() {}

func (x *ConvolutionProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_conv_autotuning_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvolutionProto.ProtoReflect.Descriptor instead.
func (*ConvolutionProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_conv_autotuning_proto_rawDescGZIP(), []int{0}
}

func (x *ConvolutionProto) GetKind() ConvolutionKind {
	if x != nil {
		return x.Kind
	}
	return ConvolutionKind(0)
}

func (x *ConvolutionProto) GetInput() *TensorDescriptorProto {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *ConvolutionProto) GetFilter() *TensorDescriptorProto {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ConvolutionProto) GetOutput() *TensorDescriptorProto {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *ConvolutionProto) GetConvDesc() *ConvolutionDescriptorProto {
	if x != nil {
		return x.ConvDesc
	}
	return nil
}

func (x *ConvolutionProto) GetConvScale() float64 {
	if x != nil {
		return x.ConvScale
	}
	return 0
}

func (x *ConvolutionProto) GetSideValueScale() float64 {
	if x != nil {
		return x.SideValueScale
	}
	return 0
}

func (x *ConvolutionProto) GetActivation() ActivationMode {
	if x != nil {
		return x.Activation
	}
	return ActivationMode(0)
}

func (x *ConvolutionProto) GetInputAddress() int64 {
	if x != nil {
		return x.InputAddress
	}
	return 0
}

func (x *ConvolutionProto) GetFilterAddress() int64 {
	if x != nil {
		return x.FilterAddress
	}
	return 0
}

func (x *ConvolutionProto) GetOutputAddress() int64 {
	if x != nil {
		return x.OutputAddress
	}
	return 0
}

func (x *ConvolutionProto) GetBiasAddress() int64 {
	if x != nil {
		return x.BiasAddress
	}
	return 0
}

func (x *ConvolutionProto) GetSideInputAddress() int64 {
	if x != nil {
		return x.SideInputAddress
	}
	return 0
}

type MatmulProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbDtype     DataType       `protobuf:"varint,1,opt,name=ab_dtype,json=abDtype,proto3,enum=dnn.DataType" json:"ab_dtype,omitempty"`
	CDtype      DataType       `protobuf:"varint,2,opt,name=c_dtype,json=cDtype,proto3,enum=dnn.DataType" json:"c_dtype,omitempty"`
	TransA      bool                           `protobuf:"varint,3,opt,name=trans_a,json=transA,proto3" json:"trans_a,omitempty"`
	TransB      bool                           `protobuf:"varint,4,opt,name=trans_b,json=transB,proto3" json:"trans_b,omitempty"`
	M           uint64                         `protobuf:"varint,5,opt,name=m,proto3" json:"m,omitempty"`
	N           uint64                         `protobuf:"varint,6,opt,name=n,proto3" json:"n,omitempty"`
	K           uint64                         `protobuf:"varint,7,opt,name=k,proto3" json:"k,omitempty"`
	Lda         int64                          `protobuf:"varint,8,opt,name=lda,proto3" json:"lda,omitempty"`
	Ldb         int64                          `protobuf:"varint,9,opt,name=ldb,proto3" json:"ldb,omitempty"`
	Ldc         int64                          `protobuf:"varint,10,opt,name=ldc,proto3" json:"ldc,omitempty"`
	Activation  ActivationMode `protobuf:"varint,11,opt,name=activation,proto3,enum=dnn.ActivationMode" json:"activation,omitempty"`
	AAddress    int64                          `protobuf:"varint,12,opt,name=a_address,json=aAddress,proto3" json:"a_address,omitempty"`
	BAddress    int64                          `protobuf:"varint,13,opt,name=b_address,json=bAddress,proto3" json:"b_address,omitempty"`
	CAddress    int64                          `protobuf:"varint,14,opt,name=c_address,json=cAddress,proto3" json:"c_address,omitempty"`
	BiasAddress int64                          `protobuf:"varint,15,opt,name=bias_address,json=biasAddress,proto3" json:"bias_address,omitempty"`
}

func (x *MatmulProto) Reset() {
	*x = MatmulProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_conv_autotuning_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatmulProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatmulProto) ProtoMessage() {}

func (x *MatmulProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_conv_autotuning_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatmulProto.ProtoReflect.Descriptor instead.
func (*MatmulProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_conv_autotuning_proto_rawDescGZIP(), []int{1}
}

func (x *MatmulProto) GetAbDtype() DataType {
	if x != nil {
		return x.AbDtype
	}
	return DataType(0)
}

func (x *MatmulProto) GetCDtype() DataType {
	if x != nil {
		return x.CDtype
	}
	return DataType(0)
}

func (x *MatmulProto) GetTransA() bool {
	if x != nil {
		return x.TransA
	}
	return false
}

func (x *MatmulProto) GetTransB() bool {
	if x != nil {
		return x.TransB
	}
	return false
}

func (x *MatmulProto) GetM() uint64 {
	if x != nil {
		return x.M
	}
	return 0
}

func (x *MatmulProto) GetN() uint64 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *MatmulProto) GetK() uint64 {
	if x != nil {
		return x.K
	}
	return 0
}

func (x *MatmulProto) GetLda() int64 {
	if x != nil {
		return x.Lda
	}
	return 0
}

func (x *MatmulProto) GetLdb() int64 {
	if x != nil {
		return x.Ldb
	}
	return 0
}

func (x *MatmulProto) GetLdc() int64 {
	if x != nil {
		return x.Ldc
	}
	return 0
}

func (x *MatmulProto) GetActivation() ActivationMode {
	if x != nil {
		return x.Activation
	}
	return ActivationMode(0)
}

func (x *MatmulProto) GetAAddress() int64 {
	if x != nil {
		return x.AAddress
	}
	return 0
}

func (x *MatmulProto) GetBAddress() int64 {
	if x != nil {
		return x.BAddress
	}
	return 0
}

func (x *MatmulProto) GetCAddress() int64 {
	if x != nil {
		return x.CAddress
	}
	return 0
}

func (x *MatmulProto) GetBiasAddress() int64 {
	if x != nil {
		return x.BiasAddress
	}
	return 0
}

var File_tensorflow_core_protobuf_conv_autotuning_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_conv_autotuning_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x5f,
	0x61, 0x75, 0x74, 0x6f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x31, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x2f, 0x64, 0x6e, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb6, 0x05, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x38, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x24, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x6e, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x40,
	0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e,
	0x64, 0x6e, 0x6e, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x42, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x2e, 0x64, 0x6e, 0x6e, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x6e, 0x6e, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x4c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x76,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x6e,
	0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x76, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x76, 0x5f, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x76,
	0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0e, 0x73, 0x69, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12,
	0x43, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x6e, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x69, 0x61, 0x73, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62,
	0x69, 0x61, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x69,
	0x64, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x69, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xd0, 0x03, 0x0a, 0x0b, 0x4d, 0x61, 0x74,
	0x6d, 0x75, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x38, 0x0a, 0x08, 0x61, 0x62, 0x5f, 0x64,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x6e, 0x6e,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x61, 0x62, 0x44, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x5f, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x64, 0x6e, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x06, 0x63, 0x44, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x5f, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x41, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x5f, 0x62, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x42, 0x12, 0x0c, 0x0a, 0x01,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x01, 0x6d, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x01, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x6b, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x01, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x64, 0x61, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x6c, 0x64, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x64, 0x62, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6c, 0x64, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x64,
	0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6c, 0x64, 0x63, 0x12, 0x43, 0x0a, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x23, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x2e, 0x64, 0x6e, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x62, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x62, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x69, 0x61, 0x73,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x62, 0x69, 0x61, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x57, 0x5a, 0x55, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_conv_autotuning_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_conv_autotuning_proto_rawDescData = file_tensorflow_core_protobuf_conv_autotuning_proto_rawDesc
)

func file_tensorflow_core_protobuf_conv_autotuning_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_conv_autotuning_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_conv_autotuning_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_conv_autotuning_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_conv_autotuning_proto_rawDescData
}

var file_tensorflow_core_protobuf_conv_autotuning_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_core_protobuf_conv_autotuning_proto_goTypes = []interface{}{
	(*ConvolutionProto)(nil),                           // 0: tensorflow.ConvolutionProto
	(*MatmulProto)(nil),                                // 1: tensorflow.MatmulProto
	(ConvolutionKind)(0),               // 2: dnn.ConvolutionKind
	(*TensorDescriptorProto)(nil),      // 3: dnn.TensorDescriptorProto
	(*ConvolutionDescriptorProto)(nil), // 4: dnn.ConvolutionDescriptorProto
	(ActivationMode)(0),                // 5: dnn.ActivationMode
	(DataType)(0),                      // 6: dnn.DataType
}
var file_tensorflow_core_protobuf_conv_autotuning_proto_depIdxs = []int32{
	2, // 0: tensorflow.ConvolutionProto.kind:type_name -> dnn.ConvolutionKind
	3, // 1: tensorflow.ConvolutionProto.input:type_name -> dnn.TensorDescriptorProto
	3, // 2: tensorflow.ConvolutionProto.filter:type_name -> dnn.TensorDescriptorProto
	3, // 3: tensorflow.ConvolutionProto.output:type_name -> dnn.TensorDescriptorProto
	4, // 4: tensorflow.ConvolutionProto.conv_desc:type_name -> dnn.ConvolutionDescriptorProto
	5, // 5: tensorflow.ConvolutionProto.activation:type_name -> dnn.ActivationMode
	6, // 6: tensorflow.MatmulProto.ab_dtype:type_name -> dnn.DataType
	6, // 7: tensorflow.MatmulProto.c_dtype:type_name -> dnn.DataType
	5, // 8: tensorflow.MatmulProto.activation:type_name -> dnn.ActivationMode
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_conv_autotuning_proto_init() }
func file_tensorflow_core_protobuf_conv_autotuning_proto_init() {
	if File_tensorflow_core_protobuf_conv_autotuning_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_conv_autotuning_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvolutionProto); i {
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
		file_tensorflow_core_protobuf_conv_autotuning_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatmulProto); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_conv_autotuning_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_conv_autotuning_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_conv_autotuning_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_conv_autotuning_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_conv_autotuning_proto = out.File
	file_tensorflow_core_protobuf_conv_autotuning_proto_rawDesc = nil
	file_tensorflow_core_protobuf_conv_autotuning_proto_goTypes = nil
	file_tensorflow_core_protobuf_conv_autotuning_proto_depIdxs = nil
}
