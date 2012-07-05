package log

import (
	//"log"
	"os"
	"strings"
	"fmt"
	"time"
)

type Flags uint32

const (
	Discovery Flags = 1 << iota
	Master
	Worker
	Signal
	Command
	Verbose
)

var flags Flags

var names = map[Flags]string {
	Discovery: "discovery",
	Master: "master",
	Worker: "worker",
	Signal: "signal",
	Command: "command",
	Verbose: "verbose",
}

func init() {
	for k, v := range names {
		uname := strings.ToUpper(v)

		if len(os.Getenv("DEBUG_" + uname)) > 0 {
			flags |= k
		}
	}
}

func C(format string, v... interface{}) {
	Message(Command, format, v...)
}


func S(format string, v... interface{}) {
	Message(Signal, format, v...)
}

func D(format string, v... interface{}) {
	Message(Discovery, format, v...)
}

func W(format string, v... interface{}) {
	Message(Worker, format, v...)
}

func M(format string, v... interface{}) {
	Message(Master, format, v...)
}

func E(format string, v... interface{}) {
	Message(0, format, v...)
}

func Message(f Flags, format string, v... interface{}) {
	if f != 0 && f & flags == 0 {
		return
	}

	nms := make([]string, 0)

	for k, v := range names {
		if k & f != 0 {
			nms = append(nms, v)
		}
	}

	now := time.Now()
	t := fmt.Sprintf("%d:%d:%d", now.Hour(), now.Minute(), now.Second())

	formatted := fmt.Sprintf(format, v...)

	fmt.Fprintf(os.Stderr, "[%s] %s: %s\n", t, strings.Join(nms, ", "), formatted)
}
