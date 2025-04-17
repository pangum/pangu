package once

import (
	"sync"
)

var Application = new(sync.Once)
