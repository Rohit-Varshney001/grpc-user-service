// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = proto.Unmarshal
var _ = proto.UnmarshalField
var _ = proto.Reset
var _ = proto.XXX_Unmarshal

type User struct {
	Id                  int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fname               string  `protobuf:"bytes,2,opt,name=fname,proto3" json:"fname,omitempty"`
	City                string  `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Phone               int64   `protobuf:"varint,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Height              float32 `protobuf:"fixed32,5,opt,name=height,proto3" json:"height,omitempty"`
	Married             bool    `protobuf:"varint,6,opt,name=married,proto3" json:"married,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e212b40445e78d4, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetFname() string {
	if m != nil {
		return m.Fname
	}
	return ""
}

func (m *User) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *User) GetPhone() int64 {
	if m != nil {
		return m.Phone
	}
	return 0
}

func (m *User) GetHeight() float32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *User) GetMarried() bool {
	if m != nil {
		return m.Married
	}
	return false
}

// UserIdRequest represents a request to get a user by ID.
type UserIdRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserIdRequest) Reset()         { *m = UserIdRequest{} }
func (m *UserIdRequest) String() string { return proto.CompactTextString(m) }
func (*UserIdRequest) ProtoMessage()    {}
func (*UserIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e212b40445e78d4, []int{1}
}

func (m *UserIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIdRequest.Unmarshal(m, b)
}
func (m *UserIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIdRequest.Marshal(b, m, deterministic)
}
func (m *UserIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIdRequest.Merge(m, src)
}
func (m *UserIdRequest) XXX_Size() int {
	return xxx_messageInfo_UserIdRequest.Size(m)
}
func (m *UserIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserIdRequest proto.InternalMessageInfo

func (m *UserIdRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// UserIdsRequest represents a request to get multiple users by IDs.
type UserIdsRequest struct {
	Ids                  []int32  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserIdsRequest) Reset()         { *m = UserIdsRequest{} }
func (m *UserIdsRequest) String() string { return proto.CompactTextString(m) }
func (*UserIdsRequest) ProtoMessage()    {}
func (*UserIdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e212b40445e78d4, []int{2}
}

func (m *UserIdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIdsRequest.Unmarshal(m, b)
}
func (m *UserIdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIdsRequest.Marshal(b, m, deterministic)
}
func (m *UserIdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIdsRequest.Merge(m, src)
}
func (m *UserIdsRequest) XXX_Size() int {
	return xxx_messageInfo_UserIdsRequest.Size(m)
}
func (m *UserIdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserIdsRequest proto.InternalMessageInfo

func (m *UserIdsRequest) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

// UserSearchRequest represents a request to search for users based on criteria.
type UserSearchRequest struct {
	Fname                string   `protobuf:"bytes,1,opt,name=fname,proto3" json:"fname,omitempty"`
	City                 string   `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"` // Ensures that this field is excluded from JSON serialization
}

var xxx_messageInfo_UserSearchRequest proto.InternalMessageInfo

func (m *UserSearchRequest) Reset()         { *m = UserSearchRequest{} }
func (m *UserSearchRequest) String() string { return proto.CompactTextString(m) }
func (*UserSearchRequest) ProtoMessage()    {}
func (*UserSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e212b40445e78d4, []int{3}
}

func (m *UserSearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSearchRequest.Unmarshal(m, b)
}
func (m *UserSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSearchRequest.Marshal(b, m, deterministic)
}
func (m *UserSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSearchRequest.Merge(m, src)
}
func (m *UserSearchRequest) XXX_Size() int {
	return xxx_messageInfo_UserSearchRequest.Size(m)
}
func (m *UserSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSearchRequest.DiscardUnknown(m)
}

func (m *UserSearchRequest) GetFname() string {
	if m != nil {
		return m.Fname
	}
	return ""
}

func (m *UserSearchRequest) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

// Users represents a collection of User messages.
type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e212b40445e78d4, []int{4}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*UserIdRequest)(nil), "user.UserIdRequest")
	proto.RegisterType((*UserIdsRequest)(nil), "user.UserIdsRequest")
	proto.RegisterType((*UserSearchRequest)(nil), "user.UserSearchRequest")
	proto.RegisterType((*Users)(nil), "user.Users")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ interface{} = (*User)(nil)
var _ interface{} = (*UserIdRequest)(nil)
var _ interface{} = (*UserIdsRequest)(nil)
var _ interface{} = (*UserSearchRequest)(nil)
var _ interface{} = (*Users)(nil)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3
