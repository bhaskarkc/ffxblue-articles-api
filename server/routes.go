package server

import (
	"github.com/bhaskarkc/ffxblue-articles-api/controllers/articles"
	"github.com/bhaskarkc/ffxblue-articles-api/controllers/healthcheck"
	"github.com/bhaskarkc/ffxblue-articles-api/controllers/tags"
)

func registerRoutes() {
	httpServer.GET("/healthcheck", healthcheck.Healthcheck)
	httpServer.HEAD("/healthcheck", healthcheck.Healthcheck)
	httpServer.GET("/articles/:article_id", articles.GetArticle)
	httpServer.POST("/articles", articles.CreateArticle)
	httpServer.GET("/tags/:tag_name/:date", tags.GetTagByDate)
}
