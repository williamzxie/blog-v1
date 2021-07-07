// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: sc.proto

package proto

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

type Response int32

const (
	Response_ResponseBeginIndex Response = 0
)

// Enum value maps for Response.
var (
	Response_name = map[int32]string{
		0: "ResponseBeginIndex",
	}
	Response_value = map[string]int32{
		"ResponseBeginIndex": 0,
	}
)

func (x Response) Enum() *Response {
	p := new(Response)
	*p = x
	return p
}

func (x Response) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Response) Descriptor() protoreflect.EnumDescriptor {
	return file_sc_proto_enumTypes[0].Descriptor()
}

func (Response) Type() protoreflect.EnumType {
	return &file_sc_proto_enumTypes[0]
}

func (x Response) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Response.Descriptor instead.
func (Response) EnumDescriptor() ([]byte, []int) {
	return file_sc_proto_rawDescGZIP(), []int{0}
}

type ResultCode int32

const (
	ResultCode_Success   ResultCode = 0 //协议请求成功，其余失败
	ResultCode_Fail      ResultCode = 1
	ResultCode_SuccessOK ResultCode = 10000
)

// Enum value maps for ResultCode.
var (
	ResultCode_name = map[int32]string{
		0:     "Success",
		1:     "Fail",
		10000: "SuccessOK",
	}
	ResultCode_value = map[string]int32{
		"Success":   0,
		"Fail":      1,
		"SuccessOK": 10000,
	}
)

func (x ResultCode) Enum() *ResultCode {
	p := new(ResultCode)
	*p = x
	return p
}

func (x ResultCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultCode) Descriptor() protoreflect.EnumDescriptor {
	return file_sc_proto_enumTypes[1].Descriptor()
}

func (ResultCode) Type() protoreflect.EnumType {
	return &file_sc_proto_enumTypes[1]
}

func (x ResultCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultCode.Descriptor instead.
func (ResultCode) EnumDescriptor() ([]byte, []int) {
	return file_sc_proto_rawDescGZIP(), []int{1}
}

type ResponsePkg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CmdId                   Response                 `protobuf:"varint,1,opt,name=cmdId,proto3,enum=proto.Response" json:"cmdId,omitempty"` // 协议ID
	Code                    ResultCode               `protobuf:"varint,2,opt,name=code,proto3,enum=proto.ResultCode" json:"code,omitempty"` //返回码
	Categories              []*Category              `protobuf:"bytes,3,rep,name=categories,proto3" json:"categories,omitempty"`
	ArticlesByCategoryOrTag *ArticlesByCategoryOrTag `protobuf:"bytes,4,opt,name=articlesByCategoryOrTag,proto3" json:"articlesByCategoryOrTag,omitempty"`
	Tags                    []*Tag                   `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Messages                []*Message               `protobuf:"bytes,6,rep,name=messages,proto3" json:"messages,omitempty"`
	FriendLinks             []*FriendLink            `protobuf:"bytes,7,rep,name=friendLinks,proto3" json:"friendLinks,omitempty"`
	CommentInfo             *CommentInfo             `protobuf:"bytes,8,opt,name=commentInfo,proto3" json:"commentInfo,omitempty"`
	LoginResponse           *LoginResponse           `protobuf:"bytes,9,opt,name=loginResponse,proto3" json:"loginResponse,omitempty"` //token
	Message                 string                   `protobuf:"bytes,10,opt,name=message,proto3" json:"message,omitempty"`            //消息
	ServerTime              int64                    `protobuf:"varint,11,opt,name=serverTime,proto3" json:"serverTime,omitempty"`     //服务器时间
	ArticleList             []*Article               `protobuf:"bytes,12,rep,name=articleList,proto3" json:"articleList,omitempty"`
	BlogHomeInfo            *BlogHomeInfo            `protobuf:"bytes,13,opt,name=blogHomeInfo,proto3" json:"blogHomeInfo,omitempty"`
	ArticleInfo             *ArticleInfo             `protobuf:"bytes,14,opt,name=articleInfo,proto3" json:"articleInfo,omitempty"`
	ArchiveInfo             *Archives                `protobuf:"bytes,15,opt,name=archiveInfo,proto3" json:"archiveInfo,omitempty"`
	About                   *About                   `protobuf:"bytes,16,opt,name=about,proto3" json:"about,omitempty"`
	ReplyList               []*Reply                 `protobuf:"bytes,17,rep,name=replyList,proto3" json:"replyList,omitempty"`
	ScChat                  *ScChat                  `protobuf:"bytes,18,opt,name=scChat,proto3" json:"scChat,omitempty"`
	UserMenu                []*ScUserMenuMessage     `protobuf:"bytes,19,rep,name=userMenu,proto3" json:"userMenu,omitempty"`
}

func (x *ResponsePkg) Reset() {
	*x = ResponsePkg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponsePkg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponsePkg) ProtoMessage() {}

func (x *ResponsePkg) ProtoReflect() protoreflect.Message {
	mi := &file_sc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponsePkg.ProtoReflect.Descriptor instead.
func (*ResponsePkg) Descriptor() ([]byte, []int) {
	return file_sc_proto_rawDescGZIP(), []int{0}
}

func (x *ResponsePkg) GetCmdId() Response {
	if x != nil {
		return x.CmdId
	}
	return Response_ResponseBeginIndex
}

func (x *ResponsePkg) GetCode() ResultCode {
	if x != nil {
		return x.Code
	}
	return ResultCode_Success
}

func (x *ResponsePkg) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *ResponsePkg) GetArticlesByCategoryOrTag() *ArticlesByCategoryOrTag {
	if x != nil {
		return x.ArticlesByCategoryOrTag
	}
	return nil
}

func (x *ResponsePkg) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ResponsePkg) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *ResponsePkg) GetFriendLinks() []*FriendLink {
	if x != nil {
		return x.FriendLinks
	}
	return nil
}

func (x *ResponsePkg) GetCommentInfo() *CommentInfo {
	if x != nil {
		return x.CommentInfo
	}
	return nil
}

func (x *ResponsePkg) GetLoginResponse() *LoginResponse {
	if x != nil {
		return x.LoginResponse
	}
	return nil
}

func (x *ResponsePkg) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResponsePkg) GetServerTime() int64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *ResponsePkg) GetArticleList() []*Article {
	if x != nil {
		return x.ArticleList
	}
	return nil
}

func (x *ResponsePkg) GetBlogHomeInfo() *BlogHomeInfo {
	if x != nil {
		return x.BlogHomeInfo
	}
	return nil
}

func (x *ResponsePkg) GetArticleInfo() *ArticleInfo {
	if x != nil {
		return x.ArticleInfo
	}
	return nil
}

func (x *ResponsePkg) GetArchiveInfo() *Archives {
	if x != nil {
		return x.ArchiveInfo
	}
	return nil
}

func (x *ResponsePkg) GetAbout() *About {
	if x != nil {
		return x.About
	}
	return nil
}

func (x *ResponsePkg) GetReplyList() []*Reply {
	if x != nil {
		return x.ReplyList
	}
	return nil
}

func (x *ResponsePkg) GetScChat() *ScChat {
	if x != nil {
		return x.ScChat
	}
	return nil
}

func (x *ResponsePkg) GetUserMenu() []*ScUserMenuMessage {
	if x != nil {
		return x.UserMenu
	}
	return nil
}

type ScChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          uint32         `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	ScChatOnline  *ScChatOnline  `protobuf:"bytes,2,opt,name=scChatOnline,proto3" json:"scChatOnline,omitempty"`
	ScChatMessage *ScChatMessage `protobuf:"bytes,3,opt,name=scChatMessage,proto3" json:"scChatMessage,omitempty"`
}

func (x *ScChat) Reset() {
	*x = ScChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScChat) ProtoMessage() {}

func (x *ScChat) ProtoReflect() protoreflect.Message {
	mi := &file_sc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScChat.ProtoReflect.Descriptor instead.
func (*ScChat) Descriptor() ([]byte, []int) {
	return file_sc_proto_rawDescGZIP(), []int{1}
}

func (x *ScChat) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ScChat) GetScChatOnline() *ScChatOnline {
	if x != nil {
		return x.ScChatOnline
	}
	return nil
}

func (x *ScChat) GetScChatMessage() *ScChatMessage {
	if x != nil {
		return x.ScChatMessage
	}
	return nil
}

type ScChatOnline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Online uint32 `protobuf:"varint,2,opt,name=online,proto3" json:"online,omitempty"`
}

func (x *ScChatOnline) Reset() {
	*x = ScChatOnline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScChatOnline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScChatOnline) ProtoMessage() {}

func (x *ScChatOnline) ProtoReflect() protoreflect.Message {
	mi := &file_sc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScChatOnline.ProtoReflect.Descriptor instead.
func (*ScChatOnline) Descriptor() ([]byte, []int) {
	return file_sc_proto_rawDescGZIP(), []int{2}
}

func (x *ScChatOnline) GetOnline() uint32 {
	if x != nil {
		return x.Online
	}
	return 0
}

type ScChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname   string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar     string `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Content    string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	UserId     uint32 `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
	Type       uint32 `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	IpAddr     string `protobuf:"bytes,6,opt,name=ipAddr,proto3" json:"ipAddr,omitempty"`
	IpSource   string `protobuf:"bytes,7,opt,name=ipSource,proto3" json:"ipSource,omitempty"`
	CreateTime int64  `protobuf:"varint,8,opt,name=createTime,proto3" json:"createTime,omitempty"`
}

func (x *ScChatMessage) Reset() {
	*x = ScChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScChatMessage) ProtoMessage() {}

func (x *ScChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_sc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScChatMessage.ProtoReflect.Descriptor instead.
func (*ScChatMessage) Descriptor() ([]byte, []int) {
	return file_sc_proto_rawDescGZIP(), []int{3}
}

func (x *ScChatMessage) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ScChatMessage) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *ScChatMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ScChatMessage) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ScChatMessage) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ScChatMessage) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *ScChatMessage) GetIpSource() string {
	if x != nil {
		return x.IpSource
	}
	return ""
}

func (x *ScChatMessage) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type ScUserMenuMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path      string               `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Component string               `protobuf:"bytes,3,opt,name=component,proto3" json:"component,omitempty"`
	Icon      string               `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	IsHidden  bool                 `protobuf:"varint,5,opt,name=isHidden,proto3" json:"isHidden,omitempty"`
	Children  []*ScUserMenuMessage `protobuf:"bytes,6,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *ScUserMenuMessage) Reset() {
	*x = ScUserMenuMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScUserMenuMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScUserMenuMessage) ProtoMessage() {}

func (x *ScUserMenuMessage) ProtoReflect() protoreflect.Message {
	mi := &file_sc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScUserMenuMessage.ProtoReflect.Descriptor instead.
func (*ScUserMenuMessage) Descriptor() ([]byte, []int) {
	return file_sc_proto_rawDescGZIP(), []int{4}
}

func (x *ScUserMenuMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScUserMenuMessage) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ScUserMenuMessage) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *ScUserMenuMessage) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *ScUserMenuMessage) GetIsHidden() bool {
	if x != nil {
		return x.IsHidden
	}
	return false
}

func (x *ScUserMenuMessage) GetChildren() []*ScUserMenuMessage {
	if x != nil {
		return x.Children
	}
	return nil
}

var File_sc_proto protoreflect.FileDescriptor

var file_sc_proto_rawDesc = []byte{
	0x0a, 0x08, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x07,
	0x0a, 0x0b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x6b, 0x67, 0x12, 0x25, 0x0a,
	0x05, 0x63, 0x6d, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x63,
	0x6d, 0x64, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x17,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x4f, 0x72, 0x54, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x42, 0x79,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4f, 0x72, 0x54, 0x61, 0x67, 0x52, 0x17, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x4f, 0x72, 0x54, 0x61, 0x67, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x67,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x33, 0x0a, 0x0b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0b, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a,
	0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0d, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x67, 0x48, 0x6f, 0x6d,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x48, 0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0c, 0x62, 0x6c, 0x6f, 0x67, 0x48, 0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34,
	0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x0b, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x73, 0x52, 0x0b, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x05, 0x61, 0x62, 0x6f, 0x75, 0x74,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x62, 0x6f, 0x75, 0x74, 0x52, 0x05, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x09, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x09, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x63, 0x43, 0x68, 0x61,
	0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x63, 0x43, 0x68, 0x61, 0x74, 0x52, 0x06, 0x73, 0x63, 0x43, 0x68, 0x61, 0x74, 0x12, 0x34,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x6e, 0x75, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x65, 0x6e, 0x75, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4d, 0x65, 0x6e, 0x75, 0x22, 0x91, 0x01, 0x0a, 0x06, 0x53, 0x63, 0x43, 0x68, 0x61, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x73, 0x63, 0x43, 0x68, 0x61, 0x74, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x63, 0x43, 0x68, 0x61, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x0c,
	0x73, 0x63, 0x43, 0x68, 0x61, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x3a, 0x0a, 0x0d,
	0x73, 0x63, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x43, 0x68,
	0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x73, 0x63, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x0c, 0x53, 0x63, 0x43, 0x68,
	0x61, 0x74, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x22, 0xdd, 0x01, 0x0a, 0x0d, 0x53, 0x63, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0xbf, 0x01, 0x0a, 0x11, 0x53, 0x63, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x6e, 0x75, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x34, 0x0a, 0x08,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x6e,
	0x75, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x2a, 0x22, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x10, 0x00, 0x2a, 0x33, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x09, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4f, 0x4b, 0x10, 0x90, 0x4e, 0x42, 0x13, 0x5a, 0x11, 0x2e,
	0x2e, 0x2f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sc_proto_rawDescOnce sync.Once
	file_sc_proto_rawDescData = file_sc_proto_rawDesc
)

func file_sc_proto_rawDescGZIP() []byte {
	file_sc_proto_rawDescOnce.Do(func() {
		file_sc_proto_rawDescData = protoimpl.X.CompressGZIP(file_sc_proto_rawDescData)
	})
	return file_sc_proto_rawDescData
}

var file_sc_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_sc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sc_proto_goTypes = []interface{}{
	(Response)(0),                   // 0: proto.Response
	(ResultCode)(0),                 // 1: proto.ResultCode
	(*ResponsePkg)(nil),             // 2: proto.ResponsePkg
	(*ScChat)(nil),                  // 3: proto.ScChat
	(*ScChatOnline)(nil),            // 4: proto.ScChatOnline
	(*ScChatMessage)(nil),           // 5: proto.ScChatMessage
	(*ScUserMenuMessage)(nil),       // 6: proto.ScUserMenuMessage
	(*Category)(nil),                // 7: proto.Category
	(*ArticlesByCategoryOrTag)(nil), // 8: proto.ArticlesByCategoryOrTag
	(*Tag)(nil),                     // 9: proto.Tag
	(*Message)(nil),                 // 10: proto.Message
	(*FriendLink)(nil),              // 11: proto.FriendLink
	(*CommentInfo)(nil),             // 12: proto.CommentInfo
	(*LoginResponse)(nil),           // 13: proto.LoginResponse
	(*Article)(nil),                 // 14: proto.Article
	(*BlogHomeInfo)(nil),            // 15: proto.BlogHomeInfo
	(*ArticleInfo)(nil),             // 16: proto.ArticleInfo
	(*Archives)(nil),                // 17: proto.Archives
	(*About)(nil),                   // 18: proto.About
	(*Reply)(nil),                   // 19: proto.Reply
}
var file_sc_proto_depIdxs = []int32{
	0,  // 0: proto.ResponsePkg.cmdId:type_name -> proto.Response
	1,  // 1: proto.ResponsePkg.code:type_name -> proto.ResultCode
	7,  // 2: proto.ResponsePkg.categories:type_name -> proto.Category
	8,  // 3: proto.ResponsePkg.articlesByCategoryOrTag:type_name -> proto.ArticlesByCategoryOrTag
	9,  // 4: proto.ResponsePkg.tags:type_name -> proto.Tag
	10, // 5: proto.ResponsePkg.messages:type_name -> proto.Message
	11, // 6: proto.ResponsePkg.friendLinks:type_name -> proto.FriendLink
	12, // 7: proto.ResponsePkg.commentInfo:type_name -> proto.CommentInfo
	13, // 8: proto.ResponsePkg.loginResponse:type_name -> proto.LoginResponse
	14, // 9: proto.ResponsePkg.articleList:type_name -> proto.Article
	15, // 10: proto.ResponsePkg.blogHomeInfo:type_name -> proto.BlogHomeInfo
	16, // 11: proto.ResponsePkg.articleInfo:type_name -> proto.ArticleInfo
	17, // 12: proto.ResponsePkg.archiveInfo:type_name -> proto.Archives
	18, // 13: proto.ResponsePkg.about:type_name -> proto.About
	19, // 14: proto.ResponsePkg.replyList:type_name -> proto.Reply
	3,  // 15: proto.ResponsePkg.scChat:type_name -> proto.ScChat
	6,  // 16: proto.ResponsePkg.userMenu:type_name -> proto.ScUserMenuMessage
	4,  // 17: proto.ScChat.scChatOnline:type_name -> proto.ScChatOnline
	5,  // 18: proto.ScChat.scChatMessage:type_name -> proto.ScChatMessage
	6,  // 19: proto.ScUserMenuMessage.children:type_name -> proto.ScUserMenuMessage
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_sc_proto_init() }
func file_sc_proto_init() {
	if File_sc_proto != nil {
		return
	}
	file_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponsePkg); i {
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
		file_sc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScChat); i {
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
		file_sc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScChatOnline); i {
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
		file_sc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScChatMessage); i {
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
		file_sc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScUserMenuMessage); i {
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
			RawDescriptor: file_sc_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sc_proto_goTypes,
		DependencyIndexes: file_sc_proto_depIdxs,
		EnumInfos:         file_sc_proto_enumTypes,
		MessageInfos:      file_sc_proto_msgTypes,
	}.Build()
	File_sc_proto = out.File
	file_sc_proto_rawDesc = nil
	file_sc_proto_goTypes = nil
	file_sc_proto_depIdxs = nil
}
