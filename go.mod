module github.com/pangum/pangu

go 1.21

require (
	github.com/goexl/gfx v0.1.7
	github.com/goexl/gox v0.5.5
	github.com/pelletier/go-toml v1.9.5
	github.com/urfave/cli/v2 v2.27.4
	go.uber.org/dig v1.18.0
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/drone/envsubst v1.0.3
	github.com/goexl/env v0.0.2
	github.com/goexl/exception v0.0.1
	github.com/goexl/log v0.0.7
	github.com/goexl/mengpo v0.2.6
	github.com/goexl/xiren v0.0.6
	github.com/olekukonko/tablewriter v0.0.5
	github.com/zs5460/art v0.3.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.5 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.22.0 // indirect
	github.com/goexl/baozheng v0.0.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rs/xid v1.5.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/net v0.28.0 // indirect
)

// v1 项目从原来的storezhang/pangu迁移过来，原来的版本号不再使用，直到最新发布到该版本
retract [v1.0.0, v1.3.9]
