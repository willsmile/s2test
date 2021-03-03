package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// version of s2test
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

			plan, err := LoadPlan(path)
			if err != nil {
				return err
			}

			store, err := LoadStore(plan.TargetPath)
			if err != nil {
				return err
			}

			report := plan.Execute(store)
			report.Print()

			return nil
		},
	}

	error := app.Run(os.Args)
	if error != nil {
		log.Fatal("[Cli Error] ", error)
	}
}
