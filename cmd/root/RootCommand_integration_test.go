// +build integration

package root_test

import (
	"fops/cmd/root"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetRootCommand(t *testing.T) {
	cmd, err := root.GetRootCommand()
	assert.Equal(t, "*cobra.Command", reflect.TypeOf(cmd).String())
	assert.NoError(t, err)
}

func TestSetRootCommand(t *testing.T) {
	cmd := &cobra.Command{}
	setErr := root.SetRootCommand(cmd)
	rootCmd, getErr := root.GetRootCommand()
	assert.NoError(t, setErr)
	assert.NoError(t, getErr)
	assert.Equal(t, rootCmd, cmd)
}