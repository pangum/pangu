package config

import (
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/goexl/exception"
	"github.com/goexl/gfx"
	"github.com/goexl/gox"
	"github.com/goexl/mengpo"
	"github.com/goexl/xiren"
	"github.com/pangum/pangu/internal"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/internal/config/internal/callback"
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
	// 从环境变量中加载配置
	defer g.processEnvironmentConfig(reflect.ValueOf(config).Elem(), gox.NewSlice[string](), g.getEnvironmentConfig)

	if _, se := os.Stat(path); nil != se && os.IsNotExist(se) && g.params.Nullable { // 允许不使用配置文件
		// 空实现，纯占位
	} else if le := g.params.Load(path, config); nil != le { // 从路径中加载数据
		err = le
	}

	return
}

func (g *Getter) processEnvironmentConfig(object reflect.Value, names gox.Slice[string], set callback.SetValue) {
	typ := object.Type()
	for index := 0; index < object.NumField(); index++ {
		field := object.Field(index)
		name := typ.Field(index).Name
		kind := field.Kind()

		if reflect.Struct == kind { // 如果字段为结构体，则递归处理
			names = append(names, name)
			g.processEnvironmentConfig(field, names, set)
		} else if reflect.Ptr == kind { // 如果是指针，初始化
			field.Set(reflect.New(field.Type().Elem()))
		} else if g.zero(field) { // 如果字段为零值，则通过回调函数获取新值并设置回原来的字段
			copies := names.Clone() // !复制一份字段名列表，防止干扰其它结构体字段
			set(append(copies, name), field)
		}
	}
}

func (g *Getter) getEnvironmentConfig(names gox.Slice[string], field reflect.Value) {
	switch field.Kind() {
	case reflect.Bool:
		g.setEnvironmentConfigValue(names, field, g.bool)
	case reflect.Int:
		g.setEnvironmentConfigValue(names, field, g.int)
	case reflect.Int8:
		g.setEnvironmentConfigValue(names, field, g.int8)
	case reflect.Int16:
		g.setEnvironmentConfigValue(names, field, g.int16)
	case reflect.Int32:
		g.setEnvironmentConfigValue(names, field, g.int32)
	case reflect.Int64:
		g.setEnvironmentConfigValue(names, field, g.int64)
	case reflect.Uint:
		g.setEnvironmentConfigValue(names, field, g.uint)
	case reflect.Uint8:
		g.setEnvironmentConfigValue(names, field, g.uint8)
	case reflect.Uint16:
		g.setEnvironmentConfigValue(names, field, g.uint16)
	case reflect.Uint32:
		g.setEnvironmentConfigValue(names, field, g.uint32)
	case reflect.Uint64:
		g.setEnvironmentConfigValue(names, field, g.uint64)
	case reflect.Uintptr:
		g.setEnvironmentConfigValue(names, field, g.uintPtr)
	case reflect.Float32:
		g.setEnvironmentConfigValue(names, field, g.float32)
	case reflect.Float64:
		g.setEnvironmentConfigValue(names, field, g.float64)
	case reflect.String:
		g.setEnvironmentConfigValue(names, field, g.string)
	default:
		g.setEnvironmentConfigValue(names, field, g.string)
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
	} else if !g.params.Nullable {
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

func (g *Getter) setEnvironmentConfigValue(names gox.Slice[string], field reflect.Value, convert callback.Convert) {
	// 将所有名称转换为大写，符合环境变量的定义
	keys := make([]string, names.Length()) // 复制一份，不影响原来的字段名
	for index, name := range names {
		keys[index] = strings.ToUpper(name)
	}

	key := strings.Join(keys, constant.Underline)
	if environment, ok := os.LookupEnv(key); !ok {
		// TODO 记录日志
	} else if ce := convert(environment, field); nil != ce {
		// TODO 记录日志
	}
}

func (g *Getter) zero(field reflect.Value) bool {
	return reflect.DeepEqual(reflect.Zero(field.Type()).Interface(), field.Interface())
}

func (g *Getter) clone(from []string) (to []string) {
	to = make([]string, len(from))
	copy(to, from)

	return
}

func (g *Getter) bool(from string, field reflect.Value) (err error) {
	if value, pbe := strconv.ParseBool(from); nil == pbe {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pbe
	}

	return
}

func (g *Getter) float32(from string, field reflect.Value) (err error) {
	if value, pfe := strconv.ParseFloat(from, 32); nil == pfe {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pfe
	}

	return
}

func (g *Getter) float64(from string, field reflect.Value) (err error) {
	if value, pfe := strconv.ParseFloat(from, 64); nil == pfe {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pfe
	}

	return
}

func (g *Getter) int(from string, field reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, strconv.IntSize); nil == pie {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pie
	}

	return
}

func (g *Getter) int8(from string, field reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 8); nil == pie {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pie
	}

	return
}

func (g *Getter) int16(from string, field reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 16); nil == pie {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pie
	}

	return
}

func (g *Getter) int32(from string, field reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 32); nil == pie {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pie
	}

	return
}

func (g *Getter) int64(from string, field reflect.Value) (err error) {
	var value any
	switch field.Interface().(type) {
	case time.Duration:
		value, err = time.ParseDuration(from)
	case gox.Bytes:
		value, err = gox.ParseBytes(from)
	default:
		value, err = strconv.ParseInt(from, 0, 64)
	}
	if nil != err {
		field.Set(reflect.ValueOf(value))
	}

	return
}
func (g *Getter) uint(from string, field reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, strconv.IntSize); nil == pie {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pie
	}

	return
}

func (g *Getter) uint8(from string, field reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 8); nil == pie {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pie
	}

	return
}

func (g *Getter) uint16(from string, field reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 16); nil == pie {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pie
	}

	return
}

func (g *Getter) uint32(from string, field reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 32); nil == pie {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pie
	}

	return
}

func (g *Getter) uint64(from string, field reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 64); nil == pie {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pie
	}

	return
}

func (g *Getter) uintPtr(from string, field reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, strconv.IntSize); nil == pie {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pie
	}

	return
}

func (g *Getter) string(from string, field reflect.Value) (err error) {
	field.Set(reflect.ValueOf(from).Convert(field.Type()))

	return
}
