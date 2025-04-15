package test

import (
	"testing"

	"github.com/pangum/core"
)

func TestConfigPrefix(t *testing.T) {
	t.Setenv("PANGU_BANNER_DATA", "prefix")
	t.Setenv("PANGU_CODE_PANIC", "5")
	core.New().Name("config.prefix").Get().Run(NewBootstrap)
}
