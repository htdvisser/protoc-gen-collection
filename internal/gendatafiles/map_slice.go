package gendatafiles

import (
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
)

type MapItem struct {
	Key   string
	Value interface{}
}

type MapSlice []MapItem

func (m MapSlice) MarshalYAML() (interface{}, error) {
	yamlMapSlice := make(yaml.MapSlice, len(m))
	for i, v := range m {
		yamlMapSlice[i] = yaml.MapItem{
			Key:   v.Key,
			Value: v.Value,
		}
	}
	return yamlMapSlice, nil
}

var jsonEncoder = jsoniter.ConfigCompatibleWithStandardLibrary

func (m MapSlice) MarshalJSON() ([]byte, error) {
	stream := jsonEncoder.BorrowStream(nil)
	defer jsonEncoder.ReturnStream(stream)
	stream.WriteObjectStart()
	for i, v := range m {
		if i > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField(v.Key)
		stream.WriteVal(v.Value)
	}
	stream.WriteObjectEnd()
	return stream.Buffer(), nil
}

type mapSliceByKey MapSlice

func (a mapSliceByKey) Len() int           { return len(a) }
func (a mapSliceByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a mapSliceByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }
