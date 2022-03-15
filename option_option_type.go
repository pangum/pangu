package pangu

var (
	_              = Append
	_              = Override
	_ optionOption = (*optionConfigType)(nil)
)

type optionConfigType struct {
	typ optionType
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

func (ct *optionConfigType) applyOption(options *optionOptions) {
	options.typ = ct.typ
}
