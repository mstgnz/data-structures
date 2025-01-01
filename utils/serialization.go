package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Serializer interface defines methods for serialization and deserialization
type Serializer interface {
	Serialize(v interface{}) ([]byte, error)
	Deserialize(data []byte, v interface{}) error
}

// JSONSerializer implements JSON serialization
type JSONSerializer struct {
	PrettyPrint bool
}

func (s JSONSerializer) Serialize(v interface{}) ([]byte, error) {
	if s.PrettyPrint {
		return json.MarshalIndent(v, "", "  ")
	}
	return json.Marshal(v)
}

func (s JSONSerializer) Deserialize(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// XMLSerializer implements XML serialization
type XMLSerializer struct {
	PrettyPrint bool
}

func (s XMLSerializer) Serialize(v interface{}) ([]byte, error) {
	if s.PrettyPrint {
		return xml.MarshalIndent(v, "", "  ")
	}
	return xml.Marshal(v)
}

func (s XMLSerializer) Deserialize(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

// GobSerializer implements Gob serialization
type GobSerializer struct{}

func (s GobSerializer) Serialize(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s GobSerializer) Deserialize(data []byte, v interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(v)
}

// SerializationFormat represents supported serialization formats
type SerializationFormat string

const (
	FormatJSON SerializationFormat = "json"
	FormatXML  SerializationFormat = "xml"
	FormatGOB  SerializationFormat = "gob"
)

// GetSerializer returns a serializer for the specified format
func GetSerializer(format SerializationFormat, prettyPrint bool) (Serializer, error) {
	switch format {
	case FormatJSON:
		return JSONSerializer{PrettyPrint: prettyPrint}, nil
	case FormatXML:
		return XMLSerializer{PrettyPrint: prettyPrint}, nil
	case FormatGOB:
		return GobSerializer{}, nil
	default:
		return nil, fmt.Errorf("unsupported serialization format: %s", format)
	}
}

// SerializeToFile serializes data to a string in the specified format
func SerializeToString(v interface{}, format SerializationFormat, prettyPrint bool) (string, error) {
	serializer, err := GetSerializer(format, prettyPrint)
	if err != nil {
		return "", err
	}

	data, err := serializer.Serialize(v)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// DeserializeFromString deserializes data from a string in the specified format
func DeserializeFromString(data string, v interface{}, format SerializationFormat) error {
	serializer, err := GetSerializer(format, false)
	if err != nil {
		return err
	}

	return serializer.Deserialize([]byte(data), v)
}

// Example usage types
type ExampleStruct struct {
	ID       int                    `json:"id" xml:"id"`
	Name     string                 `json:"name" xml:"name"`
	Tags     []string               `json:"tags" xml:"tags"`
	Metadata map[string]interface{} `json:"metadata" xml:"metadata"`
}

// RegisterTypes registers types for Gob serialization
func RegisterTypes() {
	gob.Register(ExampleStruct{})
	gob.Register(map[string]interface{}{})
	gob.Register([]interface{}{})
}

// SerializationHelper provides helper methods for common serialization tasks
type SerializationHelper struct {
	Format      SerializationFormat
	PrettyPrint bool
}

func NewSerializationHelper(format SerializationFormat, prettyPrint bool) *SerializationHelper {
	return &SerializationHelper{
		Format:      format,
		PrettyPrint: prettyPrint,
	}
}

func (h *SerializationHelper) Serialize(v interface{}) (string, error) {
	return SerializeToString(v, h.Format, h.PrettyPrint)
}

func (h *SerializationHelper) Deserialize(data string, v interface{}) error {
	return DeserializeFromString(data, v, h.Format)
}

// DeepCopy creates a deep copy of a value using serialization
func DeepCopy(src, dst interface{}) error {
	serializer := GobSerializer{}

	// Serialize the source value
	data, err := serializer.Serialize(src)
	if err != nil {
		return fmt.Errorf("serialization error: %v", err)
	}

	// Deserialize into the destination value
	if err := serializer.Deserialize(data, dst); err != nil {
		return fmt.Errorf("deserialization error: %v", err)
	}

	return nil
}
