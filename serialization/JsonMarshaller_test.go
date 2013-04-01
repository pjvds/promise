package serialization

import (
	"testing"
)

type message struct {
	Name  string
	Value int
}

func TestJsonMarshallerIsMarshaller(t *testing.T) {
	m := NewJsonMarshaller()
	i := interface{}(m)

	if _, ok := i.(Marshaller); !ok {
		t.Error("JsonMarshaller doesn't implement interface Marshaller")
	}
}

func TestMarshallUnformatted(t *testing.T) {
	marshaller := getUnformattedMarshaller()
	message := message{
		Name:  "Pieter Joost",
		Value: 42,
	}

	data, _ := marshaller.Marshal(message)
	json := string(data)

	expected := "{\"Name\":\"Pieter Joost\",\"Value\":42}"
	if json != expected {
		t.Errorf("invalid json output: '%v' should be '%v", json, expected)
	}
}

func TestMarshallFormatted(t *testing.T) {
	marshaller := getFormattedMarshaller()
	message := message{
		Name:  "Pieter Joost",
		Value: 42,
	}

	data, _ := marshaller.Marshal(message)
	json := string(data)

	expected := "{\n\t\"Name\": \"Pieter Joost\",\n\t\"Value\": 42\n}"
	if json != expected {
		t.Errorf("invalid json output: '%v' should be '%v", json, expected)
	}
}

func TestUnMarshall(t *testing.T) {
	marshaller := getUnformattedMarshaller()
	json := []byte("{\"Name\":\"Pieter Joost\",\"Value\":42}")
	var message message

	err := marshaller.Unmarshal(json, &message)

	if err != nil {
		t.Errorf("error while unmarshalling: %v", err.Error())
	}

	expectedName := "Pieter Joost"
	if message.Name != expectedName {
		t.Errorf("field name didn't unmarshal correctly: %v should be %v", message.Name, expectedName)
	}

	expectedNumber := 42
	if message.Value != expectedNumber {
		t.Errorf("field value didn't unmarshal correctly: %v should be %v", message.Value, expectedNumber)
	}
}

func getUnformattedMarshaller() JsonMarshaller {
	m := NewJsonMarshaller()
	m.Format = false

	return m
}

func getFormattedMarshaller() JsonMarshaller {
	m := NewJsonMarshaller()
	m.Format = true

	return m
}
