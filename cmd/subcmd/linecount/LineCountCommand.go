package linecount

import (
	"bufio"
	"errors"
	"fmt"
	"fops/cmd/subcmd"
	"fops/service"
	"github.com/spf13/cobra"
	"io"
)
var getFile = service.GetFile
var usage string = "linecount"
var shortName string = "Print line count of file"
type checksum struct {
	subcmd.SubCommand
}
func GetSubCommand() (*cobra.Command, error){
	return subcmd.GetCommand(&checksum{subcmd.SubCommand{
		Usage:     usage,
		ShortName: shortName,
	}})
}
func (c *checksum) SetRequired(cmd *cobra.Command) error {
	if markErr := cmd.MarkFlagRequired("file"); markErr != nil {
		return markErr
	}
	return nil
}
func (c *checksum) SetFlags(cmd *cobra.Command) error {
	cmd.Flags().StringP("file", "f", "", "the input file")
	return nil
}
func (c *checksum) RUN (cmd *cobra.Command, args []string) error {
	filepath, pathErr := cmd.Flags().GetString("file")
	if pathErr != nil {
		return pathErr
	}
	if len(filepath) > 0 {
		file, _, fileErr := getFile(filepath)
		if fileErr != nil {
			fmt.Println(fileErr)
			return nil
		}
		fmt.Println(getFileLineCount(file))
	}else{
		return errors.New("Missing Argument")
	}
	return nil
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