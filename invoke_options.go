package pangu

type invokeOptions struct {
	name string
}

func defaultInvokeOptions() *invokeOptions {
	return &invokeOptions{}
}
