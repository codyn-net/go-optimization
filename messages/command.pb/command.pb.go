package command

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"
import "ponyo.epfl.ch/go/get/optimization-go/optimization/messages/task.pb"

var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CommandType int32

const (
	CommandType_List		CommandType	= 0
	CommandType_Info		CommandType	= 1
	CommandType_Kill		CommandType	= 2
	CommandType_SetPriority		CommandType	= 3
	CommandType_Authenticate	CommandType	= 4
	CommandType_Progress		CommandType	= 5
	CommandType_Idle		CommandType	= 6
	CommandType_NumCommands		CommandType	= 7
)

var CommandType_name = map[int32]string{
	0:	"List",
	1:	"Info",
	2:	"Kill",
	3:	"SetPriority",
	4:	"Authenticate",
	5:	"Progress",
	6:	"Idle",
	7:	"NumCommands",
}
var CommandType_value = map[string]int32{
	"List":		0,
	"Info":		1,
	"Kill":		2,
	"SetPriority":	3,
	"Authenticate":	4,
	"Progress":	5,
	"Idle":		6,
	"NumCommands":	7,
}

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}
func (x CommandType) String() string {
	return proto.EnumName(CommandType_name, int32(x))
}
func (x CommandType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *CommandType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommandType_value, data, "CommandType")
	if err != nil {
		return err
	}
	*x = CommandType(value)
	return nil
}

type KillCommand struct {
	Id			*uint32	`protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *KillCommand) Reset()		{ *this = KillCommand{} }
func (this *KillCommand) String() string	{ return proto.CompactTextString(this) }
func (*KillCommand) ProtoMessage()		{}

func (this *KillCommand) GetId() uint32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

type SetPriorityCommand struct {
	Id			*uint32		`protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Priority		*float64	`protobuf:"fixed64,2,req,name=priority" json:"priority,omitempty"`
	XXX_unrecognized	[]byte		`json:"-"`
}

func (this *SetPriorityCommand) Reset()		{ *this = SetPriorityCommand{} }
func (this *SetPriorityCommand) String() string	{ return proto.CompactTextString(this) }
func (*SetPriorityCommand) ProtoMessage()	{}

func (this *SetPriorityCommand) GetId() uint32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *SetPriorityCommand) GetPriority() float64 {
	if this != nil && this.Priority != nil {
		return *this.Priority
	}
	return 0
}

type AuthenticateCommand struct {
	Token			*string	`protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *AuthenticateCommand) Reset()		{ *this = AuthenticateCommand{} }
func (this *AuthenticateCommand) String() string	{ return proto.CompactTextString(this) }
func (*AuthenticateCommand) ProtoMessage()		{}

func (this *AuthenticateCommand) GetToken() string {
	if this != nil && this.Token != nil {
		return *this.Token
	}
	return ""
}

type ListCommand struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *ListCommand) Reset()		{ *this = ListCommand{} }
func (this *ListCommand) String() string	{ return proto.CompactTextString(this) }
func (*ListCommand) ProtoMessage()		{}

type InfoCommand struct {
	Id			*uint32	`protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *InfoCommand) Reset()		{ *this = InfoCommand{} }
func (this *InfoCommand) String() string	{ return proto.CompactTextString(this) }
func (*InfoCommand) ProtoMessage()		{}

func (this *InfoCommand) GetId() uint32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

type ProgressCommand struct {
	Id			*uint32	`protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *ProgressCommand) Reset()		{ *this = ProgressCommand{} }
func (this *ProgressCommand) String() string	{ return proto.CompactTextString(this) }
func (*ProgressCommand) ProtoMessage()		{}

func (this *ProgressCommand) GetId() uint32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

type IdleCommand struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *IdleCommand) Reset()		{ *this = IdleCommand{} }
func (this *IdleCommand) String() string	{ return proto.CompactTextString(this) }
func (*IdleCommand) ProtoMessage()		{}

type Command struct {
	Type			*CommandType		`protobuf:"varint,1,req,name=type,enum=optimization.messages.command.CommandType" json:"type,omitempty"`
	List			*ListCommand		`protobuf:"bytes,2,opt,name=list" json:"list,omitempty"`
	Info			*InfoCommand		`protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
	Kill			*KillCommand		`protobuf:"bytes,4,opt,name=kill" json:"kill,omitempty"`
	Setpriority		*SetPriorityCommand	`protobuf:"bytes,5,opt,name=setpriority" json:"setpriority,omitempty"`
	Authenticate		*AuthenticateCommand	`protobuf:"bytes,6,opt,name=authenticate" json:"authenticate,omitempty"`
	Progress		*ProgressCommand	`protobuf:"bytes,7,opt,name=progress" json:"progress,omitempty"`
	Idle			*IdleCommand		`protobuf:"bytes,8,opt,name=idle" json:"idle,omitempty"`
	XXX_unrecognized	[]byte			`json:"-"`
}

func (this *Command) Reset()		{ *this = Command{} }
func (this *Command) String() string	{ return proto.CompactTextString(this) }
func (*Command) ProtoMessage()		{}

func (this *Command) GetType() CommandType {
	if this != nil && this.Type != nil {
		return *this.Type
	}
	return 0
}

func (this *Command) GetList() *ListCommand {
	if this != nil {
		return this.List
	}
	return nil
}

func (this *Command) GetInfo() *InfoCommand {
	if this != nil {
		return this.Info
	}
	return nil
}

func (this *Command) GetKill() *KillCommand {
	if this != nil {
		return this.Kill
	}
	return nil
}

func (this *Command) GetSetpriority() *SetPriorityCommand {
	if this != nil {
		return this.Setpriority
	}
	return nil
}

func (this *Command) GetAuthenticate() *AuthenticateCommand {
	if this != nil {
		return this.Authenticate
	}
	return nil
}

func (this *Command) GetProgress() *ProgressCommand {
	if this != nil {
		return this.Progress
	}
	return nil
}

func (this *Command) GetIdle() *IdleCommand {
	if this != nil {
		return this.Idle
	}
	return nil
}

type Job struct {
	Id			*uint32		`protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name			*string		`protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	User			*string		`protobuf:"bytes,3,req,name=user" json:"user,omitempty"`
	Priority		*float64	`protobuf:"fixed64,4,req,name=priority" json:"priority,omitempty"`
	Started			*uint64		`protobuf:"varint,5,req,name=started" json:"started,omitempty"`
	Lastupdate		*uint64		`protobuf:"varint,6,req,name=lastupdate" json:"lastupdate,omitempty"`
	Progress		*float64	`protobuf:"fixed64,7,req,name=progress" json:"progress,omitempty"`
	Taskssuccess		*uint32		`protobuf:"varint,8,req,name=taskssuccess" json:"taskssuccess,omitempty"`
	Tasksfailed		*uint32		`protobuf:"varint,9,req,name=tasksfailed" json:"tasksfailed,omitempty"`
	Runtime			*float64	`protobuf:"fixed64,10,req,name=runtime" json:"runtime,omitempty"`
	XXX_unrecognized	[]byte		`json:"-"`
}

func (this *Job) Reset()		{ *this = Job{} }
func (this *Job) String() string	{ return proto.CompactTextString(this) }
func (*Job) ProtoMessage()		{}

func (this *Job) GetId() uint32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *Job) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func (this *Job) GetUser() string {
	if this != nil && this.User != nil {
		return *this.User
	}
	return ""
}

func (this *Job) GetPriority() float64 {
	if this != nil && this.Priority != nil {
		return *this.Priority
	}
	return 0
}

func (this *Job) GetStarted() uint64 {
	if this != nil && this.Started != nil {
		return *this.Started
	}
	return 0
}

func (this *Job) GetLastupdate() uint64 {
	if this != nil && this.Lastupdate != nil {
		return *this.Lastupdate
	}
	return 0
}

func (this *Job) GetProgress() float64 {
	if this != nil && this.Progress != nil {
		return *this.Progress
	}
	return 0
}

func (this *Job) GetTaskssuccess() uint32 {
	if this != nil && this.Taskssuccess != nil {
		return *this.Taskssuccess
	}
	return 0
}

func (this *Job) GetTasksfailed() uint32 {
	if this != nil && this.Tasksfailed != nil {
		return *this.Tasksfailed
	}
	return 0
}

func (this *Job) GetRuntime() float64 {
	if this != nil && this.Runtime != nil {
		return *this.Runtime
	}
	return 0
}

type InfoResponse struct {
	Job			*Job	`protobuf:"bytes,1,req,name=job" json:"job,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *InfoResponse) Reset()		{ *this = InfoResponse{} }
func (this *InfoResponse) String() string	{ return proto.CompactTextString(this) }
func (*InfoResponse) ProtoMessage()		{}

func (this *InfoResponse) GetJob() *Job {
	if this != nil {
		return this.Job
	}
	return nil
}

type ListResponse struct {
	Jobs			[]*Job	`protobuf:"bytes,1,rep,name=jobs" json:"jobs,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *ListResponse) Reset()		{ *this = ListResponse{} }
func (this *ListResponse) String() string	{ return proto.CompactTextString(this) }
func (*ListResponse) ProtoMessage()		{}

type AuthenticateResponse struct {
	Challenge		*string	`protobuf:"bytes,1,req,name=challenge" json:"challenge,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *AuthenticateResponse) Reset()		{ *this = AuthenticateResponse{} }
func (this *AuthenticateResponse) String() string	{ return proto.CompactTextString(this) }
func (*AuthenticateResponse) ProtoMessage()		{}

func (this *AuthenticateResponse) GetChallenge() string {
	if this != nil && this.Challenge != nil {
		return *this.Challenge
	}
	return ""
}

type KillResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *KillResponse) Reset()		{ *this = KillResponse{} }
func (this *KillResponse) String() string	{ return proto.CompactTextString(this) }
func (*KillResponse) ProtoMessage()		{}

type SetPriorityResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *SetPriorityResponse) Reset()		{ *this = SetPriorityResponse{} }
func (this *SetPriorityResponse) String() string	{ return proto.CompactTextString(this) }
func (*SetPriorityResponse) ProtoMessage()		{}

type ProgressResponse struct {
	Fitnesses		[]*task.Identify_Fitness	`protobuf:"bytes,1,rep,name=fitnesses" json:"fitnesses,omitempty"`
	Items			[]*task.Progress		`protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized	[]byte				`json:"-"`
}

func (this *ProgressResponse) Reset()		{ *this = ProgressResponse{} }
func (this *ProgressResponse) String() string	{ return proto.CompactTextString(this) }
func (*ProgressResponse) ProtoMessage()		{}

type IdleResponse struct {
	Seconds			*uint64	`protobuf:"varint,1,req,name=seconds" json:"seconds,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *IdleResponse) Reset()		{ *this = IdleResponse{} }
func (this *IdleResponse) String() string	{ return proto.CompactTextString(this) }
func (*IdleResponse) ProtoMessage()		{}

func (this *IdleResponse) GetSeconds() uint64 {
	if this != nil && this.Seconds != nil {
		return *this.Seconds
	}
	return 0
}

type Response struct {
	Type			*CommandType		`protobuf:"varint,1,req,name=type,enum=optimization.messages.command.CommandType" json:"type,omitempty"`
	Status			*bool			`protobuf:"varint,2,req,name=status" json:"status,omitempty"`
	Message			*string			`protobuf:"bytes,3,req,name=message" json:"message,omitempty"`
	List			*ListResponse		`protobuf:"bytes,4,opt,name=list" json:"list,omitempty"`
	Info			*InfoResponse		`protobuf:"bytes,5,opt,name=info" json:"info,omitempty"`
	Kill			*KillResponse		`protobuf:"bytes,6,opt,name=kill" json:"kill,omitempty"`
	Setpriority		*SetPriorityResponse	`protobuf:"bytes,7,opt,name=setpriority" json:"setpriority,omitempty"`
	Authenticate		*AuthenticateResponse	`protobuf:"bytes,8,opt,name=authenticate" json:"authenticate,omitempty"`
	Progress		*ProgressResponse	`protobuf:"bytes,9,opt,name=progress" json:"progress,omitempty"`
	Idle			*IdleResponse		`protobuf:"bytes,10,opt,name=idle" json:"idle,omitempty"`
	XXX_unrecognized	[]byte			`json:"-"`
}

func (this *Response) Reset()		{ *this = Response{} }
func (this *Response) String() string	{ return proto.CompactTextString(this) }
func (*Response) ProtoMessage()		{}

func (this *Response) GetType() CommandType {
	if this != nil && this.Type != nil {
		return *this.Type
	}
	return 0
}

func (this *Response) GetStatus() bool {
	if this != nil && this.Status != nil {
		return *this.Status
	}
	return false
}

func (this *Response) GetMessage() string {
	if this != nil && this.Message != nil {
		return *this.Message
	}
	return ""
}

func (this *Response) GetList() *ListResponse {
	if this != nil {
		return this.List
	}
	return nil
}

func (this *Response) GetInfo() *InfoResponse {
	if this != nil {
		return this.Info
	}
	return nil
}

func (this *Response) GetKill() *KillResponse {
	if this != nil {
		return this.Kill
	}
	return nil
}

func (this *Response) GetSetpriority() *SetPriorityResponse {
	if this != nil {
		return this.Setpriority
	}
	return nil
}

func (this *Response) GetAuthenticate() *AuthenticateResponse {
	if this != nil {
		return this.Authenticate
	}
	return nil
}

func (this *Response) GetProgress() *ProgressResponse {
	if this != nil {
		return this.Progress
	}
	return nil
}

func (this *Response) GetIdle() *IdleResponse {
	if this != nil {
		return this.Idle
	}
	return nil
}

func init() {
	proto.RegisterEnum("optimization.messages.command.CommandType", CommandType_name, CommandType_value)
}
