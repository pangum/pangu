package pangu

type (
	provideOption interface {
		applyProvide(options *provideOptions)
	}

	provideOptions struct {
		name string
	}
)

func defaultProvideOptions() *provideOptions {
	return &provideOptions{}
}
