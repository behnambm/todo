// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: protobuf/todo.proto

package todoservice

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

type ItemMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Priority int64  `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *ItemMessage) Reset() {
	*x = ItemMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemMessage) ProtoMessage() {}

func (x *ItemMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemMessage.ProtoReflect.Descriptor instead.
func (*ItemMessage) Descriptor() ([]byte, []int) {
	return file_protobuf_todo_proto_rawDescGZIP(), []int{0}
}

func (x *ItemMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ItemMessage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ItemMessage) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type TodoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *TodoMessage) Reset() {
	*x = TodoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoMessage) ProtoMessage() {}

func (x *TodoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoMessage.ProtoReflect.Descriptor instead.
func (*TodoMessage) Descriptor() ([]byte, []int) {
	return file_protobuf_todo_proto_rawDescGZIP(), []int{1}
}

func (x *TodoMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TodoMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TodoMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UserTodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserTodosRequest) Reset() {
	*x = UserTodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTodosRequest) ProtoMessage() {}

func (x *UserTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTodosRequest.ProtoReflect.Descriptor instead.
func (*UserTodosRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_todo_proto_rawDescGZIP(), []int{2}
}

func (x *UserTodosRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type TodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TodoId string `protobuf:"bytes,1,opt,name=todoId,proto3" json:"todoId,omitempty"`
}

func (x *TodoRequest) Reset() {
	*x = TodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoRequest) ProtoMessage() {}

func (x *TodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoRequest.ProtoReflect.Descriptor instead.
func (*TodoRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_todo_proto_rawDescGZIP(), []int{3}
}

func (x *TodoRequest) GetTodoId() string {
	if x != nil {
		return x.TodoId
	}
	return ""
}

type UserTodosWithItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserTodosWithItemsRequest) Reset() {
	*x = UserTodosWithItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTodosWithItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTodosWithItemsRequest) ProtoMessage() {}

func (x *UserTodosWithItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTodosWithItemsRequest.ProtoReflect.Descriptor instead.
func (*UserTodosWithItemsRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_todo_proto_rawDescGZIP(), []int{4}
}

func (x *UserTodosWithItemsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type TodoWithItemsMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Items       []*ItemMessage `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *TodoWithItemsMessage) Reset() {
	*x = TodoWithItemsMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoWithItemsMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoWithItemsMessage) ProtoMessage() {}

func (x *TodoWithItemsMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoWithItemsMessage.ProtoReflect.Descriptor instead.
func (*TodoWithItemsMessage) Descriptor() ([]byte, []int) {
	return file_protobuf_todo_proto_rawDescGZIP(), []int{5}
}

func (x *TodoWithItemsMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TodoWithItemsMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TodoWithItemsMessage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TodoWithItemsMessage) GetItems() []*ItemMessage {
	if x != nil {
		return x.Items
	}
	return nil
}

type UserTodosReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*TodoMessage `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *UserTodosReply) Reset() {
	*x = UserTodosReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTodosReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTodosReply) ProtoMessage() {}

func (x *UserTodosReply) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTodosReply.ProtoReflect.Descriptor instead.
func (*UserTodosReply) Descriptor() ([]byte, []int) {
	return file_protobuf_todo_proto_rawDescGZIP(), []int{6}
}

func (x *UserTodosReply) GetTodos() []*TodoMessage {
	if x != nil {
		return x.Todos
	}
	return nil
}

type UserTodosWithItemsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*TodoWithItemsMessage `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *UserTodosWithItemsReply) Reset() {
	*x = UserTodosWithItemsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_todo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTodosWithItemsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTodosWithItemsReply) ProtoMessage() {}

func (x *UserTodosWithItemsReply) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_todo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTodosWithItemsReply.ProtoReflect.Descriptor instead.
func (*UserTodosWithItemsReply) Descriptor() ([]byte, []int) {
	return file_protobuf_todo_proto_rawDescGZIP(), []int{7}
}

func (x *UserTodosWithItemsReply) GetTodos() []*TodoWithItemsMessage {
	if x != nil {
		return x.Todos
	}
	return nil
}

var File_protobuf_todo_proto protoreflect.FileDescriptor

var file_protobuf_todo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0b, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x53, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x10, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x64, 0x6f, 0x49, 0x64, 0x22, 0x33,
	0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x57, 0x69, 0x74, 0x68, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x14, 0x54, 0x6f, 0x64, 0x6f, 0x57, 0x69, 0x74, 0x68,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x34, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x64, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x22, 0x46, 0x0a, 0x17,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x57, 0x69, 0x74, 0x68, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x57, 0x69, 0x74,
	0x68, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x74,
	0x6f, 0x64, 0x6f, 0x73, 0x32, 0xbf, 0x01, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x34, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x11, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x64, 0x6f, 0x73, 0x57, 0x69, 0x74, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x11, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x57, 0x69, 0x74, 0x68, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x0c, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x54, 0x6f, 0x64, 0x6f, 0x57, 0x69, 0x74, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x68, 0x6e, 0x61, 0x6d, 0x62, 0x6d, 0x2f, 0x74, 0x6f,
	0x64, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_todo_proto_rawDescOnce sync.Once
	file_protobuf_todo_proto_rawDescData = file_protobuf_todo_proto_rawDesc
)

func file_protobuf_todo_proto_rawDescGZIP() []byte {
	file_protobuf_todo_proto_rawDescOnce.Do(func() {
		file_protobuf_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_todo_proto_rawDescData)
	})
	return file_protobuf_todo_proto_rawDescData
}

var file_protobuf_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protobuf_todo_proto_goTypes = []interface{}{
	(*ItemMessage)(nil),               // 0: ItemMessage
	(*TodoMessage)(nil),               // 1: TodoMessage
	(*UserTodosRequest)(nil),          // 2: UserTodosRequest
	(*TodoRequest)(nil),               // 3: TodoRequest
	(*UserTodosWithItemsRequest)(nil), // 4: UserTodosWithItemsRequest
	(*TodoWithItemsMessage)(nil),      // 5: TodoWithItemsMessage
	(*UserTodosReply)(nil),            // 6: UserTodosReply
	(*UserTodosWithItemsReply)(nil),   // 7: UserTodosWithItemsReply
}
var file_protobuf_todo_proto_depIdxs = []int32{
	0, // 0: TodoWithItemsMessage.items:type_name -> ItemMessage
	1, // 1: UserTodosReply.todos:type_name -> TodoMessage
	5, // 2: UserTodosWithItemsReply.todos:type_name -> TodoWithItemsMessage
	2, // 3: Todo.GetUserTodos:input_type -> UserTodosRequest
	2, // 4: Todo.GetUserTodosWithItems:input_type -> UserTodosRequest
	3, // 5: Todo.GetTodoWithItems:input_type -> TodoRequest
	6, // 6: Todo.GetUserTodos:output_type -> UserTodosReply
	7, // 7: Todo.GetUserTodosWithItems:output_type -> UserTodosWithItemsReply
	5, // 8: Todo.GetTodoWithItems:output_type -> TodoWithItemsMessage
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protobuf_todo_proto_init() }
func file_protobuf_todo_proto_init() {
	if File_protobuf_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemMessage); i {
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
		file_protobuf_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoMessage); i {
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
		file_protobuf_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTodosRequest); i {
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
		file_protobuf_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoRequest); i {
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
		file_protobuf_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTodosWithItemsRequest); i {
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
		file_protobuf_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoWithItemsMessage); i {
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
		file_protobuf_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTodosReply); i {
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
		file_protobuf_todo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTodosWithItemsReply); i {
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
			RawDescriptor: file_protobuf_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_todo_proto_goTypes,
		DependencyIndexes: file_protobuf_todo_proto_depIdxs,
		MessageInfos:      file_protobuf_todo_proto_msgTypes,
	}.Build()
	File_protobuf_todo_proto = out.File
	file_protobuf_todo_proto_rawDesc = nil
	file_protobuf_todo_proto_goTypes = nil
	file_protobuf_todo_proto_depIdxs = nil
}