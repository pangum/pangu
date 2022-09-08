package arg

type (
	option interface {
		apply(options *options)
	}

	options struct {
		aliases     []string
		usage       string
		envs        []string
		required    bool
		hidden      bool
		defaultText string
	}
)
