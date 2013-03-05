package net

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type Address struct {
	Protocol string
	Host     string
	Port     string
}

func (a *Address) Dial() (net.Conn, error) {
	var ads string

	if a.Port != "" {
		ads = net.JoinHostPort(a.Host, a.Port)
	} else {
		ads = a.Host
	}

	return net.Dial(a.Protocol, ads)
}

func (a *Address) underlyingProtocol() string {
	protocol := a.Protocol

	if protocol == "multicast" {
		protocol = "udp"
	}

	return protocol
}

func (a *Address) Listen() (net.Listener, error) {
	var ads string

	if a.Port != "" {
		ads = net.JoinHostPort(a.Host, a.Port)
	} else {
		ads = a.Host
	}

	listener, err := net.Listen(a.underlyingProtocol(), ads)

	if a.Port == "0" {
		// check real assigned port here if needed!

		if tcp, ok := listener.Addr().(*net.TCPAddr); ok {
			a.Port = fmt.Sprintf("%v", tcp.Port)
		}
	}


	return listener, err
}

func (a *Address) IPAddr() (*net.IPAddr, error) {
	return net.ResolveIPAddr(a.underlyingProtocol(), a.Host)
}

func (a *Address) Resolve() error {
	if a.Protocol == "unix" {
		return nil
	}

	isip := (net.ParseIP(a.Host) != nil)

	var ip *net.IPAddr
	ip = nil

	if isip {
		var err error

		ip, err = a.IPAddr()

		if err != nil {
			return err
		}
	}

	if !isip || ip.IP.IsUnspecified() {
		var name string

		if ip != nil && ip.IP.IsUnspecified() {
			var err error

			name, err = os.Hostname()

			if err != nil {
				return err
			}
		} else {
			name = a.Host
		}

		addrs, err := net.LookupHost(name)

		if err != nil {
			return err
		}

		if len(addrs) != 0 {
			a.Host = addrs[0]
		}
	} else {
		a.Host = ip.String()
	}

	ret, err := net.LookupPort(a.underlyingProtocol(), a.Port)

	if err != nil {
		return err
	}

	a.Port = fmt.Sprintf("%v", ret)

	return nil
}

func (a *Address) String() string {
	return fmt.Sprintf("%v://%v", a.Protocol, net.JoinHostPort(a.Host, a.Port))
}

func ParseAddress(constr string) *Address {
	return ParseAddressWithDefaultProtocol(constr, "tcp")
}

func ParseAddressWithDefaultProtocol(constr string, defaultProtocol string) *Address {
	ret := new(Address)

	idx := strings.Index(constr, "://")

	if idx == -1 {
		ret.Protocol = defaultProtocol
	} else {
		ret.Protocol = constr[0:idx]
		constr = constr[idx+3:]
	}

	ret.Host, ret.Port, _ = net.SplitHostPort(constr)

	if ret.Host == "" {
		ret.Host = "0.0.0.0"
	}

	return ret
}
