package plugin

import (
	`github.com/pangum/pangu`
)

func init() {
	app := pangu.New()
	if err := app.Provides(newTest); nil != err {
		panic(err)
	}
}
