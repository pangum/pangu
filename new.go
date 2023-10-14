package pangu

import (
	"github.com/pangum/pangu/internal/builder"
)

var _ = New

func New() *builder.Application {
	return builder.NewApplication()
}
