// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.2
// source: emoji.proto

package emoji

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

type EmojizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *EmojizeRequest) Reset() {
	*x = EmojizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoji_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmojizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmojizeRequest) ProtoMessage() {}

func (x *EmojizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_emoji_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmojizeRequest.ProtoReflect.Descriptor instead.
func (*EmojizeRequest) Descriptor() ([]byte, []int) {
	return file_emoji_proto_rawDescGZIP(), []int{0}
}

func (x *EmojizeRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type EmojizeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmojizedText string `protobuf:"bytes,1,opt,name=emojized_text,json=emojizedText,proto3" json:"emojized_text,omitempty"`
}

func (x *EmojizeReply) Reset() {
	*x = EmojizeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoji_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmojizeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmojizeReply) ProtoMessage() {}

func (x *EmojizeReply) ProtoReflect() protoreflect.Message {
	mi := &file_emoji_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmojizeReply.ProtoReflect.Descriptor instead.
func (*EmojizeReply) Descriptor() ([]byte, []int) {
	return file_emoji_proto_rawDescGZIP(), []int{1}
}

func (x *EmojizeReply) GetEmojizedText() string {
	if x != nil {
		return x.EmojizedText
	}
	return ""
}

var File_emoji_proto protoreflect.FileDescriptor

var file_emoji_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x6d, 0x6f, 0x6a, 0x69, 0x22, 0x24, 0x0a, 0x0e, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x7a, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x33, 0x0a, 0x0c, 0x45, 0x6d,
	0x6f, 0x6a, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d,
	0x6f, 0x6a, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x7a, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x32,
	0x45, 0x0a, 0x0c, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x35, 0x0a, 0x07, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x7a, 0x65, 0x12, 0x15, 0x2e, 0x65, 0x6d, 0x6f,
	0x6a, 0x69, 0x2e, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x75, 0x68, 0x75, 0x69, 0x7a, 0x75, 0x6f, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x77, 0x65, 0x62, 0x2d, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2f, 0x65, 0x6d, 0x6f,
	0x6a, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_emoji_proto_rawDescOnce sync.Once
	file_emoji_proto_rawDescData = file_emoji_proto_rawDesc
)

func file_emoji_proto_rawDescGZIP() []byte {
	file_emoji_proto_rawDescOnce.Do(func() {
		file_emoji_proto_rawDescData = protoimpl.X.CompressGZIP(file_emoji_proto_rawDescData)
	})
	return file_emoji_proto_rawDescData
}

var file_emoji_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_emoji_proto_goTypes = []interface{}{
	(*EmojizeRequest)(nil), // 0: emoji.EmojizeRequest
	(*EmojizeReply)(nil),   // 1: emoji.EmojizeReply
}
var file_emoji_proto_depIdxs = []int32{
	0, // 0: emoji.EmojiService.Emojize:input_type -> emoji.EmojizeRequest
	1, // 1: emoji.EmojiService.Emojize:output_type -> emoji.EmojizeReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_emoji_proto_init() }
func file_emoji_proto_init() {
	if File_emoji_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_emoji_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmojizeRequest); i {
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
		file_emoji_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmojizeReply); i {
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
			RawDescriptor: file_emoji_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_emoji_proto_goTypes,
		DependencyIndexes: file_emoji_proto_depIdxs,
		MessageInfos:      file_emoji_proto_msgTypes,
	}.Build()
	File_emoji_proto = out.File
	file_emoji_proto_rawDesc = nil
	file_emoji_proto_goTypes = nil
	file_emoji_proto_depIdxs = nil
}