module github.com/pangum/pangu

go 1.16

require (
 	github.com/drone/envsubst v1.0.3
	github.com/common-nighthawk/go-figure v0.0.0-20200609044655-c4b36f998cf2
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	github.com/pelletier/go-toml v1.2.0
	github.com/storezhang/dig v0.0.1
	github.com/storezhang/gox v1.8.1
	github.com/storezhang/guc v0.0.2
	github.com/storezhang/mengpo v0.1.0
	github.com/storezhang/simaqian v0.0.6
	github.com/storezhang/validatorx v1.0.9
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/crypto v0.0.0-20211115234514-b4de73f9ece8 // indirect
	golang.org/x/sys v0.0.0-20211113001501-0c823b97ae02 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

// replace github.com/storezhang/gox => ../gox
// replace github.com/storezhang/god => ../../storezhang/god
// replace github.com/storezhang/simaqian => ../glog
// replace github.com/storezhang/validatorx => ../../storezhang/validatorx
