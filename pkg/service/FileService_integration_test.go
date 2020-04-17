// +build integration

package service_test

import (
	"fmt"
	"fops/pkg/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFile(t *testing.T) {
	initErr := service.InitialSystemService()
	_, _, err := service.GetFile("../../test/myfile.txt")
	assert.NoError(t, initErr)
	assert.NoError(t, err)
}

func TestGetFileWithInvalidPath(t *testing.T) {
	initErr := service.InitialSystemService()
	path := "../../test/test.txt"
	_, _, err := service.GetFile(path)
	assert.NoError(t, initErr)
	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf("error: No such file %s", path), err.Error())
}