package application

import (
	"context"
)

type Before interface {
	Before(context.Context) error
}
