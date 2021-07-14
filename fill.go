package pangu

import (
	`encoding/json`

	`github.com/mcuadros/go-defaults`
	`github.com/storezhang/validatorx`
)

// Fill 填充数据
func Fill(obj interface{}, data []byte) (err error) {
	// 反序列化
	if err = json.Unmarshal(data, obj); nil != err {
		return
	}
	// 处理默认值
	defaults.SetDefaults(obj)
	// 数据验证
	err = validatorx.Struct(obj)

	return
}
