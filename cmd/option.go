package cmd

type (
	option interface {
		apply(options *options)
	}

	options struct {
		aliases     []string
		usage       string
		description string
	}
)

func defaultOption() *options {
	return new(options)
}
