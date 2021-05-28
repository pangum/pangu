package pangu

type invokeOption interface {
	applyInvoke(options *invokeOptions)
}
