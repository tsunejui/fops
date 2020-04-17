package linecount

import (
	"bufio"
	"errors"
	"fmt"
	"fops/cmd/subcmd"
	"fops/pkg/service"
	"github.com/gabriel-vasile/mimetype"
	"github.com/spf13/cobra"
	"io"
)
var getFile = service.GetFile
var usage string = "linecount"
var shortName string = "Print line count of file"
var getFileArgs = func(cmd *cobra.Command) (string, error){
	return cmd.Flags().GetString("file")
}
var markFlag = func(cmd *cobra.Command, name string) error {
	return cmd.MarkFlagRequired(name)
}
type linecount struct {
	subcmd.SubCommand
}
func GetSubCommand() (*cobra.Command, error){
	return subcmd.GetCommand(&linecount{subcmd.SubCommand{
		Usage:     usage,
		ShortName: shortName,
	}})
}
func (c *linecount) SetRequired(cmd *cobra.Command) error {
	if markErr := markFlag(cmd, "file"); markErr != nil {
		return markErr
	}
	return nil
}
func (c *linecount) SetFlags(cmd *cobra.Command) error {
	cmd.Flags().StringP("file", "f", "", "the input file")
	return nil
}
func (c *linecount) RUN (cmd *cobra.Command, args []string) error {
	filepath, pathErr := getFileArgs(cmd)
	if pathErr != nil {
		return pathErr
	}
	if len(filepath) > 0 {
		file, _, fileErr := getFile(filepath)
		if fileErr != nil {
			return fileErr
		}
		binaryFile, binaryErr := isBinary(filepath)
		if binaryErr != nil {
			return binaryErr
		}
		if binaryFile {
			return fmt.Errorf("error: Cannot do linecount for binary file '%s' ", file.Name())
		}
		fmt.Println(getFileLineCount(file))
	}else{
		return errors.New("Missing Argument")
	}
	return nil
}

var isBinary = func(filePath string) (bool, error) {
	mime, err := mimetype.DetectFile(filePath)
	if err != nil {
		return false, err
	}
	binary := true
	for mime := mime; mime != nil; mime = mime.Parent() {
		if mime.Is("text/plain") {
			binary = false
		}
	}
	return binary, nil
}

var getFileLineCount = func (file io.Reader) int {
	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return len(result)
}