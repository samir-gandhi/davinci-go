package davinci

import (
	"fmt"
	"reflect"
)

var (
	_ CodecContext = &EncoderContext{}
)

type EncoderContext struct {
	typeEncoders map[reflect.Type]ValueEncoder
	KindEncoders map[reflect.Kind]ValueEncoder
	encoders     map[string]ValueEncoder
	Opts         ExportCmpOpts
}

func NewEncoderContext(opts ExportCmpOpts) *EncoderContext {
	return &EncoderContext{
		Opts: opts,
	}
}

func (o *EncoderContext) InitEncoders() {

	o.encoders = map[string]ValueEncoder{
		"json":   NewJSONEncoder(),
		"map":    NewMapEncoder(o),
		"ptr":    NewPtrEncoder(o),
		"slice":  NewSliceEncoder(o),
		"struct": NewStructEncoder(o),
	}

	o.typeEncoders = map[reflect.Type]ValueEncoder{
		reflect.TypeOf(EpochTime{}): o.encoders["json"],
	}

	o.KindEncoders = map[reflect.Kind]ValueEncoder{
		reflect.Ptr:       o.encoders["ptr"],
		reflect.Struct:    o.encoders["struct"],
		reflect.Slice:     o.encoders["slice"],
		reflect.Map:       o.encoders["map"],
		reflect.Interface: o.encoders["json"],
		reflect.String:    o.encoders["json"],
		reflect.Bool:      o.encoders["json"],
		reflect.Int:       o.encoders["json"],
		reflect.Int8:      o.encoders["json"],
		reflect.Int16:     o.encoders["json"],
		reflect.Int32:     o.encoders["json"],
		reflect.Int64:     o.encoders["json"],
		reflect.Uint:      o.encoders["json"],
		reflect.Uint8:     o.encoders["json"],
		reflect.Uint16:    o.encoders["json"],
		reflect.Uint32:    o.encoders["json"],
		reflect.Uint64:    o.encoders["json"],
		reflect.Float32:   o.encoders["json"],
		reflect.Float64:   o.encoders["json"],
	}
}

func (o EncoderContext) Encode(v any) (bytes []byte, err error) {

	// Check if there's a custom implementation and run it if there is
	if marshaler, ok := v.(Marshaler); ok {
		return marshaler.MarshalDavinci(o.Opts)
	}

	rtyp := reflect.TypeOf(v)

	if !reflect.ValueOf(v).IsValid() {
		return []byte("null"), nil
	}

	// Lookup the relevant encoder for the value type
	encoder, err := o.GetEncoder(rtyp)
	if err != nil {
		return nil, err
	}

	// Decode the value with the encoder
	return encoder.EncodeValue(reflect.ValueOf(v))
}

func (o EncoderContext) GetEncoder(valueType reflect.Type) (ValueEncoder, error) {
	if valueType == nil {
		return nil, fmt.Errorf("nil value type")
	}

	if encoder, ok := o.LookupEncoderByType(valueType); ok {
		if encoder == nil {
			return nil, fmt.Errorf("no encoder found for type %s", valueType)
		}
		return encoder, nil
	}

	if encoder, ok := o.LookupEncoderByKind(valueType.Kind()); ok {
		if encoder == nil {
			return nil, fmt.Errorf("no encoder found for kind %s", valueType)
		}
		return encoder, nil
	}

	return nil, fmt.Errorf("no encoder lookup implementation for type %s", valueType)
}

func (o EncoderContext) LookupEncoderByType(valueType reflect.Type) (ValueEncoder, bool) {

	if o.typeEncoders != nil && o.typeEncoders[valueType] != nil {
		return o.typeEncoders[valueType], true
	}

	return nil, false
}

func (o EncoderContext) LookupEncoderByKind(valueKind reflect.Kind) (ValueEncoder, bool) {

	if o.KindEncoders != nil && o.KindEncoders[valueKind] != nil {
		return o.KindEncoders[valueKind], true
	}

	return nil, false
}

func (o EncoderContext) GetOpts() ExportCmpOpts {
	return o.Opts
}
