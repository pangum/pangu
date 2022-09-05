package pangu

var (
	_              = ConfigWatcher
	_ configOption = (*optionConfigWatcher)(nil)
)

type optionConfigWatcher struct {
	watcher configWatcher
}

// ConfigWatcher 配置监控
func ConfigWatcher(watcher configWatcher) *optionConfigWatcher {
	return &optionConfigWatcher{
		watcher: watcher,
	}
}

func (cw *optionConfigWatcher) applyConfig(options *configOptions) {
	options.watcher = cw.watcher
}
