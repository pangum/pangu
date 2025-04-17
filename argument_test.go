package boot_test

import (
	"testing"

	"github.com/harluo/boot"
	"github.com/stretchr/testify/assert"
)

func TestNewArgument(t *testing.T) {
	target := ""
	argument := boot.NewArgument("test", &target).Build()
	assert.NotNil(t, argument)
}
