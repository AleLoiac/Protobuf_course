// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0--rc3
// source: complex/complex.proto

package complex

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

// nested messages
type ComplexMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	One_Message     *BasicMessage   `protobuf:"bytes,2,opt,name=one_Message,json=oneMessage,proto3" json:"one_Message,omitempty"`
	MultipleMessage []*BasicMessage `protobuf:"bytes,3,rep,name=multiple_message,json=multipleMessage,proto3" json:"multiple_message,omitempty"`
}

func (x *ComplexMessage) Reset() {
	*x = ComplexMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_complex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexMessage) ProtoMessage() {}

func (x *ComplexMessage) ProtoReflect() protoreflect.Message {
	mi := &file_complex_complex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexMessage.ProtoReflect.Descriptor instead.
func (*ComplexMessage) Descriptor() ([]byte, []int) {
	return file_complex_complex_proto_rawDescGZIP(), []int{0}
}

func (x *ComplexMessage) GetOne_Message() *BasicMessage {
	if x != nil {
		return x.One_Message
	}
	return nil
}

func (x *ComplexMessage) GetMultipleMessage() []*BasicMessage {
	if x != nil {
		return x.MultipleMessage
	}
	return nil
}

type BasicMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BasicMessage) Reset() {
	*x = BasicMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_complex_complex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicMessage) ProtoMessage() {}

func (x *BasicMessage) ProtoReflect() protoreflect.Message {
	mi := &file_complex_complex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicMessage.ProtoReflect.Descriptor instead.
func (*BasicMessage) Descriptor() ([]byte, []int) {
	return file_complex_complex_proto_rawDescGZIP(), []int{1}
}

func (x *BasicMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BasicMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_complex_complex_proto protoreflect.FileDescriptor

var file_complex_complex_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x22, 0x8a, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x0a, 0x6f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x10, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x2e,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x32, 0x0a,
	0x0c, 0x42, 0x61, 0x73, 0x69, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_complex_complex_proto_rawDescOnce sync.Once
	file_complex_complex_proto_rawDescData = file_complex_complex_proto_rawDesc
)

func file_complex_complex_proto_rawDescGZIP() []byte {
	file_complex_complex_proto_rawDescOnce.Do(func() {
		file_complex_complex_proto_rawDescData = protoimpl.X.CompressGZIP(file_complex_complex_proto_rawDescData)
	})
	return file_complex_complex_proto_rawDescData
}

var file_complex_complex_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_complex_complex_proto_goTypes = []interface{}{
	(*ComplexMessage)(nil), // 0: complex.ComplexMessage
	(*BasicMessage)(nil),   // 1: complex.BasicMessage
}
var file_complex_complex_proto_depIdxs = []int32{
	1, // 0: complex.ComplexMessage.one_Message:type_name -> complex.BasicMessage
	1, // 1: complex.ComplexMessage.multiple_message:type_name -> complex.BasicMessage
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_complex_complex_proto_init() }
func file_complex_complex_proto_init() {
	if File_complex_complex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_complex_complex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexMessage); i {
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
		file_complex_complex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicMessage); i {
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
			RawDescriptor: file_complex_complex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_complex_complex_proto_goTypes,
		DependencyIndexes: file_complex_complex_proto_depIdxs,
		MessageInfos:      file_complex_complex_proto_msgTypes,
	}.Build()
	File_complex_complex_proto = out.File
	file_complex_complex_proto_rawDesc = nil
	file_complex_complex_proto_goTypes = nil
	file_complex_complex_proto_depIdxs = nil
}
