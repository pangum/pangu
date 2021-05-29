package conf

import (
	`github.com/storezhang/pangu`
)

func Provides(application *pangu.Application) error {
	return application.Provides(config, example, http)
}
