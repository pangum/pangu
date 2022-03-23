package pangu

const (
	exceptionBootstrapMustHasDependencies = `启动器构造方法必须有依赖项`
	exceptionBoostrapMustReturnBootstrap  = `启动器构造方法必须返回pangu.Bootstrap`

	exceptionConstructorMustReturn = `构造方法必须有返回值`
	exceptionConstructorMustFunc   = `构造方法必须是方法不能是其它类型`
)
