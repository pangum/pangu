package pangu

var _ option = (*optionSubcommandHelpTemplate)(nil)

type optionSubcommandHelpTemplate struct {
	template string
}

// SubcommandHelpTemplate 配置命令帮助信息
func SubcommandHelpTemplate(template string) *optionSubcommandHelpTemplate {
	return &optionSubcommandHelpTemplate{
		template: template,
	}
}

func (b *optionSubcommandHelpTemplate) apply(options *options) {
	options.helpSubcommandTemplate = b.template
}
