// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/teams.proto

package team_pb

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

type Member struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint64                 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Admin         bool                   `protobuf:"varint,3,opt,name=admin,proto3" json:"admin,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Member) Reset() {
	*x = Member{}
	mi := &file_proto_teams_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_proto_teams_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Member) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Member) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

type Team struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Members       []*Member              `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Team) Reset() {
	*x = Team{}
	mi := &file_proto_teams_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_proto_msgTypes[1]
	if x != nil {
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
	return file_proto_teams_proto_rawDescGZIP(), []int{1}
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

func (x *Team) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Team) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

type TeamList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Teams         []*Team                `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TeamList) Reset() {
	*x = TeamList{}
	mi := &file_proto_teams_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TeamList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamList) ProtoMessage() {}

func (x *TeamList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamList.ProtoReflect.Descriptor instead.
func (*TeamList) Descriptor() ([]byte, []int) {
	return file_proto_teams_proto_rawDescGZIP(), []int{2}
}

func (x *TeamList) GetTeams() []*Team {
	if x != nil {
		return x.Teams
	}
	return nil
}

type CreateTeamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Members       []*Member              `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTeamRequest) Reset() {
	*x = CreateTeamRequest{}
	mi := &file_proto_teams_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeamRequest) ProtoMessage() {}

func (x *CreateTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeamRequest.ProtoReflect.Descriptor instead.
func (*CreateTeamRequest) Descriptor() ([]byte, []int) {
	return file_proto_teams_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTeamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTeamRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTeamRequest) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

type GetTeamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TeamId        string                 `protobuf:"bytes,1,opt,name=teamId,proto3" json:"teamId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTeamRequest) Reset() {
	*x = GetTeamRequest{}
	mi := &file_proto_teams_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeamRequest) ProtoMessage() {}

func (x *GetTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_proto_msgTypes[4]
	if x != nil {
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
	return file_proto_teams_proto_rawDescGZIP(), []int{4}
}

func (x *GetTeamRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

type DeleteTeamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TeamId        string                 `protobuf:"bytes,1,opt,name=teamId,proto3" json:"teamId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTeamRequest) Reset() {
	*x = DeleteTeamRequest{}
	mi := &file_proto_teams_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeamRequest) ProtoMessage() {}

func (x *DeleteTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeamRequest.ProtoReflect.Descriptor instead.
func (*DeleteTeamRequest) Descriptor() ([]byte, []int) {
	return file_proto_teams_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTeamRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

type UpdateTeamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TeamId        string                 `protobuf:"bytes,1,opt,name=teamId,proto3" json:"teamId,omitempty"`
	Update        *Team                  `protobuf:"bytes,2,opt,name=update,proto3" json:"update,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTeamRequest) Reset() {
	*x = UpdateTeamRequest{}
	mi := &file_proto_teams_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeamRequest) ProtoMessage() {}

func (x *UpdateTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeamRequest.ProtoReflect.Descriptor instead.
func (*UpdateTeamRequest) Descriptor() ([]byte, []int) {
	return file_proto_teams_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTeamRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *UpdateTeamRequest) GetUpdate() *Team {
	if x != nil {
		return x.Update
	}
	return nil
}

type TeamResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TeamResponse) Reset() {
	*x = TeamResponse{}
	mi := &file_proto_teams_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamResponse) ProtoMessage() {}

func (x *TeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamResponse.ProtoReflect.Descriptor instead.
func (*TeamResponse) Descriptor() ([]byte, []int) {
	return file_proto_teams_proto_rawDescGZIP(), []int{7}
}

func (x *TeamResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *TeamResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TeamResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_teams_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_teams_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_teams_proto_rawDescGZIP(), []int{8}
}

var File_proto_teams_proto protoreflect.FileDescriptor

const file_proto_teams_proto_rawDesc = "" +
	"\n" +
	"\x11proto/teams.proto\x12\x04team\"J\n" +
	"\x06Member\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\x04R\x06userId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05admin\x18\x03 \x01(\bR\x05admin\"t\n" +
	"\x04Team\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12&\n" +
	"\amembers\x18\x04 \x03(\v2\f.team.MemberR\amembers\",\n" +
	"\bTeamList\x12 \n" +
	"\x05teams\x18\x01 \x03(\v2\n" +
	".team.TeamR\x05teams\"q\n" +
	"\x11CreateTeamRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12&\n" +
	"\amembers\x18\x03 \x03(\v2\f.team.MemberR\amembers\"(\n" +
	"\x0eGetTeamRequest\x12\x16\n" +
	"\x06teamId\x18\x01 \x01(\tR\x06teamId\"+\n" +
	"\x11DeleteTeamRequest\x12\x16\n" +
	"\x06teamId\x18\x01 \x01(\tR\x06teamId\"O\n" +
	"\x11UpdateTeamRequest\x12\x16\n" +
	"\x06teamId\x18\x01 \x01(\tR\x06teamId\x12\"\n" +
	"\x06update\x18\x02 \x01(\v2\n" +
	".team.TeamR\x06update\"R\n" +
	"\fTeamResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\tR\x02id\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\"\a\n" +
	"\x05Empty2\x94\x02\n" +
	"\vTeamService\x129\n" +
	"\n" +
	"CreateTeam\x12\x17.team.CreateTeamRequest\x1a\x12.team.TeamResponse\x12+\n" +
	"\aGetTeam\x12\x14.team.GetTeamRequest\x1a\n" +
	".team.Team\x12'\n" +
	"\bGetTeams\x12\v.team.Empty\x1a\x0e.team.TeamList\x129\n" +
	"\n" +
	"DeleteTeam\x12\x17.team.DeleteTeamRequest\x1a\x12.team.TeamResponse\x129\n" +
	"\n" +
	"UpdateTeam\x12\x17.team.UpdateTeamRequest\x1a\x12.team.TeamResponseB\fZ\n" +
	"pb/team_pbb\x06proto3"

var (
	file_proto_teams_proto_rawDescOnce sync.Once
	file_proto_teams_proto_rawDescData []byte
)

func file_proto_teams_proto_rawDescGZIP() []byte {
	file_proto_teams_proto_rawDescOnce.Do(func() {
		file_proto_teams_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_teams_proto_rawDesc), len(file_proto_teams_proto_rawDesc)))
	})
	return file_proto_teams_proto_rawDescData
}

var file_proto_teams_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_teams_proto_goTypes = []any{
	(*Member)(nil),            // 0: team.Member
	(*Team)(nil),              // 1: team.Team
	(*TeamList)(nil),          // 2: team.TeamList
	(*CreateTeamRequest)(nil), // 3: team.CreateTeamRequest
	(*GetTeamRequest)(nil),    // 4: team.GetTeamRequest
	(*DeleteTeamRequest)(nil), // 5: team.DeleteTeamRequest
	(*UpdateTeamRequest)(nil), // 6: team.UpdateTeamRequest
	(*TeamResponse)(nil),      // 7: team.TeamResponse
	(*Empty)(nil),             // 8: team.Empty
}
var file_proto_teams_proto_depIdxs = []int32{
	0, // 0: team.Team.members:type_name -> team.Member
	1, // 1: team.TeamList.teams:type_name -> team.Team
	0, // 2: team.CreateTeamRequest.members:type_name -> team.Member
	1, // 3: team.UpdateTeamRequest.update:type_name -> team.Team
	3, // 4: team.TeamService.CreateTeam:input_type -> team.CreateTeamRequest
	4, // 5: team.TeamService.GetTeam:input_type -> team.GetTeamRequest
	8, // 6: team.TeamService.GetTeams:input_type -> team.Empty
	5, // 7: team.TeamService.DeleteTeam:input_type -> team.DeleteTeamRequest
	6, // 8: team.TeamService.UpdateTeam:input_type -> team.UpdateTeamRequest
	7, // 9: team.TeamService.CreateTeam:output_type -> team.TeamResponse
	1, // 10: team.TeamService.GetTeam:output_type -> team.Team
	2, // 11: team.TeamService.GetTeams:output_type -> team.TeamList
	7, // 12: team.TeamService.DeleteTeam:output_type -> team.TeamResponse
	7, // 13: team.TeamService.UpdateTeam:output_type -> team.TeamResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_teams_proto_init() }
func file_proto_teams_proto_init() {
	if File_proto_teams_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_teams_proto_rawDesc), len(file_proto_teams_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_teams_proto_goTypes,
		DependencyIndexes: file_proto_teams_proto_depIdxs,
		MessageInfos:      file_proto_teams_proto_msgTypes,
	}.Build()
	File_proto_teams_proto = out.File
	file_proto_teams_proto_goTypes = nil
	file_proto_teams_proto_depIdxs = nil
}
