package main

import (
	"log"
	"os"
)

func main() {
	if err := New().Run(os.Args); err != nil {
		log.Fatal("[Error] ", err)
	}
}
