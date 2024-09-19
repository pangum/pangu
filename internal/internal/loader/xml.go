package loader

import (
	"context"
	"encoding/xml"
	"path/filepath"
	"strings"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/pangum/config"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/runtime"
)

var _ config.Loader = (*Xml)(nil)

type Xml struct {
	targets map[runtime.Pointer]bool
}

func NewXml() *Xml {
	return &Xml{
		targets: make(map[runtime.Pointer]bool),
	}
}

func (*Xml) Local() bool {
	return true
}

func (*Xml) Extensions() []string {
	return []string{
		constant.ExtensionXml,
	}
}

func (x *Xml) Load(ctx context.Context, target *map[string]any) (loaded bool, err error) {
	if path, pok := ctx.Value(config.ContextFilepath).(string); !pok {
		err = exception.New().Message("未指定配置文件路径").Field(field.New("loader", "xml")).Build()
	} else if bytes, bok := ctx.Value(config.ContextBytes).([]byte); !bok {
		err = exception.New().Message("配置文件无内容").Field(field.New("loader", "xml")).Build()
	} else {
		loaded, err = x.load(&path, &bytes, target)
	}

	return
}

func (x *Xml) load(path *string, bytes *[]byte, target *map[string]any) (loaded bool, err error) {
	loadable := false
	if constant.ExtensionXml == strings.ToLower(filepath.Ext(*path)) {
		loadable = true
		err = xml.Unmarshal(*bytes, target)
	}
	if nil == err && loadable {
		loaded = true
	}

	return
}
