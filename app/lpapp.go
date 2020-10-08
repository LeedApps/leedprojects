package app

import (
	"github.com/gin-gonic/gin"
	"github.com/leedapps/leedprojects/app/config"
	"github.com/leedapps/leedprojects/app/routes"
)

// LPApp stores all application resources
type LPApp struct {
	Router   *gin.Engine
	DBConfig config.LPConfig
}

// InitializeApp creates the base application
func InitializeApp(loadFullServer bool) LPApp {
	config.LoadEnv()
	lpapp := LPApp{
		DBConfig: config.DBConfig,
	}
	if loadFullServer {
		lpapp.Router = routes.LPRouter()
	}
	lpapp.DBConfig.Load()
	return lpapp
}
