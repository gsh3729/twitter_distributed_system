// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: connect.proto

package connect

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

type FollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User1 string `protobuf:"bytes,1,opt,name=user1,proto3" json:"user1,omitempty"`
	User2 string `protobuf:"bytes,2,opt,name=user2,proto3" json:"user2,omitempty"`
}

func (x *FollowRequest) Reset() {
	*x = FollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequest) ProtoMessage() {}

func (x *FollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequest.ProtoReflect.Descriptor instead.
func (*FollowRequest) Descriptor() ([]byte, []int) {
	return file_connect_proto_rawDescGZIP(), []int{0}
}

func (x *FollowRequest) GetUser1() string {
	if x != nil {
		return x.User1
	}
	return ""
}

func (x *FollowRequest) GetUser2() string {
	if x != nil {
		return x.User2
	}
	return ""
}

type FollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *FollowResponse) Reset() {
	*x = FollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowResponse) ProtoMessage() {}

func (x *FollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowResponse.ProtoReflect.Descriptor instead.
func (*FollowResponse) Descriptor() ([]byte, []int) {
	return file_connect_proto_rawDescGZIP(), []int{1}
}

func (x *FollowResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UnfollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User1 string `protobuf:"bytes,1,opt,name=user1,proto3" json:"user1,omitempty"`
	User2 string `protobuf:"bytes,2,opt,name=user2,proto3" json:"user2,omitempty"`
}

func (x *UnfollowRequest) Reset() {
	*x = UnfollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnfollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfollowRequest) ProtoMessage() {}

func (x *UnfollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connect_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfollowRequest.ProtoReflect.Descriptor instead.
func (*UnfollowRequest) Descriptor() ([]byte, []int) {
	return file_connect_proto_rawDescGZIP(), []int{2}
}

func (x *UnfollowRequest) GetUser1() string {
	if x != nil {
		return x.User1
	}
	return ""
}

func (x *UnfollowRequest) GetUser2() string {
	if x != nil {
		return x.User2
	}
	return ""
}

type UnfollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *UnfollowResponse) Reset() {
	*x = UnfollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnfollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfollowResponse) ProtoMessage() {}

func (x *UnfollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connect_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfollowResponse.ProtoReflect.Descriptor instead.
func (*UnfollowResponse) Descriptor() ([]byte, []int) {
	return file_connect_proto_rawDescGZIP(), []int{3}
}

func (x *UnfollowResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_connect_proto protoreflect.FileDescriptor

var file_connect_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x22, 0x3b, 0x0a, 0x0d, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x31, 0x12,
	0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x32, 0x22, 0x2a, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x3d, 0x0a, 0x0f, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x32,
	0x22, 0x2c, 0x0a, 0x10, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x90,
	0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3b, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x08, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x55,
	0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connect_proto_rawDescOnce sync.Once
	file_connect_proto_rawDescData = file_connect_proto_rawDesc
)

func file_connect_proto_rawDescGZIP() []byte {
	file_connect_proto_rawDescOnce.Do(func() {
		file_connect_proto_rawDescData = protoimpl.X.CompressGZIP(file_connect_proto_rawDescData)
	})
	return file_connect_proto_rawDescData
}

var file_connect_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_connect_proto_goTypes = []interface{}{
	(*FollowRequest)(nil),    // 0: connect.FollowRequest
	(*FollowResponse)(nil),   // 1: connect.FollowResponse
	(*UnfollowRequest)(nil),  // 2: connect.UnfollowRequest
	(*UnfollowResponse)(nil), // 3: connect.UnfollowResponse
}
var file_connect_proto_depIdxs = []int32{
	0, // 0: connect.ConnectService.Follow:input_type -> connect.FollowRequest
	2, // 1: connect.ConnectService.Unfollow:input_type -> connect.UnfollowRequest
	1, // 2: connect.ConnectService.Follow:output_type -> connect.FollowResponse
	3, // 3: connect.ConnectService.Unfollow:output_type -> connect.UnfollowResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_connect_proto_init() }
func file_connect_proto_init() {
	if File_connect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRequest); i {
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
		file_connect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowResponse); i {
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
		file_connect_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnfollowRequest); i {
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
		file_connect_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnfollowResponse); i {
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
			RawDescriptor: file_connect_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_connect_proto_goTypes,
		DependencyIndexes: file_connect_proto_depIdxs,
		MessageInfos:      file_connect_proto_msgTypes,
	}.Build()
	File_connect_proto = out.File
	file_connect_proto_rawDesc = nil
	file_connect_proto_goTypes = nil
	file_connect_proto_depIdxs = nil
}
