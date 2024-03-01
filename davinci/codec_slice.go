package davinci

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var _ ValueEncoder = &SliceDecoder{}
var _ ValueDecoder = &SliceDecoder{}

type SliceDecoder struct {
	dCtx *DecoderContext
}

func NewSliceDecoder(dCtx *DecoderContext) SliceDecoder {
	return SliceDecoder{
		dCtx: dCtx,
	}
}

func (SliceDecoder) String() string {
	return "davinci.SliceDecoder"
}

func (d SliceDecoder) DecodeValue(data []byte, v reflect.Value) error {
	if !v.IsValid() || !v.CanSet() || v.Kind() != reflect.Slice {
		return fmt.Errorf("invalid slice value to decode")
	}

	// Get the type of the slice
	typ := v.Type().Elem()

	// Unmarshal the data into a slice
	var tempSlice []interface{}
	if err := json.Unmarshal(data, &tempSlice); err != nil {
		return err
	}

	for _, item := range tempSlice {
		// Create a new value of the slice element type
		elem := reflect.New(typ).Elem()

		// Convert the map value to a JSON byte slice
		jsonValueBytes, err := json.Marshal(item)
		if err != nil {
			return err
		}

		// Decode the value into the new element
		if err := d.dCtx.Decode(jsonValueBytes, elem.Addr().Interface()); err != nil {
			return err
		}

		// Append the new element to the slice
		v.Set(reflect.Append(v, elem))
	}

	return nil
}

func (d SliceDecoder) EncodeValue(data interface{}, v reflect.Value) error {
	return json.Unmarshal(data.([]byte), v.Addr().Interface())
}
