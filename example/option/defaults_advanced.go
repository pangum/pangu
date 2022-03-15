package option

var _ = (*plugin)(nil)

type plugin struct {
	// 目录
	Folder string `default:"${PLUGIN_FOLDER=${FOLDER=.}}" validate:"required"`
	// 脚本列表
	Scripts []string `default:"${PLUGIN_SCRIPTS=${SCRIPTS=['build']}}" validate:"required,dive"`
}
