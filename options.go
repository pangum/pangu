package pangu

type options struct {
	// 应用名称
	name string
	// 应用描述
	description string
	// 标志
	banner banner
	// 是否处理默认值
	isDefault bool
	// 是否验证数据
	isValidate bool
}

func defaultOptions() options {
	return options{
		isDefault:  true,
		isValidate: true,
	}
}
