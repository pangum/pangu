package arg

var (
	_        = Aliases
	_ option = (*optionAliases)(nil)
)

type optionAliases struct {
	aliases []string
}

// Aliases 别名列表
func Aliases(aliases ...string) *optionAliases {
	return &optionAliases{
		aliases: aliases,
	}
}

func (a *optionAliases) apply(options *options) {
	options.aliases = a.aliases
}
