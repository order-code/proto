// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.15.8
// source: authenticator.proto

package authenticator

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

type TokenVerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtToken string `protobuf:"bytes,1,opt,name=jwt_token,json=jwtToken,proto3" json:"jwt_token,omitempty"`
}

func (x *TokenVerifyRequest) Reset() {
	*x = TokenVerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenVerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenVerifyRequest) ProtoMessage() {}

func (x *TokenVerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenVerifyRequest.ProtoReflect.Descriptor instead.
func (*TokenVerifyRequest) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{0}
}

func (x *TokenVerifyRequest) GetJwtToken() string {
	if x != nil {
		return x.JwtToken
	}
	return ""
}

type TokenVerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	AccountId   string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Status      int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	StaffId     string `protobuf:"bytes,4,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
	WorkspaceId string `protobuf:"bytes,5,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
}

func (x *TokenVerifyResponse) Reset() {
	*x = TokenVerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authenticator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenVerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenVerifyResponse) ProtoMessage() {}

func (x *TokenVerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authenticator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenVerifyResponse.ProtoReflect.Descriptor instead.
func (*TokenVerifyResponse) Descriptor() ([]byte, []int) {
	return file_authenticator_proto_rawDescGZIP(), []int{1}
}

func (x *TokenVerifyResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TokenVerifyResponse) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *TokenVerifyResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TokenVerifyResponse) GetStaffId() string {
	if x != nil {
		return x.StaffId
	}
	return ""
}

func (x *TokenVerifyResponse) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

var File_authenticator_proto protoreflect.FileDescriptor

var file_authenticator_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6d, 0x65, 0x78, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x31, 0x0a, 0x12, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x77, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa6,
	0x01, 0x0a, 0x13, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x66, 0x66, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x32, 0x7a, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x62, 0x0a, 0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x28,
	0x2e, 0x6d, 0x65, 0x78, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d, 0x65, 0x78, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authenticator_proto_rawDescOnce sync.Once
	file_authenticator_proto_rawDescData = file_authenticator_proto_rawDesc
)

func file_authenticator_proto_rawDescGZIP() []byte {
	file_authenticator_proto_rawDescOnce.Do(func() {
		file_authenticator_proto_rawDescData = protoimpl.X.CompressGZIP(file_authenticator_proto_rawDescData)
	})
	return file_authenticator_proto_rawDescData
}

var file_authenticator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_authenticator_proto_goTypes = []interface{}{
	(*TokenVerifyRequest)(nil),  // 0: mex.authenticator.v1.TokenVerifyRequest
	(*TokenVerifyResponse)(nil), // 1: mex.authenticator.v1.TokenVerifyResponse
}
var file_authenticator_proto_depIdxs = []int32{
	0, // 0: mex.authenticator.v1.AuthenticatorService.TokenVerify:input_type -> mex.authenticator.v1.TokenVerifyRequest
	1, // 1: mex.authenticator.v1.AuthenticatorService.TokenVerify:output_type -> mex.authenticator.v1.TokenVerifyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authenticator_proto_init() }
func file_authenticator_proto_init() {
	if File_authenticator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authenticator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenVerifyRequest); i {
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
		file_authenticator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenVerifyResponse); i {
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
			RawDescriptor: file_authenticator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authenticator_proto_goTypes,
		DependencyIndexes: file_authenticator_proto_depIdxs,
		MessageInfos:      file_authenticator_proto_msgTypes,
	}.Build()
	File_authenticator_proto = out.File
	file_authenticator_proto_rawDesc = nil
	file_authenticator_proto_goTypes = nil
	file_authenticator_proto_depIdxs = nil
}
