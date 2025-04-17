package command

import (
	"context"

	"github.com/harluo/boot/internal/application"
	"github.com/harluo/boot/internal/internal/param"
)

// Default 命令基础
// 可以通过匿名继承的方式快速创建命令
type Default struct {
	params *param.Command
}

func NewDefault(params *param.Command) *Default {
	return &Default{
		params: params,
	}
}

func (d *Default) Name() string {
	return d.params.Name
}

func (d *Default) Aliases() []string {
	return d.params.Aliases
}

func (d *Default) Usage() string {
	return d.params.Usage
}

func (d *Default) Arguments() (arguments application.Arguments) {
	return
}

func (d *Default) Subcommands() (subcommands application.Commands) {
	return
}

func (d *Default) Description() string {
	return d.params.Description
}

func (d *Default) Category() string {
	return d.params.Category
}

func (d *Default) Hidden() bool {
	return d.params.Hidden
}

func (d *Default) Before(_ context.Context) (err error) {
	return
}

func (d *Default) Stop(_ context.Context) (err error) {
	return
}

func (d *Default) After(_ context.Context) (err error) {
	return
}
