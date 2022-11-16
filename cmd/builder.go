package cmd

type builder struct {
	command *Command
}

func New(name string) *builder {
	return &builder{
		command: &Command{
			name: name,
		},
	}
}

func (b *builder) Build() *Command {
	return b.command
}

func (b *builder) Aliases(aliases ...string) *builder {
	b.command.aliases = aliases

	return b
}

func (b *builder) Usage(usage string) *builder {
	b.command.usage = usage

	return b
}
