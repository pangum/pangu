package builder

import (
	"github.com/pangum/pangu/internal"
	"github.com/pangum/pangu/internal/builder/internal/function"
	"github.com/pangum/pangu/internal/core"
	"github.com/pangum/pangu/internal/internal/app"
	"github.com/pangum/pangu/internal/param"
)

var shadow *Application

type Application struct {
	params *param.Application

	config  *Config
	timeout *Timeout
	banner  *Banner
	help    *Help
}

// NewApplication !基于单例实现，保证每次只生成一个可配置项
func NewApplication() (application *Application) {
	once.Do(newApplication)
	application = shadow

	return
}

func newApplication() {
	shadow = new(Application)
	shadow.params = param.NewApplication()

	// !预创建，保证单例
	shadow.config = newConfig(shadow)
	shadow.timeout = newTimeout(shadow)
	shadow.banner = newBanner(shadow)
	shadow.help = newHelp(shadow)

	return
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
	return a.timeout
}

func (a *Application) Banner() *Banner {
	return a.banner
}

func (a *Application) Config() *Config {
	return a.config
}

func (a *Application) Help() *Help {
	return a.help
}

func (a *Application) Get() *core.Application {
	return core.New(a.params)
}

func (a *Application) set(set function.Set) (application *Application) {
	set()
	application = a

	return
}
