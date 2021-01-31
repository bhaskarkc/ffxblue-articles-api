package services

import (
	"github.com/bhaskarkc/ffxblue-article-api/domain/articles"
	"github.com/bhaskarkc/ffxblue-article-api/utils/date"
	"github.com/bhaskarkc/ffxblue-article-api/utils/errors"
)

var ArticleService articleServiceInterface = &articleService{}

type articleServiceInterface interface {
	Get(int64) (*articles.Article, *errors.RestErr)
	Create(articles.Article) (*articles.Article, *errors.RestErr)
	RelatedTags([]int64) ([]string, *errors.RestErr)
}

type articleService struct{}

func (s *articleService) Get(Id int64) (*articles.Article, *errors.RestErr) {
	article := &articles.Article{Id: Id}
	article.Date = date.GetNowString()
	if err := article.Get(); err != nil {
		return nil, err
	}
	return article, nil
}

func (s *articleService) Create(article articles.Article) (*articles.Article, *errors.RestErr) {
	if err := article.Validate(); err != nil {
		return nil, err
	}

	article.Date = date.GetNowDBFormat()
	if err := article.Save(); err != nil {
		return nil, err
	}
	// TODO: insert Tags and tags relation.
	return &article, nil
}

func (s *articleService) RelatedTags(Ids []int64) ([]string, *errors.RestErr) {
	return []string{}, nil
}
