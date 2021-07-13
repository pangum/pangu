package pangu

// Author 作者
type Author struct {
	// 名字
	Name string `json:"name" yaml:"name" xml:"name"`
	// 邮箱
	Email string `json:"email" yaml:"email" xml:"email"`
}
