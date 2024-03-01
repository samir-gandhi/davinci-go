package davinci

func Marshal(v any, opts ExportCmpOpts) (bytes []byte, err error) {

	// init the encoders
	encoderCtx := NewEncoderContext(opts)
	encoderCtx.InitEncoders()

	// Encode the thing
	return encoderCtx.Encode(v)
}
