package davinci

func Unmarshal(bytes []byte, v any, opts ExportCmpOpts) (err error) {

	// init the decoders
	decoderCtx := NewDecoderContext(opts)
	decoderCtx.InitDecoders()

	// Decode the thing
	return decoderCtx.Decode(bytes, v)
}
