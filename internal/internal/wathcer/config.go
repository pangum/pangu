package wathcer

import (
	"reflect"

	"github.com/goexl/gfx"
	"github.com/pangum/pangu/internal/config"
	"github.com/pangum/pangu/internal/param"
)

var _ gfx.Watcher = (*Config)(nil)

type Config struct {
	config  any
	path    string
	watcher config.Watcher
	param   *param.Config
}

func NewConfig(config any, path string, watcher config.Watcher, options *param.Config) *Config {
	return &Config{
		config:  config,
		path:    path,
		watcher: watcher,
		param:   options,
	}
}

func (c *Config) OnChanged(path string) {
	if c.path != path {
		return
	}

	to := reflect.New(reflect.TypeOf(c.config).Elem()).Interface()
	if err := c.param.Load(path, to); nil == err && nil != c.watcher {
		go c.watcher.OnChanged(c.path, c.config, to)
		c.config = to
	}
}

func (c *Config) OnRenamed(_ string) {}

func (c *Config) OnDeleted(path string) {
	if nil != c.param.Watcher {
		go c.watcher.OnDeleted(path)
	}
}

func (c *Config) OnCreated(_ string) {}

func (c *Config) OnPermissionChanged(_ string) {}

func (c *Config) OnError(err error) {
	c.watcher.OnError(c.path, err)
}
