package arg

type (
	option interface {
		apply(options *options)
	}

	options struct {
		_default    any
		destination interface{}
		aliases     []string
		usage       string
		envs        []string
		required    bool
		hidden      bool
		dt          string
	}
)

func defaultOption() *options {
	return &options{
		required: false,
		hidden:   false,
	}
}
