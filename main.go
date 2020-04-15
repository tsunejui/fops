package main

import (
	"fops/cmd"
	"fops/expection"
	"fops/service"
	"log"
)

func main() {
	defer expection.UnexpectedHandle()
	if initSystemErr := service.InitialSystemService(); initSystemErr != nil {
		log.Println("Failed to initial system service", initSystemErr)
	}
	if executeErr := cmd.Execute(); executeErr != nil {
		//log.Println("Failed to execute root command:", executeErr)
	}
	defer service.GetSystemService().Shutdown()
}