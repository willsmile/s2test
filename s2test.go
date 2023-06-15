package main

import (
	"os"

	"github.com/willsmile/s2test/internal/cli"
)

func main() {
	if err := cli.New().Run(os.Args); err != nil {
		cli.Log(err)
	}
}
