// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: user.proto

package dto

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
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

type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SummaryLength  int32  `protobuf:"varint,1,opt,name=summary_length,json=summaryLength,proto3" json:"summary_length,omitempty"`
	SummaryPrompt  string `protobuf:"bytes,2,opt,name=summary_prompt,json=summaryPrompt,proto3" json:"summary_prompt,omitempty"`
	EmphasisPrompt string `protobuf:"bytes,3,opt,name=emphasis_prompt,json=emphasisPrompt,proto3" json:"emphasis_prompt,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *Settings) GetSummaryLength() int32 {
	if x != nil {
		return x.SummaryLength
	}
	return 0
}

func (x *Settings) GetSummaryPrompt() string {
	if x != nil {
		return x.SummaryPrompt
	}
	return ""
}

func (x *Settings) GetEmphasisPrompt() string {
	if x != nil {
		return x.EmphasisPrompt
	}
	return ""
}

type UsersCreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UsersCreateUserRequest) Reset() {
	*x = UsersCreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersCreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersCreateUserRequest) ProtoMessage() {}

func (x *UsersCreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersCreateUserRequest.ProtoReflect.Descriptor instead.
func (*UsersCreateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UsersCreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UsersCreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UsersCreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pk      int32  `protobuf:"varint,1,opt,name=pk,proto3" json:"pk,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UsersCreateUserResponse) Reset() {
	*x = UsersCreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersCreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersCreateUserResponse) ProtoMessage() {}

func (x *UsersCreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersCreateUserResponse.ProtoReflect.Descriptor instead.
func (*UsersCreateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UsersCreateUserResponse) GetPk() int32 {
	if x != nil {
		return x.Pk
	}
	return 0
}

func (x *UsersCreateUserResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UsersCreateUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UsersLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UsersLoginRequest) Reset() {
	*x = UsersLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersLoginRequest) ProtoMessage() {}

func (x *UsersLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersLoginRequest.ProtoReflect.Descriptor instead.
func (*UsersLoginRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UsersLoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UsersLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UsersLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UsersLoginResponse) Reset() {
	*x = UsersLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersLoginResponse) ProtoMessage() {}

func (x *UsersLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersLoginResponse.ProtoReflect.Descriptor instead.
func (*UsersLoginResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *UsersLoginResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UsersLogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UsersLogoutRequest) Reset() {
	*x = UsersLogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersLogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersLogoutRequest) ProtoMessage() {}

func (x *UsersLogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersLogoutRequest.ProtoReflect.Descriptor instead.
func (*UsersLogoutRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

type UsersLogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UsersLogoutResponse) Reset() {
	*x = UsersLogoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersLogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersLogoutResponse) ProtoMessage() {}

func (x *UsersLogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersLogoutResponse.ProtoReflect.Descriptor instead.
func (*UsersLogoutResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *UsersLogoutResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UsersChangePwdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldPassword string `protobuf:"bytes,1,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	Id          string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UsersChangePwdRequest) Reset() {
	*x = UsersChangePwdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersChangePwdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersChangePwdRequest) ProtoMessage() {}

func (x *UsersChangePwdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersChangePwdRequest.ProtoReflect.Descriptor instead.
func (*UsersChangePwdRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *UsersChangePwdRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *UsersChangePwdRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *UsersChangePwdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UsersChangePwdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UsersChangePwdResponse) Reset() {
	*x = UsersChangePwdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersChangePwdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersChangePwdResponse) ProtoMessage() {}

func (x *UsersChangePwdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersChangePwdResponse.ProtoReflect.Descriptor instead.
func (*UsersChangePwdResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *UsersChangePwdResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UsersGetSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UsersGetSettingsRequest) Reset() {
	*x = UsersGetSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersGetSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersGetSettingsRequest) ProtoMessage() {}

func (x *UsersGetSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersGetSettingsRequest.ProtoReflect.Descriptor instead.
func (*UsersGetSettingsRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

type UsersGetSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings *Settings `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	Message  string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UsersGetSettingsResponse) Reset() {
	*x = UsersGetSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersGetSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersGetSettingsResponse) ProtoMessage() {}

func (x *UsersGetSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersGetSettingsResponse.ProtoReflect.Descriptor instead.
func (*UsersGetSettingsResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *UsersGetSettingsResponse) GetSettings() *Settings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *UsersGetSettingsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UsersUpdateSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings *Settings `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *UsersUpdateSettingsRequest) Reset() {
	*x = UsersUpdateSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersUpdateSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersUpdateSettingsRequest) ProtoMessage() {}

func (x *UsersUpdateSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersUpdateSettingsRequest.ProtoReflect.Descriptor instead.
func (*UsersUpdateSettingsRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{11}
}

func (x *UsersUpdateSettingsRequest) GetSettings() *Settings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type UsersUpdateSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UsersUpdateSettingsResponse) Reset() {
	*x = UsersUpdateSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersUpdateSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersUpdateSettingsResponse) ProtoMessage() {}

func (x *UsersUpdateSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersUpdateSettingsResponse.ProtoReflect.Descriptor instead.
func (*UsersUpdateSettingsResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{12}
}

func (x *UsersUpdateSettingsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x64, 0x74,
	0x6f, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6d, 0x70, 0x68, 0x61, 0x73, 0x69, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6d, 0x70, 0x68,
	0x61, 0x73, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x22, 0x50, 0x0a, 0x16, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x53, 0x0a, 0x17,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x70, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x4b, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2e,
	0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x14,
	0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6d, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x73, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x77, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x73, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x50, 0x77, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x5f, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x73, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a, 0x1a, 0x55, 0x73, 0x65, 0x72, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x37, 0x0a,
	0x1b, 0x55, 0x73, 0x65, 0x72, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd7, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x4e, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1b, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e,
	0x64, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12,
	0x17, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x50, 0x77, 0x64, 0x12, 0x1a, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x77, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x77, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x51, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x73, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1f, 0x2e, 0x64, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x64, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_user_proto_goTypes = []interface{}{
	(*Settings)(nil),                    // 0: dto.Settings
	(*UsersCreateUserRequest)(nil),      // 1: dto.UsersCreateUserRequest
	(*UsersCreateUserResponse)(nil),     // 2: dto.UsersCreateUserResponse
	(*UsersLoginRequest)(nil),           // 3: dto.UsersLoginRequest
	(*UsersLoginResponse)(nil),          // 4: dto.UsersLoginResponse
	(*UsersLogoutRequest)(nil),          // 5: dto.UsersLogoutRequest
	(*UsersLogoutResponse)(nil),         // 6: dto.UsersLogoutResponse
	(*UsersChangePwdRequest)(nil),       // 7: dto.UsersChangePwdRequest
	(*UsersChangePwdResponse)(nil),      // 8: dto.UsersChangePwdResponse
	(*UsersGetSettingsRequest)(nil),     // 9: dto.UsersGetSettingsRequest
	(*UsersGetSettingsResponse)(nil),    // 10: dto.UsersGetSettingsResponse
	(*UsersUpdateSettingsRequest)(nil),  // 11: dto.UsersUpdateSettingsRequest
	(*UsersUpdateSettingsResponse)(nil), // 12: dto.UsersUpdateSettingsResponse
}
var file_user_proto_depIdxs = []int32{
	0,  // 0: dto.UsersGetSettingsResponse.settings:type_name -> dto.Settings
	0,  // 1: dto.UsersUpdateSettingsRequest.settings:type_name -> dto.Settings
	1,  // 2: dto.User.UsersCreateUser:input_type -> dto.UsersCreateUserRequest
	3,  // 3: dto.User.UsersLogin:input_type -> dto.UsersLoginRequest
	5,  // 4: dto.User.UsersLogout:input_type -> dto.UsersLogoutRequest
	7,  // 5: dto.User.UsersChangePwd:input_type -> dto.UsersChangePwdRequest
	9,  // 6: dto.User.UsersGetSettings:input_type -> dto.UsersGetSettingsRequest
	11, // 7: dto.User.UsersUpdateSettings:input_type -> dto.UsersUpdateSettingsRequest
	2,  // 8: dto.User.UsersCreateUser:output_type -> dto.UsersCreateUserResponse
	4,  // 9: dto.User.UsersLogin:output_type -> dto.UsersLoginResponse
	6,  // 10: dto.User.UsersLogout:output_type -> dto.UsersLogoutResponse
	8,  // 11: dto.User.UsersChangePwd:output_type -> dto.UsersChangePwdResponse
	10, // 12: dto.User.UsersGetSettings:output_type -> dto.UsersGetSettingsResponse
	12, // 13: dto.User.UsersUpdateSettings:output_type -> dto.UsersUpdateSettingsResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersCreateUserRequest); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersCreateUserResponse); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersLoginRequest); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersLoginResponse); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersLogoutRequest); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersLogoutResponse); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersChangePwdRequest); i {
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
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersChangePwdResponse); i {
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
		file_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersGetSettingsRequest); i {
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
		file_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersGetSettingsResponse); i {
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
		file_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersUpdateSettingsRequest); i {
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
		file_user_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersUpdateSettingsResponse); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
