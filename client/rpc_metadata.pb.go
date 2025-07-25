// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: client/rpc_metadata.proto

package clientpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type InstallerInitRequest_State int32

const (
	InstallerInitRequest_STATE_UNSPECIFIED InstallerInitRequest_State = 0
	InstallerInitRequest_STATE_INIT        InstallerInitRequest_State = 1
	InstallerInitRequest_STATE_DONE        InstallerInitRequest_State = 2
)

// Enum value maps for InstallerInitRequest_State.
var (
	InstallerInitRequest_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_INIT",
		2: "STATE_DONE",
	}
	InstallerInitRequest_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STATE_INIT":        1,
		"STATE_DONE":        2,
	}
)

func (x InstallerInitRequest_State) Enum() *InstallerInitRequest_State {
	p := new(InstallerInitRequest_State)
	*p = x
	return p
}

func (x InstallerInitRequest_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstallerInitRequest_State) Descriptor() protoreflect.EnumDescriptor {
	return file_client_rpc_metadata_proto_enumTypes[0].Descriptor()
}

func (InstallerInitRequest_State) Type() protoreflect.EnumType {
	return &file_client_rpc_metadata_proto_enumTypes[0]
}

func (x InstallerInitRequest_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstallerInitRequest_State.Descriptor instead.
func (InstallerInitRequest_State) EnumDescriptor() ([]byte, []int) {
	return file_client_rpc_metadata_proto_rawDescGZIP(), []int{0, 0}
}

type InstallerInitRequest struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Status        InstallerInitRequest_State `protobuf:"varint,1,opt,name=status,proto3,enum=client.InstallerInitRequest_State" json:"status,omitempty"`
	Metadata      string                     `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InstallerInitRequest) Reset() {
	*x = InstallerInitRequest{}
	mi := &file_client_rpc_metadata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstallerInitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallerInitRequest) ProtoMessage() {}

func (x *InstallerInitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_rpc_metadata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallerInitRequest.ProtoReflect.Descriptor instead.
func (*InstallerInitRequest) Descriptor() ([]byte, []int) {
	return file_client_rpc_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *InstallerInitRequest) GetStatus() InstallerInitRequest_State {
	if x != nil {
		return x.Status
	}
	return InstallerInitRequest_STATE_UNSPECIFIED
}

func (x *InstallerInitRequest) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type InstallerInitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InstallerInitResponse) Reset() {
	*x = InstallerInitResponse{}
	mi := &file_client_rpc_metadata_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstallerInitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallerInitResponse) ProtoMessage() {}

func (x *InstallerInitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_rpc_metadata_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallerInitResponse.ProtoReflect.Descriptor instead.
func (*InstallerInitResponse) Descriptor() ([]byte, []int) {
	return file_client_rpc_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *InstallerInitResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *InstallerInitResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type MetadataRepoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RepoUrl       string                 `protobuf:"bytes,1,opt,name=repo_url,json=repoUrl,proto3" json:"repo_url,omitempty"`
	Metadata      string                 `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetadataRepoRequest) Reset() {
	*x = MetadataRepoRequest{}
	mi := &file_client_rpc_metadata_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetadataRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataRepoRequest) ProtoMessage() {}

func (x *MetadataRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_rpc_metadata_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataRepoRequest.ProtoReflect.Descriptor instead.
func (*MetadataRepoRequest) Descriptor() ([]byte, []int) {
	return file_client_rpc_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *MetadataRepoRequest) GetRepoUrl() string {
	if x != nil {
		return x.RepoUrl
	}
	return ""
}

func (x *MetadataRepoRequest) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type MetadataRepoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetadataRepoResponse) Reset() {
	*x = MetadataRepoResponse{}
	mi := &file_client_rpc_metadata_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetadataRepoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataRepoResponse) ProtoMessage() {}

func (x *MetadataRepoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_rpc_metadata_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataRepoResponse.ProtoReflect.Descriptor instead.
func (*MetadataRepoResponse) Descriptor() ([]byte, []int) {
	return file_client_rpc_metadata_proto_rawDescGZIP(), []int{3}
}

func (x *MetadataRepoResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *MetadataRepoResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_client_rpc_metadata_proto protoreflect.FileDescriptor

const file_client_rpc_metadata_proto_rawDesc = "" +
	"\n" +
	"\x19client/rpc_metadata.proto\x12\x06client\x1a\x1fgoogle/protobuf/timestamp.proto\"\xae\x01\n" +
	"\x14InstallerInitRequest\x12:\n" +
	"\x06status\x18\x01 \x01(\x0e2\".client.InstallerInitRequest.StateR\x06status\x12\x1a\n" +
	"\bmetadata\x18\x02 \x01(\tR\bmetadata\">\n" +
	"\x05State\x12\x15\n" +
	"\x11STATE_UNSPECIFIED\x10\x00\x12\x0e\n" +
	"\n" +
	"STATE_INIT\x10\x01\x12\x0e\n" +
	"\n" +
	"STATE_DONE\x10\x02\"j\n" +
	"\x15InstallerInitResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x129\n" +
	"\n" +
	"created_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\"L\n" +
	"\x13MetadataRepoRequest\x12\x19\n" +
	"\brepo_url\x18\x01 \x01(\tR\arepoUrl\x12\x1a\n" +
	"\bmetadata\x18\x02 \x01(\tR\bmetadata\"i\n" +
	"\x14MetadataRepoResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x129\n" +
	"\n" +
	"created_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt2\xad\x01\n" +
	"\x0fMetadataService\x12M\n" +
	"\fInitMetadata\x12\x1c.client.InstallerInitRequest\x1a\x1d.client.InstallerInitResponse\"\x00\x12K\n" +
	"\fRepoMetadata\x12\x1b.client.MetadataRepoRequest\x1a\x1c.client.MetadataRepoResponse\"\x00B\x19Z\x17axilock-runner/clientpbb\x06proto3"

var (
	file_client_rpc_metadata_proto_rawDescOnce sync.Once
	file_client_rpc_metadata_proto_rawDescData []byte
)

func file_client_rpc_metadata_proto_rawDescGZIP() []byte {
	file_client_rpc_metadata_proto_rawDescOnce.Do(func() {
		file_client_rpc_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_client_rpc_metadata_proto_rawDesc), len(file_client_rpc_metadata_proto_rawDesc)))
	})
	return file_client_rpc_metadata_proto_rawDescData
}

var file_client_rpc_metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_client_rpc_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_client_rpc_metadata_proto_goTypes = []any{
	(InstallerInitRequest_State)(0), // 0: client.InstallerInitRequest.State
	(*InstallerInitRequest)(nil),    // 1: client.InstallerInitRequest
	(*InstallerInitResponse)(nil),   // 2: client.InstallerInitResponse
	(*MetadataRepoRequest)(nil),     // 3: client.MetadataRepoRequest
	(*MetadataRepoResponse)(nil),    // 4: client.MetadataRepoResponse
	(*timestamppb.Timestamp)(nil),   // 5: google.protobuf.Timestamp
}
var file_client_rpc_metadata_proto_depIdxs = []int32{
	0, // 0: client.InstallerInitRequest.status:type_name -> client.InstallerInitRequest.State
	5, // 1: client.InstallerInitResponse.created_at:type_name -> google.protobuf.Timestamp
	5, // 2: client.MetadataRepoResponse.created_at:type_name -> google.protobuf.Timestamp
	1, // 3: client.MetadataService.InitMetadata:input_type -> client.InstallerInitRequest
	3, // 4: client.MetadataService.RepoMetadata:input_type -> client.MetadataRepoRequest
	2, // 5: client.MetadataService.InitMetadata:output_type -> client.InstallerInitResponse
	4, // 6: client.MetadataService.RepoMetadata:output_type -> client.MetadataRepoResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_client_rpc_metadata_proto_init() }
func file_client_rpc_metadata_proto_init() {
	if File_client_rpc_metadata_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_client_rpc_metadata_proto_rawDesc), len(file_client_rpc_metadata_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_rpc_metadata_proto_goTypes,
		DependencyIndexes: file_client_rpc_metadata_proto_depIdxs,
		EnumInfos:         file_client_rpc_metadata_proto_enumTypes,
		MessageInfos:      file_client_rpc_metadata_proto_msgTypes,
	}.Build()
	File_client_rpc_metadata_proto = out.File
	file_client_rpc_metadata_proto_goTypes = nil
	file_client_rpc_metadata_proto_depIdxs = nil
}
