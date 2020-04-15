package checksum

import (
	"flag"
	"fops/commands"
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
	"reflect"
	"rsc.io/getopt"
	"testing"
)

func TestGetCommand(t *testing.T) {
	assert.Equal(t, "*checksum.command", reflect.TypeOf(GetCommand()).String())
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

func TestHandleWithMd5(t *testing.T) {
	defer gostub.StubFunc(&getFileData, []byte{}, nil).Reset()
	assert.NoError(t, (&command{
		SubCommand: commands.SubCommand{
			Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "",
		},
		path: "test",
		md5Algorithm: true,
	}).Handle([]string{}))
}

func TestHandleWithSha1(t *testing.T) {
	defer gostub.StubFunc(&getFileData, []byte{}, nil).Reset()
	assert.NoError(t, (&command{
		SubCommand: commands.SubCommand{
			Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "",
		},
		path: "test",
		sha1Algorithm: true,
	}).Handle([]string{}))
}

func TestHandleWithSha256(t *testing.T) {
	defer gostub.StubFunc(&getFileData, []byte{}, nil).Reset()
	assert.NoError(t, (&command{
		SubCommand: commands.SubCommand{
			Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "",
		},
		path: "test",
		sha256Algorithm: true,
	}).Handle([]string{}))
}

func TestHandleWithNoAlgorithmHasError(t *testing.T) {
	defer gostub.StubFunc(&getFileData, []byte{}, nil).Reset()
	assert.Error(t, (&command{
		SubCommand: commands.SubCommand{
			Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "",
		},
		path: "test",
	}).Handle([]string{}))
}

func TestHandleWithNoOathHasError(t *testing.T) {
	defer gostub.StubFunc(&getFileData, []byte{}, nil).Reset()
	assert.Error(t, (&command{
		SubCommand: commands.SubCommand{
			Flag:        getopt.NewFlagSet("test", flag.ExitOnError),
			Description: "",
			CommandName: "",
		},
	}).Handle([]string{}))
}