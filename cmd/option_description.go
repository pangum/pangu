package cmd

var (
	_        = Description
	_ option = (*optionDescription)(nil)
)

type optionDescription struct {
	description string
}

// Description 命令描述
func Description(description string) *optionDescription {
	return &optionDescription{
		description: description,
	}
}

func (u *optionDescription) apply(options *options) {
	options.description = u.description
}
