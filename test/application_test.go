package test

import (
	"testing"

	"github.com/pangum/pangu"
)

func TestApplication(t *testing.T) {
	pangu.New().Get().Run(NewBootstrap())
}
