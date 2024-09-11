package builder

import (
	"github.com/pangum/pangu/internal/param"
)

type Code struct {
	params      *param.Code
	application *Application
}

func NewCode(application *Application) (code *Code) {
	code = new(Code)
	code.params = application.params.Code
	code.application = application

	return
}

func (c *Code) Success(success int) (code *Code) {
	c.params.Success = success
	code = c

	return
}

func (c *Code) Failed(failed int) (code *Code) {
	c.params.Failed = failed
	code = c

	return
}

func (c *Code) Panic(panic int) (code *Code) {
	c.params.Panic = panic
	code = c

	return
}

func (c *Code) Build() (application *Application) {
	c.application.params.Code = c.params
	application = c.application

	return
}
