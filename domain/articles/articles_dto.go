package articles

import (
	"strings"

	"github.com/bhaskarkc/ffxblue-article-api/utils/errors"
)

type Article struct {
	Id    int64    `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

type Articles []Article

func (article *Article) Validate() *errors.RestErr {
	article.Title = strings.TrimSpace(article.Title)
	article.Date = strings.TrimSpace(article.Date)
	article.Body = strings.TrimSpace(article.Body)
	return nil
}
