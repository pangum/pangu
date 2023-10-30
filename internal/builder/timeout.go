package builder

import (
	"time"

	"github.com/pangum/pangu/internal/param"
)

type Timeout struct {
	params      *param.Timeout
	application *Application
}

func NewTimeout(application *Application) (timeout *Timeout) {
	timeout = new(Timeout)
	timeout.params = application.params.Timeout
	timeout.application = application

	return
}

func (t *Timeout) Boot(duration time.Duration) (timeout *Timeout) {
	t.params.Boot = duration
	timeout = t

	return
}

func (t *Timeout) Quit(duration time.Duration) (timeout *Timeout) {
	t.params.Quit = duration
	timeout = t

	return
}

func (t *Timeout) Build() (application *Application) {
	t.params.Set = true
	t.application.params.Timeout = t.params
	application = t.application

	return
}
