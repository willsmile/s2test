package main

import (
	"github.com/urfave/cli/v2"
)

// version of s2test
const version = "0.1.0"

func New() *cli.App {
	app := &cli.App{
		Name:    "s2test",
		Usage:   "A Simple Smoke Test Tool",
		Version: version,
		Authors: []*cli.Author{
			{
				Name:  "Wei Chen (willsmile)",
				Email: "willsmile.me@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Value:   "",
				Usage:   "Path of test plan to execute",
			},
		},
		Action: func(c *cli.Context) error {
			var (
				plan  Plan
				store Store
				err   error
			)

			path := c.String("path")
			err = LoadJSON(path, &plan)
			if err != nil {
				return err
			}

			err = LoadJSON(plan.TargetPath, &store)
			if err != nil {
				return err
			}

			err = report.Print()
			if err != nil {
				return err
			}

			return nil
		},
	}

	return app
}
