package core

import (
	"context"

	"github.com/harluo/boot/internal/internal/kernel"
)

type Typer struct {
	value any
}

func NewTyper(value any) *Typer {
	return &Typer{
		value: value,
	}
}

func (t *Typer) Usage() (usage string) {
	if convert, ok := t.value.(kernel.Usage); ok {
		usage = convert.Usage()
	}

	return
}

func (t *Typer) Aliases() (aliases []string) {
	if convert, ok := t.value.(kernel.Aliases); ok {
		aliases = convert.Aliases()
	}

	return
}

func (t *Typer) Text() (text string) {
	if convert, ok := t.value.(kernel.Text); ok {
		text = convert.Text()
	}

	return
}

func (t *Typer) Description() (description string) {
	if convert, ok := t.value.(kernel.Description); ok {
		description = convert.Description()
	}

	return
}

func (t *Typer) Category() (category string) {
	if convert, ok := t.value.(kernel.Category); ok {
		category = convert.Category()
	}

	return
}

func (t *Typer) Required() (required bool) {
	if convert, ok := t.value.(kernel.Required); ok {
		required = convert.Required()
	}

	return
}

func (t *Typer) Hidden() (hidden bool) {
	if convert, ok := t.value.(kernel.Hidden); ok {
		hidden = convert.Hidden()
	}

	return
}

func (t *Typer) Addable() (addable bool) {
	if convert, ok := t.value.(kernel.Addable); ok {
		addable = convert.Addable()
	}

	return
}

func (t *Typer) Default() (defaults any) {
	if convert, ok := t.value.(kernel.Default); ok {
		defaults = convert.Default()
	}

	return
}

func (t *Typer) Environments() (environments []string) {
	if convert, ok := t.value.(kernel.Environments); ok {
		environments = convert.Environments()
	}

	return
}

func (t *Typer) Run() (run func(context.Context) error) {
	if convert, ok := t.value.(kernel.Run); ok {
		run = convert.Run
	}

	return
}
