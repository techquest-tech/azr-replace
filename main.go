package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:        "Azr Replace",
		Usage:       "-",
		Description: "cli to run Azure DevOps replace locally",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "env",
				Aliases: []string{"e"},
				Usage:   "env settings",
			},
			&cli.StringSliceFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Value:   cli.NewStringSlice(".env"),
			},
			&cli.StringFlag{
				Name:     "source",
				Aliases:  []string{"s"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "output",
				Value:   "-",
				Aliases: []string{"o"},
				Usage:   "output to file or default - to STDOUT",
			},
			&cli.StringFlag{
				Name:    "log",
				Aliases: []string{"l"},
				EnvVars: []string{"LOG_LEVEL"},
				Value:   "info",
			},
			&cli.StringFlag{
				Name:  "start",
				Value: "#{",
			},
			&cli.StringFlag{
				Name:  "end",
				Value: "}#",
			},
		},
		Before: initApp,
		Action: fire,
	}
	err := app.Run(os.Args)
	if err != nil {
		logrus.WithField("err", err).Fatalf("task failed. %s", err)
	}
}

func initApp(c *cli.Context) error {
	log := c.String("log")
	level, _ := logrus.ParseLevel(log)
	logrus.SetLevel(level)
	logrus.SetOutput(os.Stdout)

	return nil
}
