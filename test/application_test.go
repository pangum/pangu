package test

import (
	"testing"

	"github.com/pangum/core"
)

func TestApplication(t *testing.T) {
	core.New().Name("test").Banner().Ascii("Test").Build().Get().Run(NewBootstrap)
}
