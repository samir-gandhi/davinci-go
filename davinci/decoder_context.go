package davinci

import (
	"fmt"
	"reflect"
)

var (
	_ CodecContext = &DecoderContext{}
)

type DecoderContext struct {
	typeDecoders map[reflect.Type]ValueDecoder
	KindDecoders map[reflect.Kind]ValueDecoder
	decoders     map[string]ValueDecoder
	Opts         ExportCmpOpts
}

func NewDecoderContext(opts ExportCmpOpts) *DecoderContext {
	return &DecoderContext{
		Opts: opts,
	}
}

func (o *DecoderContext) InitDecoders() {

	o.decoders = map[string]ValueDecoder{
		"slice":  NewSliceDecoder(o),
		"ptr":    NewPtrDecoder(o),
		"struct": NewStructDecoder(o),
		"json":   NewJSONDecoder(),
	}

	o.typeDecoders = map[reflect.Type]ValueDecoder{
		reflect.TypeOf(EpochTime{}): o.decoders["json"],
	}

	o.KindDecoders = map[reflect.Kind]ValueDecoder{
		reflect.Ptr:       o.decoders["ptr"],
		reflect.Struct:    o.decoders["struct"],
		reflect.Slice:     o.decoders["slice"],
		reflect.Interface: o.decoders["json"],
		reflect.String:    o.decoders["json"],
		reflect.Bool:      o.decoders["json"],
		reflect.Int:       o.decoders["json"],
		reflect.Int8:      o.decoders["json"],
		reflect.Int16:     o.decoders["json"],
		reflect.Int32:     o.decoders["json"],
		reflect.Int64:     o.decoders["json"],
		reflect.Uint:      o.decoders["json"],
		reflect.Uint8:     o.decoders["json"],
		reflect.Uint16:    o.decoders["json"],
		reflect.Uint32:    o.decoders["json"],
		reflect.Uint64:    o.decoders["json"],
		reflect.Float32:   o.decoders["json"],
		reflect.Float64:   o.decoders["json"],
	}
}

func (o DecoderContext) Decode(bytes []byte, v any) (err error) {
	// Check if there's a custom implementation and run it if there is
	if unmarshaler, ok := v.(Unmarshaler); ok {
		return unmarshaler.UnmarshalDavinci(bytes, o.Opts)
	}

	// Check if the value to unmarshal to is a pointer or a map
	rval := reflect.ValueOf(v)
	switch rval.Kind() {
	case reflect.Ptr:
		if rval.IsNil() {
			return fmt.Errorf("Cannot decode to nil pointer")
		}
		rval = rval.Elem()
	case reflect.Map:
		if rval.IsNil() {
			return fmt.Errorf("Cannot decode to nil map")
		}
	default:
		return fmt.Errorf("The object to unmarshal to must be a pointer or a map. Got: %v", rval)
	}

	// Lookup the relevant decoder for the value type
	decoder, err := o.GetDecoder(rval.Type())
	if err != nil {
		return err
	}

	// Decode the value with the decoder
	return decoder.DecodeValue(bytes, rval)
}

func (o DecoderContext) GetDecoder(valueType reflect.Type) (ValueDecoder, error) {
	if valueType == nil {
		return nil, fmt.Errorf("nil value type")
	}

	if decoder, ok := o.LookupDecoderByType(valueType); ok {
		if decoder == nil {
			return nil, fmt.Errorf("no decoder found for type %s", valueType)
		}
		return decoder, nil
	}

	if decoder, ok := o.LookupDecoderByKind(valueType.Kind()); ok {
		if decoder == nil {
			return nil, fmt.Errorf("no decoder found for kind %s", valueType)
		}
		return decoder, nil
	}

	return nil, fmt.Errorf("no decoder lookup implementation for type %s", valueType)
}

func (o DecoderContext) LookupDecoderByType(valueType reflect.Type) (ValueDecoder, bool) {

	if o.typeDecoders != nil && o.typeDecoders[valueType] != nil {
		return o.typeDecoders[valueType], true
	}

	return nil, false
}

func (o DecoderContext) LookupDecoderByKind(valueKind reflect.Kind) (ValueDecoder, bool) {

	if o.KindDecoders != nil && o.KindDecoders[valueKind] != nil {
		return o.KindDecoders[valueKind], true
	}

	return nil, false
}

func (o DecoderContext) GetOpts() ExportCmpOpts {
	return o.Opts
}
