package builder

import (
	"github.com/pangum/pangu/internal"
	"github.com/pangum/pangu/internal/builder/internal/function"
	"github.com/pangum/pangu/internal/core"
	"github.com/pangum/pangu/internal/internal/app"
	"github.com/pangum/pangu/internal/param"
)

type Application struct {
	params *param.Application
}

func NewApplication() *Application {
	return &Application{
		params: param.NewApplication(),
	}
}

func (a *Application) Verify() *Application {
	return a.set(func() {
		a.params.Verify = true
	})
}

func (a *Application) Author(name string, email string) *Application {
	return a.set(func() {
		a.params.Authors = append(a.params.Authors, app.NewAuthor(name, email))
	})
}

func (a *Application) Copyright(copyright string) *Application {
	return a.set(func() {
		a.params.Copyright = copyright
	})
}

func (a *Application) Description(description string) *Application {
	return a.set(func() {
		a.params.Description = description
	})
}

func (a *Application) Usage(usage string) *Application {
	return a.set(func() {
		a.params.Usage = usage
	})
}

func (a *Application) Metadata(key string, value any) *Application {
	return a.set(func() {
		a.params.Metadata[key] = value
	})
}

func (a *Application) Name(name string) *Application {
	return a.set(func() {
		internal.Name = name
	})
}

func (a *Application) Timeout() *Timeout {
	return NewTimeout(a)
}

func (a *Application) Banner() *Banner {
	return NewBanner(a)
}

func (a *Application) Config() *Config {
	return NewConfig(a)
}

func (a *Application) Help() *Help {
	return NewHelp(a)
}

func (a *Application) Get() *core.Application {
	return core.New(a.params)
}

func (a *Application) set(set function.Set) (application *Application) {
	set()
	a.params.Set = true
	application = a

	return
}
