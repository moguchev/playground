// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.5
// source: bar.proto

package bar

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

type GetFooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFooRequest) Reset() {
	*x = GetFooRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFooRequest) ProtoMessage() {}

func (x *GetFooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFooRequest.ProtoReflect.Descriptor instead.
func (*GetFooRequest) Descriptor() ([]byte, []int) {
	return file_bar_proto_rawDescGZIP(), []int{0}
}

func (x *GetFooRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetFooResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foos []*Object `protobuf:"bytes,1,rep,name=foos,proto3" json:"foos,omitempty"`
}

func (x *GetFooResponse) Reset() {
	*x = GetFooResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFooResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFooResponse) ProtoMessage() {}

func (x *GetFooResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFooResponse.ProtoReflect.Descriptor instead.
func (*GetFooResponse) Descriptor() ([]byte, []int) {
	return file_bar_proto_rawDescGZIP(), []int{1}
}

func (x *GetFooResponse) GetFoos() []*Object {
	if x != nil {
		return x.Foos
	}
	return nil
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_bar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_bar_proto_rawDescGZIP(), []int{2}
}

func (x *Object) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_bar_proto protoreflect.FileDescriptor

var file_bar_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x62, 0x61, 0x72,
	0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x31, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x61, 0x72, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04,
	0x66, 0x6f, 0x6f, 0x73, 0x22, 0x18, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x3d,
	0x0a, 0x03, 0x42, 0x61, 0x72, 0x12, 0x36, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x42,
	0x61, 0x72, 0x12, 0x12, 0x2e, 0x62, 0x61, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x62, 0x61, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x67, 0x75,
	0x63, 0x68, 0x65, 0x76, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bar_proto_rawDescOnce sync.Once
	file_bar_proto_rawDescData = file_bar_proto_rawDesc
)

func file_bar_proto_rawDescGZIP() []byte {
	file_bar_proto_rawDescOnce.Do(func() {
		file_bar_proto_rawDescData = protoimpl.X.CompressGZIP(file_bar_proto_rawDescData)
	})
	return file_bar_proto_rawDescData
}

var file_bar_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bar_proto_goTypes = []interface{}{
	(*GetFooRequest)(nil),  // 0: bar.GetFooRequest
	(*GetFooResponse)(nil), // 1: bar.GetFooResponse
	(*Object)(nil),         // 2: bar.Object
}
var file_bar_proto_depIdxs = []int32{
	2, // 0: bar.GetFooResponse.foos:type_name -> bar.Object
	0, // 1: bar.Bar.GetFooBar:input_type -> bar.GetFooRequest
	1, // 2: bar.Bar.GetFooBar:output_type -> bar.GetFooResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bar_proto_init() }
func file_bar_proto_init() {
	if File_bar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFooRequest); i {
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
		file_bar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFooResponse); i {
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
		file_bar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
			RawDescriptor: file_bar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bar_proto_goTypes,
		DependencyIndexes: file_bar_proto_depIdxs,
		MessageInfos:      file_bar_proto_msgTypes,
	}.Build()
	File_bar_proto = out.File
	file_bar_proto_rawDesc = nil
	file_bar_proto_goTypes = nil
	file_bar_proto_depIdxs = nil
}
