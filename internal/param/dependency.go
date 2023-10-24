package param

type Dependency struct {
	Puts []*Put
	Gets []*Get
}

func NewDependency() *Dependency {
	return new(Dependency)
}
