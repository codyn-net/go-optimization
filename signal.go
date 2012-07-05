package optimization

import (
	"fmt"
	"ponyo.epfl.ch/git/optimization/go/optimization/log"
	"reflect"
	"runtime"
)

type sigType struct {
	Type  reflect.Type
	Value reflect.Value
}

type Signal struct {
	signature reflect.Value
	funcs     []sigType
	sync      bool
}

func newSignal(sync bool, sig interface{}) *Signal {
	v := reflect.ValueOf(sig)
	t := v.Type()

	if t.Kind() != reflect.Func {
		return nil
	}

	return &Signal{
		signature: v,
		sync:      sync,
		funcs:     make([]sigType, 0),
	}
}

func NewSignal(sig interface{}) *Signal {
	return newSignal(false, sig)
}

func NewSyncSignal(sig interface{}) *Signal {
	return newSignal(true, sig)
}

func (s *Signal) Connect(f interface{}) {
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)

	if t.Kind() != reflect.Func {
		_, fname, line, _ := runtime.Caller(1)

		log.S("%s:%d: Could not connect `%v', not a function", fname, line, t)
		return
	}

	// check args
	st := s.signature.Type()

	if t.NumIn() != st.NumIn() {
		_, fname, line, _ := runtime.Caller(1)

		log.S("%s:%d: Could not connect `%v', incorrect number of arguments (%d but expected %d)",
			fname,
			line,
			t,
			t.NumIn(),
			st.NumIn())

		return
	}

	for i := 0; i < t.NumIn(); i++ {
		k := t.In(i).Kind()
		kst := st.In(i).Kind()

		if k == reflect.Interface || kst == reflect.Interface {
			continue
		}

		if t.In(i).Kind() != st.In(i).Kind() {
			_, fname, line, _ := runtime.Caller(1)

			log.S("%s:%d: Could not connect `%v', incorrect argument %d type: %v but expected %v",
				fname,
				line,
				t,
				i+1,
				t.In(i).Kind(),
				st.In(i).Kind())

			return
		}
	}

	s.funcs = append(s.funcs, sigType{
		Type:  t,
		Value: v,
	})
}

func (s *Signal) Emit(args ...interface{}) {
	s.emit(s.sync, args...)
}

func (s *Signal) EmitSync(args ...interface{}) {
	s.emit(true, args...)
}

func (s *Signal) EmitAsync(args ...interface{}) {
	s.emit(false, args...)
}

func (s *Signal) emit(sync bool, args ...interface{}) {
	ot := s.signature.Type()

	if ot.NumIn() != len(args) {
		_, fname, line, _ := runtime.Caller(2)

		log.S("%s:%d: Could not emit signal, incorrect number of arguments (%d but expected %d)",
			fname,
			line,
			len(args),
			ot.NumIn())

		return
	}

	// convert args to values
	vals := make([]reflect.Value, len(args))

	for i, v := range args {
		vals[i] = reflect.ValueOf(v)

		if !vals[i].Type().AssignableTo(ot.In(i)) {
			return
		}
	}

	f := func() {
		for _, r := range s.funcs {
			rval := r.Value.Call(vals)

			for _, v := range rval {
				fmt.Println(v)
			}
		}
	}

	if sync {
		Events <- f
	} else {
		f()
	}
}
