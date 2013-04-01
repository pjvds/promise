package serialization

import (
	"bytes"
	log "code.google.com/p/log4go"
	"encoding/json"
)

type JsonMarshaller struct {
	Format bool
}

func NewJsonMarshaller() *JsonMarshaller {
	return &JsonMarshaller{
		Format: true,
	}
}

func (m JsonMarshaller) Marshal(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)

	if err == nil && m.Format {
		var buffer bytes.Buffer
		err = json.Indent(&buffer, data, "", "\t")
		data = buffer.Bytes()
	}

	if data != nil {
		log.Debug("marshalled to json: ", string(data))
	}
	return data, err
}

func (m JsonMarshaller) Unmarshal(data []byte, v interface{}) error {
	log.Debug("unmarshalling json: '%v'", string(data))

	return json.Unmarshal(data, v)
}
