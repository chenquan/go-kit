package graceful

import (
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShutdown(t *testing.T) {
	val := 1

	called := AddShutdownListener(func() {
		val += 2
	})
	//Shutdown()
	_ = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)

	called()
	assert.Equal(t, 3, val)
}
