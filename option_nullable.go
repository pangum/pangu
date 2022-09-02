package pangu

var (
	_              = Nullable
	_              = Required
	_ option       = (*optionNullable)(nil)
	_ configOption = (*optionNullable)(nil)
)

type optionNullable struct {
	nullable bool
}

// Nullable 是否可空
func Nullable() *optionNullable {
	return nullable(true)
}

// Required 不允许为空
func Required() *optionNullable {
	return nullable(false)
}

func nullable(nullable bool) *optionNullable {
	return &optionNullable{
		nullable: nullable,
	}
}

func (n *optionNullable) apply(options *options) {
	options.nullable = n.nullable
}

func (n *optionNullable) applyConfig(options *configOptions) {
	options.nullable = n.nullable
}
