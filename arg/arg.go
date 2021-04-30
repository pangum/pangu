package arg

// Arg 描述一个参数
type Arg struct {
	// 名称
	Name string
	// 别名
	Aliases []string
	// 使用方法
	Usage string
	// 环境变量
	EnvVars []string
	// 是否是必须
	Required bool
	// 是否隐藏
	Hidden bool
	// 默认值
	DefaultText string
}

func (a *Arg) GetName() string {
	return a.Name
}

func (a *Arg) GetAliases() []string {
	return a.Aliases
}

func (a *Arg) GetUsage() string {
	return a.Usage
}

func (a *Arg) GetDefaultText() string {
	return a.DefaultText
}
