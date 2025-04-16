package internal

import (
	"sync"
)

var Once = new(sync.Once)
