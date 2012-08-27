package optimization

import (
	"fmt"
	"net"
	discovery "optimization_messages_discovery"
	"ponyo.epfl.ch/go/get/optimization/go/optimization/log"
	"strconv"
)

var _ = fmt.Println

type Discovered struct {
	Connection string
	Host       string
}

type Greeting struct {
	Connection string
	Host       string
}

type Discovery struct {
	conn *net.UDPConn

	Namespace string
	Host      string
	Port      uint

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
				Connection: msg.GetGreeting().GetConnection(),
				Host:       addr.IP.String(),
			}

			switch msg.GetType() {
			case discovery.Discovery_TypeGreeting:
				log.D("Received greeting from %s", disc.Host)

				Events <- func() {
					for _, g := range d.Greeting {
						g(disc)
					}
				}

			case discovery.Discovery_TypeWakeup:
				log.D("Received wakeup from %s", disc.Host)

				Events <- func() {
					for _, w := range d.Wakeup {
						w(disc)
					}
				}
			}
		})
	}
}

func NewDiscovery(host string, port uint, namespace string) (*Discovery, error) {
	ret := &Discovery{
		Host:      host,
		Port:      port,
		Namespace: namespace,
	}

	ret.Wakeup = []func(disc *Discovered){}
	ret.Greeting = []func(disc *Discovered){}

	host = host + ":" + strconv.FormatUint(uint64(port), 10)

	addr, err := net.ResolveUDPAddr("udp", host)

	if err != nil {
		return nil, err
	}

	ret.conn, err = net.ListenMulticastUDP("udp", nil, addr)

	if err != nil {
		return nil, err
	}

	go ret.read()

	return ret, nil
}

func (d *Discovery) SendWakeup() {
	s := fmt.Sprintf("%v:%v", d.Host, d.Port)

	addr, err := net.ResolveUDPAddr("udp", s)

	if err != nil {
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)

	if err != nil {
		return
	}

	cl := NewClientConnection(conn, new(discovery.Discovery))

	disc := new(discovery.Discovery)
	tp := discovery.Discovery_TypeWakeup

	disc.Type = &tp
	disc.Namespace = &d.Namespace

	disc.Wakeup = new(discovery.Wakeup)

	wad := "multicast://" + s
	disc.Wakeup.Connection = &wad

	cl.Send(disc)
}

func (d *Discovery) Close() {
	d.conn.Close()
}
