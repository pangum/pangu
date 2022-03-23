package conf

import (
	`github.com/pangum/pangu`
)

func init() {
	pangu.New().Dependencies(config, example)
}
