package builder

import (
	"time"

	"github.com/pangum/core/internal/internal/param"
)

type Timeout struct {
	params      *param.Timeout
	application *Application
}

func newTimeout(application *Application) *Timeout {
	return &Timeout{
		params:      application.params.Timeout,
		application: application,
	}
}

func (t *Timeout) Boot(duration time.Duration) (timeout *Timeout) {
	t.params.Startup = duration
	timeout = t

	return
}

func (t *Timeout) Quit(duration time.Duration) (timeout *Timeout) {
	t.params.Quit = duration
	timeout = t

	return
}

func (t *Timeout) Build() (application *Application) {
	t.application.params.Timeout = t.params
	application = t.application

	return
}
