package pangu

import (
	`github.com/urfave/cli/v2`
)

func newApp() (app *cli.App) {
	app = cli.NewApp()
	app.EnableBashCompletion = true

	return
}
