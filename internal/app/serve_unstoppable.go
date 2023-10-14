package app

// UnstoppableServe 不需要停止的服务
type UnstoppableServe struct{}

func (us UnstoppableServe) Stop() (err error) {
	return
}
