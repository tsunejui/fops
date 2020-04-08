package providers

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetChecksumCommand(t *testing.T) {
	assert.Equal(t, "*checksum.command", reflect.TypeOf(GetChecksumCommand()).String())
}

func TestGetVersionCommand(t *testing.T) {
	assert.Equal(t, "*version.command", reflect.TypeOf(GetVersionCommand()).String())
}

func TestGetLineCountCommand(t *testing.T) {
	assert.Equal(t, "*linecount.command", reflect.TypeOf(GetLineCountCommand()).String())
}

func TestGetHelpCommand(t *testing.T) {
	assert.Equal(t, "*help.command", reflect.TypeOf(GetHelpCommand()).String())
}