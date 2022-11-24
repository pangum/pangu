package pangu

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
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

type defaultConfigLoader struct {
	options *configOptions
	data    []byte
}

func newDefaultConfigLoader(options *configOptions) *defaultConfigLoader {
	return &defaultConfigLoader{
		options: options,
		data:    []byte(""),
	}
}

func (dcl *defaultConfigLoader) Load(path string, config any) (err error) {
	if _, exist := gfx.Exists(path); !exist && !dcl.options.nullable {
		err = exc.NewField("配置文件不存在", field.New("path", path))
	} else if exist {
		dcl.data, err = os.ReadFile(path)
	}
	if nil != err {
		return
	}

	// 处理环境变量，不能修改原始数据，复制一份原始数据做修改
	var _data string
	if _data, err = envsubst.Eval(string(dcl.data), dcl.options.environmentGetter); nil != err {
		return
	}

	switch strings.ToLower(filepath.Ext(path)) {
	case ymlExt:
		fallthrough
	case yamlExt:
		err = yaml.Unmarshal([]byte(_data), config)
	case jsonExt:
		err = json.Unmarshal([]byte(_data), config)
	case tomlExt:
		err = toml.Unmarshal([]byte(_data), config)
	case xmlExt:
		err = xml.Unmarshal([]byte(_data), config)
	default:
		err = yaml.Unmarshal([]byte(_data), config)
	}
	if nil != err {
		return
	}

	return
}
