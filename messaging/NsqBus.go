package messaging

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/pjvds/promise/serialization"
	"net/http"
)

type NsqBus struct {
	marshaller serialization.Marshaller
}

func NewNsqBus() *NsqBus {
	return &NsqBus{
		marshaller: serialization.NewJsonMarshaller(),
	}
}

func (bus *NsqBus) Publish(message *Message) error {
	body, err := bus.marshaller.Marshal(message)

	if err != nil {
		return errors.New("unable to marshall message: " + err.Error())
	}

	httpclient := &http.Client{}
	port := 4151
	method := "PUT"
	topic := "default"
	endpoint := fmt.Sprintf("http://127.0.0.1:%d/%s?topic=%s", port, method, topic)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	resp, err := httpclient.Do(req)
	if err != nil {
		return errors.New("unable to send request: " + err.Error())
	}
	resp.Body.Close()

	message.published = true
	return nil
}
