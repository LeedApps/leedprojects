package tasks

import "github.com/urfave/cli/v2"

// AllTasks collects all tasks required for tasker
func AllTasks() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "db",
			Usage:       "Database related tasks",
			Subcommands: DBCommands(),
		},
	}
}
