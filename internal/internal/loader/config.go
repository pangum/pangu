package loader

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"path/filepath"
	"strings"

	"github.com/drone/envsubst"
	"github.com/goexl/exception"
	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
	"github.com/pangum/pangu/internal/callback/getter"
	"github.com/pangum/pangu/internal/config"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/internal/message"
	"github.com/pangum/pangu/internal/runtime"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

var _ config.Loader = (*Config)(nil)

type Config struct {
	getter   getter.Environment
	nullable bool
	data     []byte
}

func NewConfig(getter getter.Environment, nullable bool) *Config {
	return &Config{
		getter:   getter,
		nullable: nullable,
		data:     []byte(""),
	}
}

func (c *Config) Load(path string, config runtime.Pointer) (err error) {
	if re := c.read(path); nil != re {
		err = re
	} else if data, ge := envsubst.Eval(string(c.data), c.getter); nil != ge {
		err = ge
	} else {
		err = c.load(path, data, config)
	}

	return
}

func (c *Config) read(path string) (err error) {
	if _, exist := gfx.Exists(path); !exist && !c.nullable {
		err = exception.New().Message(message.ConfigFileNotfound).Field(field.New("path", path)).Build()
	} else if exist {
		c.data, err = os.ReadFile(path)
	}

	return
}

func (c *Config) load(path string, content string, config runtime.Pointer) (err error) {
	switch strings.ToLower(filepath.Ext(path)) {
	case constant.ExtensionYml:
		fallthrough
	case constant.ExtensionYaml:
		err = yaml.Unmarshal([]byte(content), config)
	case constant.ExtensionJson:
		err = json.Unmarshal([]byte(content), config)
	case constant.ExtensionToml:
		err = toml.Unmarshal([]byte(content), config)
	case constant.ExtensionXml:
		err = xml.Unmarshal([]byte(content), config)
	default:
		err = yaml.Unmarshal([]byte(content), config)
	}

	return
}
