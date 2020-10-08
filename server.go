package main

import (
	"github.com/leedapps/leedprojects/app"
	"github.com/leedapps/leedprojects/app/config"
)

func main() {
	lpapp := app.InitializeApp(true)
	lpconfig := config.DBConfig
	lpconfig.Load()
	lpapp.Router.Run()
}
