module github.com/pangum/pangu

go 1.19

require (
	github.com/goexl/exc v0.0.5
	github.com/goexl/gfx v0.1.7
	github.com/goexl/gox v0.2.0
	github.com/goexl/mengpo v0.2.4
	github.com/goexl/simaqian v0.2.4
	github.com/goexl/xiren v0.0.5
	github.com/pelletier/go-toml v1.9.5
	github.com/storezhang/dig v0.0.1
	github.com/urfave/cli/v2 v2.25.1
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/drone/envsubst v1.0.3
	github.com/goexl/env v0.0.2
	github.com/zs5460/art v0.2.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/goexl/baozheng v0.0.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/rs/xid v1.4.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
)

// v1 项目从原来的storezhang/pangu迁移过来，原来的版本号不再使用，直到最新发布到该版本
retract [v1.0.0, v1.3.9]
