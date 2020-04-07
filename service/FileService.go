package service

import (
	"errors"
	"log"
	"os"
)

func GetFile(path string) (*os.File, os.FileInfo, error) {
	var info os.FileInfo
	var infoErr error
	file, openErr := os.Open(path)
	if openErr != nil {
		if os.IsNotExist(openErr) {
			return file, info, errors.New("error: No such file '" + path +  "'")
		}
	}
	if addError := service.AddPreload(func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Println("Failed to close file", closeErr)
		}
	}); addError != nil {
		return file, info, addError
	}
	info, infoErr = file.Stat()
	if infoErr != nil {
		return file, info, errors.New("Failed to get file info: " + infoErr.Error())
	}
	if info.Mode().IsDir() {
		return file, info, errors.New("error: Expected file got directory '" + path + "'")
	}
	return file, info, nil
}