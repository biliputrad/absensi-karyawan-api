package config

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(config Config) *gin.Engine {
	if config.GinMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	return router
}
