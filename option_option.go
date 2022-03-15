package pangu

type (
	optionOption interface {
		applyOption(options *optionOptions)
	}

	optionOptions struct {
		typ optionType
	}
)

func defaultOptionOptions() *optionOptions {
	return &optionOptions{
		typ: ConfigTypeAppend,
	}
}
