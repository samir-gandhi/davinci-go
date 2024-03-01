package davinci

import (
	"encoding/json"
	"reflect"
)

var _ ValueEncoder = &JSONDecoder{}
var _ ValueDecoder = &JSONDecoder{}

type JSONDecoder struct{}

func NewJSONDecoder() JSONDecoder {
	return JSONDecoder{}
}

func (JSONDecoder) String() string {
	return "davinci.JSONDecoder"
}

func (d JSONDecoder) DecodeValue(data []byte, v reflect.Value) error {
	return json.Unmarshal(data, v.Addr().Interface())
}

func (d JSONDecoder) EncodeValue(data interface{}, v reflect.Value) error {
	return json.Unmarshal(data.([]byte), v.Addr().Interface())
}
