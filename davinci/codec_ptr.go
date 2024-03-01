package davinci

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var _ ValueEncoder = &PtrDecoder{}
var _ ValueDecoder = &PtrDecoder{}

type PtrDecoder struct {
	dCtx *DecoderContext
}

func NewPtrDecoder(dCtx *DecoderContext) PtrDecoder {
	return PtrDecoder{
		dCtx: dCtx,
	}
}

func (PtrDecoder) String() string {
	return "davinci.PtrDecoder"
}

func (d PtrDecoder) DecodeValue(data []byte, v reflect.Value) error {
	if !v.IsValid() || !v.CanSet() || v.Kind() != reflect.Ptr {
		return fmt.Errorf("invalid pointer value to decode")
	}

	typ := v.Type()

	if v.IsNil() {
		v.Set(reflect.New(typ.Elem()))
	}

	// Decode the value into the struct field
	return d.dCtx.Decode(data, v.Elem().Addr().Interface())
}

func (d PtrDecoder) EncodeValue(data interface{}, v reflect.Value) error {
	return json.Unmarshal(data.([]byte), v.Addr().Interface())
}
