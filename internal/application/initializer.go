package application

import (
	"context"
)

type Initializer interface {
	Initialize(context.Context) error
}
