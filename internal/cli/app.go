package cli

import (
	"log"

	"github.com/urfave/cli/v2"

	"github.com/willsmile/s2test/internal/config"
	"github.com/willsmile/s2test/internal/executor"
)

const (
	Name    = "s2test"
	version = "0.3.0"
)

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
				plan  executor.Plan
				store executor.Endpoints
				err   error
			)

			path := c.String("path")
			err = config.LoadJSON(path, &plan)
			if err != nil {
				return err
			}

			err = config.LoadJSON(plan.TargetPath, &store)
			if err != nil {
				return err
			}

			report, err := plan.Execute(&store)
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

func Log(err error) {
	switch err {
	case executor.ErrNoTasksToExecute, executor.ErrEmptyReport:
		log.Print("[INFO] ", err)
	default:
		log.Fatal("[ERROR] ", err)
	}
}
