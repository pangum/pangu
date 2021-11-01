package command

import (
	`github.com/pangum/pangu`
)

func Provides(application *pangu.Application) error {
	return application.Provides(newTest)
}
