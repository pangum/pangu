package arg

var _ = New[int]

type builder[T argumentType] struct {
	argument *argument[T]
}

func New[T argumentType](name string, target *T) *builder[T] {
	return &builder[T]{
		argument: &argument[T]{
			name:   name,
			target: target,
		},
	}
}

func (b *builder[T]) Build() *argument[T] {
	return b.argument
}

func (b *builder[T]) Default(value T) *builder[T] {
	b.argument.value = value

	return b
}

func (b *builder[T]) Aliases(aliases ...string) *builder[T] {
	b.argument.aliases = aliases

	return b
}

func (b *builder[T]) Usage(usage string) *builder[T] {
	b.argument.usage = usage

	return b
}

func (b *builder[T]) Envs(envs ...string) *builder[T] {
	b.argument.envs = envs

	return b
}

func (b *builder[T]) Required() *builder[T] {
	b.argument.required = true

	return b
}

func (b *builder[T]) Hidden() *builder[T] {
	b.argument.hidden = true

	return b
}

func (b *builder[T]) Text(text string) *builder[T] {
	b.argument.text = text

	return b
}

func (b *builder[T]) Action(action action[T]) *builder[T] {
	b.argument.action = action

	return b
}
