package core

import (
	"os"
	"strings"

	"github.com/goexl/gox"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/internal/config/internal/core/internal"
	"github.com/pangum/pangu/internal/runtime"
)

type Environment struct {
	// 无字段
}

func NewEnvironment() *Environment {
	return new(Environment)
}

func (e *Environment) Process(config runtime.Pointer) (err error) {
	environments := os.Environ()
	values := make(map[string]any, len(environments))
	for _, environment := range environments { // 将数据形成映射值
		e.fill(&values, environment)
	}
	err = internal.NewDecoder(&values).Decode(config)

	return
}

func (e *Environment) fill(data *map[string]any, environment string) {
	values := strings.Split(environment, constant.StringEqual)
	if 2 != len(values) {
		return
	}

	keys := gox.String(values[0]).Split().Named().Build().Apply()
	length := len(keys)
	current := *data
	for index := 0; index < length; index++ {
		field := keys[index]
		if value, cached := current[field]; length-1 > index && !cached { // 如果是第一处遇到键，初始化并切换处理对象
			current[field] = make(map[string]any)
			current = current[field].(map[string]any) // !切换当前处理对象
		} else if length-1 > index && cached { // 如果之前初始化，切换处理对象
			current = value.(map[string]any) // !切换当前处理对象
		} else if length-1 == index { // 处理最后一位键值
			current[field] = values[1]
		}
	}

	return
}
