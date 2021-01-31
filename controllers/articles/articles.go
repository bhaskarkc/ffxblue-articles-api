package controllers

import (
	"net/http"

	"github.com/bhaskarkc/ffxblue-article-api/domain/articles"
	"github.com/bhaskarkc/ffxblue-article-api/services"
	"github.com/bhaskarkc/ffxblue-article-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var article articles.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.ArticleService.Create(article)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}
