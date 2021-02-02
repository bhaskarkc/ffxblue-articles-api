package articles

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bhaskarkc/ffxblue-article-api/domain/articles"
	"github.com/bhaskarkc/ffxblue-article-api/services"
	"github.com/bhaskarkc/ffxblue-article-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	articleId, err := strconv.ParseInt(
		c.Param("article_id"), 10, 64,
	)
	if err != nil {
		err := errors.NewBadRequestError(
			fmt.Sprintf("non numeric article %d passed", articleId),
		)
		c.JSON(err.Status, err)
		return
	}

	article, getErr := services.ArticleService.Get(articleId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, article)
}

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
