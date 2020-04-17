// +build integration

package exception_test

import (
	"fops/pkg/exception"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

func TestPanicHandle(t *testing.T) {
	if os.Getenv("TEST_PANIC_HANDLE") == "1" {
		func() {
			defer exception.PanicHandle()
			panic("test")
		}()
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestPanicHandle")
	cmd.Env = append(os.Environ(), "TEST_PANIC_HANDLE=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestPanicHandleNotPanic(t *testing.T) {
	assert.NotPanics(t, exception.PanicHandle)
}

func TestUnexpectedHandle(t *testing.T) {
	assert.NotPanics(t, func() {
		defer exception.UnexpectedHandle()
		panic("test")
	})
}