module github.com/pangum/pangu

go 1.16

require (
	github.com/common-nighthawk/go-figure v0.0.0-20200609044655-c4b36f998cf2
	github.com/mcuadros/go-defaults v1.2.0
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/storezhang/dig v0.0.1
	github.com/storezhang/gox v1.7.9
	github.com/storezhang/guc v0.0.2
	github.com/storezhang/simaqian v0.0.3
	github.com/storezhang/validatorx v1.0.5
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

// replace github.com/storezhang/gox => ../gox
// replace github.com/storezhang/simaqian => ../glog
// replace github.com/storezhang/validatorx => ../../storezhang/validatorx
