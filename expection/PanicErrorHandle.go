package expection

import (
	"log"
	"os"
)

func PanicHandle() {
	if err := recover(); err != nil {
		log.Println("Exception Error", err)
		os.Exit(1)
	}
}

func UnexpectedHandle() {
	if err := recover(); err != nil {
		log.Println("Unexpecte Error", err)
	}
}