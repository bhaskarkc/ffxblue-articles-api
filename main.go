package main

import (
	"github.com/bhaskarkc/ffxblue-articles-api/logger"
	"github.com/bhaskarkc/ffxblue-articles-api/server"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env", ".env.example"); err != nil {
		logger.Error("ENV vars not loaded", err)
	}
}

func main() {
	server.Start()
}
