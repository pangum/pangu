package get

import (
	"github.com/goexl/log"
	"go.uber.org/dig"
)

type Logger struct {
	dig.In

	Optional log.Logger `optional:"true"`
}
