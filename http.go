package pangu

import (
	`github.com/storezhang/gox`
)

func getClientConfig(config gox.HttpConfig) gox.HttpClientConfig {
	return config.Client
}

func getServerConfig(config gox.HttpConfig) gox.HttpServerConfig {
	return config.Server
}
