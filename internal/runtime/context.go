package runtime

import (
	"context"
	"time"

	"github.com/goexl/gox"
	"github.com/urfave/cli/v2"
)

var _ context.Context = (*Context)(nil)

// Context 描述上下文
type Context struct {
	context *cli.Context
}

// NewContext 创建上下文
func NewContext(ctx *cli.Context) *Context {
	return &Context{
		context: ctx,
	}
}

func (c *Context) Deadline() (time.Time, bool) {
	return c.context.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.context.Done()
}

func (c *Context) Err() error {
	return c.context.Err()
}

func (c *Context) String(key string) string {
	return c.context.String(key)
}

func (c *Context) Set(key string, value string) error {
	return c.context.Set(key, value)
}

func (c *Context) Value(key any) any {
	return c.context.Value(gox.ToString(key))
}
