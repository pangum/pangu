package argument

import (
	"github.com/pangum/pangu/internal/argument/internal"
)

var _ = New[int]

type Builder[T Type] struct {
	argument *Default[T]
}

func New[T Type](name string, target *T) *Builder[T] {
	return &Builder[T]{
		argument: NewDefault(name, target),
	}
}

func (b *Builder[T]) Build() *Default[T] {
	return b.argument
}

func (b *Builder[T]) Default(value T) (builder *Builder[T]) {
	b.argument.value = value
	builder = b

	return
}

func (b *Builder[T]) Addable() (builder *Builder[T]) {
	b.argument.addable = true
	builder = b

	return
}

func (b *Builder[T]) Aliases(aliases ...string) (builder *Builder[T]) {
	b.argument.aliases = aliases
	builder = b

	return
}

func (b *Builder[T]) Usage(usage string) (builder *Builder[T]) {
	b.argument.usage = usage
	builder = b

	return
}

func (b *Builder[T]) Environment(environment ...string) (builder *Builder[T]) {
	b.argument.environments = environment
	builder = b

	return
}

func (b *Builder[T]) Required() (builder *Builder[T]) {
	b.argument.required = true
	builder = b

	return
}

func (b *Builder[T]) Hidden() (builder *Builder[T]) {
	b.argument.hidden = true
	builder = b

	return
}

func (b *Builder[T]) Text(text string) (builder *Builder[T]) {
	b.argument.text = text
	builder = b

	return
}

func (b *Builder[T]) Action(action internal.Action[T]) (builder *Builder[T]) {
	b.argument.action = action
	builder = b

	return
}
