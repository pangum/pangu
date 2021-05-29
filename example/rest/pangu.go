package rest

import (
	`github.com/storezhang/pangu`
)

func Provides(application *pangu.Application) error {
	return application.Provides(newServer)
}
