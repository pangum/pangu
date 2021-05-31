package command

import (
	`github.com/storezhang/pangu`
)

func Provides(application *pangu.Application) error {
	return application.Provides(newTest)
}
