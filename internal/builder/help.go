package builder

import (
	"github.com/pangum/pangu/internal/param"
)

type Help struct {
	params      *param.Help
	application *Application
}

func NewHelp(application *Application) *Help {
	return &Help{
		params:      param.NewHelp(),
		application: application,
	}
}

func (h *Help) App(tooltip string) (help *Help) {
	h.params.App = tooltip
	help = h

	return
}

func (h *Help) Command(tooltip string) (help *Help) {
	h.params.Command = tooltip
	help = h

	return
}

func (h *Help) Subcommand(tooltip string) (help *Help) {
	h.params.Subcommand = tooltip
	help = h

	return
}

func (h *Help) Build() (application *Application) {
	h.application.params.Help = h.params
	application = h.application

	return
}
