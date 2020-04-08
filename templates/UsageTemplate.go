package templates

import (
	"log"
	"os"
	"text/template"
)

type CommandTemplate struct {
	Name string
	Description string
}

func RootCommandUsage(subcommand []CommandTemplate) {
	tmpl, tmplErr := template.New("rootcommand").Parse(
`File Ops
Usage:
  fops [flags] 
  fops [command] 
Available Commands:
  {{range .Commands}}{{.Name}} {{.Description}}
  {{end}}
Flags:`)
	if tmplErr != nil {
		log.Println("Failed to new template: ", tmplErr)
	}
	if executeErr := tmpl.Execute(os.Stdout, struct {
		Commands []CommandTemplate
	}{
		Commands: subcommand,
	}); executeErr != nil {
		log.Println("Failed to execute template: ", executeErr)
	}
}

func HelpCommandUsage(description string) {
	tmpl, tmplErr := template.New("helpcommand").Parse(
		`{{.Description}}
Usage:
  fops help [command]
`)
	if tmplErr != nil {
		log.Println("Failed to new template: ", tmplErr)
	}
	if executeErr := tmpl.Execute(os.Stdout, struct {
		Description string
	}{
		Description: description,
	}); executeErr != nil {
		log.Println("Failed to execute template: ", executeErr)
	}
}

func SubCommandUsage(description string, commandName string) {
	tmpl, tmplErr := template.New("subcommand").Parse(
		`{{.Description}}
Usage:
  fops {{.Name}} [flags] 
Flags:`)
	if tmplErr != nil {
		log.Println("Failed to new template: ", tmplErr)
	}
	if executeErr := tmpl.Execute(os.Stdout, struct {
		Name string
		Description string
	}{
		Name: commandName,
		Description: description,
	}); executeErr != nil {
		log.Println("Failed to execute template: ", executeErr)
	}
}