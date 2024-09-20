package core

import (
	"context"
	"os"
	"reflect"

	"github.com/drone/envsubst"
	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/log"
	"github.com/pangum/config"
	"github.com/pangum/pangu/internal/internal/config/internal/core/internal"
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
)

type Loader struct {
	paths   *gox.Slice[string]
	params  *param.Config
	targets []runtime.Pointer
	logger  *log.Logger
}

func NewLoader(paths *gox.Slice[string], params *param.Config, logger *log.Logger) *Loader {
	return &Loader{
		paths:   paths,
		params:  params,
		targets: make([]runtime.Pointer, 0),
		logger:  logger,
	}
}

func (l *Loader) Load(target runtime.Pointer) (err error) {
	for _, path := range *l.paths {
		err = l.load(path, target)
	}

	return
}

func (l *Loader) load(path string, target runtime.Pointer) (err error) {
	if ctx, has, lce := l.loadLocalContext(path); nil != lce {
		err = lce
	} else {
		err = l.fill(ctx, target, has)
	}

	return
}

func (l *Loader) Wrote() {
	if 0 == len(l.params.Changers) {
		return
	}

	for _, target := range l.targets {
		newTarget := reflect.New(reflect.TypeOf(target)).Elem().Interface()
		if le := l.Load(newTarget); nil != le {
			// todo
		} else if !reflect.DeepEqual(target, newTarget) { // 如果配置有变化
			// todo
		}
	}
}

func (l *Loader) loadLocalContext(path string) (ctx context.Context, populated bool, err error) {
	if _, se := os.Stat(path); nil != se && os.IsNotExist(se) && !l.params.Nullable { // 没有配置文件
		err = exception.New().Message("缺少配置文件").Build()
	} else if bytes, rfe := os.ReadFile(path); nil != rfe {
		err = rfe
	} else if eval, ee := envsubst.Eval(string(bytes), l.params.EnvironmentGetter); nil != ee {
		err = ee
	} else {
		ctx = context.Background()
		ctx = context.WithValue(ctx, config.ContextFilepath, path)
		ctx = context.WithValue(ctx, config.ContextBytes, []byte(eval))
		populated = 0 != len(bytes)
	}

	return
}

func (l *Loader) fill(localContext context.Context, target runtime.Pointer, populated bool) (err error) {
	networkContext := context.Background()
	modules := l.modules(target)
	for _, loader := range l.params.Loaders {
		if loader.Local() && !populated {
			continue
		}

		ctx := localContext
		if !loader.Local() { // 默认为本地上下文，如果为网络加载器，切换为网络上下文
			ctx = networkContext
		}

		value := make(map[string]any)
		if loaded, le := loader.Load(ctx, &value, modules); nil != le {
			err = le
		} else if loaded && 0 != len(value) { // 确实加载了配置数据
			err = internal.NewDecoder(&value).Decode(target)
		}

		if nil != err {
			break
		}
	}

	if nil == err {
		l.targets = append(l.targets, target)
	}

	return
}

func (l *Loader) modules(target runtime.Pointer) (modules []string) {
	typeOfTarget := reflect.TypeOf(target)
	if reflect.Ptr == typeOfTarget.Kind() {
		typeOfTarget = typeOfTarget.Elem()
	}
	if 1 != typeOfTarget.NumField() {
		return
	}

	names := make(map[string]*gox.Empty)
	name := typeOfTarget.Field(0).Name

	swh := gox.String(name).Switch()
	names[swh.Strike().Build().Case()] = nil
	names[swh.Underscore().Build().Case()] = nil
	names[swh.Camel().Build().Case()] = nil

	modules = make([]string, 0, len(names))
	for key := range names {
		modules = append(modules, key)
	}

	return
}
