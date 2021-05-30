// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: chat/v1/profile.proto

package v1

import (
	v1 "github.com/harmony-development/chapati/gen/harmonytypes/v1"
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

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_profile_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName   string        `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserAvatar string        `protobuf:"bytes,2,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
	UserStatus v1.UserStatus `protobuf:"varint,3,opt,name=user_status,json=userStatus,proto3,enum=protocol.harmonytypes.v1.UserStatus" json:"user_status,omitempty"`
	IsBot      bool          `protobuf:"varint,4,opt,name=is_bot,json=isBot,proto3" json:"is_bot,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_chat_v1_profile_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetUserResponse) GetUserAvatar() string {
	if x != nil {
		return x.UserAvatar
	}
	return ""
}

func (x *GetUserResponse) GetUserStatus() v1.UserStatus {
	if x != nil {
		return x.UserStatus
	}
	return v1.UserStatus_USER_STATUS_ONLINE_UNSPECIFIED
}

func (x *GetUserResponse) GetIsBot() bool {
	if x != nil {
		return x.IsBot
	}
	return false
}

type GetUserBulkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []uint64 `protobuf:"varint,1,rep,packed,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *GetUserBulkRequest) Reset() {
	*x = GetUserBulkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBulkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBulkRequest) ProtoMessage() {}

func (x *GetUserBulkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBulkRequest.ProtoReflect.Descriptor instead.
func (*GetUserBulkRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_profile_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserBulkRequest) GetUserIds() []uint64 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type GetUserBulkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*GetUserResponse `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUserBulkResponse) Reset() {
	*x = GetUserBulkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserBulkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserBulkResponse) ProtoMessage() {}

func (x *GetUserBulkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserBulkResponse.ProtoReflect.Descriptor instead.
func (*GetUserBulkResponse) Descriptor() ([]byte, []int) {
	return file_chat_v1_profile_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserBulkResponse) GetUsers() []*GetUserResponse {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetUserMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *GetUserMetadataRequest) Reset() {
	*x = GetUserMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserMetadataRequest) ProtoMessage() {}

func (x *GetUserMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserMetadataRequest.ProtoReflect.Descriptor instead.
func (*GetUserMetadataRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_profile_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserMetadataRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type GetUserMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata string `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *GetUserMetadataResponse) Reset() {
	*x = GetUserMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserMetadataResponse) ProtoMessage() {}

func (x *GetUserMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserMetadataResponse.ProtoReflect.Descriptor instead.
func (*GetUserMetadataResponse) Descriptor() ([]byte, []int) {
	return file_chat_v1_profile_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserMetadataResponse) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type ProfileUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewUsername    string        `protobuf:"bytes,1,opt,name=new_username,json=newUsername,proto3" json:"new_username,omitempty"`
	UpdateUsername bool          `protobuf:"varint,2,opt,name=update_username,json=updateUsername,proto3" json:"update_username,omitempty"`
	NewAvatar      string        `protobuf:"bytes,3,opt,name=new_avatar,json=newAvatar,proto3" json:"new_avatar,omitempty"`
	UpdateAvatar   bool          `protobuf:"varint,4,opt,name=update_avatar,json=updateAvatar,proto3" json:"update_avatar,omitempty"`
	NewStatus      v1.UserStatus `protobuf:"varint,5,opt,name=new_status,json=newStatus,proto3,enum=protocol.harmonytypes.v1.UserStatus" json:"new_status,omitempty"`
	UpdateStatus   bool          `protobuf:"varint,6,opt,name=update_status,json=updateStatus,proto3" json:"update_status,omitempty"`
	IsBot          bool          `protobuf:"varint,7,opt,name=is_bot,json=isBot,proto3" json:"is_bot,omitempty"`
	UpdateIsBot    bool          `protobuf:"varint,8,opt,name=update_is_bot,json=updateIsBot,proto3" json:"update_is_bot,omitempty"`
}

func (x *ProfileUpdateRequest) Reset() {
	*x = ProfileUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileUpdateRequest) ProtoMessage() {}

func (x *ProfileUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileUpdateRequest.ProtoReflect.Descriptor instead.
func (*ProfileUpdateRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_profile_proto_rawDescGZIP(), []int{6}
}

func (x *ProfileUpdateRequest) GetNewUsername() string {
	if x != nil {
		return x.NewUsername
	}
	return ""
}

func (x *ProfileUpdateRequest) GetUpdateUsername() bool {
	if x != nil {
		return x.UpdateUsername
	}
	return false
}

func (x *ProfileUpdateRequest) GetNewAvatar() string {
	if x != nil {
		return x.NewAvatar
	}
	return ""
}

func (x *ProfileUpdateRequest) GetUpdateAvatar() bool {
	if x != nil {
		return x.UpdateAvatar
	}
	return false
}

func (x *ProfileUpdateRequest) GetNewStatus() v1.UserStatus {
	if x != nil {
		return x.NewStatus
	}
	return v1.UserStatus_USER_STATUS_ONLINE_UNSPECIFIED
}

func (x *ProfileUpdateRequest) GetUpdateStatus() bool {
	if x != nil {
		return x.UpdateStatus
	}
	return false
}

func (x *ProfileUpdateRequest) GetIsBot() bool {
	if x != nil {
		return x.IsBot
	}
	return false
}

func (x *ProfileUpdateRequest) GetUpdateIsBot() bool {
	if x != nil {
		return x.UpdateIsBot
	}
	return false
}

var File_chat_v1_profile_proto protoreflect.FileDescriptor

var file_chat_v1_profile_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x68, 0x61, 0x72, 0x6d, 0x6f,
	0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0xad, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x45, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73,
	0x5f, 0x62, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x42, 0x6f,
	0x74, 0x22, 0x2f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x75, 0x6c, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x22, 0x4e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x75, 0x6c,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x2f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xcb, 0x02, 0x0a, 0x14, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x23,
	0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x43, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x6e,
	0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a,
	0x06, 0x69, 0x73, 0x5f, 0x62, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69,
	0x73, 0x42, 0x6f, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x73, 0x5f, 0x62, 0x6f, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x73, 0x42, 0x6f, 0x74, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x79, 0x2d, 0x64,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x70, 0x61,
	0x74, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_v1_profile_proto_rawDescOnce sync.Once
	file_chat_v1_profile_proto_rawDescData = file_chat_v1_profile_proto_rawDesc
)

func file_chat_v1_profile_proto_rawDescGZIP() []byte {
	file_chat_v1_profile_proto_rawDescOnce.Do(func() {
		file_chat_v1_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_v1_profile_proto_rawDescData)
	})
	return file_chat_v1_profile_proto_rawDescData
}

var file_chat_v1_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_chat_v1_profile_proto_goTypes = []interface{}{
	(*GetUserRequest)(nil),          // 0: protocol.chat.v1.GetUserRequest
	(*GetUserResponse)(nil),         // 1: protocol.chat.v1.GetUserResponse
	(*GetUserBulkRequest)(nil),      // 2: protocol.chat.v1.GetUserBulkRequest
	(*GetUserBulkResponse)(nil),     // 3: protocol.chat.v1.GetUserBulkResponse
	(*GetUserMetadataRequest)(nil),  // 4: protocol.chat.v1.GetUserMetadataRequest
	(*GetUserMetadataResponse)(nil), // 5: protocol.chat.v1.GetUserMetadataResponse
	(*ProfileUpdateRequest)(nil),    // 6: protocol.chat.v1.ProfileUpdateRequest
	(v1.UserStatus)(0),              // 7: protocol.harmonytypes.v1.UserStatus
}
var file_chat_v1_profile_proto_depIdxs = []int32{
	7, // 0: protocol.chat.v1.GetUserResponse.user_status:type_name -> protocol.harmonytypes.v1.UserStatus
	1, // 1: protocol.chat.v1.GetUserBulkResponse.users:type_name -> protocol.chat.v1.GetUserResponse
	7, // 2: protocol.chat.v1.ProfileUpdateRequest.new_status:type_name -> protocol.harmonytypes.v1.UserStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chat_v1_profile_proto_init() }
func file_chat_v1_profile_proto_init() {
	if File_chat_v1_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_v1_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_chat_v1_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
		file_chat_v1_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserBulkRequest); i {
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
		file_chat_v1_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserBulkResponse); i {
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
		file_chat_v1_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserMetadataRequest); i {
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
		file_chat_v1_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserMetadataResponse); i {
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
		file_chat_v1_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileUpdateRequest); i {
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
			RawDescriptor: file_chat_v1_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_v1_profile_proto_goTypes,
		DependencyIndexes: file_chat_v1_profile_proto_depIdxs,
		MessageInfos:      file_chat_v1_profile_proto_msgTypes,
	}.Build()
	File_chat_v1_profile_proto = out.File
	file_chat_v1_profile_proto_rawDesc = nil
	file_chat_v1_profile_proto_goTypes = nil
	file_chat_v1_profile_proto_depIdxs = nil
}