package loader

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"path/filepath"
	"strings"

	"github.com/drone/envsubst"
	"github.com/goexl/exc"
	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
	"github.com/pangum/pangu/internal/config"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/param"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

var _ config.Loader[int] = (*Config[int])(nil)

type Config[T any] struct {
	params *param.Config
	data   []byte
}

func NewConfig[T any](params *param.Config) *Config[T] {
	return &Config[T]{
		params: params,
		data:   []byte(""),
	}
}

func (c *Config[T]) Load(path string) (value *T, err error) {
	if re := c.read(path); nil != re {
		err = re
	} else if data, ge := envsubst.Eval(string(c.data), c.params.EnvironmentGetter); nil != ge {
		err = ge
	} else {
		value, err = c.load(path, data)
	}

	return
}

func (c *Config[T]) read(path string) (err error) {
	if _, exist := gfx.Exists(path); !exist && !c.params.Nullable {
		err = exc.NewField("配置文件不存在", field.New("path", path))
	} else if exist {
		c.data, err = os.ReadFile(path)
	}

	return
}

func (c *Config[T]) load(path string, content string) (value *T, err error) {
	switch strings.ToLower(filepath.Ext(path)) {
	case constant.ExtensionYml:
		fallthrough
	case constant.ExtensionYaml:
		err = yaml.Unmarshal([]byte(content), value)
	case constant.ExtensionJson:
		err = json.Unmarshal([]byte(content), value)
	case constant.ExtensionToml:
		err = toml.Unmarshal([]byte(content), value)
	case constant.ExtensionXml:
		err = xml.Unmarshal([]byte(content), value)
	default:
		err = yaml.Unmarshal([]byte(content), value)
	}

	return
}
