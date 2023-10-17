package pangu

// Bootstrapper 基础启动器
type Bootstrapper struct {
	// 用于编写代码时从本结构体继承（匿名内部类）从而少写部分方法
}

func (b *Bootstrapper) Before() (err error) {
	return
}

func (b *Bootstrapper) After() (err error) {
	return
}
