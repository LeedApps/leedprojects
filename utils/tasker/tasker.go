package main

import (
	"log"
	"os"

	"github.com/leedapps/leedprojects/app"
	"github.com/leedapps/leedprojects/app/tasks"

	"github.com/urfave/cli/v2"
)

func main() {
	lpapp := app.InitializeApp(false)
	tasker := &cli.App{
		Name:     "tasker",
		Usage:    "Run project related tasks",
		Commands: tasks.AllTasks(lpapp),
	}
	err := tasker.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
