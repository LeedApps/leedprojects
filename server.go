package main

import (
	"github.com/leedapps/leedprojects/app"
)

func main() {
	lpapp := app.InitializeApp()
	lpapp.Router.Run()
}
