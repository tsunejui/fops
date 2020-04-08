package linecount

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"fops/commands"
	"fops/service"
	"os"
	"path/filepath"
	"rsc.io/getopt"
)
var getFile = service.GetFile
type command struct {
	commands.SubCommand
	path string
}
func (c *command) Initial () error {
	c.GetField().StringVar(&c.path, "file", "", " the input file")
	c.GetField().Alias("f", "file")
	return nil
}
func (c *command) Handle (args []string) error {
	if len(c.path) > 0 {
		file, _, fileErr := getFile(c.path)
		if fileErr != nil {
			fmt.Println(fileErr)
			return nil
		}
		//edge case
		if filepath.Base(c.path) == filepath.Base(os.Args[0]) {
			fmt.Printf("error: Cannot do linecount for binary file'%s'", filepath.Base(c.path))
			return nil
		}

		fmt.Println(getFileLineCount(file))
	}else{
		return errors.New("Missing Argument")
	}
	return nil
}
func GetCommand() commands.CommandInterface {
	linecount := getopt.NewFlagSet("linecount", flag.ExitOnError)
	return &command{
		SubCommand: commands.SubCommand{
			Flag: linecount,
			CommandName: "linecount",
			Description: "Print line count of file",
		},
	}
}

func getFileLineCount(file *os.File) int {
	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return len(result)
}