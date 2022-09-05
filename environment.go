package pangu

import (
	"strings"
)

type environment struct {
	key   string
	value string
}

func parseEnv(from string) (env *environment) {
	data := strings.Split(from, envSeparator)
	if envCount != len(data) {
		return
	}

	env = &environment{
		key:   data[0],
		value: data[1],
	}

	return
}
