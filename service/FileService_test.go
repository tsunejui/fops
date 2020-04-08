package service

import (
	"errors"
	"github.com/prashantv/gostub"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetFileWithNoSuchFileError(t *testing.T) {
	defer gostub.StubFunc(&openFile, &os.File{}, errors.New("test")).Reset()
	defer gostub.StubFunc(&fileIsNotExist, true).Reset()
	_, _, getFileErr := GetFile("test")
	assert.Error(t, getFileErr)
}

func TestGetFileWithOpenError(t *testing.T) {
	defer gostub.StubFunc(&openFile, &os.File{}, errors.New("test")).Reset()
	defer gostub.StubFunc(&fileIsNotExist, false).Reset()
	_, _, getFileErr := GetFile("test")
	assert.Error(t, getFileErr)
}

func TestGetFileWithAddPreloadError(t *testing.T) {
	defer gostub.StubFunc(&openFile, &os.File{}, nil).Reset()
	defer gostub.StubFunc(&addSystemPreload, errors.New("test")).Reset()
	_, _, getFileErr := GetFile("test")
	assert.Error(t, getFileErr)
}

func TestGetFileWithStatError(t *testing.T) {
	var info os.FileInfo
	defer gostub.StubFunc(&openFile, &os.File{}, nil).Reset()
	defer gostub.StubFunc(&addSystemPreload, nil).Reset()
	defer gostub.StubFunc(&getFileStat, info, errors.New("test")).Reset()
	_, _, getFileErr := GetFile("test")
	assert.Error(t, getFileErr)
}

func TestGetFileWithInfoDirError(t *testing.T) {
	var info os.FileInfo
	defer gostub.StubFunc(&openFile, &os.File{}, nil).Reset()
	defer gostub.StubFunc(&addSystemPreload, nil).Reset()
	defer gostub.StubFunc(&getFileStat, info, nil).Reset()
	defer gostub.StubFunc(&infoIsDir, true).Reset()
	_, _, getFileErr := GetFile("test")
	assert.Error(t, getFileErr)
}

func TestGetFileWithNoError(t *testing.T) {
	var info os.FileInfo
	defer gostub.StubFunc(&openFile, &os.File{}, nil).Reset()
	defer gostub.StubFunc(&addSystemPreload, nil).Reset()
	defer gostub.StubFunc(&getFileStat, info, nil).Reset()
	defer gostub.StubFunc(&infoIsDir, false).Reset()
	_, _, getFileErr := GetFile("test")
	assert.NoError(t, getFileErr)
}