package pangu

type (
	configOption interface {
		applyConfig(options *configOptions)
	}

	configOptions struct {
		typ configType
	}
)

func defaultConfigOptions() *configOptions {
	return &configOptions{
		typ: ConfigTypeAppend,
	}
}
