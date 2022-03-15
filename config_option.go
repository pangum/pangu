package pangu

type (
	configOption interface {
		applyConfig(options *configOptions)
	}

	configOptions struct {
		paths      []string
		extensions []string

		defaults bool
		validate bool

		tag  tag
		envs []*env
	}
)

func defaultConfigOptions() *configOptions {
	return &configOptions{
		paths: []string{},
		extensions: []string{
			ymlExt,
			yamlExt,
			tomlExt,
			jsonExt,
			xmlExt,
		},

		defaults: true,
		validate: true,

		tag: tag{
			defaults: `default`,
		},
	}
}