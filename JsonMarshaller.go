package main

import (
	"bytes"
	"encoding/json"
)

type JsonMarshaller struct {
	Format bool
}

func NewJsonMarshaller() JsonMarshaller {
	return JsonMarshaller{
		Format: true,
	}
}

type Marshaller interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

func (m JsonMarshaller) Marshal(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)

	if err == nil {
		var buffer bytes.Buffer
		err = json.Indent(&buffer, data, "", "\t")
		data = buffer.Bytes()
	}

	return data, err
}

func (m JsonMarshaller) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
