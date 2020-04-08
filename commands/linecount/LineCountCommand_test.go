package linecount

import (
	"flag"
	"fops/commands"
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"rsc.io/getopt"
	"testing"
)

func TestGetCommand(t *testing.T) {
	assert.Equal(t, "*linecount.command", reflect.TypeOf(GetCommand()).String())
}

func TestInitial(t *testing.T) {
	assert.NoError(t, (&command{
		SubCommand: commands.SubCommand{
			Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "",
		},
	}).Initial())
}

func TestHandleHasPath(t *testing.T) {
	mockFile, mockFileErr := os.Create("test")
	var info os.FileInfo
	defer gostub.StubFunc(&getFile, mockFile, info, nil).Reset()
	assert.NoError(t, (&command{
		SubCommand: commands.SubCommand{
			Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "",
		},
		path: "test",
	}).Handle([]string{"test"}))
	assert.NoError(t, mockFileErr)
}

func TestHandleHasNotPath(t *testing.T) {
	mockFile, mockFileErr := os.Create("test")
	var info os.FileInfo
	defer gostub.StubFunc(&getFile, mockFile, info, nil).Reset()
	assert.Error(t, (&command{
		SubCommand: commands.SubCommand{
			Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "",
		},
	}).Handle([]string{"test"}))
	assert.NoError(t, mockFileErr)
}