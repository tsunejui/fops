package checksum

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

func TestChecksumRUN(t *testing.T) {
	var mockFilePath string = "test"
	var mockFIle []byte
	var mockCmd *cobra.Command
	defer gostub.StubFunc(&getFileArgs, mockFilePath, nil).Reset()
	defer gostub.StubFunc(&getFileData, mockFIle, nil).Reset()
	defer gostub.StubFunc(&getFlagName, true, nil).Reset()
	initErr := initialAlgorithm()
	assert.NoError(t, (&checksum{
		subcmd.SubCommand{
			ShortName: "",
			Usage:     "",
		},
	}).RUN(mockCmd, []string{}))
	assert.NoError(t, initErr)
}

func TestChecksumSetFlags(t *testing.T) {
	mockCmd := &cobra.Command{}
	assert.NoError(t, (&checksum{
		subcmd.SubCommand{
			ShortName: "",
			Usage:     "",
		},
	}).SetFlags(mockCmd))
}

func TestChecksumSetRequired(t *testing.T) {
	defer gostub.StubFunc(&markFlag, nil).Reset()
	cmd := &cobra.Command{}
	assert.NoError(t, (&checksum{
		subcmd.SubCommand{
			ShortName: "",
			Usage:     "",
		},
	}).SetRequired(cmd))
}