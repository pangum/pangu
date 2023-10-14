package command

type Builder struct {
	base *Base
}

func (b *Builder) Build() *Base {
	return b.base
}

func (b *Builder) Aliases(aliases ...string) *Builder {
	b.base.aliases = aliases

	return b
}

func (b *Builder) Usage(usage string) *Builder {
	b.base.usage = usage

	return b
}

func (b *Builder) Category(category string) *Builder {
	b.base.category = category

	return b
}

func (b *Builder) Hidden() *Builder {
	b.base.hidden = true

	return b
}
