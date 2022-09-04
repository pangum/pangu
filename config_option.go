package pangu

import (
	"path/filepath"

	"github.com/goexl/env"
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
		nullable  bool

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
			filepath.Join(configurationDir, applicationName),
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
		nullable:  true,

		tag: tag{
			defaults: defaultsTag,
		},

		environmentGetter: env.Get,
		environments:      make([]*environment, 0),
	}
}
