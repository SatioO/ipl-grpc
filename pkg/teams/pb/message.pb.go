// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: pkg/teams/proto/message.proto

package teams_pb

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

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Abbr  string `protobuf:"bytes,3,opt,name=abbr,proto3" json:"abbr,omitempty"`
	Color string `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	Logo  string `protobuf:"bytes,5,opt,name=logo,proto3" json:"logo,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_teams_proto_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_teams_proto_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_pkg_teams_proto_message_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Team) GetAbbr() string {
	if x != nil {
		return x.Abbr
	}
	return ""
}

func (x *Team) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Team) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

type GetTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTeamRequest) Reset() {
	*x = GetTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_teams_proto_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeamRequest) ProtoMessage() {}

func (x *GetTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_teams_proto_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeamRequest.ProtoReflect.Descriptor instead.
func (*GetTeamRequest) Descriptor() ([]byte, []int) {
	return file_pkg_teams_proto_message_proto_rawDescGZIP(), []int{1}
}

type GetTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Team `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *GetTeamResponse) Reset() {
	*x = GetTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_teams_proto_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeamResponse) ProtoMessage() {}

func (x *GetTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_teams_proto_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeamResponse.ProtoReflect.Descriptor instead.
func (*GetTeamResponse) Descriptor() ([]byte, []int) {
	return file_pkg_teams_proto_message_proto_rawDescGZIP(), []int{2}
}

func (x *GetTeamResponse) GetItems() []*Team {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateTeamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTeamsRequest) Reset() {
	*x = CreateTeamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_teams_proto_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamsRequest) ProtoMessage() {}

func (x *CreateTeamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_teams_proto_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamsRequest.ProtoReflect.Descriptor instead.
func (*CreateTeamsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_teams_proto_message_proto_rawDescGZIP(), []int{3}
}

type CreateTeamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateTeamsResponse) Reset() {
	*x = CreateTeamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_teams_proto_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTeamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamsResponse) ProtoMessage() {}

func (x *CreateTeamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_teams_proto_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamsResponse.ProtoReflect.Descriptor instead.
func (*CreateTeamsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_teams_proto_message_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTeamsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_teams_proto_message_proto protoreflect.FileDescriptor

var file_pkg_teams_proto_message_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x5f, 0x70, 0x62, 0x22, 0x68, 0x0a, 0x04, 0x54, 0x65, 0x61,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x62, 0x62, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x62, 0x62, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x6f, 0x67, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x5f,
	0x70, 0x62, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x14,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x74, 0x65, 0x61, 0x6d, 0x73,
	0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_teams_proto_message_proto_rawDescOnce sync.Once
	file_pkg_teams_proto_message_proto_rawDescData = file_pkg_teams_proto_message_proto_rawDesc
)

func file_pkg_teams_proto_message_proto_rawDescGZIP() []byte {
	file_pkg_teams_proto_message_proto_rawDescOnce.Do(func() {
		file_pkg_teams_proto_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_teams_proto_message_proto_rawDescData)
	})
	return file_pkg_teams_proto_message_proto_rawDescData
}

var file_pkg_teams_proto_message_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_teams_proto_message_proto_goTypes = []interface{}{
	(*Team)(nil),                // 0: teams_pb.Team
	(*GetTeamRequest)(nil),      // 1: teams_pb.GetTeamRequest
	(*GetTeamResponse)(nil),     // 2: teams_pb.GetTeamResponse
	(*CreateTeamsRequest)(nil),  // 3: teams_pb.CreateTeamsRequest
	(*CreateTeamsResponse)(nil), // 4: teams_pb.CreateTeamsResponse
}
var file_pkg_teams_proto_message_proto_depIdxs = []int32{
	0, // 0: teams_pb.GetTeamResponse.Items:type_name -> teams_pb.Team
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_teams_proto_message_proto_init() }
func file_pkg_teams_proto_message_proto_init() {
	if File_pkg_teams_proto_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_teams_proto_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Team); i {
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
		file_pkg_teams_proto_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeamRequest); i {
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
		file_pkg_teams_proto_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeamResponse); i {
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
		file_pkg_teams_proto_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamsRequest); i {
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
		file_pkg_teams_proto_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTeamsResponse); i {
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
			RawDescriptor: file_pkg_teams_proto_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_teams_proto_message_proto_goTypes,
		DependencyIndexes: file_pkg_teams_proto_message_proto_depIdxs,
		MessageInfos:      file_pkg_teams_proto_message_proto_msgTypes,
	}.Build()
	File_pkg_teams_proto_message_proto = out.File
	file_pkg_teams_proto_message_proto_rawDesc = nil
	file_pkg_teams_proto_message_proto_goTypes = nil
	file_pkg_teams_proto_message_proto_depIdxs = nil
}
