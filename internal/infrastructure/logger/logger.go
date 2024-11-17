package logger

import (
	"log"
)

func Init() {
	// You can extend this function to use a more sophisticated logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Logger initialized")
}
