package pangu

type options struct {
	// 应用名称
	name string
	// 标志
	banner banner
	// 是否处理默认值
	isDefault bool
}

func defaultOptions() options {
	return options{
		isDefault: true,
	}
}
