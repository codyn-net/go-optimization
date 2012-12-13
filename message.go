package optimization

import (
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"io"
	task "optimization/messages/task.pb"
	"strconv"
)

var _ = fmt.Println

func EncodeMessage(msg proto.Message) ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	data, err := proto.Marshal(msg)

	if err != nil {
		return nil, err
	}

	buf.WriteString(strconv.Itoa(len(data)))
	buf.WriteByte(' ')
	buf.Write(data)

	return buf.Bytes(), nil
}

func EncodeCommunication(msg proto.Message) ([]byte, error) {
	ret := new(task.Communication)

	switch x := msg.(type) {
		case *task.Batch:
			tp := task.Communication_CommunicationBatch

			ret.Type = &tp
			ret.Batch = x
		case *task.Task:
			tp := task.Communication_CommunicationTask

			ret.Type = &tp
			ret.Task = x
		case *task.Response:
			tp := task.Communication_CommunicationResponse

			ret.Type = &tp
			ret.Response = x
		case *task.Token:
			tp := task.Communication_CommunicationToken

			ret.Type = &tp
			ret.Token = x
		case *task.Cancel:
			tp := task.Communication_CommunicationCancel

			ret.Type = &tp
			ret.Cancel = x
		case *task.Identify:
			tp := task.Communication_CommunicationIdentify

			ret.Type = &tp
			ret.Identify = x
		case *task.Ping:
			tp := task.Communication_CommunicationPing

			ret.Type = &tp
			ret.Ping = x
		case *task.Progress:
			tp := task.Communication_CommunicationProgress

			ret.Type = &tp
			ret.Progress = x
		case *task.Notification:
			tp := task.Communication_CommunicationNotification

			ret.Type = &tp
			ret.Notification = x
		default:
			return nil, fmt.Errorf("Unknown communication type")
	}

	return EncodeMessage(ret)
}

func NewCommunication(tp task.Communication_Type, ft func(*task.Communication)) *task.Communication {
	ret := new(task.Communication)

	ret.Type = &tp
	ft(ret)

	return ret
}

func ExtractMessages(data []byte, ret proto.Message, cb func()) int {
	// Read until space
	buf := bytes.NewBuffer(data)
	n := 0

	for {
		num, err := buf.ReadString(' ')

		if err != nil {
			break
		}

		val, err := strconv.ParseInt(num[:len(num)-1], 10, 32)

		if err != nil {
			break
		}

		if buf.Len() < int(val) {
			break
		}

		msg := make([]byte, val)
		nn, err := buf.Read(msg)

		if err != nil || nn != len(msg) {
			break
		}

		if err = proto.Unmarshal(msg, ret); err != nil {
			break
		}

		n += nn + len(num)

		cb()

		ret.Reset()
	}

	return n
}

func ReadMessages(reader io.Reader, ret proto.Message, cb func(interface{}, error) bool) {
	buf := new(bytes.Buffer)
	data := make([]byte, 512)

	for {
		n, err := reader.Read(data)

		// append to the buffer
		buf.Write(data[:n])

		b := buf.Bytes()
		var cont bool
		cont = true

		n = ExtractMessages(b, ret, func() {
			cnt := cb(proto.Clone(ret), nil)

			if !cnt {
				cont = false
			}
		})

		if n > 0 {
			buf = bytes.NewBuffer(b[n:])
		}

		if err != nil {
			if err == io.EOF {
				err = nil
			}

			cb(nil, err)
			break
		}

		if !cont {
			break
		}
	}
}

func ReadCommunication(reader io.Reader, ret proto.Message, cb func(interface{}, error) bool) {
	comm := new(task.Communication)

	ReadMessages(reader, comm, func(cc interface{}, err error) bool {
		if err != nil {
			return cb(nil, err)
		}

		if cc == nil {
			return cb(nil, nil)
		}

		c := cc.(*task.Communication)

		var tp task.Communication_Type
		var msg interface{}

		// Convert to correct type
		switch ret.(type) {
			case *task.Batch:
				tp = task.Communication_CommunicationBatch
				msg = c.Batch
			case *task.Task:
				tp = task.Communication_CommunicationTask
				msg = c.Task
			case *task.Response:
				tp = task.Communication_CommunicationResponse

				msg = c.Response
			case *task.Token:
				tp = task.Communication_CommunicationToken
				msg = c.Token
			case *task.Cancel:
				tp = task.Communication_CommunicationCancel
				msg = c.Cancel
			case *task.Identify:
				tp = task.Communication_CommunicationIdentify
				msg = c.Identify
			case *task.Ping:
				tp = task.Communication_CommunicationPing
				msg = c.Ping
			case *task.Progress:
				tp = task.Communication_CommunicationProgress
				msg = c.Progress
			case *task.Notification:
				tp = task.Communication_CommunicationNotification
				msg = c.Notification
			default:
				return cb(nil, fmt.Errorf("Unknown communication type"))
		}

		if tp != c.GetType() {
			return cb(nil, fmt.Errorf("Communication type %v is not the desired type %v", c.GetType(), tp))
		}

		return cb(msg, nil)
	})
}
