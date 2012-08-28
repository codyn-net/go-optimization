package net

import (
	"strings"
	"net"
	"fmt"
	"strconv"
)

type Address struct {
	Protocol string
	Host string
	Port int64
}

func (a *Address) Dial() (net.Conn, error) {
	var ads string

	if a.Port != -1 {
		ads = fmt.Sprintf("%v:%v", a.Host, a.Port)
	} else {
		ads = a.Host
	}

	return net.Dial(a.Protocol, ads)
}

func (a *Address) String() string {
	return fmt.Sprintf("%v://%v:%v", a.Protocol, a.Host, a.Port)
}

func ParseAddress(constr string) *Address {
	return ParseAddressWithDefaultProtocol(constr, "tcp");
}

func ParseAddressWithDefaultProtocol(constr string, defaultProtocol string) *Address {
	ret := new(Address)

	idx := strings.Index(constr, "://")

	if idx == -1 {
		ret.Protocol = defaultProtocol
	} else {
		ret.Protocol = constr[0:idx]
		constr = constr[idx + 3:]
	}

	lastcol := strings.LastIndex(constr, ":")

	if lastcol == -1 {
		ret.Host = constr
		ret.Port = -1
	} else {
		ret.Host = constr[0:lastcol]
		ret.Port, _ = strconv.ParseInt(constr[lastcol + 1:], 10, 32)
	}

	return ret
}
