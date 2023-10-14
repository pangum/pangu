package internal

import (
	"os"
)

type Environments []*Environment

func (e Environments) Set() (err error) {
	for _, environment := range e {
		err = os.Setenv(environment.Key, environment.Value)
	}

	return
}
