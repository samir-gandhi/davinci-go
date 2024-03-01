package davinci

import (
	"fmt"
	"reflect"
)

var _ ValueEncoder = &PtrCodec{}
var _ ValueDecoder = &PtrCodec{}

type PtrCodec struct {
	dCtx *DecoderContext
	eCtx *EncoderContext
}

func NewPtrDecoder(dCtx *DecoderContext) ValueDecoder {
	return PtrCodec{
		dCtx: dCtx,
	}
}

func NewPtrEncoder(eCtx *EncoderContext) ValueEncoder {
	return PtrCodec{
		eCtx: eCtx,
	}
}

func (PtrCodec) String() string {
	return "davinci.PtrCodec"
}

func (d PtrCodec) DecodeValue(data []byte, v reflect.Value) error {
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

func (d PtrCodec) EncodeValue(v reflect.Value) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}
