package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/willsmile/s2test/lib"
)

// Version of s2test
const version = "0.1.0"

func main() {
	var path string

	app := &cli.App{
		Name:    "s2test",
		Usage:   "A Simple Smoke Test Tool",
		Version: version,
		Authors: []*cli.Author{
			{
				Name:  "willsmile",
				Email: "smile.v.chen@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "p",
				Value:       "",
				Usage:       "path of test plan to execute",
				Destination: &path,
			},
		},
		Action: func(c *cli.Context) error {
			plan := lib.LoadPlan(path)
			store := lib.LoadAPIStore(plan.TargetPath)
			lib.Execute(plan, store)
			return nil
		},
	}

	error := app.Run(os.Args)
	if error != nil {
		log.Fatal("[Cli Error] ", error)
	}
}
