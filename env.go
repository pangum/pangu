package pangu

import (
	`strings`
)

type env struct {
	key   string
	value string
}

func parseEnv(from string) (_env *env) {
	data := strings.Split(from, envSeparator)
	if envCount != len(data) {
		return
	}

	_env = &env{
		key:   data[0],
		value: data[1],
	}

	return
}
