package runtime

import (
	"github.com/pangum/pangu/internal/constant"
	"github.com/urfave/cli/v2"
)

type Shadow struct {
	*cli.App
}

func NewShadow() (shadow *Shadow) {
	app := cli.NewApp()
	app.EnableBashCompletion = false
	app.UseShortOptionHandling = false
	app.Commands = append(app.Commands, &cli.Command{
		Name:   constant.CommandSilent,
		Action: shadow.silent,
	})

	shadow = new(Shadow)
	shadow.App = app

	return
}

func (s *Shadow) silent(_ *cli.Context) (err error) {
	return
}
