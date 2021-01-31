package server

import (
	"github.com/bhaskarkc/ffxblue-article-api/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	httpServer := gin.Default()
	logger.Info("Starting http server.")
	httpServer.Run(":7001")
}
