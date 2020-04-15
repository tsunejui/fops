package linecount

import (
	"bytes"
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
	var info os.FileInfo
	defer gostub.StubFunc(&getFile, &os.File{}, info, nil).Reset()
	defer gostub.StubFunc(&getFileLineCount, 0).Reset()
	assert.NoError(t, (&command{
		SubCommand: commands.SubCommand{
			Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "",
		},
		path: "test",
	}).Handle([]string{"test"}))
}

func TestHandleHasNotPath(t *testing.T) {
	var info os.FileInfo
	defer gostub.StubFunc(&getFile, &os.File{}, info, nil).Reset()
	assert.Error(t, (&command{
		SubCommand: commands.SubCommand{
			Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "",
		},
	}).Handle([]string{"test"}))
}

func TestGetFileLineCount(t *testing.T) {
	assert.Equal(t, "int", reflect.TypeOf(getFileLineCount(bytes.NewBufferString("foo\nbar"))).String())
}