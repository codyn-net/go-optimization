package optimization

import (
	"code.google.com/p/goprotobuf/proto"
	"errors"
	"fmt"
	"net"
	optinet "ponyo.epfl.ch/go/get/optimization/go/optimization/net"
	"strings"
)

type State uint32

const (
	_ State = iota // skip 0
	Connecting
	Connected
	Disconnected
)

type Client struct {
	Comm       net.Conn
	Connection string
	Host       string
	State      State

	MessageTemplate proto.Message

	OnMessage *Signal
	OnState   *Signal
}

func NewClientConnection(comm net.Conn, msg proto.Message) *Client {
	ret := NewClient(msg)

	ret.Comm = comm
	ret.State = Connected

	go ret.readLoop()

	return ret
}

func NewClient(msg proto.Message) *Client {
	return &Client{
		State:           Disconnected,
		OnMessage:       NewSyncSignal(func(proto.Message) {}),
		OnState:         NewSignal(func() {}),
		MessageTemplate: msg,
	}
}

func (c *Client) setState(state State) {
	if c.State == state {
		return
	}

	c.State = state

	c.OnState.Emit()
}

func (c *Client) Disconnect() {
	if c.State != Connected {
		return
	}

	c.Comm.Close()
}

func (c *Client) Connect(connection string, connected func(error)) {
	addr := optinet.ParseAddress(connection)

	if addr == nil {
		if connected != nil {
			connected(errors.New(fmt.Sprintf("Could not resolve address: %s", connection)))
		}

		return
	}

	if c.State != Disconnected {
		if connected != nil {
			connected(errors.New("Already connected"))
		}

		return
	}

	c.setState(Connecting)

	// Connect async in a go routine
	go func() {
		comm, err := addr.Dial()

		Events <- func() {
			if comm != nil {
				res, _ := net.LookupAddr(addr.Host)

				if len(res) > 0 {
					c.Host = res[0]
				}

				c.Comm = comm
				c.setState(Connected)

				c.Connection = connection

				if connected != nil {
					connected(nil)
				}

				// Start the read loop in a go routine
				go c.readLoop()
			} else {
				if err != nil && connected != nil {
					connected(err)
				}

				c.setState(Disconnected)
			}
		}
	}()
}

func (c *Client) Send(msg proto.Message) {
	data, err := EncodeMessage(msg)

	if err != nil {
		return
	}

	go c.Comm.Write(data)
}

func (c *Client) readLoop() {
	if len(c.Host) == 0 {
		hname := c.Comm.RemoteAddr().String()
		pos := strings.Index(hname, ":")

		if pos >= 0 {
			hname = hname[:pos]
		}

		res, _ := net.LookupAddr(hname)

		if len(res) > 0 {
			c.Host = res[0]
		}
	}

	ReadMessages(c.Comm, c.MessageTemplate, func(msg interface{}, err error) bool {
		if err != nil {
			Events <- func() {
				c.setState(Disconnected)
			}
		} else {
			c.OnMessage.Emit(msg)
		}

		return true
	})
}
