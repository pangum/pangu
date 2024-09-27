package os

import (
	"fmt"
	"os"

	"github.com/goexl/gox"
	"github.com/pangum/pangu/internal/internal/constant"
)

func GetEnvironment(key string, current string, defaults string) string {
	return gox.Ift(current != defaults, getEnvironment(key, current), current)
}

func getEnvironment(key string, dv string) (final string) {
	if first, fk := os.LookupEnv(key); fk {
		final = first
	} else if second, sk := os.LookupEnv(fmt.Sprintf(constant.EnvironmentFormatter, key)); sk {
		final = second
	} else {
		final = dv
	}

	return
}
