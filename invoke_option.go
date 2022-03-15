package pangu

type (
	invokeOption interface {
		applyInvoke(options *invokeOptions)
	}

	invokeOptions struct {
		name string
	}
)

func defaultInvokeOptions() *invokeOptions {
	return &invokeOptions{}
}
