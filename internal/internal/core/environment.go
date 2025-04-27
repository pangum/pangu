package core

type Environment struct {
	Key   string
	Value string
}

func NewEnvironment(key string, value string) *Environment {
	return &Environment{
		Key:   key,
		Value: value,
	}
}
