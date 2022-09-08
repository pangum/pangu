package arg

var (
	_        = Hidden
	_ option = (*optionHidden)(nil)
)

type optionHidden struct{}

func Hidden() *optionHidden {
	return new(optionHidden)
}

func (h *optionHidden) apply(options *options) {
	options.required = true
}
