package plugin

import (
	`github.com/storezhang/pangu`
)

func init() {
	app := pangu.New()
	if err := app.Provides(newTest); nil != err {
		panic(err)
	}
}
