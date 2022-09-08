package arg

var (
	_        = Required
	_ option = (*optionRequired)(nil)
)

type optionRequired struct{}

func Required() *optionRequired {
	return new(optionRequired)
}

func (r *optionRequired) apply(options *options) {
	options.required = true
}
