package param

type Injection struct {
	Name  string
	Group string
}

func NewInjection() *Injection {
	return new(Injection)
}
