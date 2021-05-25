module github.com/storezhang/pangu

go 1.16

require (
	github.com/common-nighthawk/go-figure v0.0.0-20200609044655-c4b36f998cf2
	github.com/go-resty/resty/v2 v2.6.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/snappy v0.0.3 // indirect
	github.com/mcuadros/go-defaults v1.2.0
	github.com/onsi/ginkgo v1.15.0 // indirect
	github.com/onsi/gomega v1.10.5 // indirect
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rubenv/sql-migrate v0.0.0-20210408115534-a32ed26c37ea
	github.com/storezhang/glog v1.0.5
	github.com/storezhang/gox v1.4.10
	github.com/storezhang/validatorx v1.0.5
	github.com/urfave/cli/v2 v2.3.0
	go.uber.org/dig v1.10.0
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20210521195947-fe42d452be8f // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	xorm.io/builder v0.3.9
	xorm.io/xorm v1.1.0
)

// replace github.com/storezhang/gox => ../gox
// replace github.com/storezhang/glog => ../glog
// replace github.com/class100/sdk-go => ../../class100/sdk-go
// replace github.com/storezhang/validatorx => ../../storezhang/validatorx
// replace github.com/class100/yunke-core => ../../../github/yunke-core
