// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: helix_api.proto

package rpc

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

type UsersAndGamesParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIDs []string `protobuf:"bytes,1,rep,name=UserIDs,proto3" json:"UserIDs,omitempty"`
	GameIDs []string `protobuf:"bytes,2,rep,name=GameIDs,proto3" json:"GameIDs,omitempty"`
}

func (x *UsersAndGamesParams) Reset() {
	*x = UsersAndGamesParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helix_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersAndGamesParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersAndGamesParams) ProtoMessage() {}

func (x *UsersAndGamesParams) ProtoReflect() protoreflect.Message {
	mi := &file_helix_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersAndGamesParams.ProtoReflect.Descriptor instead.
func (*UsersAndGamesParams) Descriptor() ([]byte, []int) {
	return file_helix_api_proto_rawDescGZIP(), []int{0}
}

func (x *UsersAndGamesParams) GetUserIDs() []string {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

func (x *UsersAndGamesParams) GetGameIDs() []string {
	if x != nil {
		return x.GameIDs
	}
	return nil
}

type UsersAndGames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users *Users `protobuf:"bytes,1,opt,name=Users,proto3" json:"Users,omitempty"`
	Games *Games `protobuf:"bytes,2,opt,name=Games,proto3" json:"Games,omitempty"`
}

func (x *UsersAndGames) Reset() {
	*x = UsersAndGames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helix_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersAndGames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersAndGames) ProtoMessage() {}

func (x *UsersAndGames) ProtoReflect() protoreflect.Message {
	mi := &file_helix_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersAndGames.ProtoReflect.Descriptor instead.
func (*UsersAndGames) Descriptor() ([]byte, []int) {
	return file_helix_api_proto_rawDescGZIP(), []int{1}
}

func (x *UsersAndGames) GetUsers() *Users {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *UsersAndGames) GetGames() *Games {
	if x != nil {
		return x.Games
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Login           string `protobuf:"bytes,2,opt,name=Login,proto3" json:"Login,omitempty"`
	DisplayName     string `protobuf:"bytes,3,opt,name=DisplayName,proto3" json:"DisplayName,omitempty"`
	Description     string `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	ProfileImageURL string `protobuf:"bytes,5,opt,name=ProfileImageURL,proto3" json:"ProfileImageURL,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helix_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_helix_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_helix_api_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *User) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *User) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *User) GetProfileImageURL() string {
	if x != nil {
		return x.ProfileImageURL
	}
	return ""
}

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	BoxArtURL string `protobuf:"bytes,3,opt,name=BoxArtURL,proto3" json:"BoxArtURL,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helix_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_helix_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_helix_api_proto_rawDescGZIP(), []int{3}
}

func (x *Game) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Game) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Game) GetBoxArtURL() string {
	if x != nil {
		return x.BoxArtURL
	}
	return ""
}

type Stream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID   string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	UserName string `protobuf:"bytes,3,opt,name=UserName,proto3" json:"UserName,omitempty"`
	GameID   string `protobuf:"bytes,4,opt,name=GameID,proto3" json:"GameID,omitempty"`
	Title    string `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *Stream) Reset() {
	*x = Stream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helix_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream) ProtoMessage() {}

func (x *Stream) ProtoReflect() protoreflect.Message {
	mi := &file_helix_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream.ProtoReflect.Descriptor instead.
func (*Stream) Descriptor() ([]byte, []int) {
	return file_helix_api_proto_rawDescGZIP(), []int{4}
}

func (x *Stream) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Stream) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Stream) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Stream) GetGameID() string {
	if x != nil {
		return x.GameID
	}
	return ""
}

func (x *Stream) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	DisplayName  string `protobuf:"bytes,2,opt,name=DisplayName,proto3" json:"DisplayName,omitempty"`
	Title        string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	ThumbnailURL string `protobuf:"bytes,4,opt,name=ThumbnailURL,proto3" json:"ThumbnailURL,omitempty"`
	IsLive       bool   `protobuf:"varint,5,opt,name=IsLive,proto3" json:"IsLive,omitempty"`
}

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helix_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_helix_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_helix_api_proto_rawDescGZIP(), []int{5}
}

func (x *Channel) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Channel) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Channel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Channel) GetThumbnailURL() string {
	if x != nil {
		return x.ThumbnailURL
	}
	return ""
}

func (x *Channel) GetIsLive() bool {
	if x != nil {
		return x.IsLive
	}
	return false
}

type MultiQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDs   []string `protobuf:"bytes,1,rep,name=IDs,proto3" json:"IDs,omitempty"`
	Names []string `protobuf:"bytes,2,rep,name=Names,proto3" json:"Names,omitempty"`
}

func (x *MultiQuery) Reset() {
	*x = MultiQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helix_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiQuery) ProtoMessage() {}

func (x *MultiQuery) ProtoReflect() protoreflect.Message {
	mi := &file_helix_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiQuery.ProtoReflect.Descriptor instead.
func (*MultiQuery) Descriptor() ([]byte, []int) {
	return file_helix_api_proto_rawDescGZIP(), []int{6}
}

func (x *MultiQuery) GetIDs() []string {
	if x != nil {
		return x.IDs
	}
	return nil
}

func (x *MultiQuery) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type ChannelSearchQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	LiveOnly bool   `protobuf:"varint,2,opt,name=LiveOnly,proto3" json:"LiveOnly,omitempty"`
}

func (x *ChannelSearchQuery) Reset() {
	*x = ChannelSearchQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helix_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelSearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelSearchQuery) ProtoMessage() {}

func (x *ChannelSearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_helix_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelSearchQuery.ProtoReflect.Descriptor instead.
func (*ChannelSearchQuery) Descriptor() ([]byte, []int) {
	return file_helix_api_proto_rawDescGZIP(), []int{7}
}

func (x *ChannelSearchQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChannelSearchQuery) GetLiveOnly() bool {
	if x != nil {
		return x.LiveOnly
	}
	return false
}

type Users struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
}

func (x *Users) Reset() {
	*x = Users{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helix_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Users) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Users) ProtoMessage() {}

func (x *Users) ProtoReflect() protoreflect.Message {
	mi := &file_helix_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Users.ProtoReflect.Descriptor instead.
func (*Users) Descriptor() ([]byte, []int) {
	return file_helix_api_proto_rawDescGZIP(), []int{8}
}

func (x *Users) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type Games struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Games []*Game `protobuf:"bytes,1,rep,name=Games,proto3" json:"Games,omitempty"`
}

func (x *Games) Reset() {
	*x = Games{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helix_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Games) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Games) ProtoMessage() {}

func (x *Games) ProtoReflect() protoreflect.Message {
	mi := &file_helix_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Games.ProtoReflect.Descriptor instead.
func (*Games) Descriptor() ([]byte, []int) {
	return file_helix_api_proto_rawDescGZIP(), []int{9}
}

func (x *Games) GetGames() []*Game {
	if x != nil {
		return x.Games
	}
	return nil
}

type Streams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Streams []*Stream `protobuf:"bytes,1,rep,name=Streams,proto3" json:"Streams,omitempty"`
}

func (x *Streams) Reset() {
	*x = Streams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helix_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Streams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Streams) ProtoMessage() {}

func (x *Streams) ProtoReflect() protoreflect.Message {
	mi := &file_helix_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Streams.ProtoReflect.Descriptor instead.
func (*Streams) Descriptor() ([]byte, []int) {
	return file_helix_api_proto_rawDescGZIP(), []int{10}
}

func (x *Streams) GetStreams() []*Stream {
	if x != nil {
		return x.Streams
	}
	return nil
}

type Channels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channels []*Channel `protobuf:"bytes,1,rep,name=Channels,proto3" json:"Channels,omitempty"`
}

func (x *Channels) Reset() {
	*x = Channels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helix_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channels) ProtoMessage() {}

func (x *Channels) ProtoReflect() protoreflect.Message {
	mi := &file_helix_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channels.ProtoReflect.Descriptor instead.
func (*Channels) Descriptor() ([]byte, []int) {
	return file_helix_api_proto_rawDescGZIP(), []int{11}
}

func (x *Channels) GetChannels() []*Channel {
	if x != nil {
		return x.Channels
	}
	return nil
}

type StatusParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusParams) Reset() {
	*x = StatusParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helix_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusParams) ProtoMessage() {}

func (x *StatusParams) ProtoReflect() protoreflect.Message {
	mi := &file_helix_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusParams.ProtoReflect.Descriptor instead.
func (*StatusParams) Descriptor() ([]byte, []int) {
	return file_helix_api_proto_rawDescGZIP(), []int{12}
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alloc     uint64 `protobuf:"varint,1,opt,name=Alloc,proto3" json:"Alloc,omitempty"`
	Sys       uint64 `protobuf:"varint,2,opt,name=Sys,proto3" json:"Sys,omitempty"`
	NumGC     uint32 `protobuf:"varint,3,opt,name=NumGC,proto3" json:"NumGC,omitempty"`
	CacheSize uint32 `protobuf:"varint,4,opt,name=CacheSize,proto3" json:"CacheSize,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helix_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_helix_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_helix_api_proto_rawDescGZIP(), []int{13}
}

func (x *Status) GetAlloc() uint64 {
	if x != nil {
		return x.Alloc
	}
	return 0
}

func (x *Status) GetSys() uint64 {
	if x != nil {
		return x.Sys
	}
	return 0
}

func (x *Status) GetNumGC() uint32 {
	if x != nil {
		return x.NumGC
	}
	return 0
}

func (x *Status) GetCacheSize() uint32 {
	if x != nil {
		return x.CacheSize
	}
	return 0
}

var File_helix_api_proto protoreflect.FileDescriptor

var file_helix_api_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x65, 0x6c, 0x69, 0x78, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0x49, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41,
	0x6e, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x61, 0x6d, 0x65, 0x49,
	0x44, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44,
	0x73, 0x22, 0x53, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x6e, 0x64, 0x47, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x20, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x05, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x05, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x52,
	0x05, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x52, 0x4c, 0x22, 0x48, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x42, 0x6f, 0x78, 0x41, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x42, 0x6f, 0x78, 0x41, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x22, 0x7a, 0x0a,
	0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x47,
	0x61, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x61, 0x6d,
	0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x07, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x52, 0x4c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x52,
	0x4c, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x4c, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x49, 0x73, 0x4c, 0x69, 0x76, 0x65, 0x22, 0x34, 0x0a, 0x0a, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x44, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x49, 0x44, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22,
	0x44, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x69, 0x76,
	0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x4c, 0x69, 0x76,
	0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0x28, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1f,
	0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22,
	0x28, 0x0a, 0x05, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x47, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x05, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x07, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x12, 0x25, 0x0a, 0x07, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x07, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x34, 0x0a, 0x08, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x28, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x08, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x22, 0x64, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x41, 0x6c, 0x6c, 0x6f,
	0x63, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x53, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x75, 0x6d, 0x47, 0x43, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x4e, 0x75, 0x6d, 0x47, 0x43, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x32, 0xb5, 0x02, 0x0a, 0x0b, 0x54, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x41, 0x6e, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x6e, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x41, 0x6e, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x27, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x0f,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0c, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x38, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x17, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x1a, 0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x73, 0x12, 0x2b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6e,
	0x73, 0x67, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2d, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x3b, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_helix_api_proto_rawDescOnce sync.Once
	file_helix_api_proto_rawDescData = file_helix_api_proto_rawDesc
)

func file_helix_api_proto_rawDescGZIP() []byte {
	file_helix_api_proto_rawDescOnce.Do(func() {
		file_helix_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_helix_api_proto_rawDescData)
	})
	return file_helix_api_proto_rawDescData
}

var file_helix_api_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_helix_api_proto_goTypes = []interface{}{
	(*UsersAndGamesParams)(nil), // 0: rpc.UsersAndGamesParams
	(*UsersAndGames)(nil),       // 1: rpc.UsersAndGames
	(*User)(nil),                // 2: rpc.User
	(*Game)(nil),                // 3: rpc.Game
	(*Stream)(nil),              // 4: rpc.Stream
	(*Channel)(nil),             // 5: rpc.Channel
	(*MultiQuery)(nil),          // 6: rpc.MultiQuery
	(*ChannelSearchQuery)(nil),  // 7: rpc.ChannelSearchQuery
	(*Users)(nil),               // 8: rpc.Users
	(*Games)(nil),               // 9: rpc.Games
	(*Streams)(nil),             // 10: rpc.Streams
	(*Channels)(nil),            // 11: rpc.Channels
	(*StatusParams)(nil),        // 12: rpc.StatusParams
	(*Status)(nil),              // 13: rpc.Status
}
var file_helix_api_proto_depIdxs = []int32{
	8,  // 0: rpc.UsersAndGames.Users:type_name -> rpc.Users
	9,  // 1: rpc.UsersAndGames.Games:type_name -> rpc.Games
	2,  // 2: rpc.Users.Users:type_name -> rpc.User
	3,  // 3: rpc.Games.Games:type_name -> rpc.Game
	4,  // 4: rpc.Streams.Streams:type_name -> rpc.Stream
	5,  // 5: rpc.Channels.Channels:type_name -> rpc.Channel
	0,  // 6: rpc.TwitchCache.GetUsersAndGames:input_type -> rpc.UsersAndGamesParams
	6,  // 7: rpc.TwitchCache.GetUsers:input_type -> rpc.MultiQuery
	6,  // 8: rpc.TwitchCache.GetGames:input_type -> rpc.MultiQuery
	6,  // 9: rpc.TwitchCache.GetStreams:input_type -> rpc.MultiQuery
	7,  // 10: rpc.TwitchCache.SearchChannels:input_type -> rpc.ChannelSearchQuery
	12, // 11: rpc.TwitchCache.GetStatus:input_type -> rpc.StatusParams
	1,  // 12: rpc.TwitchCache.GetUsersAndGames:output_type -> rpc.UsersAndGames
	8,  // 13: rpc.TwitchCache.GetUsers:output_type -> rpc.Users
	9,  // 14: rpc.TwitchCache.GetGames:output_type -> rpc.Games
	10, // 15: rpc.TwitchCache.GetStreams:output_type -> rpc.Streams
	11, // 16: rpc.TwitchCache.SearchChannels:output_type -> rpc.Channels
	13, // 17: rpc.TwitchCache.GetStatus:output_type -> rpc.Status
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_helix_api_proto_init() }
func file_helix_api_proto_init() {
	if File_helix_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helix_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersAndGamesParams); i {
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
		file_helix_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersAndGames); i {
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
		file_helix_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_helix_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Game); i {
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
		file_helix_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stream); i {
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
		file_helix_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channel); i {
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
		file_helix_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiQuery); i {
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
		file_helix_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelSearchQuery); i {
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
		file_helix_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Users); i {
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
		file_helix_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Games); i {
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
		file_helix_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Streams); i {
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
		file_helix_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channels); i {
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
		file_helix_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusParams); i {
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
		file_helix_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_helix_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helix_api_proto_goTypes,
		DependencyIndexes: file_helix_api_proto_depIdxs,
		MessageInfos:      file_helix_api_proto_msgTypes,
	}.Build()
	File_helix_api_proto = out.File
	file_helix_api_proto_rawDesc = nil
	file_helix_api_proto_goTypes = nil
	file_helix_api_proto_depIdxs = nil
}
