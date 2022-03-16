package guide

type constructor struct{}

func newConstructor() *constructor {
	return new(constructor)
}
