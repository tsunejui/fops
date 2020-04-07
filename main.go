package main

import (
	"fops/commands/root"
	"fops/expection"
	"fops/handler"
	"fops/providers"
	"fops/service"
	"log"
	"os"
)

func main() {
	defer expection.UnexpectedHandle()
	if initSystemErr := service.InitialSystemService(); initSystemErr != nil {
		log.Println("Failed to initial system service", initSystemErr)
	}
	if initRootErr := root.InitialRootCommand(); initRootErr != nil {
		log.Println("Failed to initial root command", initRootErr)
	}
	if initHandlerErr := handler.InitialCommandHandler(); initHandlerErr != nil {
		log.Println("Failed to initial command handler", initHandlerErr)
	}
	if handlerErr := handler.GetCommandHandler().
		SetRootCommand(root.GetRootCommand()).
		AddCommand(providers.GetHelpCommand()).
		AddCommand(providers.GetVersionCommand()).
		AddCommand(providers.GetChecksumCommand()).
		AddCommand(providers.GetLineCountCommand()).
		Do(os.Args); handlerErr != nil {
		log.Println("Failed to parse all sub commands", handlerErr)
	}
	defer service.GetSystemService().Shutdown()
}