package optimization

import (
	"ponyo.epfl.ch/go/get/optimization/go/optimization/messages/task.pb"
	"os"
	"os/signal"
	"syscall"
	"time"
	"errors"
	"io"
	"fmt"
)

type dispatcherResponse struct {
	Response *task.Response
	Error error
}

type Dispatcher struct {
	reader io.ReadCloser
	writer io.WriteCloser

	Task *task.Task
	Response *task.Response

	Parameters map[string]float64
	Settings map[string]string

	Shutdown chan bool
}

type DispatcherError struct {
	Type task.Response_Failure_Type
	Message string
}

func (x *DispatcherError) Error() string {
	return fmt.Sprintf("[%s] %s", x.Type.String(), x.Message)
}

func (x *DispatcherError) ToFailure() *task.Response_Failure {
	return &task.Response_Failure {
		Type: &x.Type,
		Message: &x.Message,
	}
}

func NewDispatcherError(tp task.Response_Failure_Type, message string) error {
	return &DispatcherError {
		Type: tp,
		Message: message,
	}
}

type DispatchHandler func(x *Dispatcher) error

func NewDispatcher(reader io.ReadCloser, writer io.WriteCloser) *Dispatcher {
	return &Dispatcher {
		reader: reader,
		writer: writer,

		Settings: make(map[string]string),
		Parameters: make(map[string]float64),
	}
}

func (x *Dispatcher) SetFitness(fitness map[string]float64) {
	if x.Response == nil {
		return
	}

	x.Response.Fitness = make([]*task.Response_Fitness, 0, len(fitness))

	for k, v := range fitness {
		x.AddFitness(k, v)
	}
}

func (x *Dispatcher) AddFitness(name string, value float64) {
	if x.Response == nil {
		return
	}

	x.Response.Fitness = append(x.Response.Fitness, &task.Response_Fitness {
		Name: &name,
		Value: &value,
	})
}

func (x *Dispatcher) SetData(data map[string]string) {
	if x.Response == nil {
		return
	}

	x.Response.Data = make([]*task.Response_KeyValue, 0, len(data))

	for k, v := range data {
		x.AddData(k, v)
	}
}

func (x *Dispatcher) AddData(name string, value string) {
	if x.Response == nil {
		return
	}

	x.Response.Data = append(x.Response.Data, &task.Response_KeyValue {
		Key: &name,
		Value: &value,
	})
}

func (x *Dispatcher) extract() {
	if x.Task == nil {
		return
	}

	for _, s := range x.Task.Settings {
		name := s.GetKey()
		value := s.GetValue()

		x.Settings[name] = value
	}

	for _, p := range x.Task.Parameters {
		name := p.GetName()
		value := p.GetValue()

		x.Parameters[name] = value
	}
}

func (x *Dispatcher) ReadTask() error {
	if x.Task != nil {
		return nil
	}

	return ReadCommunication(x.reader, new(task.Task), func(tiface interface{}, err error) bool {
		if err != nil {
			return false
		}

		if tiface == nil {
			return true
		}

		x.Task = tiface.(*task.Task)
		x.Response = new(task.Response)

		x.Response.Id = x.Task.Id
		x.Response.Uniqueid = x.Task.Uniqueid

		x.extract()

		return false
	})
}

func (x *Dispatcher) SetResponse(r *task.Response) {
	x.Response.Status = r.Status

	x.Response.Fitness = r.Fitness
	x.Response.Data = r.Data
	x.Response.Failure = r.Failure
}

func (x *Dispatcher) WriteResponse() error {
	b, err := EncodeCommunication(x.Response)

	if err != nil {
		return err
	}

	_, err = x.writer.Write(b)

	if err != nil {
		return err
	}

	return nil
}

func (x *Dispatcher) runHandler(handler DispatchHandler, response chan error, shutdown chan bool) {
	go func() {
		if err := x.ReadTask(); err != nil {
			response <- err
		} else if x.Task == nil {
			response <- errors.New("Cancelled")
		} else {
			response <- handler(x)
		}
	}()

	select {
	case <-shutdown:
		x.reader.Close()
		x.writer.Close()

		if x.Shutdown != nil {
			x.Shutdown <- true
		} else {
			response <- errors.New("Cancelled")
		}
	}
}

func (x *Dispatcher) handleResponse(err error) error {
	if x.Task == nil || x.Response == nil {
		return nil
	}

	x.Response.Id = x.Task.Id
	x.Response.Uniqueid = x.Task.Uniqueid

	var r *DispatcherError

	if err != nil {
		var ok bool

		// Send failed response
		if r, ok = err.(*DispatcherError); !ok {
			r = NewDispatcherError(task.Response_Failure_Unknown, err.Error()).(*DispatcherError)
		}

		status := task.Response_Failed
		x.Response.Status = &status
		x.Response.Failure = r.ToFailure()
	} else {
		// Woot woot, send succces!
		status := task.Response_Success
		x.Response.Status = &status
	}

	if err := x.WriteResponse(); err != nil {
		return err
	}

	return r
}

func (x *Dispatcher) Run(handler DispatchHandler) error {
	// Setup signals
	signals := make(chan os.Signal, 10)
	signal.Notify(signals, os.Signal(syscall.SIGTERM), os.Interrupt)

	finished := make(chan error, 10)
	shutdown := make(chan bool, 10)

	go x.runHandler(handler, finished, shutdown)

	for {
		select {
		case <-signals:
			// Cancel the dispatcher. This should in turn send something on
			// the finished channel, but install a timeout just in
			// case
			shutdown <- true

			go func() {
				time.Sleep(2 * time.Second)
				os.Exit(1)
			}()
		case response := <-finished:
			return x.handleResponse(response)
		}
	}

	return nil
}
