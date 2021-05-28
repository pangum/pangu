package pangu

type provideOption interface {
	applyProvide(options *provideOptions)
}
