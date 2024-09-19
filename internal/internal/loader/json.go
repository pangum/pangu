package loader

import (
	"context"
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/pangum/config"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/internal/loader/internal"
)

var _ config.Loader = (*Json)(nil)

type Json struct {
	jsonc *internal.Jsonc
}

func NewJson() *Json {
	return &Json{
		jsonc: internal.NewJsonc(),
	}
}

func (*Json) Local() bool {
	return true
}

func (*Json) Extensions() []string {
	return []string{
		constant.ExtensionJson,
		constant.ExtensionJsonc,
		constant.ExtensionJson5,
	}
}

func (j *Json) Load(ctx context.Context, target *map[string]any) (loaded bool, err error) {
	if path, pok := ctx.Value(config.ContextFilepath).(string); !pok {
		err = exception.New().Message("未指定配置文件路径").Field(field.New("loader", "json")).Build()
	} else if bytes, bok := ctx.Value(config.ContextBytes).([]byte); !bok {
		err = exception.New().Message("配置文件无内容").Field(field.New("loader", "json")).Build()
	} else {
		loaded, err = j.load(&path, &bytes, target)
	}

	return
}

func (j *Json) load(path *string, bytes *[]byte, target *map[string]any) (loaded bool, err error) {
	loadable := false
	ext := strings.ToLower(filepath.Ext(*path))
	if constant.ExtensionJson5 == ext || constant.ExtensionJsonc == ext {
		to := j.jsonc.Clear(string(*bytes))
		*bytes = []byte(to)
	}

	if constant.ExtensionJson == ext || constant.ExtensionJson5 == ext || constant.ExtensionJsonc == ext {
		loadable = true
		err = json.Unmarshal(*bytes, target)
	}
	if nil == err && loadable {
		loaded = true
	}

	return
}
