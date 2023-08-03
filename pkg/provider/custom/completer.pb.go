
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: completer.proto

package custom

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

type Role int32

const (
	Role_ROLE_UNSPECIFIED Role = 0
	Role_ROLE_SYSTEM      Role = 1
	Role_ROLE_USER        Role = 2
	Role_ROLE_ASSISTANT   Role = 3
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "ROLE_SYSTEM",
		2: "ROLE_USER",
		3: "ROLE_ASSISTANT",
	}
	Role_value = map[string]int32{
		"ROLE_UNSPECIFIED": 0,
		"ROLE_SYSTEM":      1,
		"ROLE_USER":        2,
		"ROLE_ASSISTANT":   3,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_completer_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_completer_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_completer_proto_rawDescGZIP(), []int{0}
}

type Completion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Completion) Reset() {
	*x = Completion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_completer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Completion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Completion) ProtoMessage() {}

func (x *Completion) ProtoReflect() protoreflect.Message {
	mi := &file_completer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Completion.ProtoReflect.Descriptor instead.
func (*Completion) Descriptor() ([]byte, []int) {
	return file_completer_proto_rawDescGZIP(), []int{0}
}

func (x *Completion) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Completion) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role    Role   `protobuf:"varint,1,opt,name=role,proto3,enum=completer.Role" json:"role,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_completer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_completer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_completer_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CompletionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model       string     `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Messages    []*Message `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	Stop        []string   `protobuf:"bytes,5,rep,name=stop,proto3" json:"stop,omitempty"`
	MaxTokens   *int32     `protobuf:"varint,4,opt,name=max_tokens,json=maxTokens,proto3,oneof" json:"max_tokens,omitempty"`
	Temperature *float32   `protobuf:"fixed32,3,opt,name=temperature,proto3,oneof" json:"temperature,omitempty"`
}

func (x *CompletionRequest) Reset() {
	*x = CompletionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_completer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionRequest) ProtoMessage() {}

func (x *CompletionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_completer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionRequest.ProtoReflect.Descriptor instead.
func (*CompletionRequest) Descriptor() ([]byte, []int) {
	return file_completer_proto_rawDescGZIP(), []int{2}
}

func (x *CompletionRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CompletionRequest) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *CompletionRequest) GetStop() []string {
	if x != nil {
		return x.Stop
	}
	return nil
}

func (x *CompletionRequest) GetMaxTokens() int32 {
	if x != nil && x.MaxTokens != nil {
		return *x.MaxTokens
	}
	return 0
}

func (x *CompletionRequest) GetTemperature() float32 {
	if x != nil && x.Temperature != nil {
		return *x.Temperature
	}
	return 0
}

var File_completer_proto protoreflect.FileDescriptor

var file_completer_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x22, 0x4a, 0x0a, 0x0a,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x48, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0xd7, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x2e,
	0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x74,
	0x6f, 0x70, 0x12, 0x22, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x48, 0x01, 0x52, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2a, 0x50, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f,
	0x4c, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x52,
	0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f,
	0x4c, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x49, 0x53, 0x54, 0x41, 0x4e, 0x54, 0x10, 0x03, 0x32, 0x50,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x08, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x64, 0x72, 0x69, 0x61, 0x6e, 0x6c, 0x69, 0x65, 0x63, 0x68, 0x74, 0x69, 0x2f, 0x6c, 0x6c, 0x61,
	0x6d, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_completer_proto_rawDescOnce sync.Once
	file_completer_proto_rawDescData = file_completer_proto_rawDesc
)

func file_completer_proto_rawDescGZIP() []byte {
	file_completer_proto_rawDescOnce.Do(func() {
		file_completer_proto_rawDescData = protoimpl.X.CompressGZIP(file_completer_proto_rawDescData)
	})
	return file_completer_proto_rawDescData
}

var file_completer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_completer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_completer_proto_goTypes = []interface{}{
	(Role)(0),                 // 0: completer.Role
	(*Completion)(nil),        // 1: completer.Completion
	(*Message)(nil),           // 2: completer.Message
	(*CompletionRequest)(nil), // 3: completer.CompletionRequest
}
var file_completer_proto_depIdxs = []int32{
	2, // 0: completer.Completion.message:type_name -> completer.Message
	0, // 1: completer.Message.role:type_name -> completer.Role
	2, // 2: completer.CompletionRequest.messages:type_name -> completer.Message
	3, // 3: completer.completer.Complete:input_type -> completer.CompletionRequest
	1, // 4: completer.completer.Complete:output_type -> completer.Completion
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_completer_proto_init() }
func file_completer_proto_init() {
	if File_completer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_completer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Completion); i {
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
		file_completer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_completer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletionRequest); i {
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
	file_completer_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_completer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_completer_proto_goTypes,
		DependencyIndexes: file_completer_proto_depIdxs,
		EnumInfos:         file_completer_proto_enumTypes,
		MessageInfos:      file_completer_proto_msgTypes,
	}.Build()
	File_completer_proto = out.File
	file_completer_proto_rawDesc = nil
	file_completer_proto_goTypes = nil
	file_completer_proto_depIdxs = nil
}