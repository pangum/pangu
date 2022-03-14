package pangu

var (
	_              = Append
	_              = Override
	_ configOption = (*optionConfigType)(nil)
)

type optionConfigType struct {
	typ configType
}

// Append 追加
func Append() *optionConfigType {
	return &optionConfigType{
		typ: ConfigTypeAppend,
	}
}

// Override 覆盖
func Override() *optionConfigType {
	return &optionConfigType{
		typ: ConfigTypeOverride,
	}
}

func (ct *optionConfigType) applyConfig(options *configOptions) {
	options.typ = ct.typ
}
