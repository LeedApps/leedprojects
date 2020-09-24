package main

import (
	"log"
	"os"

	"github.com/leedapps/leedprojects/app/tasks"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "tasker",
		Usage:    "Run project related tasks",
		Commands: tasks.AllTasks(),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
