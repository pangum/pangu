package pangu

import (
	`context`

	`github.com/storezhang/gox`
)

// NewRedis 创建Redis客户端
func NewRedis(config gox.RedisConfig) (client *redis.Client, err error) {
	options := &redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	}
	client = redis.NewClient(options)

	// 测试Redis连接
	_, err = client.Ping(context.Background()).Result()

	return
}
