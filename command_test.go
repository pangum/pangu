package boot_test

import (
	"testing"

	"github.com/harluo/boot"
	"github.com/stretchr/testify/assert"
)

func TestNewCommand(t *testing.T) {
	command := boot.NewCommand("test")
	assert.NotNil(t, command)
}
