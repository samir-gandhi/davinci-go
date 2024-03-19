package davinci

import "reflect"

type ValueDecoder interface {
	DecodeValue([]byte, reflect.Value) error
}

type ValueEncoder interface {
	EncodeValue(reflect.Value) ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalDavinci([]byte, ExportCmpOpts) error
}

type Marshaler interface {
	MarshalDavinci(ExportCmpOpts) ([]byte, error)
}

type CodecContext interface {
	GetOpts() ExportCmpOpts
}
