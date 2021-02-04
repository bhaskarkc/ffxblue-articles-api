package tags

import (
	"fmt"
	"net/http"

	"github.com/bhaskarkc/ffxblue-articles-api/services"
	"github.com/bhaskarkc/ffxblue-articles-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetTagByDate(c *gin.Context) {
	tagName := c.Param("tag_name")
	if tagName == "" {
		err := errors.NewBadRequestError(
			fmt.Sprint("invalid tagName"),
		)
		c.JSON(err.Status, err)
		return
	}

	dateString := c.Param("date")
	if dateString == "" {
		err := errors.NewBadRequestError(
			fmt.Sprint("invalid dateString"),
		)
		c.JSON(err.Status, err)
		return
	}

	tagByDate, err := services.TagService.GetTagByDate(tagName, dateString)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, tagByDate)
}
