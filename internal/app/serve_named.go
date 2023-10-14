package app

var _ = NewNamedServe

// NamedServe 命名的服务
type NamedServe struct {
	UnstoppableServe

	name string
}

// NewNamedServe 创建命名服务
func NewNamedServe(name string) NamedServe {
	return NamedServe{
		name: name,
	}
}

func (ns NamedServe) Name() string {
	return ns.name
}
