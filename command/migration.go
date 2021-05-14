package command

// 升级接口
// 属于内部接口，纯粹是为了方便使用而写，无实际用处
// 为了保持不暴露内部实现，也就是不暴露pangu.migration，又要在command包中使用pangu.migration，特意做的接口，保持和pangu.migration一致
type migration interface {
	// Migrate 执行升级
	Migrate() error
}
