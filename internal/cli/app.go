package cli

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"

	"github.com/willsmile/s2test/internal/config"
	"github.com/willsmile/s2test/internal/executor"
	"github.com/willsmile/s2test/internal/reporter"
	"github.com/willsmile/s2test/internal/storage"
)

const (
	appName = "s2test"
	version = "0.6.0"
)

func New() *cli.App {
	app := &cli.App{
		Name:    appName,
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
				Name:    "plan",
				Aliases: []string{"p"},
				Value:   "",
				Usage:   "Path of plan to execute",
			},
			&cli.StringFlag{
				Name:    "api",
				Aliases: []string{"a"},
				Value:   "",
				Usage:   "Path of API endpoints",
			},
			&cli.StringFlag{
				Name:    "mode",
				Aliases: []string{"m"},
				Value:   "normal",
				Usage:   "Select print mode",
			},
		},
		Action: func(c *cli.Context) error {
			var (
				plan      executor.Plan
				endpoints storage.Endpoints
			)

			planPath := c.String("plan")
			if err := config.LoadJSON(planPath, &plan); err != nil {
				return err
			}

			apiPath := plan.GetEndpointsPath(c.String("api"))
			if err := config.LoadJSON(apiPath, &endpoints); err != nil {
				return err
			}

			reports, err := plan.Execute(&endpoints, appInfo())
			if err != nil {
				return err
			}

			printMode, err := reporter.NewPrintMode(c.String("mode"))
			if err != nil {
				return err
			}
			if err := reports.Print(printMode); err != nil {
				return err
			}

			return nil
		},
	}

	return app
}

func Log(err error) {
	switch err {
	case executor.ErrNoTasksToExecute, reporter.ErrEmptyReport:
		log.Print("[INFO] ", err)
	default:
		log.Fatal("[ERROR] ", err)
	}
}

func appInfo() string {
	return fmt.Sprintf("%s %s", appName, version)
}
