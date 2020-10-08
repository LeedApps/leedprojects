package tasks

import (
	"github.com/leedapps/leedprojects/app"
	"github.com/urfave/cli/v2"
)

// AllTasks collects all tasks required for tasker
func AllTasks(lpapp app.LPApp) []*cli.Command {
	return []*cli.Command{
		{
			Name:        "db",
			Usage:       "Database related tasks",
			Subcommands: DBCommands(lpapp),
		},
	}
}
