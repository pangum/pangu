package os

import (
	"fmt"
	"os"

	"github.com/pangum/pangu/internal/constant"
)

func GetEnvironment(key string) (final string) {
	if first, fk := os.LookupEnv(key); fk {
		final = first
	} else if second, sk := os.LookupEnv(fmt.Sprintf(constant.EnvironmentFormatter, key)); sk {
		final = second
	} else {
		final = constant.EnvironmentNotSet
	}

	return
}
