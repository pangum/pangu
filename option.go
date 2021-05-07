package pangu

type option interface {
	apply(options *options)
}
