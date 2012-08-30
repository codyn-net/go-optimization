package optimization

import (
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	task "optimization/messages/task.pb"
	"strconv"
	"io"
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
			cb(nil, err)
			break
		}

		if !cont {
			break
		}
	}
}
