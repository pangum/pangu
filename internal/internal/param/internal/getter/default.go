package getter

type Default struct {
	get func(string) string
}

func NewDefault(get func(string) string) *Default {
	return &Default{
		get: get,
	}
}

func (g *Default) Get(key string) string {
	return g.get(key)
}
