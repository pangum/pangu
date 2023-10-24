package runtime

import (
	"runtime"
)

func FuncForPC(pc uintptr) *runtime.Func {
	return runtime.FuncForPC(pc)
}
