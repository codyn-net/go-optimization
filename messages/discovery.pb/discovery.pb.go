package discovery

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Discovery_Type int32

const (
	Discovery_TypeGreeting	Discovery_Type	= 0
	Discovery_TypeWakeup	Discovery_Type	= 1
)

var Discovery_Type_name = map[int32]string{
	0:	"TypeGreeting",
	1:	"TypeWakeup",
}
var Discovery_Type_value = map[string]int32{
	"TypeGreeting":	0,
	"TypeWakeup":	1,
}

func (x Discovery_Type) Enum() *Discovery_Type {
	p := new(Discovery_Type)
	*p = x
	return p
}
func (x Discovery_Type) String() string {
	return proto.EnumName(Discovery_Type_name, int32(x))
}
func (x Discovery_Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Discovery_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Discovery_Type_value, data, "Discovery_Type")
	if err != nil {
		return err
	}
	*x = Discovery_Type(value)
	return nil
}

type Greeting struct {
	Connection		*string	`protobuf:"bytes,1,req,name=connection" json:"connection,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *Greeting) Reset()		{ *this = Greeting{} }
func (this *Greeting) String() string	{ return proto.CompactTextString(this) }
func (*Greeting) ProtoMessage()		{}

func (this *Greeting) GetConnection() string {
	if this != nil && this.Connection != nil {
		return *this.Connection
	}
	return ""
}

type Wakeup struct {
	Connection		*string	`protobuf:"bytes,1,req,name=connection" json:"connection,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (this *Wakeup) Reset()		{ *this = Wakeup{} }
func (this *Wakeup) String() string	{ return proto.CompactTextString(this) }
func (*Wakeup) ProtoMessage()		{}

func (this *Wakeup) GetConnection() string {
	if this != nil && this.Connection != nil {
		return *this.Connection
	}
	return ""
}

type Discovery struct {
	Type			*Discovery_Type	`protobuf:"varint,1,req,name=type,enum=optimization.messages.discovery.Discovery_Type" json:"type,omitempty"`
	Greeting		*Greeting	`protobuf:"bytes,2,opt,name=greeting" json:"greeting,omitempty"`
	Wakeup			*Wakeup		`protobuf:"bytes,3,opt,name=wakeup" json:"wakeup,omitempty"`
	Namespace		*string		`protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
	XXX_unrecognized	[]byte		`json:"-"`
}

func (this *Discovery) Reset()		{ *this = Discovery{} }
func (this *Discovery) String() string	{ return proto.CompactTextString(this) }
func (*Discovery) ProtoMessage()	{}

func (this *Discovery) GetType() Discovery_Type {
	if this != nil && this.Type != nil {
		return *this.Type
	}
	return 0
}

func (this *Discovery) GetGreeting() *Greeting {
	if this != nil {
		return this.Greeting
	}
	return nil
}

func (this *Discovery) GetWakeup() *Wakeup {
	if this != nil {
		return this.Wakeup
	}
	return nil
}

func (this *Discovery) GetNamespace() string {
	if this != nil && this.Namespace != nil {
		return *this.Namespace
	}
	return ""
}

func init() {
	proto.RegisterEnum("optimization.messages.discovery.Discovery_Type", Discovery_Type_name, Discovery_Type_value)
}
