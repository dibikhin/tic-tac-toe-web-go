// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: tttwebgrpc/tictactoe.proto

package tttwebgrpc

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tttwebgrpc_tictactoe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_tttwebgrpc_tictactoe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_tttwebgrpc_tictactoe_proto_rawDescGZIP(), []int{0}
}

type StatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Action string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *StatusReply) Reset() {
	*x = StatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tttwebgrpc_tictactoe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusReply) ProtoMessage() {}

func (x *StatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_tttwebgrpc_tictactoe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusReply.ProtoReflect.Descriptor instead.
func (*StatusReply) Descriptor() ([]byte, []int) {
	return file_tttwebgrpc_tictactoe_proto_rawDescGZIP(), []int{1}
}

func (x *StatusReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StatusReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

var File_tttwebgrpc_tictactoe_proto protoreflect.FileDescriptor

var file_tttwebgrpc_tictactoe_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x74, 0x74, 0x77, 0x65, 0x62, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x69, 0x63,
	0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x69,
	0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x3d, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0x3f, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74,
	0x6f, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x69, 0x62, 0x69, 0x6b, 0x68, 0x69, 0x6e, 0x2f, 0x74, 0x69, 0x63, 0x2d, 0x74, 0x61, 0x63, 0x2d,
	0x74, 0x6f, 0x65, 0x2d, 0x77, 0x65, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x74, 0x74, 0x77, 0x65,
	0x62, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tttwebgrpc_tictactoe_proto_rawDescOnce sync.Once
	file_tttwebgrpc_tictactoe_proto_rawDescData = file_tttwebgrpc_tictactoe_proto_rawDesc
)

func file_tttwebgrpc_tictactoe_proto_rawDescGZIP() []byte {
	file_tttwebgrpc_tictactoe_proto_rawDescOnce.Do(func() {
		file_tttwebgrpc_tictactoe_proto_rawDescData = protoimpl.X.CompressGZIP(file_tttwebgrpc_tictactoe_proto_rawDescData)
	})
	return file_tttwebgrpc_tictactoe_proto_rawDescData
}

var file_tttwebgrpc_tictactoe_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tttwebgrpc_tictactoe_proto_goTypes = []interface{}{
	(*Empty)(nil),       // 0: tictactoe.Empty
	(*StatusReply)(nil), // 1: tictactoe.StatusReply
}
var file_tttwebgrpc_tictactoe_proto_depIdxs = []int32{
	0, // 0: tictactoe.Game.GetStatus:input_type -> tictactoe.Empty
	1, // 1: tictactoe.Game.GetStatus:output_type -> tictactoe.StatusReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tttwebgrpc_tictactoe_proto_init() }
func file_tttwebgrpc_tictactoe_proto_init() {
	if File_tttwebgrpc_tictactoe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tttwebgrpc_tictactoe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_tttwebgrpc_tictactoe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusReply); i {
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
			RawDescriptor: file_tttwebgrpc_tictactoe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tttwebgrpc_tictactoe_proto_goTypes,
		DependencyIndexes: file_tttwebgrpc_tictactoe_proto_depIdxs,
		MessageInfos:      file_tttwebgrpc_tictactoe_proto_msgTypes,
	}.Build()
	File_tttwebgrpc_tictactoe_proto = out.File
	file_tttwebgrpc_tictactoe_proto_rawDesc = nil
	file_tttwebgrpc_tictactoe_proto_goTypes = nil
	file_tttwebgrpc_tictactoe_proto_depIdxs = nil
}
