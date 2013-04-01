package eventing

import (
	log "code.google.com/p/log4go"
	"github.com/pjvds/promise/identification"
	"reflect"
)

type EventHandler func(event interface{}) error

type EventSource struct {
	Id       identification.Identifier
	handlers map[string]EventHandler
}

func NewEventSource() EventSource {
	source := EventSource{
		Id:       identification.NewId(),
		handlers: map[string]EventHandler{},
	}

	return source
}

type EventSourcer interface {
	Apply(e *EventMessage)
}

type EventMessage struct {
	Id []byte
}

func getEventKey(event interface{}) string {
	eType := reflect.TypeOf(event)
	return eType.PkgPath() + "/" + eType.Name()
}

func (source EventSource) RegisterHandler(event interface{}, handler EventHandler) {
	key := getEventKey(event)
	log.Trace("Registering handler for type %v", key)

	source.handlers[key] = handler
}

func (source EventSource) getHandler(key string) EventHandler {
	log.Trace("Getting handler for key %v.", key)

	handler := source.handlers[key]

	if handler == nil {
		log.Warn("No handler found.")
	}

	return handler
}

func (source EventSource) Apply(e interface{}) {
	key := getEventKey(e)

	log.Trace("Invoking handler now.")
	handler := source.getHandler(key)
	err := handler(e)

	if err != nil {
		log.Warn("Handler returned error: %v", err.Error())
	} else {
		log.Trace("Handler executed successfully")
	}

	// refl := reflect.ValueOf(source)
	// methNum := refl.NumMethod()

	// log.Trace("found %v methods on %v", methNum, refl.Type())

	// for i := 0; i < methNum; i++ {
	// 	method := refl.Method(i)
	// 	methodInfo := method.Type()

	// 	log.Trace("[%v/%v] looking at method %v.", i+1, methNum, methodInfo.String())

	// 	if methodInfo.NumIn() != 1 {
	// 		log.Trace("skipping method because it has %v arguments", methodInfo.NumIn())
	// 		continue
	// 	}

	// 	if methodInfo.In(0).Kind() == reflect.TypeOf(e).Kind() {
	// 		args := make([]reflect.Value, 1)
	// 		args[0] = reflect.ValueOf(e)
	// 		///method.Call(args)

	// 		log.Trace("event has been applied")
	// 	}
	// }
}
