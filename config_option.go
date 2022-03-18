package pangu

import (
	`path/filepath`

	`github.com/goexl/env`
)

type (
	configOption interface {
		applyConfig(options *configOptions)
	}

	configOptions struct {
		paths      []string
		extensions []string

		defaults  bool
		validates bool

		tag               tag
		environmentGetter environmentGetter
		environments      []*environment
	}
)

func defaultConfigOptions() *configOptions {
	return &configOptions{
		paths: []string{
			applicationName,
			filepath.Join(configDir, applicationName),
			filepath.Join(confDir, applicationName),
		},
		extensions: []string{
			ymlExt,
			yamlExt,
			tomlExt,
			jsonExt,
			xmlExt,
		},

		defaults:  true,
		validates: true,

		tag: tag{
			defaults: defaultsTag,
		},

		environmentGetter: env.Get,
		environments:      make([]*environment, 0),
	}
}
