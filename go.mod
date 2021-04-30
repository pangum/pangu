module github.com/storezhang/pangu

go 1.16

require (
	github.com/go-redis/redis/v8 v8.8.0
	github.com/go-resty/resty/v2 v2.5.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/wire v0.5.0
	github.com/olivere/elastic/v7 v7.0.24
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/rubenv/sql-migrate v0.0.0-20210408115534-a32ed26c37ea
	github.com/storezhang/glog v1.0.5
	github.com/storezhang/gox v1.4.1
	github.com/urfave/cli/v2 v2.3.0
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
	xorm.io/builder v0.3.7
	xorm.io/core v0.7.3
	xorm.io/xorm v1.0.5
)

replace xorm.io/xorm => gitea.com/storezhang/xorm v1.0.7

replace github.com/storezhang/gox => ../gox

// replace github.com/storezhang/glog => ../glog
// replace github.com/class100/sdk-go => ../../class100/sdk-go
// replace github.com/storezhang/validatorx => ../../storezhang/validatorx
// replace github.com/class100/yunke-core => ../../../github/yunke-core
