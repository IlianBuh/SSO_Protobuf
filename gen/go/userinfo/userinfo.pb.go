// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: userinfo.proto

package userinfov1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UsersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuids         []int32                `protobuf:"varint,1,rep,packed,name=uuids,proto3" json:"uuids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UsersRequest) Reset() {
	*x = UsersRequest{}
	mi := &file_userinfo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersRequest) ProtoMessage() {}

func (x *UsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersRequest.ProtoReflect.Descriptor instead.
func (*UsersRequest) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{0}
}

func (x *UsersRequest) GetUuids() []int32 {
	if x != nil {
		return x.Uuids
	}
	return nil
}

type UsersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*User                `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UsersResponse) Reset() {
	*x = UsersResponse{}
	mi := &file_userinfo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersResponse) ProtoMessage() {}

func (x *UsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersResponse.ProtoReflect.Descriptor instead.
func (*UsersResponse) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{1}
}

func (x *UsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          int32                  `protobuf:"varint,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_userinfo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{2}
}

func (x *UserRequest) GetUuid() int32 {
	if x != nil {
		return x.Uuid
	}
	return 0
}

type UserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	mi := &file_userinfo_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{3}
}

func (x *UserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          int32                  `protobuf:"varint,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Login         string                 `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_userinfo_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetUuid() int32 {
	if x != nil {
		return x.Uuid
	}
	return 0
}

func (x *User) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UsersExistRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          []int32                `protobuf:"varint,1,rep,packed,name=uuid,proto3" json:"uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UsersExistRequest) Reset() {
	*x = UsersExistRequest{}
	mi := &file_userinfo_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersExistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersExistRequest) ProtoMessage() {}

func (x *UsersExistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersExistRequest.ProtoReflect.Descriptor instead.
func (*UsersExistRequest) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{5}
}

func (x *UsersExistRequest) GetUuid() []int32 {
	if x != nil {
		return x.Uuid
	}
	return nil
}

type UsersExistResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Exist         bool                   `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UsersExistResponse) Reset() {
	*x = UsersExistResponse{}
	mi := &file_userinfo_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersExistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersExistResponse) ProtoMessage() {}

func (x *UsersExistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersExistResponse.ProtoReflect.Descriptor instead.
func (*UsersExistResponse) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{6}
}

func (x *UsersExistResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

var File_userinfo_proto protoreflect.FileDescriptor

const file_userinfo_proto_rawDesc = "" +
	"\n" +
	"\x0euserinfo.proto\x12\buserinfo\"$\n" +
	"\fUsersRequest\x12\x14\n" +
	"\x05uuids\x18\x01 \x03(\x05R\x05uuids\"5\n" +
	"\rUsersResponse\x12$\n" +
	"\x05users\x18\x01 \x03(\v2\x0e.userinfo.UserR\x05users\"!\n" +
	"\vUserRequest\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\x05R\x04uuid\"2\n" +
	"\fUserResponse\x12\"\n" +
	"\x04user\x18\x01 \x01(\v2\x0e.userinfo.UserR\x04user\"F\n" +
	"\x04User\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\x05R\x04uuid\x12\x14\n" +
	"\x05login\x18\x02 \x01(\tR\x05login\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\"'\n" +
	"\x11UsersExistRequest\x12\x12\n" +
	"\x04uuid\x18\x01 \x03(\x05R\x04uuid\"*\n" +
	"\x12UsersExistResponse\x12\x14\n" +
	"\x05exist\x18\x01 \x01(\bR\x05exist2\xc4\x01\n" +
	"\bUserInfo\x128\n" +
	"\x05Users\x12\x16.userinfo.UsersRequest\x1a\x17.userinfo.UsersResponse\x125\n" +
	"\x04User\x12\x15.userinfo.UserRequest\x1a\x16.userinfo.UserResponse\x12G\n" +
	"\n" +
	"UsersExist\x12\x1b.userinfo.UsersExistRequest\x1a\x1c.userinfo.UsersExistResponseB!Z\x1fIlianBuh.userinfo.v1;userinfov1b\x06proto3"

var (
	file_userinfo_proto_rawDescOnce sync.Once
	file_userinfo_proto_rawDescData []byte
)

func file_userinfo_proto_rawDescGZIP() []byte {
	file_userinfo_proto_rawDescOnce.Do(func() {
		file_userinfo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_userinfo_proto_rawDesc), len(file_userinfo_proto_rawDesc)))
	})
	return file_userinfo_proto_rawDescData
}

var file_userinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_userinfo_proto_goTypes = []any{
	(*UsersRequest)(nil),       // 0: userinfo.UsersRequest
	(*UsersResponse)(nil),      // 1: userinfo.UsersResponse
	(*UserRequest)(nil),        // 2: userinfo.UserRequest
	(*UserResponse)(nil),       // 3: userinfo.UserResponse
	(*User)(nil),               // 4: userinfo.User
	(*UsersExistRequest)(nil),  // 5: userinfo.UsersExistRequest
	(*UsersExistResponse)(nil), // 6: userinfo.UsersExistResponse
}
var file_userinfo_proto_depIdxs = []int32{
	4, // 0: userinfo.UsersResponse.users:type_name -> userinfo.User
	4, // 1: userinfo.UserResponse.user:type_name -> userinfo.User
	0, // 2: userinfo.UserInfo.Users:input_type -> userinfo.UsersRequest
	2, // 3: userinfo.UserInfo.User:input_type -> userinfo.UserRequest
	5, // 4: userinfo.UserInfo.UsersExist:input_type -> userinfo.UsersExistRequest
	1, // 5: userinfo.UserInfo.Users:output_type -> userinfo.UsersResponse
	3, // 6: userinfo.UserInfo.User:output_type -> userinfo.UserResponse
	6, // 7: userinfo.UserInfo.UsersExist:output_type -> userinfo.UsersExistResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_userinfo_proto_init() }
func file_userinfo_proto_init() {
	if File_userinfo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_userinfo_proto_rawDesc), len(file_userinfo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userinfo_proto_goTypes,
		DependencyIndexes: file_userinfo_proto_depIdxs,
		MessageInfos:      file_userinfo_proto_msgTypes,
	}.Build()
	File_userinfo_proto = out.File
	file_userinfo_proto_goTypes = nil
	file_userinfo_proto_depIdxs = nil
}
