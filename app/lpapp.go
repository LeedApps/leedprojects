package app

import (
	"github.com/gin-gonic/gin"
	"github.com/leedapps/leedprojects/app/routes"
)

// LPApp stores all application resources
type LPApp struct {
	Router *gin.Engine
}

// InitializeApp creates the base application
func InitializeApp() LPApp {
	lpapp := LPApp{
		Router: routes.LPRouter(),
	}
	return lpapp
}
