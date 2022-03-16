package guide

import (
	`github.com/pangum/pangu`
)

func init() {
	pangu.New().Musts(
		newConstructor,
		newAnswer,
		newQa,
		newServer,
		newIn,
	)
}
