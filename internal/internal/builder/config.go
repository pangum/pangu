package builder

import (
	"github.com/goexl/gox"
	"github.com/heluon/boot/internal/internal/callback"
	"github.com/heluon/boot/internal/internal/kernel"
	"github.com/heluon/boot/internal/internal/param"

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

func (c *Config[T]) Filepath(required string, optionals ...string) (t *T) {
	c.param.Paths = append(c.param.Paths, required)
	c.param.Paths = append(c.param.Paths, optionals...)
	t = c.from

	return
}

func (c *Config[T]) Environment(key string, value string) (t *T) {
	c.param.Environments = append(c.param.Environments, kernel.NewEnvironment(key, value))
	t = c.from

	return
}

func (c *Config[T]) Loader(loader config.Loader, loaders ...config.Loader) (t *T) {
	c.param.Loaders = append(c.param.Loaders, loader)
	c.param.Loaders = append(c.param.Loaders, loaders...)
	t = c.from

	return
}

func (c *Config[T]) Getter(getter callback.Getter, getters ...callback.Getter) (t *T) {
	c.param.Getters[getter] = new(gox.Empty)
	for _, _getter := range getters {
		c.param.Getters[_getter] = new(gox.Empty)
	}
	t = c.from

	return
}
