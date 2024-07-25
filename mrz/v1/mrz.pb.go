// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: mrz/v1/mrz.proto

package mrz_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Args *anypb.Any `protobuf:"bytes,2,opt,name=Args,proto3" json:"Args,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mrz_v1_mrz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_mrz_v1_mrz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_mrz_v1_mrz_proto_rawDescGZIP(), []int{0}
}

func (x *Model) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Model) GetArgs() *anypb.Any {
	if x != nil {
		return x.Args
	}
	return nil
}

type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Resp *anypb.Any `protobuf:"bytes,2,opt,name=resp,proto3" json:"resp,omitempty"`
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mrz_v1_mrz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_mrz_v1_mrz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_mrz_v1_mrz_proto_rawDescGZIP(), []int{1}
}

func (x *Res) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Res) GetResp() *anypb.Any {
	if x != nil {
		return x.Resp
	}
	return nil
}

var File_mrz_v1_mrz_proto protoreflect.FileDescriptor

var file_mrz_v1_mrz_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x72, 0x7a, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x72, 0x7a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x6d, 0x72, 0x7a, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x28,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x04, 0x41, 0x72, 0x67, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x41, 0x72,
	0x67, 0x73, 0x22, 0x59, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x72, 0x65, 0x73, 0x70, 0x42, 0x5a, 0x0a,
	0x20, 0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x66, 0x73, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x6c, 0x65, 0x72,
	0x2e, 0x6d, 0x71, 0x74, 0x74, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x72, 0x7a, 0x2e, 0x76,
	0x31, 0x42, 0x08, 0x4d, 0x72, 0x7a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x6c, 0x65,
	0x72, 0x2f, 0x6d, 0x71, 0x74, 0x74, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x72, 0x7a, 0x2f,
	0x76, 0x31, 0x3b, 0x6d, 0x72, 0x7a, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mrz_v1_mrz_proto_rawDescOnce sync.Once
	file_mrz_v1_mrz_proto_rawDescData = file_mrz_v1_mrz_proto_rawDesc
)

func file_mrz_v1_mrz_proto_rawDescGZIP() []byte {
	file_mrz_v1_mrz_proto_rawDescOnce.Do(func() {
		file_mrz_v1_mrz_proto_rawDescData = protoimpl.X.CompressGZIP(file_mrz_v1_mrz_proto_rawDescData)
	})
	return file_mrz_v1_mrz_proto_rawDescData
}

var file_mrz_v1_mrz_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mrz_v1_mrz_proto_goTypes = []interface{}{
	(*Model)(nil),     // 0: mrz.v1.Model
	(*Res)(nil),       // 1: mrz.v1.Res
	(*anypb.Any)(nil), // 2: google.protobuf.Any
}
var file_mrz_v1_mrz_proto_depIdxs = []int32{
	2, // 0: mrz.v1.Model.data:type_name -> google.protobuf.Any
	2, // 1: mrz.v1.Model.Args:type_name -> google.protobuf.Any
	2, // 2: mrz.v1.Res.data:type_name -> google.protobuf.Any
	2, // 3: mrz.v1.Res.resp:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mrz_v1_mrz_proto_init() }
func file_mrz_v1_mrz_proto_init() {
	if File_mrz_v1_mrz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mrz_v1_mrz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
		file_mrz_v1_mrz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
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
			RawDescriptor: file_mrz_v1_mrz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mrz_v1_mrz_proto_goTypes,
		DependencyIndexes: file_mrz_v1_mrz_proto_depIdxs,
		MessageInfos:      file_mrz_v1_mrz_proto_msgTypes,
	}.Build()
	File_mrz_v1_mrz_proto = out.File
	file_mrz_v1_mrz_proto_rawDesc = nil
	file_mrz_v1_mrz_proto_goTypes = nil
	file_mrz_v1_mrz_proto_depIdxs = nil
}