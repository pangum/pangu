package guide

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newConstructor,
		newAnswer,
		newQa,
		newServer,
		newIn,
	)
}
