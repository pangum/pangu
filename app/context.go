package app

import (
	"context"

	"github.com/urfave/cli/v2"
)

// Context 描述上下文
type Context struct {
	context.Context

	context *cli.Context
}

// NewContext 创建上下文
func NewContext(ctx *cli.Context) *Context {
	return &Context{
		Context: context.Background(),

		context: ctx,
	}
}

func (c *Context) String(name string) string {
	return c.context.String(name)
}

func (c *Context) Set(name string, value string) error {
	return c.context.Set(name, value)
}

func (c *Context) Value(key any) (value any) {
	if _key, ok := key.(string); ok {
		value = c.context.Value(_key)
	}
	if nil == value {
		value = c.Context.Value(key)
	}

	return
}
