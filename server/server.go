package server

import (
	"github.com/bhaskarkc/ffxblue-articles-api/utils"
	"github.com/gin-gonic/gin"
)

var httpServer = gin.Default()

func Start() {
	registerRoutes()
	port := utils.JoinString(":", utils.GetEnv("SERVER_PORT", "8080"))
	httpServer.Run(port)
}
