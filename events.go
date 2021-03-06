package optimization

import (
	"time"
)

type Event func()
type EventQueue chan Event

var Events EventQueue

func init() {
	Events = make(EventQueue, 1024)
}

func (e EventQueue) Loop() {
	for {
		select {
		case event, _ := <-e:
			event()
		}
	}
}

func (e EventQueue) Timeout(ms time.Duration, f Event) {
	go func() {
		time.Sleep(ms)
		e <- f
	}()
}
