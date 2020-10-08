package tasks

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/leedapps/leedprojects/app"
	"github.com/leedapps/leedprojects/app/config"
	_ "github.com/lib/pq" // Loads postgres sql drivers

	"github.com/urfave/cli/v2"
)

// DBCommands returns db related commands for tasker
func DBCommands(lpapp app.LPApp) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "init",
			Usage:   "Create the database required for the environment",
			Aliases: []string{"i"},
			Action: func(c *cli.Context) error {
				return DbInit(lpapp)
			},
		},
	}
}

// DbConnectionString produces the connection string required for using the postges sql driver
func DbConnectionString(dbconfig config.LPConfig, includeDatabase bool) string {
	if includeDatabase {
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbconfig.Values["User"], dbconfig.Values["Password"], dbconfig.Values["Host"], dbconfig.Values["Port"], dbconfig.Values["Database"], dbconfig.Values["SslMode"])
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=%s", dbconfig.Values["User"], dbconfig.Values["Password"], dbconfig.Values["Host"], dbconfig.Values["Port"], dbconfig.Values["SslMode"])
}

// DbInit creates the database(s) required for the project
func DbInit(lpapp app.LPApp) error {
	connectionString := DbConnectionString(lpapp.DBConfig, false)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", lpapp.DBConfig.Values["Database"]))
	if err == nil {
		log.Println("Database created successfully")
	}
	return err
}
