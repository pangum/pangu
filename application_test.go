package boot_test

import (
	"testing"

	"github.com/harluo/boot"
	"github.com/stretchr/testify/assert"
)

func TestNewApplication(t *testing.T) {
	application := boot.New().Get()
	assert.NotNil(t, application)
}
