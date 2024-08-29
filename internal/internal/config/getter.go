package config

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/goexl/exception"
	"github.com/goexl/gfx"
	"github.com/goexl/mengpo"
	"github.com/goexl/xiren"
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
	} else if fe := g.Fill(path, target); nil != fe { // 加载数据
		err = fe
	} else if nil != g.params.Watcher { // 配置文件监控
		// TODO err = g.Watch(target, g.params.Watcher)
	} else {
		g.path = path
	}

	return
}

func (g *Getter) Fill(path string, config runtime.Pointer) (err error) {
	if le := g.load(path, config); nil != le { // 从路径中加载数据
		err = le
	} else if g.params.Default { // 处理默认值
		// !此处逻辑不能往前，原因是如果对象里面包含指针，那么只能在包含指针的结构体被解析后才能去设置默认值，不然指针将被会设置成空值
		err = mengpo.New().Tag(g.params.Tag.Default).Build().Set(config)
	} else if g.params.Validate { // 数据验证
		err = xiren.Struct(config)
	}

	return
}

func (g *Getter) load(path string, config runtime.Pointer) (err error) {
	if _, se := os.Stat(path); nil != se && os.IsNotExist(se) && g.params.Nullable { // 允许不使用配置文件
		// 空实现，纯占位
	} else if le := g.params.Load(path, config); nil != le { // 从路径中加载数据
		err = le
	}

	return
}

func (g *Getter) handleFields(object reflect.Value, prefix string, callback func(fieldName string) interface{}) {
	typ := object.Type()
	for index := 0; index < object.NumField(); index++ {
		field := object.Field(index)
		fieldName := typ.Field(index).Name

		// 获取字段的 JSON 标签，如果存在则使用标签作为字段名
		jsonTag := typ.Field(index).Tag.Get("json")
		if jsonTag != "" {
			jsonTags := strings.Split(jsonTag, ",")
			if len(jsonTags) > 0 {
				fieldName = jsonTags[0]
			}
		}

		// 构建带有下划线的字段名
		fieldNameWithPrefix := fieldName
		if prefix != "" {
			fieldNameWithPrefix = prefix + "_" + fieldName
		}

		// 如果字段为结构体，则递归处理
		if field.Kind() == reflect.Struct {
			g.handleFields(field, fieldNameWithPrefix, callback)
		} else {
			// 如果字段为零值，则通过回调函数获取新值并设置回原来的字段
			if reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
				newValue := callback(fieldNameWithPrefix)
				field.Set(reflect.ValueOf(newValue))
			}
		}
	}
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
	} else if !g.params.Nullable && !exists {
		err = exception.New().Message("配置文件不存在").Build()
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
