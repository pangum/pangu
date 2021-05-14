package pangu

var _ option = (*optionCommandHelpTemplate)(nil)

type optionCommandHelpTemplate struct {
	template string
}

// CommandHelpTemplate 配置命令帮助信息
func CommandHelpTemplate(template string) *optionCommandHelpTemplate {
	return &optionCommandHelpTemplate{
		template: template,
	}
}

func (b *optionCommandHelpTemplate) apply(options *options) {
	options.helpCommandTemplate = b.template
}
