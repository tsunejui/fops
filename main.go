package main

import (
	"fops/cmd"
	"fops/pkg/exception"
	"fops/pkg/service"
	"log"
	"os"
)

func main() {
	defer exception.UnexpectedHandle()
	if initSystemErr := service.InitialSystemService(); initSystemErr != nil {
		log.Println("Failed to initial system service", initSystemErr)
	}
	if executeErr := cmd.Execute(); executeErr != nil {
		os.Exit(1)
	}
	defer service.GetSystemService().Shutdown()
}
