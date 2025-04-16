package get

import (
	"github.com/goexl/log"
	"github.com/heluon/di"
)

type Logger struct {
	di.Get

	Optional log.Logger `optional:"true"`
}
