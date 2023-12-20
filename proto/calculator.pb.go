// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: proto/calculator.proto

package calculator

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

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 float64 `protobuf:"fixed64,1,opt,name=Num1,proto3" json:"Num1,omitempty"`
	Num2 float64 `protobuf:"fixed64,2,opt,name=Num2,proto3" json:"Num2,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetNum1() float64 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *AddRequest) GetNum2() float64 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type SubtractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 float64 `protobuf:"fixed64,1,opt,name=Num1,proto3" json:"Num1,omitempty"`
	Num2 float64 `protobuf:"fixed64,2,opt,name=Num2,proto3" json:"Num2,omitempty"`
}

func (x *SubtractRequest) Reset() {
	*x = SubtractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubtractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubtractRequest) ProtoMessage() {}

func (x *SubtractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubtractRequest.ProtoReflect.Descriptor instead.
func (*SubtractRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *SubtractRequest) GetNum1() float64 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *SubtractRequest) GetNum2() float64 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type MultiplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 float64 `protobuf:"fixed64,1,opt,name=Num1,proto3" json:"Num1,omitempty"`
	Num2 float64 `protobuf:"fixed64,2,opt,name=Num2,proto3" json:"Num2,omitempty"`
}

func (x *MultiplyRequest) Reset() {
	*x = MultiplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplyRequest) ProtoMessage() {}

func (x *MultiplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplyRequest.ProtoReflect.Descriptor instead.
func (*MultiplyRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *MultiplyRequest) GetNum1() float64 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *MultiplyRequest) GetNum2() float64 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type DivideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 float64 `protobuf:"fixed64,1,opt,name=Num1,proto3" json:"Num1,omitempty"`
	Num2 float64 `protobuf:"fixed64,2,opt,name=Num2,proto3" json:"Num2,omitempty"`
}

func (x *DivideRequest) Reset() {
	*x = DivideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivideRequest) ProtoMessage() {}

func (x *DivideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivideRequest.ProtoReflect.Descriptor instead.
func (*DivideRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *DivideRequest) GetNum1() float64 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *DivideRequest) GetNum2() float64 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response float64 `protobuf:"fixed64,1,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetResponse() float64 {
	if x != nil {
		return x.Response
	}
	return 0
}

var File_proto_calculator_proto protoreflect.FileDescriptor

var file_proto_calculator_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x22, 0x34, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x04, 0x4e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x75, 0x6d, 0x32, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x4e, 0x75, 0x6d, 0x32, 0x22, 0x39, 0x0a, 0x0f, 0x53, 0x75,
	0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x4e, 0x75, 0x6d,
	0x31, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x04, 0x4e, 0x75, 0x6d, 0x32, 0x22, 0x39, 0x0a, 0x0f, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x75, 0x6d, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x4e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x4e, 0x75, 0x6d, 0x32,
	0x22, 0x37, 0x0a, 0x0d, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x04, 0x4e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x04, 0x4e, 0x75, 0x6d, 0x32, 0x22, 0x26, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xfa, 0x01, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x33, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x16, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53,
	0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79,
	0x12, 0x1b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x12, 0x19, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d,
	0x5a, 0x0b, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_calculator_proto_rawDescOnce sync.Once
	file_proto_calculator_proto_rawDescData = file_proto_calculator_proto_rawDesc
)

func file_proto_calculator_proto_rawDescGZIP() []byte {
	file_proto_calculator_proto_rawDescOnce.Do(func() {
		file_proto_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_calculator_proto_rawDescData)
	})
	return file_proto_calculator_proto_rawDescData
}

var file_proto_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_calculator_proto_goTypes = []interface{}{
	(*AddRequest)(nil),      // 0: calculator.AddRequest
	(*SubtractRequest)(nil), // 1: calculator.SubtractRequest
	(*MultiplyRequest)(nil), // 2: calculator.MultiplyRequest
	(*DivideRequest)(nil),   // 3: calculator.DivideRequest
	(*Response)(nil),        // 4: calculator.Response
}
var file_proto_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.Calculator.Add:input_type -> calculator.AddRequest
	1, // 1: calculator.Calculator.Subtract:input_type -> calculator.SubtractRequest
	2, // 2: calculator.Calculator.Multiply:input_type -> calculator.MultiplyRequest
	3, // 3: calculator.Calculator.Divide:input_type -> calculator.DivideRequest
	4, // 4: calculator.Calculator.Add:output_type -> calculator.Response
	4, // 5: calculator.Calculator.Subtract:output_type -> calculator.Response
	4, // 6: calculator.Calculator.Multiply:output_type -> calculator.Response
	4, // 7: calculator.Calculator.Divide:output_type -> calculator.Response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_calculator_proto_init() }
func file_proto_calculator_proto_init() {
	if File_proto_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_proto_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubtractRequest); i {
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
		file_proto_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiplyRequest); i {
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
		file_proto_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivideRequest); i {
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
		file_proto_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_calculator_proto_goTypes,
		DependencyIndexes: file_proto_calculator_proto_depIdxs,
		MessageInfos:      file_proto_calculator_proto_msgTypes,
	}.Build()
	File_proto_calculator_proto = out.File
	file_proto_calculator_proto_rawDesc = nil
	file_proto_calculator_proto_goTypes = nil
	file_proto_calculator_proto_depIdxs = nil
}
