// +build unit

package version

import (
	"fops/cmd/subcmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetSubCommand(t *testing.T) {
	cmd, cmdErr := GetSubCommand()
	assert.NoError(t, cmdErr)
	assert.Equal(t, "*cobra.Command", reflect.TypeOf(cmd).String())
}

func TestVersionSetFlags(t *testing.T) {
	cmd := &cobra.Command{}
	assert.NoError(t, (&versionCmd{
		subcmd.SubCommand{
			ShortName: "",
			Usage:     "",
		},
	}).SetFlags(cmd))
}

func TestVersionCmdSetRequired(t *testing.T) {
	cmd := &cobra.Command{}
	assert.NoError(t, (&versionCmd{
		subcmd.SubCommand{
			ShortName: "",
			Usage:     "",
		},
	}).SetRequired(cmd))
}

func TestVersionCmdRUN(t *testing.T) {
	cmd := &cobra.Command{}
	assert.NoError(t, (&versionCmd{
		subcmd.SubCommand{
			ShortName: "",
			Usage:     "",
		},
	}).RUN(cmd, []string{}))
}