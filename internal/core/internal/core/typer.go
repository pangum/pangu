package core

import (
	"context"

	"github.com/harluo/boot/internal/application"
	"github.com/harluo/boot/internal/internal/checker"
)

type Typer struct {
	value any
}

func NewTyper(value any) *Typer {
	return &Typer{
		value: value,
	}
}

func (t *Typer) Usage() (value string) {
	if convert, ok := t.value.(checker.Usage); ok {
		value = convert.Usage()
	}

	return
}

func (t *Typer) Aliases() (value []string) {
	if convert, ok := t.value.(checker.Aliases); ok {
		value = convert.Aliases()
	}

	return
}

func (t *Typer) Text() (value string) {
	if convert, ok := t.value.(checker.Text); ok {
		value = convert.Text()
	}

	return
}

func (t *Typer) Description() (value string) {
	if description, dlk := t.value.(checker.Description); dlk {
		value = description.Description()
	} else if desc, dsk := t.value.(checker.Desc); dsk {
		value = desc.Desc()
	}

	return
}

func (t *Typer) Category() (value string) {
	if convert, ok := t.value.(checker.Category); ok {
		value = convert.Category()
	}

	return
}

func (t *Typer) Required() (value bool) {
	if convert, ok := t.value.(checker.Required); ok {
		value = convert.Required()
	}

	return
}

func (t *Typer) Hidden() (value bool) {
	if convert, ok := t.value.(checker.Hidden); ok {
		value = convert.Hidden()
	}

	return
}

func (t *Typer) Addable() (value bool) {
	if convert, ok := t.value.(checker.Addable); ok {
		value = convert.Addable()
	}

	return
}

func (t *Typer) Default(callback func(value any)) {
	if convert, ok := t.value.(checker.Default); ok {
		callback(convert.Default())
	}

	return
}

func (t *Typer) Environments() (value []string) {
	if environments, elk := t.value.(checker.Environments); elk {
		value = environments.Environments()
	} else if envs, esk := t.value.(checker.Envs); esk {
		value = envs.Envs()
	}

	return
}

func (t *Typer) Arguments() (value []application.Argument) {
	if arguments, alk := t.value.(checker.Arguments); alk {
		value = arguments.Arguments()
	}

	return
}

func (t *Typer) Commands() (value []application.Command) {
	if commands, ck := t.value.(checker.Commands); ck {
		value = commands.Subcommands()
	}

	return
}

func (t *Typer) Subcommands() (value []application.Command) {
	if subcommands, sk := t.value.(checker.Subcommands); sk {
		value = subcommands.Subcommands()
	}

	return
}

func (t *Typer) Run(ctx context.Context) (err error) {
	if convert, ok := t.value.(checker.Run); ok {
		err = convert.Run(ctx)
	}

	return
}
