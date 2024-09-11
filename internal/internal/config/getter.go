package config

import (
	"context"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/drone/envsubst"
	"github.com/goexl/exception"
	"github.com/goexl/gfx"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/goexl/mengpo"
	"github.com/goexl/xiren"
	"github.com/pangum/config"
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

	logger *log.Logger
	once   *sync.Once
}

func newGetter(params *param.Config, logger *log.Logger) *Getter {
	return &Getter{
		params: params,

		logger: logger,
		once:   new(sync.Once),
	}
}

func (g *Getter) Get(target runtime.Pointer) (err error) {
	if path, fpe := g.filepath(); nil != fpe { // 获取最终的配置文件路径
		err = fpe
	} else if fe := g.Fill(path, target); nil != fe { // 加载数据
		err = fe
	} else { // 保存已经使用过的配置文件路径
		g.path = path
	}

	return
}

func (g *Getter) Fill(path string, target runtime.Pointer) (err error) {
	if le := g.load(path, target); nil != le { // 从路径中加载数据
		err = le
	}

	if nil == err && g.params.Default { // 处理默认值
		// !此处逻辑不能往前，原因是如果对象里面包含指针，那么只能在包含指针的结构体被解析后才能去设置默认值，不然指针将被会设置成空值
		err = mengpo.New().Tag(g.params.Tag.Default).Build().Set(target)
	}

	// 从环境变量中加载配置
	g.processEnvironmentConfig(reflect.ValueOf(target).Elem(), gox.NewSlice[string](), g.setEnvironmentConfig)

	if nil == err && g.params.Validate { // 数据验证
		err = xiren.Struct(target)
	}

	return
}

func (g *Getter) load(path string, target runtime.Pointer) (err error) {
	if _, se := os.Stat(path); nil != se && os.IsNotExist(se) && g.params.Nullable { // 允许不使用配置文件
		// 空实现，纯占位
	} else if le := g.loadFromContext(path, target); nil != le { // 从路径中加载数据
		err = le
	}

	return
}

func (g *Getter) loadFromContext(path string, target runtime.Pointer) (err error) {
	localContext := context.Background()
	if bytes, rfe := os.ReadFile(path); nil != rfe {
		err = rfe
	} else if evaled, ee := envsubst.Eval(string(bytes), g.params.EnvironmentGetter); nil != ee {
		err = ee
	} else {
		localContext = context.WithValue(localContext, config.ContextFilepath, path)
		localContext = context.WithValue(localContext, config.ContextBytes, []byte(evaled))
	}
	if nil != err {
		return
	}

	networkContext := context.Background()
	for _, loader := range g.params.Loaders {
		ctx := localContext
		if !loader.Local() { // 默认为本地上下文，如果确实为网络加载器，切换为网络上下文
			ctx = networkContext
		}
		if _, le := loader.Load(ctx, target); nil != le {
			err = le
		}

		if nil != err {
			break
		}
	}

	return
}

func (g *Getter) processEnvironmentConfig(config reflect.Value, names gox.Slice[string], set callback.SetValue) {
	typ := config.Type()
	for index := 0; index < config.NumField(); index++ {
		target := config.Field(index)
		name := typ.Field(index).Name
		kind := target.Kind()

		if reflect.Struct == kind { // 如果字段为结构体才递归处理
			names = append(names, name)
			g.processEnvironmentConfig(target, names, set)
		} else if reflect.Ptr == kind && target.Elem().Kind() == reflect.Struct { // 如果是结构体指针才进行递归处理
			if target.IsNil() { // 仅对空指针进行初始化
				target.Set(reflect.New(target.Type().Elem()))
			}
			names = append(names, name)
			g.processEnvironmentConfig(target.Elem(), names, set)
		} else { // 不管是否被设置，都尝试从环境变量中设置变量值
			copies := names.Clone() // !复制一份字段名列表，防止干扰其它结构体字段
			set(append(copies, name), target)
		}
	}
}

func (g *Getter) setEnvironmentConfig(names gox.Slice[string], target reflect.Value) {
	switch target.Kind() {
	case reflect.Bool:
		g.setEnvironmentConfigValue(names, target, g.bool)
	case reflect.Int:
		g.setEnvironmentConfigValue(names, target, g.int)
	case reflect.Int8:
		g.setEnvironmentConfigValue(names, target, g.int8)
	case reflect.Int16:
		g.setEnvironmentConfigValue(names, target, g.int16)
	case reflect.Int32:
		g.setEnvironmentConfigValue(names, target, g.int32)
	case reflect.Int64:
		g.setEnvironmentConfigValue(names, target, g.int64)
	case reflect.Uint:
		g.setEnvironmentConfigValue(names, target, g.uint)
	case reflect.Uint8:
		g.setEnvironmentConfigValue(names, target, g.uint8)
	case reflect.Uint16:
		g.setEnvironmentConfigValue(names, target, g.uint16)
	case reflect.Uint32:
		g.setEnvironmentConfigValue(names, target, g.uint32)
	case reflect.Uint64:
		g.setEnvironmentConfigValue(names, target, g.uint64)
	case reflect.Uintptr:
		g.setEnvironmentConfigValue(names, target, g.uintPtr)
	case reflect.Float32:
		g.setEnvironmentConfigValue(names, target, g.float32)
	case reflect.Float64:
		g.setEnvironmentConfigValue(names, target, g.float64)
	case reflect.String:
		g.setEnvironmentConfigValue(names, target, g.string)
	default:
		g.setEnvironmentConfigValue(names, target, g.string)
	}
}

func (g *Getter) filepath() (path string, err error) {
	if "" != g.path {
		path = g.path
	} else {
		path, err = g.detectFilepath()
	}

	return
}

func (g *Getter) detectFilepath() (path string, err error) {
	exists := gfx.Exists().Filename(constant.ApplicationName)
	if constant.EnvironmentNotSet != internal.Name { // 如果配置了应用名称，可以使用应用名称的配置文件
		exists.Filename(internal.Name)
	}
	// 配置所有可能的配置目录
	exists.Directory("config")
	exists.Directory("conf")
	exists.Directory("configuration")

	if final, checked := exists.Build().Check(); checked {
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
	config.Usage = "指定配置文件路径"
	config.Destination = &g.path

	shell.Flags = append(shell.Flags, config)
	shadow.Flags = append(shadow.Flags, config)
}

func (g *Getter) setEnvironmentConfigValue(names gox.Slice[string], target reflect.Value, convert callback.Convert) {
	// 将所有名称转换为大写，符合环境变量的定义
	keys := make([]string, names.Length()) // 复制一份，不影响原来的字段名
	for index, name := range names {
		keys[index] = strings.ToUpper(name)
	}

	key := strings.Join(keys, constant.Underline)
	fields := gox.Fields[any]{
		field.New("field", strings.Join(names, constant.Dot)),
		field.New("environment", key),
	}
	if environment, ok := os.LookupEnv(key); !ok {
		(*g.logger).Debug("未在环境变量中发现配置", fields...)
	} else if ce := convert(environment, target); nil != ce {
		(*g.logger).Debug("从环境变量中设置配置值出错", fields.Add(field.Error(ce))...)
	} else {
		(*g.logger).Debug("环境变量中设置配置值", fields...)
	}
}

func (g *Getter) bool(from string, target reflect.Value) (err error) {
	if value, pbe := strconv.ParseBool(from); nil == pbe {
		target.Set(reflect.ValueOf(value))
	} else {
		err = pbe
	}

	return
}

func (g *Getter) float32(from string, target reflect.Value) (err error) {
	if value, pfe := strconv.ParseFloat(from, 32); nil == pfe {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pfe
	}

	return
}

func (g *Getter) float64(from string, target reflect.Value) (err error) {
	if value, pfe := strconv.ParseFloat(from, 64); nil == pfe {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pfe
	}

	return
}

func (g *Getter) int(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, strconv.IntSize); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (g *Getter) int8(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 8); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (g *Getter) int16(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 16); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (g *Getter) int32(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 32); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (g *Getter) int64(from string, target reflect.Value) (err error) {
	var value any
	switch target.Interface().(type) {
	case time.Duration:
		value, err = time.ParseDuration(from)
	case gox.Bytes:
		value, err = gox.ParseBytes(from)
	default:
		value, err = strconv.ParseInt(from, 0, 64)
	}
	if nil != err {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	}

	return
}
func (g *Getter) uint(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, strconv.IntSize); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (g *Getter) uint8(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 8); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (g *Getter) uint16(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 16); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (g *Getter) uint32(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 32); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (g *Getter) uint64(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 64); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (g *Getter) uintPtr(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, strconv.IntSize); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (g *Getter) string(from string, target reflect.Value) (err error) {
	target.Set(reflect.ValueOf(from).Convert(target.Type()))

	return
}
