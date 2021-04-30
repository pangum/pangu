package pangu

import (
	`github.com/storezhang/gox`
)

// NewElasticsearch 创建Elasticsearch客户端
func NewElasticsearch(config gox.ElasticsearchConfig) (client *elastic.Client, err error) {
	options := []elastic.ClientOptionFunc{
		elastic.SetURL(config.Address),
		elastic.SetGzip(config.Gzip),
	}

	if "" != config.Username {
		options = append(options, elastic.SetBasicAuth(config.Username, config.Password))
	}
	if 0 != len(config.Headers) {
		options = append(options, elastic.SetHeaders(config.Headers))
	}
	if 0 != len(config.Plugins) {
		options = append(options, elastic.SetRequiredPlugins(config.Plugins...))
	}

	// 嗅探器
	if 0 != config.Sniffer.Interval || 0 != config.Sniffer.Timeout || 0 != config.Sniffer.StartupTimeout {
		options = append(options, elastic.SetSniff(true))
	}
	if 0 != config.Sniffer.Interval {
		options = append(options, elastic.SetSnifferInterval(config.Sniffer.Interval))
	}
	if 0 != config.Sniffer.Timeout {
		options = append(options, elastic.SetSnifferTimeout(config.Sniffer.Timeout))
	}
	if 0 != config.Sniffer.StartupTimeout {
		options = append(options, elastic.SetSnifferTimeoutStartup(config.Sniffer.StartupTimeout))
	}

	// 健康检查
	if 0 != config.Health.Interval || 0 != config.Health.Timeout || 0 != config.Health.StartupTimeout {
		options = append(options, elastic.SetHealthcheck(true))
	}
	if 0 != config.Health.Interval {
		options = append(options, elastic.SetHealthcheckInterval(config.Health.Interval))
	}
	if 0 != config.Health.Timeout {
		options = append(options, elastic.SetHealthcheckTimeout(config.Health.Timeout))
	}
	if 0 != config.Health.StartupTimeout {
		options = append(options, elastic.SetHealthcheckTimeoutStartup(config.Health.StartupTimeout))
	}

	// 创建客户端，有问题将错误返回，不能不处理错误或者将错误吃掉
	if client, err = elastic.NewClient(options...); nil != err {
		return
	}

	return
}
