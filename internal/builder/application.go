package builder

import (
	internal2 "github.com/pangum/pangu/internal/core"
	"github.com/pangum/pangu/internal/internal"
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

func (a *Application) Verify() (application *Application) {
	a.params.Verify = true
	application = a

	return
}

func (a *Application) Author(name string, email string) (application *Application) {
	a.params.Authors = append(a.params.Authors, internal.NewAuthor(name, email))
	application = a

	return
}

func (a *Application) Copyright(copyright string) (application *Application) {
	a.params.Copyright = copyright
	application = a

	return
}

func (a *Application) Description(description string) (application *Application) {
	a.params.Description = description
	application = a

	return
}

func (a *Application) Metadata(key string, value any) (application *Application) {
	a.params.Metadata[key] = value
	application = a

	return
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

func (a *Application) Get() *internal2.Application {
	return internal2.New(a.params)
}
