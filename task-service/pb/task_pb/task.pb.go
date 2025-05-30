// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: task.proto

package task_pb

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

// --- Основная сущность ---
type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        uint64                 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	TeamId        string                 `protobuf:"bytes,3,opt,name=teamId,proto3" json:"teamId,omitempty"`
	Title         string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Status        string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`     // pending, in-progress, done
	Priority      string                 `protobuf:"bytes,7,opt,name=priority,proto3" json:"priority,omitempty"` // low, medium, high
	Deadline      *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Tags          []string               `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	Folder        string                 `protobuf:"bytes,10,opt,name=folder,proto3" json:"folder,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Task) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Task) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

func (x *Task) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

func (x *Task) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Task) GetFolder() string {
	if x != nil {
		return x.Folder
	}
	return ""
}

func (x *Task) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Task) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// --- Список задач ---
type TaskList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*Task                `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskList) Reset() {
	*x = TaskList{}
	mi := &file_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskList) ProtoMessage() {}

func (x *TaskList) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskList.ProtoReflect.Descriptor instead.
func (*TaskList) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskList) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

// --- Запросы ---
type CreateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint64                 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	TeamId        string                 `protobuf:"bytes,2,opt,name=teamId,proto3" json:"teamId,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Priority      string                 `protobuf:"bytes,6,opt,name=priority,proto3" json:"priority,omitempty"`
	Deadline      *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Tags          []string               `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	Folder        string                 `protobuf:"bytes,9,opt,name=folder,proto3" json:"folder,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	mi := &file_task_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTaskRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateTaskRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTaskRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateTaskRequest) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

func (x *CreateTaskRequest) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

func (x *CreateTaskRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateTaskRequest) GetFolder() string {
	if x != nil {
		return x.Folder
	}
	return ""
}

type GetTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        uint64                 `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	mi := &file_task_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{3}
}

func (x *GetTaskRequest) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        uint64                 `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	mi := &file_task_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTaskRequest) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        uint64                 `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Update        *Task                  `protobuf:"bytes,2,opt,name=update,proto3" json:"update,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	mi := &file_task_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTaskRequest) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *UpdateTaskRequest) GetUpdate() *Task {
	if x != nil {
		return x.Update
	}
	return nil
}

// --- Ответы ---
type TaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Id            uint64                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Task          *Task                  `protobuf:"bytes,4,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	mi := &file_task_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{6}
}

func (x *TaskResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *TaskResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

// --- Пустое сообщение ---
type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_task_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[7]
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
	return file_task_proto_rawDescGZIP(), []int{7}
}

var File_task_proto protoreflect.FileDescriptor

const file_task_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"task.proto\x12\x04task\x1a\x1fgoogle/protobuf/timestamp.proto\"\x8a\x03\n" +
	"\x04Task\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x16\n" +
	"\x06userId\x18\x02 \x01(\x04R\x06userId\x12\x16\n" +
	"\x06teamId\x18\x03 \x01(\tR\x06teamId\x12\x14\n" +
	"\x05title\x18\x04 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12\x16\n" +
	"\x06status\x18\x06 \x01(\tR\x06status\x12\x1a\n" +
	"\bpriority\x18\a \x01(\tR\bpriority\x126\n" +
	"\bdeadline\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\bdeadline\x12\x12\n" +
	"\x04tags\x18\t \x03(\tR\x04tags\x12\x16\n" +
	"\x06folder\x18\n" +
	" \x01(\tR\x06folder\x128\n" +
	"\tcreatedAt\x18\v \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x128\n" +
	"\tupdatedAt\x18\f \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\",\n" +
	"\bTaskList\x12 \n" +
	"\x05tasks\x18\x01 \x03(\v2\n" +
	".task.TaskR\x05tasks\"\x93\x02\n" +
	"\x11CreateTaskRequest\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\x04R\x06userId\x12\x16\n" +
	"\x06teamId\x18\x02 \x01(\tR\x06teamId\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12\x16\n" +
	"\x06status\x18\x05 \x01(\tR\x06status\x12\x1a\n" +
	"\bpriority\x18\x06 \x01(\tR\bpriority\x126\n" +
	"\bdeadline\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\bdeadline\x12\x12\n" +
	"\x04tags\x18\b \x03(\tR\x04tags\x12\x16\n" +
	"\x06folder\x18\t \x01(\tR\x06folder\"(\n" +
	"\x0eGetTaskRequest\x12\x16\n" +
	"\x06taskId\x18\x01 \x01(\x04R\x06taskId\"+\n" +
	"\x11DeleteTaskRequest\x12\x16\n" +
	"\x06taskId\x18\x01 \x01(\x04R\x06taskId\"O\n" +
	"\x11UpdateTaskRequest\x12\x16\n" +
	"\x06taskId\x18\x01 \x01(\x04R\x06taskId\x12\"\n" +
	"\x06update\x18\x02 \x01(\v2\n" +
	".task.TaskR\x06update\"r\n" +
	"\fTaskResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\x04R\x02id\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\x12\x1e\n" +
	"\x04task\x18\x04 \x01(\v2\n" +
	".task.TaskR\x04task\"\a\n" +
	"\x05Empty2\x94\x02\n" +
	"\vTaskService\x129\n" +
	"\n" +
	"CreateTask\x12\x17.task.CreateTaskRequest\x1a\x12.task.TaskResponse\x12+\n" +
	"\aGetTask\x12\x14.task.GetTaskRequest\x1a\n" +
	".task.Task\x12'\n" +
	"\bGetTasks\x12\v.task.Empty\x1a\x0e.task.TaskList\x129\n" +
	"\n" +
	"DeleteTask\x12\x17.task.DeleteTaskRequest\x1a\x12.task.TaskResponse\x129\n" +
	"\n" +
	"UpdateTask\x12\x17.task.UpdateTaskRequest\x1a\x12.task.TaskResponseB\x11Z\x0ftask/pb/task_pbb\x06proto3"

var (
	file_task_proto_rawDescOnce sync.Once
	file_task_proto_rawDescData []byte
)

func file_task_proto_rawDescGZIP() []byte {
	file_task_proto_rawDescOnce.Do(func() {
		file_task_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_task_proto_rawDesc), len(file_task_proto_rawDesc)))
	})
	return file_task_proto_rawDescData
}

var file_task_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_task_proto_goTypes = []any{
	(*Task)(nil),                  // 0: task.Task
	(*TaskList)(nil),              // 1: task.TaskList
	(*CreateTaskRequest)(nil),     // 2: task.CreateTaskRequest
	(*GetTaskRequest)(nil),        // 3: task.GetTaskRequest
	(*DeleteTaskRequest)(nil),     // 4: task.DeleteTaskRequest
	(*UpdateTaskRequest)(nil),     // 5: task.UpdateTaskRequest
	(*TaskResponse)(nil),          // 6: task.TaskResponse
	(*Empty)(nil),                 // 7: task.Empty
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_task_proto_depIdxs = []int32{
	8,  // 0: task.Task.deadline:type_name -> google.protobuf.Timestamp
	8,  // 1: task.Task.createdAt:type_name -> google.protobuf.Timestamp
	8,  // 2: task.Task.updatedAt:type_name -> google.protobuf.Timestamp
	0,  // 3: task.TaskList.tasks:type_name -> task.Task
	8,  // 4: task.CreateTaskRequest.deadline:type_name -> google.protobuf.Timestamp
	0,  // 5: task.UpdateTaskRequest.update:type_name -> task.Task
	0,  // 6: task.TaskResponse.task:type_name -> task.Task
	2,  // 7: task.TaskService.CreateTask:input_type -> task.CreateTaskRequest
	3,  // 8: task.TaskService.GetTask:input_type -> task.GetTaskRequest
	7,  // 9: task.TaskService.GetTasks:input_type -> task.Empty
	4,  // 10: task.TaskService.DeleteTask:input_type -> task.DeleteTaskRequest
	5,  // 11: task.TaskService.UpdateTask:input_type -> task.UpdateTaskRequest
	6,  // 12: task.TaskService.CreateTask:output_type -> task.TaskResponse
	0,  // 13: task.TaskService.GetTask:output_type -> task.Task
	1,  // 14: task.TaskService.GetTasks:output_type -> task.TaskList
	6,  // 15: task.TaskService.DeleteTask:output_type -> task.TaskResponse
	6,  // 16: task.TaskService.UpdateTask:output_type -> task.TaskResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_task_proto_init() }
func file_task_proto_init() {
	if File_task_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_task_proto_rawDesc), len(file_task_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_proto_goTypes,
		DependencyIndexes: file_task_proto_depIdxs,
		MessageInfos:      file_task_proto_msgTypes,
	}.Build()
	File_task_proto = out.File
	file_task_proto_goTypes = nil
	file_task_proto_depIdxs = nil
}
