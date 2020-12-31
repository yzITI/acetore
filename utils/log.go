package utils

import (
	"log"
	"os"
)

func init() {
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
}

// Log ...
func Log(msg string) {
	log.Println(msg)
}
