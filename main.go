package main

import (
	"log"
	"os"
)

func main() {
	if err := New().Run(os.Args); err != nil {
		switch err {
		case ErrNoTasksToExecute, ErrEmptyReport:
			log.Print("[INFO] ", err)
		default:
			log.Fatal("[ERROR] ", err)
		}
	}
}
