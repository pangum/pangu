module github.com/storezhang/pangu

go 1.16

require (
	github.com/common-nighthawk/go-figure v0.0.0-20200609044655-c4b36f998cf2
	github.com/mcuadros/go-defaults v1.2.0
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/storezhang/glog v1.0.8
	github.com/storezhang/gox v1.5.3
	github.com/storezhang/validatorx v1.0.5
	github.com/urfave/cli/v2 v2.3.0
	go.uber.org/dig v1.10.0
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/sys v0.0.0-20210423082822-04245dca01da // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/tools v0.0.0-20201224043029-2b0845dc783e // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

// replace github.com/storezhang/gox => ../gox
// replace github.com/storezhang/glog => ../glog
// replace github.com/storezhang/validatorx => ../../storezhang/validatorx
