package monitor

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type MessageType int32

const (
	MessageType_Register	MessageType	= 0
	MessageType_Unregister	MessageType	= 1
	MessageType_Measurement	MessageType	= 2
)

var MessageType_name = map[int32]string{
	0:	"Register",
	1:	"Unregister",
	2:	"Measurement",
}
var MessageType_value = map[string]int32{
	"Register":	0,
	"Unregister":	1,
	"Measurement":	2,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (x MessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}

type RegisterMessage struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *RegisterMessage) Reset()		{ *this = RegisterMessage{} }
func (this *RegisterMessage) String() string	{ return proto.CompactTextString(this) }
func (*RegisterMessage) ProtoMessage()		{}

type UnregisterMessage struct {
	XXX_unrecognized []byte `json:"-"`
}

func (this *UnregisterMessage) Reset()		{ *this = UnregisterMessage{} }
func (this *UnregisterMessage) String() string	{ return proto.CompactTextString(this) }
func (*UnregisterMessage) ProtoMessage()	{}

type MeasurementMessage struct {
	Time			*float64			`protobuf:"fixed64,1,req,name=time" json:"time,omitempty"`
	Cpu			[]*MeasurementMessage_CPU	`protobuf:"bytes,2,rep,name=cpu" json:"cpu,omitempty"`
	Memory			*MeasurementMessage_Memory	`protobuf:"bytes,3,req,name=memory" json:"memory,omitempty"`
	XXX_unrecognized	[]byte				`json:"-"`
}

func (this *MeasurementMessage) Reset()		{ *this = MeasurementMessage{} }
func (this *MeasurementMessage) String() string	{ return proto.CompactTextString(this) }
func (*MeasurementMessage) ProtoMessage()	{}

func (this *MeasurementMessage) GetTime() float64 {
	if this != nil && this.Time != nil {
		return *this.Time
	}
	return 0
}

func (this *MeasurementMessage) GetMemory() *MeasurementMessage_Memory {
	if this != nil {
		return this.Memory
	}
	return nil
}

type MeasurementMessage_Memory struct {
	Total			*uint64	`protobuf:"varint,1,req,name=total" json:"total,omitempty"`
	Used			*uint64	`protobuf:"varint,2,req,name=used" json:"used,omitempty"`
	Free			*uint64	`protobuf:"varint,3,req,name=free" json:"free,omitempty"`
	Shared			*uint64	`protobuf:"varint,4,req,name=shared" json:"shared,omitempty"`
	Buffer			*uint64	`protobuf:"varint,5,req,name=buffer" json:"buffer,omitempty"`
	Cached			*uint64	`protobuf:"varint,6,req,name=cached" json:"cached,omitempty"`
	User			*uint64	`protobuf:"varint,7,req,name=user" json:"user,omitempty"`
	Locked			*uint64	`protobuf:"varint,8,req,name=locked" json:"locked,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *MeasurementMessage_Memory) Reset()		{ *this = MeasurementMessage_Memory{} }
func (this *MeasurementMessage_Memory) String() string	{ return proto.CompactTextString(this) }
func (*MeasurementMessage_Memory) ProtoMessage()	{}

func (this *MeasurementMessage_Memory) GetTotal() uint64 {
	if this != nil && this.Total != nil {
		return *this.Total
	}
	return 0
}

func (this *MeasurementMessage_Memory) GetUsed() uint64 {
	if this != nil && this.Used != nil {
		return *this.Used
	}
	return 0
}

func (this *MeasurementMessage_Memory) GetFree() uint64 {
	if this != nil && this.Free != nil {
		return *this.Free
	}
	return 0
}

func (this *MeasurementMessage_Memory) GetShared() uint64 {
	if this != nil && this.Shared != nil {
		return *this.Shared
	}
	return 0
}

func (this *MeasurementMessage_Memory) GetBuffer() uint64 {
	if this != nil && this.Buffer != nil {
		return *this.Buffer
	}
	return 0
}

func (this *MeasurementMessage_Memory) GetCached() uint64 {
	if this != nil && this.Cached != nil {
		return *this.Cached
	}
	return 0
}

func (this *MeasurementMessage_Memory) GetUser() uint64 {
	if this != nil && this.User != nil {
		return *this.User
	}
	return 0
}

func (this *MeasurementMessage_Memory) GetLocked() uint64 {
	if this != nil && this.Locked != nil {
		return *this.Locked
	}
	return 0
}

type MeasurementMessage_CPU struct {
	Total			*uint64	`protobuf:"varint,1,req,name=total" json:"total,omitempty"`
	User			*uint64	`protobuf:"varint,2,req,name=user" json:"user,omitempty"`
	Nice			*uint64	`protobuf:"varint,3,req,name=nice" json:"nice,omitempty"`
	Sys			*uint64	`protobuf:"varint,4,req,name=sys" json:"sys,omitempty"`
	Idle			*uint64	`protobuf:"varint,5,req,name=idle" json:"idle,omitempty"`
	Iowait			*uint64	`protobuf:"varint,6,req,name=iowait" json:"iowait,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *MeasurementMessage_CPU) Reset()		{ *this = MeasurementMessage_CPU{} }
func (this *MeasurementMessage_CPU) String() string	{ return proto.CompactTextString(this) }
func (*MeasurementMessage_CPU) ProtoMessage()		{}

func (this *MeasurementMessage_CPU) GetTotal() uint64 {
	if this != nil && this.Total != nil {
		return *this.Total
	}
	return 0
}

func (this *MeasurementMessage_CPU) GetUser() uint64 {
	if this != nil && this.User != nil {
		return *this.User
	}
	return 0
}

func (this *MeasurementMessage_CPU) GetNice() uint64 {
	if this != nil && this.Nice != nil {
		return *this.Nice
	}
	return 0
}

func (this *MeasurementMessage_CPU) GetSys() uint64 {
	if this != nil && this.Sys != nil {
		return *this.Sys
	}
	return 0
}

func (this *MeasurementMessage_CPU) GetIdle() uint64 {
	if this != nil && this.Idle != nil {
		return *this.Idle
	}
	return 0
}

func (this *MeasurementMessage_CPU) GetIowait() uint64 {
	if this != nil && this.Iowait != nil {
		return *this.Iowait
	}
	return 0
}

type Message struct {
	Type			*MessageType		`protobuf:"varint,1,req,name=type,enum=optimization.messages.monitor.MessageType" json:"type,omitempty"`
	Register		*RegisterMessage	`protobuf:"bytes,2,opt,name=register" json:"register,omitempty"`
	Unregister		*UnregisterMessage	`protobuf:"bytes,3,opt,name=unregister" json:"unregister,omitempty"`
	Measurement		*MeasurementMessage	`protobuf:"bytes,4,opt,name=measurement" json:"measurement,omitempty"`
	XXX_unrecognized	[]byte			`json:"-"`
}

func (this *Message) Reset()		{ *this = Message{} }
func (this *Message) String() string	{ return proto.CompactTextString(this) }
func (*Message) ProtoMessage()		{}

func (this *Message) GetType() MessageType {
	if this != nil && this.Type != nil {
		return *this.Type
	}
	return 0
}

func (this *Message) GetRegister() *RegisterMessage {
	if this != nil {
		return this.Register
	}
	return nil
}

func (this *Message) GetUnregister() *UnregisterMessage {
	if this != nil {
		return this.Unregister
	}
	return nil
}

func (this *Message) GetMeasurement() *MeasurementMessage {
	if this != nil {
		return this.Measurement
	}
	return nil
}

func init() {
	proto.RegisterEnum("optimization.messages.monitor.MessageType", MessageType_name, MessageType_value)
}
