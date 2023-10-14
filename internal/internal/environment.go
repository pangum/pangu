package internal

import (
	"strings"

	"github.com/pangum/pangu/internal/constant"
)

type Environment struct {
	Key   string
	Value string
}

func ParseEnvironment(from string) (env *Environment) {
	data := strings.Split(from, constant.EnvironmentSeparator)
	if constant.EnvironmentCount != len(data) {
		return
	}

	env = &Environment{
		Key:   data[0],
		Value: data[1],
	}

	return
}
