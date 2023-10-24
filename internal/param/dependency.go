package param

type Dependency struct {
	Puts   []*Put
	Gets   []*Get
	Verify bool
}

func NewDependency(verify bool) *Dependency {
	return &Dependency{
		Verify: verify,
	}
}

func (d *Dependency) Clear() {
	d.Puts = make([]*Put, 0, 1)
	d.Gets = make([]*Get, 0, 1)
}
