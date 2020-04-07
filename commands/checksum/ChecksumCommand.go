package checksum

import (
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"flag"
	"fmt"
	"fops/commands"
	"fops/service"
	"io/ioutil"
	"rsc.io/getopt"
	"crypto/md5"
)

type command struct {
	commands.SubCommand
}
var path string
var md5Algorithm bool
var sha1Algorithm bool
var sha256Algorithm bool
func (c *command) Initial () {
	c.Flag.StringVar(&path, "file", "", " the input file")
	c.Flag.Alias("f", "file")
	c.Flag.BoolVar(&md5Algorithm,"md5",false, "generate md5 checksum")
	c.Flag.BoolVar(&sha1Algorithm,"sha1",false, "generate sha1 checksum")
	c.Flag.BoolVar(&sha256Algorithm,"sha256",false, "generate sha256 checksum")
}
func (c *command) Handle (args []string) error {
	if len(path) > 0 {
		_, _, fileErr := service.GetFile(path)
		if fileErr != nil {
			fmt.Println(fileErr)
		}
		data, readErr := ioutil.ReadFile(path)
		if readErr != nil {
			return errors.New("Failed to read file: " + readErr.Error())
		}
		switch true {
		case md5Algorithm:
			fmt.Println(fmt.Sprintf("%x", md5.Sum(data)))
			break
		case sha1Algorithm:
			fmt.Println(fmt.Sprintf("%x", sha1.Sum(data)))
			break
		case sha256Algorithm:
			fmt.Println(fmt.Sprintf("%x", sha256.Sum256(data)))
			break
		default:
			return errors.New("Invalid Argument")
		}
	}else{
		return errors.New("Missing Argument")
	}
	return nil
}
func GetCommand() commands.CommandInterface {
	checksum := getopt.NewFlagSet("checksum", flag.ExitOnError)
	return &command{
		commands.SubCommand{
			Flag: checksum,
			CommandName: "checksum",
			Description: "Print checksum of file",
		},
	}
}