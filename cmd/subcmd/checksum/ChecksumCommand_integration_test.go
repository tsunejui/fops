// +build integration

package checksum_test

import (
	"fops/cmd/subcmd/checksum"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetSubCommand(t *testing.T) {
	cmd, cmdErr := checksum.GetSubCommand()
	assert.NoError(t, cmdErr)
	assert.Equal(t, "*cobra.Command", reflect.TypeOf(cmd).String())
}