package davinci

import (
	"encoding/json"
	"reflect"
)

var _ ValueEncoder = &JSONCodec{}
var _ ValueDecoder = &JSONCodec{}

type JSONCodec struct{}

func NewJSONDecoder() ValueDecoder {
	return JSONCodec{}
}

func NewJSONEncoder() ValueEncoder {
	return JSONCodec{}
}

func (JSONCodec) String() string {
	return "davinci.JSONCodec"
}

func (d JSONCodec) DecodeValue(data []byte, v reflect.Value) error {
	return json.Unmarshal(data, v.Addr().Interface())
}

func (d JSONCodec) EncodeValue(v reflect.Value) ([]byte, error) {
	return json.Marshal(v.Interface())
}
