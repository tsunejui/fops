package root

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetRootCommand(t *testing.T) {
	cmd, err := GetRootCommand()
	assert.Equal(t, "*cobra.Command", reflect.TypeOf(cmd).String())
	assert.NoError(t, err)
}

func TestSetRootCommand(t *testing.T) {
	cmd := &cobra.Command{}
	setErr := SetRootCommand(cmd)
	rootCmd, getErr := GetRootCommand()
	assert.NoError(t, setErr)
	assert.NoError(t, getErr)
	assert.Equal(t, rootCmd, cmd)
}