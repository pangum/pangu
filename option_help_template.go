package pangu

var _ option = (*optionHelpTemplate)(nil)

type optionHelpTemplate struct {
	template string
}

// HelpTemplate 配置应用描述
func HelpTemplate(template string) *optionHelpTemplate {
	return &optionHelpTemplate{
		template: template,
	}
}

func (b *optionHelpTemplate) apply(options *options) {
	options.helpTemplate = b.template
}
