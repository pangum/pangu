package constant

const (
	ApplicationName        = "application"
	ApplicationStacktrace  = 10
	ApplicationSkip        = 1
	ApplicationCodeSuccess = 0
	ApplicationCodeFailed  = 1
	ApplicationCodePanic   = -1
	ApplicationDefaultName = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.Name=$NAME"来注入值`
)
