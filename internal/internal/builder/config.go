package builder

import (
	"github.com/pangum/config"
	"github.com/pangum/pangu/internal/internal/app"
	"github.com/pangum/pangu/internal/param"
)

type Config[T any] struct {
	param *param.Config
	from  *T
}

func NewConfig[T any](param *param.Config, from *T) *Config[T] {
	return &Config[T]{
		param: param,
		from:  from,
	}
}

func (c *Config[T]) Nullable() (t *T) {
	c.param.Nullable = true
	t = c.from

	return
}

func (c *Config[T]) Required() (t *T) {
	c.param.Nullable = false
	t = c.from

	return
}

func (c *Config[T]) Default() (t *T) {
	c.param.Default = true
	t = c.from

	return
}

func (c *Config[T]) Environment(key string, value string) (t *T) {
	c.param.Environments = append(c.param.Environments, app.NewEnvironment(key, value))
	t = c.from

	return
}

func (c *Config[T]) Loader(loader config.Loader, loaders ...config.Loader) (t *T) {
	c.param.Loaders = append(c.param.Loaders, loader)
	c.param.Loaders = append(c.param.Loaders, loaders...)
	t = c.from

	return
}
