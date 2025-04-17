package get

import (
	"github.com/goexl/log"
	"github.com/harluo/di"
)

type Logger struct {
	di.Get

	Optional log.Logger `optional:"true"`
}
