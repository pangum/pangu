package conf

type Example struct {
	// 管理后台密码
	Password string `validate:"required"`
}
