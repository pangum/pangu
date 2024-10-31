package command

type Builder struct {
	base *Default
}

func (b *Builder) Build() *Default {
	return b.base
}

func (b *Builder) Aliases(aliases ...string) (builder *Builder) {
	b.base.aliases = append(b.base.aliases, aliases...)
	builder = b

	return
}

func (b *Builder) Usage(usage string) (builder *Builder) {
	b.base.usage = usage
	builder = b

	return
}

func (b *Builder) Category(category string) (builder *Builder) {
	b.base.category = category
	builder = b

	return
}

func (b *Builder) Hidden() (builder *Builder) {
	b.base.hidden = true
	builder = b

	return
}
