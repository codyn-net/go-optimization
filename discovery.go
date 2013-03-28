package optimization

import (
	"fmt"
	"net"
	"ponyo.epfl.ch/go/get/optimization-go/optimization/messages/discovery.pb"
	"ponyo.epfl.ch/go/get/optimization-go/optimization/log"
)

var _ = fmt.Println

type Discovered struct {
	Connection string
	Host       string
}

type Discovery struct {
	conn *net.UDPConn

	Namespace string
	Address   string

	Wakeup   []func(disc *Discovered)
	Greeting []func(disc *Discovered)
}

func (d *Discovery) read() {
	buf := make([]byte, 512)

	for {
		n, addr, err := d.conn.ReadFromUDP(buf)

		if err != nil {
			break
		}

		msg := new(discovery.Discovery)

		ExtractMessages(buf[0:n], msg, func() {
			if msg.GetNamespace() != d.Namespace {
				log.Message(log.Discovery|log.Verbose,
					"Discovery does not match namespace (%s but need %s)",
					msg.GetNamespace(), d.Namespace)
				return
			}

			disc := &Discovered{
				Host: addr.IP.String(),
			}

			switch msg.GetType() {
			case discovery.Discovery_TypeGreeting:
				disc.Connection = msg.GetGreeting().GetConnection()

				log.D("Received greeting from %s (connection: %s)",
					disc.Host,
					disc.Connection)

				Events <- func() {
					for _, g := range d.Greeting {
						g(disc)
					}
				}

			case discovery.Discovery_TypeWakeup:
				disc.Connection = msg.GetWakeup().GetConnection()

				log.D("Received wakeup from %s (connection: %s)",
					disc.Host,
					disc.Connection)

				Events <- func() {
					for _, w := range d.Wakeup {
						w(disc)
					}
				}
			}
		})
	}
}

func NewDiscovery(address string, namespace string) (*Discovery, error) {
	ret := &Discovery{
		Namespace: namespace,
	}

	ret.Wakeup = []func(disc *Discovered){}
	ret.Greeting = []func(disc *Discovered){}

	addr, err := net.ResolveUDPAddr("udp", address)

	ret.Address = address
	ret.conn, err = net.ListenMulticastUDP("udp", nil, addr)

	if err != nil {
		return nil, err
	}

	go ret.read()

	return ret, nil
}

func (d *Discovery) connect() *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp", d.Address)

	if err != nil {
		return nil
	}

	conn, err := net.DialUDP("udp", nil, addr)

	if err != nil {
		return nil
	}

	return conn
}

func (d *Discovery) SendGreeting(connection string) {
	conn := d.connect()

	if conn == nil {
		return
	}

	cl := NewClientConnection(conn, new(discovery.Discovery))

	disc := new(discovery.Discovery)
	tp := discovery.Discovery_TypeGreeting

	disc.Type = &tp
	disc.Namespace = &d.Namespace

	disc.Greeting = new(discovery.Greeting)
	disc.Greeting.Connection = &connection

	cl.Send(disc)
}

func (d *Discovery) SendWakeup() {
	conn := d.connect()

	if conn == nil {
		return
	}

	cl := NewClientConnection(conn, new(discovery.Discovery))

	disc := new(discovery.Discovery)
	tp := discovery.Discovery_TypeWakeup

	disc.Type = &tp
	disc.Namespace = &d.Namespace

	disc.Wakeup = new(discovery.Wakeup)

	disc.Wakeup.Connection = &d.Address

	cl.Send(disc)
}

func (d *Discovery) Close() {
	d.conn.Close()
}
