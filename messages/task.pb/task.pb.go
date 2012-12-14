package task

import proto "code.google.com/p/goprotobuf/proto"
import "math"

var _ = proto.GetString
var _ = math.Inf

type Identify_Fitness_Type int32

const (
	Identify_Fitness_Maximize	Identify_Fitness_Type	= 0
	Identify_Fitness_Minimize	Identify_Fitness_Type	= 1
)

var Identify_Fitness_Type_name = map[int32]string{
	0:	"Maximize",
	1:	"Minimize",
}
var Identify_Fitness_Type_value = map[string]int32{
	"Maximize":	0,
	"Minimize":	1,
}

func NewIdentify_Fitness_Type(x Identify_Fitness_Type) *Identify_Fitness_Type {
	e := Identify_Fitness_Type(x)
	return &e
}
func (x Identify_Fitness_Type) Enum() *Identify_Fitness_Type {
	p := new(Identify_Fitness_Type)
	*p = x
	return p
}
func (x Identify_Fitness_Type) String() string {
	return proto.EnumName(Identify_Fitness_Type_name, int32(x))
}

type Response_Status int32

const (
	Response_Success	Response_Status	= 0
	Response_Failed		Response_Status	= 1
	Response_Challenge	Response_Status	= 2
)

var Response_Status_name = map[int32]string{
	0:	"Success",
	1:	"Failed",
	2:	"Challenge",
}
var Response_Status_value = map[string]int32{
	"Success":	0,
	"Failed":	1,
	"Challenge":	2,
}

func NewResponse_Status(x Response_Status) *Response_Status {
	e := Response_Status(x)
	return &e
}
func (x Response_Status) Enum() *Response_Status {
	p := new(Response_Status)
	*p = x
	return p
}
func (x Response_Status) String() string {
	return proto.EnumName(Response_Status_name, int32(x))
}

type Response_Failure_Type int32

const (
	Response_Failure_Timeout		Response_Failure_Type	= 0
	Response_Failure_DispatcherNotFound	Response_Failure_Type	= 1
	Response_Failure_NoResponse		Response_Failure_Type	= 2
	Response_Failure_Dispatcher		Response_Failure_Type	= 3
	Response_Failure_Unknown		Response_Failure_Type	= 4
	Response_Failure_WrongRequest		Response_Failure_Type	= 5
	Response_Failure_Disconnected		Response_Failure_Type	= 6
)

var Response_Failure_Type_name = map[int32]string{
	0:	"Timeout",
	1:	"DispatcherNotFound",
	2:	"NoResponse",
	3:	"Dispatcher",
	4:	"Unknown",
	5:	"WrongRequest",
	6:	"Disconnected",
}
var Response_Failure_Type_value = map[string]int32{
	"Timeout":		0,
	"DispatcherNotFound":	1,
	"NoResponse":		2,
	"Dispatcher":		3,
	"Unknown":		4,
	"WrongRequest":		5,
	"Disconnected":		6,
}

func NewResponse_Failure_Type(x Response_Failure_Type) *Response_Failure_Type {
	e := Response_Failure_Type(x)
	return &e
}
func (x Response_Failure_Type) Enum() *Response_Failure_Type {
	p := new(Response_Failure_Type)
	*p = x
	return p
}
func (x Response_Failure_Type) String() string {
	return proto.EnumName(Response_Failure_Type_name, int32(x))
}

type Notification_Type int32

const (
	Notification_Info	Notification_Type	= 0
	Notification_Error	Notification_Type	= 1
	Notification_Warning	Notification_Type	= 2
)

var Notification_Type_name = map[int32]string{
	0:	"Info",
	1:	"Error",
	2:	"Warning",
}
var Notification_Type_value = map[string]int32{
	"Info":		0,
	"Error":	1,
	"Warning":	2,
}

func NewNotification_Type(x Notification_Type) *Notification_Type {
	e := Notification_Type(x)
	return &e
}
func (x Notification_Type) Enum() *Notification_Type {
	p := new(Notification_Type)
	*p = x
	return p
}
func (x Notification_Type) String() string {
	return proto.EnumName(Notification_Type_name, int32(x))
}

type Communication_Type int32

const (
	Communication_CommunicationBatch	Communication_Type	= 0
	Communication_CommunicationTask		Communication_Type	= 1
	Communication_CommunicationResponse	Communication_Type	= 2
	Communication_CommunicationToken	Communication_Type	= 3
	Communication_CommunicationCancel	Communication_Type	= 4
	Communication_CommunicationIdentify	Communication_Type	= 5
	Communication_CommunicationPing		Communication_Type	= 6
	Communication_CommunicationProgress	Communication_Type	= 7
	Communication_CommunicationNotification	Communication_Type	= 8
)

var Communication_Type_name = map[int32]string{
	0:	"CommunicationBatch",
	1:	"CommunicationTask",
	2:	"CommunicationResponse",
	3:	"CommunicationToken",
	4:	"CommunicationCancel",
	5:	"CommunicationIdentify",
	6:	"CommunicationPing",
	7:	"CommunicationProgress",
	8:	"CommunicationNotification",
}
var Communication_Type_value = map[string]int32{
	"CommunicationBatch":		0,
	"CommunicationTask":		1,
	"CommunicationResponse":	2,
	"CommunicationToken":		3,
	"CommunicationCancel":		4,
	"CommunicationIdentify":	5,
	"CommunicationPing":		6,
	"CommunicationProgress":	7,
	"CommunicationNotification":	8,
}

func NewCommunication_Type(x Communication_Type) *Communication_Type {
	e := Communication_Type(x)
	return &e
}
func (x Communication_Type) Enum() *Communication_Type {
	p := new(Communication_Type)
	*p = x
	return p
}
func (x Communication_Type) String() string {
	return proto.EnumName(Communication_Type_name, int32(x))
}

type Token struct {
	Id			*uint32	`protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Response		*string	`protobuf:"bytes,2,req,name=response" json:"response,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *Token) Reset()		{ *this = Token{} }
func (this *Token) String() string	{ return proto.CompactTextString(this) }
func (*Token) ProtoMessage()		{}

func (this *Token) GetId() uint32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *Token) GetResponse() string {
	if this != nil && this.Response != nil {
		return *this.Response
	}
	return ""
}

type Cancel struct {
	Id			*uint32	`protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *Cancel) Reset()		{ *this = Cancel{} }
func (this *Cancel) String() string	{ return proto.CompactTextString(this) }
func (*Cancel) ProtoMessage()		{}

func (this *Cancel) GetId() uint32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

type Ping struct {
	Id			*uint32	`protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *Ping) Reset()		{ *this = Ping{} }
func (this *Ping) String() string	{ return proto.CompactTextString(this) }
func (*Ping) ProtoMessage()		{}

func (this *Ping) GetId() uint32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

type Task struct {
	Id			*uint32				`protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Dispatcher		*string				`protobuf:"bytes,2,req,name=dispatcher" json:"dispatcher,omitempty"`
	Job			*string				`protobuf:"bytes,3,req,name=job" json:"job,omitempty"`
	Optimizer		*string				`protobuf:"bytes,4,req,name=optimizer" json:"optimizer,omitempty"`
	Parameters		[]*Task_Parameter		`protobuf:"bytes,5,rep,name=parameters" json:"parameters,omitempty"`
	Settings		[]*Task_KeyValue		`protobuf:"bytes,6,rep,name=settings" json:"settings,omitempty"`
	Data			[]*Task_KeyValue		`protobuf:"bytes,7,rep,name=data" json:"data,omitempty"`
	Uniqueid		*uint64				`protobuf:"varint,8,opt,name=uniqueid" json:"uniqueid,omitempty"`
	XXX_extensions		map[int32]proto.Extension	`json:"-"`
	XXX_unrecognized	[]byte				`json:"-"`
}

func (this *Task) Reset()		{ *this = Task{} }
func (this *Task) String() string	{ return proto.CompactTextString(this) }
func (*Task) ProtoMessage()		{}

var extRange_Task = []proto.ExtensionRange{
	{1000, 536870911},
}

func (*Task) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Task
}
func (this *Task) ExtensionMap() map[int32]proto.Extension {
	if this.XXX_extensions == nil {
		this.XXX_extensions = make(map[int32]proto.Extension)
	}
	return this.XXX_extensions
}

func (this *Task) GetId() uint32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *Task) GetDispatcher() string {
	if this != nil && this.Dispatcher != nil {
		return *this.Dispatcher
	}
	return ""
}

func (this *Task) GetJob() string {
	if this != nil && this.Job != nil {
		return *this.Job
	}
	return ""
}

func (this *Task) GetOptimizer() string {
	if this != nil && this.Optimizer != nil {
		return *this.Optimizer
	}
	return ""
}

func (this *Task) GetUniqueid() uint64 {
	if this != nil && this.Uniqueid != nil {
		return *this.Uniqueid
	}
	return 0
}

type Task_Parameter struct {
	Name			*string		`protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value			*float64	`protobuf:"fixed64,2,req,name=value" json:"value,omitempty"`
	Min			*float64	`protobuf:"fixed64,3,req,name=min" json:"min,omitempty"`
	Max			*float64	`protobuf:"fixed64,4,req,name=max" json:"max,omitempty"`
	XXX_unrecognized	[]byte		`json:"-"`
}

func (this *Task_Parameter) Reset()		{ *this = Task_Parameter{} }
func (this *Task_Parameter) String() string	{ return proto.CompactTextString(this) }
func (*Task_Parameter) ProtoMessage()		{}

func (this *Task_Parameter) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func (this *Task_Parameter) GetValue() float64 {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return 0
}

func (this *Task_Parameter) GetMin() float64 {
	if this != nil && this.Min != nil {
		return *this.Min
	}
	return 0
}

func (this *Task_Parameter) GetMax() float64 {
	if this != nil && this.Max != nil {
		return *this.Max
	}
	return 0
}

type Task_KeyValue struct {
	Key			*string	`protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value			*string	`protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *Task_KeyValue) Reset()		{ *this = Task_KeyValue{} }
func (this *Task_KeyValue) String() string	{ return proto.CompactTextString(this) }
func (*Task_KeyValue) ProtoMessage()		{}

func (this *Task_KeyValue) GetKey() string {
	if this != nil && this.Key != nil {
		return *this.Key
	}
	return ""
}

func (this *Task_KeyValue) GetValue() string {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return ""
}

type Identify struct {
	Name			*string			`protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	User			*string			`protobuf:"bytes,2,req,name=user" json:"user,omitempty"`
	Priority		*float64		`protobuf:"fixed64,3,req,name=priority" json:"priority,omitempty"`
	Timeout			*float64		`protobuf:"fixed64,4,opt,name=timeout" json:"timeout,omitempty"`
	Version			*uint32			`protobuf:"varint,5,opt,name=version" json:"version,omitempty"`
	Fitness			[]*Identify_Fitness	`protobuf:"bytes,6,rep,name=fitness" json:"fitness,omitempty"`
	XXX_unrecognized	[]byte			`json:"-"`
}

func (this *Identify) Reset()		{ *this = Identify{} }
func (this *Identify) String() string	{ return proto.CompactTextString(this) }
func (*Identify) ProtoMessage()		{}

func (this *Identify) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func (this *Identify) GetUser() string {
	if this != nil && this.User != nil {
		return *this.User
	}
	return ""
}

func (this *Identify) GetPriority() float64 {
	if this != nil && this.Priority != nil {
		return *this.Priority
	}
	return 0
}

func (this *Identify) GetTimeout() float64 {
	if this != nil && this.Timeout != nil {
		return *this.Timeout
	}
	return 0
}

func (this *Identify) GetVersion() uint32 {
	if this != nil && this.Version != nil {
		return *this.Version
	}
	return 0
}

type Identify_Fitness struct {
	Type			*Identify_Fitness_Type	`protobuf:"varint,1,req,name=type,enum=optimization.messages.task.Identify_Fitness_Type" json:"type,omitempty"`
	Name			*string			`protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	XXX_unrecognized	[]byte			`json:"-"`
}

func (this *Identify_Fitness) Reset()		{ *this = Identify_Fitness{} }
func (this *Identify_Fitness) String() string	{ return proto.CompactTextString(this) }
func (*Identify_Fitness) ProtoMessage()		{}

func (this *Identify_Fitness) GetType() Identify_Fitness_Type {
	if this != nil && this.Type != nil {
		return *this.Type
	}
	return 0
}

func (this *Identify_Fitness) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

type Batch struct {
	Tasks			[]*Task		`protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
	Progress		*float64	`protobuf:"fixed64,2,req,name=progress" json:"progress,omitempty"`
	XXX_unrecognized	[]byte		`json:"-"`
}

func (this *Batch) Reset()		{ *this = Batch{} }
func (this *Batch) String() string	{ return proto.CompactTextString(this) }
func (*Batch) ProtoMessage()		{}

func (this *Batch) GetProgress() float64 {
	if this != nil && this.Progress != nil {
		return *this.Progress
	}
	return 0
}

type Progress struct {
	Tick			*uint64			`protobuf:"varint,1,req,name=tick" json:"tick,omitempty"`
	Terms			[]*Progress_Term	`protobuf:"bytes,2,rep,name=terms" json:"terms,omitempty"`
	XXX_unrecognized	[]byte			`json:"-"`
}

func (this *Progress) Reset()		{ *this = Progress{} }
func (this *Progress) String() string	{ return proto.CompactTextString(this) }
func (*Progress) ProtoMessage()		{}

func (this *Progress) GetTick() uint64 {
	if this != nil && this.Tick != nil {
		return *this.Tick
	}
	return 0
}

type Progress_Term struct {
	Best			*float64	`protobuf:"fixed64,1,req,name=best" json:"best,omitempty"`
	Mean			*float64	`protobuf:"fixed64,2,req,name=mean" json:"mean,omitempty"`
	XXX_unrecognized	[]byte		`json:"-"`
}

func (this *Progress_Term) Reset()		{ *this = Progress_Term{} }
func (this *Progress_Term) String() string	{ return proto.CompactTextString(this) }
func (*Progress_Term) ProtoMessage()		{}

func (this *Progress_Term) GetBest() float64 {
	if this != nil && this.Best != nil {
		return *this.Best
	}
	return 0
}

func (this *Progress_Term) GetMean() float64 {
	if this != nil && this.Mean != nil {
		return *this.Mean
	}
	return 0
}

type Response struct {
	Id			*uint32				`protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Status			*Response_Status		`protobuf:"varint,2,req,name=status,enum=optimization.messages.task.Response_Status" json:"status,omitempty"`
	Fitness			[]*Response_Fitness		`protobuf:"bytes,3,rep,name=fitness" json:"fitness,omitempty"`
	Challenge		*string				`protobuf:"bytes,4,opt,name=challenge" json:"challenge,omitempty"`
	Failure			*Response_Failure		`protobuf:"bytes,5,opt,name=failure" json:"failure,omitempty"`
	Data			[]*Response_KeyValue		`protobuf:"bytes,6,rep,name=data" json:"data,omitempty"`
	Uniqueid		*uint64				`protobuf:"varint,7,opt,name=uniqueid" json:"uniqueid,omitempty"`
	XXX_extensions		map[int32]proto.Extension	`json:"-"`
	XXX_unrecognized	[]byte				`json:"-"`
}

func (this *Response) Reset()		{ *this = Response{} }
func (this *Response) String() string	{ return proto.CompactTextString(this) }
func (*Response) ProtoMessage()		{}

var extRange_Response = []proto.ExtensionRange{
	{1000, 536870911},
}

func (*Response) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Response
}
func (this *Response) ExtensionMap() map[int32]proto.Extension {
	if this.XXX_extensions == nil {
		this.XXX_extensions = make(map[int32]proto.Extension)
	}
	return this.XXX_extensions
}

func (this *Response) GetId() uint32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *Response) GetStatus() Response_Status {
	if this != nil && this.Status != nil {
		return *this.Status
	}
	return 0
}

func (this *Response) GetChallenge() string {
	if this != nil && this.Challenge != nil {
		return *this.Challenge
	}
	return ""
}

func (this *Response) GetFailure() *Response_Failure {
	if this != nil {
		return this.Failure
	}
	return nil
}

func (this *Response) GetUniqueid() uint64 {
	if this != nil && this.Uniqueid != nil {
		return *this.Uniqueid
	}
	return 0
}

type Response_Fitness struct {
	Name			*string		`protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value			*float64	`protobuf:"fixed64,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized	[]byte		`json:"-"`
}

func (this *Response_Fitness) Reset()		{ *this = Response_Fitness{} }
func (this *Response_Fitness) String() string	{ return proto.CompactTextString(this) }
func (*Response_Fitness) ProtoMessage()		{}

func (this *Response_Fitness) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func (this *Response_Fitness) GetValue() float64 {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return 0
}

type Response_KeyValue struct {
	Key			*string	`protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value			*string	`protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *Response_KeyValue) Reset()		{ *this = Response_KeyValue{} }
func (this *Response_KeyValue) String() string	{ return proto.CompactTextString(this) }
func (*Response_KeyValue) ProtoMessage()	{}

func (this *Response_KeyValue) GetKey() string {
	if this != nil && this.Key != nil {
		return *this.Key
	}
	return ""
}

func (this *Response_KeyValue) GetValue() string {
	if this != nil && this.Value != nil {
		return *this.Value
	}
	return ""
}

type Response_Failure struct {
	Type			*Response_Failure_Type	`protobuf:"varint,1,req,name=type,enum=optimization.messages.task.Response_Failure_Type" json:"type,omitempty"`
	Message			*string			`protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized	[]byte			`json:"-"`
}

func (this *Response_Failure) Reset()		{ *this = Response_Failure{} }
func (this *Response_Failure) String() string	{ return proto.CompactTextString(this) }
func (*Response_Failure) ProtoMessage()		{}

func (this *Response_Failure) GetType() Response_Failure_Type {
	if this != nil && this.Type != nil {
		return *this.Type
	}
	return 0
}

func (this *Response_Failure) GetMessage() string {
	if this != nil && this.Message != nil {
		return *this.Message
	}
	return ""
}

type Notification struct {
	Type			*Notification_Type	`protobuf:"varint,1,req,name=type,enum=optimization.messages.task.Notification_Type" json:"type,omitempty"`
	Message			*string			`protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	XXX_unrecognized	[]byte			`json:"-"`
}

func (this *Notification) Reset()		{ *this = Notification{} }
func (this *Notification) String() string	{ return proto.CompactTextString(this) }
func (*Notification) ProtoMessage()		{}

func (this *Notification) GetType() Notification_Type {
	if this != nil && this.Type != nil {
		return *this.Type
	}
	return 0
}

func (this *Notification) GetMessage() string {
	if this != nil && this.Message != nil {
		return *this.Message
	}
	return ""
}

type Communication struct {
	Type			*Communication_Type	`protobuf:"varint,1,req,name=type,enum=optimization.messages.task.Communication_Type" json:"type,omitempty"`
	Batch			*Batch			`protobuf:"bytes,2,opt,name=batch" json:"batch,omitempty"`
	Task			*Task			`protobuf:"bytes,3,opt,name=task" json:"task,omitempty"`
	Response		*Response		`protobuf:"bytes,4,opt,name=response" json:"response,omitempty"`
	Token			*Token			`protobuf:"bytes,5,opt,name=token" json:"token,omitempty"`
	Cancel			*Cancel			`protobuf:"bytes,6,opt,name=cancel" json:"cancel,omitempty"`
	Identify		*Identify		`protobuf:"bytes,7,opt,name=identify" json:"identify,omitempty"`
	Ping			*Ping			`protobuf:"bytes,8,opt,name=ping" json:"ping,omitempty"`
	Progress		*Progress		`protobuf:"bytes,9,opt,name=progress" json:"progress,omitempty"`
	Notification		*Notification		`protobuf:"bytes,10,opt,name=notification" json:"notification,omitempty"`
	XXX_unrecognized	[]byte			`json:"-"`
}

func (this *Communication) Reset()		{ *this = Communication{} }
func (this *Communication) String() string	{ return proto.CompactTextString(this) }
func (*Communication) ProtoMessage()		{}

func (this *Communication) GetType() Communication_Type {
	if this != nil && this.Type != nil {
		return *this.Type
	}
	return 0
}

func (this *Communication) GetBatch() *Batch {
	if this != nil {
		return this.Batch
	}
	return nil
}

func (this *Communication) GetTask() *Task {
	if this != nil {
		return this.Task
	}
	return nil
}

func (this *Communication) GetResponse() *Response {
	if this != nil {
		return this.Response
	}
	return nil
}

func (this *Communication) GetToken() *Token {
	if this != nil {
		return this.Token
	}
	return nil
}

func (this *Communication) GetCancel() *Cancel {
	if this != nil {
		return this.Cancel
	}
	return nil
}

func (this *Communication) GetIdentify() *Identify {
	if this != nil {
		return this.Identify
	}
	return nil
}

func (this *Communication) GetPing() *Ping {
	if this != nil {
		return this.Ping
	}
	return nil
}

func (this *Communication) GetProgress() *Progress {
	if this != nil {
		return this.Progress
	}
	return nil
}

func (this *Communication) GetNotification() *Notification {
	if this != nil {
		return this.Notification
	}
	return nil
}

func init() {
	proto.RegisterEnum("optimization.messages.task.Identify_Fitness_Type", Identify_Fitness_Type_name, Identify_Fitness_Type_value)
	proto.RegisterEnum("optimization.messages.task.Response_Status", Response_Status_name, Response_Status_value)
	proto.RegisterEnum("optimization.messages.task.Response_Failure_Type", Response_Failure_Type_name, Response_Failure_Type_value)
	proto.RegisterEnum("optimization.messages.task.Notification_Type", Notification_Type_name, Notification_Type_value)
	proto.RegisterEnum("optimization.messages.task.Communication_Type", Communication_Type_name, Communication_Type_value)
}
