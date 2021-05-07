package pangu

type options struct {
	// 应用名称
	name string
	// 标志
	banner banner
}

func defaultOptions() options {
	return options{}
}
