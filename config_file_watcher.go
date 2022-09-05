package pangu

import (
	"reflect"

	"github.com/goexl/gfx"
)

var _ gfx.Watcher = (*configFileWatcher)(nil)

type configFileWatcher struct {
	config  any
	path    string
	watcher configWatcher
	options *configOptions
}

func newConfigFileWatcher(config any, path string, watcher configWatcher, options *configOptions) *configFileWatcher {
	return &configFileWatcher{
		config:  config,
		path:    path,
		watcher: watcher,
		options: options,
	}
}

func (cfw *configFileWatcher) OnChanged(path string) {
	if cfw.path != path {
		return
	}

	to := reflect.New(reflect.TypeOf(cfw.config).Elem()).Interface()
	if err := cfw.options.Load(path, to); nil == err && nil != cfw.watcher {
		go cfw.watcher.OnChanged(cfw.path, cfw.config, to)
		cfw.config = to
	}
}

func (cfw *configFileWatcher) OnRenamed(_ string) {}

func (cfw *configFileWatcher) OnDeleted(path string) {
	if nil != cfw.options.watcher {
		go cfw.watcher.OnDeleted(path)
	}
}

func (cfw *configFileWatcher) OnCreated(_ string) {}

func (cfw *configFileWatcher) OnPermissionChanged(_ string) {}

func (cfw *configFileWatcher) OnError(err error) {
	cfw.watcher.OnError(cfw.path, err)
}
