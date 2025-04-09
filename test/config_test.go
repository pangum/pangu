package test

import (
	"testing"

	"github.com/pangum/pangu"
)

func TestConfigPrefix(t *testing.T) {
	t.Setenv("PANGU_BANNER_DATA", "prefix")
	pangu.New().Name("config.prefix").Get().Run(NewBootstrap)
}
