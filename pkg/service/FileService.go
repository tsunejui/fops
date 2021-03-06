package service

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	openFile = os.Open
	fileIsNotExist = os.IsNotExist
	addSystemPreload = func (load func()) error {
		return service.AddPreload(load)
	}
	getFileStat = func(file *os.File) (os.FileInfo, error) {
		return file.Stat()
	}
	infoIsDir = func(info os.FileInfo) bool {
		return info.Mode().IsDir()
	}
)

func GetFile(path string) (*os.File, os.FileInfo, error) {
	var info os.FileInfo
	var infoErr error
	file, openErr := openFile(path)
	if openErr != nil {
		if fileIsNotExist(openErr) {
			return file, info, errors.New(fmt.Sprintf("error: No such file %s", path))
		}
		return file, info, openErr
	}
	if addError := addSystemPreload(func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Println("Failed to close file", closeErr)
		}
	}); addError != nil {
		return file, info, addError
	}
	info, infoErr = getFileStat(file)
	if infoErr != nil {
		return file, info, fmt.Errorf("Failed to get file info:  %s", infoErr.Error())
	}
	if infoIsDir(info) {
		return file, info, fmt.Errorf("error: Expected file got directory %s", path)
	}
	return file, info, nil
}