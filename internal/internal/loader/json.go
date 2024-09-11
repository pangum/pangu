package loader

import (
	"context"
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/pangum/config"
	"github.com/pangum/pangu/internal/runtime"
)

var _ config.Loader = (*Json)(nil)

type Json struct {
	targets map[runtime.Pointer]bool
}

func NewJson() *Json {
	return &Json{
		targets: make(map[runtime.Pointer]bool),
	}
}

func (*Json) Local() bool {
	return true
}

func (j *Json) Load(ctx context.Context, target runtime.Pointer) (loaded bool, err error) {
	if path, pok := ctx.Value(config.ContextFilepath).(string); !pok {
		err = exception.New().Message("未指定配置文件路径").Field(field.New("loader", "json")).Build()
	} else if bytes, bok := ctx.Value(config.ContextBytes).([]byte); !bok {
		err = exception.New().Message("配置文件无内容").Field(field.New("loader", "json")).Build()
	} else {
		loaded, err = j.load(&path, &bytes, target)
	}

	return
}

func (j *Json) load(path *string, bytes *[]byte, target runtime.Pointer) (loaded bool, err error) {
	loadable := false
	if ".json" == strings.ToLower(filepath.Ext(*path)) {
		loadable = true
		err = json.Unmarshal(*bytes, target)
	}
	if nil == err && loadable {
		loaded = true
	}

	return
}
