package optional

type External struct {
	Logger bool
}

func NewExternal() *External {
	return &External{
		Logger: false,
	}
}
