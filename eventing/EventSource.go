package eventing

import (
	"github.com/pjvds/promise/identification"
	"reflect"
)

type eventSource struct {
	Id identification.Identifier
}

type EventSource interface {
	Apply(e *EventMessage)
}

type EventMessage struct {
	Id []byte
}

func (source eventSource) Apply(e *EventMessage) {
	refl := reflect.ValueOf(&source)
	methNum := refl.NumMethod()

	for i := 0; i < methNum; i++ {
		meth := refl.Method(i)

		if meth.NumField() != 0 {
			continue
		}

		if meth.Field(0).Type() == reflect.TypeOf(e) {
			args := make([]reflect.Value, 1)
			args[0] = reflect.ValueOf(e)
			meth.Call(args)
		}
	}
}
