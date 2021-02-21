package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
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
				Name:        "path",
				Aliases:     []string{"p"},
				Value:       "",
				Usage:       "path of test plan to execute",
				Destination: &path,
			},
		},
		Action: func(c *cli.Context) error {
			if c.String("path") == "" {
				path = c.Args().First()
			}

			plan := LoadPlan(path)
			store := LoadStore(plan.TargetPath)
			plan.Execute(store)
			return nil
		},
	}

	error := app.Run(os.Args)
	if error != nil {
		log.Fatal("[Cli Error] ", error)
	}
}
