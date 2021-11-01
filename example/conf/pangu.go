package conf

import (
	`github.com/pangum/pangu`
)

func Provides(application *pangu.Application) error {
	return application.Provides(config, example, http)
}
