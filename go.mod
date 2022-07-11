module github.com/pangum/pangu

go 1.17

require (
	github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be
	github.com/goexl/exc v0.0.4
	github.com/goexl/gfx v0.0.8
	github.com/goexl/gox v0.0.5
	github.com/goexl/mengpo v0.1.7
	github.com/goexl/simaqian v0.1.3
	github.com/goexl/xiren v0.0.3
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/pelletier/go-toml v1.9.5
	github.com/storezhang/dig v0.0.1
	github.com/urfave/cli/v2 v2.11.0
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064 // indirect
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/drone/envsubst v1.0.3
	github.com/goexl/env v0.0.2
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/goexl/baozheng v0.0.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/rs/xid v1.4.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
)

// v1 项目从原来的storezhang/pangu迁移过来，原来的版本号不再使用，直到最新发布到该版本
retract [v1.0.0, v1.3.9]

// replace github.com/storezhang/gox => ../gox
// replace github.com/goexl/gfx => ../../goexl/gfx
// replace github.com/goexl/exc => ../../goexl/exc
