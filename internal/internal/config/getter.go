package config

import (
	"path/filepath"

	"github.com/goexl/gfx"
	"github.com/pangum/pangu/internal"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
	"github.com/urfave/cli/v2"
)

type Getter struct {
	path   string
	params *param.Config
}

func newConfig(params *param.Config) *Getter {
	return &Getter{
		params: params,
	}
}

func (g *Getter) Get(target runtime.Pointer) (err error) {
	if path, fpe := g.filepath(); nil != fpe {
		err = fpe
	} else if fe := g.params.Fill(path, target); nil != fe { // 加载数据
		err = fe
	} else if nil != g.params.Watcher { // 配置文件监控
		// TODO err = g.Watch(target, g.params.Watcher)
	} else {
		g.path = path
	}

	return
}

func (g *Getter) filepath() (path string, err error) {
	gfxOptions := gfx.NewExistsOptions(
		gfx.Paths(g.params.Paths...),
		gfx.Extensions(g.params.Extensions...),
	)
	// 如果配置了应用名称，可以使用应用名称的配置文件
	if constant.EnvironmentNotSet != internal.Name {
		gfxOptions = append(gfxOptions, gfx.Paths(
			internal.Name,
			filepath.Join(constant.ConfigDir, internal.Name),
			filepath.Join(constant.ConfigConfDir, internal.Name),
			filepath.Join(constant.ConfigConfigurationDir, internal.Name),
		))
	}

	if final, exists := gfx.Exists(g.path, gfxOptions...); exists {
		path = final
	} else { // 如果找不到配置文件，则所用默认的配置文件
		path = g.path
	}

	return
}

func (g *Getter) bind(shell *runtime.Shell, shadow *runtime.Shadow) {
	config := new(cli.StringFlag)
	config.Name = constant.ConfigName
	config.Aliases = []string{
		constant.ConfigAliasC,
		constant.ConfigAliasConf,
		constant.ConfigAliasConfiguration,
	}
	config.Value = constant.ConfigDefaultFilepath
	config.Usage = "指定配置文件路径"
	config.Destination = &g.path

	shell.Flags = append(shell.Flags, config)
	shadow.Flags = append(shadow.Flags, config)
}
