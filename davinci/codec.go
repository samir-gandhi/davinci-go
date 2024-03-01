package davinci

import "reflect"

type ValueDecoder interface {
	DecodeValue([]byte, reflect.Value) error
}

type ValueEncoder interface {
	EncodeValue(interface{}, reflect.Value) error
}

type Unmarshaler interface {
	UnmarshalDavinci([]byte, ExportCmpOpts) error
}
