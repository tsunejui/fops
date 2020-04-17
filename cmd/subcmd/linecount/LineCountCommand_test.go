// +build unit

package linecount

import (
	"fops/cmd/subcmd"
	"github.com/prashantv/gostub"
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

func TestLineCountRUN(t *testing.T) {
	var mockFilePath string = "test"
	var mockCmd *cobra.Command
	defer gostub.StubFunc(&getFileArgs, mockFilePath, nil).Reset()
	defer gostub.StubFunc(&getFile, nil, nil, nil).Reset()
	defer gostub.StubFunc(&getFileLineCount, 0).Reset()
	defer gostub.StubFunc(&isBinary, false, nil).Reset()
	assert.NoError(t, (&linecount{
		subcmd.SubCommand{
			ShortName: "",
			Usage:     "",
		},
	}).RUN(mockCmd, []string{}))
}
func TestLinecountSetFlags(t *testing.T) {
	mockCmd := &cobra.Command{}
	assert.NoError(t, (&linecount{
		subcmd.SubCommand{
			ShortName: "",
			Usage:     "",
		},
	}).SetFlags(mockCmd))
}
func TestLineCountSetRequired(t *testing.T) {
	defer gostub.StubFunc(&markFlag, nil).Reset()
	cmd := &cobra.Command{}
	assert.NoError(t, (&linecount{
		subcmd.SubCommand{
			ShortName: "",
			Usage:     "",
		},
	}).SetRequired(cmd))
}