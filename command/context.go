package command

// Context 描述上下文
type Context struct {
	context *cli.Context
}

// NewContext 创建一个上下文
func NewContext(ctx *cli.Context) *Context {
	return &Context{context: ctx}
}

func (c *Context) String(name string) string {
	return c.context.String(name)
}
